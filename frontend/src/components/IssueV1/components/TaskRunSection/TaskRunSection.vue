<template>
  <div v-if="flattenTaskRunList.length > 0" class="px-4 py-2 space-y-4">
    <TaskRunTable :task-run-list="flattenTaskRunList" />
  </div>
</template>

<script lang="ts" setup>
import { create } from "@bufbuild/protobuf";
import { computed, watchEffect } from "vue";
import {
  useIssueContext,
  taskRunListForTask,
} from "@/components/IssueV1/logic";
import { rolloutServiceClientConnect } from "@/grpcweb";
import { GetTaskRunLogRequestSchema } from "@/types/proto-es/v1/rollout_service_pb";
import { TaskRun_Status } from "@/types/proto-es/v1/rollout_service_pb";
import TaskRunTable from "./TaskRunTable.vue";

const { issue, selectedTask } = useIssueContext();

const flattenTaskRunList = computed(() => {
  return taskRunListForTask(issue.value, selectedTask.value);
});

watchEffect(async () => {
  // Fetching the latest task run log for running task runs of selected task.
  for (const taskRun of flattenTaskRunList.value) {
    if (taskRun.status === TaskRun_Status.RUNNING) {
      const request = create(GetTaskRunLogRequestSchema, {
        parent: taskRun.name,
      });
      const response = await rolloutServiceClientConnect.getTaskRunLog(request);
      taskRun.taskRunLog = response;
    }
  }
});
</script>
