<template>
  <div class="max-w-3xl mx-auto divide-y">
    <DeploymentConfig
      v-if="schedule"
      :schedule="schedule"
      :allow-edit="allowEdit"
      :label-list="availableLabelList"
      :database-list="databaseList"
    />
    <div v-if="allowEdit" class="pt-4 flex justify-between items-center">
      <button class="btn-normal" @click="addStage">Add Stage</button>
      <button class="btn-primary">{{ $t("common.update") }}</button>
    </div>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, PropType, ref, watchEffect } from "vue";
import { useStore } from "vuex";
import {
  ProjectPatch,
  Project,
  Principal,
  Environment,
  Label,
  Database,
  AvailableLabel,
  DeploymentSchedule,
} from "../types";
import DeploymentConfig from "./DeploymentConfig";

export default defineComponent({
  name: "ProjectDeploymentConfigurationPanel",
  components: { DeploymentConfig },
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

    const prepareList = () => {
      store.dispatch("environment/fetchEnvironmentList");
      store.dispatch("label/fetchLabelList");
      store.dispatch("database/fetchDatabaseListByProjectId", props.project.id);
    };

    const environmentList = computed(
      () => store.getters["environment/environmentList"]() as Environment[]
    );

    const labelList = computed(
      () => store.getters["label/labelList"]() as Label[]
    );

    const databaseList = computed(
      () =>
        store.getters["database/databaseListByProjectId"](
          props.project.id
        ) as Database[]
    );

    watchEffect(prepareList);

    const availableLabelList = computed((): AvailableLabel[] => {
      const list: AvailableLabel[] = labelList.value.map((label) => {
        return { key: label.key, valueList: [...label.valueList] };
      });
      const environmentLabel: AvailableLabel = {
        key: "environment",
        valueList: environmentList.value.map((env) => env.name),
      };
      list.push(environmentLabel);

      return list;
    });

    const generateDefaultSchedule = () => {
      const schedule: DeploymentSchedule = {
        deployments: [],
      };
      environmentList.value.forEach((env) => {
        schedule.deployments.push({
          spec: {
            selector: {
              matchExpressions: [
                {
                  key: "environment",
                  operator: "In",
                  values: [env.name],
                },
              ],
            },
          },
        });
      });
      return schedule;
    };

    const schedule = ref<DeploymentSchedule>();

    watchEffect(() => {
      if (environmentList.value.length > 0) {
        schedule.value = generateDefaultSchedule();
      }
    });

    const addStage = () => {
      schedule.value.deployments.push({
        spec: {
          selector: {
            matchExpressions: [],
          },
        },
      });
    };

    return {
      environmentList,
      labelList,
      databaseList,
      availableLabelList,
      schedule,
      addStage,
    };
  },
});
</script>
