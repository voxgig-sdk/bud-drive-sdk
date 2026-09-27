// Typed models for the BudDrive SDK (JSDoc typedefs).
//
// GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
// params (op.<name>.points[].g.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Annotations only — no runtime effect. Do not
// edit by hand.

/**
 * @typedef {Object} DriveCampaignsApi
 * @property {string} [audience]
 * @property {string} [branding]
 * @property {string} campaign_goal
 * @property {string} campaign_id
 * @property {Object} [channels]
 * @property {string} [completed_at]
 * @property {string} [context]
 * @property {string} [created_at]
 * @property {string} [created_by]
 * @property {string} [description]
 * @property {Object} [duration]
 * @property {Array} [examples]
 * @property {string} [extra_context]
 * @property {string} [group]
 * @property {Object} [headline]
 * @property {string} [id]
 * @property {string} [image_style]
 * @property {string} [image_url]
 * @property {string} [label]
 * @property {string} [message]
 * @property {string} [name]
 * @property {Array} [parameters]
 * @property {Object} [primary_button]
 * @property {number} [priority_order]
 * @property {Object} [result]
 * @property {Object} [schedule]
 * @property {string} [started_at]
 * @property {string} [status]
 * @property {Object} [styling]
 * @property {string} [template_id]
 * @property {Object} [template_variables]
 * @property {Array} [templates]
 * @property {string} [tone_of_voice]
 * @property {Object} [trend]
 * @property {string} [type]
 */

/**
 * @typedef {Object} DriveCampaignsApiLoadMatch
 * @property {string} campaign_id
 */

/**
 * @typedef {Object} DriveCampaignsApiListMatch
 * @property {string} [created_by]
 * @property {number} [items_per_page]
 * @property {string} [page_token]
 * @property {string} [status]
 */

/**
 * @typedef {Object} DriveCampaignsApiCreateData
 * @property {string} [audience]
 * @property {string} [branding]
 * @property {string} campaign_goal
 * @property {string} campaign_id
 * @property {Object} [channels]
 * @property {string} [completed_at]
 * @property {string} [context]
 * @property {string} [created_at]
 * @property {string} [created_by]
 * @property {string} [description]
 * @property {Object} [duration]
 * @property {Array} [examples]
 * @property {string} [extra_context]
 * @property {string} [group]
 * @property {Object} [headline]
 * @property {string} [id]
 * @property {string} [image_style]
 * @property {string} [image_url]
 * @property {string} [label]
 * @property {string} [message]
 * @property {string} [name]
 * @property {Array} [parameters]
 * @property {Object} [primary_button]
 * @property {number} [priority_order]
 * @property {Object} [result]
 * @property {Object} [schedule]
 * @property {string} [started_at]
 * @property {string} [status]
 * @property {Object} [styling]
 * @property {string} [template_id]
 * @property {Object} [template_variables]
 * @property {Array} [templates]
 * @property {string} [tone_of_voice]
 * @property {Object} [trend]
 * @property {string} [type]
 */

/**
 * @typedef {Object} DriveCampaignsApiUpdateData
 * @property {string} campaign_id
 * @property {string} [audience]
 * @property {string} [branding]
 * @property {string} [campaign_goal]
 * @property {Object} [channels]
 * @property {string} [completed_at]
 * @property {string} [context]
 * @property {string} [created_at]
 * @property {string} [created_by]
 * @property {string} [description]
 * @property {Object} [duration]
 * @property {Array} [examples]
 * @property {string} [extra_context]
 * @property {string} [group]
 * @property {Object} [headline]
 * @property {string} [id]
 * @property {string} [image_style]
 * @property {string} [image_url]
 * @property {string} [label]
 * @property {string} [message]
 * @property {string} [name]
 * @property {Array} [parameters]
 * @property {Object} [primary_button]
 * @property {number} [priority_order]
 * @property {Object} [result]
 * @property {Object} [schedule]
 * @property {string} [started_at]
 * @property {string} [status]
 * @property {Object} [styling]
 * @property {string} [template_id]
 * @property {Object} [template_variables]
 * @property {Array} [templates]
 * @property {string} [tone_of_voice]
 * @property {Object} [trend]
 * @property {string} [type]
 */

