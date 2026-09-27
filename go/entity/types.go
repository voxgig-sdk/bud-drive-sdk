// Typed models for the BudDrive SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
// params (op.<name>.points[].g.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import (
	"encoding/json"

	"github.com/voxgig-sdk/bud-drive-sdk/go/core"
)

// DriveCampaignsApi is the typed data model for the drive_campaigns_api entity.
type DriveCampaignsApi struct {
}

// DriveCampaignsApiLoadMatch is the typed request payload for DriveCampaignsApi.LoadTyped.
type DriveCampaignsApiLoadMatch struct {
	CampaignId string `json:"campaign_id"`
}

// DriveCampaignsApiListMatch is the typed request payload for DriveCampaignsApi.ListTyped.
type DriveCampaignsApiListMatch struct {
	CreatedBy *string `json:"created_by,omitempty"`
	ItemsPerPage *int `json:"items_per_page,omitempty"`
	PageToken *string `json:"page_token,omitempty"`
	Status *string `json:"status,omitempty"`
}

// DriveCampaignsApiCreateData is the typed request payload for DriveCampaignsApi.CreateTyped.
type DriveCampaignsApiCreateData struct {
	Audience *string `json:"audience,omitempty"`
	Branding *string `json:"branding,omitempty"`
	CampaignGoal string `json:"campaign_goal"`
	CampaignId string `json:"campaign_id"`
	Channels *map[string]any `json:"channels,omitempty"`
	CompletedAt *string `json:"completed_at,omitempty"`
	Context *string `json:"context,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	Description *string `json:"description,omitempty"`
	Duration *map[string]any `json:"duration,omitempty"`
	Examples *[]any `json:"examples,omitempty"`
	ExtraContext *string `json:"extra_context,omitempty"`
	Group *string `json:"group,omitempty"`
	Headline *map[string]any `json:"headline,omitempty"`
	Id *string `json:"id,omitempty"`
	ImageStyle *string `json:"image_style,omitempty"`
	ImageUrl *string `json:"image_url,omitempty"`
	Label *string `json:"label,omitempty"`
	Message *string `json:"message,omitempty"`
	Name *string `json:"name,omitempty"`
	Parameters *[]any `json:"parameters,omitempty"`
	PrimaryButton *map[string]any `json:"primary_button,omitempty"`
	PriorityOrder *int `json:"priority_order,omitempty"`
	Result *map[string]any `json:"result,omitempty"`
	Schedule *map[string]any `json:"schedule,omitempty"`
	StartedAt *string `json:"started_at,omitempty"`
	Status *string `json:"status,omitempty"`
	Styling *map[string]any `json:"styling,omitempty"`
	TemplateId *string `json:"template_id,omitempty"`
	TemplateVariables *map[string]any `json:"template_variables,omitempty"`
	Templates *[]any `json:"templates,omitempty"`
	ToneOfVoice *string `json:"tone_of_voice,omitempty"`
	Trend *map[string]any `json:"trend,omitempty"`
	Type *string `json:"type,omitempty"`
}

// DriveCampaignsApiUpdateData is the typed request payload for DriveCampaignsApi.UpdateTyped.
type DriveCampaignsApiUpdateData struct {
	CampaignId string `json:"campaign_id"`
	Audience *string `json:"audience,omitempty"`
	Branding *string `json:"branding,omitempty"`
	CampaignGoal *string `json:"campaign_goal,omitempty"`
	Channels *map[string]any `json:"channels,omitempty"`
	CompletedAt *string `json:"completed_at,omitempty"`
	Context *string `json:"context,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	Description *string `json:"description,omitempty"`
	Duration *map[string]any `json:"duration,omitempty"`
	Examples *[]any `json:"examples,omitempty"`
	ExtraContext *string `json:"extra_context,omitempty"`
	Group *string `json:"group,omitempty"`
	Headline *map[string]any `json:"headline,omitempty"`
	Id *string `json:"id,omitempty"`
	ImageStyle *string `json:"image_style,omitempty"`
	ImageUrl *string `json:"image_url,omitempty"`
	Label *string `json:"label,omitempty"`
	Message *string `json:"message,omitempty"`
	Name *string `json:"name,omitempty"`
	Parameters *[]any `json:"parameters,omitempty"`
	PrimaryButton *map[string]any `json:"primary_button,omitempty"`
	PriorityOrder *int `json:"priority_order,omitempty"`
	Result *map[string]any `json:"result,omitempty"`
	Schedule *map[string]any `json:"schedule,omitempty"`
	StartedAt *string `json:"started_at,omitempty"`
	Status *string `json:"status,omitempty"`
	Styling *map[string]any `json:"styling,omitempty"`
	TemplateId *string `json:"template_id,omitempty"`
	TemplateVariables *map[string]any `json:"template_variables,omitempty"`
	Templates *[]any `json:"templates,omitempty"`
	ToneOfVoice *string `json:"tone_of_voice,omitempty"`
	Trend *map[string]any `json:"trend,omitempty"`
	Type *string `json:"type,omitempty"`
}

// DriveCampaignsApiRemoveMatch is the typed request payload for DriveCampaignsApi.RemoveTyped.
type DriveCampaignsApiRemoveMatch struct {
	CampaignId string `json:"campaign_id"`
}

// DriveMcpApi is the typed data model for the drive_mcp_api entity.
type DriveMcpApi struct {
}

