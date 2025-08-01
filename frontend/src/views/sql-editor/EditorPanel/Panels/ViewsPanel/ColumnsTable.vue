<template>
  <div ref="containerElRef" class="w-full h-full px-2 py-2 overflow-x-auto">
    <NDataTable
      v-bind="$attrs"
      ref="dataTableRef"
      size="small"
      :row-key="(column) => column.name"
      :columns="columns"
      :data="layoutReady ? filteredColumns : []"
      :max-height="tableBodyHeight"
      :virtual-scroll="true"
      :striped="true"
      :bordered="true"
      :bottom-bordered="true"
      row-class-name="cursor-default"
    />
  </div>
</template>

<script lang="ts" setup>
import type { DataTableColumn } from "naive-ui";
import { NCheckbox, NDataTable } from "naive-ui";
import { computed, h, watch } from "vue";
import { useI18n } from "vue-i18n";
import { DefaultValueCell } from "@/components/SchemaEditorLite/Panels/TableColumnEditor/components";
import type { ComposedDatabase } from "@/types";
import type {
  ColumnMetadata,
  DatabaseMetadata,
  SchemaMetadata,
  ViewMetadata,
} from "@/types/proto-es/v1/database_service_pb";
import { getHighlightHTMLByRegExp, useAutoHeightDataTable } from "@/utils";
import { EllipsisCell } from "../../common";
import { useCurrentTabViewStateContext } from "../../context/viewState";

const props = defineProps<{
  db: ComposedDatabase;
  database: DatabaseMetadata;
  schema: SchemaMetadata;
  view: ViewMetadata;
  keyword?: string;
}>();

const { viewState } = useCurrentTabViewStateContext();
const {
  dataTableRef,
  containerElRef,
  virtualListRef,
  tableBodyHeight,
  layoutReady,
} = useAutoHeightDataTable();
const { t } = useI18n();

const filteredColumns = computed(() => {
  const keyword = props.keyword?.trim().toLowerCase();
  if (keyword) {
    return props.view.columns.filter((column) =>
      column.name.toLowerCase().includes(keyword)
    );
  }
  return props.view.columns;
});

const columns = computed(() => {
  const engine = props.db.instanceResource.engine;
  const downGrade = filteredColumns.value.length > 50;
  const columns: (DataTableColumn<ColumnMetadata> & { hide?: boolean })[] = [
    {
      key: "name",
      title: t("schema-editor.column.name"),
      resizable: true,
      minWidth: 140,
      className: "truncate",
      render: (column) => {
        return h("span", {
          innerHTML: getHighlightHTMLByRegExp(column.name, props.keyword ?? ""),
        });
      },
    },
    {
      key: "type",
      title: t("schema-editor.column.type"),
      resizable: true,
      minWidth: 140,
      maxWidth: 320,
      className: "truncate",
    },
    {
      key: "default-value",
      title: t("schema-editor.column.default"),
      resizable: true,
      minWidth: 140,
      maxWidth: 320,
      className: "input-cell",
      render: (column) => {
        return h(DefaultValueCell, {
          column,
          disabled: true,
          engine: engine,
        });
      },
    },
    {
      key: "comment",
      title: t("schema-editor.column.comment"),
      resizable: true,
      minWidth: 140,
      maxWidth: 320,
      className: "overflow-hidden",
      render: (column) => {
        return h(EllipsisCell, {
          content: column.userComment,
          downGrade,
        });
      },
    },
    {
      key: "not-null",
      title: t("schema-editor.column.not-null"),
      resizable: true,
      minWidth: 80,
      maxWidth: 160,
      className: "checkbox-cell",
      render: (column) => {
        return h(NCheckbox, {
          checked: !column.nullable,
          readonly: true,
        });
      },
    },
  ];
  return columns.filter((header) => !header.hide);
});

watch(
  [() => viewState.value?.detail.column, virtualListRef],
  ([column, vl]) => {
    if (column && vl) {
      requestAnimationFrame(() => {
        vl.scrollTo({ key: column });
      });
    }
  },
  { immediate: true }
);
</script>

<style lang="postcss" scoped>
:deep(.n-data-table-th .n-data-table-resize-button::after) {
  @apply bg-control-bg h-2/3;
}
:deep(.n-data-table-td.input-cell) {
  @apply pl-0.5 pr-1 py-0;
}

:deep(.n-data-table-td.input-cell .n-input__placeholder),
:deep(.n-data-table-td.input-cell .n-base-selection-placeholder) {
  @apply italic;
}
:deep(.n-data-table-td.checkbox-cell) {
  @apply pr-1 py-0;
}
:deep(.n-data-table-td.text-cell) {
  @apply pr-1 py-0;
}
</style>
