import { BudDriveEntityBase } from '../BudDriveEntityBase';
import type { BudDriveSDK } from '../BudDriveSDK';
import type { Control } from '../types';
import type { DriveCampaignsApi, DriveCampaignsApiLoadMatch, DriveCampaignsApiListMatch, DriveCampaignsApiCreateData, DriveCampaignsApiUpdateData, DriveCampaignsApiRemoveMatch } from '../BudDriveTypes';
declare class DriveCampaignsApiEntity extends BudDriveEntityBase<DriveCampaignsApi> {
    constructor(client: BudDriveSDK, entopts: any);
    make(this: DriveCampaignsApiEntity): DriveCampaignsApiEntity;
    load(this: any, reqmatch?: DriveCampaignsApiLoadMatch, ctrl?: Control): Promise<DriveCampaignsApiEntity>;
    list(this: any, reqmatch?: DriveCampaignsApiListMatch, ctrl?: Control): Promise<DriveCampaignsApiEntity[]>;
    create(this: any, reqdata?: DriveCampaignsApiCreateData, ctrl?: Control): Promise<DriveCampaignsApiEntity>;
    update(this: any, reqdata?: DriveCampaignsApiUpdateData, ctrl?: Control): Promise<DriveCampaignsApiEntity>;
    remove(this: any, reqmatch?: DriveCampaignsApiRemoveMatch, ctrl?: Control): Promise<DriveCampaignsApiEntity>;
}
export { DriveCampaignsApiEntity };
