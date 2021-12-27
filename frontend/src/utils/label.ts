import { countBy, uniqBy } from "lodash-es";
import {
  Database,
  DatabaseLabel,
  Label,
  LabelSelector,
  LabelSelectorRequirement,
} from "../types";

export const RESERVED_LABEL_ID = 0;

export const isReservedLabel = (label: Label): boolean => {
  return label.id === RESERVED_LABEL_ID;
};

export const isReservedDatabaseLabel = (
  dbLabel: DatabaseLabel,
  labelList: Label[]
): boolean => {
  const label = labelList.find((label) => label.key === dbLabel.key);
  if (!label) return false;
  return label.id === RESERVED_LABEL_ID;
};

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

export const filterDatabaseListByLabelSelector = (
  databaseList: Database[],
  labelSelector: LabelSelector
): Database[] => {
  return databaseList.filter((db) =>
    isDatabaseMatchesSelector(db, labelSelector)
  );
};

export const isDatabaseMatchesSelector = (
  database: Database,
  selector: LabelSelector
): boolean => {
  const rules = selector.matchExpressions;
  return rules.every((rule) => {
    switch (rule.operator) {
      case "In":
        return checkLabelIn(database, rule);
      case "Exists":
        return checkLabelExists(database, rule);
      default:
        // unknown operators are taken as mismatch
        console.warn(`known operator "${rule.operator}"`);
        return false;
    }
  });
};

const checkLabelIn = (
  db: Database,
  rule: LabelSelectorRequirement
): boolean => {
  if (rule.key === "bb.environment") {
    return rule.values.some((env) => env === db.instance.environment.name);
  }

  const label = db.labels.find((label) => label.key === rule.key);
  if (!label) return false;

  return rule.values.some((value) => value === label.value);
};

const checkLabelExists = (
  db: Database,
  rule: LabelSelectorRequirement
): boolean => {
  if (rule.key === "environment") {
    return true;
  }

  return db.labels.some((label) => label.key === rule.key);
};
