# BudDrive SDK exists test

import pytest
from buddrive_sdk import BudDriveSDK


class TestExists:

    def test_should_create_test_sdk(self):
        testsdk = BudDriveSDK.test(None, None)
        assert testsdk is not None
