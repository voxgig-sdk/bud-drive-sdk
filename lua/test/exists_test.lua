-- BudDrive SDK exists test

local sdk = require("bud-drive_sdk")

describe("BudDriveSDK", function()
  it("should create test SDK", function()
    local testsdk = sdk.test(nil, nil)
    assert.is_not_nil(testsdk)
  end)
end)
