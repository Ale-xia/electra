import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgRecord3 } from "./types/electra/meter/tx";
import { MsgDeletePowerPurchaseContract } from "./types/electra/meter/tx";
import { MsgCreatePowerPurchaseContract } from "./types/electra/meter/tx";
import { MsgUpdatePowerPurchaseContract } from "./types/electra/meter/tx";
import { MsgRecord } from "./types/electra/meter/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/electra.meter.MsgRecord3", MsgRecord3],
    ["/electra.meter.MsgDeletePowerPurchaseContract", MsgDeletePowerPurchaseContract],
    ["/electra.meter.MsgCreatePowerPurchaseContract", MsgCreatePowerPurchaseContract],
    ["/electra.meter.MsgUpdatePowerPurchaseContract", MsgUpdatePowerPurchaseContract],
    ["/electra.meter.MsgRecord", MsgRecord],
    
];

export { msgTypes }