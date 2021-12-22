<template>
  <DatabaseMatrix
    v-for="dbGroup in databaseListGroupByName"
    :key="dbGroup.name"
    :name="dbGroup.name"
    :database-list="dbGroup.databaseList"
    :environment-list="environmentList"
    :label-list="labelList"
  />
</template>

<script lang="ts">
import { computed, defineComponent, PropType, watchEffect } from "vue";
import { useStore } from "vuex";
import { useRouter } from "vue-router";
import { Database, Environment, Label } from "../../types";
import InstanceEngineIcon from "../InstanceEngineIcon.vue";
import { groupBy } from "lodash-es";
import { useI18n } from "vue-i18n";
import DatabaseMatrix from "./DatabaseMatrix.vue";

type Mode = "ALL" | "ALL_SHORT" | "INSTANCE" | "PROJECT" | "PROJECT_SHORT";

type DatabaseGroupByName = {
  name: string;
  databaseList: Database[];
};

export default defineComponent({
  name: "TenantDatabaseTable",
  components: { InstanceEngineIcon, DatabaseMatrix },
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
    databaseList: {
      type: Object as PropType<Database[]>,
      required: true,
    },
  },
  emits: ["select-database"],
  setup(props, { emit }) {
    const store = useStore();
    const router = useRouter();
    const { t } = useI18n();

    const prepareList = () => {
      store.dispatch("environment/fetchEnvironmentList");
      store.dispatch("label/fetchLabelList");
    };
    watchEffect(prepareList);

    const labelList = computed(
      () => store.getters["label/labelList"]() as Label[]
    );
    const environmentList = computed(
      () => store.getters["environment/environmentList"]() as Environment[]
    );

    const databaseListGroupByName = computed((): DatabaseGroupByName[] => {
      const dict = groupBy(props.databaseList, "name");
      return Object.keys(dict).map((name) => ({
        name,
        databaseList: dict[name],
      }));
    });

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
      labelList,
      environmentList,
      databaseListGroupByName,

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
