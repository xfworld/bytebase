<template>
  <div class="deployment-stage flex w-full relative">
    <div
      v-if="allowEdit"
      class="reorder flex flex-col items-center justify-between pr-2 py-0"
    >
      <button
        class="text-control hover:text-control-hover"
        :class="[index > 0 ? 'visible' : 'invisible']"
        @click="$emit('prev')"
      >
        <heroicons-solid:arrow-circle-up class="w-6 h-6" />
      </button>
      <button
        class="text-control hover:text-control-hover"
        :class="[index < max - 1 ? 'visible' : 'invisible']"
        @click="$emit('next')"
      >
        <heroicons-solid:arrow-circle-down class="w-6 h-6" />
      </button>
    </div>
    <div class="main flex-1 space-y-2 py-2 overflow-hidden">
      <slot name="header" />
      <div
        v-for="(selector, j) in deployment.spec.selector.matchExpressions"
        :key="j"
        class="flex content-start"
      >
        <SelectorItem
          :editable="allowEdit"
          :selector="selector"
          :label-list="labelList"
          @remove="removeSelector(selector)"
        />
      </div>
      <button v-if="allowEdit" class="btn-add btn-normal" @click="addSelector">
        {{ $t("deployment-config.add-selector") }}
      </button>
    </div>

    <span
      class="absolute right-2 top-2 p-1 text-control cursor-pointer hover:bg-gray-200 rounded-sm"
      @click="$emit('remove')"
    >
      <heroicons:solid:x class="w-4 h-4" />
    </span>
  </div>

  <div>
    <div v-for="db in filteredDatabaseList" :key="db.id">
      - name: {{ db.name }}, instance: {{ db.instance.name }}, env:
      {{ db.instance.environment.name }}, labels:
      {{ JSON.stringify(db.labels) }}
    </div>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, PropType } from "vue";
import {
  AvailableLabel,
  Database,
  Deployment,
  LabelSelectorRequirement,
} from "../../types";
import { filterDatabaseListByLabelSelector } from "../../utils";
import SelectorItem from "./SelectorItem.vue";

export default defineComponent({
  name: "DeploymentStage",
  components: { SelectorItem },
  props: {
    deployment: {
      type: Object as PropType<Deployment>,
      required: true,
    },
    index: {
      type: Number,
      required: true,
    },
    max: {
      type: Number,
      required: true,
    },
    allowEdit: {
      type: Boolean,
      default: false,
    },
    databaseList: {
      type: Array as PropType<Database[]>,
      default: () => [],
    },
    labelList: {
      type: Array as PropType<AvailableLabel[]>,
      default: () => [],
    },
  },
  emits: ["remove", "prev", "next"],
  setup(props) {
    const removeSelector = (selector: LabelSelectorRequirement) => {
      const array = props.deployment.spec.selector.matchExpressions;
      const index = array.indexOf(selector);
      if (index >= 0) {
        array.splice(index, 1);
      }
    };

    const addSelector = () => {
      const array = props.deployment.spec.selector.matchExpressions;
      array.push({
        key: props.labelList[0]?.key || "",
        operator: "In",
        values: [],
      });
    };
    const filteredDatabaseList = computed(() => {
      return filterDatabaseListByLabelSelector(
        props.databaseList,
        props.deployment.spec.selector
      );
    });

    return {
      removeSelector,
      addSelector,
      filteredDatabaseList,
    };
  },
});
</script>

<style scoped>
.btn-add {
  @apply py-1.5 !important;
}
</style>
