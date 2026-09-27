# BudDrive Python SDK Reference

Complete API reference for the BudDrive Python SDK.


## BudDriveSDK

### Constructor

```python
from buddrive_sdk import BudDriveSDK

client = BudDriveSDK(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `dict` | SDK configuration options. |
| `options["apikey"]` | `str` | API key for authentication. |
| `options["base"]` | `str` | Base URL for API requests. |
| `options["prefix"]` | `str` | URL prefix appended after base. |
| `options["suffix"]` | `str` | URL suffix appended after path. |
| `options["headers"]` | `dict` | Custom headers for all requests. |
| `options["feature"]` | `dict` | Feature configuration. |
| `options["system"]` | `dict` | System overrides (e.g. custom fetch). |


### Static Methods

#### `BudDriveSDK.test(testopts=None, sdkopts=None)`

Create a test client with mock features active. Both arguments may be `None`.

```python
client = BudDriveSDK.test()
```


### Instance Methods

#### `DriveCampaignsApi(data=None)`

Create a new `DriveCampaignsApiEntity` instance. Pass `None` for no initial data.

#### `DriveMcpApi(data=None)`

Create a new `DriveMcpApiEntity` instance. Pass `None` for no initial data.

#### `DriveSegmentsApi(data=None)`

Create a new `DriveSegmentsApiEntity` instance. Pass `None` for no initial data.

#### `options_map() -> dict`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs=None) -> dict`

Make a direct HTTP request to any API endpoint. Returns a result `dict` with `ok`, `status`, `headers`, and `data` (or `err` on failure). This escape hatch never raises — branch on `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `str` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `str` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `dict` | Path parameter values. |
| `fetchargs["query"]` | `dict` | Query string parameters. |
| `fetchargs["headers"]` | `dict` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (dicts are JSON-serialized). |

**Returns:** `result_dict`

#### `prepare(fetchargs=None) -> dict`

Prepare a fetch definition without sending. Returns the `fetchdef` and raises on error.


---

## DriveCampaignsApiEntity

```python
drive_campaigns_api = client.DriveCampaignsApi()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `audience` | `str` | No | The ID of the segment this campaign targets. |
| `branding` | `str` | No | Overrides your configured value for this run only. |
| `campaign_goal` | `str` | Yes | The business goal the campaign should serve, in plain language. |
| `campaign_id` | `str` | Yes |  |
| `channels` | `dict` | No |  |
| `completed_at` | `str` | No |  |
| `context` | `str` | No | Overrides your configured value for this run only. |
| `created_at` | `str` | No |  |
| `created_by` | `str` | No |  |
| `description` | `str` | No |  |
| `duration` | `dict` | No |  |
| `examples` | `list` | No |  |
| `extra_context` | `str` | No | Overrides your configured value for this run only. |
| `group` | `str` | No |  |
| `headline` | `dict` | No |  |
| `id` | `str` | No |  |
| `image_style` | `str` | No | Overrides your configured value for this run only. |
| `image_url` | `str` | No |  |
| `label` | `str` | No |  |
| `message` | `str` | No |  |
| `name` | `str` | No |  |
| `parameters` | `list` | No |  |
| `primary_button` | `dict` | No |  |
| `priority_order` | `int` | No |  |
| `result` | `dict` | No | The campaign a completed run produced. |
| `schedule` | `dict` | No | When the campaign is published. |
| `started_at` | `str` | No |  |
| `status` | `str` | No | The campaign's current status. |
| `styling` | `dict` | No |  |
| `template_id` | `str` | No |  |
| `template_variables` | `dict` | No |  |
| `templates` | `list` | No |  |
| `tone_of_voice` | `str` | No | Overrides your configured value for this run only. |
| `trend` | `dict` | No |  |
| `type` | `str` | No |  |

