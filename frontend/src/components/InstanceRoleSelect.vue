<template>
  <NSelect
    v-bind="$attrs"
    :value="role"
    :options="options"
    :placeholder="$t('instance.select-database-user')"
    :filter="filterByTitle"
    :filterable="true"
    :virtual-scroll="true"
    :fallback-option="false"
    @update:value="$emit('update:instance-role', $event)"
  />
</template>

<script lang="ts" setup>
import type { SelectOption } from "naive-ui";
import { computed, ref, watch } from "vue";
import { useInstanceV1Store } from "@/store";
import type { InstanceRole } from "@/types/proto/v1/instance_role_service";

interface InstanceRoleSelectOption extends SelectOption {
  value: string;
  instanceRole: InstanceRole;
}

const props = defineProps<{
  role?: string;
  instance?: string;
  filter: (role: InstanceRole) => boolean;
}>();

const emit = defineEmits<{
  (event: "update:instance-role", role: string | undefined): void;
}>();

const instanceV1Store = useInstanceV1Store();
const instanceRoleList = ref<InstanceRole[]>([]);

watch(
  () => props.instance,
  async (instanceName) => {
    if (instanceName) {
      const instance = instanceV1Store.getInstanceByName(instanceName);
      instanceRoleList.value =
        await instanceV1Store.fetchInstanceRoleListByName(instance.name);
      emit("update:instance-role", undefined);
    }
  },
  { immediate: true }
);

const options = computed(() => {
  return filteredInstanceRoleList.value.map<InstanceRoleSelectOption>(
    (instanceRole) => {
      return {
        instanceRole,
        value: instanceRole.name,
        label: instanceRole.roleName,
      };
    }
  );
});

const filteredInstanceRoleList = computed(() => {
  if (!props.filter) return instanceRoleList.value;
  return instanceRoleList.value.filter(props.filter);
});

const filterByTitle = (pattern: string, option: SelectOption) => {
  const { instanceRole } = option as InstanceRoleSelectOption;
  return instanceRole.roleName.toLowerCase().includes(pattern.toLowerCase());
};
</script>
