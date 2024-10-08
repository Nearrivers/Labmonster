<template>
  <GraphNode :id="id" :data="data" v-slot="isNodeSelected">
    <textarea
      ref="input"
      :id="props.id"
      class="resize-none bg-transparent px-2 py-4 outline-none"
      rows="1"
      :class="{ 'cursor-grab': !isNodeSelected }"
      v-model="nodeText"
      @input="handleUpdate"
      @keypress.enter="input?.blur()"
      autocomplete="off"
      :style="dimensions"
    />
  </GraphNode>
</template>

<script setup lang="ts">
import { Styles, useNode, useVueFlow } from '@vue-flow/core';
import { computed, ref } from 'vue';
import { CustomNodeData } from '@/types/CustomNodeData';
import GraphNode from '@/components/ui/GraphNode.vue';

const props = defineProps<{
  id: string;
  data: CustomNodeData;
}>();

const { node } = useNode(props.id);
const nodeText = defineModel<string>('text');
const input = ref<HTMLTextAreaElement | null>(null);
const { updateNode } = useVueFlow();

const dimensions = computed(() =>
  node.style
    ? {
        width: (node.style as Styles).width + 'px',
        height: (node.style as Styles).height + 'px',
      }
    : {},
);

function handleUpdate() {
  updateNode<Partial<CustomNodeData>>(props.id, {
    data: {
      ...props.data,
      text: nodeText.value,
    },
  });
}
</script>
