from setuptools import setup

packages = [
    'notifications_utils',
    'notifications_utils.clients',
    'notifications_utils.clients.deskpro',
    'notifications_utils.clients.redis',
    'notifications_utils.clients.statsd'
]

install_requires =  [
    'bleach>=3.1,<4.0',
    'boto3>=1.13,<2.0',
    'botocore>=1.16,<2.0',
    'certifi>=2020.4,<2021.0',
    'chardet>=3.0,<4.0',
    'click>=7.1,<8.0',
    'flask-redis>=0.4.0,<0.5.0',
    'flask>=1.1,<2.0',
    'idna>=2.9,<3.0',
    'itsdangerous>=1.1,<2.0',
    'jinja2>=2.11,<3.0',
    'jmespath>=0.9.5,<0.10.0',
    'markupsafe>=1.1,<2.0',
    'mistune>=0.8.4,<0.9.0',
    'monotonic>=1.5,<2.0',
    'orderedset>=2.0,<3.0',
    'phonenumbers>=8.12,<9.0',
    'poetry>=1.0.5,<2.0.0',
    'pypdf2>=1.26,<2.0',
    'python-dateutil>=2.8,<3.0',
    'python-json-logger>=0.1.11,<0.2.0',
    'pytz>=2020.1,<2021.0',
    'pyyaml>=5.3,<6.0',
    'redis>=3.5,<4.0',
    'requests>=2.23,<3.0',
    's3transfer>=0.3.3,<0.4.0',
    'six>=1.14,<2.0',
    'smartypants>=2.0,<3.0',
    'statsd>=3.3,<4.0',
    'urllib3>=1.25,<2.0',
    'webencodings>=0.5.1,<0.6.0',
    'werkzeug>=1.0,<2.0'
]

setup_requires = [
    'pytest-runner'
]

tests_require = [
    "pytest",
    "pytest-mock",
    "pytest-cov",
    "requests-mock",
    "freezegun",
    "flake8",
]

setup_kwargs = {
    'name': 'notifications-utils',
    'version': '26.2.1',
    'description': 'Shared python code for Notify',
    'long_description': None,
    'author': 'DTA',
    'author_email': 'notify-dev@dta.gov.au',
    'maintainer': None,
    'maintainer_email': None,
    'url': None,
    'packages': packages,
    'include_package_data': True,
    'install_requires': install_requires,
    'setup_requires': setup_requires,
    'tests_require': tests_require,
    'python_requires': '>=3.7,<4.0',
}

setup(**setup_kwargs)
