# BudDrive SDK feature factory

from buddrive_sdk.feature.base_feature import BudDriveBaseFeature
from buddrive_sdk.feature.debug_feature import BudDriveDebugFeature
from buddrive_sdk.feature.idempotency_feature import BudDriveIdempotencyFeature
from buddrive_sdk.feature.metrics_feature import BudDriveMetricsFeature
from buddrive_sdk.feature.paging_feature import BudDrivePagingFeature
from buddrive_sdk.feature.ratelimit_feature import BudDriveRatelimitFeature
from buddrive_sdk.feature.retry_feature import BudDriveRetryFeature
from buddrive_sdk.feature.test_feature import BudDriveTestFeature
from buddrive_sdk.feature.timeout_feature import BudDriveTimeoutFeature


_FEATURES = {
    "base": lambda: BudDriveBaseFeature(),
    "debug": lambda: BudDriveDebugFeature(),
    "idempotency": lambda: BudDriveIdempotencyFeature(),
    "metrics": lambda: BudDriveMetricsFeature(),
    "paging": lambda: BudDrivePagingFeature(),
    "ratelimit": lambda: BudDriveRatelimitFeature(),
    "retry": lambda: BudDriveRetryFeature(),
    "test": lambda: BudDriveTestFeature(),
    "timeout": lambda: BudDriveTimeoutFeature(),
}


def _make_feature(name):
    factory = _FEATURES.get(name)
    if factory is not None:
        return factory()
    return _FEATURES["base"]()


# True when this SDK was generated with the named feature class - the
# constructor's tolerance for extend-carried features reads this (an
# active name with no generated class must not become a BaseFeature
# stray when an extend instance carries it).
def _has_feature(name):
    return name in _FEATURES
