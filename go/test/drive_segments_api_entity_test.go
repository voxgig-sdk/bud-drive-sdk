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

func TestDriveSegmentsApiEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.DriveSegmentsApi(nil)
		if ent == nil {
			t.Fatal("expected non-nil DriveSegmentsApiEntity")
		}
	})

	// Feature #4: the entity Stream(action, ...) method runs the op pipeline and
	// returns a channel over result items. With the streaming feature active it
	// yields the feature's incremental output; otherwise it falls back to the
	// materialised list so Stream always yields.
	t.Run("stream", func(t *testing.T) {
		seed := map[string]any{
			"entity": map[string]any{
				"drive_segments_api": map[string]any{
					"s1": map[string]any{"id": "s1"},
					"s2": map[string]any{"id": "s2"},
					"s3": map[string]any{"id": "s3"},
				},
			},
		}

		// Fallback: streaming inactive -> yields the materialised list items.
		base := sdk.TestSDK(seed, nil)
		var seen []any
		for item := range base.DriveSegmentsApi(nil).Stream("list", nil, nil) {
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
			for item := range streamSdk.DriveSegmentsApi(nil).Stream("list", nil, nil) {
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
		setup := drive_segments_apiBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "list", "update", "load", "remove"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "drive_segments_api." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set BUD_DRIVE_TEST_DRIVE_SEGMENTS_API_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		driveSegmentsApiRef01Ent := client.DriveSegmentsApi(nil)
		driveSegmentsApiRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath(setup.data, []any{"new", "drive_segments_api"}), "drive_segments_api_ref01"))
		driveSegmentsApiRef01Data["segment_id"] = setup.idmap["segment01"]

		driveSegmentsApiRef01DataResult, err := driveSegmentsApiRef01Ent.Create(driveSegmentsApiRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		driveSegmentsApiRef01Data = core.ToMapAny(entityData(driveSegmentsApiRef01DataResult))
		if driveSegmentsApiRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}
		if driveSegmentsApiRef01Data["id"] == nil {
			t.Fatal("expected created entity to have an id")
		}

		// LIST
		driveSegmentsApiRef01Match := map[string]any{}

		driveSegmentsApiRef01ListResult, err := driveSegmentsApiRef01Ent.List(driveSegmentsApiRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		driveSegmentsApiRef01List, driveSegmentsApiRef01ListOk := driveSegmentsApiRef01ListResult.([]any)
		if !driveSegmentsApiRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", driveSegmentsApiRef01ListResult)
		}

		foundItem := vs.Select(entityListToData(driveSegmentsApiRef01List), map[string]any{"id": driveSegmentsApiRef01Data["id"]})
		if vs.IsEmpty(foundItem) {
			t.Fatal("expected to find created entity in list")
		}

		// UPDATE
		driveSegmentsApiRef01DataUp0Up := map[string]any{
			"id": driveSegmentsApiRef01Data["id"],
		}

		driveSegmentsApiRef01MarkdefUp0Name := "created_at"
		driveSegmentsApiRef01MarkdefUp0Value := fmt.Sprintf("Mark01-drive_segments_api_ref01_%d", setup.now)
		driveSegmentsApiRef01DataUp0Up[driveSegmentsApiRef01MarkdefUp0Name] = driveSegmentsApiRef01MarkdefUp0Value

		driveSegmentsApiRef01ResdataUp0Result, err := driveSegmentsApiRef01Ent.Update(driveSegmentsApiRef01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		driveSegmentsApiRef01ResdataUp0 := core.ToMapAny(entityData(driveSegmentsApiRef01ResdataUp0Result))
		if driveSegmentsApiRef01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if driveSegmentsApiRef01ResdataUp0["id"] != driveSegmentsApiRef01DataUp0Up["id"] {
			t.Fatal("expected update result id to match")
		}
		if driveSegmentsApiRef01ResdataUp0[driveSegmentsApiRef01MarkdefUp0Name] != driveSegmentsApiRef01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", driveSegmentsApiRef01MarkdefUp0Name, driveSegmentsApiRef01ResdataUp0[driveSegmentsApiRef01MarkdefUp0Name])
		}

		// LOAD
		driveSegmentsApiRef01MatchDt0 := map[string]any{
			"id": driveSegmentsApiRef01Data["id"],
		}
		driveSegmentsApiRef01DataDt0Loaded, err := driveSegmentsApiRef01Ent.Load(driveSegmentsApiRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		driveSegmentsApiRef01DataDt0LoadResult := core.ToMapAny(entityData(driveSegmentsApiRef01DataDt0Loaded))
		if driveSegmentsApiRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if driveSegmentsApiRef01DataDt0LoadResult["id"] != driveSegmentsApiRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

		// REMOVE
		driveSegmentsApiRef01MatchRm0 := map[string]any{
			"id": driveSegmentsApiRef01Data["id"],
		}
		_, err = driveSegmentsApiRef01Ent.Remove(driveSegmentsApiRef01MatchRm0, nil)
		if err != nil {
			t.Fatalf("remove failed: %v", err)
		}

		// LIST
		driveSegmentsApiRef01MatchRt0 := map[string]any{}

		driveSegmentsApiRef01ListRt0Result, err := driveSegmentsApiRef01Ent.List(driveSegmentsApiRef01MatchRt0, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		driveSegmentsApiRef01ListRt0, driveSegmentsApiRef01ListRt0Ok := driveSegmentsApiRef01ListRt0Result.([]any)
		if !driveSegmentsApiRef01ListRt0Ok {
			t.Fatalf("expected list result to be an array, got %T", driveSegmentsApiRef01ListRt0Result)
		}

		notFoundItem := vs.Select(entityListToData(driveSegmentsApiRef01ListRt0), map[string]any{"id": driveSegmentsApiRef01Data["id"]})
		if !vs.IsEmpty(notFoundItem) {
			t.Fatal("expected removed entity to not be in list")
		}

	})
}

func drive_segments_apiBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "drive_segments_api", "DriveSegmentsApiTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read drive_segments_api test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse drive_segments_api test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap, _ := vs.Transform(
		[]any{"drive_segments_api01", "drive_segments_api02", "drive_segments_api03", "segment01"},
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
	entidEnvRaw := os.Getenv("BUD_DRIVE_TEST_DRIVE_SEGMENTS_API_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"BUD_DRIVE_TEST_DRIVE_SEGMENTS_API_ENTID": idmap,
		"BUD_DRIVE_TEST_LIVE":      "FALSE",
		"BUD_DRIVE_TEST_EXPLAIN":   "FALSE",
		"BUD_DRIVE_APIKEY":         "",
	})

	idmapResolved := core.ToMapAny(env["BUD_DRIVE_TEST_DRIVE_SEGMENTS_API_ENTID"])
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
