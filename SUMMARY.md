# Drive API

Drive lets you understand your customers using the flags it compiles from data across every source you bring in, not just transactions, build segments from those flags, and run personalised campaigns against them. Drive works in a loop: **understand** your customers (via the flags that make up Drive&#39;s criteria), **segment** them (Drive Segments API), and **engage** them (Drive Campaigns API). Measuring a campaign&#39;s performance feeds back into segmentation, so the next audience is sharper than the last. Campaigns don&#39;t have to be built step by step, either: describe the outcome you want in plain language and Drive&#39;s campaign agent finds the matching audience, writes the copy, generates the artwork, and estimates the opportunity for you. Segments V2 endpoints return non-standard vendor content types (`application/vnd.bud.drive.notification+json`, `.criteria+json`, `.visualisation+json`, `.criteria-options+json`) rather than plain `application/json`, check the `Content-Type` of these responses rather than assuming JSON.

## Start here

This guide introduces the API, the client libraries, and the companion tools in this repository. Start with the API capabilities, choose a client for your application, and use the linked reference when you need exact request and response details.

The selected API surface contains 3 entities and 40 HTTP routes. There are 6 SDK targets and 2 companion tools.

An entity groups related API operations. An operation can have several routes with different inputs or authentication requirements. The SDK exposes the entity and its operations using the conventions of the selected language.

## What the API provides

### [DriveCampaignsApi](docs/api/drive_campaigns_api.html)

Results: Created; Accepted; OK; Deleted; No Content; Updated; Status Updated.

SDK operations: `create`, `list`, `load`, `remove`, `update`.

Key fields to recognise:

- `audience`: The ID of the segment this campaign targets.
- `branding`: Overrides your configured value for this run only.
- `campaign_goal`: The business goal the campaign should serve, in plain language.
- `context`: Overrides your configured value for this run only.
- `description`: Customer-facing body copy.

### [DriveMcpApi](docs/api/drive_mcp_api.html)

Results: A JSON-RPC response from the MCP server.

SDK operations: `create`.

### [DriveSegmentsApi](docs/api/drive_segments_api.html)

Results: OK; OK. Response uses the `application/vnd.bud.drive.visualisation+json` media type.; Created. Response uses the `application/vnd.bud.drive.notification+json` media type.; OK. Response uses the `application/vnd.bud.drive.criteria+json` media type.; OK. Response uses the `application/vnd.bud.drive.criteria-options+json` media type.; OK. Response uses the `application/vnd.bud.drive.notification+json` media type.; OK. Response uses the `application/vnd.bud.drive.visualisation+json` media type, updating a draft re-runs it and returns the resulting count.

SDK operations: `create`, `list`, `load`, `patch`, `remove`, `update`.

Key fields to recognise:

- `criteria_from`: The ID of the segment whose criteria should be copied.
- `customers`: Customer IDs to include in the fixed customer list.
- `customers_from`: The ID of the segment whose current members should be copied.
- `source`: How the segment was created. Potential values are subject to change and may be converted to enums at a future unspecified date. | Value | Description | |-----------------|-------------------------------------------------| | `manual` | Built from criteria/groups | | `uploaded_list` | A fixed customer list from an uploaded file |
- `type`: The shape of the statistic&#39;s values. Potential values are subject to change and may be converted to enums at a future unspecified date. | Value | Description | |---------------|---------------------------------------| | `aggregate` | A single summary value | | `time_series` | A series of values over time |

### Route map

Use this map to locate a capability. Consult the entity reference before supplying request data; routes for the same operation can require different fields.

| Entity | SDK operation | HTTP route | Authentication |
| --- | --- | --- | --- |
| [DriveCampaignsApi](docs/api/drive_campaigns_api.html) | `create` | `POST /drive-api/v1/campaigns` | Required |
| [DriveCampaignsApi](docs/api/drive_campaigns_api.html) | `create` | `POST /drive-api/v1/campaigns/agent` | Required |
| [DriveCampaignsApi](docs/api/drive_campaigns_api.html) | `list` | `GET /drive-api/v1/campaigns` | Required |
| [DriveCampaignsApi](docs/api/drive_campaigns_api.html) | `list` | `GET /drive-api/v1/campaigns/banner-widgets` | Required |
| [DriveCampaignsApi](docs/api/drive_campaigns_api.html) | `list` | `GET /drive-api/v1/campaigns/content-variables` | Required |
| [DriveCampaignsApi](docs/api/drive_campaigns_api.html) | `load` | `GET /drive-api/v1/campaigns/{campaign_id}/insights/performance/trends` | Required |
| [DriveCampaignsApi](docs/api/drive_campaigns_api.html) | `load` | `GET /drive-api/v1/campaigns/{campaign_id}` | Required |
| [DriveCampaignsApi](docs/api/drive_campaigns_api.html) | `load` | `GET /drive-api/v1/campaigns/{campaign_id}/content` | Required |
| [DriveCampaignsApi](docs/api/drive_campaigns_api.html) | `load` | `GET /drive-api/v1/campaigns/agent/{generation_id}` | Required |
| [DriveCampaignsApi](docs/api/drive_campaigns_api.html) | `load` | `GET /drive-api/v1/campaign-templates/compiled-templates` | Required |
| [DriveCampaignsApi](docs/api/drive_campaigns_api.html) | `remove` | `DELETE /drive-api/v1/campaigns/{campaign_id}` | Required |
| [DriveCampaignsApi](docs/api/drive_campaigns_api.html) | `remove` | `DELETE /drive-api/v1/campaign-templates/templates/{template_id}` | Required |
| [DriveCampaignsApi](docs/api/drive_campaigns_api.html) | `update` | `PUT /drive-api/v1/campaigns/{campaign_id}/content` | Required |
| [DriveCampaignsApi](docs/api/drive_campaigns_api.html) | `update` | `PUT /drive-api/v1/campaigns/{campaign_id}/status` | Required |
| [DriveMcpApi](docs/api/drive_mcp_api.html) | `create` | `POST /drive-mcp/v1/campaigns` | Required |
| [DriveMcpApi](docs/api/drive_mcp_api.html) | `create` | `POST /drive-mcp/v1/copilot` | Required |
| [DriveMcpApi](docs/api/drive_mcp_api.html) | `create` | `POST /drive-mcp/v1/segments` | Required |
| [DriveSegmentsApi](docs/api/drive_segments_api.html) | `create` | `POST /drive-api/v1/segments/{segment_id}/query` | Required |
| [DriveSegmentsApi](docs/api/drive_segments_api.html) | `create` | `POST /drive-api/v2/criteria/drafts/{criteria_id}/run` | Required |
| [DriveSegmentsApi](docs/api/drive_segments_api.html) | `create` | `POST /drive-api/v2/segments/{segment_id}/criteria/run` | Required |
| [DriveSegmentsApi](docs/api/drive_segments_api.html) | `create` | `POST /drive-api/v1/segments/from-existing/criteria` | Required |
| [DriveSegmentsApi](docs/api/drive_segments_api.html) | `create` | `POST /drive-api/v1/segments/from-existing/fixed-customers` | Required |
| [DriveSegmentsApi](docs/api/drive_segments_api.html) | `create` | `POST /drive-api/v1/segments/upload` | Required |
| [DriveSegmentsApi](docs/api/drive_segments_api.html) | `create` | `POST /drive-api/v2/criteria/drafts` | Required |
| [DriveSegmentsApi](docs/api/drive_segments_api.html) | `create` | `POST /drive-api/v2/criteria/run` | Required |
| [DriveSegmentsApi](docs/api/drive_segments_api.html) | `create` | `POST /drive-api/v2/criteria/suggestions/segment-name` | Required |
| [DriveSegmentsApi](docs/api/drive_segments_api.html) | `create` | `POST /drive-api/v2/segments` | Required |
| [DriveSegmentsApi](docs/api/drive_segments_api.html) | `list` | `GET /drive-api/v1/segments/{segment_id}/statistics/overview` | Required |
| [DriveSegmentsApi](docs/api/drive_segments_api.html) | `list` | `GET /drive-api/v2/segments` | Required |
| [DriveSegmentsApi](docs/api/drive_segments_api.html) | `list` | `GET /drive-api/v1/tags` | Required |
| [DriveSegmentsApi](docs/api/drive_segments_api.html) | `load` | `GET /drive-api/v1/segments/{segment_id}/similarity/{other_segment_id}` | Required |
| [DriveSegmentsApi](docs/api/drive_segments_api.html) | `load` | `GET /drive-api/v2/criteria/drafts/{criteria_id}` | Required |
| [DriveSegmentsApi](docs/api/drive_segments_api.html) | `load` | `GET /drive-api/v2/criteria/options` | Required |
| [DriveSegmentsApi](docs/api/drive_segments_api.html) | `load` | `GET /drive-api/v1/segments/{segment_id}` | Required |
| [DriveSegmentsApi](docs/api/drive_segments_api.html) | `load` | `GET /drive-api/v2/segments/{segment_id}/criteria` | Required |
| [DriveSegmentsApi](docs/api/drive_segments_api.html) | `load` | `GET /drive-api/v2/segments/{segment_id}/criteria/active-options` | Required |
| [DriveSegmentsApi](docs/api/drive_segments_api.html) | `patch` | `PATCH /drive-api/v1/segments/{segment_id}` | Required |
| [DriveSegmentsApi](docs/api/drive_segments_api.html) | `remove` | `DELETE /drive-api/v1/segments/{segment_id}` | Required |
| [DriveSegmentsApi](docs/api/drive_segments_api.html) | `update` | `PUT /drive-api/v2/segments/{segment_id}/criteria/{criteria_id}` | Required |
| [DriveSegmentsApi](docs/api/drive_segments_api.html) | `update` | `PUT /drive-api/v2/criteria/drafts/{criteria_id}` | Required |

## Connect to the API

- Bud sandbox: `https://api-sandbox.thisisbud.com`

The default credential is sent in the `Authorization` header with the `Bearer` prefix.

Authentication flow: 1. Perform OAuth2 Client Credentials authentication using API Credentials (`client_id`,`client_secret`) to obtain an `access_token` against `/v1/oauth/token` endpoint, 2. Use `access_token` as Bearer Authorisation for every other API request, 3. Include `X-Client-Id` (=client_id) within the header of every API request, 4. Note that some of the requests may also require `X-Customer-Id` to be provided within the request header. ### Examples Obtain OAuth2 `access_token` and `refresh_token` using `grant_type=client_credentials` and HTTP Basic auth header ``` curl --basic --user &#123;&#123;client_id&#125;&#125;:&#123;&#123;client_secret&#125;&#125; \ -X POST https://api-sandbox.thisisbud.com/v1/oauth/token \ -H &#39;Content-Type: application/x-www-form-urlencoded&#39; \ -d grant_type=client_credentials ``` Successful response: ``` &#123; &quot;operation_id&quot;: &quot;oauth_token_post&quot;, &quot;data&quot;: &#123; &quot;access_token&quot;: &quot;dd0c17e3fd6d2ce94aa091257a3ea393b4f9b5cf3d3e998f07dc9826da86ff15&quot;, &quot;token_type&quot;: &quot;bearer&quot;, &quot;expires_in&quot;: 3600, &quot;refresh_token&quot;: &quot;fac32cca7559d9f6e8f1dfe9a99c71fa1dcfeb482bedf287d7934d2667ae54b3&quot; &#125; &#125; ``` Refresh `access_token` token using `refresh_token` against `/v1/oauth/token` endpoint with `grant_type=refresh_token` ``` curl -X POST \ https://api-sandbox.thisisbud.com/v1/oauth/token \ -H &#39;Content-Type: application/x-www-form-urlencoded&#39; \ -H &#39;X-Client-Id: &#123;&#123;client_id&#125;&#125;&#39; \ -d &#39;grant_type=refresh_token&amp;refresh_token=&#123;&#123;refresh_token&#125;&#125;&#39; ``` Successful response: ``` &#123; &quot;operation_id&quot;: &quot;oauth_token_post&quot;, &quot;data&quot;: &#123; &quot;access_token&quot;: &quot;cc0c17e3fd6d2ce94aa091257a3ea393b4f9b5cf3d3e998f07dc9826da86ff94&quot;, &quot;token_type&quot;: &quot;bearer&quot;, &quot;expires_in&quot;: 3600, &quot;refresh_token&quot;: &quot;ffc30cca7559d9f6e8f1dfe9a99c71fa1dcfeb482bedf287d7934d2667ae54b3&quot; &#125; &#125; ```

Check authentication for the route you plan to call. A route that declares no authentication can be used without credentials; this does not change the requirements of other routes. Keep credentials in environment variables or a configured secret provider, and keep them out of source control and logs.

## Make a first request

1. Choose the API server and an operation that matches your task.
2. Check the operation’s required input and authentication. Use values valid for your account and environment.
3. Send one request and inspect the returned data before adding retries, concurrency, or a larger batch.

For an SDK call, install or build the chosen client, create a client instance with its documented configuration, and call the required entity operation. Language references describe the argument shape, asynchronous behaviour, and returned values.

## Choose an SDK

Choose the language already used by your application or service. The clients represent the same API model, while package setup, naming, and return types follow each language. Check the selected client’s reference and tests before integrating it into an existing application.

| Client | Repository directory | Distribution |
| --- | --- | --- |
| [Golang](docs/sdks/go.html) | `go/` | Build from source |
| [JavaScript](docs/sdks/js.html) | `js/` | Build from source |
| [Lua](docs/sdks/lua.html) | `lua/` | Build from source |
| [PHP](docs/sdks/php.html) | `php/` | Build from source |
| [Python](docs/sdks/py.html) | `py/` | Build from source |
| [TypeScript](docs/sdks/ts.html) | `ts/` | Build from source |

Build-from-source entries are not marked as published in the project model. Follow the build instructions in that target’s README, then consume the resulting package using your language’s local dependency mechanism. Published entries give the installation command recorded for that client.

## Companion tools

These targets provide another way to use the API. Their available commands or tools can cover a smaller set of operations than the client libraries.

### [Go CLI](docs/tools/go-cli.html)

Use the command-line interface for shell-based tasks and scripts.

Repository directory: `go-cli/`. Not published. Build from the go-cli directory.


### [Go MCP server](docs/tools/go-mcp.html)

Use the MCP server to expose supported API operations to an MCP client.

Repository directory: `go-mcp/`. Not published. Build from the go-mcp directory.

- `bud-drive_list`: List records for an entity. Supported entities: `drive_campaigns_api`, `drive_segments_api`.
- `bud-drive_load`: Load one record for an entity. Supported entities: `drive_campaigns_api`, `drive_segments_api`.

## Operational features

Features supply behaviour around API calls, such as request handling, diagnostics, or local testing. Inclusion in this project does not mean a feature is enabled at runtime. Check the selected SDK’s supported features and configuration defaults, then enable the behaviour your application needs.

- [`debug`](docs/features/debug.html): Request/response capture ring buffer for debugging
- [`idempotency`](docs/features/idempotency.html): Idempotency keys for safe retries of mutating operations
- [`metrics`](docs/features/metrics.html): Statistics capture: per-operation counters and latency
- [`paging`](docs/features/paging.html): Pagination signals for list operations
- [`ratelimit`](docs/features/ratelimit.html): Client-side rate limiting via a token bucket
- [`retry`](docs/features/retry.html): Automatic retry of transient failures with exponential backoff
- [`test`](docs/features/test.html): In-memory mock transport for testing without a live server
- [`timeout`](docs/features/timeout.html): Per-request timeout with transport abort

Start with the default client configuration. Add request limits and diagnostics as needed, test error paths, and review retry behaviour before using operations that change data. A retry can repeat an operation unless the API provides a suitable guarantee.

## Continue with the documentation

- Follow the [first-call guide](docs/guides/first-call.html) for the setup sequence.
- Read the [authentication guide](docs/guides/authentication.html) before using protected routes.
- Use the [API reference](docs/api/index.html) for request schemas, response formats, and status codes.
- Check the chosen SDK or companion tool reference for its configuration and supported operations.

