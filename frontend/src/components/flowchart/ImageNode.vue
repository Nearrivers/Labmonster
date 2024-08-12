<template>
  <GraphNode :id="id" :data="data" isResizable class="p-2">
    <label :for="props.id">
      <input
        ref="input"
        type="file"
        :id="props.id"
        class="hidden bg-transparent px-2 py-4 outline-none"
        @input="handleUpdate"
        @change="handleUpdate"
        autocomplete="off"
      />
    </label>
    <div class="nodrag resize overflow-auto">
      <img
        :src="imgSrc"
        alt="Image du setup"
        class="object-contain"
        :style="{ width, height }"
      />
    </div>
  </GraphNode>
</template>

<script setup lang="ts">
import { useVueFlow } from '@vue-flow/core';
import { onMounted, ref, watch } from 'vue';
import { CustomNodeData } from '@/types/CustomNodeData';
import GraphNode from '../ui/GraphNode.vue';

const props = defineProps<{
  id: string;
  width: number;
  height: number;
  data: CustomNodeData;
}>();

const imgSrc = ref('');
const input = ref<HTMLInputElement | null>(null);
const { updateNode } = useVueFlow();

watch(imgSrc, handleUpdate);

onMounted(() => {
  if (!props.data.image) {
    return;
  }

  if (typeof props.data.image === 'string') {
    console.log('test');
    imgSrc.value = props.data.image;
    return;
  }

  const reader = new FileReader();
  reader.onload = function (e) {
    imgSrc.value = e.target?.result as string;
  };
  reader.readAsDataURL(props.data.image);
});

function handleUpdate() {
  updateNode<Partial<CustomNodeData>>(props.id, {
    data: {
      text: '',
      image: imgSrc.value,
      hasFrameDataSection: props.data.hasFrameDataSection,
    },
  });
}
</script>
