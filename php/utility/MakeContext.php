<?php
declare(strict_types=1);

// BudDrive SDK utility: make_context

require_once __DIR__ . '/../core/Context.php';

class BudDriveMakeContext
{
    public static function call(array $ctxmap, ?BudDriveContext $basectx): BudDriveContext
    {
        return new BudDriveContext($ctxmap, $basectx);
    }
}
