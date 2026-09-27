// Typed models for the BudDrive SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
// params (op.<name>.points[].g.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.

export interface DriveCampaignsApi {
  audience?: string
  branding?: string
  campaign_goal: string
  campaign_id: string
  channels?: Record<string, any>
  completed_at?: string
  context?: string
  created_at?: string
  created_by?: string
  description?: string
  duration?: Record<string, any>
  examples?: any[]
  extra_context?: string
  group?: string
  headline?: Record<string, any>
  id?: string
  image_style?: string
  image_url?: string
  label?: string
  message?: string
  name?: string
  parameters?: any[]
  primary_button?: Record<string, any>
  priority_order?: number
  result?: Record<string, any>
  schedule?: Record<string, any>
  started_at?: string
  status?: string
  styling?: Record<string, any>
  template_id?: string
  template_variables?: Record<string, any>
  templates?: any[]
  tone_of_voice?: string
  trend?: Record<string, any>
  type?: string
}

export interface DriveCampaignsApiLoadMatch {
  campaign_id: string

  // Selects a custom action instead of the plain load:
  //   'content'
  // The remaining keys are that action's own payload.
  $action?: string
  [action: string]: any
}

export interface DriveCampaignsApiListMatch {
  created_by?: string
  items_per_page?: number
  page_token?: string
  status?: string
}

export interface DriveCampaignsApiCreateData {
  audience?: string
  branding?: string
  campaign_goal: string
  campaign_id: string
  channels?: Record<string, any>
  completed_at?: string
  context?: string
  created_at?: string
  created_by?: string
  description?: string
  duration?: Record<string, any>
  examples?: any[]
  extra_context?: string
  group?: string
  headline?: Record<string, any>
  id?: string
  image_style?: string
  image_url?: string
  label?: string
  message?: string
  name?: string
  parameters?: any[]
  primary_button?: Record<string, any>
  priority_order?: number
  result?: Record<string, any>
  schedule?: Record<string, any>
  started_at?: string
  status?: string
  styling?: Record<string, any>
  template_id?: string
  template_variables?: Record<string, any>
  templates?: any[]
  tone_of_voice?: string
  trend?: Record<string, any>
  type?: string
}

export interface DriveCampaignsApiUpdateData {
  campaign_id: string
  audience?: string
  branding?: string
  campaign_goal?: string
  channels?: Record<string, any>
  completed_at?: string
  context?: string
  created_at?: string
  created_by?: string
  description?: string
  duration?: Record<string, any>
  examples?: any[]
  extra_context?: string
  group?: string
  headline?: Record<string, any>
  id?: string
  image_style?: string
  image_url?: string
  label?: string
  message?: string
  name?: string
  parameters?: any[]
  primary_button?: Record<string, any>
  priority_order?: number
  result?: Record<string, any>
  schedule?: Record<string, any>
  started_at?: string
  status?: string
  styling?: Record<string, any>
  template_id?: string
  template_variables?: Record<string, any>
  templates?: any[]
  tone_of_voice?: string
  trend?: Record<string, any>
  type?: string

  // Selects a custom action instead of the plain update:
  //   'content' | 'status'
  // The remaining keys are that action's own payload.
  $action?: string
  [action: string]: any
}

export interface DriveCampaignsApiRemoveMatch {
  campaign_id: string
}

export interface DriveMcpApi {
  error?: Record<string, any>
  id: string
  jsonrpc: string
  method: string
  params?: Record<string, any>
  result?: Record<string, any>
}

export interface DriveMcpApiCreateData {
  error?: Record<string, any>
  id: string
  jsonrpc: string
  method: string
  params?: Record<string, any>
  result?: Record<string, any>
}

export interface DriveSegmentsApi {
  created_at?: string
  created_by?: string
  criteria: any
  criteria_from: string
  customers: any[]
  customers_from: string
  description?: string
  fixed_customer_list?: boolean
  fixed_customer_list_size?: number
  from_criteria_id?: string
  id: string
  intersection_count?: number
  jaccard_index?: number
  name: string
  parameters?: Record<string, any>
  source?: string
  statistics_id?: string
  suggestion?: Record<string, any>
  tag?: string
  tags?: any[]
  type?: string
  union_count?: number
  upload?: Record<string, any>
  values?: any[]
}

export interface DriveSegmentsApiLoadMatch {
  segment_id: string
}

export interface DriveSegmentsApiListMatch {
  max_per_page?: number
  page?: number
  search?: string
}

export interface DriveSegmentsApiCreateData {
  created_at?: string
  created_by?: string
  criteria: any
  criteria_from: string
  customers: any[]
  customers_from: string
  description?: string
  fixed_customer_list?: boolean
  fixed_customer_list_size?: number
  from_criteria_id?: string
  id: string
  intersection_count?: number
  jaccard_index?: number
  name: string
  parameters?: Record<string, any>
  source?: string
  statistics_id?: string
  suggestion?: Record<string, any>
  tag?: string
  tags?: any[]
  type?: string
  union_count?: number
  upload?: Record<string, any>
  values?: any[]

  // Selects a custom action instead of the plain create:
  //   'query' | 'run'
  // The remaining keys are that action's own payload.
  $action?: string
  [action: string]: any
}

export interface DriveSegmentsApiUpdateData {
  criteria_id: string
  segment_id?: string
  created_at?: string
  created_by?: string
  criteria?: any
  criteria_from?: string
  customers?: any[]
  customers_from?: string
  description?: string
  fixed_customer_list?: boolean
  fixed_customer_list_size?: number
  from_criteria_id?: string
  id?: string
  intersection_count?: number
  jaccard_index?: number
  name?: string
  parameters?: Record<string, any>
  source?: string
  statistics_id?: string
  suggestion?: Record<string, any>
  tag?: string
  tags?: any[]
  type?: string
  union_count?: number
  upload?: Record<string, any>
  values?: any[]
}

export interface DriveSegmentsApiRemoveMatch {
  segment_id: string
}

