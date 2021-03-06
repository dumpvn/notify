from app.notify_client import NotifyAdminAPIClient


class PlatformStatsAPIClient(NotifyAdminAPIClient):
    # Fudge assert in the super __init__ so
    # we can set those variables later.
    def __init__(self):
        super().__init__("a" * 73, "b")


    def get_aggregate_platform_stats(self, params_dict=None):
        return self.get("/platform-stats", params=params_dict)

    def get_billing_for_all_services(self, params_dict=None):
        return self.get("/platform-stats/billing-for-all-services", params=params_dict)

    def get_usage_for_all_services(self, params_dict=None):
        return self.get("/platform-stats/usage-for-all-services", params=params_dict)


platform_stats_api_client = PlatformStatsAPIClient()
