<template>
  <div v-if="available">
    <NButton text type="primary" :size="size" @click="onClick">
      {{ $t("sql-editor.request-query") }}
    </NButton>

    <RequestQueryPanel
      :show="showPanel"
      :project-id="database?.projectEntity.uid"
      :database="database"
      :placement="panelPlacement"
      @close="showPanel = false"
    />
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from "vue";
import RequestQueryPanel from "@/components/Issue/panel/RequestQueryPanel/index.vue";
import { useCurrentUserV1 } from "@/store";
import { UNKNOWN_ID, type ComposedDatabase } from "@/types";
import { hasPermissionToCreateRequestGrantIssue } from "@/utils";

const props = withDefaults(
  defineProps<{
    database: ComposedDatabase;
    panelPlacement: "left" | "right";
    size?: "tiny" | "medium";
  }>(),
  {
    panelPlacement: "right",
    size: "medium",
  }
);

const me = useCurrentUserV1();
const showPanel = ref(false);

const available = computed(() => {
  if (props.database.uid === String(UNKNOWN_ID)) {
    return false;
  }

  return hasPermissionToCreateRequestGrantIssue(props.database, me.value);
});

const onClick = (e: MouseEvent) => {
  e.stopPropagation();
  e.preventDefault();
  showPanel.value = true;
};
</script>
