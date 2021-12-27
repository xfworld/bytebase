<template>
  <div>{{ name }}</div>
  <BBTable
    :column-list="columnList"
    :data-source="filteredMatrices"
    :show-header="true"
    :custom-header="true"
    :left-bordered="bordered"
    :right-bordered="bordered"
    :top-bordered="bordered"
    :bottom-bordered="bordered"
    :row-clickable="false"
  >
    <template #header>
      <BBTableHeaderCell
        :left-padding="4"
        :title="columnList[0].title + 'asdf'"
        class="w-1/12 pr-2"
      >
        <div class="relative flex items-center">
          <span>{{ firstByLabel }}</span>
          <heroicons-solid:selector class="h-4 w-4 text-control-light" />
          <select
            v-model="firstByLabel"
            class="absolute w-full h-full inset-0 opacity-0"
          >
            <option
              v-for="label in labelList"
              :key="label.key"
              :value="label.key"
            >
              {{ label.key }}
            </option>
          </select>
        </div>
      </BBTableHeaderCell>

      <BBTableHeaderCell
        v-for="xValue in filteredXAxisValues"
        :key="xValue"
        :title="xValue"
        :style="{
          width: `${(100 - 1 / 12) / filteredXAxisValues.length - 1}%`,
        }"
        class="text-center"
      >
        <template v-if="xValue">{{ xValue }}</template>
        <template v-else>{{ $t("database.empty-label-value") }}</template>
      </BBTableHeaderCell>
    </template>

    <template #body="{ rowData: matrix }">
      <BBTableCell
        :left-padding="4"
        class="pr-2 whitespace-nowrap"
        :class="{
          'text-control-placeholder': !matrix.labelValue,
        }"
      >
        <template v-if="matrix.labelValue">{{ matrix.labelValue }}</template>
        <template v-else>{{ $t("database.empty-label-value") }}</template>
      </BBTableCell>
      <BBTableCell v-for="(dbList, i) in matrix.databaseMatrix" :key="i">
        <div class="flex flex-col items-center space-y-1">
          <DatabaseMatrixItem
            v-for="db in dbList"
            :key="db.id"
            :database="db"
            :custom-click="customClick"
            @select-database="(db) => $emit('select-database', db)"
          />
          <span v-if="dbList.length === 0">-</span>
        </div>
      </BBTableCell>
    </template>
  </BBTable>
  <!-- <h1>{{ databaseListGroupByName.map((g) => g.name) }}</h1> -->
</template>

<script lang="ts">
import { computed, defineComponent, PropType, ref, watchEffect } from "vue";
import {
  Database,
  Environment,
  Label,
  LabelKeyType,
  LabelValueType,
} from "../../types";
import { BBTableColumn } from "../../bbkit/types";
import DatabaseMatrixItem from "./DatabaseMatrixItem.vue";
import { groupBy } from "lodash-es";
import { findDefaultGroupByLabel } from "../../utils";

type Mode = "ALL" | "ALL_SHORT" | "INSTANCE" | "PROJECT" | "PROJECT_SHORT";

type DatabaseMatrix = {
  labelValue: LabelValueType;
  databaseMatrix: Database[][];
};

