<template>
  <div class="max-w-3xl mx-auto">
    <BBAttention
      v-if="deployment?.id === EMPTY_ID"
      :style="'WARN'"
      :title="$t('common.deployment-config')"
      :description="$t('deployment-config.this-is-example-deployment-config')"
    >
    </BBAttention>
    <div class="divide-y">
      <DeploymentConfigTool
        v-if="deployment"
        :schedule="deployment.schedule"
        :allow-edit="allowEdit"
        :label-list="availableLabelList"
        :database-list="databaseList"
      />
      <div v-if="allowEdit" class="pt-4 flex justify-between items-center">
        <button class="btn-normal" @click="addStage">Add Stage</button>
        <button class="btn-primary" @click="update">
          {{ $t("common.update") }}
        </button>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, PropType, ref, watchEffect } from "vue";
import { useStore } from "vuex";
import {
  Project,
  Environment,
  Label,
  Database,
  AvailableLabel,
  DeploymentConfig,
  UNKNOWN_ID,
  EMPTY_ID,
  empty,
  DeploymentConfigPatch,
} from "../types";
import DeploymentConfigTool from "./DeploymentConfigTool";
import { cloneDeep } from "lodash-es";
import { useI18n } from "vue-i18n";
import { generateDefaultSchedule } from "../utils";

export default defineComponent({
  name: "ProjectDeploymentConfigurationPanel",
  components: { DeploymentConfigTool },
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
    const { t } = useI18n();
    const deployment = ref<DeploymentConfig>();

    const prepareList = () => {
      store.dispatch("environment/fetchEnvironmentList");
      store.dispatch("label/fetchLabelList");
      store.dispatch("database/fetchDatabaseListByProjectId", props.project.id);
      store.dispatch(
        "deployment/fetchDeploymentConfigByProjectId",
        props.project.id
      );
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
      return labelList.value.map((label) => {
        return { key: label.key, valueList: [...label.valueList] };
      });
    });

    watchEffect(() => {
      const dep = store.getters["deployment/deploymentConfigByProjectId"](
        props.project.id
      ) as DeploymentConfig;
      if (dep.id === UNKNOWN_ID) {
        // if the project has no related deployment-config
        // just generate a "staged-by-env" example to users
        // this is not saved immediately, it's a draft
        // users need to edit and save it before creating a deployment issue
        if (environmentList.value.length > 0) {
          deployment.value = empty("DEPLOYMENT_CONFIG") as DeploymentConfig;
          deployment.value.schedule = generateDefaultSchedule(
            environmentList.value
          );
        }
      } else {
        // otherwise we clone the saved deployment-config
        // <DeploymentConfigTool /> will mutate `deployment.value` directly
        // when update button clicked, we save the draft to backend
        // we don't show a "cancel" button because if users don't want to save
        //   the draft, they can just leave the page without any saving action
        // even more we may deliver a confirm modal when leaving the page with a
        //   dirty but not saved draft
        deployment.value = cloneDeep(dep);
      }
    });

    const addStage = () => {
      if (!deployment.value) return;

      deployment.value.schedule.deployments.push({
        spec: {
          selector: {
            matchExpressions: [],
          },
        },
      });
    };

    const update = () => {
      if (!deployment.value) return;

      const deploymentConfigPatch: DeploymentConfigPatch = {
        payload: JSON.stringify(deployment.value.schedule),
      };
      store.dispatch("deployment/patchDeploymentConfigByProjectId", {
        projectId: props.project.id,
        deploymentConfigPatch,
      });
      store.dispatch("notification/pushNotification", {
        module: "bytebase",
        style: "SUCCESS",
        title: t("deployment-config.update-success"),
      });
    };

    return {
      EMPTY_ID,
      environmentList,
      labelList,
      databaseList,
      availableLabelList,
      deployment,
      addStage,
      update,
    };
  },
});
</script>
