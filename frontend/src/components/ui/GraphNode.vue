<template>
  <NodeResizer
    :minWidth="100"
    :minHeight="60"
    :handleClassName="'opacity-0 z-20'"
    :lineClassName="'z-20 !border-transparent rounded-lg !border'"
    @resizeEnd="rememberDimensions"
  />
  <div
    class="relative h-full rounded-lg bg-background text-popover ring-2 ring-accent transition-all"
    :class="[{ '!ring-popover': isNodeSelected }, props.class]"
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
  <Handle
    :id="props.id + 'top'"
    type="source"
    :position="Position.Top"
    :class="'z-30 h-4 w-4 !cursor-pointer bg-accent-foreground opacity-0 hover:opacity-100'"
  />
  <Handle
    :id="props.id + 'right'"
    type="source"
    :position="Position.Right"
    :class="'z-30 h-4 w-4 !cursor-pointer bg-accent-foreground opacity-0 hover:opacity-100'"
  />
  <Handle
    :id="props.id + 'left'"
    type="source"
    :position="Position.Left"
    :class="'z-30 h-4 w-4 !cursor-pointer bg-accent-foreground opacity-0 hover:opacity-100'"
  />
  <Handle
    :id="props.id + 'bot'"
    type="source"
    :position="Position.Bottom"
    :class="'z-30 h-4 w-4 !cursor-pointer bg-accent-foreground opacity-0 hover:opacity-100'"
  />
</template>

<script setup lang="ts">
import { Handle, Position, useNode, useVueFlow } from '@vue-flow/core';
import { computed, ref } from 'vue';
import { CustomNodeData } from '@/types/CustomNodeData';
import NodeToolbar from '../flowchart/NodeToolbar.vue';
import FrameData from '../flowchart/FrameData.vue';
import { NodeResizer, OnResizeStart } from '@vue-flow/node-resizer';
import '@vue-flow/node-resizer/dist/style.css';

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
