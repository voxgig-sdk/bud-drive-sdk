# BudDrive Python SDK



The Python SDK for the BudDrive API — an entity-oriented client following Pythonic conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `client.DriveCampaignsApi()` — each
carrying a small, uniform set of operations (`list`, `load`, `create`, `update`, `remove`, `patch`) instead of raw URL
paths and query strings. You work with named resources and verbs, which
keeps the cognitive load low.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to PyPI. Install it from the GitHub
release tag (`py/vX.Y.Z`, see [Releases](https://github.com/voxgig-sdk/bud-drive-sdk/releases)) or
from a source checkout:

```bash
pip install -e .
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```python
import os
from buddrive_sdk import BudDriveSDK

client = BudDriveSDK({
    "apikey": os.environ.get("BUD_DRIVE_APIKEY"),
})
```

### 2. List drivecampaignsapi records

`list()` returns a `list` of records (each a `dict`) and raises on
error — iterate it directly.

```python
try:
    drivecampaignsapis = client.DriveCampaignsApi().list()
    for drivecampaignsapi in drivecampaignsapis:
        print(drivecampaignsapi)
except Exception as err:
    print(f"list failed: {err}")
```

### 3. Load a drivecampaignsapi

`load()` returns the ENTITY — call data_get() for the record — and raises on error.

```python
try:
    drivecampaignsapi = client.DriveCampaignsApi().load({"campaign_id": "example_campaign_id"})
    print(drivecampaignsapi)
except Exception as err:
    print(f"load failed: {err}")
```

### 4. Create, update, and remove

```python
# Create — returns the ENTITY (call data_get() for the record)
created = client.DriveCampaignsApi().create({"campaign_goal": "example_campaign_goal", "campaign_id": "example_campaign_id"})

# Update — the created record's id is a plain dict key
client.DriveCampaignsApi().update({"campaign_id": "example_campaign_id", "audience": "example_audience"})

# Remove
client.DriveCampaignsApi().remove({"campaign_id": "example_campaign_id"})
```


## Error handling

Entity operations raise on failure, so wrap them in `try` / `except`:

```python
try:
    drivecampaignsapis = client.DriveCampaignsApi().list()
    print(drivecampaignsapis)
except Exception as err:
    print(f"list failed: {err}")
```

`direct()` does **not** raise — it returns the result envelope. Branch
on `ok`; on failure `status` holds the HTTP status (for error responses)
and `err` holds a transport error, so read both defensively:

```python
result = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example_id"},
})

if not result["ok"]:
    print("request failed:", result.get("status"), result.get("err"))
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```python
result = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example"},
})

if result["ok"]:
    print(result["status"])  # 200
    print(result["data"])    # response body
else:
    # A non-2xx response carries status + data (the error body); a
    # transport-level failure carries err instead. Only one is present, so
    # read both with .get() rather than indexing a key that may be absent.
    print(result.get("status"), result.get("err"))
```

### Prepare a request without sending it

```python
# prepare() returns the fetch definition and raises on error.
fetchdef = client.prepare({
    "path": "/api/resource/{id}",
    "method": "DELETE",
    "params": {"id": "example"},
})

print(fetchdef["url"])
print(fetchdef["method"])
print(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```python
client = BudDriveSDK.test()

# Entity ops return the ENTITY and raises on error;
# call data_get() for the record.
drivecampaignsapi = client.DriveCampaignsApi().list()
# drivecampaignsapi contains the mock response record
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```python
def mock_fetch(url, init):
    return {
        "status": 200,
        "statusText": "OK",
        "headers": {},
        "json": lambda: {"id": "mock01"},
    }, None

client = BudDriveSDK({
    "base": "http://localhost:8080",
    "system": {
        "fetch": mock_fetch,
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
cd py && pytest test/
```


## Reference

### BudDriveSDK

```python
from buddrive_sdk import BudDriveSDK

client = BudDriveSDK(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `str` | API key for authentication. |
| `base` | `str` | Base URL of the API server. |
| `prefix` | `str` | URL path prefix prepended to all requests. |
| `suffix` | `str` | URL path suffix appended to all requests. |
| `feature` | `dict` | Feature activation flags. |
| `extend` | `list` | Additional Feature instances to load. |
| `system` | `dict` | System overrides (e.g. custom `fetch` function). |

### test

```python
client = BudDriveSDK.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `None`.

### BudDriveSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> dict` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> dict` | Build an HTTP request definition without sending. Raises on error. |
| `direct` | `(fetchargs) -> dict` | Build and send an HTTP request. Returns a result dict (branch on `ok`). |
| `DriveCampaignsApi` | `(data) -> DriveCampaignsApiEntity` | Create a DriveCampaignsApi entity instance. |
| `DriveMcpApi` | `(data) -> DriveMcpApiEntity` | Create a DriveMcpApi entity instance. |
| `DriveSegmentsApi` | `(data) -> DriveSegmentsApiEntity` | Create a DriveSegmentsApi entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `(reqmatch, ctrl) -> any` | Load a single entity by match criteria. Raises on error. |
| `list` | `(reqmatch, ctrl) -> list` | List entities matching the criteria. Raises on error. |
| `create` | `(reqdata, ctrl) -> any` | Create a new entity. Raises on error. |
| `update` | `(reqdata, ctrl) -> any` | Update an existing entity. Raises on error. |
| `remove` | `(reqmatch, ctrl) -> any` | Remove an entity. Raises on error. |
| `data_get` | `() -> dict` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> dict` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> str` | Return the entity name. |

### Result shape

Entity operations return the ENTITY (call data_get() for the record) (a `dict` for single-entity
ops, a `list` for `list`) and raise on error. Wrap calls in
`try`/`except` to handle failures.

The `direct()` escape hatch never raises — it returns a result `dict`
you branch on via `result["ok"]`:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `bool` | `True` if the HTTP status is 2xx. |
| `status` | `int` | HTTP status code. |
| `headers` | `dict` | Response headers. |
| `data` | `any` | Parsed JSON response body. |

On error, `ok` is `False` and `err` contains the error value.

### Entities

#### DriveCampaignsApi

| Field | Description |
| --- | --- |
| `audience` | The ID of the segment this campaign targets. |
| `branding` | Overrides your configured value for this run only. |
| `campaign_goal` | The business goal the campaign should serve, in plain language. |
| `campaign_id` |  |
| `channels` |  |
| `completed_at` |  |
| `context` | Overrides your configured value for this run only. |
| `created_at` |  |
| `created_by` |  |
| `description` |  |
| `duration` |  |
| `examples` |  |
| `extra_context` | Overrides your configured value for this run only. |
| `group` |  |
| `headline` |  |
| `id` |  |
| `image_style` | Overrides your configured value for this run only. |
| `image_url` |  |
| `label` |  |
| `message` |  |
| `name` |  |
| `parameters` |  |
| `primary_button` |  |
| `priority_order` |  |
| `result` | The campaign a completed run produced. |
| `schedule` | When the campaign is published. |
| `started_at` |  |
| `status` | The campaign's current status. |
| `styling` |  |
| `template_id` |  |
| `template_variables` |  |
| `templates` |  |
| `tone_of_voice` | Overrides your configured value for this run only. |
| `trend` |  |
| `type` |  |

Operations: Create, List, Load, Remove, Update.

API path: `/drive-api/v1/campaigns`

#### DriveMcpApi

| Field | Description |
| --- | --- |
| `error` |  |
| `id` |  |
| `jsonrpc` |  |
| `method` |  |
| `params` |  |
| `result` |  |

Operations: Create.

API path: `/drive-mcp/v1/campaigns`

#### DriveSegmentsApi

| Field | Description |
| --- | --- |
| `created_at` |  |
| `created_by` |  |
| `criteria` |  |
| `criteria_from` | The ID of the segment whose criteria should be copied. |
| `customers` | Customer IDs to include in the fixed customer list. |
| `customers_from` | The ID of the segment whose current members should be copied. |
| `description` |  |
| `fixed_customer_list` |  |
| `fixed_customer_list_size` |  |
| `from_criteria_id` |  |
| `id` |  |
| `intersection_count` |  |
| `jaccard_index` |  |
| `name` |  |
| `parameters` |  |
| `source` | How the segment was created. |
| `statistics_id` |  |
| `suggestion` |  |
| `tag` |  |
| `tags` |  |
| `type` | The shape of the statistic's values. |
| `union_count` |  |
| `upload` |  |
| `values` |  |

Operations: Create, List, Load, Patch, Remove, Update.

API path: `/drive-api/v1/segments/{segment_id}/query`



## Entities


### DriveCampaignsApi

