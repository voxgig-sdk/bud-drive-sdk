# BudDrive PHP SDK Reference

Complete API reference for the BudDrive PHP SDK.


## BudDriveSDK

### Constructor

```php
require_once __DIR__ . '/buddrive_sdk.php';

$client = new BudDriveSDK($options);
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$options` | `array` | SDK configuration options. |
| `$options["apikey"]` | `string` | API key for authentication. |
| `$options["base"]` | `string` | Base URL for API requests. |
| `$options["prefix"]` | `string` | URL prefix appended after base. |
| `$options["suffix"]` | `string` | URL suffix appended after path. |
| `$options["headers"]` | `array` | Custom headers for all requests. |
| `$options["feature"]` | `array` | Feature configuration. |
| `$options["system"]` | `array` | System overrides (e.g. custom fetch). |


### Static Methods

#### `BudDriveSDK::test($testopts = null, $sdkopts = null)`

Create a test client with mock features active. Both arguments may be `null`.

```php
$client = BudDriveSDK::test();
```


### Instance Methods

#### `DriveCampaignsApi($data = null)`

Create a new `DriveCampaignsApiEntity` instance. Pass `null` for no initial data.

#### `DriveMcpApi($data = null)`

Create a new `DriveMcpApiEntity` instance. Pass `null` for no initial data.

#### `DriveSegmentsApi($data = null)`

Create a new `DriveSegmentsApiEntity` instance. Pass `null` for no initial data.

#### `options_map(): array`

Return a deep copy of the current SDK options.

#### `get_utility(): BudDriveUtility`

Return a copy of the SDK utility object.

#### `direct(array $fetchargs = []): array`

Make a direct HTTP request to any API endpoint. This is the raw-HTTP escape
hatch: it does **not** throw. It returns a result array
`["ok" => bool, "status" => int, "headers" => array, "data" => mixed]`, or
`["ok" => false, "err" => \Exception]` on failure. Branch on `$result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `$fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `$fetchargs["params"]` | `array` | Path parameter values for `{param}` substitution. |
| `$fetchargs["query"]` | `array` | Query string parameters. |
| `$fetchargs["headers"]` | `array` | Request headers (merged with defaults). |
| `$fetchargs["body"]` | `mixed` | Request body (arrays are JSON-serialized). |
| `$fetchargs["ctrl"]` | `array` | Control options. |

**Returns:** `array` — the result dict (see above); never throws.

#### `prepare(array $fetchargs = []): mixed`

Prepare a fetch definition without sending the request. Returns the
`$fetchdef` array. Throws on error.


---

## DriveCampaignsApiEntity

```php
$drive_campaigns_api = $client->DriveCampaignsApi();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `audience` | `string` | No | The ID of the segment this campaign targets. |
| `branding` | `string` | No | Overrides your configured value for this run only. |
| `campaign_goal` | `string` | Yes | The business goal the campaign should serve, in plain language. |
| `campaign_id` | `string` | Yes |  |
| `channels` | `array` | No |  |
| `completed_at` | `string` | No |  |
| `context` | `string` | No | Overrides your configured value for this run only. |
| `created_at` | `string` | No |  |
| `created_by` | `string` | No |  |
| `description` | `string` | No |  |
| `duration` | `array` | No |  |
| `examples` | `array` | No |  |
| `extra_context` | `string` | No | Overrides your configured value for this run only. |
| `group` | `string` | No |  |
| `headline` | `array` | No |  |
| `id` | `string` | No |  |
| `image_style` | `string` | No | Overrides your configured value for this run only. |
| `image_url` | `string` | No |  |
| `label` | `string` | No |  |
| `message` | `string` | No |  |
| `name` | `string` | No |  |
| `parameters` | `array` | No |  |
| `primary_button` | `array` | No |  |
| `priority_order` | `int` | No |  |
| `result` | `array` | No | The campaign a completed run produced. |
| `schedule` | `array` | No | When the campaign is published. |
| `started_at` | `string` | No |  |
| `status` | `string` | No | The campaign's current status. |
| `styling` | `array` | No |  |
| `template_id` | `string` | No |  |
| `template_variables` | `array` | No |  |
| `templates` | `array` | No |  |
| `tone_of_voice` | `string` | No | Overrides your configured value for this run only. |
| `trend` | `array` | No |  |
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

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->DriveCampaignsApi()->create([
  "campaign_goal" => null, // string
  "campaign_id" => null, // string
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->DriveCampaignsApi()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->DriveCampaignsApi()->load(["campaign_id" => "campaign_id"]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->DriveCampaignsApi()->remove(["campaign_id" => "campaign_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->DriveCampaignsApi()->update([
  "campaign_id" => "campaign_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): DriveCampaignsApiEntity`

Create a new `DriveCampaignsApiEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## DriveMcpApiEntity

```php
$drive_mcp_api = $client->DriveMcpApi();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `error` | `array` | No |  |
| `id` | `string` | Yes |  |
| `jsonrpc` | `string` | Yes |  |
| `method` | `string` | Yes |  |
| `params` | `array` | No |  |
| `result` | `array` | No |  |

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

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->DriveMcpApi()->create([
  "id" => null, // string
  "jsonrpc" => null, // string
  "method" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): DriveMcpApiEntity`

Create a new `DriveMcpApiEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## DriveSegmentsApiEntity

```php
$drive_segments_api = $client->DriveSegmentsApi();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | No |  |
| `created_by` | `string` | No |  |
| `criteria` | `mixed` | Yes |  |
| `criteria_from` | `string` | Yes | The ID of the segment whose criteria should be copied. |
| `customers` | `array` | Yes | Customer IDs to include in the fixed customer list. |
| `customers_from` | `string` | Yes | The ID of the segment whose current members should be copied. |
| `description` | `string` | No |  |
| `fixed_customer_list` | `bool` | No |  |
| `fixed_customer_list_size` | `int` | No |  |
| `from_criteria_id` | `string` | No |  |
| `id` | `string` | Yes |  |
| `intersection_count` | `int` | No |  |
| `jaccard_index` | `float` | No |  |
| `name` | `string` | Yes |  |
| `parameters` | `array` | No |  |
| `source` | `string` | No | How the segment was created. |
| `statistics_id` | `string` | No |  |
| `suggestion` | `array` | No |  |
| `tag` | `string` | No |  |
| `tags` | `array` | No |  |
| `type` | `string` | No | The shape of the statistic's values. |
| `union_count` | `int` | No |  |
| `upload` | `array` | No |  |
| `values` | `array` | No |  |

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

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->DriveSegmentsApi()->create([
  "criteria" => null, // mixed
  "criteria_from" => null, // string
  "customers" => null, // array
  "customers_from" => null, // string
  "id" => null, // string
  "name" => null, // string
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->DriveSegmentsApi()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->DriveSegmentsApi()->load(["segment_id" => "segment_id"]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->DriveSegmentsApi()->remove(["segment_id" => "segment_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->DriveSegmentsApi()->update([
  "criteria_id" => "criteria_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): DriveSegmentsApiEntity`

Create a new `DriveSegmentsApiEntity` instance with the same client and
options.

#### `get_name(): string`

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

```php
$client = new BudDriveSDK([
  "feature" => [
    "debug" => ["active" => true],
    "idempotency" => ["active" => true],
    "metrics" => ["active" => true],
    "paging" => ["active" => true],
    "ratelimit" => ["active" => true],
    "retry" => ["active" => true],
    "test" => ["active" => true],
    "timeout" => ["active" => true],
  ],
]);
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

