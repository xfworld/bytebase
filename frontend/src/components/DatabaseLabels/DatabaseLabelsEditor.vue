<template>
  <div class="database-labels-editor flex items-center">
    <DatabaseLabels :labels="state.labels" :editable="state.mode === 'EDIT'" />
    <div
      v-if="state.mode === 'VIEW' && state.labels.length === 0"
      class="text-sm text-control-placeholder"
    >
      {{ $t("database.no-label") }}
    </div>
    <div class="buttons flex items-center gap-1 ml-1 text-control">
      <template v-if="state.mode === 'VIEW'">
        <div class="icon-btn lite" @click="state.mode = 'EDIT'">
          <heroicons-outline:pencil class="w-4 h-4" />
        </div>
      </template>
      <template v-else>
        <div class="icon-btn text-error" @click="cancel">
          <heroicons-solid:x class="w-4 h-4" />
        </div>

        <div class="icon-btn text-success" @click="save">
          <heroicons-solid:check class="w-4 h-4" />
        </div>
      </template>
    </div>
  </div>
</template>

<script lang="ts">
import { cloneDeep } from "lodash-es";
import { defineComponent, PropType, reactive, watch } from "vue";
import { DatabaseLabel } from "../../types";

type LocalState = {
  mode: "VIEW" | "EDIT";
  labels: DatabaseLabel[];
};

export default defineComponent({
  name: "DatabaseLabelsEditor",
  props: {
    labels: {
      type: Array as PropType<DatabaseLabel[]>,
      default: () => [],
    },
  },
  emits: ["save"],
  setup(props, { emit }) {
    const state = reactive<LocalState>({
      mode: "VIEW",
      labels: cloneDeep(props.labels),
    });

    watch(
      () => props.labels,
      (labels) => {
        state.labels = cloneDeep(props.labels);
      }
    );

    const cancel = () => {
      state.labels = cloneDeep(props.labels);
      state.mode = "VIEW";
    };
    const save = () => {
      emit("save", state.labels);
      state.mode = "VIEW";
    };

    return {
      state,
      cancel,
      save,
    };
  },
});
</script>

<style scoped>
.icon-btn {
  @apply px-2 py-1 inline-flex items-center
    rounded bg-white border border-control-border
    hover:border-control-hover
    cursor-pointer;
}
.icon-btn.lite {
  @apply px-1 border-none hover:bg-control-bg-hover;
}
</style>