Create an instance: `drive_campaigns_api = client.DriveCampaignsApi()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `audience` | `str` | The ID of the segment this campaign targets. |
| `branding` | `str` | Overrides your configured value for this run only. |
| `campaign_goal` | `str` | The business goal the campaign should serve, in plain language. |
| `campaign_id` | `str` |  |
| `channels` | `dict` |  |
| `completed_at` | `str` |  |
| `context` | `str` | Overrides your configured value for this run only. |
| `created_at` | `str` |  |
| `created_by` | `str` |  |
| `description` | `str` |  |
| `duration` | `dict` |  |
| `examples` | `list` |  |
| `extra_context` | `str` | Overrides your configured value for this run only. |
| `group` | `str` |  |
| `headline` | `dict` |  |
| `id` | `str` |  |
| `image_style` | `str` | Overrides your configured value for this run only. |
| `image_url` | `str` |  |
| `label` | `str` |  |
| `message` | `str` |  |
| `name` | `str` |  |
| `parameters` | `list` |  |
| `primary_button` | `dict` |  |
| `priority_order` | `int` |  |
| `result` | `dict` | The campaign a completed run produced. |
| `schedule` | `dict` | When the campaign is published. |
| `started_at` | `str` |  |
| `status` | `str` | The campaign's current status. |
| `styling` | `dict` |  |
| `template_id` | `str` |  |
| `template_variables` | `dict` |  |
| `templates` | `list` |  |
| `tone_of_voice` | `str` | Overrides your configured value for this run only. |
| `trend` | `dict` |  |
| `type` | `str` |  |

#### Example: Load

```python
drive_campaigns_api = client.DriveCampaignsApi().load({"campaign_id": "campaign_id"})
```

#### Example: List

```python
drive_campaigns_apis = client.DriveCampaignsApi().list()
```

#### Example: Create

```python
drive_campaigns_api = client.DriveCampaignsApi().create({
    "campaign_goal": "example_campaign_goal",  # str
    "campaign_id": "example_campaign_id",  # str
})
```


### DriveMcpApi

Create an instance: `drive_mcp_api = client.DriveMcpApi()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `error` | `dict` |  |
| `id` | `str` |  |
| `jsonrpc` | `str` |  |
| `method` | `str` |  |
| `params` | `dict` |  |
| `result` | `dict` |  |

#### Example: Create

```python
drive_mcp_api = client.DriveMcpApi().create({
    "id": "example_id",  # str
    "jsonrpc": "example_jsonrpc",  # str
    "method": "example_method",  # str
})
```


### DriveSegmentsApi

Create an instance: `drive_segments_api = client.DriveSegmentsApi()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `str` |  |
| `created_by` | `str` |  |
| `criteria` | `Any` |  |
| `criteria_from` | `str` | The ID of the segment whose criteria should be copied. |
| `customers` | `list` | Customer IDs to include in the fixed customer list. |
| `customers_from` | `str` | The ID of the segment whose current members should be copied. |
| `description` | `str` |  |
| `fixed_customer_list` | `bool` |  |
| `fixed_customer_list_size` | `int` |  |
| `from_criteria_id` | `str` |  |
| `id` | `str` |  |
| `intersection_count` | `int` |  |
| `jaccard_index` | `float` |  |
| `name` | `str` |  |
| `parameters` | `dict` |  |
| `source` | `str` | How the segment was created. |
| `statistics_id` | `str` |  |
| `suggestion` | `dict` |  |
| `tag` | `str` |  |
| `tags` | `list` |  |
| `type` | `str` | The shape of the statistic's values. |
| `union_count` | `int` |  |
| `upload` | `dict` |  |
| `values` | `list` |  |

#### Example: Load

```python
drive_segments_api = client.DriveSegmentsApi().load({"segment_id": "segment_id"})
```

#### Example: List

```python
drive_segments_apis = client.DriveSegmentsApi().list()
```

#### Example: Create

```python
drive_segments_api = client.DriveSegmentsApi().create({
    "criteria": "example_criteria",  # Any
    "criteria_from": "example_criteria_from",  # str
    "customers": [],  # list
    "customers_from": "example_customers_from",  # str
    "id": "example_id",  # str
    "name": "example_name",  # str
})
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

Features are the extension mechanism. A feature is a Python class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

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

### Data as dicts

The Python SDK uses plain dicts throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `helpers.to_map()` to safely validate that a value is a dict.

### Module structure

```
py/
├── buddrive_sdk.py         -- Main SDK module
├── config.py                    -- Configuration
├── schema.py                    -- Generated option + entity specs
├── features.py                  -- Feature factory
├── core/                        -- Core types and context
├── entity/                      -- Entity implementations
├── feature/                     -- Built-in features (Base, Test, Log)
├── utility/                     -- Utility functions and struct library
└── test/                        -- Test suites
```

The main module (`buddrive_sdk`) exports the SDK class.
Import entity or utility modules directly only when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```python
drivecampaignsapi = client.DriveCampaignsApi()
drivecampaignsapi.list()

# drivecampaignsapi.data_get() now returns the drivecampaignsapi data from the last list
# drivecampaignsapi.match_get() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
