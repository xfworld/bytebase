<template>
  <div ref="containerRef" class="h-full flex flex-col">
    <div v-if="isLoading" class="flex items-center justify-center h-full">
      <BBSpin />
    </div>
    <template v-else>
      <div class="border-b">
        <BannerSection v-if="!isCreating" />
        <FeatureAttention :feature="PlanFeature.FEATURE_APPROVAL_WORKFLOW" />

        <HeaderSection />
      </div>
      <div class="flex-1 flex flex-row">
        <div
          class="flex-1 flex flex-col hide-scrollbar divide-y overflow-x-hidden"
        >
          <StageSection />

          <TaskListSection />

          <TaskRunSection v-if="!isCreating" />

          <SQLCheckSection v-if="isCreating" />
          <PlanCheckSection v-if="!isCreating" />

          <StatementSection />

          <DescriptionSection />

          <IssueCommentSection v-if="!isCreating" />
        </div>

        <div
          v-if="sidebarMode == 'DESKTOP'"
          class="hide-scrollbar border-l"
          :style="{
            width: `${desktopSidebarWidth}px`,
          }"
        >
          <Sidebar />
        </div>
      </div>

      <template v-if="sidebarMode === 'MOBILE'">
        <!-- mobile sidebar -->
        <Drawer :show="mobileSidebarOpen" @close="mobileSidebarOpen = false">
          <div
            style="
              min-width: 240px;
              width: 80vw;
              max-width: 320px;
              padding: 0.5rem;
            "
          >
            <Sidebar v-if="sidebarMode === 'MOBILE'" />
          </div>
        </Drawer>
      </template>

      <IssueReviewActionPanel
        :action="ongoingIssueReviewAction?.action"
        @close="ongoingIssueReviewAction = undefined"
      />
      <IssueStatusActionPanel
        :action="ongoingIssueStatusAction?.action"
        @close="ongoingIssueStatusAction = undefined"
      />
      <TaskRolloutActionPanel
        :action="ongoingTaskRolloutAction?.action"
        :task-list="ongoingTaskRolloutAction?.taskList ?? []"
        @close="ongoingTaskRolloutAction = undefined"
      />
    </template>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from "vue";
import { useRoute, useRouter } from "vue-router";
import { BBSpin } from "@/bbkit";
import { FeatureAttention } from "@/components/FeatureGuard";
import { useIssueLayoutVersion } from "@/composables/useIssueLayoutVersion";
import {
  PROJECT_V1_ROUTE_ISSUE_DETAIL_V1,
  PROJECT_V1_ROUTE_PLAN_DETAIL_SPECS,
} from "@/router/dashboard/projectV1";
import { useCurrentProjectV1 } from "@/store";
import { Issue_Type } from "@/types/proto-es/v1/issue_service_pb";
import type { Plan, Plan_Spec } from "@/types/proto-es/v1/plan_service_pb";
import { type Task } from "@/types/proto-es/v1/rollout_service_pb";
import { PlanFeature } from "@/types/proto-es/v1/subscription_service_pb";
import { extractIssueUID, extractProjectResourceName } from "@/utils";
import { SQLCheckSection } from "../Plan/components";
import { providePlanSQLCheckContext } from "../Plan/components/SQLCheckSection";
import { provideSidebarContext } from "../Plan/logic";
import { Drawer } from "../v2";
import {
  BannerSection,
  HeaderSection,
  StageSection,
  TaskListSection,
  TaskRunSection,
  PlanCheckSection,
  StatementSection,
  DescriptionSection,
  Sidebar,
  IssueReviewActionPanel,
  IssueStatusActionPanel,
  TaskRolloutActionPanel,
  IssueCommentSection,
} from "./components";
import type {
  IssueReviewAction,
  IssueStatusAction,
  TaskRolloutAction,
} from "./logic";
import { specForTask, useIssueContext, usePollIssue } from "./logic";

const router = useRouter();
const route = useRoute();
const { enabledNewLayout } = useIssueLayoutVersion();
const { isCreating, issue, selectedTask, events } = useIssueContext();
const { project } = useCurrentProjectV1();
const containerRef = ref<HTMLElement>();
const isLoading = ref(true);

const ongoingIssueReviewAction = ref<{
  action: IssueReviewAction;
}>();
const ongoingIssueStatusAction = ref<{
  action: IssueStatusAction;
}>();
const ongoingTaskRolloutAction = ref<{
  action: TaskRolloutAction;
  taskList: Task[];
}>();

usePollIssue();

events.on("perform-issue-review-action", ({ action }) => {
  ongoingIssueReviewAction.value = {
    action,
  };
});

events.on("perform-issue-status-action", ({ action }) => {
  ongoingIssueStatusAction.value = {
    action,
  };
});

events.on("perform-task-rollout-action", async ({ action, tasks }) => {
  ongoingTaskRolloutAction.value = {
    action,
    taskList: tasks,
  };
});

providePlanSQLCheckContext({
  project,
  plan: computed(() => issue.value.planEntity as Plan),
  selectedSpec: computed(
    () =>
      specForTask(
        issue.value.planEntity as Plan,
        selectedTask.value
      ) as Plan_Spec
  ),
  selectedTask: selectedTask,
});

const {
  mode: sidebarMode,
  desktopSidebarWidth,
  mobileSidebarOpen,
} = provideSidebarContext(containerRef);

onMounted(() => {
  if (
    enabledNewLayout.value &&
    issue.value.type === Issue_Type.DATABASE_CHANGE &&
    issue.value.planEntity?.specs.every(
      (spec) => spec.config.case === "changeDatabaseConfig"
    )
  ) {
    if (isCreating.value) {
      // Redirect to plans creation page with original query parameters
      router.replace({
        name: PROJECT_V1_ROUTE_PLAN_DETAIL_SPECS,
        params: {
          projectId: extractProjectResourceName(project.value.name),
          planId: "create",
        },
        query: route.query,
      });
    } else {
      router.replace({
        name: PROJECT_V1_ROUTE_ISSUE_DETAIL_V1,
        params: {
          projectId: extractProjectResourceName(project.value.name),
          issueId: extractIssueUID(issue.value.name),
        },
        query: route.query,
      });
    }
  } else {
    isLoading.value = false;
  }
});
</script>
