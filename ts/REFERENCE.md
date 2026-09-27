# BudDrive TypeScript SDK Reference

Complete API reference for the BudDrive TypeScript SDK.


## BudDriveSDK

### Constructor

```ts
new BudDriveSDK(options?: object)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `object` | SDK configuration options. |
| `options.apikey` | `string` | API key for authentication. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `object` | Custom headers for all requests. |
| `options.feature` | `object` | Feature configuration. |
| `options.system` | `object` | System overrides (e.g. custom fetch). |


### Static Methods

#### `BudDriveSDK.test(testopts?, sdkopts?)`

Create a test client with mock features active.

```ts
const client = BudDriveSDK.test()
```

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `testopts` | `object` | Test feature options. |
| `sdkopts` | `object` | Additional SDK options merged with test defaults. |

**Returns:** `BudDriveSDK` instance in test mode.


### Instance Methods

#### `DriveCampaignsApi(data?: object)`

Create a new `DriveCampaignsApi` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `DriveCampaignsApiEntity` instance.

#### `DriveMcpApi(data?: object)`

Create a new `DriveMcpApi` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `DriveMcpApiEntity` instance.

#### `DriveSegmentsApi(data?: object)`

Create a new `DriveSegmentsApi` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `DriveSegmentsApiEntity` instance.

#### `options()`

Return a deep copy of the current SDK options.

**Returns:** `object`

#### `utility()`

Return a copy of the SDK utility object.

**Returns:** `object`

#### `direct(fetchargs?: object)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs.path` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs.method` | `string` | HTTP method (default: `GET`). |
| `fetchargs.params` | `object` | Path parameter values for `{param}` substitution. |
| `fetchargs.query` | `object` | Query string parameters. |
| `fetchargs.headers` | `object` | Request headers (merged with defaults). |
| `fetchargs.body` | `any` | Request body (objects are JSON-serialized). |
| `fetchargs.ctrl` | `object` | Control options (e.g. `{ explain: true }`). |

**Returns:** `Promise<{ ok, status, headers, data } | Error>`

#### `prepare(fetchargs?: object)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `Promise<{ url, method, headers, body } | Error>`

#### `tester(testopts?, sdkopts?)`

Alias for `BudDriveSDK.test()`.

**Returns:** `BudDriveSDK` instance in test mode.


---

## DriveCampaignsApiEntity

```ts
const drive_campaigns_api = client.DriveCampaignsApi()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `audience` | `string` | No | The ID of the segment this campaign targets. |
| `branding` | `string` | No | Overrides your configured value for this run only. |
| `campaign_goal` | `string` | Yes | The business goal the campaign should serve, in plain language. |
| `campaign_id` | `string` | Yes |  |
| `channels` | `Record<string, any>` | No |  |
| `completed_at` | `string` | No |  |
| `context` | `string` | No | Overrides your configured value for this run only. |
| `created_at` | `string` | No |  |
| `created_by` | `string` | No |  |
| `description` | `string` | No |  |
| `duration` | `Record<string, any>` | No |  |
| `examples` | `any[]` | No |  |
| `extra_context` | `string` | No | Overrides your configured value for this run only. |
| `group` | `string` | No |  |
| `headline` | `Record<string, any>` | No |  |
| `id` | `string` | No |  |
| `image_style` | `string` | No | Overrides your configured value for this run only. |
| `image_url` | `string` | No |  |
| `label` | `string` | No |  |
| `message` | `string` | No |  |
| `name` | `string` | No |  |
| `parameters` | `any[]` | No |  |
| `primary_button` | `Record<string, any>` | No |  |
| `priority_order` | `number` | No |  |
| `result` | `Record<string, any>` | No | The campaign a completed run produced. |
| `schedule` | `Record<string, any>` | No | When the campaign is published. |
| `started_at` | `string` | No |  |
| `status` | `string` | No | The campaign's current status. |
| `styling` | `Record<string, any>` | No |  |
| `template_id` | `string` | No |  |
| `template_variables` | `Record<string, any>` | No |  |
| `templates` | `any[]` | No |  |
| `tone_of_voice` | `string` | No | Overrides your configured value for this run only. |
| `trend` | `Record<string, any>` | No |  |
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

### Actions

This entity exposes custom API actions in addition to the standard
operations. Select one with `$action` in the call's argument; the
remaining keys are sent as that action's payload.

| Action | Route | Call |
| --- | --- | --- |
| `content` | `/drive-api/v1/campaigns/{campaign_id}/content` | `client.DriveCampaignsApi().load({ $action: 'content', ... })` |
| `content` | `/drive-api/v1/campaigns/{campaign_id}/content` | `client.DriveCampaignsApi().update({ $action: 'content', ... })` |
| `status` | `/drive-api/v1/campaigns/{campaign_id}/status` | `client.DriveCampaignsApi().update({ $action: 'status', ... })` |

An action returns that action's OWN response, which is not necessarily a
DriveCampaignsApi record — check the API definition for its shape.

```ts
const result = await client.DriveCampaignsApi().load({
  $action: 'content',
  /* ...the action's own arguments */
})
```

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.DriveCampaignsApi().create({
  campaign_goal: 'example_campaign_goal',
  campaign_id: 'example_campaign_id',
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.DriveCampaignsApi().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.DriveCampaignsApi().load({ campaign_id: 'campaign_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.DriveCampaignsApi().remove({ campaign_id: 'campaign_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.DriveCampaignsApi().update({
  campaign_id: 'campaign_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `DriveCampaignsApiEntity` instance with the same client and
options.

#### `client()`

Return the parent `BudDriveSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## DriveMcpApiEntity