// DriveMcpApiCreateData is the typed request payload for DriveMcpApi.CreateTyped.
type DriveMcpApiCreateData struct {
	Error *map[string]any `json:"error,omitempty"`
	Id string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Method string `json:"method"`
	Params *map[string]any `json:"params,omitempty"`
	Result *map[string]any `json:"result,omitempty"`
}

// DriveSegmentsApi is the typed data model for the drive_segments_api entity.
type DriveSegmentsApi struct {
}

// DriveSegmentsApiLoadMatch is the typed request payload for DriveSegmentsApi.LoadTyped.
type DriveSegmentsApiLoadMatch struct {
	SegmentId string `json:"segment_id"`
}

// DriveSegmentsApiListMatch is the typed request payload for DriveSegmentsApi.ListTyped.
type DriveSegmentsApiListMatch struct {
	MaxPerPage *int `json:"max_per_page,omitempty"`
	Page *int `json:"page,omitempty"`
	Search *string `json:"search,omitempty"`
}

// DriveSegmentsApiCreateData is the typed request payload for DriveSegmentsApi.CreateTyped.
type DriveSegmentsApiCreateData struct {
	CreatedAt *string `json:"created_at,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	Criteria any `json:"criteria"`
	CriteriaFrom string `json:"criteria_from"`
	Customers []any `json:"customers"`
	CustomersFrom string `json:"customers_from"`
	Description *string `json:"description,omitempty"`
	FixedCustomerList *bool `json:"fixed_customer_list,omitempty"`
	FixedCustomerListSize *int `json:"fixed_customer_list_size,omitempty"`
	FromCriteriaId *string `json:"from_criteria_id,omitempty"`
	Id string `json:"id"`
	IntersectionCount *int `json:"intersection_count,omitempty"`
	JaccardIndex *float64 `json:"jaccard_index,omitempty"`
	Name string `json:"name"`
	Parameters *map[string]any `json:"parameters,omitempty"`
	Source *string `json:"source,omitempty"`
	StatisticsId *string `json:"statistics_id,omitempty"`
	Suggestion *map[string]any `json:"suggestion,omitempty"`
	Tag *string `json:"tag,omitempty"`
	Tags *[]any `json:"tags,omitempty"`
	Type *string `json:"type,omitempty"`
	UnionCount *int `json:"union_count,omitempty"`
	Upload *map[string]any `json:"upload,omitempty"`
	Values *[]any `json:"values,omitempty"`
}

// DriveSegmentsApiUpdateData is the typed request payload for DriveSegmentsApi.UpdateTyped.
type DriveSegmentsApiUpdateData struct {
	CriteriaId string `json:"criteria_id"`
	SegmentId *string `json:"segment_id,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	Criteria *any `json:"criteria,omitempty"`
	CriteriaFrom *string `json:"criteria_from,omitempty"`
	Customers *[]any `json:"customers,omitempty"`
	CustomersFrom *string `json:"customers_from,omitempty"`
	Description *string `json:"description,omitempty"`
	FixedCustomerList *bool `json:"fixed_customer_list,omitempty"`
	FixedCustomerListSize *int `json:"fixed_customer_list_size,omitempty"`
	FromCriteriaId *string `json:"from_criteria_id,omitempty"`
	Id *string `json:"id,omitempty"`
	IntersectionCount *int `json:"intersection_count,omitempty"`
	JaccardIndex *float64 `json:"jaccard_index,omitempty"`
	Name *string `json:"name,omitempty"`
	Parameters *map[string]any `json:"parameters,omitempty"`
	Source *string `json:"source,omitempty"`
	StatisticsId *string `json:"statistics_id,omitempty"`
	Suggestion *map[string]any `json:"suggestion,omitempty"`
	Tag *string `json:"tag,omitempty"`
	Tags *[]any `json:"tags,omitempty"`
	Type *string `json:"type,omitempty"`
	UnionCount *int `json:"union_count,omitempty"`
	Upload *map[string]any `json:"upload,omitempty"`
	Values *[]any `json:"values,omitempty"`
}

// DriveSegmentsApiRemoveMatch is the typed request payload for DriveSegmentsApi.RemoveTyped.
type DriveSegmentsApiRemoveMatch struct {
	SegmentId string `json:"segment_id"`
}

// asMap turns a typed request/data struct into the map[string]any the
// runtime op pipeline consumes, honouring the json tags above.
func asMap(v any) map[string]any {
	out := map[string]any{}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// entityData unwraps an entity to its data map.
//
// Operations resolve to the ENTITY, not the raw data (see AGENTS.md), and an
// entity's fields are UNEXPORTED — marshalling one directly yields `{}`, so
// every typed accessor would silently hand back a zero-valued struct. The
// typed boundary therefore takes the data hop first.
func entityData(v any) any {
	if ent, ok := v.(core.Entity); ok {
		return ent.Data()
	}
	return v
}

// typedFrom decodes a runtime value (an entity, or the map[string]any the op
// pipeline produced) into a typed model T via a JSON round-trip. On any error
// it returns the zero value of T; the op's own (value, error) tuple carries
// the real error.
func typedFrom[T any](v any) T {
	var out T
	v = entityData(v)
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedSliceFrom decodes a runtime list value into a typed slice []T via a
// JSON round-trip, for list ops. `list` resolves to a slice of ENTITY
// instances, so each element takes the data hop.
func typedSliceFrom[T any](v any) []T {
	var out []T
	if v == nil {
		return out
	}
	if list, ok := v.([]any); ok {
		unwrapped := make([]any, 0, len(list))
		for _, item := range list {
			unwrapped = append(unwrapped, entityData(item))
		}
		v = unwrapped
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}
