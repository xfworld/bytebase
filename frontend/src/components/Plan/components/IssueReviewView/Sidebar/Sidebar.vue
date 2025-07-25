<template>
  <div class="w-full flex flex-col p-4 gap-4">
    <IssueStatusSection
      v-if="issue.approvalTemplates.length > 0"
      :issue="issue"
    />

    <ApprovalFlowSection :issue="issue" />

    <IssueLabels
      :project="project"
      :value="issue.labels || []"
      :disabled="!allowChange"
      @update:value="onIssueLabelsUpdate"
    />

    <div class="w-full flex flex-col">
      <h3 class="textlabel mb-1">
        {{ $t("common.description") }}
      </h3>
      <NInput
        v-model:value="issue.description"
        type="textarea"
        :placeholder="$t('issue.add-some-description')"
        :disabled="!allowChange"
        :autosize="false"
        :resizable="false"
        @blur="onIssueDescriptionUpdate(issue.description)"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { create } from "@bufbuild/protobuf";
import { NInput } from "naive-ui";
import { computed } from "vue";
import { useI18n } from "vue-i18n";
import IssueLabels from "@/components/IssueV1/components/Sidebar/IssueLabels.vue";
import { useResourcePoller } from "@/components/Plan/logic/poller";
import { issueServiceClientConnect } from "@/grpcweb";
import {
  extractUserId,
  pushNotification,
  useCurrentProjectV1,
  useCurrentUserV1,
} from "@/store";
import {
  UpdateIssueRequestSchema,
  IssueStatus,
} from "@/types/proto-es/v1/issue_service_pb";
import { hasProjectPermissionV2 } from "@/utils";
import { usePlanContextWithIssue } from "../../../logic/context";
import ApprovalFlowSection from "./ApprovalFlowSection/ApprovalFlowSection.vue";
import IssueStatusSection from "./IssueStatusSection.vue";

const { t } = useI18n();
const { issue } = usePlanContextWithIssue();
const currentUser = useCurrentUserV1();
const { project } = useCurrentProjectV1();
const { refreshResources } = useResourcePoller();

const allowChange = computed(() => {
  if (issue.value.status !== IssueStatus.OPEN) {
    return false;
  }

  // Allowed if current user is the creator.
  if (extractUserId(issue.value.creator) === currentUser.value.email) {
    return true;
  }

  return hasProjectPermissionV2(project.value, "bb.issues.update");
});

const onIssueLabelsUpdate = async (labels: string[]) => {
  const issuePatch = {
    ...issue.value,
    labels,
  };
  const request = create(UpdateIssueRequestSchema, {
    issue: issuePatch,
    updateMask: { paths: ["labels"] },
  });
  await issueServiceClientConnect.updateIssue(request);
  refreshResources(["issue"], true /** force */);
  pushNotification({
    module: "bytebase",
    style: "SUCCESS",
    title: t("common.updated"),
  });
};

const onIssueDescriptionUpdate = async (description: string) => {
  const issuePatch = {
    ...issue.value,
    description,
  };
  const request = create(UpdateIssueRequestSchema, {
    issue: issuePatch,
    updateMask: { paths: ["description"] },
  });
  await issueServiceClientConnect.updateIssue(request);
  // After running checks, we need to refresh the plan and plan check runs.
  refreshResources(["issue"], true /** force */);
  pushNotification({
    module: "bytebase",
    style: "SUCCESS",
    title: t("common.updated"),
  });
};
</script>
