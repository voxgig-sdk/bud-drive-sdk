# BudDrive Golang SDK



The Golang SDK for the BudDrive API — an entity-oriented client using standard Go conventions. No generics required; data flows as `map[string]any`.

It exposes the API as capitalised, semantic **Entities** — e.g. `client.DriveCampaignsApi(nil)` — each with the same small set of operations (`List`, `Load`, `Create`, `Update`, `Remove`, `Patch`) instead of raw URL paths and query strings. You call meaning, not endpoints, which keeps the cognitive load low.

> Also generated from this model: `go-cli`, `go-mcp`, `js`, `lua`, `php`, `py`, `ts` — see
> the [top-level README](../README.md).


## Install
```bash
go get github.com/voxgig-sdk/bud-drive-sdk/go@latest
```

The Go module proxy resolves the version from the `go/vX.Y.Z` GitHub
release tag — see [Releases](https://github.com/voxgig-sdk/bud-drive-sdk/releases) for the available versions.

To vendor from a local checkout instead, clone this repo alongside your
project and add a `replace` directive pointing at the checked-out
`go/` directory:

```bash
go mod edit -replace github.com/voxgig-sdk/bud-drive-sdk/go=../bud-drive-sdk/go
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### Quickstart

A complete program: create a client, then call the entity operations.
Each operation returns `(value, error)` — the value is the data itself
(there is no `{ok, data}` wrapper), so check `err` and use the value
directly.

```go
package main

import (
    "fmt"
    "os"
    sdk "github.com/voxgig-sdk/bud-drive-sdk/go"
)

func main() {
    client := sdk.NewBudDriveSDK(map[string]any{
        "apikey": os.Getenv("BUD_DRIVE_APIKEY"),
    })

    // List driveCampaignsApi records — the value is the array of records itself.
    driveCampaignsApis, err := client.DriveCampaignsApi(nil).List(nil, nil)
    if err != nil {
        panic(err)
    }
    for _, item := range driveCampaignsApis.([]any) {
        fmt.Println(item)
    }

    // Load a single driveCampaignsApi — the value is the loaded record.
    driveCampaignsApi, err := client.DriveCampaignsApi(nil).Load(map[string]any{"campaign_id": "example_campaign_id"}, nil)
    if err != nil {
        panic(err)
    }
    fmt.Println(driveCampaignsApi)

    // Create a driveCampaignsApi.
    created, err := client.DriveCampaignsApi(nil).Create(map[string]any{"campaign_goal": "example_campaign_goal", "campaign_id": "example_campaign_id"}, nil)
    if err != nil {
        panic(err)
    }
    fmt.Println(created)

    // Update a driveCampaignsApi.
    updated, err := client.DriveCampaignsApi(nil).Update(map[string]any{"campaign_id": "example_campaign_id", "audience": "example_audience"}, nil)
    if err != nil {
        panic(err)
    }
    fmt.Println(updated)

    // Remove a driveCampaignsApi.
    removed, err := client.DriveCampaignsApi(nil).Remove(map[string]any{"campaign_id": "example_campaign_id"}, nil)
    if err != nil {
        panic(err)
    }
    fmt.Println(removed)
}
```


## Error handling

Every entity operation returns `(value, error)`. Check `err` before
using the value — there is no exception to catch:

```go
drivecampaignsapis, err := client.DriveCampaignsApi(nil).List(nil, nil)
if err != nil {
    // handle err
    return
}
_ = drivecampaignsapis
```

`Direct` follows the same `(value, error)` convention:

```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example_id"},
})
if err != nil {
    // handle err
}
_ = result
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example"},
})
if err != nil {
    panic(err)
}

if result["ok"] == true {
    fmt.Println(result["status"]) // 200
    fmt.Println(result["data"])   // response body
}
```

### Prepare a request without sending it

```go
fetchdef, err := client.Prepare(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "DELETE",
    "params": map[string]any{"id": "example"},
})
if err != nil {
    panic(err)
}