export default defineComponent({
  name: "TenantDatabaseMatrix",
  components: { DatabaseMatrixItem },
  props: {
    bordered: {
      default: true,
      type: Boolean,
    },
    mode: {
      default: "ALL",
      type: String as PropType<Mode>,
    },
    customClick: {
      default: false,
      type: Boolean,
    },
    name: {
      type: String,
      required: true,
    },
    databaseList: {
      type: Object as PropType<Database[]>,
      required: true,
    },
    environmentList: {
      type: Array as PropType<Environment[]>,
      required: true,
    },
    labelList: {
      type: Array as PropType<Label[]>,
      required: true,
    },
  },
  emits: ["select-database"],
  setup(props) {
    /**
     * `firstByLabel` defines the y-axis of matrix
     * `thenByLabel` defines the x-axis of matrix
     * for now, `thenByLabel` is specified as 'bb.environment'
     */
    const firstByLabel = ref<LabelKeyType>();
    const thenByLabel = ref<LabelKeyType>("bb.environment");

    const filteredXAxisValues = ref<LabelValueType[]>([]);
    const filteredMatrices = ref<DatabaseMatrix[]>([]);

    // find the default label key to firstBy
    watchEffect(() => {
      // excluding "bb.environment" because the x-axis is already specified
      //   as "bb.environment"
      firstByLabel.value = findDefaultGroupByLabel(
        props.labelList,
        props.databaseList,
        [thenByLabel.value]
      );
    });

    // pre-filtered y-axis values
    const yAxisValues = computed((): string[] => {
      const key = firstByLabel.value;
      if (!key) {
        // y-axis is undefined
        return [];
      }

      // order based on label.valueList
      // plus one more "<empty value>"
      const label = props.labelList.find((label) => label.key === key);
      if (!label) return [];
      return [...label.valueList, ""];
    });

    // the matrix rows
    const firstGroupedBy = computed(() => {
      const key = firstByLabel.value;
      if (!key) {
        // y-axis is undefined
        return [];
      }
      const dict = groupBy(props.databaseList, (db) => getLabelValue(db, key));
      return yAxisValues.value.map((labelValue) => {
        const databaseList = dict[labelValue] || [];
        return {
          labelValue,
          databaseList,
        };
      });
      // .filter((pair) => pair.databaseList.length > 0);
    });

    // pre-filtered x-axis values
    const xAxisValues = computed(() => {
      // order based on label.valueList
      // plus one more "<empty value>"
      const key = thenByLabel.value;
      const label = props.labelList.find((label) => label.key === key);
      if (!label) return [];
      return [...label.valueList, ""];
    });

    // pre-filtered database matrices
    const matrices = computed(() => {
      const key = thenByLabel.value;
      return firstGroupedBy.value.map(
        ({ labelValue: yValue, databaseList }) => {
          const databaseMatrix: Database[][] = xAxisValues.value.map((xValue) =>
            databaseList.filter((db) => getLabelValue(db, key) === xValue)
          );
          return {
            labelValue: yValue,
            databaseMatrix,
          };
        }
      );
    });

    // now filter the axes and matrices
    // we only remove rows/cols of "<empty value>"
    // but keep other empty rows/cols because we want to give a whole view to
    //   the project.
    // e.g. if we hide "Prod" because there are no databases labeled as
    //   "Prod", we are not able to judge whether there is no environment
    //   named "Prod" or "Prod" has no databases.
    watchEffect(() => {
      filteredXAxisValues.value = [...xAxisValues.value];
      // if every row's "<empty value>" has no databases
      // we should hide the "<empty value>" col
      const shouldClipLastCol = matrices.value.every(
        (row) =>
          row.databaseMatrix.length > 0 &&
          row.databaseMatrix[row.databaseMatrix.length - 1].length === 0
      );
      if (shouldClipLastCol) filteredXAxisValues.value.pop();

      filteredMatrices.value = [];
      for (let i = 0; i < matrices.value.length; i++) {
        const row = {
          labelValue: matrices.value[i].labelValue,
          databaseMatrix: [...matrices.value[i].databaseMatrix],
        };
        if (shouldClipLastCol) {
          row.databaseMatrix.pop();
        }
        filteredMatrices.value.push(row);
      }
      // if every col of "<empty value>" row is empty
      // we should hide the "<empty value>" row
      if (
        filteredMatrices.value.length > 0 &&
        filteredMatrices.value[
          filteredMatrices.value.length - 1
        ].databaseMatrix.every((item) => item.length === 0)
      ) {
        filteredMatrices.value.pop();
      }
    });

    const columnList = computed((): BBTableColumn[] => [
      { title: firstByLabel.value || "" },
      ...filteredXAxisValues.value.map((xValue) => ({ title: xValue })),
    ]);

    return {
      firstByLabel,
      columnList,
      firstGroupedBy,
      yAxisValues,
      xAxisValues,
      matrices,
      filteredXAxisValues,
      filteredMatrices,
    };
  },
});

const getLabelValue = (db: Database, key: LabelKeyType): LabelValueType => {
  const label = db.labels.find((target) => target.key === key);
  if (!label) return "";
  return label.value;
};
</script>
