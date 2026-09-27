<?php
declare(strict_types=1);

// Typed models for the BudDrive SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
// params (op.<name>.points[].g.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
//
// These are documentation-grade value objects (PHP 8 typed properties),
// registered on the composer classmap autoload. The SDK boundary exchanges
// assoc-arrays; these classes name the shapes for tooling and typed callers.

/** DriveCampaignsApi entity data model. */
class DriveCampaignsApi
{
    public ?string $audience = null;
    public ?string $branding = null;
    public string $campaign_goal;
    public string $campaign_id;
    public ?array $channels = null;
    public ?string $completed_at = null;
    public ?string $context = null;
    public ?string $created_at = null;
    public ?string $created_by = null;
    public ?string $description = null;
    public ?array $duration = null;
    public ?array $examples = null;
    public ?string $extra_context = null;
    public ?string $group = null;
    public ?array $headline = null;
    public ?string $id = null;
    public ?string $image_style = null;
    public ?string $image_url = null;
    public ?string $label = null;
    public ?string $message = null;
    public ?string $name = null;
    public ?array $parameters = null;
    public ?array $primary_button = null;
    public ?int $priority_order = null;
    public ?array $result = null;
    public ?array $schedule = null;
    public ?string $started_at = null;
    public ?string $status = null;
    public ?array $styling = null;
    public ?string $template_id = null;
    public ?array $template_variables = null;
    public ?array $templates = null;
    public ?string $tone_of_voice = null;
    public ?array $trend = null;
    public ?string $type = null;
}

/** Request payload for DriveCampaignsApi#load. */
class DriveCampaignsApiLoadMatch
{
    public string $campaign_id;
}

/** Request payload for DriveCampaignsApi#list. */
class DriveCampaignsApiListMatch
{
    public ?string $created_by = null;
    public ?int $items_per_page = null;
    public ?string $page_token = null;
    public ?string $status = null;
}

/** Request payload for DriveCampaignsApi#create. */
class DriveCampaignsApiCreateData
{
    public ?string $audience = null;
    public ?string $branding = null;
    public string $campaign_goal;
    public string $campaign_id;
    public ?array $channels = null;
    public ?string $completed_at = null;
    public ?string $context = null;
    public ?string $created_at = null;
    public ?string $created_by = null;
    public ?string $description = null;
    public ?array $duration = null;
    public ?array $examples = null;
    public ?string $extra_context = null;
    public ?string $group = null;
    public ?array $headline = null;
    public ?string $id = null;
    public ?string $image_style = null;
    public ?string $image_url = null;
    public ?string $label = null;
    public ?string $message = null;
    public ?string $name = null;
    public ?array $parameters = null;
    public ?array $primary_button = null;
    public ?int $priority_order = null;
    public ?array $result = null;
    public ?array $schedule = null;
    public ?string $started_at = null;
    public ?string $status = null;
    public ?array $styling = null;
    public ?string $template_id = null;
    public ?array $template_variables = null;
    public ?array $templates = null;
    public ?string $tone_of_voice = null;
    public ?array $trend = null;
    public ?string $type = null;
}

/** Request payload for DriveCampaignsApi#update. */
class DriveCampaignsApiUpdateData
{
    public string $campaign_id;
    public ?string $audience = null;
    public ?string $branding = null;
    public ?string $campaign_goal = null;
    public ?array $channels = null;
    public ?string $completed_at = null;
    public ?string $context = null;
    public ?string $created_at = null;
    public ?string $created_by = null;
    public ?string $description = null;
    public ?array $duration = null;
    public ?array $examples = null;
    public ?string $extra_context = null;
    public ?string $group = null;
    public ?array $headline = null;
    public ?string $id = null;
    public ?string $image_style = null;
    public ?string $image_url = null;
    public ?string $label = null;
    public ?string $message = null;
    public ?string $name = null;
    public ?array $parameters = null;
    public ?array $primary_button = null;
    public ?int $priority_order = null;
    public ?array $result = null;
    public ?array $schedule = null;
    public ?string $started_at = null;
    public ?string $status = null;
    public ?array $styling = null;
    public ?string $template_id = null;
    public ?array $template_variables = null;
    public ?array $templates = null;
    public ?string $tone_of_voice = null;
    public ?array $trend = null;
    public ?string $type = null;
}

