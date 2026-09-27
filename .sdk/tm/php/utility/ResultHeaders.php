<?php
declare(strict_types=1);

// BudDrive SDK utility: result_headers

class BudDriveResultHeaders
{
    public static function call(BudDriveContext $ctx): ?BudDriveResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result) {
            if ($response && is_array($response->headers)) {
                $result->headers = $response->headers;
            } else {
                $result->headers = [];
            }
        }
        return $result;
    }
}
