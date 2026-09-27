<?php
declare(strict_types=1);

// BudDrive SDK utility: feature_hook

class BudDriveFeatureHook
{
    public static function call(BudDriveContext $ctx, string $name): void
    {
        if (!$ctx->client) {
            return;
        }
        $features = $ctx->client->features ?? null;
        if (!$features) {
            return;
        }
        foreach ($features as $f) {
            if (method_exists($f, $name)) {
                $f->$name($ctx);
            }
        }
    }
}
