<template>
  <div>{{ name }}</div>
  <BBTable
    :column-list="columnList"
    :data-source="databaseMatrices"
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
          <span>{{ groupByLabel }}</span>
          <heroicons-solid:selector class="h-4 w-4 text-control-light" />
          <select
            v-model="groupByLabel"
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
        v-for="env in environmentList"
        :key="env.id"
        :title="env.name"
        :style="{ width: `${(100 - 1 / 12) / environmentList.length - 1}%` }"
        class="text-center"
      />
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

type DatabaseListGroupByLabel = {
  labelValue: LabelValueType;
  databaseList: Database[];
};

type DatabaseMatrixGroupByLabelAndEnvironment = {
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
    const groupByLabel = ref<LabelKeyType>();
    watchEffect(() => {
      // find the default label key to groupBy
      groupByLabel.value = findDefaultGroupByLabel(
        props.labelList,
        props.databaseList
      );
    });

    const databaseListGroupByLabel = computed(
      (): DatabaseListGroupByLabel[] => {
        if (!groupByLabel.value) {
          // no field for grouping
          // this is nonsense
          return [];
        } else {
          const dict = groupBy(props.databaseList, (db) => {
            const label = db.labels.find(
              (target) => target.key === groupByLabel.value
            );
            if (!label) return "";
            return label.value;
          });
          return Object.keys(dict).map((value) => ({
            labelValue: value,
            databaseList: dict[value],
          }));
        }
      }
    );

    const databaseMatrices = computed(
      (): DatabaseMatrixGroupByLabelAndEnvironment[] => {
        return databaseListGroupByLabel.value.map(
          ({ labelValue, databaseList }) => {
            const databaseDictGroupByEnvironment = groupBy(
              databaseList,
              (db) => db.instance.environment.id
            );
            const databaseMatrix: Database[][] = [];
            for (let i = 0; i < props.environmentList.length; i++) {
              const env = props.environmentList[i];
              databaseMatrix[i] = databaseDictGroupByEnvironment[env.id] || [];
            }
            return {
              labelValue,
              databaseMatrix,
            };
          }
        );
      }
    );

    const columnList = computed((): BBTableColumn[] => [
      { title: groupByLabel.value || "" },
      ...props.environmentList.map((env) => ({ title: env.name })),
    ]);

    return {
      groupByLabel,
      databaseListGroupByLabel,
      databaseMatrices,
      columnList,
    };
  },
});
</script>
