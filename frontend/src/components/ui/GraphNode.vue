<template>
  <div
    class="relative min-w-40 rounded-lg bg-background text-primary shadow-md ring-2 ring-border transition-all dark:shadow-none dark:ring-border"
    :class="[{ '!ring-primary': isNodeSelected }, props.class]"
    @click.right.stop="console.log('noeud cliqué')"
  >
    <slot :isNodeSelected="isNodeSelected"></slot>
    <FrameData v-if="data.hasFrameDataSection" />
    <Transition
      enter-active-class="transition-all duration-200-"
      leave-active-class="transition-all duration-200"
      enter-from-class="opacity-0 translate-y-12 scale-75"
      leave-to-class="opacity-0 translate-y-12 scale-75"
    >
      <NodeToolbar
        :nodeId="props.id"
        @edit="emit('edit')"
        v-if="isNodeSelected && !isDragging"
      />
    </Transition>
  </div>
  <Handle :id="props.id + 'top'" type="source" :position="Position.Top" />
  <Handle :id="props.id + 'right'" type="source" :position="Position.Right" />
  <Handle :id="props.id + 'left'" type="source" :position="Position.Left" />
  <Handle :id="props.id + 'bot'" type="source" :position="Position.Bottom" />
</template>

<script setup lang="ts">
import { Handle, Position, useVueFlow } from '@vue-flow/core';
import { computed, ref } from 'vue';
import { CustomNodeData } from '@/types/CustomNodeData';
import NodeToolbar from '../flowchart/NodeToolbar.vue';
import FrameData from '../flowchart/FrameData.vue';

const props = defineProps<{
  id: string;
  data: CustomNodeData;
  isResizable?: boolean;
  class?: string;
}>();

const emit = defineEmits<{
  (e: 'edit'): void;
}>();

const isDragging = ref(false);
const { getSelectedNodes, onNodeDragStart, onNodeDragStop, selectNodesOnDrag } =
  useVueFlow();

const isNodeSelected = computed(() =>
  getSelectedNodes.value.some((n) => n.id === props.id),
);

onNodeDragStart((_) => {
  isDragging.value = true;
});

onNodeDragStop((_) => {
  isDragging.value = false;
});
</script>
