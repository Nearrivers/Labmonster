<template>
  <Transition
    enter-active-class="transition-all duration-200-"
    leave-active-class="transition-all duration-200"
    enter-from-class="opacity-0 translate-y-12 scale-75"
    leave-to-class="opacity-0 translate-y-12 scale-75"
  >
    <div
      class="absolute -top-20 left-1/2 flex -translate-x-1/2 scale-100 rounded-md border bg-background p-1"
      v-if="isEdgeSelected"
    >
      <ToolbarButton @click="removeEdges(edge.id)">
        <template #icon>
          <Trash2 class="h-5 w-5" />
        </template>
        <template #tooltip> Supprimer </template>
      </ToolbarButton>
      <ToolbarButton
        @click="
          fitView({
            nodes: [edge.source, edge.target],
            duration: 200,
          })
        "
      >
        <template #icon><ScanSearch class="h-5 w-5" /> </template>
        <template #tooltip> Zoomer sur le noeud </template>
      </ToolbarButton>
      <ToolbarButton @click="emit('edit')">
        <template #icon><SquarePen class="h-5 w-5" /> </template>
        <template #tooltip> Modifier le label </template>
      </ToolbarButton>
      <ToolbarButton
        @click="emit('remove')"
        v-if="edge.data.label && edge.data.label.length > 0"
      >
        <template #icon><PenOff class="h-5 w-5" /> </template>
        <template #tooltip> Enlever le label </template>
      </ToolbarButton>
    </div>
  </Transition>
</template>

<script setup lang="ts">
import { useEdge, useVueFlow } from '@vue-flow/core';
import ToolbarButton from './ToolbarButton.vue';
import { PenOff, ScanSearch, SquarePen, Trash2 } from 'lucide-vue-next';
import { computed } from 'vue';

const emit = defineEmits<{
  (e: 'edit'): void;
  (e: 'remove'): void;
}>();
const { edge } = useEdge();
const { removeEdges, fitView, getSelectedEdges } = useVueFlow();
const isEdgeSelected = computed(() =>
  getSelectedEdges.value.some((e) => e.id === edge.id),
);
</script>
