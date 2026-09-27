package sdktest

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	sdk "github.com/voxgig-sdk/bud-drive-sdk/go"
	"github.com/voxgig-sdk/bud-drive-sdk/go/core"

	vs "github.com/voxgig-sdk/bud-drive-sdk/go/utility/struct"
)

func TestDriveMcpApiEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.DriveMcpApi(nil)
		if ent == nil {
			t.Fatal("expected non-nil DriveMcpApiEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := drive_mcp_apiBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "drive_mcp_api." + _op, _mode); _shouldSkip {
				if _reason == "" {
					_reason = "skipped via sdk-test-control.json"
				}
				t.Skip(_reason)
				return
			}
		}
		// The basic flow consumes synthetic IDs from the fixture. In live mode
		// without an *_ENTID env override, those IDs hit the live API and 4xx.
		if setup.syntheticOnly {
			t.Skip("live entity test uses synthetic IDs from fixture — set BUD_DRIVE_TEST_DRIVE_MCP_API_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		driveMcpApiRef01Ent := client.DriveMcpApi(nil)
		driveMcpApiRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath(setup.data, []any{"new", "drive_mcp_api"}), "drive_mcp_api_ref01"))

		driveMcpApiRef01DataResult, err := driveMcpApiRef01Ent.Create(driveMcpApiRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		driveMcpApiRef01Data = core.ToMapAny(entityData(driveMcpApiRef01DataResult))
		if driveMcpApiRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}
		if driveMcpApiRef01Data["id"] == nil {
			t.Fatal("expected created entity to have an id")
		}

	})
}

func drive_mcp_apiBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "drive_mcp_api", "DriveMcpApiTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read drive_mcp_api test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse drive_mcp_api test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap, _ := vs.Transform(
		[]any{"drive_mcp_api01", "drive_mcp_api02", "drive_mcp_api03"},
		map[string]any{
			"`$PACK`": []any{"", map[string]any{
				"`$KEY`": "`$COPY`",
				"`$VAL`": []any{"`$FORMAT`", "upper", "`$COPY`"},
			}},
		},
	)

	// Detect ENTID env override before envOverride consumes it. When live
	// mode is on without a real override, the basic test runs against synthetic
	// IDs from the fixture and 4xx's. Surface this so the test can skip.
	entidEnvRaw := os.Getenv("BUD_DRIVE_TEST_DRIVE_MCP_API_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"BUD_DRIVE_TEST_DRIVE_MCP_API_ENTID": idmap,
		"BUD_DRIVE_TEST_LIVE":      "FALSE",
		"BUD_DRIVE_TEST_EXPLAIN":   "FALSE",
		"BUD_DRIVE_APIKEY":         "",
	})

	idmapResolved := core.ToMapAny(env["BUD_DRIVE_TEST_DRIVE_MCP_API_ENTID"])
	if idmapResolved == nil {
		idmapResolved = core.ToMapAny(idmap)
	}

	if env["BUD_DRIVE_TEST_LIVE"] == "TRUE" {
		// An empty map, not a nil one: Merge returns nil when its last entry
		// is nil, and BasicSetup is normally called with no extras - so a
		// bare nil silently discarded the apikey and server values below.
		extraOpts := extra
		if extraOpts == nil {
			extraOpts = map[string]any{}
		}

		mergedOpts := vs.Merge([]any{
			// liveClientOptions() FIRST, so the generated fields below win:
			// sdk-test-control.json's test.client.options adds to the live
			// client, it does not redirect it.
			liveClientOptions(),
			map[string]any{
				"apikey": env["BUD_DRIVE_APIKEY"],
			},
			extraOpts,
		})
		client = sdk.NewBudDriveSDK(core.ToMapAny(mergedOpts))
	}

	live := env["BUD_DRIVE_TEST_LIVE"] == "TRUE"
	return &entityTestSetup{
		client:        client,
		data:          entityData,
		idmap:         idmapResolved,
		env:           env,
		explain:       env["BUD_DRIVE_TEST_EXPLAIN"] == "TRUE",
		live:          live,
		syntheticOnly: live && !idmapOverridden,
		now:           time.Now().UnixMilli(),
	}
}
