import uuid
from datetime import datetime, timedelta
from decimal import Decimal
import functools

import flask
from flask import current_app
import click
from click_datetime import Datetime as click_dt

from app import db
from app.dao.monthly_billing_dao import (
    create_or_update_monthly_billing,
    get_monthly_billing_by_notification_type,
    get_service_ids_that_need_billing_populated
)
from app.models import PROVIDERS, User, SMS_TYPE, EMAIL_TYPE
from app.dao.services_dao import (
    delete_service_and_all_associated_db_objects,
    dao_fetch_all_services_by_user
)
from app.dao.provider_rates_dao import create_provider_rates as dao_create_provider_rates
from app.dao.users_dao import (delete_model_user, delete_user_verify_codes)
from app.utils import get_midnight_for_day_before, get_london_midnight_in_utc
from app.performance_platform.processing_time import (send_processing_time_for_start_and_end)
from app.celery.scheduled_tasks import send_total_sent_notifications_to_performance_platform


@click.group(name='command', help='Additional commands')
def command_group():
    pass


class notify_command:
    def __init__(self, name=None):
        self.name = name

    def __call__(self, func):
        # we need to call the flask with_appcontext decorator to ensure the config is loaded, db connected etc etc.
        # we also need to use functools.wraps to carry through the names and docstrings etc of the functions.
        # Then we need to turn it into a click.Command - that's what command_group.add_command expects.
        @click.command(name=self.name)
        @functools.wraps(func)
        @flask.cli.with_appcontext
        def wrapper(*args, **kwargs):
            return func(*args, **kwargs)

        command_group.add_command(wrapper)

        return wrapper


@notify_command()
@click.option('-p', '--provider_name', required=True, type=click.Choice(PROVIDERS))
@click.option('-c', '--cost', required=True, help='Cost (pence) per message including decimals', type=float)
@click.option('-d', '--valid_from', required=True, type=click_dt(format='%Y-%m-%dT%H:%M:%S'))
def create_provider_rates(provider_name, cost, valid_from):
    """
    Backfill rates for a given provider
    """
    cost = Decimal(cost)
    dao_create_provider_rates(provider_name, valid_from, cost)


@notify_command()
@click.option('-u', '--user_email_prefix', required=True, help="""
    Functional test user email prefix. eg "notify-test-preview"
""")  # noqa
def purge_functional_test_data(user_email_prefix):
    """
    Remove non-seeded functional test data

    users, services, etc. Give an email prefix. Probably "notify-test-preview".
    """
    users = User.query.filter(User.email_address.like("{}%".format(user_email_prefix))).all()
    for usr in users:
        # Make sure the full email includes a uuid in it
        # Just in case someone decides to use a similar email address.
        try:
            uuid.UUID(usr.email_address.split("@")[0].split('+')[1])
        except ValueError:
            print("Skipping {} as the user email doesn't contain a UUID.".format(usr.email_address))
        else:
            services = dao_fetch_all_services_by_user(usr.id)
            if services:
                for service in services:
                    delete_service_and_all_associated_db_objects(service)
            else:
                delete_user_verify_codes(usr)
                delete_model_user(usr)


@notify_command()
def backfill_notification_statuses():
    """
    DEPRECATED. Populates notification_status.

    This will be used to populate the new `Notification._status_fkey` with the old
    `Notification._status_enum`
    """
    LIMIT = 250000
    subq = "SELECT id FROM notification_history WHERE notification_status is NULL LIMIT {}".format(LIMIT)
    update = "UPDATE notification_history SET notification_status = status WHERE id in ({})".format(subq)
    result = db.session.execute(subq).fetchall()

    while len(result) > 0:
        db.session.execute(update)
        print('commit {} updates at {}'.format(LIMIT, datetime.utcnow()))
        db.session.commit()
        result = db.session.execute(subq).fetchall()


@notify_command()
def update_notification_international_flag():
    """
    DEPRECATED. Set notifications.international=false.
    """
    # 250,000 rows takes 30 seconds to update.
    subq = "select id from notifications where international is null limit 250000"
    update = "update notifications set international = False where id in ({})".format(subq)
    result = db.session.execute(subq).fetchall()

    while len(result) > 0:
        db.session.execute(update)
        print('commit 250000 updates at {}'.format(datetime.utcnow()))
        db.session.commit()
        result = db.session.execute(subq).fetchall()

    # Now update notification_history
    subq_history = "select id from notification_history where international is null limit 250000"
    update_history = "update notification_history set international = False where id in ({})".format(subq_history)
    result_history = db.session.execute(subq_history).fetchall()
    while len(result_history) > 0:
        db.session.execute(update_history)
        print('commit 250000 updates at {}'.format(datetime.utcnow()))
        db.session.commit()
        result_history = db.session.execute(subq_history).fetchall()


@notify_command()
def fix_notification_statuses_not_in_sync():
    """
    DEPRECATED.
    This will be used to correct an issue where Notification._status_enum and NotificationHistory._status_fkey
    became out of sync. See 979e90a.

    Notification._status_enum is the source of truth so NotificationHistory._status_fkey will be updated with
    these values.
    """
    MAX = 10000

    subq = "SELECT id FROM notifications WHERE cast (status as text) != notification_status LIMIT {}".format(MAX)
    update = "UPDATE notifications SET notification_status = status WHERE id in ({})".format(subq)
    result = db.session.execute(subq).fetchall()

    while len(result) > 0:
        db.session.execute(update)
        print('Committed {} updates at {}'.format(len(result), datetime.utcnow()))
        db.session.commit()
        result = db.session.execute(subq).fetchall()

    subq_hist = "SELECT id FROM notification_history WHERE cast (status as text) != notification_status LIMIT {}" \
        .format(MAX)
    update = "UPDATE notification_history SET notification_status = status WHERE id in ({})".format(subq_hist)
    result = db.session.execute(subq_hist).fetchall()

    while len(result) > 0:
        db.session.execute(update)
        print('Committed {} updates at {}'.format(len(result), datetime.utcnow()))
        db.session.commit()
        result = db.session.execute(subq_hist).fetchall()