/** Request payload for DriveCampaignsApi#remove. */
class DriveCampaignsApiRemoveMatch
{
    public string $campaign_id;
}

/** DriveMcpApi entity data model. */
class DriveMcpApi
{
    public ?array $error = null;
    public string $id;
    public string $jsonrpc;
    public string $method;
    public ?array $params = null;
    public ?array $result = null;
}

/** Request payload for DriveMcpApi#create. */
class DriveMcpApiCreateData
{
    public ?array $error = null;
    public string $id;
    public string $jsonrpc;
    public string $method;
    public ?array $params = null;
    public ?array $result = null;
}

/** DriveSegmentsApi entity data model. */
class DriveSegmentsApi
{
    public ?string $created_at = null;
    public ?string $created_by = null;
    public mixed $criteria;
    public string $criteria_from;
    public array $customers;
    public string $customers_from;
    public ?string $description = null;
    public ?bool $fixed_customer_list = null;
    public ?int $fixed_customer_list_size = null;
    public ?string $from_criteria_id = null;
    public string $id;
    public ?int $intersection_count = null;
    public ?float $jaccard_index = null;
    public string $name;
    public ?array $parameters = null;
    public ?string $source = null;
    public ?string $statistics_id = null;
    public ?array $suggestion = null;
    public ?string $tag = null;
    public ?array $tags = null;
    public ?string $type = null;
    public ?int $union_count = null;
    public ?array $upload = null;
    public ?array $values = null;
}

/** Request payload for DriveSegmentsApi#load. */
class DriveSegmentsApiLoadMatch
{
    public string $segment_id;
}

/** Request payload for DriveSegmentsApi#list. */
class DriveSegmentsApiListMatch
{
    public ?int $max_per_page = null;
    public ?int $page = null;
    public ?string $search = null;
}

/** Request payload for DriveSegmentsApi#create. */
class DriveSegmentsApiCreateData
{
    public ?string $created_at = null;
    public ?string $created_by = null;
    public mixed $criteria;
    public string $criteria_from;
    public array $customers;
    public string $customers_from;
    public ?string $description = null;
    public ?bool $fixed_customer_list = null;
    public ?int $fixed_customer_list_size = null;
    public ?string $from_criteria_id = null;
    public string $id;
    public ?int $intersection_count = null;
    public ?float $jaccard_index = null;
    public string $name;
    public ?array $parameters = null;
    public ?string $source = null;
    public ?string $statistics_id = null;
    public ?array $suggestion = null;
    public ?string $tag = null;
    public ?array $tags = null;
    public ?string $type = null;
    public ?int $union_count = null;
    public ?array $upload = null;
    public ?array $values = null;
}

/** Request payload for DriveSegmentsApi#update. */
class DriveSegmentsApiUpdateData
{
    public string $criteria_id;
    public ?string $segment_id = null;
    public ?string $created_at = null;
    public ?string $created_by = null;
    public mixed $criteria = null;
    public ?string $criteria_from = null;
    public ?array $customers = null;
    public ?string $customers_from = null;
    public ?string $description = null;
    public ?bool $fixed_customer_list = null;
    public ?int $fixed_customer_list_size = null;
    public ?string $from_criteria_id = null;
    public ?string $id = null;
    public ?int $intersection_count = null;
    public ?float $jaccard_index = null;
    public ?string $name = null;
    public ?array $parameters = null;
    public ?string $source = null;
    public ?string $statistics_id = null;
    public ?array $suggestion = null;
    public ?string $tag = null;
    public ?array $tags = null;
    public ?string $type = null;
    public ?int $union_count = null;
    public ?array $upload = null;
    public ?array $values = null;
}

/** Request payload for DriveSegmentsApi#remove. */
class DriveSegmentsApiRemoveMatch
{
    public string $segment_id;
}