```ts
const drive_mcp_api = client.DriveMcpApi()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `error` | `Record<string, any>` | No |  |
| `id` | `string` | Yes |  |
| `jsonrpc` | `string` | Yes |  |
| `method` | `string` | Yes |  |
| `params` | `Record<string, any>` | No |  |
| `result` | `Record<string, any>` | No |  |

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

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.DriveMcpApi().create({
  id: 'example_id',
  jsonrpc: 'example_jsonrpc',
  method: 'example_method',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `DriveMcpApiEntity` instance with the same client and
options.

#### `client()`

Return the parent `BudDriveSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## DriveSegmentsApiEntity

```ts
const drive_segments_api = client.DriveSegmentsApi()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | No |  |
| `created_by` | `string` | No |  |
| `criteria` | `any` | Yes |  |
| `criteria_from` | `string` | Yes | The ID of the segment whose criteria should be copied. |
| `customers` | `any[]` | Yes | Customer IDs to include in the fixed customer list. |
| `customers_from` | `string` | Yes | The ID of the segment whose current members should be copied. |
| `description` | `string` | No |  |
| `fixed_customer_list` | `boolean` | No |  |
| `fixed_customer_list_size` | `number` | No |  |
| `from_criteria_id` | `string` | No |  |
| `id` | `string` | Yes |  |
| `intersection_count` | `number` | No |  |
| `jaccard_index` | `number` | No |  |
| `name` | `string` | Yes |  |
| `parameters` | `Record<string, any>` | No |  |
| `source` | `string` | No | How the segment was created. |
| `statistics_id` | `string` | No |  |
| `suggestion` | `Record<string, any>` | No |  |
| `tag` | `string` | No |  |
| `tags` | `any[]` | No |  |
| `type` | `string` | No | The shape of the statistic's values. |
| `union_count` | `number` | No |  |
| `upload` | `Record<string, any>` | No |  |
| `values` | `any[]` | No |  |

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

### Actions

This entity exposes custom API actions in addition to the standard
operations. Select one with `$action` in the call's argument; the
remaining keys are sent as that action's payload.

| Action | Route | Call |
| --- | --- | --- |
| `query` | `/drive-api/v1/segments/{segment_id}/query` | `client.DriveSegmentsApi().create({ $action: 'query', ... })` |
| `run` | `/drive-api/v2/criteria/drafts/{criteria_id}/run` | `client.DriveSegmentsApi().create({ $action: 'run', ... })` |

An action returns that action's OWN response, which is not necessarily a
DriveSegmentsApi record — check the API definition for its shape.

```ts
const result = await client.DriveSegmentsApi().create({
  $action: 'query',
  /* ...the action's own arguments */
})
```

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.DriveSegmentsApi().create({
  criteria: 'example_criteria',
  criteria_from: 'example_criteria_from',
  customers: [],
  customers_from: 'example_customers_from',
  id: 'example_id',
  name: 'example_name',
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.DriveSegmentsApi().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.DriveSegmentsApi().load({ segment_id: 'segment_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.DriveSegmentsApi().remove({ segment_id: 'segment_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.DriveSegmentsApi().update({
  criteria_id: 'criteria_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `DriveSegmentsApiEntity` instance with the same client and
options.

#### `client()`

Return the parent `BudDriveSDK` instance.

#### `entopts()`

Return a copy of the entity options.


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

```ts
const client = new BudDriveSDK({
  feature: {
    debug: { active: true },
    idempotency: { active: true },
    metrics: { active: true },
    paging: { active: true },
    ratelimit: { active: true },
    retry: { active: true },
    test: { active: true },
    timeout: { active: true },
  }
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

