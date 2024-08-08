<template>
  <Panel :position="'bottom-right'" class="flex">
    <div
      class="select-none rounded-md border border-border bg-background p-1 text-sm"
    >
      {{ fileName }}
    </div>
  </Panel>
</template>

<script setup lang="ts">
import { Panel } from '@vue-flow/core';
import { computed } from 'vue';
import { useRoute } from 'vue-router';

const route = useRoute();

const fileName = computed(() =>
  route.params.path.slice(0, route.params.path.indexOf('.')),
);
</script>
