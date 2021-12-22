<template>
  <div class="bb-database-matrix-item" @click="clickDatabase">
    <div class="sync-status whitespace-nowrap">
      <span class="tooltip-wrapper">
        <heroicons-solid:check
          v-if="database.syncStatus === 'OK'"
          class="w-4 h-4 text-success"
        />
        <heroicons-outline:exclamation v-else class="w-4 h-4 text-warning" />
        <span class="tooltip">
          {{ $t("db.last-successful-sync") }}
          {{ humanizeTs(database.lastSuccessfulSyncTs) }}
        </span>
      </span>
      <span class="flex-1">{{ database.syncStatus }}</span>
    </div>

    <div class="instance whitespace-pre-wrap">
      <InstanceEngineIcon :instance="database.instance" />
      <span class="flex-1">{{ instanceName(database.instance) }}</span>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, PropType } from "vue";
import { useRouter } from "vue-router";
import { Database } from "../../types";
import { databaseSlug } from "../../utils";
import InstanceEngineIcon from "../InstanceEngineIcon.vue";

export default defineComponent({
  name: "DatabaseMatrixItem",
  components: {
    InstanceEngineIcon,
  },
  props: {
    database: {
      type: Object as PropType<Database>,
      required: true,
    },
    customClick: {
      default: false,
      type: Boolean,
    },
  },
  emits: ["select-database"],
  setup(props, { emit }) {
    const router = useRouter();

    const clickDatabase = () => {
      const { customClick, database } = props;
      if (customClick) {
        emit("select-database", database);
      } else {
        router.push(`/db/${databaseSlug(database)}`);
      }
    };

    return { clickDatabase };
  },
});
</script>

<style scoped>
.bb-database-matrix-item {
  @apply border-gray-300 border rounded px-2 py-0 divide-y cursor-pointer select-none hover:bg-gray-50;
}
.bb-database-matrix-item > * {
  @apply flex items-center py-1 gap-1;
}
</style>
