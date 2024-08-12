<template>
  <Panel :position="'bottom-right'" class="flex">
    <div
      class="flex select-none items-center gap-2 rounded-md border border-border bg-background p-1 text-sm"
    >
      <slot></slot>
      <RefreshCw class="h-4 w-4 animate-spin" v-if="isSaving" />
      <CircleCheck class="h-4 w-4" v-else />
    </div>
  </Panel>
</template>

<script setup lang="ts">
import { Panel } from '@vue-flow/core';
import { CircleCheck, RefreshCw } from 'lucide-vue-next';

defineProps<{
  isSaving: boolean;
}>();
</script>