/**
 * @typedef {Object} DriveCampaignsApiRemoveMatch
 * @property {string} campaign_id
 */

/**
 * @typedef {Object} DriveMcpApi
 * @property {Object} [error]
 * @property {string} id
 * @property {string} jsonrpc
 * @property {string} method
 * @property {Object} [params]
 * @property {Object} [result]
 */

/**
 * @typedef {Object} DriveMcpApiCreateData
 * @property {Object} [error]
 * @property {string} id
 * @property {string} jsonrpc
 * @property {string} method
 * @property {Object} [params]
 * @property {Object} [result]
 */

/**
 * @typedef {Object} DriveSegmentsApi
 * @property {string} [created_at]
 * @property {string} [created_by]
 * @property {*} criteria
 * @property {string} criteria_from
 * @property {Array} customers
 * @property {string} customers_from
 * @property {string} [description]
 * @property {boolean} [fixed_customer_list]
 * @property {number} [fixed_customer_list_size]
 * @property {string} [from_criteria_id]
 * @property {string} id
 * @property {number} [intersection_count]
 * @property {number} [jaccard_index]
 * @property {string} name
 * @property {Object} [parameters]
 * @property {string} [source]
 * @property {string} [statistics_id]
 * @property {Object} [suggestion]
 * @property {string} [tag]
 * @property {Array} [tags]
 * @property {string} [type]
 * @property {number} [union_count]
 * @property {Object} [upload]
 * @property {Array} [values]
 */

/**
 * @typedef {Object} DriveSegmentsApiLoadMatch
 * @property {string} segment_id
 */

/**
 * @typedef {Object} DriveSegmentsApiListMatch
 * @property {number} [max_per_page]
 * @property {number} [page]
 * @property {string} [search]
 */

/**
 * @typedef {Object} DriveSegmentsApiCreateData
 * @property {string} [created_at]
 * @property {string} [created_by]
 * @property {*} criteria
 * @property {string} criteria_from
 * @property {Array} customers
 * @property {string} customers_from
 * @property {string} [description]
 * @property {boolean} [fixed_customer_list]
 * @property {number} [fixed_customer_list_size]
 * @property {string} [from_criteria_id]
 * @property {string} id
 * @property {number} [intersection_count]
 * @property {number} [jaccard_index]
 * @property {string} name
 * @property {Object} [parameters]
 * @property {string} [source]
 * @property {string} [statistics_id]
 * @property {Object} [suggestion]
 * @property {string} [tag]
 * @property {Array} [tags]
 * @property {string} [type]
 * @property {number} [union_count]
 * @property {Object} [upload]
 * @property {Array} [values]
 */

/**
 * @typedef {Object} DriveSegmentsApiUpdateData
 * @property {string} criteria_id
 * @property {string} [segment_id]
 * @property {string} [created_at]
 * @property {string} [created_by]
 * @property {*} [criteria]
 * @property {string} [criteria_from]
 * @property {Array} [customers]
 * @property {string} [customers_from]
 * @property {string} [description]
 * @property {boolean} [fixed_customer_list]
 * @property {number} [fixed_customer_list_size]
 * @property {string} [from_criteria_id]
 * @property {string} [id]
 * @property {number} [intersection_count]
 * @property {number} [jaccard_index]
 * @property {string} [name]
 * @property {Object} [parameters]
 * @property {string} [source]
 * @property {string} [statistics_id]
 * @property {Object} [suggestion]
 * @property {string} [tag]
 * @property {Array} [tags]
 * @property {string} [type]
 * @property {number} [union_count]
 * @property {Object} [upload]
 * @property {Array} [values]
 */

/**
 * @typedef {Object} DriveSegmentsApiRemoveMatch
 * @property {string} segment_id
 */