fmt.Println(fetchdef["url"])
fmt.Println(fetchdef["method"])
fmt.Println(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```go
client := sdk.Test()

driveCampaignsApi, err := client.DriveCampaignsApi(nil).List(
    nil, nil,
)
if err != nil {
    panic(err)
}
fmt.Println(driveCampaignsApi) // the returned mock data
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```go
mockFetch := func(url string, init map[string]any) (map[string]any, error) {
    return map[string]any{
        "status":     200,
        "statusText": "OK",
        "headers":    map[string]any{},
        "json": (func() any)(func() any {
            return map[string]any{"id": "mock01"}
        }),
    }, nil
}

client := sdk.NewBudDriveSDK(map[string]any{
    "base": "http://localhost:8080",
    "system": map[string]any{
        "fetch": (func(string, map[string]any) (map[string]any, error))(mockFetch),
    },
})
```

### Run live tests

Create a `.env.local` file at the project root:

```
BUD_DRIVE_TEST_LIVE=TRUE
BUD_DRIVE_APIKEY=<your-key>
```

Then run:

```bash
cd go && go test ./test/...
```


## Reference

### NewBudDriveSDK

```go
func NewBudDriveSDK(options map[string]any) *BudDriveSDK
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `"apikey"` | `string` | API key for authentication. |
| `"base"` | `string` | Base URL of the API server. |
| `"prefix"` | `string` | URL path prefix prepended to all requests. |
| `"suffix"` | `string` | URL path suffix appended to all requests. |
| `"feature"` | `map[string]any` | Feature activation flags. |
| `"extend"` | `[]any` | Additional Feature instances to load. |
| `"system"` | `map[string]any` | System overrides (e.g. custom `"fetch"` function). |

### TestSDK

```go
func TestSDK(testopts map[string]any, sdkopts map[string]any) *BudDriveSDK
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### BudDriveSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `OptionsMap` | `() map[string]any` | Deep copy of current SDK options. |
| `GetUtility` | `() *Utility` | Copy of the SDK utility object. |
| `Prepare` | `(fetchargs map[string]any) (map[string]any, error)` | Build an HTTP request definition without sending. |
| `Direct` | `(fetchargs map[string]any) (map[string]any, error)` | Build and send an HTTP request. |
| `DriveCampaignsApi` | `(data map[string]any) BudDriveEntity` | Create a DriveCampaignsApi entity instance. |
| `DriveMcpApi` | `(data map[string]any) BudDriveEntity` | Create a DriveMcpApi entity instance. |
| `DriveSegmentsApi` | `(data map[string]any) BudDriveEntity` | Create a DriveSegmentsApi entity instance. |

### Entity interface (BudDriveEntity)

All entities implement the `BudDriveEntity` interface.

| Method | Signature | Description |
| --- | --- | --- |
| `Load` | `(reqmatch, ctrl map[string]any) (any, error)` | Load a single entity by match criteria. |
| `List` | `(reqmatch, ctrl map[string]any) (any, error)` | List entities matching the criteria. |
| `Create` | `(reqdata, ctrl map[string]any) (any, error)` | Create a new entity. |
| `Update` | `(reqdata, ctrl map[string]any) (any, error)` | Update an existing entity. |
| `Remove` | `(reqmatch, ctrl map[string]any) (any, error)` | Remove an entity. |
| `Data` | `(args ...any) any` | Get or set entity data. |
| `Match` | `(args ...any) any` | Get or set entity match criteria. |
| `Make` | `() Entity` | Create a new instance with the same options. |
| `GetName` | `() string` | Return the entity name. |

### Result shape

Entity operations return `(value, error)`. The `value` is the
operation's data **directly** — there is no wrapper:

| Operation | `value` |
| --- | --- |
| `Load` / `Create` / `Update` / `Remove` | the entity record (`map[string]any`) |
| `List` | a `[]any` of entity records |

Check `err` first, then use the value directly (or the typed
`...Typed` variants, which return the entity's model struct and a typed
slice):

    driveCampaignsApi, err := client.DriveCampaignsApi(nil).List(map[string]any{/* fields */}, nil)
    if err != nil { /* handle */ }
    // driveCampaignsApi is the returned record

Only `Direct()` returns a response envelope — a `map[string]any` with
`"ok"`, `"status"`, `"headers"`, and `"data"` keys.

### Entities

#### DriveCampaignsApi

| Field | Description |
| --- | --- |
| `"audience"` | The ID of the segment this campaign targets. |
| `"branding"` | Overrides your configured value for this run only. |
| `"campaign_goal"` | The business goal the campaign should serve, in plain language. |
| `"campaign_id"` |  |
| `"channels"` |  |
| `"completed_at"` |  |
| `"context"` | Overrides your configured value for this run only. |
| `"created_at"` |  |
| `"created_by"` |  |
| `"description"` |  |
| `"duration"` |  |
| `"examples"` |  |
| `"extra_context"` | Overrides your configured value for this run only. |
| `"group"` |  |
| `"headline"` |  |
| `"id"` |  |
| `"image_style"` | Overrides your configured value for this run only. |
| `"image_url"` |  |
| `"label"` |  |
| `"message"` |  |
| `"name"` |  |
| `"parameters"` |  |
| `"primary_button"` |  |
| `"priority_order"` |  |
| `"result"` | The campaign a completed run produced. |
| `"schedule"` | When the campaign is published. |
| `"started_at"` |  |
| `"status"` | The campaign's current status. |
| `"styling"` |  |
| `"template_id"` |  |
| `"template_variables"` |  |
| `"templates"` |  |
| `"tone_of_voice"` | Overrides your configured value for this run only. |
| `"trend"` |  |
| `"type"` |  |

Operations: Create, List, Load, Remove, Update.

API path: `/drive-api/v1/campaigns`

#### DriveMcpApi

| Field | Description |
| --- | --- |
| `"error"` |  |
| `"id"` |  |
| `"jsonrpc"` |  |
| `"method"` |  |
| `"params"` |  |
| `"result"` |  |

Operations: Create.

API path: `/drive-mcp/v1/campaigns`

#### DriveSegmentsApi

| Field | Description |
| --- | --- |
| `"created_at"` |  |
| `"created_by"` |  |
| `"criteria"` |  |
| `"criteria_from"` | The ID of the segment whose criteria should be copied. |
| `"customers"` | Customer IDs to include in the fixed customer list. |
| `"customers_from"` | The ID of the segment whose current members should be copied. |
| `"description"` |  |
| `"fixed_customer_list"` |  |
| `"fixed_customer_list_size"` |  |
| `"from_criteria_id"` |  |
| `"id"` |  |
| `"intersection_count"` |  |
| `"jaccard_index"` |  |
| `"name"` |  |
| `"parameters"` |  |
| `"source"` | How the segment was created. |
| `"statistics_id"` |  |
| `"suggestion"` |  |
| `"tag"` |  |
| `"tags"` |  |
| `"type"` | The shape of the statistic's values. |
| `"union_count"` |  |
| `"upload"` |  |
| `"values"` |  |

Operations: Create, List, Load, Patch, Remove, Update.

API path: `/drive-api/v1/segments/{segment_id}/query`



## Entities


### DriveCampaignsApi

Create an instance: `driveCampaignsApi := client.DriveCampaignsApi(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `audience` | `string` | The ID of the segment this campaign targets. |
| `branding` | `string` | Overrides your configured value for this run only. |
| `campaign_goal` | `string` | The business goal the campaign should serve, in plain language. |
| `campaign_id` | `string` |  |
| `channels` | `map[string]any` |  |
| `completed_at` | `string` |  |
| `context` | `string` | Overrides your configured value for this run only. |
| `created_at` | `string` |  |
| `created_by` | `string` |  |
| `description` | `string` |  |
| `duration` | `map[string]any` |  |
| `examples` | `[]any` |  |
| `extra_context` | `string` | Overrides your configured value for this run only. |
| `group` | `string` |  |
| `headline` | `map[string]any` |  |
| `id` | `string` |  |
| `image_style` | `string` | Overrides your configured value for this run only. |
| `image_url` | `string` |  |
| `label` | `string` |  |
| `message` | `string` |  |
| `name` | `string` |  |
| `parameters` | `[]any` |  |
| `primary_button` | `map[string]any` |  |
| `priority_order` | `int` |  |
| `result` | `map[string]any` | The campaign a completed run produced. |
| `schedule` | `map[string]any` | When the campaign is published. |
| `started_at` | `string` |  |
| `status` | `string` | The campaign's current status. |
| `styling` | `map[string]any` |  |
| `template_id` | `string` |  |
| `template_variables` | `map[string]any` |  |
| `templates` | `[]any` |  |
| `tone_of_voice` | `string` | Overrides your configured value for this run only. |
| `trend` | `map[string]any` |  |
| `type` | `string` |  |

#### Example: Load

```go
driveCampaignsApi, err := client.DriveCampaignsApi(nil).Load(map[string]any{"campaign_id": "campaign_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(driveCampaignsApi) // the loaded record
```

#### Example: List

```go
driveCampaignsApis, err := client.DriveCampaignsApi(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(driveCampaignsApis) // the array of records
```

#### Example: Create

```go
result, err := client.DriveCampaignsApi(nil).Create(map[string]any{
    "campaign_goal": "example_campaign_goal",
    "campaign_id": "example_campaign_id",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### DriveMcpApi

Create an instance: `driveMcpApi := client.DriveMcpApi(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `error` | `map[string]any` |  |
| `id` | `string` |  |
| `jsonrpc` | `string` |  |
| `method` | `string` |  |
| `params` | `map[string]any` |  |
| `result` | `map[string]any` |  |

#### Example: Create

```go
result, err := client.DriveMcpApi(nil).Create(map[string]any{
    "id": "example_id",
    "jsonrpc": "example_jsonrpc",
    "method": "example_method",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### DriveSegmentsApi

Create an instance: `driveSegmentsApi := client.DriveSegmentsApi(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` |  |
| `created_by` | `string` |  |
| `criteria` | `any` |  |
| `criteria_from` | `string` | The ID of the segment whose criteria should be copied. |
| `customers` | `[]any` | Customer IDs to include in the fixed customer list. |
| `customers_from` | `string` | The ID of the segment whose current members should be copied. |
| `description` | `string` |  |
| `fixed_customer_list` | `bool` |  |
| `fixed_customer_list_size` | `int` |  |
| `from_criteria_id` | `string` |  |
| `id` | `string` |  |
| `intersection_count` | `int` |  |
| `jaccard_index` | `float64` |  |
| `name` | `string` |  |
| `parameters` | `map[string]any` |  |
| `source` | `string` | How the segment was created. |
| `statistics_id` | `string` |  |
| `suggestion` | `map[string]any` |  |
| `tag` | `string` |  |
| `tags` | `[]any` |  |
| `type` | `string` | The shape of the statistic's values. |
| `union_count` | `int` |  |
| `upload` | `map[string]any` |  |
| `values` | `[]any` |  |

#### Example: Load

```go
driveSegmentsApi, err := client.DriveSegmentsApi(nil).Load(map[string]any{"segment_id": "segment_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(driveSegmentsApi) // the loaded record
```

#### Example: List

```go
driveSegmentsApis, err := client.DriveSegmentsApi(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(driveSegmentsApis) // the array of records
```

#### Example: Create

```go
result, err := client.DriveSegmentsApi(nil).Create(map[string]any{
    "criteria": "example_criteria",
    "criteria_from": "example_criteria_from",
    "customers": []any{},
    "customers_from": "example_customers_from",
    "id": "example_id",
    "name": "example_name",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

## Features

This SDK ships 8 optional features. Each is **inactive until you
switch it on**, so an SDK you have not configured behaves exactly as if none of
them existed — no retries, no cache, no logging, no measurable overhead.

Activate a feature by name in the client options, alongside the options shown
above:

| Feature | What it does |
|---|---|
| [`debug`](#debug) | Debug capture |
| [`idempotency`](#idempotency) | Idempotency |
| [`metrics`](#metrics) | Metrics |
| [`paging`](#paging) | Paging |
| [`ratelimit`](#ratelimit) | Rate limiting |
| [`retry`](#retry) | Retry |
| [`test`](#test) | Test transport |
| [`timeout`](#timeout) | Timeout |

> **Order matters for `ratelimit`, `retry`, `timeout`.** These wrap the
> transport, so each one wraps whatever is already installed: the order you
> activate them in IS the nesting order. Activating them as an ordered list
> rather than a map is what fixes that order.

### debug

Debug capture.

| Option | Default |
|---|---|
| `active` | `false` |
| `max` | `100` |
| `redact` | `['authorization', 'cookie', 'set-cookie', 'api-key', 'apikey', 'x-api-key', 'idempotency-key']` |

Set `feature.debug.active` to enable it, then override any of the options above.

### idempotency

Idempotency.

| Option | Default |
|---|---|
| `active` | `false` |
| `header` | `'Idempotency-Key'` |
| `methods` | `['POST', 'PUT', 'PATCH', 'DELETE']` |
| `ops` | `['create', 'update', 'remove']` |

Set `feature.idempotency.active` to enable it, then override any of the options above.

### metrics

Metrics.

| Option | Default |
|---|---|
| `active` | `false` |

Set `feature.metrics.active` to enable it, then override any of the options above.

### paging

Paging.

| Option | Default |
|---|---|
| `active` | `false` |
| `afterVar` | `'after'` |
| `cursorParam` | `'cursor'` |
| `firstVar` | `'first'` |
| `limitParam` | `'limit'` |
| `pageParam` | `'page'` |
| `startPage` | `1` |

Set `feature.paging.active` to enable it, then override any of the options above.

### ratelimit

Rate limiting.

| Option | Default |
|---|---|
| `active` | `false` |
| `burst` | `5` |
| `rate` | `5` |

Set `feature.ratelimit.active` to enable it, then override any of the options above.

`ratelimit` wraps the transport, so its position among the other
transport features decides what it sees. A feature activated later wraps one
activated earlier.

### retry

Retry.

| Option | Default |
|---|---|
| `active` | `false` |
| `factor` | `2` |
| `maxDelay` | `2000` |
| `minDelay` | `50` |
| `retries` | `2` |
| `statuses` | `[408, 425, 429, 500, 502, 503, 504]` |

Set `feature.retry.active` to enable it, then override any of the options above.

`retry` wraps the transport, so its position among the other
transport features decides what it sees. A feature activated later wraps one
activated earlier.

### test

Test transport.

| Option | Default |
|---|---|
| `active` | `false` |

Set `feature.test.active` to enable it, then override any of the options above.

### timeout

Timeout.

| Option | Default |
|---|---|
| `active` | `false` |
| `ms` | `30000` |

Set `feature.timeout.active` to enable it, then override any of the options above.

`timeout` wraps the transport, so its position among the other
transport features decides what it sees. A feature activated later wraps one
activated earlier.


## Open types

2 fields are carried as open values rather than typed structures.
This follows from the API definition, not from a gap in this SDK: the
definition describes them with untagged unions —
`oneOf`/`anyOf` branches with no `discriminator` — so it never states which
variant a given value is. Nothing can select a branch reliably, so the SDK
passes the value through unchanged rather than assert a shape the API does not
guarantee.

| Entity | Field | Variants | Nesting |
| --- | --- | --- | --- |
| `drive_campaigns_api` | `templates` | 3 | 3 levels |
| `drive_segments_api` | `criteria` | 3 | 0 levels |

These values round-trip unchanged — read them, modify them, send them back. If
the API adds a `discriminator` to the definition, regenerating will type them.
Every other field is typed normally.

## Advanced

> The sections above cover everyday use. The material below explains the
> SDK's internals — useful when extending it with custom features, but not
> needed for normal use.

### The operation pipeline

Every entity operation follows a six-stage pipeline. Each stage fires a
feature hook before executing:

```
PrePoint → PreSpec → PreRequest → PreResponse → PreResult → PreDone
```

- **PrePoint**: Resolves which API endpoint to call based on the
  operation name and entity configuration.
- **PreSpec**: Builds the HTTP spec — URL, method, headers, body —
  from the resolved point and the caller's parameters.
- **PreRequest**: Sends the HTTP request. Features can intercept here
  to replace the transport (as TestFeature does with mocks).
- **PreResponse**: Parses the raw HTTP response.
- **PreResult**: Extracts the business data from the parsed response.
- **PreDone**: Final stage before returning to the caller. Entity
  state (match, data) is updated here.

If any stage errors, the pipeline short-circuits and the error surfaces
to the caller — see [Error handling](#error-handling) for how that looks
in this language.

### Features and hooks

Features are the extension mechanism. A feature implements the
`Feature` interface and provides hooks — functions keyed by pipeline
stage names.

The SDK ships with built-in features:

- **DebugFeature**: Debug capture
- **IdempotencyFeature**: Idempotency
- **MetricsFeature**: Metrics
- **PagingFeature**: Paging
- **RatelimitFeature**: Rate limiting
- **RetryFeature**: Retry
- **TestFeature**: Test transport
- **TimeoutFeature**: Timeout

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as maps

The Go SDK uses `map[string]any` throughout rather than typed structs.
This mirrors the dynamic nature of the API and keeps the SDK
flexible — no code generation is needed when the API schema changes.

Use `core.ToMapAny()` to safely cast results and nested data.

### Package structure

```
github.com/voxgig-sdk/bud-drive-sdk/go/
├── bud-drive.go        # Root package — type aliases and constructors
├── core/               # SDK core — client, types, pipeline
├── entity/             # Entity implementations
├── feature/            # Built-in features (Base, Test, Log)
├── utility/            # Utility functions and struct library
└── test/               # Test suites
```

The root package (`github.com/voxgig-sdk/bud-drive-sdk/go`) re-exports everything needed
for normal use. Import sub-packages only when you need specific types
like `core.ToMapAny`.

### Entity state

Entity instances are stateful. After a successful `List`, the entity
stores the returned data and match criteria internally.

```go
drivecampaignsapi := client.DriveCampaignsApi(nil)
drivecampaignsapi.List(nil, nil)

// drivecampaignsapi.Data() now returns the drivecampaignsapi data from the last list
// drivecampaignsapi.Match() returns the last match criteria
```

Call `Make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`Direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `Prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