### Field Usage by Operation

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `audience` | - | - | Yes | - | - |
| `branding` | - | - | - | - | - |
| `campaign_goal` | - | - | - | - | - |
| `campaign_id` | - | Yes | - | - | - |
| `channels` | - | - | - | - | - |
| `completed_at` | - | - | - | - | - |
| `context` | - | - | - | - | - |
| `created_at` | - | - | - | - | - |
| `created_by` | - | - | - | - | - |
| `description` | - | - | - | - | - |
| `duration` | - | - | - | - | - |
| `examples` | - | - | - | - | - |
| `extra_context` | - | - | - | - | - |
| `group` | - | - | - | - | - |
| `headline` | - | - | - | - | - |
| `id` | - | - | - | - | - |
| `image_style` | - | - | - | - | - |
| `image_url` | - | - | - | - | - |
| `label` | - | - | - | - | - |
| `message` | - | - | - | - | - |
| `name` | - | - | Yes | - | - |
| `parameters` | - | - | - | - | - |
| `primary_button` | - | - | - | - | - |
| `priority_order` | - | - | - | - | - |
| `result` | - | - | - | - | - |
| `schedule` | - | - | - | - | - |
| `started_at` | - | - | - | - | - |
| `status` | - | - | - | - | - |
| `styling` | - | - | - | - | - |
| `template_id` | - | - | - | - | - |
| `template_variables` | - | - | - | - | - |
| `templates` | - | - | - | - | - |
| `tone_of_voice` | - | - | - | - | - |
| `trend` | - | - | - | - | - |
| `type` | - | - | - | - | - |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.DriveCampaignsApi().create({
    "campaign_goal": "example_campaign_goal",  # str
    "campaign_id": "example_campaign_id",  # str
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.DriveCampaignsApi().list()
for drive_campaigns_api in results:
    print(drive_campaigns_api)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.DriveCampaignsApi().load({"campaign_id": "campaign_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.DriveCampaignsApi().remove({"campaign_id": "campaign_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.DriveCampaignsApi().update({
    "campaign_id": "campaign_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `DriveCampaignsApiEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## DriveMcpApiEntity

```python
drive_mcp_api = client.DriveMcpApi()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `error` | `dict` | No |  |
| `id` | `str` | Yes |  |
| `jsonrpc` | `str` | Yes |  |
| `method` | `str` | Yes |  |
| `params` | `dict` | No |  |
| `result` | `dict` | No |  |

### Field Usage by Operation

| Field | create |
| --- | --- |
| `error` | - |
| `id` | Yes |
| `jsonrpc` | - |
| `method` | - |
| `params` | - |
| `result` | - |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.DriveMcpApi().create({
    "id": "example_id",  # str
    "jsonrpc": "example_jsonrpc",  # str
    "method": "example_method",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `DriveMcpApiEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## DriveSegmentsApiEntity

```python
drive_segments_api = client.DriveSegmentsApi()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `str` | No |  |
| `created_by` | `str` | No |  |
| `criteria` | `Any` | Yes |  |
| `criteria_from` | `str` | Yes | The ID of the segment whose criteria should be copied. |
| `customers` | `list` | Yes | Customer IDs to include in the fixed customer list. |
| `customers_from` | `str` | Yes | The ID of the segment whose current members should be copied. |
| `description` | `str` | No |  |
| `fixed_customer_list` | `bool` | No |  |
| `fixed_customer_list_size` | `int` | No |  |
| `from_criteria_id` | `str` | No |  |
| `id` | `str` | Yes |  |
| `intersection_count` | `int` | No |  |
| `jaccard_index` | `float` | No |  |
| `name` | `str` | Yes |  |
| `parameters` | `dict` | No |  |
| `source` | `str` | No | How the segment was created. |
| `statistics_id` | `str` | No |  |
| `suggestion` | `dict` | No |  |
| `tag` | `str` | No |  |
| `tags` | `list` | No |  |
| `type` | `str` | No | The shape of the statistic's values. |
| `union_count` | `int` | No |  |
| `upload` | `dict` | No |  |
| `values` | `list` | No |  |

### Field Usage by Operation

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `created_at` | - | - | - | - | - |
| `created_by` | - | - | - | - | - |
| `criteria` | - | - | - | - | - |
| `criteria_from` | - | - | - | - | - |
| `customers` | - | - | - | - | - |
| `customers_from` | - | - | - | - | - |
| `description` | - | - | - | - | - |
| `fixed_customer_list` | - | - | - | - | - |
| `fixed_customer_list_size` | - | - | - | - | - |
| `from_criteria_id` | - | - | - | - | - |
| `id` | - | - | - | - | - |
| `intersection_count` | - | - | - | - | - |
| `jaccard_index` | - | - | - | - | - |
| `name` | - | Yes | - | - | - |
| `parameters` | - | - | - | - | - |
| `source` | - | - | - | - | - |
| `statistics_id` | - | - | - | - | - |
| `suggestion` | - | - | - | - | - |
| `tag` | - | - | - | - | - |
| `tags` | - | - | - | - | - |
| `type` | - | - | - | - | - |
| `union_count` | - | - | - | - | - |
| `upload` | - | - | - | - | - |
| `values` | - | - | - | - | - |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.DriveSegmentsApi().create({
    "criteria": "example_criteria",  # Any
    "criteria_from": "example_criteria_from",  # str
    "customers": [],  # list
    "customers_from": "example_customers_from",  # str
    "id": "example_id",  # str
    "name": "example_name",  # str
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.DriveSegmentsApi().list()
for drive_segments_api in results:
    print(drive_segments_api)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.DriveSegmentsApi().load({"segment_id": "segment_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.DriveSegmentsApi().remove({"segment_id": "segment_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.DriveSegmentsApi().update({
    "criteria_id": "criteria_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `DriveSegmentsApiEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `debug` | 0.0.1 | Debug capture |
| `idempotency` | 0.0.1 | Idempotency |
| `metrics` | 0.0.1 | Metrics |
| `paging` | 0.0.1 | Paging |
| `ratelimit` | 0.0.1 | Rate limiting |
| `retry` | 0.0.1 | Retry |
| `test` | 0.0.1 | Test transport |
| `timeout` | 0.0.1 | Timeout |


Features are activated via the `feature` option:

```python
client = BudDriveSDK({
    "feature": {
        "debug": {"active": True},
        "idempotency": {"active": True},
        "metrics": {"active": True},
        "paging": {"active": True},
        "ratelimit": {"active": True},
        "retry": {"active": True},
        "test": {"active": True},
        "timeout": {"active": True},
    },
})
```


### Configuring features

Each feature is inactive until switched on, and an SDK with no feature
configured does no feature work at all. Every option below keeps its default
unless you name it.

The array form of \`feature\` is significant: several features wrap the
transport, and the order you list them in is the order they nest.

#### Ordering

`ratelimit`, `retry`, `timeout` wrap the transport. Each
wraps whatever is already installed, so **activation order is nesting order**:
a feature activated later sits OUTSIDE one activated earlier, and sees the call
first.

That decides behaviour, not just sequence: a feature that short-circuits the
call, such as a cache serving a hit, stops every feature nested inside it from
ever seeing that call.

`debug`, `idempotency`, `metrics`, `paging`, `test` attach to pipeline hooks
rather than the transport, so their order does not affect what they observe.

#### `debug`

Debug capture.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `max` | `100` |
| `redact` | `['authorization', 'cookie', 'set-cookie', 'api-key', 'apikey', 'x-api-key', 'idempotency-key']` |

| Option | Type |
|---|---|
| `now` | function |
| `onEntry` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.debug.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Attaches to pipeline hooks, not the transport, so activation order does
  not change what it observes.
- Inactive by default: leaving it out costs nothing at runtime.

#### `idempotency`

Idempotency.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `header` | `'Idempotency-Key'` |
| `methods` | `['POST', 'PUT', 'PATCH', 'DELETE']` |
| `ops` | `['create', 'update', 'remove']` |

| Option | Type |
|---|---|
| `keygen` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.idempotency.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Attaches to pipeline hooks, not the transport, so activation order does
  not change what it observes.
- Inactive by default: leaving it out costs nothing at runtime.

#### `metrics`

Metrics.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |

| Option | Type |
|---|---|
| `now` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.metrics.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Attaches to pipeline hooks, not the transport, so activation order does
  not change what it observes.
- Inactive by default: leaving it out costs nothing at runtime.

#### `paging`

Paging.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `afterVar` | `'after'` |
| `cursorParam` | `'cursor'` |
| `firstVar` | `'first'` |
| `limitParam` | `'limit'` |
| `pageParam` | `'page'` |
| `startPage` | `1` |

| Option | Type |
|---|---|
| `limit` | number |
| `ops` | list |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.paging.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Attaches to pipeline hooks, not the transport, so activation order does
  not change what it observes.
- Inactive by default: leaving it out costs nothing at runtime.

#### `ratelimit`

Rate limiting.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `burst` | `5` |
| `rate` | `5` |

| Option | Type |
|---|---|
| `now` | function |
| `sleep` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.ratelimit.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Wraps the transport: its place in the activation order decides what it
  sees. See [Ordering](#ordering) above.
- Inactive by default: leaving it out costs nothing at runtime.

#### `retry`

Retry.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `factor` | `2` |
| `maxDelay` | `2000` |
| `minDelay` | `50` |
| `retries` | `2` |
| `statuses` | `[408, 425, 429, 500, 502, 503, 504]` |

| Option | Type |
|---|---|
| `jitter` | boolean |
| `sleep` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.retry.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Wraps the transport: its place in the activation order decides what it
  sees. See [Ordering](#ordering) above.
- Inactive by default: leaving it out costs nothing at runtime.

#### `test`

Test transport.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |

| Option | Type |
|---|---|
| `entity` | map |
| `net` | map |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.test.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Attaches to pipeline hooks, not the transport, so activation order does
  not change what it observes.
- Installs the BASE transport that the wrapping features wrap, so it must be
  activated before them.
- Inactive by default: leaving it out costs nothing at runtime.

#### `timeout`

Timeout.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `ms` | `30000` |

| Option | Type |
|---|---|
| `clearTimer` | function |
| `setTimer` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.timeout.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Wraps the transport: its place in the activation order decides what it
  sees. See [Ordering](#ordering) above.
- Inactive by default: leaving it out costs nothing at runtime.

