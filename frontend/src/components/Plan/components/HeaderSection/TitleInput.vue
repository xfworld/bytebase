<template>
  <div class="flex-1">
    <NInput
      v-model:value="state.title"
      :style="style"
      :loading="state.isUpdating"
      :disabled="!allowEdit || state.isUpdating"
      size="medium"
      required
      class="bb-plan-title-input"
      @focus="state.isEditing = true"
      @blur="onBlur"
      @keyup.enter="onEnter"
      @update:value="onUpdateValue"
    />
  </div>
</template>

<script setup lang="ts">
import { create } from "@bufbuild/protobuf";
import { NInput } from "naive-ui";
import type { CSSProperties } from "vue";
import { computed, reactive, watch } from "vue";
import { useI18n } from "vue-i18n";
import { planServiceClientConnect } from "@/grpcweb";
import {
  pushNotification,
  useCurrentUserV1,
  extractUserId,
  useCurrentProjectV1,
} from "@/store";
import { UpdatePlanRequestSchema } from "@/types/proto-es/v1/plan_service_pb";
import { PlanSchema } from "@/types/proto-es/v1/plan_service_pb";
import { hasProjectPermissionV2 } from "@/utils";
import { usePlanContext } from "../../logic";

type ViewMode = "EDIT" | "VIEW";

const { t } = useI18n();
const currentUser = useCurrentUserV1();
const { project } = useCurrentProjectV1();
const { isCreating, plan, readonly } = usePlanContext();

const state = reactive({
  isEditing: false,
  isUpdating: false,
  title: plan.value.title,
});

// Watch for changes in issue/plan to update the title
watch(
  () => [plan.value],
  () => {
    state.title = plan.value.title;
  },
  { immediate: true }
);

const viewMode = computed((): ViewMode => {
  if (isCreating.value) return "EDIT";
  return state.isEditing ? "EDIT" : "VIEW";
});

const style = computed(() => {
  const style: CSSProperties = {
    cursor: "default",
    "--n-color-disabled": "transparent",
    "--n-font-size": "18px",
    "font-weight": "bold",
  };
  const border =
    viewMode.value === "EDIT"
      ? "1px solid rgb(var(--color-control-border))"
      : "none";
  style["--n-border"] = border;
  style["--n-border-disabled"] = border;

  return style;
});

const allowEdit = computed(() => {
  if (readonly.value) {
    return false;
  }
  if (isCreating.value) {
    return true;
  }
  // Allowed if current user is the creator.
  if (extractUserId(plan.value.creator) === currentUser.value.email) {
    return true;
  }
  // Allowed if current user has related permission.
  if (hasProjectPermissionV2(project.value, "bb.plans.update")) {
    return true;
  }
  return false;
});

const onBlur = async () => {
  const cleanup = () => {
    state.isEditing = false;
    state.isUpdating = false;
  };

  if (isCreating.value) {
    cleanup();
    return;
  }

  // Update plan title
  if (state.title === plan.value.title) {
    cleanup();
    return;
  }
  try {
    state.isUpdating = true;
    const planPatch = create(PlanSchema, {
      ...plan.value,
      title: state.title,
    });
    const request = create(UpdatePlanRequestSchema, {
      plan: planPatch,
      updateMask: { paths: ["title"] },
    });
    const response = await planServiceClientConnect.updatePlan(request);
    Object.assign(plan.value, response);
    pushNotification({
      module: "bytebase",
      style: "SUCCESS",
      title: t("common.updated"),
    });
  } finally {
    cleanup();
  }
};

const onEnter = (e: Event) => {
  const input = e.target as HTMLInputElement;
  input.blur();
};

const onUpdateValue = (value: string) => {
  if (!isCreating.value) {
    return;
  }
  // When creating, we only update plan title
  plan.value.title = value;
};
</script>

<style>
.bb-plan-title-input input {
  cursor: text !important;
  color: var(--n-text-color) !important;
  text-decoration-color: var(--n-text-color) !important;
}
</style>
