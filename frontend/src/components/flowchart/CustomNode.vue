<template>
  <div
    class="relative min-w-40 rounded-lg bg-background text-primary shadow-md ring-2 ring-border transition-all dark:shadow-none dark:ring-border"
    :class="{ '!ring-primary': isNodeSelected }"
  >
    <input
      ref="input"
      :id="props.id"
      class="bg-transparent px-2 py-4 outline-none"
      :class="{ 'cursor-grab': !isNodeSelected }"
      v-model="nodeText"
      @input="handleUpdate"
      @keypress.enter="input?.blur()"
      autocomplete="off"
    />
    <FrameData v-if="data.hasFrameDataSection" />
    <Transition
      enter-active-class="transition-all duration-200-"
      leave-active-class="transition-all duration-200"
      enter-from-class="opacity-0 translate-y-12 scale-75"
      leave-to-class="opacity-0 translate-y-12 scale-75"
    >
      <NodeToolbar
        :nodeId="props.id"
        @edit="onEdit"
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
import NodeToolbar from './NodeToolbar.vue';
import FrameData from './FrameData.vue';
import { CustomNodeData } from '@/types/CustomNodeData';

const props = defineProps<{
  id: string;
  data: CustomNodeData;
}>();
const input = ref<HTMLInputElement | null>(null);
const isDragging = ref(false);
const { updateNode, getSelectedNodes, onNodeDragStart, onNodeDragStop } =
  useVueFlow();
const nodeText = ref('');

const isNodeSelected = computed(() =>
  getSelectedNodes.value.some((n) => n.id === props.id),
);

onNodeDragStart((_) => {
  isDragging.value = true;
});

onNodeDragStop((_) => {
  isDragging.value = false;
});

function handleUpdate() {
  updateNode<Partial<CustomNodeData>>(props.id, {
    data: {
      title: nodeText.value,
      hasFrameDataSection: props.data.hasFrameDataSection,
    },
  });
}

function onEdit() {
  if (input.value) {
    input.value.focus();
  }
}
</script>
