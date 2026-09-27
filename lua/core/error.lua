-- BudDrive SDK error

local BudDriveError = {}
BudDriveError.__index = BudDriveError


function BudDriveError.new(code, msg, ctx)
  local self = setmetatable({}, BudDriveError)
  self.is_sdk_error = true
  self.sdk = "BudDrive"
  self.code = code or ""
  self.msg = msg or ""
  self.ctx = ctx
  self.result = nil
  self.spec = nil
  return self
end


function BudDriveError:error()
  return self.msg
end


function BudDriveError:__tostring()
  return self.msg
end


return BudDriveError
