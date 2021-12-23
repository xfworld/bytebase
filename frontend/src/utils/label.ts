import { uniqBy } from "lodash-es";
import { DatabaseLabel } from "../types";

export const validateLabels = (labels: DatabaseLabel[]): string | undefined => {
  for (let i = 0; i < labels.length; i++) {
    const label = labels[i];
    if (!label.key) return "database.label-error.key-necessary";
    if (!label.value) return "database.label-error.value-necessary";
  }
  if (labels.length !== uniqBy(labels, "key").length) {
    return "database.label-error.key-duplicated";
  }
  return undefined;
};
