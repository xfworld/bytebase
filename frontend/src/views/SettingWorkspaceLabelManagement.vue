<template>
  <div class="mt-2 space-y-6 divide-y divide-block-border">
    <div>
      <h3 class="text-lg leading-6 font-medium text-main">
        {{ $t("settings.workspace.label-management") }}
      </h3>
      <div class="mt-4">
        <BBTable
          class="mt-2"
          :column-list="COLUMN_LIST"
          :data-source="labelList"
          :show-header="true"
          :row-clickable="false"
        >
          <template #header>
            <BBTableHeaderCell
              class="w-36 table-cell"
              :left-padding="4"
              :title="COLUMN_LIST[0].title"
            />
            <BBTableHeaderCell
              class="w-auto table-cell"
              :title="COLUMN_LIST[1].title"
            />
          </template>
          <template #body="{ rowData: label }">
            <BBTableCell :left-padding="4" class="w-36 table-cell">
              {{ label.key }}
            </BBTableCell>
            <BBTableCell class="whitespace-nowrap">
              <div class="tags">
                <div v-for="(value, j) in label.valueList" :key="j" class="tag">
                  <span>{{ value }}</span>
                  <span
                    v-if="allowEdit"
                    class="remove"
                    @click="removeValue(label, value)"
                  >
                    <heroicons-solid:x class="w-3 h-3" />
                  </span>
                </div>
                <AddLabelValue
                  v-if="allowEdit"
                  :label="label"
                  @add="(v) => addValue(label, v)"
                />
              </div>
            </BBTableCell>
          </template>
        </BBTable>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { computed, watchEffect } from "vue";
import { useStore } from "vuex";
import { isDBAOrOwner } from "../utils";
import { Principal, Label, LabelPatch } from "../types";
import { BBTableColumn } from "../bbkit/types";
import { BBTable, BBTableCell, BBTableHeaderCell } from "../bbkit";
import { useI18n } from "vue-i18n";
import AddLabelValue from "../components/AddLabelValue.vue";

export default {
  name: "SettingWorkspaceLabelManagement",
  components: {
    BBTable,
    BBTableCell,
    BBTableHeaderCell,
    AddLabelValue,
  },
  setup() {
    const { t } = useI18n();
    const store = useStore();
    const currentUser = computed(
      () => store.getters["auth/currentUser"]() as Principal
    );

    const prepareLabelList = () => {
      store.dispatch("label/fetchLabelList");
    };

    watchEffect(prepareLabelList);

    const labelList = computed(
      () => store.getters["label/labelList"]() as Label[]
    );

    const allowEdit = computed(() => {
      return isDBAOrOwner(currentUser.value.role);
    });

    const addValue = (label: Label, value: string) => {
      const labelPatch: LabelPatch = {
        valueList: [...label.valueList, value],
      };
      store.dispatch("label/patchLabel", {
        id: label.id,
        labelPatch,
      });
    };

    const removeValue = (label: Label, value: string) => {
      const valueList = [...label.valueList];
      const index = valueList.indexOf(value);
      if (index < 0) return;
      valueList.splice(index, 1);
      const labelPatch: LabelPatch = {
        valueList,
      };
      store.dispatch("label/patchLabel", {
        id: label.id,
        labelPatch,
      });
    };

    const COLUMN_LIST = computed((): BBTableColumn[] => [
      {
        title: t("settings.label-management.table.key"),
      },
      {
        title: t("settings.label-management.table.values"),
      },
    ]);

    return {
      COLUMN_LIST,
      labelList,
      allowEdit,
      addValue,
      removeValue,
    };
  },
};
</script>

<style scoped>
.tags {
  @apply flex flex-wrap gap-2;
}
.tag {
  @apply h-6 bg-gray-200 px-2 rounded whitespace-nowrap inline-flex items-center;
}
.tag > .remove {
  @apply ml-1 -mr-0.5 p-px cursor-pointer hover:bg-gray-300 rounded-sm;
}
</style>
