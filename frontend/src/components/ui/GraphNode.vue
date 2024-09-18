<template>
  <NodeResizer
    :minWidth="100"
    :minHeight="60"
    :handleClassName="'opacity-0 z-20'"
    :lineClassName="'z-20 !border-transparent rounded-lg !border'"
    @resizeEnd="rememberDimensions"
  />
  <div
    class="relative h-full rounded-lg bg-background p-2 text-popover ring-2 ring-accent transition-all"
    :class="[{ '!ring-popover': isNodeSelected }, props.class]"
  >
    <TitlePart />
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
  <NodeHandles :nodeId="id" />
</template>

<script setup lang="ts">
import '@vue-flow/node-resizer/dist/style.css';
import { useNode, useVueFlow } from '@vue-flow/core';
import { computed, ref } from 'vue';
import { CustomNodeData } from '@/types/CustomNodeData';
import NodeToolbar from '../flowchart/NodeToolbar.vue';
import FrameData from '../flowchart/FrameData.vue';
import { NodeResizer, OnResizeStart } from '@vue-flow/node-resizer';
import TitlePart from '../flowchart/Nodes/Parts/TitlePart.vue';
import NodeHandles from './NodeHandles.vue';

const props = defineProps<{
  id: string;
  data: CustomNodeData;
  class?: string;
}>();

const emit = defineEmits<{
  (e: 'edit'): void;
}>();

const isDragging = ref(false);
const { node } = useNode(props.id);
const { getSelectedNodes, onNodeDragStart, onNodeDragStop } = useVueFlow();

const isNodeSelected = computed(() =>
  getSelectedNodes.value.some((n) => n.id === props.id),
);

function rememberDimensions(e: OnResizeStart) {
  node.style = {
    width: e.params.width,
    height: e.params.height,
  };
}

onNodeDragStart((_) => {
  isDragging.value = true;
});

onNodeDragStop((_) => {
  isDragging.value = false;
});
</script>
