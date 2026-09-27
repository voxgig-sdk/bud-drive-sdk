package sdktest

import (
	"encoding/json"
	"fmt"
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

func TestDriveCampaignsApiEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.DriveCampaignsApi(nil)
		if ent == nil {
			t.Fatal("expected non-nil DriveCampaignsApiEntity")
		}
	})

	// Feature #4: the entity Stream(action, ...) method runs the op pipeline and
	// returns a channel over result items. With the streaming feature active it
	// yields the feature's incremental output; otherwise it falls back to the
	// materialised list so Stream always yields.
	t.Run("stream", func(t *testing.T) {
		seed := map[string]any{
			"entity": map[string]any{
				"drive_campaigns_api": map[string]any{
					"s1": map[string]any{"id": "s1"},
					"s2": map[string]any{"id": "s2"},
					"s3": map[string]any{"id": "s3"},
				},
			},
		}

		// Fallback: streaming inactive -> yields the materialised list items.
		base := sdk.TestSDK(seed, nil)
		var seen []any
		for item := range base.DriveCampaignsApi(nil).Stream("list", nil, nil) {
			seen = append(seen, item)
		}
		if len(seen) != 3 {
			t.Fatalf("expected 3 streamed items, got %d", len(seen))
		}

		// Inbound: streaming active -> yields each item from the feature iterator.
		hasStreaming := false
		if fm, ok := core.SharedConfig()["feature"].(map[string]any); ok {
			_, hasStreaming = fm["streaming"]
		}
		if hasStreaming {
			streamSdk := sdk.TestSDK(seed, map[string]any{
				"feature": map[string]any{"streaming": map[string]any{"active": true}},
			})
			var got []any
			for item := range streamSdk.DriveCampaignsApi(nil).Stream("list", nil, nil) {
				if sub, ok := item.([]any); ok {
					got = append(got, sub...)
				} else {
					got = append(got, item)
				}
			}
			if len(got) != 3 {
				t.Fatalf("expected 3 items via streaming feature, got %d", len(got))
			}
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := drive_campaigns_apiBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "list", "update", "load", "remove"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "drive_campaigns_api." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set BUD_DRIVE_TEST_DRIVE_CAMPAIGNS_API_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		driveCampaignsApiRef01Ent := client.DriveCampaignsApi(nil)
		driveCampaignsApiRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath(setup.data, []any{"new", "drive_campaigns_api"}), "drive_campaigns_api_ref01"))

		driveCampaignsApiRef01DataResult, err := driveCampaignsApiRef01Ent.Create(driveCampaignsApiRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		driveCampaignsApiRef01Data = core.ToMapAny(entityData(driveCampaignsApiRef01DataResult))
		if driveCampaignsApiRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}
		if driveCampaignsApiRef01Data["id"] == nil {
			t.Fatal("expected created entity to have an id")
		}

		// LIST
		driveCampaignsApiRef01Match := map[string]any{}

		driveCampaignsApiRef01ListResult, err := driveCampaignsApiRef01Ent.List(driveCampaignsApiRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		driveCampaignsApiRef01List, driveCampaignsApiRef01ListOk := driveCampaignsApiRef01ListResult.([]any)
		if !driveCampaignsApiRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", driveCampaignsApiRef01ListResult)
		}

		foundItem := vs.Select(entityListToData(driveCampaignsApiRef01List), map[string]any{"id": driveCampaignsApiRef01Data["id"]})
		if vs.IsEmpty(foundItem) {
			t.Fatal("expected to find created entity in list")
		}

		// UPDATE
		driveCampaignsApiRef01DataUp0Up := map[string]any{
			"id": driveCampaignsApiRef01Data["id"],
		}

		driveCampaignsApiRef01MarkdefUp0Name := "audience"
		driveCampaignsApiRef01MarkdefUp0Value := fmt.Sprintf("Mark01-drive_campaigns_api_ref01_%d", setup.now)
		driveCampaignsApiRef01DataUp0Up[driveCampaignsApiRef01MarkdefUp0Name] = driveCampaignsApiRef01MarkdefUp0Value

		driveCampaignsApiRef01ResdataUp0Result, err := driveCampaignsApiRef01Ent.Update(driveCampaignsApiRef01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		driveCampaignsApiRef01ResdataUp0 := core.ToMapAny(entityData(driveCampaignsApiRef01ResdataUp0Result))
		if driveCampaignsApiRef01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if driveCampaignsApiRef01ResdataUp0["id"] != driveCampaignsApiRef01DataUp0Up["id"] {
			t.Fatal("expected update result id to match")
		}
		if driveCampaignsApiRef01ResdataUp0[driveCampaignsApiRef01MarkdefUp0Name] != driveCampaignsApiRef01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", driveCampaignsApiRef01MarkdefUp0Name, driveCampaignsApiRef01ResdataUp0[driveCampaignsApiRef01MarkdefUp0Name])
		}

		// LOAD
		driveCampaignsApiRef01MatchDt0 := map[string]any{
			"id": driveCampaignsApiRef01Data["id"],
		}
		driveCampaignsApiRef01DataDt0Loaded, err := driveCampaignsApiRef01Ent.Load(driveCampaignsApiRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		driveCampaignsApiRef01DataDt0LoadResult := core.ToMapAny(entityData(driveCampaignsApiRef01DataDt0Loaded))
		if driveCampaignsApiRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if driveCampaignsApiRef01DataDt0LoadResult["id"] != driveCampaignsApiRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

		// REMOVE
		driveCampaignsApiRef01MatchRm0 := map[string]any{
			"id": driveCampaignsApiRef01Data["id"],
		}
		_, err = driveCampaignsApiRef01Ent.Remove(driveCampaignsApiRef01MatchRm0, nil)
		if err != nil {
			t.Fatalf("remove failed: %v", err)
		}

		// LIST
		driveCampaignsApiRef01MatchRt0 := map[string]any{}

		driveCampaignsApiRef01ListRt0Result, err := driveCampaignsApiRef01Ent.List(driveCampaignsApiRef01MatchRt0, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		driveCampaignsApiRef01ListRt0, driveCampaignsApiRef01ListRt0Ok := driveCampaignsApiRef01ListRt0Result.([]any)
		if !driveCampaignsApiRef01ListRt0Ok {
			t.Fatalf("expected list result to be an array, got %T", driveCampaignsApiRef01ListRt0Result)
		}

		notFoundItem := vs.Select(entityListToData(driveCampaignsApiRef01ListRt0), map[string]any{"id": driveCampaignsApiRef01Data["id"]})
		if !vs.IsEmpty(notFoundItem) {
			t.Fatal("expected removed entity to not be in list")
		}

	})
}

func drive_campaigns_apiBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "drive_campaigns_api", "DriveCampaignsApiTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read drive_campaigns_api test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse drive_campaigns_api test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap, _ := vs.Transform(
		[]any{"drive_campaigns_api01", "drive_campaigns_api02", "drive_campaigns_api03"},
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
	entidEnvRaw := os.Getenv("BUD_DRIVE_TEST_DRIVE_CAMPAIGNS_API_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"BUD_DRIVE_TEST_DRIVE_CAMPAIGNS_API_ENTID": idmap,
		"BUD_DRIVE_TEST_LIVE":      "FALSE",
		"BUD_DRIVE_TEST_EXPLAIN":   "FALSE",
		"BUD_DRIVE_APIKEY":         "",
	})

	idmapResolved := core.ToMapAny(env["BUD_DRIVE_TEST_DRIVE_CAMPAIGNS_API_ENTID"])
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
