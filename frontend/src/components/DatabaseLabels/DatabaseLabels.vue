<template>
  <div class="flex gap-2">
    <DatabaseLabel
      v-for="(label, i) in labels"
      :key="i"
      :label="label"
      :editable="editable"
      :available-labels="availableLabels"
      @remove="removeLabel(i)"
    />
    <div v-if="editable && allowAdd" class="add-button" @click="addLabel">
      <heroicons-solid:plus class="w-4 h-4" />
    </div>
  </div>
</template>

<script lang="ts">
/* eslint-disable vue/no-mutating-props */

import { computed, defineComponent, PropType, watchEffect } from "vue";
import { useStore } from "vuex";
import { DatabaseLabel, Label } from "../../types";

const MAX_DATABASE_LABELS = 4;

export default defineComponent({
  name: "DatabaseLabels",
  props: {
    labels: {
      type: Array as PropType<DatabaseLabel[]>,
      default: () => [],
    },
    editable: {
      type: Boolean,
      default: false,
    },
  },
  setup(props, { emit }) {
    const store = useStore();

    const allowAdd = computed(() => props.labels.length < MAX_DATABASE_LABELS);

    const prepareLabelList = () => {
      store.dispatch("label/fetchLabelList");
    };
    watchEffect(prepareLabelList);

    const availableLabels = computed(
      () => store.getters["label/labelList"]() as Label[]
    );

    const addLabel = () => {
      const key = availableLabels.value[0]?.key || "";
      const value = availableLabels.value[0]?.valueList[0] || "";
      props.labels.push({
        key,
        value,
      });
    };

    const removeLabel = (index: number) => {
      props.labels.splice(index, 1);
    };

    return {
      availableLabels,
      allowAdd,
      addLabel,
      removeLabel,
    };
  },
});
</script>

<style scoped>
.add-button {
  @apply px-2 py-1 inline-flex items-center
    rounded bg-white border border-control-border
    hover:border-control-hover
    cursor-pointer;
}
.textfield {
  @apply rounded px-2 py-0 text-sm w-32;
}
.textfield.error {
  @apply border-error focus:ring-error focus:border-error;
}
.cancel {
  @apply text-error;
}
.save {
  @apply text-success;
}
</style>
