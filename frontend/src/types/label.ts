import { LabelId } from "./id";
import { Principal } from "./principal";

export type LabelKey = string;
export type LabelValue = string;

export type Label = {
  id: LabelId;

  // Standard fields
  creator: Principal;
  createdTs: number;
  updater: Principal;
  updatedTs: number;

  // Domain specific fields
  key: LabelKey;
  valueList: LabelValue[];
};

export type LabelPatch = {
  valueList?: LabelValue[];
};

export type DatabaseLabel = {
  key: LabelKey;
  value: LabelValue;
};

export type AvailableLabel = {
  key: LabelKey;
  valueList: LabelValue[];
};
