# Typed models for the BudDrive SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
# params (op.<name>.points[].g.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.
#
# These are TypedDicts, not dataclasses: the SDK ops return/accept plain dicts
# at runtime, and a TypedDict IS a dict shape, so the types match the runtime.
# Optional (req:false) keys are modelled as TypedDict key-optionality
# (total=False), split into a required base + total=False subclass when a type
# has both required and optional keys.

from __future__ import annotations

from typing import TypedDict, Any


class DriveCampaignsApiRequired(TypedDict):
    campaign_goal: str
    campaign_id: str


class DriveCampaignsApi(DriveCampaignsApiRequired, total=False):
    audience: str
    branding: str
    channels: dict
    completed_at: str
    context: str
    created_at: str
    created_by: str
    description: str
    duration: dict
    examples: list
    extra_context: str
    group: str
    headline: dict
    id: str
    image_style: str
    image_url: str
    label: str
    message: str
    name: str
    parameters: list
    primary_button: dict
    priority_order: int
    result: dict
    schedule: dict
    started_at: str
    status: str
    styling: dict
    template_id: str
    template_variables: dict
    templates: list
    tone_of_voice: str
    trend: dict
    type: str


class DriveCampaignsApiLoadMatch(TypedDict):
    campaign_id: str


class DriveCampaignsApiListMatch(TypedDict, total=False):
    created_by: str
    items_per_page: int
    page_token: str
    status: str


class DriveCampaignsApiCreateDataRequired(TypedDict):
    campaign_goal: str
    campaign_id: str


class DriveCampaignsApiCreateData(DriveCampaignsApiCreateDataRequired, total=False):
    audience: str
    branding: str
    channels: dict
    completed_at: str
    context: str
    created_at: str
    created_by: str
    description: str
    duration: dict
    examples: list
    extra_context: str
    group: str
    headline: dict
    id: str
    image_style: str
    image_url: str
    label: str
    message: str
    name: str
    parameters: list
    primary_button: dict
    priority_order: int
    result: dict
    schedule: dict
    started_at: str
    status: str
    styling: dict
    template_id: str
    template_variables: dict
    templates: list
    tone_of_voice: str
    trend: dict
    type: str


class DriveCampaignsApiUpdateDataRequired(TypedDict):
    campaign_id: str


class DriveCampaignsApiUpdateData(DriveCampaignsApiUpdateDataRequired, total=False):
    audience: str
    branding: str
    campaign_goal: str
    channels: dict
    completed_at: str
    context: str
    created_at: str
    created_by: str
    description: str
    duration: dict
    examples: list
    extra_context: str
    group: str
    headline: dict
    id: str
    image_style: str
    image_url: str
    label: str
    message: str
    name: str
    parameters: list
    primary_button: dict
    priority_order: int
    result: dict
    schedule: dict
    started_at: str
    status: str
    styling: dict
    template_id: str
    template_variables: dict
    templates: list
    tone_of_voice: str
    trend: dict
    type: str


class DriveCampaignsApiRemoveMatch(TypedDict):
    campaign_id: str


class DriveMcpApiRequired(TypedDict):
    id: str
    jsonrpc: str
    method: str


class DriveMcpApi(DriveMcpApiRequired, total=False):
    error: dict
    params: dict
    result: dict


class DriveMcpApiCreateDataRequired(TypedDict):
    id: str
    jsonrpc: str
    method: str


class DriveMcpApiCreateData(DriveMcpApiCreateDataRequired, total=False):
    error: dict
    params: dict
    result: dict


class DriveSegmentsApiRequired(TypedDict):
    criteria: Any
    criteria_from: str
    customers: list
    customers_from: str
    id: str
    name: str


class DriveSegmentsApi(DriveSegmentsApiRequired, total=False):
    created_at: str
    created_by: str
    description: str
    fixed_customer_list: bool
    fixed_customer_list_size: int
    from_criteria_id: str
    intersection_count: int
    jaccard_index: float
    parameters: dict
    source: str
    statistics_id: str
    suggestion: dict
    tag: str
    tags: list
    type: str
    union_count: int
    upload: dict
    values: list


class DriveSegmentsApiLoadMatch(TypedDict):
    segment_id: str


class DriveSegmentsApiListMatch(TypedDict, total=False):
    max_per_page: int
    page: int
    search: str


class DriveSegmentsApiCreateDataRequired(TypedDict):
    criteria: Any
    criteria_from: str
    customers: list
    customers_from: str
    id: str
    name: str


class DriveSegmentsApiCreateData(DriveSegmentsApiCreateDataRequired, total=False):
    created_at: str
    created_by: str
    description: str
    fixed_customer_list: bool
    fixed_customer_list_size: int
    from_criteria_id: str
    intersection_count: int
    jaccard_index: float
    parameters: dict
    source: str
    statistics_id: str
    suggestion: dict
    tag: str
    tags: list
    type: str
    union_count: int
    upload: dict
    values: list


class DriveSegmentsApiUpdateDataRequired(TypedDict):
    criteria_id: str


class DriveSegmentsApiUpdateData(DriveSegmentsApiUpdateDataRequired, total=False):
    segment_id: str
    created_at: str
    created_by: str
    criteria: Any
    criteria_from: str
    customers: list
    customers_from: str
    description: str
    fixed_customer_list: bool
    fixed_customer_list_size: int
    from_criteria_id: str
    id: str
    intersection_count: int
    jaccard_index: float
    name: str
    parameters: dict
    source: str
    statistics_id: str
    suggestion: dict
    tag: str
    tags: list
    type: str
    union_count: int
    upload: dict
    values: list


class DriveSegmentsApiRemoveMatch(TypedDict):
    segment_id: str
