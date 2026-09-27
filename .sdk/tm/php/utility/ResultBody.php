<?php
declare(strict_types=1);

// BudDrive SDK utility: result_body

class BudDriveResultBody
{
    public static function call(BudDriveContext $ctx): ?BudDriveResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result && $response && $response->json_func && $response->body) {
            $result->body = ($response->json_func)();
        }
        return $result;
    }
}
