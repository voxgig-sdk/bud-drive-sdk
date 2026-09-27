# BudDrive TypeScript SDK



The TypeScript SDK for the BudDrive API — a type-safe, entity-oriented client with full async/await support.

The API is exposed as capitalised, semantic **Entities** — e.g.
`client.DriveCampaignsApi()` — each with a small set of operations (`list`, `load`, `create`, `update`, `remove`, `patch`)
instead of raw URL paths and query parameters. This keeps the surface
predictable and low-friction for both humans and AI agents.

> Also generated from this model: `go`, `go-cli`, `go-mcp`, `js`, `lua`, `php`, `py` — see
> the [top-level README](../README.md).


## Install
This package is not yet published to npm. Install it from the GitHub
release tag (`ts/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/bud-drive-sdk/releases](https://github.com/voxgig-sdk/bud-drive-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```ts
import { BudDriveSDK } from '@voxgig-sdk/bud-drive-sdk'

const client = new BudDriveSDK({
  apikey: process.env.BUD_DRIVE_APIKEY,
})
```

### 2. List drivecampaignsapi records

`list()` resolves to an array of DriveCampaignsApi ENTITIES — every operation
resolves to entities, not raw records. Iterate them directly, and call
`.data()` on one for the record it holds:

```ts
const drivecampaignsapis = await client.DriveCampaignsApi().list()

for (const drivecampaignsapi of drivecampaignsapis) {
  console.log(drivecampaignsapi)
}
```

### 3. Load a drivecampaignsapi

`load()` returns the entity directly and throws on failure:

```ts
try {
  const drivecampaignsapi = await client.DriveCampaignsApi().load({ campaign_id: 'example_campaign_id' })
  console.log(drivecampaignsapi)
} catch (err) {
  console.error('load failed:', err)
}
```

### 4. Create, update, and remove

```ts
// Create — returns the created DriveCampaignsApi ENTITY (.data() for the record)
const created = await client.DriveCampaignsApi().create({
  campaign_goal: 'example_campaign_goal',
  campaign_id: 'example_campaign_id',
})

// Update
const updated = await client.DriveCampaignsApi().update({
  campaign_id: 'example_campaign_id',
  audience: 'example_audience',
})

// Remove
await client.DriveCampaignsApi().remove({
  campaign_id: 'example_campaign_id',
})
```


## Error handling

Entity operations reject on failure, so wrap them in `try` / `catch`:

```ts
try {
  const drivecampaignsapis = await client.DriveCampaignsApi().list()
  console.log(drivecampaignsapis)
} catch (err) {
  console.error('list failed:', err)
}
```

The low-level `direct()` method does **not** throw — it returns the
value or an `Error`, so check the result before using it:

```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example_id' },
})

if (result instanceof Error) {
  throw result
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example' },
})

if (result instanceof Error) {
  throw result
}
if (result.ok) {
  console.log(result.status)  // 200
  console.log(result.data)    // response body
}
```

### Prepare a request without sending it

```ts
const fetchdef = await client.prepare({
  path: '/api/resource/{id}',
  method: 'DELETE',
  params: { id: 'example' },
})

// Inspect before sending
console.log(fetchdef.url)
console.log(fetchdef.method)
console.log(fetchdef.headers)
```

### Use test mode

Create a mock client for unit testing — no server required:

```ts
const client = BudDriveSDK.test()

const drivecampaignsapi = await client.DriveCampaignsApi().list()
// drivecampaignsapi is the entity, populated with mock response data
// — call drivecampaignsapi.data() for the record itself
console.log(drivecampaignsapi)
```

You can also use the instance method:

```ts
const client = new BudDriveSDK({ apikey: '...' })
const testClient = client.tester()
```

### Retain entity state across calls

Entity instances remember their last match and data:

```ts
const entity = client.DriveCampaignsApi()

// First call runs the operation and stores its result
await entity.list()

// Subsequent calls reuse the stored state
const data = entity.data()
console.log(data.id)
```

### Add custom middleware

Pass features via the `extend` option:

```ts
const logger = {
  hooks: {
    PreRequest: (ctx: any) => {
      console.log('Requesting:', ctx.spec.method, ctx.spec.path)
    },
    PreResponse: (ctx: any) => {
      console.log('Status:', ctx.out.request?.status)
    },
  },
}

