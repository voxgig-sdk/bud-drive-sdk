import { BudDriveEntityBase } from '../BudDriveEntityBase';
import type { BudDriveSDK } from '../BudDriveSDK';
import type { Control } from '../types';
import type { DriveSegmentsApi, DriveSegmentsApiLoadMatch, DriveSegmentsApiListMatch, DriveSegmentsApiCreateData, DriveSegmentsApiUpdateData, DriveSegmentsApiRemoveMatch } from '../BudDriveTypes';
declare class DriveSegmentsApiEntity extends BudDriveEntityBase<DriveSegmentsApi> {
    constructor(client: BudDriveSDK, entopts: any);
    make(this: DriveSegmentsApiEntity): DriveSegmentsApiEntity;
    load(this: any, reqmatch?: DriveSegmentsApiLoadMatch, ctrl?: Control): Promise<DriveSegmentsApiEntity>;
    list(this: any, reqmatch?: DriveSegmentsApiListMatch, ctrl?: Control): Promise<DriveSegmentsApiEntity[]>;
    create(this: any, reqdata?: DriveSegmentsApiCreateData, ctrl?: Control): Promise<DriveSegmentsApiEntity>;
    update(this: any, reqdata?: DriveSegmentsApiUpdateData, ctrl?: Control): Promise<DriveSegmentsApiEntity>;
    remove(this: any, reqmatch?: DriveSegmentsApiRemoveMatch, ctrl?: Control): Promise<DriveSegmentsApiEntity>;
}
export { DriveSegmentsApiEntity };
