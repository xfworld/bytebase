<template>
  <div class="add-label-value">
    <div v-if="!isAdding" class="btn" @click="isAdding = true">
      <heroicons-solid:plus class="w-4 h-4" />
    </div>
    <template v-else>
      <input
        ref="input"
        v-model="text"
        type="text"
        autocomplete="off"
        class="textfield"
        :class="{ error }"
        :placeholder="$t('settings.label-management.value-placeholder')"
        @blur="cancel"
        @keyup.esc="cancel"
        @keyup.enter="tryAdd"
      />
      <div class="btn cancel" @click="cancel">
        <heroicons-solid:x class="w-4 h-4" />
      </div>
      <div class="btn save" @mousedown.prevent.stop="tryAdd">
        <heroicons-solid:check class="w-4 h-4" />
      </div>
    </template>
  </div>
</template>

<script lang="ts">
import {
  defineComponent,
  PropType,
  ref,
  watchEffect,
  nextTick,
  watch,
} from "vue";
import { Label } from "../types";

export default defineComponent({
  name: "AddLabelValue",
  props: {
    label: {
      type: Object as PropType<Label>,
      required: true,
    },
  },
  emits: ["add"],
  setup(props, { emit }) {
    const input = ref<HTMLInputElement>();
    const isAdding = ref(false);
    const text = ref("");
    const error = ref(false);

    const check = () => {
      const v = text.value.trim();
      if (!v) {
        // can't be empty
        error.value = true;
      } else if (props.label.valueList.includes(v)) {
        // must be unique
        error.value = true;
      } else {
        // ok
        error.value = false;
      }
      return !error.value;
    };

    watch(text, check);

    const cancel = () => {
      isAdding.value = false;
      text.value = "";
    };

    const tryAdd = () => {
      if (!check()) {
        return;
      }
      const v = text.value.trim();
      emit("add", v);
      cancel();
    };

    watchEffect(() => {
      if (isAdding.value) {
        // clear error status
        error.value = false;
        // auto focus if possible
        nextTick(() => input.value?.focus());
      }
    });

    return { isAdding, text, error, tryAdd, cancel, input };
  },
});
</script>

<style>
.add-label-value {
  @apply inline-flex flex-nowrap items-center gap-1 h-6;
}
.add-label-value > * {
  @apply h-full;
}

.btn {
  @apply px-1 inline-flex items-center
    rounded bg-white border border-control-border
    hover:border-control-hover
    cursor-pointer;
}
.textfield {
  @apply rounded px-2 py-0 text-sm w-32;
}
.textfield.error {
  @apply border-error focus:ring-error focus:border-error;
}
.cancel {
  @apply text-error;
}
.save {
  @apply text-success;
}
</style>
