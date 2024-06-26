<template>
  <div ref="tableRef">
    <NDataTable
      :columns="columnList"
      :data="planList"
      :striped="true"
      :bordered="true"
      :loading="loading"
      :row-key="(plan: ComposedPlan) => plan.uid"
      :row-props="rowProps"
      class="plan-data-table"
    />
  </div>
</template>

<script lang="tsx" setup>
import { useElementSize } from "@vueuse/core";
import type { DataTableColumn } from "naive-ui";
import { NPerformantEllipsis, NDataTable } from "naive-ui";
import { computed, ref } from "vue";
import { useI18n } from "vue-i18n";
import { useRouter } from "vue-router";
import { BBAvatar } from "@/bbkit";
import { ProjectNameCell } from "@/components/v2/Model/DatabaseV1Table/cells";
import { PROJECT_V1_ROUTE_PLAN_DETAIL } from "@/router/dashboard/projectV1";
import type { ComposedPlan } from "@/types/v1/issue/plan";
import { extractProjectResourceName, humanizeTs, planSlug } from "@/utils";

const { t } = useI18n();

const columnList = computed((): DataTableColumn<ComposedPlan>[] => {
  const columns: (DataTableColumn<ComposedPlan> & { hide?: boolean })[] = [
    // TODO(steven): show latest plan check run status.
    {
      key: "title",
      title: t("issue.table.name"),
      resizable: true,
      render: (plan) => {
        return (
          <div class="flex items-center overflow-hidden space-x-2">
            <div class="whitespace-nowrap text-control opacity-60">
              {plan.projectEntity.key}-{plan.uid}
            </div>
            <NPerformantEllipsis class="flex-1 truncate">
              {{
                default: () => <span>{plan.title}</span>,
                tooltip: () => (
                  <div class="whitespace-pre-wrap break-words break-all">
                    {plan.title}
                  </div>
                ),
              }}
            </NPerformantEllipsis>
          </div>
        );
      },
    },
    {
      key: "project",
      title: t("common.project"),
      width: 144,
      resizable: true,
      hide: !props.showProject,
      render: (plan) => (
        <ProjectNameCell project={plan.projectEntity} mode={"ALL_SHORT"} />
      ),
    },
    {
      key: "creator",
      resizable: true,
      width: 150,
      title: t("issue.table.creator"),
      hide: !showExtendedColumns.value,
      render: (plan) => (
        <div class="flex flex-row items-center overflow-hidden gap-x-2">
          <BBAvatar size="SMALL" username={plan.creatorEntity.title} />
          <span class="truncate">{plan.creatorEntity.title}</span>
        </div>
      ),
    },
    {
      key: "updateTime",
      title: t("issue.table.updated"),
      width: 150,
      resizable: true,
      hide: !showExtendedColumns.value,
      render: (plan) => humanizeTs((plan.updateTime?.getTime() ?? 0) / 1000),
    },
  ];
  return columns.filter((column) => !column.hide);
});

const props = withDefaults(
  defineProps<{
    planList: ComposedPlan[];
    loading?: boolean;
    showProject: boolean;
  }>(),
  {
    loading: true,
    showProject: true,
  }
);

const router = useRouter();

const tableRef = ref<HTMLDivElement>();
const { width: tableWidth } = useElementSize(tableRef);
const showExtendedColumns = computed(() => {
  return tableWidth.value > 800;
});

const rowProps = (plan: ComposedPlan) => {
  return {
    style: "cursor: pointer;",
    onClick: (e: MouseEvent) => {
      const route = router.resolve({
        name: PROJECT_V1_ROUTE_PLAN_DETAIL,
        params: {
          projectId: extractProjectResourceName(plan.project),
          planSlug: planSlug(plan.title, plan.uid),
        },
      });
      const url = route.fullPath;
      if (e.ctrlKey || e.metaKey) {
        window.open(url, "_blank");
      } else {
        router.push(url);
      }
    },
  };
};
</script>
