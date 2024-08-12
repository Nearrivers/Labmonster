<template>
  <GraphNode :id="id" :data="data" v-slot="isNodeSelected">
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
  </GraphNode>
</template>

<script setup lang="ts">
import { useVueFlow } from '@vue-flow/core';
import { ref } from 'vue';
import { CustomNodeData } from '@/types/CustomNodeData';
import GraphNode from '../ui/GraphNode.vue';

const props = defineProps<{
  id: string;
  data: CustomNodeData;
}>();

const nodeText = defineModel<string>('text');
const input = ref<HTMLInputElement | null>(null);
const { updateNode } = useVueFlow();

function handleUpdate() {
  updateNode<Partial<CustomNodeData>>(props.id, {
    data: {
      ...props.data,
      text: nodeText.value,
    },
  });
}

function onEdit() {
  if (input.value) {
    input.value.focus();
  }
}
</script>
