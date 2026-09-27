# BudDrive Lua SDK Reference

Complete API reference for the BudDrive Lua SDK.


## BudDriveSDK

### Constructor

```lua
local sdk = require("bud-drive_sdk")
local client = sdk.new(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `table` | SDK configuration options. |
| `options.apikey` | `string` | API key for authentication. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `table` | Custom headers for all requests. |
| `options.feature` | `table` | Feature configuration. |
| `options.system` | `table` | System overrides (e.g. custom fetch). |


### Static Methods

#### `sdk.test(testopts?, sdkopts?)`

Create a test client with mock features active. Both arguments are optional.

```lua
local client = sdk.test()
```


### Instance Methods

#### `DriveCampaignsApi(data)`

Create a new `DriveCampaignsApi` entity instance. Pass `nil` for no initial data.

#### `DriveMcpApi(data)`

Create a new `DriveMcpApi` entity instance. Pass `nil` for no initial data.

#### `DriveSegmentsApi(data)`

Create a new `DriveSegmentsApi` entity instance. Pass `nil` for no initial data.

#### `options_map() -> table`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs) -> table, err`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs.path` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs.method` | `string` | HTTP method (default: `"GET"`). |
| `fetchargs.params` | `table` | Path parameter values for `{param}` substitution. |
| `fetchargs.query` | `table` | Query string parameters. |
| `fetchargs.headers` | `table` | Request headers (merged with defaults). |
| `fetchargs.body` | `any` | Request body (tables are JSON-serialized). |
| `fetchargs.ctrl` | `table` | Control options (e.g. `{ explain = true }`). |

**Returns:** `table, err`

#### `prepare(fetchargs) -> table, err`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `table, err`


---

## DriveCampaignsApiEntity

```lua
local drive_campaigns_api = client:DriveCampaignsApi(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `audience` | `string` | No | The ID of the segment this campaign targets. |
| `branding` | `string` | No | Overrides your configured value for this run only. |
| `campaign_goal` | `string` | Yes | The business goal the campaign should serve, in plain language. |
| `campaign_id` | `string` | Yes |  |
| `channels` | `table` | No |  |
| `completed_at` | `string` | No |  |
| `context` | `string` | No | Overrides your configured value for this run only. |
| `created_at` | `string` | No |  |
| `created_by` | `string` | No |  |
| `description` | `string` | No |  |
| `duration` | `table` | No |  |
| `examples` | `table` | No |  |
| `extra_context` | `string` | No | Overrides your configured value for this run only. |
| `group` | `string` | No |  |
| `headline` | `table` | No |  |
| `id` | `string` | No |  |
| `image_style` | `string` | No | Overrides your configured value for this run only. |
| `image_url` | `string` | No |  |
| `label` | `string` | No |  |
| `message` | `string` | No |  |
| `name` | `string` | No |  |
| `parameters` | `table` | No |  |
| `primary_button` | `table` | No |  |
| `priority_order` | `number` | No |  |
| `result` | `table` | No | The campaign a completed run produced. |
| `schedule` | `table` | No | When the campaign is published. |
| `started_at` | `string` | No |  |
| `status` | `string` | No | The campaign's current status. |
| `styling` | `table` | No |  |
| `template_id` | `string` | No |  |
| `template_variables` | `table` | No |  |
| `templates` | `table` | No |  |
| `tone_of_voice` | `string` | No | Overrides your configured value for this run only. |
| `trend` | `table` | No |  |
| `type` | `string` | No |  |

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

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:DriveCampaignsApi():create({
  campaign_goal = --[[ string ]],
  campaign_id = --[[ string ]],
})
```

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:DriveCampaignsApi():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:DriveCampaignsApi():load({ campaign_id = "campaign_id" })
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:DriveCampaignsApi():remove({ campaign_id = "campaign_id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:DriveCampaignsApi():update({
  campaign_id = "campaign_id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `DriveCampaignsApiEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## DriveMcpApiEntity

```lua
local drive_mcp_api = client:DriveMcpApi(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `error` | `table` | No |  |
| `id` | `string` | Yes |  |
| `jsonrpc` | `string` | Yes |  |
| `method` | `string` | Yes |  |
| `params` | `table` | No |  |
| `result` | `table` | No |  |

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

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:DriveMcpApi():create({
  id = --[[ string ]],
  jsonrpc = --[[ string ]],
  method = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `DriveMcpApiEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## DriveSegmentsApiEntity

```lua
local drive_segments_api = client:DriveSegmentsApi(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | No |  |
| `created_by` | `string` | No |  |
| `criteria` | `any` | Yes |  |
| `criteria_from` | `string` | Yes | The ID of the segment whose criteria should be copied. |
| `customers` | `table` | Yes | Customer IDs to include in the fixed customer list. |
| `customers_from` | `string` | Yes | The ID of the segment whose current members should be copied. |
| `description` | `string` | No |  |
| `fixed_customer_list` | `boolean` | No |  |
| `fixed_customer_list_size` | `number` | No |  |
| `from_criteria_id` | `string` | No |  |
| `id` | `string` | Yes |  |
| `intersection_count` | `number` | No |  |
| `jaccard_index` | `number` | No |  |
| `name` | `string` | Yes |  |
| `parameters` | `table` | No |  |
| `source` | `string` | No | How the segment was created. |
| `statistics_id` | `string` | No |  |
| `suggestion` | `table` | No |  |
| `tag` | `string` | No |  |
| `tags` | `table` | No |  |
| `type` | `string` | No | The shape of the statistic's values. |
| `union_count` | `number` | No |  |
| `upload` | `table` | No |  |
| `values` | `table` | No |  |

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

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:DriveSegmentsApi():create({
  criteria = --[[ any ]],
  criteria_from = --[[ string ]],
  customers = --[[ table ]],
  customers_from = --[[ string ]],
  id = --[[ string ]],
  name = --[[ string ]],
})
```

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:DriveSegmentsApi():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:DriveSegmentsApi():load({ segment_id = "segment_id" })
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:DriveSegmentsApi():remove({ segment_id = "segment_id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:DriveSegmentsApi():update({
  criteria_id = "criteria_id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `DriveSegmentsApiEntity` instance with the same client and
options.

#### `get_name() -> string`

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

```lua
local client = sdk.new({
  feature = {
    debug = { active = true },
    idempotency = { active = true },
    metrics = { active = true },
    paging = { active = true },
    ratelimit = { active = true },
    retry = { active = true },
    test = { active = true },
    timeout = { active = true },
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

