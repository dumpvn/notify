from datetime import datetime
import json

from flask import Blueprint, jsonify, request

from app.dao.monthly_billing_dao import (
    get_billing_data_for_financial_year,
    get_monthly_billing_by_notification_type
)
from app.dao.date_util import get_months_for_financial_year
from app.errors import register_errors
from app.models import SMS_TYPE, EMAIL_TYPE, LETTER_TYPE
from app.utils import convert_utc_to_aet
from app.dao.annual_billing_dao import (dao_get_free_sms_fragment_limit_for_year,
                                        dao_get_all_free_sms_fragment_limit,
                                        dao_create_or_update_annual_billing_for_year,
                                        dao_update_annual_billing_for_future_years)
from app.errors import InvalidRequest
from app.schema_validation import validate
from app.dao.date_util import get_current_financial_year_start_year
from app.billing.billing_schemas import create_or_update_free_sms_fragment_limit_schema

billing_blueprint = Blueprint(
    'billing',
    __name__,
    url_prefix='/service/<uuid:service_id>/billing'
)

register_errors(billing_blueprint)


@billing_blueprint.route('/monthly-usage')
def get_yearly_usage_by_month(service_id):
    try:
        year = int(request.args.get('year'))
        results = []
        for month in get_months_for_financial_year(year):
            billing_for_month = get_monthly_billing_by_notification_type(service_id, month, SMS_TYPE)
            if billing_for_month:
                results.append(_transform_billing_for_month_sms(billing_for_month))
            letter_billing_for_month = get_monthly_billing_by_notification_type(service_id, month, LETTER_TYPE)
            if letter_billing_for_month:
                results.extend(_transform_billing_for_month_letters(letter_billing_for_month))
        return json.dumps(results)

    except TypeError:
        return jsonify(result='error', message='No valid year provided'), 400


@billing_blueprint.route('/yearly-usage-summary')
def get_yearly_billing_usage_summary(service_id):
    try:
        year = int(request.args.get('year'))
        billing_data = get_billing_data_for_financial_year(service_id, year)
        notification_types = [SMS_TYPE, EMAIL_TYPE, LETTER_TYPE]
        response = [
            _get_total_billable_units_and_rate_for_notification_type(billing_data, notification_type)
            for notification_type in notification_types
        ]

        return json.dumps(response)

    except TypeError:
        return jsonify(result='error', message='No valid year provided'), 400


def _get_total_billable_units_and_rate_for_notification_type(billing_data, noti_type):
    total_sent = 0
    rate = 0
    letter_total = 0
    for entry in billing_data:
        for monthly_total in entry.monthly_totals:
            if entry.notification_type == noti_type:
                if entry.notification_type == EMAIL_TYPE:
                    total_sent += monthly_total['billing_units']
                    rate = monthly_total['rate']
                elif entry.notification_type == SMS_TYPE:
                    total_sent += (monthly_total['billing_units'] * monthly_total['rate_multiplier'])
                    rate = monthly_total['rate']
                elif entry.notification_type == LETTER_TYPE:
                    total_sent += monthly_total['billing_units']
                    letter_total += (monthly_total['billing_units'] * monthly_total['rate'])

    return {
        "notification_type": noti_type,
        "billing_units": total_sent,
        "rate": rate,
        "letter_total": letter_total
    }


def _transform_billing_for_month_sms(billing_for_month):
    month_name = datetime.strftime(convert_utc_to_aet(billing_for_month.start_date), "%B")
    billing_units = rate = 0

    for total in billing_for_month.monthly_totals:
        billing_units += (total['billing_units'] * total['rate_multiplier'])
        rate = total['rate']

    return {
        "month": month_name,
        "billing_units": billing_units,
        "notification_type": billing_for_month.notification_type,
        "rate": rate
    }


def _transform_billing_for_month_letters(billing_for_month):
    month_name = datetime.strftime(convert_utc_to_aet(billing_for_month.start_date), "%B")
    x = list()

    for total in billing_for_month.monthly_totals:
        y = {
            "month": month_name,
            "billing_units": (total['billing_units'] * total['rate_multiplier']),
            "notification_type": billing_for_month.notification_type,
            "rate": total['rate']
        }
        x.append(y)
    if len(billing_for_month.monthly_totals) == 0:
        x.append({
            "month": month_name,
            "billing_units": 0,
            "notification_type": billing_for_month.notification_type,
            "rate": 0
        })
    return x


@billing_blueprint.route('/free-sms-fragment-limit', methods=["GET"])
def get_free_sms_fragment_limit(service_id):

    financial_year_start = request.args.get('financial_year_start')

    annual_billing = dao_get_free_sms_fragment_limit_for_year(service_id, financial_year_start)

    if annual_billing is None:
        # An entry does not exist in annual_billing table for that service and year. If it is a past year,
        # we return the oldest entry.
        # If it is the current or future years, we create an entry in the db table using the newest record,
        # and return that number.  If all fails, we return InvalidRequest.
        sms_list = dao_get_all_free_sms_fragment_limit(service_id)

        if not sms_list:
            raise InvalidRequest('no free-sms-fragment-limit entry for service {} in DB'.format(service_id), 404)
        else:
            if financial_year_start is None:
                financial_year_start = get_current_financial_year_start_year()

            if int(financial_year_start) < get_current_financial_year_start_year():
                # return the earliest historical entry
                annual_billing = sms_list[0]   # The oldest entry
            else:
                annual_billing = sms_list[-1]  # The newest entry

                annual_billing = dao_create_or_update_annual_billing_for_year(service_id,
                                                                              annual_billing.free_sms_fragment_limit,
                                                                              financial_year_start)

    return jsonify(annual_billing.serialize_free_sms_items()), 200


@billing_blueprint.route('/free-sms-fragment-limit', methods=["POST"])
def create_or_update_free_sms_fragment_limit(service_id):

    req_args = request.get_json()

    form = validate(req_args, create_or_update_free_sms_fragment_limit_schema)

    update_free_sms_fragment_limit_data(service_id,
                                        free_sms_fragment_limit=form.get('free_sms_fragment_limit'),
                                        financial_year_start=form.get('financial_year_start'))
    return jsonify(form), 201


def update_free_sms_fragment_limit_data(service_id, free_sms_fragment_limit, financial_year_start):
    current_year = get_current_financial_year_start_year()
    if not financial_year_start:
        financial_year_start = current_year

    dao_create_or_update_annual_billing_for_year(
        service_id,
        free_sms_fragment_limit,
        financial_year_start
    )
    # if we're trying to update historical data, don't touch other rows.
    # Otherwise, make sure that future years will get the new updated value.
    if financial_year_start >= current_year:
        dao_update_annual_billing_for_future_years(
            service_id,
            free_sms_fragment_limit,
            financial_year_start
        )
