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
      <BBTableCell :left-padding="4" class="pr-2">
        {{ matrix.labelValue }}
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
        </div>
      </BBTableCell>
    </template>
  </BBTable>
  <!-- <h1>{{ databaseListGroupByName.map((g) => g.name) }}</h1> -->
</template>

<script lang="ts">
import {
  computed,
  defineComponent,
  PropType,
  reactive,
  ref,
  watchEffect,
} from "vue";
import { useStore } from "vuex";
import { useRouter } from "vue-router";
import { consoleLink, databaseSlug } from "../../utils";
import {
  Database,
  Environment,
  Label,
  LabelKey,
  LabelValue,
} from "../../types";
import { BBTableColumn } from "../../bbkit/types";
import InstanceEngineIcon from "../InstanceEngineIcon.vue";
import DatabaseMatrixItem from "./DatabaseMatrixItem.vue";
import { cloneDeep, isEmpty, groupBy } from "lodash-es";
import { useI18n } from "vue-i18n";

type Mode = "ALL" | "ALL_SHORT" | "INSTANCE" | "PROJECT" | "PROJECT_SHORT";

type DatabaseListGroupByLabel = {
  labelValue: LabelValue;
  databaseList: Database[];
};

type DatabaseMatrixGroupByLabelAndEnvironment = {
  labelValue: LabelValue;
  databaseMatrix: Database[][];
};

export default defineComponent({
  name: "TenantDatabaseMatrix",
  components: { InstanceEngineIcon, DatabaseMatrixItem },
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
  setup(props, { emit }) {
    const store = useStore();
    const router = useRouter();
    const { t } = useI18n();

    const groupByLabel = ref<LabelKey>();
    watchEffect(() => {
      groupByLabel.value = props.labelList[0]?.key;
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
            if (!label) return t("database.group-other");
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

    // const columnListMap = computed(() => {
    //   return new Map([
    //     [
    //       "ALL",
    //       [
    //         {
    //           title: t("common.name"),
    //         },
    //         {
    //           title: t("common.project"),
    //         },
    //         {
    //           title: t("common.environment"),
    //         },
    //         {
    //           title: t("common.instance"),
    //         },
    //         {
    //           title: t("db.sync-status"),
    //         },
    //         {
    //           title: t("db.last-successful-sync"),
    //         },
    //       ],
    //     ],
    //     [
    //       "ALL_SHORT",
    //       [
    //         {
    //           title: t("common.name"),
    //         },
    //         {
    //           title: t("common.project"),
    //         },
    //         {
    //           title: t("common.environment"),
    //         },
    //         {
    //           title: t("common.instance"),
    //         },
    //       ],
    //     ],
    //     [
    //       "INSTANCE",
    //       [
    //         {
    //           title: t("common.name"),
    //         },
    //         {
    //           title: t("common.project"),
    //         },
    //         {
    //           title: t("db.sync-status"),
    //         },
    //         {
    //           title: t("db.last-successful-sync"),
    //         },
    //       ],
    //     ],
    //     [
    //       "PROJECT",
    //       [
    //         {
    //           title: t("common.name"),
    //         },
    //         {
    //           title: t("common.environment"),
    //         },
    //         {
    //           title: t("common.instance"),
    //         },
    //         {
    //           title: t("db.sync-status"),
    //         },
    //         {
    //           title: t("db.last-successful-sync"),
    //         },
    //       ],
    //     ],
    //     [
    //       "PROJECT_SHORT",
    //       [
    //         {
    //           title: t("common.name"),
    //         },
    //         {
    //           title: t("common.environment"),
    //         },
    //         {
    //           title: t("common.instance"),
    //         },
    //       ],
    //     ],
    //   ]);
    // });

    // // const currentUser = computed(() => store.getters["auth/currentUser"]());

    // const showInstanceColumn = computed(() => {
    //   return props.mode != "INSTANCE";
    // });

    // const showProjectColumn = computed(() => {
    //   return props.mode != "PROJECT" && props.mode != "PROJECT_SHORT";
    // });

    // const showEnvironmentColumn = computed(() => {
    //   return props.mode != "INSTANCE";
    // });

    // const showMiscColumn = computed(() => {
    //   return props.mode != "ALL_SHORT" && props.mode != "PROJECT_SHORT";
    // });

    // const columnList = computed(() => {
    //   var list: BBTableColumn[] = columnListMap.value.get(props.mode)!;
    //   if (showConsoleLink.value) {
    //     // Use cloneDeep, otherwise it will alter the one in columnListMap
    //     list = cloneDeep(list);
    //     list.push({ title: t("database.sql-console") });
    //   }
    //   return list;
    // });

    // const showConsoleLink = computed(() => {
    //   if (props.mode == "ALL_SHORT" || props.mode == "PROJECT_SHORT") {
    //     return false;
    //   }

    //   const consoleUrl =
    //     store.getters["setting/settingByName"]("bb.console.url").value;
    //   return !isEmpty(consoleUrl);
    // });

    // const databaseConsoleLink = (databaseName: string) => {
    //   const consoleUrl =
    //     store.getters["setting/settingByName"]("bb.console.url").value;
    //   if (!isEmpty(consoleUrl)) {
    //     return consoleLink(consoleUrl, databaseName);
    //   }
    //   return "";
    // };

    // const clickDatabase = function (section: number, row: number) {
    //   const database = props.databaseList[row];
    //   if (props.customClick) {
    //     emit("select-database", database);
    //   } else {
    //     router.push(`/db/${databaseSlug(database)}`);
    //   }
    // };

    return {
      groupByLabel,
      databaseListGroupByLabel,
      databaseMatrices,
      columnList,
      // showInstanceColumn,
      // showProjectColumn,
      // showEnvironmentColumn,
      // showMiscColumn,
      // columnList,
      // showConsoleLink,
      // databaseConsoleLink,
      // clickDatabase,
    };
  },
});
</script>
