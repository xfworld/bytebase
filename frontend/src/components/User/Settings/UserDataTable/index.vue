<template>
  <NDataTable
    :columns="columns"
    :data="userList"
    :striped="true"
    :bordered="true"
    :max-height="'calc(100vh - 20rem)'"
    virtual-scroll
  />

  <BBAlert
    v-model:show="state.showResetKeyAlert"
    type="warning"
    :ok-text="$t('settings.members.reset-service-key')"
    :title="$t('settings.members.reset-service-key')"
    :description="$t('settings.members.reset-service-key-alert')"
    @ok="resetServiceKey"
    @cancel="state.showResetKeyAlert = false"
  />
</template>

<script lang="ts" setup>
import type { DataTableColumn } from "naive-ui";
import { NDataTable } from "naive-ui";
import { computed, reactive, h } from "vue";
import { useI18n } from "vue-i18n";
import { useUserStore } from "@/store";
import type { User } from "@/types/proto/v1/auth_service";
import type { UserGroup } from "@/types/proto/v1/user_group";
import { copyServiceKeyToClipboardIfNeeded } from "../common";
import UserGroupsCell from "./cells/UserGroupsCell.vue";
import UserNameCell from "./cells/UserNameCell.vue";
import UserOperationsCell from "./cells/UserOperationsCell.vue";
import UserRolesCell from "./cells/UserRolesCell.vue";

interface LocalState {
  showResetKeyAlert: boolean;
  targetServiceAccount?: User;
}

defineProps<{
  userList: User[];
}>();

const emit = defineEmits<{
  (event: "update-user", user: User): void;
  (event: "select-group", group: UserGroup): void;
}>();

const { t } = useI18n();
const userStore = useUserStore();
const state = reactive<LocalState>({
  showResetKeyAlert: false,
});

const columns = computed(() => {
  return [
    {
      key: "account",
      title: t("settings.members.table.account"),
      width: "32rem",
      resizable: true,
      render: (user: User) => {
        return h(UserNameCell, {
          user,
          "onReset-service-key": tryResetServiceKey,
        });
      },
    },
    {
      key: "roles",
      title: t("settings.members.table.role"),
      resizable: true,
      render: (user: User) => {
        return h(UserRolesCell, {
          roles: user.roles,
        });
      },
    },
    {
      key: "groups",
      title: t("settings.members.table.groups"),
      resizable: true,
      render: (user: User) => {
        return h(UserGroupsCell, {
          user,
          "onSelect-group": (group) => emit("select-group", group),
        });
      },
    },
    {
      key: "operations",
      title: "",
      width: "4rem",
      render: (user: User) => {
        return h(UserOperationsCell, {
          user,
          "onUpdate-user": () => {
            emit("update-user", user);
          },
        });
      },
    },
  ] as DataTableColumn<User>[];
});

const tryResetServiceKey = (user: User) => {
  state.showResetKeyAlert = true;
  state.targetServiceAccount = user;
};

const resetServiceKey = () => {
  state.showResetKeyAlert = false;
  const user = state.targetServiceAccount;

  if (!user) {
    return;
  }
  userStore
    .updateUser({
      user,
      updateMask: ["service_key"],
      regenerateRecoveryCodes: false,
      regenerateTempMfaSecret: false,
    })
    .then((updatedUser) => {
      copyServiceKeyToClipboardIfNeeded(updatedUser);
    });
};
</script>
