<?php
declare(strict_types=1);

// BudDrive SDK exists test

require_once __DIR__ . '/../buddrive_sdk.php';

use PHPUnit\Framework\TestCase;

class ExistsTest extends TestCase
{
    public function test_create_test_sdk(): void
    {
        $testsdk = BudDriveSDK::test(null, null);
        $this->assertNotNull($testsdk);
    }
}
