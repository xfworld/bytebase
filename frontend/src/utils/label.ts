import { countBy, uniqBy } from "lodash-es";
import { Database, DatabaseLabel, Label } from "../types";

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

export const findDefaultGroupByLabel = (
  labelList: Label[],
  databaseList: Database[]
): string | undefined => {
  // concat all databases' keys into one array
  const databaseLabelKeys = databaseList.flatMap((db) =>
    db.labels.map((label) => label.key)
  );
  if (databaseLabelKeys.length > 0) {
    // counting up the keys' frequency
    const countsDict = countBy(databaseLabelKeys);
    const countsList = Object.keys(countsDict).map((key) => ({
      key,
      count: countsDict[key],
    }));
    // return the most frequent used key
    countsList.sort((a, b) => b.count - a.count);
    return countsList[0].key;
  } else {
    // just use the first label key
    return labelList[0]?.key;
  }
};
