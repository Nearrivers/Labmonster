<template>
  <VueFlow :nodes="nodes" class="h-[calc(100%-33px)]">
    <!-- :default-viewport="{ zoom: 0.5 }" -->
    <TopMenu @add-node="addNode" @zoom-in="zoomIn" @zoom-out="zoomOut" />
    <Background :gap="30" />
  </VueFlow>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import type { Node } from '@vue-flow/core';
import { VueFlow } from '@vue-flow/core';
import { Background } from '@vue-flow/background';
import { useTopMenuActions } from '@/composables/Flowchart/useTopMenuActions';
import TopMenu from './flowchart/TopMenu.vue';

const nodes = ref<Node[]>([
  {
    id: '1',
    position: { x: 25, y: 90 },
    data: { label: 'Test' },
    class:
      'dark:bg-secondary border border-border dark:border-zinc-700 text-primary rounded-lg shadow-md',
  },
]);

const { addNode, zoomIn, zoomOut } = useTopMenuActions(nodes);
</script>
