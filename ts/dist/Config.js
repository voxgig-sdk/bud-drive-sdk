"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.FEATURE_PLUGINS = exports.config = void 0;
const DebugFeature_1 = require("./feature/debug/DebugFeature");
const IdempotencyFeature_1 = require("./feature/idempotency/IdempotencyFeature");
const MetricsFeature_1 = require("./feature/metrics/MetricsFeature");
const PagingFeature_1 = require("./feature/paging/PagingFeature");
const RatelimitFeature_1 = require("./feature/ratelimit/RatelimitFeature");
const RetryFeature_1 = require("./feature/retry/RetryFeature");
const TestFeature_1 = require("./feature/test/TestFeature");
const TimeoutFeature_1 = require("./feature/timeout/TimeoutFeature");
const FEATURE_CLASS = {
    debug: DebugFeature_1.DebugFeature,
    idempotency: IdempotencyFeature_1.IdempotencyFeature,
    metrics: MetricsFeature_1.MetricsFeature,
    paging: PagingFeature_1.PagingFeature,
    ratelimit: RatelimitFeature_1.RatelimitFeature,
    retry: RetryFeature_1.RetryFeature,
    test: TestFeature_1.TestFeature,
    timeout: TimeoutFeature_1.TimeoutFeature,
};
const FEATURE_PLUGINS = {};
exports.FEATURE_PLUGINS = FEATURE_PLUGINS;
class Config {
    makeFeature(fn) {
        const fc = FEATURE_CLASS[fn];
        const fi = new fc();
        return fi;
    }
    // False for a feature added at runtime via options.extend (station's
    // adopt path) - the constructor uses this to skip makeFeature for names
    // no generated class backs.
    hasFeature(fn) {
        return null != FEATURE_CLASS[fn];
    }
    main = {
        name: 'BudDrive',
        slug: "bud-drive",
        version: "0.0.1",
        target: "ts",
    };
    feature = {
        debug: {
            "options": {
                "active": false,
                "max": 100,
                "redact": [
                    "authorization",
                    "cookie",
                    "set-cookie",
                    "api-key",
                    "apikey",
                    "x-api-key",
                    "idempotency-key"
                ]
            },
            "optspec": {
                "now": "`$FUNCTION`",
                "onEntry": "`$FUNCTION`"
            },
            "strict": false,
            "transport": "none"
        },
        idempotency: {
            "options": {
                "active": false,
                "header": "Idempotency-Key",
                "methods": [
                    "POST",
                    "PUT",
                    "PATCH",
                    "DELETE"
                ],
                "ops": [
                    "create",
                    "update",
                    "remove"
                ]
            },
            "optspec": {
                "keygen": "`$FUNCTION`"
            },
            "strict": false,
            "transport": "none"
        },
        metrics: {
            "options": {
                "active": false
            },
            "optspec": {
                "now": "`$FUNCTION`"
            },
            "strict": false,
            "transport": "none"
        },
        paging: {
            "options": {
                "active": false,
                "afterVar": "after",
                "cursorParam": "cursor",
                "firstVar": "first",
                "limitParam": "limit",
                "pageParam": "page",
                "startPage": 1
            },
            "optspec": {
                "limit": "`$NUMBER`",
                "ops": "`$LIST`"
            },
            "strict": false,
            "transport": "none"
        },
        ratelimit: {
            "options": {
                "active": false,
                "burst": 5,
                "rate": 5
            },
            "optspec": {
                "now": "`$FUNCTION`",
                "sleep": "`$FUNCTION`"
            },
            "strict": false,
            "transport": "wrap"
        },
        retry: {
            "options": {
                "active": false,
                "factor": 2,
                "maxDelay": 2000,
                "minDelay": 50,
                "retries": 2,
                "statuses": [
                    408,
                    425,
                    429,
                    500,
                    502,
                    503,
                    504
                ]
            },
            "optspec": {
                "jitter": "`$BOOLEAN`",
                "sleep": "`$FUNCTION`"
            },
            "strict": false,
            "transport": "wrap"
        },
        test: {
            "options": {
                "active": false
            },
            "optspec": {
                "entity": "`$MAP`",
                "net": "`$MAP`"
            },
            "strict": false,
            "transport": "base"
        },
        timeout: {
            "options": {
                "active": false,
                "ms": 30000
            },
            "optspec": {
                "clearTimer": "`$FUNCTION`",
                "setTimer": "`$FUNCTION`"
            },
            "strict": false,
            "transport": "wrap"
        },
    };
    options = {
        base: "https://api-sandbox.thisisbud.com",
        auth: {
            prefix: 'Bearer',
        },
        headers: {
            "content-type": "application/json"
        },
        entity: {
            drive_campaigns_api: {},
            drive_mcp_api: {},
            drive_segments_api: {},
        }
    };
    entity = {
        "drive_campaigns_api": {
            "fields": [
                {
                    "name": "audience",
                    "title": "Audience",
                    "type": "`$STRING`",
                    "op": {
                        "create": {
                            "req": true,
                            "type": "`$STRING`"
                        }
                    },
                    "short": "The ID of the segment this campaign targets.",
                    "format": "uuid"
                },
                {
                    "name": "branding",
                    "title": "Branding",
                    "type": "`$STRING`",
                    "short": "Overrides your configured value for this run only."
                },
                {
                    "name": "campaign_goal",
                    "title": "Campaign Goal",
                    "type": "`$STRING`",
                    "req": true,
                    "short": "The business goal the campaign should serve, in plain language."
                },
                {
                    "name": "campaign_id",
                    "title": "Campaign Id",
                    "type": "`$STRING`",
                    "req": true,
                    "op": {
                        "list": {
                            "type": "`$STRING`"
                        }
                    },
                    "format": "uuid"
                },
                {
                    "name": "channels",
                    "title": "Channels",
                    "type": "`$OBJECT`"
                },
                {
                    "name": "completed_at",
                    "title": "Completed At",
                    "type": "`$STRING`",
                    "format": "date-time"
                },
                {
                    "name": "context",
                    "title": "Context",
                    "type": "`$STRING`",
                    "short": "Overrides your configured value for this run only."
                },
                {
                    "name": "created_at",
                    "title": "Created At",
                    "type": "`$STRING`",
                    "format": "date-time"
                },
                {
                    "name": "created_by",
                    "title": "Created By",
                    "type": "`$STRING`"
                },
                {
                    "name": "description",
                    "title": "Description",
                    "type": "`$STRING`"
                },
                {
                    "name": "duration",
                    "title": "Duration",
                    "type": "`$OBJECT`"
                },
                {
                    "name": "examples",
                    "title": "Examples",
                    "type": "`$ARRAY`"
                },
                {
                    "name": "extra_context",
                    "title": "Extra Context",
                    "type": "`$STRING`",
                    "short": "Overrides your configured value for this run only."
                },
                {
                    "name": "group",
                    "title": "Group",
                    "type": "`$STRING`"
                },
                {
                    "name": "headline",
                    "title": "Headline",
                    "type": "`$OBJECT`"
                },
                {
                    "name": "id",
                    "title": "Id",
                    "type": "`$STRING`"
                },
                {
                    "name": "image_style",
                    "title": "Image Style",
                    "type": "`$STRING`",
                    "short": "Overrides your configured value for this run only."
                },
                {
                    "name": "image_url",
                    "title": "Image Url",
                    "type": "`$STRING`",
                    "format": "uri"
                },
                {
                    "name": "label",
                    "title": "Label",
                    "type": "`$STRING`"
                },
                {
                    "name": "message",
                    "title": "Message",
                    "type": "`$STRING`"
                },
                {
                    "name": "name",
                    "title": "Name",
                    "type": "`$STRING`",
                    "op": {
                        "create": {
                            "req": true,
                            "type": "`$STRING`"
                        }
                    }
                },
                {
                    "name": "parameters",
                    "title": "Parameters",
                    "type": "`$ARRAY`"
                },
                {
                    "name": "primary_button",
                    "title": "Primary Button",
                    "type": "`$OBJECT`"
                },
                {
                    "name": "priority_order",
                    "title": "Priority Order",
                    "type": "`$INTEGER`"
                },
                {
                    "name": "result",
                    "title": "Result",
                    "type": "`$OBJECT`",
                    "short": "The campaign a completed run produced."
                },
                {
                    "name": "schedule",
                    "title": "Schedule",
                    "type": "`$OBJECT`",
                    "short": "When the campaign is published."
                },
                {
                    "name": "started_at",
                    "title": "Started At",
                    "type": "`$STRING`",
                    "format": "date-time"
                },
                {
                    "name": "status",
                    "title": "Status",
                    "type": "`$STRING`",
                    "short": "The campaign's current status."
                },
                {
                    "name": "styling",
                    "title": "Styling",
                    "type": "`$OBJECT`"
                },
                {
                    "name": "template_id",
                    "title": "Template Id",
                    "type": "`$STRING`",
                    "format": "uuid"
                },
                {
                    "name": "template_variables",
                    "title": "Template Variables",
                    "type": "`$OBJECT`"
                },
                {
                    "name": "templates",
                    "title": "Templates",
                    "type": "`$ARRAY`"
                },
                {
                    "name": "tone_of_voice",
                    "title": "Tone Of Voice",
                    "type": "`$STRING`",
                    "short": "Overrides your configured value for this run only."
                },
                {
                    "name": "trend",
                    "title": "Trend",
                    "type": "`$OBJECT`"
                },
                {
                    "name": "type",
                    "title": "Type",
                    "type": "`$STRING`"
                }
            ],
            "id": {
                "field": "id",
                "name": "id"
            },
            "name": "drive_campaigns_api",
            "op": {
                "create": {
                    "input": "data",
                    "name": "create",
                    "points": [
                        {
                            "kind": "http",
                            "method": "POST",
                            "orig": "/drive-api/v1/campaigns",
                            "segments": [
                                {
                                    "lit": "drive-api"
                                },
                                {
                                    "lit": "v1"
                                },
                                {
                                    "lit": "campaigns"
                                }
                            ],
                            "parts": [
                                "drive-api",
                                "v1",
                                "campaigns"
                            ],
                            "rename": {},
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body.data`"
                            },
                            "args": {
                                "header": [
                                    {
                                        "name": "x_client_id",
                                        "orig": "x_client_id",
                                        "type": "`$STRING`",
                                        "kind": "header",
                                        "reqd": true,
                                        "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a"
                                    }
                                ]
                            },
                            "select": {
                                "exist": [
                                    "x_client_id"
                                ]
                            }
                        },
                        {
                            "kind": "http",
                            "method": "POST",
                            "orig": "/drive-api/v1/campaigns/agent",
                            "segments": [
                                {
                                    "lit": "drive-api"
                                },
                                {
                                    "lit": "v1"
                                },
                                {
                                    "lit": "campaigns"
                                },
                                {
                                    "lit": "agent"
                                }
                            ],
                            "parts": [
                                "drive-api",
                                "v1",
                                "campaigns",
                                "agent"
                            ],
                            "rename": {},
                            "transform": {
                                "req": {
                                    "branding": "`reqdata.branding`",
                                    "campaign_goal": "`reqdata.campaign_goal`",
                                    "context": "`reqdata.context`",
                                    "extra_context": "`reqdata.extra_context`",
                                    "image_style": "`reqdata.image_style`",
                                    "tone_of_voice": "`reqdata.tone_of_voice`"
                                },
                                "res": "`body`"
                            },
                            "args": {
                                "header": [
                                    {
                                        "name": "x_client_id",
                                        "orig": "x_client_id",
                                        "type": "`$STRING`",
                                        "kind": "header",
                                        "reqd": true,
                                        "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a"
                                    }
                                ]
                            },
                            "select": {
                                "exist": [
                                    "x_client_id"
                                ]
                            }
                        }
                    ]
                },
                "list": {
                    "input": "data",
                    "name": "list",
                    "points": [
                        {
                            "kind": "http",
                            "method": "GET",
                            "orig": "/drive-api/v1/campaigns",
                            "segments": [
                                {
                                    "lit": "drive-api"
                                },
                                {
                                    "lit": "v1"
                                },
                                {
                                    "lit": "campaigns"
                                }
                            ],
                            "parts": [
                                "drive-api",
                                "v1",
                                "campaigns"
                            ],
                            "rename": {},
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body`"
                            },
                            "args": {
                                "header": [
                                    {
                                        "name": "x_client_id",
                                        "orig": "x_client_id",
                                        "type": "`$STRING`",
                                        "kind": "header",
                                        "reqd": true,
                                        "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a"
                                    }
                                ],
                                "query": [
                                    {
                                        "name": "created_by",
                                        "orig": "created_by",
                                        "type": "`$STRING`",
                                        "kind": "query"
                                    },
                                    {
                                        "name": "items_per_page",
                                        "orig": "items_per_page",
                                        "type": "`$INTEGER`",
                                        "kind": "query"
                                    },
                                    {
                                        "name": "page_token",
                                        "orig": "page_token",
                                        "type": "`$STRING`",
                                        "kind": "query"
                                    },
                                    {
                                        "name": "status",
                                        "orig": "status",
                                        "type": "`$STRING`",
                                        "kind": "query"
                                    }
                                ]
                            },
                            "select": {
                                "exist": [
                                    "created_by",
                                    "items_per_page",
                                    "page_token",
                                    "status",
                                    "x_client_id"
                                ]
                            }
                        },
                        {
                            "kind": "http",
                            "method": "GET",
                            "orig": "/drive-api/v1/campaigns/banner-widgets",
                            "segments": [
                                {
                                    "lit": "drive-api"
                                },
                                {
                                    "lit": "v1"
                                },
                                {
                                    "lit": "campaigns"
                                },
                                {
                                    "lit": "banner-widgets"
                                }
                            ],
                            "parts": [
                                "drive-api",
                                "v1",
                                "campaigns",
                                "banner-widgets"
                            ],
                            "rename": {},
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body.data`"
                            },
                            "args": {
                                "header": [
                                    {
                                        "name": "x_client_id",
                                        "orig": "x_client_id",
                                        "type": "`$STRING`",
                                        "kind": "header",
                                        "reqd": true,
                                        "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a"
                                    },
                                    {
                                        "name": "x_customer_id",
                                        "orig": "x_customer_id",
                                        "type": "`$STRING`",
                                        "kind": "header",
                                        "reqd": true,
                                        "example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23"
                                    }
                                ],
                                "query": [
                                    {
                                        "name": "channel",
                                        "orig": "channel",
                                        "type": "`$STRING`",
                                        "kind": "query",
                                        "example": "action_card"
                                    }
                                ]
                            },
                            "select": {
                                "exist": [
                                    "channel",
                                    "x_client_id",
                                    "x_customer_id"
                                ]
                            }
                        },
                        {
                            "kind": "http",
                            "method": "GET",
                            "orig": "/drive-api/v1/campaigns/content-variables",
                            "segments": [
                                {
                                    "lit": "drive-api"
                                },
                                {
                                    "lit": "v1"
                                },
                                {
                                    "lit": "campaigns"
                                },
                                {
                                    "lit": "content-variables"
                                }
                            ],
                            "parts": [
                                "drive-api",
                                "v1",
                                "campaigns",
                                "content-variables"
                            ],
                            "rename": {},
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body.data`"
                            },
                            "args": {
                                "header": [
                                    {
                                        "name": "x_client_id",
                                        "orig": "x_client_id",
                                        "type": "`$STRING`",
                                        "kind": "header",
                                        "reqd": true,
                                        "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a"
                                    }
                                ]
                            },
                            "select": {
                                "exist": [
                                    "x_client_id"
                                ]
                            }
                        }
                    ]
                },
                "load": {
                    "input": "data",
                    "name": "load",
                    "points": [
                        {
                            "kind": "http",
                            "method": "GET",
                            "orig": "/drive-api/v1/campaigns/{campaign_id}/insights/performance/trends",
                            "segments": [
                                {
                                    "lit": "drive-api"
                                },
                                {
                                    "lit": "v1"
                                },
                                {
                                    "lit": "campaigns"
                                },
                                {
                                    "var": "campaign_id"
                                },
                                {
                                    "lit": "insights"
                                },
                                {
                                    "lit": "performance"
                                },
                                {
                                    "lit": "trends"
                                }
                            ],
                            "parts": [
                                "drive-api",
                                "v1",
                                "campaigns",
                                "{campaign_id}",
                                "insights",
                                "performance",
                                "trends"
                            ],
                            "rename": {},
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body.data`"
                            },
                            "args": {
                                "header": [
                                    {
                                        "name": "x_client_id",
                                        "orig": "x_client_id",
                                        "type": "`$STRING`",
                                        "kind": "header",
                                        "reqd": true,
                                        "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a"
                                    }
                                ],
                                "params": [
                                    {
                                        "name": "campaign_id",
                                        "orig": "campaign_id",
                                        "type": "`$STRING`",
                                        "kind": "param",
                                        "reqd": true,
                                        "example": "b28efd2c-5a58-42cd-bed7-5564eb877ca7"
                                    }
                                ],
                                "query": [
                                    {
                                        "name": "aggregation_type",
                                        "orig": "aggregation_type",
                                        "type": "`$STRING`",
                                        "kind": "query",
                                        "example": "daily"
                                    },
                                    {
                                        "name": "end_date",
                                        "orig": "end_date",
                                        "type": "`$STRING`",
                                        "kind": "query"
                                    },
                                    {
                                        "name": "metric",
                                        "orig": "metric",
                                        "type": "`$ARRAY`",
                                        "kind": "query",
                                        "reqd": true
                                    },
                                    {
                                        "name": "start_date",
                                        "orig": "start_date",
                                        "type": "`$STRING`",
                                        "kind": "query"
                                    }
                                ]
                            },
                            "select": {
                                "exist": [
                                    "aggregation_type",
                                    "campaign_id",
                                    "end_date",
                                    "metric",
                                    "start_date",
                                    "x_client_id"
                                ]
                            }
                        },
                        {
                            "kind": "http",
                            "method": "GET",
                            "orig": "/drive-api/v1/campaigns/{campaign_id}",
                            "segments": [
                                {
                                    "lit": "drive-api"
                                },
                                {
                                    "lit": "v1"
                                },
                                {
                                    "lit": "campaigns"
                                },
                                {
                                    "var": "campaign_id"
                                }
                            ],
                            "parts": [
                                "drive-api",
                                "v1",
                                "campaigns",
                                "{campaign_id}"
                            ],
                            "rename": {},
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body.data`"
                            },
                            "args": {
                                "header": [
                                    {
                                        "name": "x_client_id",
                                        "orig": "x_client_id",
                                        "type": "`$STRING`",
                                        "kind": "header",
                                        "reqd": true,
                                        "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a"
                                    }
                                ],
                                "params": [
                                    {
                                        "name": "campaign_id",
                                        "orig": "campaign_id",
                                        "type": "`$STRING`",
                                        "kind": "param",
                                        "reqd": true,
                                        "example": "b28efd2c-5a58-42cd-bed7-5564eb877ca7"
                                    }
                                ]
                            },
                            "select": {
                                "exist": [
                                    "campaign_id",
                                    "x_client_id"
                                ]
                            }
                        },
                        {
                            "kind": "http",
                            "method": "GET",
                            "orig": "/drive-api/v1/campaigns/{campaign_id}/content",
                            "segments": [
                                {
                                    "lit": "drive-api"
                                },
                                {
                                    "lit": "v1"
                                },
                                {
                                    "lit": "campaigns"
                                },
                                {
                                    "var": "campaign_id"
                                },
                                {
                                    "lit": "content"
                                }
                            ],
                            "parts": [
                                "drive-api",
                                "v1",
                                "campaigns",
                                "{campaign_id}",
                                "content"
                            ],
                            "rename": {},
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body.data`"
                            },
                            "args": {
                                "header": [
                                    {
                                        "name": "x_client_id",
                                        "orig": "x_client_id",
                                        "type": "`$STRING`",
                                        "kind": "header",
                                        "reqd": true,
                                        "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a"
                                    }
                                ],
                                "params": [
                                    {
                                        "name": "campaign_id",
                                        "orig": "campaign_id",
                                        "type": "`$STRING`",
                                        "kind": "param",
                                        "reqd": true,
                                        "example": "b28efd2c-5a58-42cd-bed7-5564eb877ca7"
                                    }
                                ]
                            },
                            "select": {
                                "$action": "content",
                                "exist": [
                                    "campaign_id",
                                    "x_client_id"
                                ]
                            }
                        },
                        {
                            "kind": "http",
                            "method": "GET",
                            "orig": "/drive-api/v1/campaigns/agent/{generation_id}",
                            "segments": [
                                {
                                    "lit": "drive-api"
                                },
                                {
                                    "lit": "v1"
                                },
                                {
                                    "lit": "campaigns"
                                },
                                {
                                    "lit": "agent"
                                },
                                {
                                    "var": "generation_id"
                                }
                            ],
                            "parts": [
                                "drive-api",
                                "v1",
                                "campaigns",
                                "agent",
                                "{generation_id}"
                            ],
                            "rename": {},
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body.data`"
                            },
                            "args": {
                                "header": [
                                    {
                                        "name": "x_client_id",
                                        "orig": "x_client_id",
                                        "type": "`$STRING`",
                                        "kind": "header",
                                        "reqd": true,
                                        "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a"
                                    }
                                ],
                                "params": [
                                    {
                                        "name": "generation_id",
                                        "orig": "generation_id",
                                        "type": "`$STRING`",
                                        "kind": "param",
                                        "reqd": true,
                                        "example": "ceca9875e4674c6699fc404b76f1786f"
                                    }
                                ]
                            },
                            "select": {
                                "exist": [
                                    "generation_id",
                                    "x_client_id"
                                ]
                            }
                        },
                        {
                            "kind": "http",
                            "method": "GET",
                            "orig": "/drive-api/v1/campaign-templates/compiled-templates",
                            "segments": [
                                {
                                    "lit": "drive-api"
                                },
                                {
                                    "lit": "v1"
                                },
                                {
                                    "lit": "campaign-templates"
                                },
                                {
                                    "lit": "compiled-templates"
                                }
                            ],
                            "parts": [
                                "drive-api",
                                "v1",
                                "campaign-templates",
                                "compiled-templates"
                            ],
                            "rename": {},
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body.data`"
                            },
                            "args": {
                                "header": [
                                    {
                                        "name": "x_client_id",
                                        "orig": "x_client_id",
                                        "type": "`$STRING`",
                                        "kind": "header",
                                        "reqd": true,
                                        "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a"
                                    }
                                ],
                                "query": [
                                    {
                                        "name": "source_type",
                                        "orig": "source_type",
                                        "type": "`$STRING`",
                                        "kind": "query"
                                    }
                                ]
                            },
                            "select": {
                                "exist": [
                                    "source_type",
                                    "x_client_id"
                                ]
                            }
                        }
                    ]
                },
                "remove": {
                    "input": "data",
                    "name": "remove",
                    "points": [
                        {
                            "kind": "http",
                            "method": "DELETE",
                            "orig": "/drive-api/v1/campaigns/{campaign_id}",
                            "segments": [
                                {
                                    "lit": "drive-api"
                                },
                                {
                                    "lit": "v1"
                                },
                                {
                                    "lit": "campaigns"
                                },
                                {
                                    "var": "campaign_id"
                                }
                            ],
                            "parts": [
                                "drive-api",
                                "v1",
                                "campaigns",
                                "{campaign_id}"
                            ],
                            "rename": {},
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body`"
                            },
                            "args": {
                                "header": [
                                    {
                                        "name": "x_client_id",
                                        "orig": "x_client_id",
                                        "type": "`$STRING`",
                                        "kind": "header",
                                        "reqd": true,
                                        "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a"
                                    }
                                ],
                                "params": [
                                    {
                                        "name": "campaign_id",
                                        "orig": "campaign_id",
                                        "type": "`$STRING`",
                                        "kind": "param",
                                        "reqd": true,
                                        "example": "b28efd2c-5a58-42cd-bed7-5564eb877ca7"
                                    }
                                ]
                            },
                            "select": {
                                "exist": [
                                    "campaign_id",
                                    "x_client_id"
                                ]
                            }
                        },
                        {
                            "kind": "http",
                            "method": "DELETE",
                            "orig": "/drive-api/v1/campaign-templates/templates/{template_id}",
                            "segments": [
                                {
                                    "lit": "drive-api"
                                },
                                {
                                    "lit": "v1"
                                },
                                {
                                    "lit": "campaign-templates"
                                },
                                {
                                    "lit": "templates"
                                },
                                {
                                    "var": "template_id"
                                }
                            ],
                            "parts": [
                                "drive-api",
                                "v1",
                                "campaign-templates",
                                "templates",
                                "{template_id}"
                            ],
                            "rename": {},
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body`"
                            },
                            "args": {
                                "header": [
                                    {
                                        "name": "x_client_id",
                                        "orig": "x_client_id",
                                        "type": "`$STRING`",
                                        "kind": "header",
                                        "reqd": true,
                                        "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a"
                                    }
                                ],
                                "params": [
                                    {
                                        "name": "template_id",
                                        "orig": "template_id",
                                        "type": "`$STRING`",
                                        "kind": "param",
                                        "reqd": true,
                                        "example": "89d133a9-02f8-4217-adb8-1d341fc91101"
                                    }
                                ]
                            },
                            "select": {
                                "exist": [
                                    "template_id",
                                    "x_client_id"
                                ]
                            }
                        }
                    ]
                },
                "update": {
                    "input": "data",
                    "name": "update",
                    "points": [
                        {
                            "kind": "http",
                            "method": "PUT",
                            "orig": "/drive-api/v1/campaigns/{campaign_id}/content",
                            "segments": [
                                {
                                    "lit": "drive-api"
                                },
                                {
                                    "lit": "v1"
                                },
                                {
                                    "lit": "campaigns"
                                },
                                {
                                    "var": "campaign_id"
                                },
                                {
                                    "lit": "content"
                                }
                            ],
                            "parts": [
                                "drive-api",
                                "v1",
                                "campaigns",
                                "{campaign_id}",
                                "content"
                            ],
                            "rename": {},
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body`"
                            },
                            "args": {
                                "header": [
                                    {
                                        "name": "x_client_id",
                                        "orig": "x_client_id",
                                        "type": "`$STRING`",
                                        "kind": "header",
                                        "reqd": true,
                                        "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a"
                                    }
                                ],
                                "params": [
                                    {
                                        "name": "campaign_id",
                                        "orig": "campaign_id",
                                        "type": "`$STRING`",
                                        "kind": "param",
                                        "reqd": true,
                                        "example": "b28efd2c-5a58-42cd-bed7-5564eb877ca7"
                                    }
                                ]
                            },
                            "select": {
                                "$action": "content",
                                "exist": [
                                    "campaign_id",
                                    "x_client_id"
                                ]
                            }
                        },
                        {
                            "kind": "http",
                            "method": "PUT",
                            "orig": "/drive-api/v1/campaigns/{campaign_id}/status",
                            "segments": [
                                {
                                    "lit": "drive-api"
                                },
                                {
                                    "lit": "v1"
                                },
                                {
                                    "lit": "campaigns"
                                },
                                {
                                    "var": "campaign_id"
                                },
                                {
                                    "lit": "status"
                                }
                            ],
                            "parts": [
                                "drive-api",
                                "v1",
                                "campaigns",
                                "{campaign_id}",
                                "status"
                            ],
                            "rename": {},
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body`"
                            },
                            "args": {
                                "header": [
                                    {
                                        "name": "x_client_id",
                                        "orig": "x_client_id",
                                        "type": "`$STRING`",
                                        "kind": "header",
                                        "reqd": true,
                                        "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a"
                                    }
                                ],
                                "params": [
                                    {
                                        "name": "campaign_id",
                                        "orig": "campaign_id",
                                        "type": "`$STRING`",
                                        "kind": "param",
                                        "reqd": true,
                                        "example": "b28efd2c-5a58-42cd-bed7-5564eb877ca7"
                                    }
                                ]
                            },
                            "select": {
                                "$action": "status",
                                "exist": [
                                    "campaign_id",
                                    "x_client_id"
                                ]
                            }
                        }
                    ]
                }
            },
            "relations": {
                "ancestors": []
            }
        },
        "drive_mcp_api": {
            "fields": [
                {
                    "name": "error",
                    "title": "Error",
                    "type": "`$OBJECT`"
                },
                {
                    "name": "id",
                    "title": "Id",
                    "type": "`$STRING`",
                    "req": true,
                    "op": {
                        "create": {
                            "type": "`$STRING`"
                        }
                    }
                },
                {
                    "name": "jsonrpc",
                    "title": "Jsonrpc",
                    "type": "`$STRING`",
                    "req": true
                },
                {
                    "name": "method",
                    "title": "Method",
                    "type": "`$STRING`",
                    "req": true
                },
                {
                    "name": "params",
                    "title": "Params",
                    "type": "`$OBJECT`"
                },
                {
                    "name": "result",
                    "title": "Result",
                    "type": "`$OBJECT`"
                }
            ],
            "id": {
                "field": "id",
                "name": "id"
            },
            "name": "drive_mcp_api",
            "op": {
                "create": {
                    "input": "data",
                    "name": "create",
                    "points": [
                        {
                            "kind": "http",
                            "method": "POST",
                            "orig": "/drive-mcp/v1/campaigns",
                            "segments": [
                                {
                                    "lit": "drive-mcp"
                                },
                                {
                                    "lit": "v1"
                                },
                                {
                                    "lit": "campaigns"
                                }
                            ],
                            "parts": [
                                "drive-mcp",
                                "v1",
                                "campaigns"
                            ],
                            "rename": {},
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body`"
                            },
                            "args": {
                                "header": [
                                    {
                                        "name": "x_client_id",
                                        "orig": "x_client_id",
                                        "type": "`$STRING`",
                                        "kind": "header",
                                        "reqd": true,
                                        "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a"
                                    }
                                ]
                            },
                            "select": {
                                "exist": [
                                    "x_client_id"
                                ]
                            }
                        },
                        {
                            "kind": "http",
                            "method": "POST",
                            "orig": "/drive-mcp/v1/copilot",
                            "segments": [
                                {
                                    "lit": "drive-mcp"
                                },
                                {
                                    "lit": "v1"
                                },
                                {
                                    "lit": "copilot"
                                }
                            ],
                            "parts": [
                                "drive-mcp",
                                "v1",
                                "copilot"
                            ],
                            "rename": {},
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body`"
                            },
                            "args": {
                                "header": [
                                    {
                                        "name": "x_client_id",
                                        "orig": "x_client_id",
                                        "type": "`$STRING`",
                                        "kind": "header",
                                        "reqd": true,
                                        "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a"
                                    }
                                ]
                            },
                            "select": {
                                "exist": [
                                    "x_client_id"
                                ]
                            }
                        },
                        {
                            "kind": "http",
                            "method": "POST",
                            "orig": "/drive-mcp/v1/segments",
                            "segments": [
                                {
                                    "lit": "drive-mcp"
                                },
                                {
                                    "lit": "v1"
                                },
                                {
                                    "lit": "segments"
                                }
                            ],
                            "parts": [
                                "drive-mcp",
                                "v1",
                                "segments"
                            ],
                            "rename": {},
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body`"
                            },
                            "args": {
                                "header": [
                                    {
                                        "name": "x_client_id",
                                        "orig": "x_client_id",
                                        "type": "`$STRING`",
                                        "kind": "header",
                                        "reqd": true,
                                        "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a"
                                    }
                                ]
                            },
                            "select": {
                                "exist": [
                                    "x_client_id"
                                ]
                            }
                        }
                    ]
                }
            },
            "relations": {
                "ancestors": []
            }
        },
        "drive_segments_api": {
            "fields": [
                {
                    "name": "created_at",
                    "title": "Created At",
                    "type": "`$STRING`",
                    "format": "date-time"
                },
                {
                    "name": "created_by",
                    "title": "Created By",
                    "type": "`$STRING`"
                },
                {
                    "name": "criteria",
                    "title": "Criteria",
                    "type": "`$ANY`",
                    "req": true
                },
                {
                    "name": "criteria_from",
                    "title": "Criteria From",
                    "type": "`$STRING`",
                    "req": true,
                    "short": "The ID of the segment whose criteria should be copied.",
                    "format": "uuid"
                },
                {
                    "name": "customers",
                    "title": "Customers",
                    "type": "`$ARRAY`",
                    "req": true,
                    "short": "Customer IDs to include in the fixed customer list."
                },
                {
                    "name": "customers_from",
                    "title": "Customers From",
                    "type": "`$STRING`",
                    "req": true,
                    "short": "The ID of the segment whose current members should be copied.",
                    "format": "uuid"
                },
                {
                    "name": "description",
                    "title": "Description",
                    "type": "`$STRING`"
                },
                {
                    "name": "fixed_customer_list",
                    "title": "Fixed Customer List",
                    "type": "`$BOOLEAN`"
                },
                {
                    "name": "fixed_customer_list_size",
                    "title": "Fixed Customer List Size",
                    "type": "`$INTEGER`"
                },
                {
                    "name": "from_criteria_id",
                    "title": "From Criteria Id",
                    "type": "`$STRING`",
                    "format": "uuid"
                },
                {
                    "name": "id",
                    "title": "Id",
                    "type": "`$STRING`",
                    "req": true,
                    "format": "uuid"
                },
                {
                    "name": "intersection_count",
                    "title": "Intersection Count",
                    "type": "`$INTEGER`"
                },
                {
                    "name": "jaccard_index",
                    "title": "Jaccard Index",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "name",
                    "title": "Name",
                    "type": "`$STRING`",
                    "req": true,
                    "op": {
                        "list": {
                            "type": "`$STRING`"
                        },
                        "patch": {
                            "type": "`$STRING`"
                        }
                    }
                },
                {
                    "name": "parameters",
                    "title": "Parameters",
                    "type": "`$OBJECT`"
                },
                {
                    "name": "source",
                    "title": "Source",
                    "type": "`$STRING`",
                    "short": "How the segment was created."
                },
                {
                    "name": "statistics_id",
                    "title": "Statistics Id",
                    "type": "`$STRING`"
                },
                {
                    "name": "suggestion",
                    "title": "Suggestion",
                    "type": "`$OBJECT`"
                },
                {
                    "name": "tag",
                    "title": "Tag",
                    "type": "`$STRING`"
                },
                {
                    "name": "tags",
                    "title": "Tags",
                    "type": "`$ARRAY`"
                },
                {
                    "name": "type",
                    "title": "Type",
                    "type": "`$STRING`",
                    "short": "The shape of the statistic's values."
                },
                {
                    "name": "union_count",
                    "title": "Union Count",
                    "type": "`$INTEGER`"
                },
                {
                    "name": "upload",
                    "title": "Upload",
                    "type": "`$OBJECT`"
                },
                {
                    "name": "values",
                    "title": "Values",
                    "type": "`$ARRAY`"
                }
            ],
            "id": {
                "field": "id",
                "name": "id"
            },
            "name": "drive_segments_api",
            "op": {
                "create": {
                    "input": "data",
                    "name": "create",
                    "points": [
                        {
                            "kind": "http",
                            "method": "POST",
                            "orig": "/drive-api/v1/segments/{segment_id}/query",
                            "segments": [
                                {
                                    "lit": "drive-api"
                                },
                                {
                                    "lit": "v1"
                                },
                                {
                                    "lit": "segments"
                                },
                                {
                                    "var": "segment_id"
                                },
                                {
                                    "lit": "query"
                                }
                            ],
                            "parts": [
                                "drive-api",
                                "v1",
                                "segments",
                                "{segment_id}",
                                "query"
                            ],
                            "rename": {},
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body.data`"
                            },
                            "args": {
                                "header": [
                                    {
                                        "name": "x_client_id",
                                        "orig": "x_client_id",
                                        "type": "`$STRING`",
                                        "kind": "header",
                                        "reqd": true,
                                        "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a"
                                    }
                                ],
                                "params": [
                                    {
                                        "name": "segment_id",
                                        "orig": "segment_id",
                                        "type": "`$STRING`",
                                        "kind": "param",
                                        "reqd": true,
                                        "example": "23303f83-a790-4ee9-95b6-d5cdeaf36fa2"
                                    }
                                ]
                            },
                            "select": {
                                "$action": "query",
                                "exist": [
                                    "segment_id",
                                    "x_client_id"
                                ]
                            }
                        },
                        {
                            "kind": "http",
                            "method": "POST",
                            "orig": "/drive-api/v2/criteria/drafts/{criteria_id}/run",
                            "segments": [
                                {
                                    "lit": "drive-api"
                                },
                                {
                                    "lit": "v2"
                                },
                                {
                                    "lit": "criteria"
                                },
                                {
                                    "lit": "drafts"
                                },
                                {
                                    "var": "draft_id"
                                },
                                {
                                    "lit": "run"
                                }
                            ],
                            "parts": [
                                "drive-api",
                                "v2",
                                "criteria",
                                "drafts",
                                "{draft_id}",
                                "run"
                            ],
                            "rename": {
                                "param": {
                                    "criteria_id": "draft_id"
                                }
                            },
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body`"
                            },
                            "args": {
                                "header": [
                                    {
                                        "name": "x_client_id",
                                        "orig": "x_client_id",
                                        "type": "`$STRING`",
                                        "kind": "header",
                                        "reqd": true,
                                        "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a"
                                    }
                                ],
                                "params": [
                                    {
                                        "name": "draft_id",
                                        "orig": "criteria_id",
                                        "type": "`$STRING`",
                                        "kind": "param",
                                        "reqd": true,
                                        "example": "0ba9049e-25e9-43f2-8ea8-1fcf9dde581e"
                                    }
                                ]
                            },
                            "select": {
                                "$action": "run",
                                "exist": [
                                    "draft_id",
                                    "x_client_id"
                                ]
                            }
                        },
                        {
                            "kind": "http",
                            "method": "POST",
                            "orig": "/drive-api/v2/segments/{segment_id}/criteria/run",
                            "segments": [
                                {
                                    "lit": "drive-api"
                                },
                                {
                                    "lit": "v2"
                                },
                                {
                                    "lit": "segments"
                                },
                                {
                                    "var": "segment_id"
                                },
                                {
                                    "lit": "criteria"
                                },
                                {
                                    "lit": "run"
                                }
                            ],
                            "parts": [
                                "drive-api",
                                "v2",
                                "segments",
                                "{segment_id}",
                                "criteria",
                                "run"
                            ],
                            "rename": {},
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body`"
                            },
                            "args": {
                                "header": [
                                    {
                                        "name": "x_client_id",
                                        "orig": "x_client_id",
                                        "type": "`$STRING`",
                                        "kind": "header",
                                        "reqd": true,
                                        "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a"
                                    }
                                ],
                                "params": [
                                    {
                                        "name": "segment_id",
                                        "orig": "segment_id",
                                        "type": "`$STRING`",
                                        "kind": "param",
                                        "reqd": true,
                                        "example": "23303f83-a790-4ee9-95b6-d5cdeaf36fa2"
                                    }
                                ]
                            },
                            "select": {
                                "exist": [
                                    "segment_id",
                                    "x_client_id"
                                ]
                            }
                        },
                        {
                            "kind": "http",
                            "method": "POST",
                            "orig": "/drive-api/v1/segments/from-existing/criteria",
                            "segments": [
                                {
                                    "lit": "drive-api"
                                },
                                {
                                    "lit": "v1"
                                },
                                {
                                    "lit": "segments"
                                },
                                {
                                    "lit": "from-existing"
                                },
                                {
                                    "lit": "criteria"
                                }
                            ],
                            "parts": [
                                "drive-api",
                                "v1",
                                "segments",
                                "from-existing",
                                "criteria"
                            ],
                            "rename": {},
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body.data`"
                            },
                            "args": {
                                "header": [
                                    {
                                        "name": "x_client_id",
                                        "orig": "x_client_id",
                                        "type": "`$STRING`",
                                        "kind": "header",
                                        "reqd": true,
                                        "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a"
                                    }
                                ]
                            },
                            "select": {
                                "exist": [
                                    "x_client_id"
                                ]
                            }
                        },
                        {
                            "kind": "http",
                            "method": "POST",
                            "orig": "/drive-api/v1/segments/from-existing/fixed-customers",
                            "segments": [
                                {
                                    "lit": "drive-api"
                                },
                                {
                                    "lit": "v1"
                                },
                                {
                                    "lit": "segments"
                                },
                                {
                                    "lit": "from-existing"
                                },
                                {
                                    "lit": "fixed-customers"
                                }
                            ],
                            "parts": [
                                "drive-api",
                                "v1",
                                "segments",
                                "from-existing",
                                "fixed-customers"
                            ],
                            "rename": {},
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body.data`"
                            },
                            "args": {
                                "header": [
                                    {
                                        "name": "x_client_id",
                                        "orig": "x_client_id",
                                        "type": "`$STRING`",
                                        "kind": "header",
                                        "reqd": true,
                                        "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a"
                                    }
                                ]
                            },
                            "select": {
                                "exist": [
                                    "x_client_id"
                                ]
                            }
                        },
                        {
                            "kind": "http",
                            "method": "POST",
                            "orig": "/drive-api/v1/segments/upload",
                            "segments": [
                                {
                                    "lit": "drive-api"
                                },
                                {
                                    "lit": "v1"
                                },
                                {
                                    "lit": "segments"
                                },
                                {
                                    "lit": "upload"
                                }
                            ],
                            "parts": [
                                "drive-api",
                                "v1",
                                "segments",
                                "upload"
                            ],
                            "rename": {},
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body.data`"
                            },
                            "args": {
                                "header": [
                                    {
                                        "name": "x_client_id",
                                        "orig": "x_client_id",
                                        "type": "`$STRING`",
                                        "kind": "header",
                                        "reqd": true,
                                        "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a"
                                    }
                                ]
                            },
                            "select": {
                                "exist": [
                                    "x_client_id"
                                ]
                            }
                        },
                        {
                            "kind": "http",
                            "method": "POST",
                            "orig": "/drive-api/v2/criteria/drafts",
                            "segments": [
                                {
                                    "lit": "drive-api"
                                },
                                {
                                    "lit": "v2"
                                },
                                {
                                    "lit": "criteria"
                                },
                                {
                                    "lit": "drafts"
                                }
                            ],
                            "parts": [
                                "drive-api",
                                "v2",
                                "criteria",
                                "drafts"
                            ],
                            "rename": {},
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body`"
                            },
                            "args": {
                                "header": [
                                    {
                                        "name": "x_client_id",
                                        "orig": "x_client_id",
                                        "type": "`$STRING`",
                                        "kind": "header",
                                        "reqd": true,
                                        "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a"
                                    }
                                ]
                            },
                            "select": {
                                "exist": [
                                    "x_client_id"
                                ]
                            }
                        },
                        {
                            "kind": "http",
                            "method": "POST",
                            "orig": "/drive-api/v2/criteria/run",
                            "segments": [
                                {
                                    "lit": "drive-api"
                                },
                                {
                                    "lit": "v2"
                                },
                                {
                                    "lit": "criteria"
                                },
                                {
                                    "lit": "run"
                                }
                            ],
                            "parts": [
                                "drive-api",
                                "v2",
                                "criteria",
                                "run"
                            ],
                            "rename": {},
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body`"
                            },
                            "args": {
                                "header": [
                                    {
                                        "name": "x_client_id",
                                        "orig": "x_client_id",
                                        "type": "`$STRING`",
                                        "kind": "header",
                                        "reqd": true,
                                        "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a"
                                    }
                                ]
                            },
                            "select": {
                                "exist": [
                                    "x_client_id"
                                ]
                            }
                        },
                        {
                            "kind": "http",
                            "method": "POST",
                            "orig": "/drive-api/v2/criteria/suggestions/segment-name",
                            "segments": [
                                {
                                    "lit": "drive-api"
                                },
                                {
                                    "lit": "v2"
                                },
                                {
                                    "lit": "criteria"
                                },
                                {
                                    "lit": "suggestions"
                                },
                                {
                                    "lit": "segment-name"
                                }
                            ],
                            "parts": [
                                "drive-api",
                                "v2",
                                "criteria",
                                "suggestions",
                                "segment-name"
                            ],
                            "rename": {},
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body.data`"
                            },
                            "args": {
                                "header": [
                                    {
                                        "name": "x_client_id",
                                        "orig": "x_client_id",
                                        "type": "`$STRING`",
                                        "kind": "header",
                                        "reqd": true,
                                        "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a"
                                    }
                                ]
                            },
                            "select": {
                                "exist": [
                                    "x_client_id"
                                ]
                            }
                        },
                        {
                            "kind": "http",
                            "method": "POST",
                            "orig": "/drive-api/v2/segments",
                            "segments": [
                                {
                                    "lit": "drive-api"
                                },
                                {
                                    "lit": "v2"
                                },
                                {
                                    "lit": "segments"
                                }
                            ],
                            "parts": [
                                "drive-api",
                                "v2",
                                "segments"
                            ],
                            "rename": {},
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body`"
                            },
                            "args": {
                                "header": [
                                    {
                                        "name": "x_client_id",
                                        "orig": "x_client_id",
                                        "type": "`$STRING`",
                                        "kind": "header",
                                        "reqd": true,
                                        "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a"
                                    }
                                ]
                            },
                            "select": {
                                "exist": [
                                    "x_client_id"
                                ]
                            }
                        }
                    ]
                },
                "list": {
                    "input": "data",
                    "name": "list",
                    "points": [
                        {
                            "kind": "http",
                            "method": "GET",
                            "orig": "/drive-api/v1/segments/{segment_id}/statistics/overview",
                            "segments": [
                                {
                                    "lit": "drive-api"
                                },
                                {
                                    "lit": "v1"
                                },
                                {
                                    "lit": "segments"
                                },
                                {
                                    "var": "segment_id"
                                },
                                {
                                    "lit": "statistics"
                                },
                                {
                                    "lit": "overview"
                                }
                            ],
                            "parts": [
                                "drive-api",
                                "v1",
                                "segments",
                                "{segment_id}",
                                "statistics",
                                "overview"
                            ],
                            "rename": {},
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body.data`"
                            },
                            "args": {
                                "header": [
                                    {
                                        "name": "x_client_id",
                                        "orig": "x_client_id",
                                        "type": "`$STRING`",
                                        "kind": "header",
                                        "reqd": true,
                                        "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a"
                                    }
                                ],
                                "params": [
                                    {
                                        "name": "segment_id",
                                        "orig": "segment_id",
                                        "type": "`$STRING`",
                                        "kind": "param",
                                        "reqd": true,
                                        "example": "23303f83-a790-4ee9-95b6-d5cdeaf36fa2"
                                    }
                                ],
                                "query": [
                                    {
                                        "name": "date_to",
                                        "orig": "date_to",
                                        "type": "`$STRING`",
                                        "kind": "query"
                                    },
                                    {
                                        "name": "significant_date",
                                        "orig": "significant_date",
                                        "type": "`$STRING`",
                                        "kind": "query"
                                    }
                                ]
                            },
                            "select": {
                                "exist": [
                                    "date_to",
                                    "segment_id",
                                    "significant_date",
                                    "x_client_id"
                                ]
                            }
                        },
                        {
                            "kind": "http",
                            "method": "GET",
                            "orig": "/drive-api/v2/segments",
                            "segments": [
                                {
                                    "lit": "drive-api"
                                },
                                {
                                    "lit": "v2"
                                },
                                {
                                    "lit": "segments"
                                }
                            ],
                            "parts": [
                                "drive-api",
                                "v2",
                                "segments"
                            ],
                            "rename": {},
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body`"
                            },
                            "args": {
                                "header": [
                                    {
                                        "name": "x_client_id",
                                        "orig": "x_client_id",
                                        "type": "`$STRING`",
                                        "kind": "header",
                                        "reqd": true,
                                        "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a"
                                    }
                                ],
                                "query": [
                                    {
                                        "name": "max_per_page",
                                        "orig": "max_per_page",
                                        "type": "`$INTEGER`",
                                        "kind": "query"
                                    },
                                    {
                                        "name": "page",
                                        "orig": "page",
                                        "type": "`$INTEGER`",
                                        "kind": "query"
                                    },
                                    {
                                        "name": "search",
                                        "orig": "search",
                                        "type": "`$STRING`",
                                        "kind": "query"
                                    }
                                ]
                            },
                            "select": {
                                "exist": [
                                    "max_per_page",
                                    "page",
                                    "search",
                                    "x_client_id"
                                ]
                            }
                        },
                        {
                            "kind": "http",
                            "method": "GET",
                            "orig": "/drive-api/v1/tags",
                            "segments": [
                                {
                                    "lit": "drive-api"
                                },
                                {
                                    "lit": "v1"
                                },
                                {
                                    "lit": "tags"
                                }
                            ],
                            "parts": [
                                "drive-api",
                                "v1",
                                "tags"
                            ],
                            "rename": {},
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body`"
                            },
                            "args": {
                                "header": [
                                    {
                                        "name": "x_client_id",
                                        "orig": "x_client_id",
                                        "type": "`$STRING`",
                                        "kind": "header",
                                        "reqd": true,
                                        "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a"
                                    }
                                ]
                            },
                            "select": {
                                "exist": [
                                    "x_client_id"
                                ]
                            }
                        }
                    ]
                },
                "load": {
                    "input": "data",
                    "name": "load",
                    "points": [
                        {
                            "kind": "http",
                            "method": "GET",
                            "orig": "/drive-api/v1/segments/{segment_id}/similarity/{other_segment_id}",
                            "segments": [
                                {
                                    "lit": "drive-api"
                                },
                                {
                                    "lit": "v1"
                                },
                                {
                                    "lit": "segments"
                                },
                                {
                                    "var": "segment_id"
                                },
                                {
                                    "lit": "similarity"
                                },
                                {
                                    "var": "other_segment_id"
                                }
                            ],
                            "parts": [
                                "drive-api",
                                "v1",
                                "segments",
                                "{segment_id}",
                                "similarity",
                                "{other_segment_id}"
                            ],
                            "rename": {},
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body.data`"
                            },
                            "args": {
                                "header": [
                                    {
                                        "name": "x_client_id",
                                        "orig": "x_client_id",
                                        "type": "`$STRING`",
                                        "kind": "header",
                                        "reqd": true,
                                        "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a"
                                    }
                                ],
                                "params": [
                                    {
                                        "name": "other_segment_id",
                                        "orig": "other_segment_id",
                                        "type": "`$STRING`",
                                        "kind": "param",
                                        "reqd": true,
                                        "example": "7e8aedf9-6be5-4807-afe0-aa292fb2cde2"
                                    },
                                    {
                                        "name": "segment_id",
                                        "orig": "segment_id",
                                        "type": "`$STRING`",
                                        "kind": "param",
                                        "reqd": true,
                                        "example": "23303f83-a790-4ee9-95b6-d5cdeaf36fa2"
                                    }
                                ],
                                "query": [
                                    {
                                        "name": "refresh",
                                        "orig": "refresh",
                                        "type": "`$BOOLEAN`",
                                        "kind": "query"
                                    }
                                ]
                            },
                            "select": {
                                "exist": [
                                    "other_segment_id",
                                    "refresh",
                                    "segment_id",
                                    "x_client_id"
                                ]
                            }
                        },
                        {
                            "kind": "http",
                            "method": "GET",
                            "orig": "/drive-api/v2/criteria/drafts/{criteria_id}",
                            "segments": [
                                {
                                    "lit": "drive-api"
                                },
                                {
                                    "lit": "v2"
                                },
                                {
                                    "lit": "criteria"
                                },
                                {
                                    "lit": "drafts"
                                },
                                {
                                    "var": "criteria_id"
                                }
                            ],
                            "parts": [
                                "drive-api",
                                "v2",
                                "criteria",
                                "drafts",
                                "{criteria_id}"
                            ],
                            "rename": {},
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body`"
                            },
                            "args": {
                                "header": [
                                    {
                                        "name": "x_client_id",
                                        "orig": "x_client_id",
                                        "type": "`$STRING`",
                                        "kind": "header",
                                        "reqd": true,
                                        "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a"
                                    }
                                ],
                                "params": [
                                    {
                                        "name": "criteria_id",
                                        "orig": "criteria_id",
                                        "type": "`$STRING`",
                                        "kind": "param",
                                        "reqd": true,
                                        "example": "0ba9049e-25e9-43f2-8ea8-1fcf9dde581e"
                                    }
                                ]
                            },
                            "select": {
                                "exist": [
                                    "criteria_id",
                                    "x_client_id"
                                ]
                            }
                        },
                        {
                            "kind": "http",
                            "method": "GET",
                            "orig": "/drive-api/v2/criteria/options",
                            "segments": [
                                {
                                    "lit": "drive-api"
                                },
                                {
                                    "lit": "v2"
                                },
                                {
                                    "lit": "criteria"
                                },
                                {
                                    "lit": "options"
                                }
                            ],
                            "parts": [
                                "drive-api",
                                "v2",
                                "criteria",
                                "options"
                            ],
                            "rename": {},
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body`"
                            },
                            "args": {
                                "header": [
                                    {
                                        "name": "x_client_id",
                                        "orig": "x_client_id",
                                        "type": "`$STRING`",
                                        "kind": "header",
                                        "reqd": true,
                                        "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a"
                                    }
                                ],
                                "query": [
                                    {
                                        "name": "filter",
                                        "orig": "filter",
                                        "type": "`$STRING`",
                                        "kind": "query"
                                    }
                                ]
                            },
                            "select": {
                                "exist": [
                                    "filter",
                                    "x_client_id"
                                ]
                            }
                        },
                        {
                            "kind": "http",
                            "method": "GET",
                            "orig": "/drive-api/v1/segments/{segment_id}",
                            "segments": [
                                {
                                    "lit": "drive-api"
                                },
                                {
                                    "lit": "v1"
                                },
                                {
                                    "lit": "segments"
                                },
                                {
                                    "var": "segment_id"
                                }
                            ],
                            "parts": [
                                "drive-api",
                                "v1",
                                "segments",
                                "{segment_id}"
                            ],
                            "rename": {},
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body.data`"
                            },
                            "args": {
                                "header": [
                                    {
                                        "name": "x_client_id",
                                        "orig": "x_client_id",
                                        "type": "`$STRING`",
                                        "kind": "header",
                                        "reqd": true,
                                        "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a"
                                    }
                                ],
                                "params": [
                                    {
                                        "name": "segment_id",
                                        "orig": "segment_id",
                                        "type": "`$STRING`",
                                        "kind": "param",
                                        "reqd": true,
                                        "example": "23303f83-a790-4ee9-95b6-d5cdeaf36fa2"
                                    }
                                ]
                            },
                            "select": {
                                "exist": [
                                    "segment_id",
                                    "x_client_id"
                                ]
                            }
                        },
                        {
                            "kind": "http",
                            "method": "GET",
                            "orig": "/drive-api/v2/segments/{segment_id}/criteria",
                            "segments": [
                                {
                                    "lit": "drive-api"
                                },
                                {
                                    "lit": "v2"
                                },
                                {
                                    "lit": "segments"
                                },
                                {
                                    "var": "segment_id"
                                },
                                {
                                    "lit": "criteria"
                                }
                            ],
                            "parts": [
                                "drive-api",
                                "v2",
                                "segments",
                                "{segment_id}",
                                "criteria"
                            ],
                            "rename": {},
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body`"
                            },
                            "args": {
                                "header": [
                                    {
                                        "name": "x_client_id",
                                        "orig": "x_client_id",
                                        "type": "`$STRING`",
                                        "kind": "header",
                                        "reqd": true,
                                        "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a"
                                    }
                                ],
                                "params": [
                                    {
                                        "name": "segment_id",
                                        "orig": "segment_id",
                                        "type": "`$STRING`",
                                        "kind": "param",
                                        "reqd": true,
                                        "example": "23303f83-a790-4ee9-95b6-d5cdeaf36fa2"
                                    }
                                ]
                            },
                            "select": {
                                "exist": [
                                    "segment_id",
                                    "x_client_id"
                                ]
                            }
                        },
                        {
                            "kind": "http",
                            "method": "GET",
                            "orig": "/drive-api/v2/segments/{segment_id}/criteria/active-options",
                            "segments": [
                                {
                                    "lit": "drive-api"
                                },
                                {
                                    "lit": "v2"
                                },
                                {
                                    "lit": "segments"
                                },
                                {
                                    "var": "segment_id"
                                },
                                {
                                    "lit": "criteria"
                                },
                                {
                                    "lit": "active-options"
                                }
                            ],
                            "parts": [
                                "drive-api",
                                "v2",
                                "segments",
                                "{segment_id}",
                                "criteria",
                                "active-options"
                            ],
                            "rename": {},
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body`"
                            },
                            "args": {
                                "header": [
                                    {
                                        "name": "x_client_id",
                                        "orig": "x_client_id",
                                        "type": "`$STRING`",
                                        "kind": "header",
                                        "reqd": true,
                                        "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a"
                                    }
                                ],
                                "params": [
                                    {
                                        "name": "segment_id",
                                        "orig": "segment_id",
                                        "type": "`$STRING`",
                                        "kind": "param",
                                        "reqd": true,
                                        "example": "23303f83-a790-4ee9-95b6-d5cdeaf36fa2"
                                    }
                                ]
                            },
                            "select": {
                                "exist": [
                                    "segment_id",
                                    "x_client_id"
                                ]
                            }
                        }
                    ]
                },
                "patch": {
                    "input": "data",
                    "name": "patch",
                    "points": [
                        {
                            "kind": "http",
                            "method": "PATCH",
                            "orig": "/drive-api/v1/segments/{segment_id}",
                            "segments": [
                                {
                                    "lit": "drive-api"
                                },
                                {
                                    "lit": "v1"
                                },
                                {
                                    "lit": "segments"
                                },
                                {
                                    "var": "segment_id"
                                }
                            ],
                            "parts": [
                                "drive-api",
                                "v1",
                                "segments",
                                "{segment_id}"
                            ],
                            "rename": {},
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body.data`"
                            },
                            "args": {
                                "header": [
                                    {
                                        "name": "x_client_id",
                                        "orig": "x_client_id",
                                        "type": "`$STRING`",
                                        "kind": "header",
                                        "reqd": true,
                                        "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a"
                                    }
                                ],
                                "params": [
                                    {
                                        "name": "segment_id",
                                        "orig": "segment_id",
                                        "type": "`$STRING`",
                                        "kind": "param",
                                        "reqd": true,
                                        "example": "23303f83-a790-4ee9-95b6-d5cdeaf36fa2"
                                    }
                                ]
                            },
                            "select": {
                                "exist": [
                                    "segment_id",
                                    "x_client_id"
                                ]
                            }
                        }
                    ]
                },
                "remove": {
                    "input": "data",
                    "name": "remove",
                    "points": [
                        {
                            "kind": "http",
                            "method": "DELETE",
                            "orig": "/drive-api/v1/segments/{segment_id}",
                            "segments": [
                                {
                                    "lit": "drive-api"
                                },
                                {
                                    "lit": "v1"
                                },
                                {
                                    "lit": "segments"
                                },
                                {
                                    "var": "segment_id"
                                }
                            ],
                            "parts": [
                                "drive-api",
                                "v1",
                                "segments",
                                "{segment_id}"
                            ],
                            "rename": {},
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body.metadata`"
                            },
                            "args": {
                                "header": [
                                    {
                                        "name": "x_client_id",
                                        "orig": "x_client_id",
                                        "type": "`$STRING`",
                                        "kind": "header",
                                        "reqd": true,
                                        "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a"
                                    }
                                ],
                                "params": [
                                    {
                                        "name": "segment_id",
                                        "orig": "segment_id",
                                        "type": "`$STRING`",
                                        "kind": "param",
                                        "reqd": true,
                                        "example": "23303f83-a790-4ee9-95b6-d5cdeaf36fa2"
                                    }
                                ]
                            },
                            "select": {
                                "exist": [
                                    "segment_id",
                                    "x_client_id"
                                ]
                            }
                        }
                    ]
                },
                "update": {
                    "input": "data",
                    "name": "update",
                    "points": [
                        {
                            "kind": "http",
                            "method": "PUT",
                            "orig": "/drive-api/v2/segments/{segment_id}/criteria/{criteria_id}",
                            "segments": [
                                {
                                    "lit": "drive-api"
                                },
                                {
                                    "lit": "v2"
                                },
                                {
                                    "lit": "segments"
                                },
                                {
                                    "var": "segment_id"
                                },
                                {
                                    "lit": "criteria"
                                },
                                {
                                    "var": "criteria_id"
                                }
                            ],
                            "parts": [
                                "drive-api",
                                "v2",
                                "segments",
                                "{segment_id}",
                                "criteria",
                                "{criteria_id}"
                            ],
                            "rename": {},
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body`"
                            },
                            "args": {
                                "header": [
                                    {
                                        "name": "x_client_id",
                                        "orig": "x_client_id",
                                        "type": "`$STRING`",
                                        "kind": "header",
                                        "reqd": true,
                                        "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a"
                                    }
                                ],
                                "params": [
                                    {
                                        "name": "criteria_id",
                                        "orig": "criteria_id",
                                        "type": "`$STRING`",
                                        "kind": "param",
                                        "reqd": true,
                                        "example": "0ba9049e-25e9-43f2-8ea8-1fcf9dde581e"
                                    },
                                    {
                                        "name": "segment_id",
                                        "orig": "segment_id",
                                        "type": "`$STRING`",
                                        "kind": "param",
                                        "reqd": true,
                                        "example": "23303f83-a790-4ee9-95b6-d5cdeaf36fa2"
                                    }
                                ]
                            },
                            "select": {
                                "exist": [
                                    "criteria_id",
                                    "segment_id",
                                    "x_client_id"
                                ]
                            }
                        },
                        {
                            "kind": "http",
                            "method": "PUT",
                            "orig": "/drive-api/v2/criteria/drafts/{criteria_id}",
                            "segments": [
                                {
                                    "lit": "drive-api"
                                },
                                {
                                    "lit": "v2"
                                },
                                {
                                    "lit": "criteria"
                                },
                                {
                                    "lit": "drafts"
                                },
                                {
                                    "var": "criteria_id"
                                }
                            ],
                            "parts": [
                                "drive-api",
                                "v2",
                                "criteria",
                                "drafts",
                                "{criteria_id}"
                            ],
                            "rename": {},
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body`"
                            },
                            "args": {
                                "header": [
                                    {
                                        "name": "x_client_id",
                                        "orig": "x_client_id",
                                        "type": "`$STRING`",
                                        "kind": "header",
                                        "reqd": true,
                                        "example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a"
                                    }
                                ],
                                "params": [
                                    {
                                        "name": "criteria_id",
                                        "orig": "criteria_id",
                                        "type": "`$STRING`",
                                        "kind": "param",
                                        "reqd": true,
                                        "example": "0ba9049e-25e9-43f2-8ea8-1fcf9dde581e"
                                    }
                                ]
                            },
                            "select": {
                                "exist": [
                                    "criteria_id",
                                    "x_client_id"
                                ]
                            }
                        }
                    ]
                }
            },
            "relations": {
                "ancestors": []
            }
        }
    };
}
const config = new Config();
exports.config = config;
//# sourceMappingURL=Config.js.map