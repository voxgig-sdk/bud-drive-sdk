-- DriveMcpApi entity test

local json = require("dkjson")
local vs = require("utility.struct.struct")
local sdk = require("bud-drive_sdk")
local helpers = require("core.helpers")
local runner = require("test.runner")

local _test_dir = debug.getinfo(1, "S").source:match("^@(.+/)")  or "./"

describe("DriveMcpApiEntity", function()
  it("should create instance", function()
    local testsdk = sdk.test(nil, nil)
    local ent = testsdk:DriveMcpApi(nil)
    assert.is_not_nil(ent)
  end)

  it("should run basic flow", function()
    local setup = drive_mcp_api_basic_setup(nil)
    -- Per-op sdk-test-control.json skip.
    local _live = setup.live or false
    for _, _op in ipairs({"create"}) do
      local _should_skip, _reason = runner.is_control_skipped("entityOp", "drive_mcp_api." .. _op, _live and "live" or "unit")
      if _should_skip then
        pending(_reason or "skipped via sdk-test-control.json")
        return
      end
    end
    -- The basic flow consumes synthetic IDs from the fixture. In live mode
    -- without an *_ENTID env override, those IDs hit the live API and 4xx.
    if setup.synthetic_only then
      pending("live entity test uses synthetic IDs from fixture — set BUD_DRIVE_TEST_DRIVE_MCP_API_ENTID JSON to run live")
      return
    end
    local client = setup.client

    -- CREATE
    local drive_mcp_api_ref01_ent = client:DriveMcpApi(nil)
    local drive_mcp_api_ref01_data = helpers.to_map(vs.getprop(
      vs.getpath(setup.data, "new.drive_mcp_api"), "drive_mcp_api_ref01"))

    local drive_mcp_api_ref01_data_result, err = drive_mcp_api_ref01_ent:create(drive_mcp_api_ref01_data, nil)
    assert.is_nil(err)
    drive_mcp_api_ref01_data = helpers.to_map(type(drive_mcp_api_ref01_data_result) == 'table' and drive_mcp_api_ref01_data_result.data_get and drive_mcp_api_ref01_data_result:data_get() or drive_mcp_api_ref01_data_result)
    assert.is_not_nil(drive_mcp_api_ref01_data)
    assert.is_not_nil(drive_mcp_api_ref01_data["id"])

  end)
end)

function drive_mcp_api_basic_setup(extra)
  runner.load_env_local()

  local entity_data_file = _test_dir .. "../../.sdk/test/entity/drive_mcp_api/DriveMcpApiTestData.json"
  local f = io.open(entity_data_file, "r")
  if f == nil then
    error("failed to read drive_mcp_api test data: " .. entity_data_file)
  end
  local entity_data_source = f:read("*a")
  f:close()

  local entity_data = json.decode(entity_data_source)

  local options = {}
  options["entity"] = entity_data["existing"]

  local client = sdk.test(options, extra)

  -- Generate idmap via transform.
  local idmap = vs.transform(
    { "drive_mcp_api01", "drive_mcp_api02", "drive_mcp_api03" },
    {
      ["`$PACK`"] = { "", {
        ["`$KEY`"] = "`$COPY`",
        ["`$VAL`"] = { "`$FORMAT`", "upper", "`$COPY`" },
      }},
    }
  )

  -- Detect ENTID env override before envOverride consumes it. When live
  -- mode is on without a real override, the basic test runs against synthetic
  -- IDs from the fixture and 4xx's. Surface this so the test can skip.
  local entid_env_raw = os.getenv("BUD_DRIVE_TEST_DRIVE_MCP_API_ENTID")
  local idmap_overridden = entid_env_raw ~= nil and entid_env_raw:match("^%s*{") ~= nil

  local env = runner.env_override({
    ["BUD_DRIVE_TEST_DRIVE_MCP_API_ENTID"] = idmap,
    ["BUD_DRIVE_TEST_LIVE"] = "FALSE",
    ["BUD_DRIVE_TEST_EXPLAIN"] = "FALSE",
    ["BUD_DRIVE_APIKEY"] = "",
  })

  local idmap_resolved = helpers.to_map(
    env["BUD_DRIVE_TEST_DRIVE_MCP_API_ENTID"])
  if idmap_resolved == nil then
    idmap_resolved = helpers.to_map(idmap)
  end

  if env["BUD_DRIVE_TEST_LIVE"] == "TRUE" then
    local merged_opts = vs.merge({
      -- FIRST, so the generated fields below win: sdk-test-control.json's
      -- test.client.options adds to the live client, it does not redirect it.
      runner.live_client_options(),
      {
        apikey = env["BUD_DRIVE_APIKEY"],
      },
      extra or {},
    })
    client = sdk.new(helpers.to_map(merged_opts))
  end

  local live = env["BUD_DRIVE_TEST_LIVE"] == "TRUE"
  return {
    client = client,
    data = entity_data,
    idmap = idmap_resolved,
    env = env,
    explain = env["BUD_DRIVE_TEST_EXPLAIN"] == "TRUE",
    live = live,
    synthetic_only = live and not idmap_overridden,
    now = os.time() * 1000,
  }
end
