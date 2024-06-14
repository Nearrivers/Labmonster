<script setup lang="ts">
import { ref } from 'vue';
import type { Node } from '@vue-flow/core';
import { VueFlow, Panel } from '@vue-flow/core';
import { Background } from '@vue-flow/background';

const nodes = ref<Node[]>([
  {
    id: '1',
    position: { x: 25, y: 90 },
    data: { label: 'Test' },
  },
]);

function addNode() {
  const id = Date.now().toString();

  nodes.value.push({
    id,
    position: { x: 150, y: 50 },
    data: { label: `Node ${id}` },
  });
}
</script>

<template>
  <VueFlow :nodes="nodes" class="h-48">
    <Background />
    <Panel :position="'top-left'">
      <button type="button" @click="addNode">Add a node</button>
    </Panel>
  </VueFlow>
</template>
