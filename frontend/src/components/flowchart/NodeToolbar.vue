<template>
  <div
    class="absolute -top-20 left-1/2 flex -translate-x-1/2 rounded-md border bg-background p-1"
  >
    <ToolbarButton @click="removeNodes(nodeId)">
      <template #icon>
        <Trash2 class="h-5 w-5" />
      </template>
      <template #tooltip> Supprimer </template>
    </ToolbarButton>
    <ToolbarButton @click="fitView({ nodes: [props.nodeId], duration: 200 })">
      <template #icon><ScanSearch class="h-5 w-5" /> </template>
      <template #tooltip> Zoomer sur le noeud </template>
    </ToolbarButton>
    <ToolbarButton @click="emit('edit')">
      <template #icon><SquarePen class="h-5 w-5" /> </template>
      <template #tooltip> Modifier </template>
    </ToolbarButton>
  </div>
</template>

<script setup lang="ts">
import { ScanSearch, SquarePen, Trash2 } from 'lucide-vue-next';
import ToolbarButton from './ToolbarButton.vue';
import { useVueFlow } from '@vue-flow/core';

const props = defineProps<{
  nodeId: string;
}>();

const emit = defineEmits<{
  (e: 'edit'): void;
}>();

const { removeNodes, fitView } = useVueFlow();
</script>
