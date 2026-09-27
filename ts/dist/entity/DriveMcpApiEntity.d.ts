import { BudDriveEntityBase } from '../BudDriveEntityBase';
import type { BudDriveSDK } from '../BudDriveSDK';
import type { Control } from '../types';
import type { DriveMcpApi, DriveMcpApiCreateData } from '../BudDriveTypes';
declare class DriveMcpApiEntity extends BudDriveEntityBase<DriveMcpApi> {
    constructor(client: BudDriveSDK, entopts: any);
    make(this: DriveMcpApiEntity): DriveMcpApiEntity;
    create(this: any, reqdata?: DriveMcpApiCreateData, ctrl?: Control): Promise<DriveMcpApiEntity>;
}
export { DriveMcpApiEntity };
