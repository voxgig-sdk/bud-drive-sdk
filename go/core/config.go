package core

import (
	"sync"
)

// MakeConfig builds a fresh, fully materialised config map. Every call
// rebuilds the whole structure, so prefer SharedConfig unless you need a
// private copy you intend to mutate.
func MakeConfig() map[string]any {
	return map[string]any{
		"main": map[string]any{
			"name": "BudDrive",
			"slug": "bud-drive",
			"version": "0.0.1",
			"target": "go",
		},
		"feature": map[string]any{
			"debug": map[string]any{
				"options": map[string]any{
					"active": false,
					"max": 100,
					"redact": []any{
						"authorization",
						"cookie",
						"set-cookie",
						"api-key",
						"apikey",
						"x-api-key",
						"idempotency-key",
					},
				},
				"optspec": map[string]any{
					"now": "`$FUNCTION`",
					"onEntry": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "none",
			},
			"idempotency": map[string]any{
				"options": map[string]any{
					"active": false,
					"header": "Idempotency-Key",
					"methods": []any{
						"POST",
						"PUT",
						"PATCH",
						"DELETE",
					},
					"ops": []any{
						"create",
						"update",
						"remove",
					},
				},
				"optspec": map[string]any{
					"keygen": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "none",
			},
			"metrics": map[string]any{
				"options": map[string]any{
					"active": false,
				},
				"optspec": map[string]any{
					"now": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "none",
			},
			"paging": map[string]any{
				"options": map[string]any{
					"active": false,
					"afterVar": "after",
					"cursorParam": "cursor",
					"firstVar": "first",
					"limitParam": "limit",
					"pageParam": "page",
					"startPage": 1,
				},
				"optspec": map[string]any{
					"limit": "`$NUMBER`",
					"ops": "`$LIST`",
				},
				"strict": false,
				"transport": "none",
			},
			"ratelimit": map[string]any{
				"options": map[string]any{
					"active": false,
					"burst": 5,
					"rate": 5,
				},
				"optspec": map[string]any{
					"now": "`$FUNCTION`",
					"sleep": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
			"retry": map[string]any{
				"options": map[string]any{
					"active": false,
					"factor": 2,
					"maxDelay": 2000,
					"minDelay": 50,
					"retries": 2,
					"statuses": []any{
						408,
						425,
						429,
						500,
						502,
						503,
						504,
					},
				},
				"optspec": map[string]any{
					"jitter": "`$BOOLEAN`",
					"sleep": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
				"optspec": map[string]any{
					"entity": "`$MAP`",
					"net": "`$MAP`",
				},
				"strict": false,
				"transport": "base",
			},
			"timeout": map[string]any{
				"options": map[string]any{
					"active": false,
					"ms": 30000,
				},
				"optspec": map[string]any{
					"clearTimer": "`$FUNCTION`",
					"setTimer": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
		},
		"options": map[string]any{
			"base": "https://api-sandbox.thisisbud.com",
			"auth": map[string]any{
				"prefix": "Bearer",
			},
			"headers": map[string]any{
				"content-type": "application/json",
			},
			"entity": map[string]any{
				"drive_campaigns_api": map[string]any{},
				"drive_mcp_api": map[string]any{},
				"drive_segments_api": map[string]any{},
			},
		},
		"entity": map[string]any{
			"drive_campaigns_api": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "audience",
						"title": "Audience",
						"type": "`$STRING`",
						"op": map[string]any{
							"create": map[string]any{
								"req": true,
								"type": "`$STRING`",
							},
						},
						"short": "The ID of the segment this campaign targets.",
						"format": "uuid",
					},
					map[string]any{
						"name": "branding",
						"title": "Branding",
						"type": "`$STRING`",
						"short": "Overrides your configured value for this run only.",
					},
					map[string]any{
						"name": "campaign_goal",
						"title": "Campaign Goal",
						"type": "`$STRING`",
						"req": true,
						"short": "The business goal the campaign should serve, in plain language.",
					},
					map[string]any{
						"name": "campaign_id",
						"title": "Campaign Id",
						"type": "`$STRING`",
						"req": true,
						"op": map[string]any{
							"list": map[string]any{
								"type": "`$STRING`",
							},
						},
						"format": "uuid",
					},
					map[string]any{
						"name": "channels",
						"title": "Channels",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "completed_at",
						"title": "Completed At",
						"type": "`$STRING`",
						"format": "date-time",
					},
					map[string]any{
						"name": "context",
						"title": "Context",
						"type": "`$STRING`",
						"short": "Overrides your configured value for this run only.",
					},
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"format": "date-time",
					},
					map[string]any{
						"name": "created_by",
						"title": "Created By",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "description",
						"title": "Description",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "duration",
						"title": "Duration",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "examples",
						"title": "Examples",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "extra_context",
						"title": "Extra Context",
						"type": "`$STRING`",
						"short": "Overrides your configured value for this run only.",
					},
					map[string]any{
						"name": "group",
						"title": "Group",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "headline",
						"title": "Headline",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "image_style",
						"title": "Image Style",
						"type": "`$STRING`",
						"short": "Overrides your configured value for this run only.",
					},
					map[string]any{
						"name": "image_url",
						"title": "Image Url",
						"type": "`$STRING`",
						"format": "uri",
					},
					map[string]any{
						"name": "label",
						"title": "Label",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "message",
						"title": "Message",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
						"op": map[string]any{
							"create": map[string]any{
								"req": true,
								"type": "`$STRING`",
							},
						},
					},
					map[string]any{
						"name": "parameters",
						"title": "Parameters",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "primary_button",
						"title": "Primary Button",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "priority_order",
						"title": "Priority Order",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "result",
						"title": "Result",
						"type": "`$OBJECT`",
						"short": "The campaign a completed run produced.",
					},
					map[string]any{
						"name": "schedule",
						"title": "Schedule",
						"type": "`$OBJECT`",
						"short": "When the campaign is published.",
					},
					map[string]any{
						"name": "started_at",
						"title": "Started At",
						"type": "`$STRING`",
						"format": "date-time",
					},
					map[string]any{
						"name": "status",
						"title": "Status",
						"type": "`$STRING`",
						"short": "The campaign's current status.",
					},
					map[string]any{
						"name": "styling",
						"title": "Styling",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "template_id",
						"title": "Template Id",
						"type": "`$STRING`",
						"format": "uuid",
					},
					map[string]any{
						"name": "template_variables",
						"title": "Template Variables",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "templates",
						"title": "Templates",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "tone_of_voice",
						"title": "Tone Of Voice",
						"type": "`$STRING`",
						"short": "Overrides your configured value for this run only.",
					},
					map[string]any{
						"name": "trend",
						"title": "Trend",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "type",
						"title": "Type",
						"type": "`$STRING`",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "drive_campaigns_api",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/drive-api/v1/campaigns",
								"segments": []any{
									map[string]any{
										"lit": "drive-api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "campaigns",
									},
								},
								"parts": []any{
									"drive-api",
									"v1",
									"campaigns",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"x_client_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/drive-api/v1/campaigns/agent",
								"segments": []any{
									map[string]any{
										"lit": "drive-api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "campaigns",
									},
									map[string]any{
										"lit": "agent",
									},
								},
								"parts": []any{
									"drive-api",
									"v1",
									"campaigns",
									"agent",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"branding": "`reqdata.branding`",
										"campaign_goal": "`reqdata.campaign_goal`",
										"context": "`reqdata.context`",
										"extra_context": "`reqdata.extra_context`",
										"image_style": "`reqdata.image_style`",
										"tone_of_voice": "`reqdata.tone_of_voice`",
									},
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"x_client_id",
									},
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/drive-api/v1/campaigns",
								"segments": []any{
									map[string]any{
										"lit": "drive-api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "campaigns",
									},
								},
								"parts": []any{
									"drive-api",
									"v1",
									"campaigns",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
									"query": []any{
										map[string]any{
											"name": "created_by",
											"orig": "created_by",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "items_per_page",
											"orig": "items_per_page",
											"type": "`$INTEGER`",
											"kind": "query",
										},
										map[string]any{
											"name": "page_token",
											"orig": "page_token",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "status",
											"orig": "status",
											"type": "`$STRING`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"created_by",
										"items_per_page",
										"page_token",
										"status",
										"x_client_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/drive-api/v1/campaigns/banner-widgets",
								"segments": []any{
									map[string]any{
										"lit": "drive-api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "campaigns",
									},
									map[string]any{
										"lit": "banner-widgets",
									},
								},
								"parts": []any{
									"drive-api",
									"v1",
									"campaigns",
									"banner-widgets",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
										map[string]any{
											"name": "x_customer_id",
											"orig": "x_customer_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "c3339af9-2426-44d4-9c4d-ddb8fd281e23",
										},
									},
									"query": []any{
										map[string]any{
											"name": "channel",
											"orig": "channel",
											"type": "`$STRING`",
											"kind": "query",
											"example": "action_card",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"channel",
										"x_client_id",
										"x_customer_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/drive-api/v1/campaigns/content-variables",
								"segments": []any{
									map[string]any{
										"lit": "drive-api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "campaigns",
									},
									map[string]any{
										"lit": "content-variables",
									},
								},
								"parts": []any{
									"drive-api",
									"v1",
									"campaigns",
									"content-variables",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"x_client_id",
									},
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/drive-api/v1/campaigns/{campaign_id}/insights/performance/trends",
								"segments": []any{
									map[string]any{
										"lit": "drive-api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "campaigns",
									},
									map[string]any{
										"var": "campaign_id",
									},
									map[string]any{
										"lit": "insights",
									},
									map[string]any{
										"lit": "performance",
									},
									map[string]any{
										"lit": "trends",
									},
								},
								"parts": []any{
									"drive-api",
									"v1",
									"campaigns",
									"{campaign_id}",
									"insights",
									"performance",
									"trends",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
									"params": []any{
										map[string]any{
											"name": "campaign_id",
											"orig": "campaign_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "b28efd2c-5a58-42cd-bed7-5564eb877ca7",
										},
									},
									"query": []any{
										map[string]any{
											"name": "aggregation_type",
											"orig": "aggregation_type",
											"type": "`$STRING`",
											"kind": "query",
											"example": "daily",
										},
										map[string]any{
											"name": "end_date",
											"orig": "end_date",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "metric",
											"orig": "metric",
											"type": "`$ARRAY`",
											"kind": "query",
											"reqd": true,
										},
										map[string]any{
											"name": "start_date",
											"orig": "start_date",
											"type": "`$STRING`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"aggregation_type",
										"campaign_id",
										"end_date",
										"metric",
										"start_date",
										"x_client_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/drive-api/v1/campaigns/{campaign_id}",
								"segments": []any{
									map[string]any{
										"lit": "drive-api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "campaigns",
									},
									map[string]any{
										"var": "campaign_id",
									},
								},
								"parts": []any{
									"drive-api",
									"v1",
									"campaigns",
									"{campaign_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
									"params": []any{
										map[string]any{
											"name": "campaign_id",
											"orig": "campaign_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "b28efd2c-5a58-42cd-bed7-5564eb877ca7",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"campaign_id",
										"x_client_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/drive-api/v1/campaigns/{campaign_id}/content",
								"segments": []any{
									map[string]any{
										"lit": "drive-api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "campaigns",
									},
									map[string]any{
										"var": "campaign_id",
									},
									map[string]any{
										"lit": "content",
									},
								},
								"parts": []any{
									"drive-api",
									"v1",
									"campaigns",
									"{campaign_id}",
									"content",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
									"params": []any{
										map[string]any{
											"name": "campaign_id",
											"orig": "campaign_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "b28efd2c-5a58-42cd-bed7-5564eb877ca7",
										},
									},
								},
								"select": map[string]any{
									"$action": "content",
									"exist": []any{
										"campaign_id",
										"x_client_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/drive-api/v1/campaigns/agent/{generation_id}",
								"segments": []any{
									map[string]any{
										"lit": "drive-api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "campaigns",
									},
									map[string]any{
										"lit": "agent",
									},
									map[string]any{
										"var": "generation_id",
									},
								},
								"parts": []any{
									"drive-api",
									"v1",
									"campaigns",
									"agent",
									"{generation_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
									"params": []any{
										map[string]any{
											"name": "generation_id",
											"orig": "generation_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "ceca9875e4674c6699fc404b76f1786f",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"generation_id",
										"x_client_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/drive-api/v1/campaign-templates/compiled-templates",
								"segments": []any{
									map[string]any{
										"lit": "drive-api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "campaign-templates",
									},
									map[string]any{
										"lit": "compiled-templates",
									},
								},
								"parts": []any{
									"drive-api",
									"v1",
									"campaign-templates",
									"compiled-templates",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
									"query": []any{
										map[string]any{
											"name": "source_type",
											"orig": "source_type",
											"type": "`$STRING`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"source_type",
										"x_client_id",
									},
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/drive-api/v1/campaigns/{campaign_id}",
								"segments": []any{
									map[string]any{
										"lit": "drive-api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "campaigns",
									},
									map[string]any{
										"var": "campaign_id",
									},
								},
								"parts": []any{
									"drive-api",
									"v1",
									"campaigns",
									"{campaign_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
									"params": []any{
										map[string]any{
											"name": "campaign_id",
											"orig": "campaign_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "b28efd2c-5a58-42cd-bed7-5564eb877ca7",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"campaign_id",
										"x_client_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/drive-api/v1/campaign-templates/templates/{template_id}",
								"segments": []any{
									map[string]any{
										"lit": "drive-api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "campaign-templates",
									},
									map[string]any{
										"lit": "templates",
									},
									map[string]any{
										"var": "template_id",
									},
								},
								"parts": []any{
									"drive-api",
									"v1",
									"campaign-templates",
									"templates",
									"{template_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
									"params": []any{
										map[string]any{
											"name": "template_id",
											"orig": "template_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "89d133a9-02f8-4217-adb8-1d341fc91101",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"template_id",
										"x_client_id",
									},
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PUT",
								"orig": "/drive-api/v1/campaigns/{campaign_id}/content",
								"segments": []any{
									map[string]any{
										"lit": "drive-api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "campaigns",
									},
									map[string]any{
										"var": "campaign_id",
									},
									map[string]any{
										"lit": "content",
									},
								},
								"parts": []any{
									"drive-api",
									"v1",
									"campaigns",
									"{campaign_id}",
									"content",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
									"params": []any{
										map[string]any{
											"name": "campaign_id",
											"orig": "campaign_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "b28efd2c-5a58-42cd-bed7-5564eb877ca7",
										},
									},
								},
								"select": map[string]any{
									"$action": "content",
									"exist": []any{
										"campaign_id",
										"x_client_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "PUT",
								"orig": "/drive-api/v1/campaigns/{campaign_id}/status",
								"segments": []any{
									map[string]any{
										"lit": "drive-api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "campaigns",
									},
									map[string]any{
										"var": "campaign_id",
									},
									map[string]any{
										"lit": "status",
									},
								},
								"parts": []any{
									"drive-api",
									"v1",
									"campaigns",
									"{campaign_id}",
									"status",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
									"params": []any{
										map[string]any{
											"name": "campaign_id",
											"orig": "campaign_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "b28efd2c-5a58-42cd-bed7-5564eb877ca7",
										},
									},
								},
								"select": map[string]any{
									"$action": "status",
									"exist": []any{
										"campaign_id",
										"x_client_id",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"drive_mcp_api": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "error",
						"title": "Error",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"req": true,
						"op": map[string]any{
							"create": map[string]any{
								"type": "`$STRING`",
							},
						},
					},
					map[string]any{
						"name": "jsonrpc",
						"title": "Jsonrpc",
						"type": "`$STRING`",
						"req": true,
					},
					map[string]any{
						"name": "method",
						"title": "Method",
						"type": "`$STRING`",
						"req": true,
					},
					map[string]any{
						"name": "params",
						"title": "Params",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "result",
						"title": "Result",
						"type": "`$OBJECT`",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "drive_mcp_api",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/drive-mcp/v1/campaigns",
								"segments": []any{
									map[string]any{
										"lit": "drive-mcp",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "campaigns",
									},
								},
								"parts": []any{
									"drive-mcp",
									"v1",
									"campaigns",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"x_client_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/drive-mcp/v1/copilot",
								"segments": []any{
									map[string]any{
										"lit": "drive-mcp",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "copilot",
									},
								},
								"parts": []any{
									"drive-mcp",
									"v1",
									"copilot",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"x_client_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/drive-mcp/v1/segments",
								"segments": []any{
									map[string]any{
										"lit": "drive-mcp",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "segments",
									},
								},
								"parts": []any{
									"drive-mcp",
									"v1",
									"segments",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"x_client_id",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"drive_segments_api": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"format": "date-time",
					},
					map[string]any{
						"name": "created_by",
						"title": "Created By",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "criteria",
						"title": "Criteria",
						"type": "`$ANY`",
						"req": true,
					},
					map[string]any{
						"name": "criteria_from",
						"title": "Criteria From",
						"type": "`$STRING`",
						"req": true,
						"short": "The ID of the segment whose criteria should be copied.",
						"format": "uuid",
					},
					map[string]any{
						"name": "customers",
						"title": "Customers",
						"type": "`$ARRAY`",
						"req": true,
						"short": "Customer IDs to include in the fixed customer list.",
					},
					map[string]any{
						"name": "customers_from",
						"title": "Customers From",
						"type": "`$STRING`",
						"req": true,
						"short": "The ID of the segment whose current members should be copied.",
						"format": "uuid",
					},
					map[string]any{
						"name": "description",
						"title": "Description",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "fixed_customer_list",
						"title": "Fixed Customer List",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "fixed_customer_list_size",
						"title": "Fixed Customer List Size",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "from_criteria_id",
						"title": "From Criteria Id",
						"type": "`$STRING`",
						"format": "uuid",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"req": true,
						"format": "uuid",
					},
					map[string]any{
						"name": "intersection_count",
						"title": "Intersection Count",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "jaccard_index",
						"title": "Jaccard Index",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
						"req": true,
						"op": map[string]any{
							"list": map[string]any{
								"type": "`$STRING`",
							},
							"patch": map[string]any{
								"type": "`$STRING`",
							},
						},
					},
					map[string]any{
						"name": "parameters",
						"title": "Parameters",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "source",
						"title": "Source",
						"type": "`$STRING`",
						"short": "How the segment was created.",
					},
					map[string]any{
						"name": "statistics_id",
						"title": "Statistics Id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "suggestion",
						"title": "Suggestion",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tag",
						"title": "Tag",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tags",
						"title": "Tags",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "type",
						"title": "Type",
						"type": "`$STRING`",
						"short": "The shape of the statistic's values.",
					},
					map[string]any{
						"name": "union_count",
						"title": "Union Count",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "upload",
						"title": "Upload",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "values",
						"title": "Values",
						"type": "`$ARRAY`",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "drive_segments_api",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/drive-api/v1/segments/{segment_id}/query",
								"segments": []any{
									map[string]any{
										"lit": "drive-api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "segments",
									},
									map[string]any{
										"var": "segment_id",
									},
									map[string]any{
										"lit": "query",
									},
								},
								"parts": []any{
									"drive-api",
									"v1",
									"segments",
									"{segment_id}",
									"query",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
									"params": []any{
										map[string]any{
											"name": "segment_id",
											"orig": "segment_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "23303f83-a790-4ee9-95b6-d5cdeaf36fa2",
										},
									},
								},
								"select": map[string]any{
									"$action": "query",
									"exist": []any{
										"segment_id",
										"x_client_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/drive-api/v2/criteria/drafts/{criteria_id}/run",
								"segments": []any{
									map[string]any{
										"lit": "drive-api",
									},
									map[string]any{
										"lit": "v2",
									},
									map[string]any{
										"lit": "criteria",
									},
									map[string]any{
										"lit": "drafts",
									},
									map[string]any{
										"var": "draft_id",
									},
									map[string]any{
										"lit": "run",
									},
								},
								"parts": []any{
									"drive-api",
									"v2",
									"criteria",
									"drafts",
									"{draft_id}",
									"run",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"criteria_id": "draft_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
									"params": []any{
										map[string]any{
											"name": "draft_id",
											"orig": "criteria_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "0ba9049e-25e9-43f2-8ea8-1fcf9dde581e",
										},
									},
								},
								"select": map[string]any{
									"$action": "run",
									"exist": []any{
										"draft_id",
										"x_client_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/drive-api/v2/segments/{segment_id}/criteria/run",
								"segments": []any{
									map[string]any{
										"lit": "drive-api",
									},
									map[string]any{
										"lit": "v2",
									},
									map[string]any{
										"lit": "segments",
									},
									map[string]any{
										"var": "segment_id",
									},
									map[string]any{
										"lit": "criteria",
									},
									map[string]any{
										"lit": "run",
									},
								},
								"parts": []any{
									"drive-api",
									"v2",
									"segments",
									"{segment_id}",
									"criteria",
									"run",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
									"params": []any{
										map[string]any{
											"name": "segment_id",
											"orig": "segment_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "23303f83-a790-4ee9-95b6-d5cdeaf36fa2",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"segment_id",
										"x_client_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/drive-api/v1/segments/from-existing/criteria",
								"segments": []any{
									map[string]any{
										"lit": "drive-api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "segments",
									},
									map[string]any{
										"lit": "from-existing",
									},
									map[string]any{
										"lit": "criteria",
									},
								},
								"parts": []any{
									"drive-api",
									"v1",
									"segments",
									"from-existing",
									"criteria",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"x_client_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/drive-api/v1/segments/from-existing/fixed-customers",
								"segments": []any{
									map[string]any{
										"lit": "drive-api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "segments",
									},
									map[string]any{
										"lit": "from-existing",
									},
									map[string]any{
										"lit": "fixed-customers",
									},
								},
								"parts": []any{
									"drive-api",
									"v1",
									"segments",
									"from-existing",
									"fixed-customers",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"x_client_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/drive-api/v1/segments/upload",
								"segments": []any{
									map[string]any{
										"lit": "drive-api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "segments",
									},
									map[string]any{
										"lit": "upload",
									},
								},
								"parts": []any{
									"drive-api",
									"v1",
									"segments",
									"upload",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"x_client_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/drive-api/v2/criteria/drafts",
								"segments": []any{
									map[string]any{
										"lit": "drive-api",
									},
									map[string]any{
										"lit": "v2",
									},
									map[string]any{
										"lit": "criteria",
									},
									map[string]any{
										"lit": "drafts",
									},
								},
								"parts": []any{
									"drive-api",
									"v2",
									"criteria",
									"drafts",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"x_client_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/drive-api/v2/criteria/run",
								"segments": []any{
									map[string]any{
										"lit": "drive-api",
									},
									map[string]any{
										"lit": "v2",
									},
									map[string]any{
										"lit": "criteria",
									},
									map[string]any{
										"lit": "run",
									},
								},
								"parts": []any{
									"drive-api",
									"v2",
									"criteria",
									"run",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"x_client_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/drive-api/v2/criteria/suggestions/segment-name",
								"segments": []any{
									map[string]any{
										"lit": "drive-api",
									},
									map[string]any{
										"lit": "v2",
									},
									map[string]any{
										"lit": "criteria",
									},
									map[string]any{
										"lit": "suggestions",
									},
									map[string]any{
										"lit": "segment-name",
									},
								},
								"parts": []any{
									"drive-api",
									"v2",
									"criteria",
									"suggestions",
									"segment-name",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"x_client_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/drive-api/v2/segments",
								"segments": []any{
									map[string]any{
										"lit": "drive-api",
									},
									map[string]any{
										"lit": "v2",
									},
									map[string]any{
										"lit": "segments",
									},
								},
								"parts": []any{
									"drive-api",
									"v2",
									"segments",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"x_client_id",
									},
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/drive-api/v1/segments/{segment_id}/statistics/overview",
								"segments": []any{
									map[string]any{
										"lit": "drive-api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "segments",
									},
									map[string]any{
										"var": "segment_id",
									},
									map[string]any{
										"lit": "statistics",
									},
									map[string]any{
										"lit": "overview",
									},
								},
								"parts": []any{
									"drive-api",
									"v1",
									"segments",
									"{segment_id}",
									"statistics",
									"overview",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
									"params": []any{
										map[string]any{
											"name": "segment_id",
											"orig": "segment_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "23303f83-a790-4ee9-95b6-d5cdeaf36fa2",
										},
									},
									"query": []any{
										map[string]any{
											"name": "date_to",
											"orig": "date_to",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "significant_date",
											"orig": "significant_date",
											"type": "`$STRING`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"date_to",
										"segment_id",
										"significant_date",
										"x_client_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/drive-api/v2/segments",
								"segments": []any{
									map[string]any{
										"lit": "drive-api",
									},
									map[string]any{
										"lit": "v2",
									},
									map[string]any{
										"lit": "segments",
									},
								},
								"parts": []any{
									"drive-api",
									"v2",
									"segments",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
									"query": []any{
										map[string]any{
											"name": "max_per_page",
											"orig": "max_per_page",
											"type": "`$INTEGER`",
											"kind": "query",
										},
										map[string]any{
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
											"kind": "query",
										},
										map[string]any{
											"name": "search",
											"orig": "search",
											"type": "`$STRING`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"max_per_page",
										"page",
										"search",
										"x_client_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/drive-api/v1/tags",
								"segments": []any{
									map[string]any{
										"lit": "drive-api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "tags",
									},
								},
								"parts": []any{
									"drive-api",
									"v1",
									"tags",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"x_client_id",
									},
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/drive-api/v1/segments/{segment_id}/similarity/{other_segment_id}",
								"segments": []any{
									map[string]any{
										"lit": "drive-api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "segments",
									},
									map[string]any{
										"var": "segment_id",
									},
									map[string]any{
										"lit": "similarity",
									},
									map[string]any{
										"var": "other_segment_id",
									},
								},
								"parts": []any{
									"drive-api",
									"v1",
									"segments",
									"{segment_id}",
									"similarity",
									"{other_segment_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
									"params": []any{
										map[string]any{
											"name": "other_segment_id",
											"orig": "other_segment_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "7e8aedf9-6be5-4807-afe0-aa292fb2cde2",
										},
										map[string]any{
											"name": "segment_id",
											"orig": "segment_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "23303f83-a790-4ee9-95b6-d5cdeaf36fa2",
										},
									},
									"query": []any{
										map[string]any{
											"name": "refresh",
											"orig": "refresh",
											"type": "`$BOOLEAN`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"other_segment_id",
										"refresh",
										"segment_id",
										"x_client_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/drive-api/v2/criteria/drafts/{criteria_id}",
								"segments": []any{
									map[string]any{
										"lit": "drive-api",
									},
									map[string]any{
										"lit": "v2",
									},
									map[string]any{
										"lit": "criteria",
									},
									map[string]any{
										"lit": "drafts",
									},
									map[string]any{
										"var": "criteria_id",
									},
								},
								"parts": []any{
									"drive-api",
									"v2",
									"criteria",
									"drafts",
									"{criteria_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
									"params": []any{
										map[string]any{
											"name": "criteria_id",
											"orig": "criteria_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "0ba9049e-25e9-43f2-8ea8-1fcf9dde581e",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"criteria_id",
										"x_client_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/drive-api/v2/criteria/options",
								"segments": []any{
									map[string]any{
										"lit": "drive-api",
									},
									map[string]any{
										"lit": "v2",
									},
									map[string]any{
										"lit": "criteria",
									},
									map[string]any{
										"lit": "options",
									},
								},
								"parts": []any{
									"drive-api",
									"v2",
									"criteria",
									"options",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
									"query": []any{
										map[string]any{
											"name": "filter",
											"orig": "filter",
											"type": "`$STRING`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"filter",
										"x_client_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/drive-api/v1/segments/{segment_id}",
								"segments": []any{
									map[string]any{
										"lit": "drive-api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "segments",
									},
									map[string]any{
										"var": "segment_id",
									},
								},
								"parts": []any{
									"drive-api",
									"v1",
									"segments",
									"{segment_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
									"params": []any{
										map[string]any{
											"name": "segment_id",
											"orig": "segment_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "23303f83-a790-4ee9-95b6-d5cdeaf36fa2",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"segment_id",
										"x_client_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/drive-api/v2/segments/{segment_id}/criteria",
								"segments": []any{
									map[string]any{
										"lit": "drive-api",
									},
									map[string]any{
										"lit": "v2",
									},
									map[string]any{
										"lit": "segments",
									},
									map[string]any{
										"var": "segment_id",
									},
									map[string]any{
										"lit": "criteria",
									},
								},
								"parts": []any{
									"drive-api",
									"v2",
									"segments",
									"{segment_id}",
									"criteria",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
									"params": []any{
										map[string]any{
											"name": "segment_id",
											"orig": "segment_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "23303f83-a790-4ee9-95b6-d5cdeaf36fa2",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"segment_id",
										"x_client_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/drive-api/v2/segments/{segment_id}/criteria/active-options",
								"segments": []any{
									map[string]any{
										"lit": "drive-api",
									},
									map[string]any{
										"lit": "v2",
									},
									map[string]any{
										"lit": "segments",
									},
									map[string]any{
										"var": "segment_id",
									},
									map[string]any{
										"lit": "criteria",
									},
									map[string]any{
										"lit": "active-options",
									},
								},
								"parts": []any{
									"drive-api",
									"v2",
									"segments",
									"{segment_id}",
									"criteria",
									"active-options",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
									"params": []any{
										map[string]any{
											"name": "segment_id",
											"orig": "segment_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "23303f83-a790-4ee9-95b6-d5cdeaf36fa2",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"segment_id",
										"x_client_id",
									},
								},
							},
						},
					},
					"patch": map[string]any{
						"input": "data",
						"name": "patch",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/drive-api/v1/segments/{segment_id}",
								"segments": []any{
									map[string]any{
										"lit": "drive-api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "segments",
									},
									map[string]any{
										"var": "segment_id",
									},
								},
								"parts": []any{
									"drive-api",
									"v1",
									"segments",
									"{segment_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
									"params": []any{
										map[string]any{
											"name": "segment_id",
											"orig": "segment_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "23303f83-a790-4ee9-95b6-d5cdeaf36fa2",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"segment_id",
										"x_client_id",
									},
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/drive-api/v1/segments/{segment_id}",
								"segments": []any{
									map[string]any{
										"lit": "drive-api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "segments",
									},
									map[string]any{
										"var": "segment_id",
									},
								},
								"parts": []any{
									"drive-api",
									"v1",
									"segments",
									"{segment_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.metadata`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
									"params": []any{
										map[string]any{
											"name": "segment_id",
											"orig": "segment_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "23303f83-a790-4ee9-95b6-d5cdeaf36fa2",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"segment_id",
										"x_client_id",
									},
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PUT",
								"orig": "/drive-api/v2/segments/{segment_id}/criteria/{criteria_id}",
								"segments": []any{
									map[string]any{
										"lit": "drive-api",
									},
									map[string]any{
										"lit": "v2",
									},
									map[string]any{
										"lit": "segments",
									},
									map[string]any{
										"var": "segment_id",
									},
									map[string]any{
										"lit": "criteria",
									},
									map[string]any{
										"var": "criteria_id",
									},
								},
								"parts": []any{
									"drive-api",
									"v2",
									"segments",
									"{segment_id}",
									"criteria",
									"{criteria_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
									"params": []any{
										map[string]any{
											"name": "criteria_id",
											"orig": "criteria_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "0ba9049e-25e9-43f2-8ea8-1fcf9dde581e",
										},
										map[string]any{
											"name": "segment_id",
											"orig": "segment_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "23303f83-a790-4ee9-95b6-d5cdeaf36fa2",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"criteria_id",
										"segment_id",
										"x_client_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "PUT",
								"orig": "/drive-api/v2/criteria/drafts/{criteria_id}",
								"segments": []any{
									map[string]any{
										"lit": "drive-api",
									},
									map[string]any{
										"lit": "v2",
									},
									map[string]any{
										"lit": "criteria",
									},
									map[string]any{
										"lit": "drafts",
									},
									map[string]any{
										"var": "criteria_id",
									},
								},
								"parts": []any{
									"drive-api",
									"v2",
									"criteria",
									"drafts",
									"{criteria_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_client_id",
											"orig": "x_client_id",
											"type": "`$STRING`",
											"kind": "header",
											"reqd": true,
											"example": "8711bbef-b357-4c2d-97ca-0d9df4206e9a",
										},
									},
									"params": []any{
										map[string]any{
											"name": "criteria_id",
											"orig": "criteria_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "0ba9049e-25e9-43f2-8ea8-1fcf9dde581e",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"criteria_id",
										"x_client_id",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
		},
	}
}

// The plugin definitions the model selected per feature, as []any so a
// feature package can consume them without core naming its types. Empty
// when no active feature declares active plugin groups for this target.
var featurePlugins = map[string][]any{
}

// FeaturePlugins is the definitions list for one feature's chain.
func FeaturePlugins(name string) []any {
	return featurePlugins[name]
}

var (
	sharedConfigOnce sync.Once
	sharedConfigVal  map[string]any
)

// SharedConfig returns the process-wide config, built once on first use.
// The SDK reads the config on every request and never writes to it, so one
// instance is shared by every client rather than rebuilt per client.
//
// The returned map is shared: treat it as read-only. Callers that need to
// mutate should use MakeConfig, which always returns a fresh copy.
func SharedConfig() map[string]any {
	sharedConfigOnce.Do(func() {
		sharedConfigVal = MakeConfig()
	})
	return sharedConfigVal
}

func makeFeature(name string) Feature {
	switch name {
	case "debug":
		if NewDebugFeatureFunc != nil {
			return NewDebugFeatureFunc()
		}
	case "idempotency":
		if NewIdempotencyFeatureFunc != nil {
			return NewIdempotencyFeatureFunc()
		}
	case "metrics":
		if NewMetricsFeatureFunc != nil {
			return NewMetricsFeatureFunc()
		}
	case "paging":
		if NewPagingFeatureFunc != nil {
			return NewPagingFeatureFunc()
		}
	case "ratelimit":
		if NewRatelimitFeatureFunc != nil {
			return NewRatelimitFeatureFunc()
		}
	case "retry":
		if NewRetryFeatureFunc != nil {
			return NewRetryFeatureFunc()
		}
	case "test":
		if NewTestFeatureFunc != nil {
			return NewTestFeatureFunc()
		}
	case "timeout":
		if NewTimeoutFeatureFunc != nil {
			return NewTimeoutFeatureFunc()
		}
	default:
		if NewBaseFeatureFunc != nil {
			return NewBaseFeatureFunc()
		}
	}
	return nil
}