@notify_command()
@click.option('-y', '--year', required=True, help="e.g. 2017", type=int)
def populate_monthly_billing(year):
    """
    Populate monthly billing table for all services for a given year.
    """
    def populate(service_id, year, month):
        create_or_update_monthly_billing(service_id, datetime(year, month, 1))
        sms_res = get_monthly_billing_by_notification_type(
            service_id, datetime(year, month, 1), SMS_TYPE
        )
        email_res = get_monthly_billing_by_notification_type(
            service_id, datetime(year, month, 1), EMAIL_TYPE
        )
        letter_res = get_monthly_billing_by_notification_type(
            service_id, datetime(year, month, 1), 'letter'
        )

        print("Finished populating data for {} for service id {}".format(month, str(service_id)))
        print('SMS: {}'.format(sms_res.monthly_totals))
        print('Email: {}'.format(email_res.monthly_totals))
        print('Letter: {}'.format(letter_res.monthly_totals))

    service_ids = get_service_ids_that_need_billing_populated(
        start_date=datetime(2016, 5, 1), end_date=datetime(2017, 8, 16)
    )
    start, end = 1, 13

    if year == 2016:
        start = 4

    for service_id in service_ids:
        print('Starting to populate data for service {}'.format(str(service_id)))
        print('Starting populating monthly billing for {}'.format(year))
        for i in range(start, end):
            print('Population for {}-{}'.format(i, year))
            populate(service_id, year, i)


@notify_command()
@click.option('-s', '--start_date', required=True, help="start date inclusive", type=click_dt(format='%Y-%m-%d'))
@click.option('-e', '--end_date', required=True, help="end date inclusive", type=click_dt(format='%Y-%m-%d'))
def backfill_performance_platform_totals(start_date, end_date):
    """
    Send historical total messages sent to Performance Platform.

    WARNING: This does not overwrite existing data. You need to delete
             the existing data or Performance Platform will double-count.
    """

    delta = end_date - start_date

    print('Sending total messages sent for all days between {} and {}'.format(start_date, end_date))

    for i in range(delta.days + 1):

        process_date = start_date + timedelta(days=i)

        print('Sending total messages sent for {}'.format(
            process_date.isoformat()
        ))

        send_total_sent_notifications_to_performance_platform(process_date)


@notify_command()
@click.option('-s', '--start_date', required=True, help="start date inclusive", type=click_dt(format='%Y-%m-%d'))
@click.option('-e', '--end_date', required=True, help="end date inclusive", type=click_dt(format='%Y-%m-%d'))
def backfill_processing_time(start_date, end_date):
    """
    Send historical processing time to Performance Platform.
    """

    delta = end_date - start_date

    print('Sending notification processing-time data for all days between {} and {}'.format(start_date, end_date))

    for i in range(delta.days + 1):
        # because the tz conversion funcs talk about midnight, and the midnight before last,
        # we want to pretend we're running this from the next morning, so add one.
        process_date = start_date + timedelta(days=i + 1)

        process_start_date = get_midnight_for_day_before(process_date)
        process_end_date = get_london_midnight_in_utc(process_date)

        print('Sending notification processing-time for {} - {}'.format(
            process_start_date.isoformat(),
            process_end_date.isoformat()
        ))
        send_processing_time_for_start_and_end(process_start_date, process_end_date)


@notify_command()
def populate_annual_billing():
    """
    add annual_billing for 2016, 2017 and 2018.
    """
    financial_year = [2016, 2017, 2018]

    for fy in financial_year:
        populate_data = """
        INSERT INTO annual_billing(id, service_id, free_sms_fragment_limit, financial_year_start,
                created_at, updated_at)
            SELECT uuid_in(md5(random()::text || now()::text)::cstring), id, 250000, {}, '{}', '{}'
            FROM services
            WHERE id NOT IN(
                SELECT service_id
                FROM annual_billing
                WHERE financial_year_start={})
        """.format(fy, datetime.utcnow(), datetime.utcnow(), fy)

        services_result1 = db.session.execute(populate_data)
        db.session.commit()

        print("Populated annual billing {} for {} services".format(fy, services_result1.rowcount))


@notify_command(name='list-routes')
def list_routes():
    """List URLs of all application routes."""
    for rule in sorted(current_app.url_map.iter_rules(), key=lambda r: r.rule):
        print("{:10} {}".format(", ".join(rule.methods - set(['OPTIONS', 'HEAD'])), rule.rule))


@notify_command(name='insert-inbound-numbers')
@click.option('-f', '--file_name', required=True,
              help="""Full path of the file to upload, file is a contains inbound numbers,
              one number per line. The number must have the format of 07... not 447....""")
def insert_inbound_numbers_from_file(file_name):
    print("Inserting inbound numbers from {}".format(file_name))
    file = open(file_name)
    sql = "insert into inbound_numbers values('{}', '{}', 'mmg', null, True, now(), null);"

    for line in file:
        print(line)
        db.session.execute(sql.format(uuid.uuid4(), line.strip()))
        db.session.commit()
    file.close()


def setup_commands(application):
    application.cli.add_command(command_group)
