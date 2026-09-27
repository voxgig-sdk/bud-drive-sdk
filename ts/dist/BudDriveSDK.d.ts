import { DriveCampaignsApiEntity } from './entity/DriveCampaignsApiEntity';
import { DriveMcpApiEntity } from './entity/DriveMcpApiEntity';
import { DriveSegmentsApiEntity } from './entity/DriveSegmentsApiEntity';
export type * from './BudDriveTypes';
import { inspect } from 'node:util';
import type { Context, Feature } from './types';
import { config } from './Config';
import { BudDriveEntityBase } from './BudDriveEntityBase';
import { Utility } from './utility/Utility';
import { BaseFeature } from './feature/base/BaseFeature';
declare const stdutil: Utility;
declare class BudDriveSDK {
    _mode: string;
    _options: any;
    _utility: Utility;
    _features: Feature[];
    _rootctx: Context;
    constructor(options?: any);
    options(): any;
    utility(): any;
    prepare(fetchargs?: any): Promise<any>;
    direct(fetchargs?: any): Promise<Error | {
        ok: boolean;
        status: number;
        headers: any;
        data: any;
        err?: undefined;
    } | {
        ok: boolean;
        err: any;
        status?: undefined;
        headers?: undefined;
        data?: undefined;
    }>;
    _rawRequest(fetchargs?: any): Promise<Error | {
        ok: boolean;
        status: number;
        headers: any;
        data: any;
        err?: undefined;
    } | {
        ok: boolean;
        err: any;
        status?: undefined;
        headers?: undefined;
        data?: undefined;
    }>;
    graphql(query: string, variables?: any, ctrl?: any): Promise<any>;
    DriveCampaignsApi(entopts?: Record<string, any>): DriveCampaignsApiEntity;
    DriveMcpApi(entopts?: Record<string, any>): DriveMcpApiEntity;
    DriveSegmentsApi(entopts?: Record<string, any>): DriveSegmentsApiEntity;
    static test(testoptsarg?: any, sdkoptsarg?: any): BudDriveSDK;
    tester(testopts?: any, sdkopts?: any): BudDriveSDK;
    toJSON(): {
        name: string;
    };
    toString(): string;
    [inspect.custom](): string;
}
declare const SDK: typeof BudDriveSDK;
export { stdutil, config, BaseFeature, BudDriveEntityBase, BudDriveSDK, SDK, };