const client = new BudDriveSDK({
  apikey: '...',
  extend: [logger],
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
cd ts && npm test
```

Live entity tests continue independent operations after errors and attempt
supported cleanup. Their final result reports failures and missing prerequisites
after the remaining work completes. The model and test inputs determine which
API operations the generated scenarios cover.


## Reference

### BudDriveSDK

#### Constructor

```ts
new BudDriveSDK(options?: {
  apikey?: string
  base?: string
  prefix?: string
  suffix?: string
  feature?: Record<string, { active: boolean }>
  extend?: Feature[]
})
```

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `string` | API key for authentication. |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `object` | Feature activation flags (e.g. `{ test: { active: true } }`). |
| `extend` | `Feature[]` | Additional feature instances to load. |

#### Methods

| Method | Returns | Description |
| --- | --- | --- |
| `options()` | `object` | Deep copy of current SDK options. |
| `utility()` | `Utility` | Deep copy of the SDK utility object. |
| `prepare(fetchargs?)` | `Promise<FetchDef>` | Build an HTTP request definition without sending it. |
| `direct(fetchargs?)` | `Promise<DirectResult>` | Build and send an HTTP request. |
| `DriveCampaignsApi(data?)` | `DriveCampaignsApiEntity` | Create a DriveCampaignsApi entity instance. |
| `DriveMcpApi(data?)` | `DriveMcpApiEntity` | Create a DriveMcpApi entity instance. |
| `DriveSegmentsApi(data?)` | `DriveSegmentsApiEntity` | Create a DriveSegmentsApi entity instance. |
| `tester(testopts?, sdkopts?)` | `BudDriveSDK` | Create a test-mode client instance. |

#### Static methods

| Method | Returns | Description |
| --- | --- | --- |
| `BudDriveSDK.test(testopts?, sdkopts?)` | `BudDriveSDK` | Create a test-mode client. |

### Entity interface

All entities share the same interface.

#### Methods

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `load(reqmatch?, ctrl?): Promise<Entity>` | Load a single entity by match criteria. |
| `list` | `list(reqmatch?, ctrl?): Promise<Entity[]>` | List entities matching the criteria. |
| `create` | `create(reqdata?, ctrl?): Promise<Entity>` | Create a new entity. |
| `update` | `update(reqdata?, ctrl?): Promise<Entity>` | Update an existing entity. |
| `remove` | `remove(reqmatch?, ctrl?): Promise<void>` | Remove an entity. |
| `data` | `data(data?: Partial<Entity>): Entity` | Get or set entity data. |
| `match` | `match(match?: Partial<Entity>): Partial<Entity>` | Get or set entity match criteria. |
| `make` | `make(): Entity` | Create a new instance with the same options. |
| `client` | `client(): BudDriveSDK` | Return the parent SDK client. |
| `entopts` | `entopts(): object` | Return a copy of the entity options. |

#### Return values

Entity operations resolve to the entity data directly — there is no
result envelope:

- `load`, `create` and `update` resolve to a single entity object.
- `list` resolves to an **array** of entity objects (iterate it directly;
  there is no `.data` and no `.ok`).
- `remove` resolves to `void`.

On a failed request these methods **throw**, so wrap calls in
`try`/`catch` to handle errors. Only `direct()` returns the result
envelope described below.

### DirectResult shape

The `direct()` method returns:

```ts
{
  ok: boolean
  status: number
  headers: object
  data: any
}
```

On error, `ok` is `false` and an `err` property contains the error.

### FetchDef shape

The `prepare()` method returns:

```ts
{
  url: string
  method: string
  headers: Record<string, string>
  body?: any
}
```

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

Operations: create, list, load, remove, update.

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

Operations: create.

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

Operations: create, list, load, patch, remove, update.

API path: `/drive-api/v1/segments/{segment_id}/query`



## Entities


### DriveCampaignsApi

Create an instance: `const drive_campaigns_api = client.DriveCampaignsApi()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `audience` | `string` | The ID of the segment this campaign targets. |
| `branding` | `string` | Overrides your configured value for this run only. |
| `campaign_goal` | `string` | The business goal the campaign should serve, in plain language. |
| `campaign_id` | `string` |  |
| `channels` | `Record<string, any>` |  |
| `completed_at` | `string` |  |
| `context` | `string` | Overrides your configured value for this run only. |
| `created_at` | `string` |  |
| `created_by` | `string` |  |
| `description` | `string` |  |
| `duration` | `Record<string, any>` |  |
| `examples` | `any[]` |  |
| `extra_context` | `string` | Overrides your configured value for this run only. |
| `group` | `string` |  |
| `headline` | `Record<string, any>` |  |
| `id` | `string` |  |
| `image_style` | `string` | Overrides your configured value for this run only. |
| `image_url` | `string` |  |
| `label` | `string` |  |
| `message` | `string` |  |
| `name` | `string` |  |
| `parameters` | `any[]` |  |
| `primary_button` | `Record<string, any>` |  |
| `priority_order` | `number` |  |
| `result` | `Record<string, any>` | The campaign a completed run produced. |
| `schedule` | `Record<string, any>` | When the campaign is published. |
| `started_at` | `string` |  |
| `status` | `string` | The campaign's current status. |
| `styling` | `Record<string, any>` |  |
| `template_id` | `string` |  |
| `template_variables` | `Record<string, any>` |  |
| `templates` | `any[]` |  |
| `tone_of_voice` | `string` | Overrides your configured value for this run only. |
| `trend` | `Record<string, any>` |  |
| `type` | `string` |  |

#### Example: Load

```ts
const drive_campaigns_api = await client.DriveCampaignsApi().load({ campaign_id: 'campaign_id' })
```

#### Example: List

```ts
const drive_campaigns_apis = await client.DriveCampaignsApi().list()
```

#### Example: Create

```ts
const drive_campaigns_api = await client.DriveCampaignsApi().create({
  campaign_goal: 'example_campaign_goal',
  campaign_id: 'example_campaign_id',
})
```


### DriveMcpApi

Create an instance: `const drive_mcp_api = client.DriveMcpApi()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `error` | `Record<string, any>` |  |
| `id` | `string` |  |
| `jsonrpc` | `string` |  |
| `method` | `string` |  |
| `params` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |

#### Example: Create

```ts
const drive_mcp_api = await client.DriveMcpApi().create({
  id: 'example_id',
  jsonrpc: 'example_jsonrpc',
  method: 'example_method',
})
```


### DriveSegmentsApi

Create an instance: `const drive_segments_api = client.DriveSegmentsApi()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` |  |
| `created_by` | `string` |  |
| `criteria` | `any` |  |
| `criteria_from` | `string` | The ID of the segment whose criteria should be copied. |
| `customers` | `any[]` | Customer IDs to include in the fixed customer list. |
| `customers_from` | `string` | The ID of the segment whose current members should be copied. |
| `description` | `string` |  |
| `fixed_customer_list` | `boolean` |  |
| `fixed_customer_list_size` | `number` |  |
| `from_criteria_id` | `string` |  |
| `id` | `string` |  |
| `intersection_count` | `number` |  |
| `jaccard_index` | `number` |  |
| `name` | `string` |  |
| `parameters` | `Record<string, any>` |  |
| `source` | `string` | How the segment was created. |
| `statistics_id` | `string` |  |
| `suggestion` | `Record<string, any>` |  |
| `tag` | `string` |  |
| `tags` | `any[]` |  |
| `type` | `string` | The shape of the statistic's values. |
| `union_count` | `number` |  |
| `upload` | `Record<string, any>` |  |
| `values` | `any[]` |  |

#### Example: Load

```ts
const drive_segments_api = await client.DriveSegmentsApi().load({ segment_id: 'segment_id' })
```

#### Example: List

```ts
const drive_segments_apis = await client.DriveSegmentsApi().list()
```

#### Example: Create

```ts
const drive_segments_api = await client.DriveSegmentsApi().create({
  criteria: 'example_criteria',
  criteria_from: 'example_criteria_from',
  customers: [],
  customers_from: 'example_customers_from',
  id: 'example_id',
  name: 'example_name',
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

Features are the extension mechanism. A feature is an object with a
`hooks` map. Each hook key is a pipeline stage name, and the value is
a function that receives the context.

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

### Module structure

```
bud-drive/
├── src/
│   ├── BudDriveSDK.ts        # Main SDK class
│   ├── entity/             # Entity implementations
│   ├── feature/            # Built-in features (Base, Test, Log)
│   └── utility/            # Utility functions
├── test/                   # Test suites
└── dist/                   # Compiled output
```

Import the SDK from the package root:

```ts
import { BudDriveSDK } from '@voxgig-sdk/bud-drive-sdk'
```

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally. Subsequent
calls on the same instance can rely on this state.

```ts
const drivecampaignsapi = client.DriveCampaignsApi()
await drivecampaignsapi.list()

// drivecampaignsapi.data() now returns the drivecampaignsapi data from the last `list`
// drivecampaignsapi.match() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

The `direct` method gives full control over the HTTP request. Use it
for non-standard endpoints, bulk operations, or any path not modelled
as an entity. The `prepare` method is useful for debugging — it
shows exactly what `direct` would send.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
