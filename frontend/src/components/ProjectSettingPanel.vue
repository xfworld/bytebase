<template>
  <div class="max-w-3xl mx-auto space-y-4">
    <h1>deployment config</h1>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, PropType } from "vue";
import { useStore } from "vuex";
import { isProjectOwner } from "../utils";
import { ProjectPatch, Project } from "../types";

export default defineComponent({
  name: "ProjectSettingPanel",
  props: {
    project: {
      required: true,
      type: Object as PropType<Project>,
    },
    allowEdit: {
      default: true,
      type: Boolean,
    },
  },
  setup(props) {
    const store = useStore();

    return {};
  },
});
</script>
