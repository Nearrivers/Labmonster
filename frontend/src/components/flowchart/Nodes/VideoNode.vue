<template>
  <GraphNode :id="id" :data="data">
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
    <div class="nodrag resize overflow-auto" :style="dimensions">
      <video
        controls
        :src="imgSrc"
        alt="Vidéo du setup"
        class="h-full w-full object-contain"
      ></video>
    </div>
  </GraphNode>
</template>

<script setup lang="ts">
import { Styles, useNode, useVueFlow } from '@vue-flow/core';
import { computed, onMounted, onUnmounted, ref } from 'vue';
import { CustomNodeData } from '@/types/CustomNodeData';
import { OpenMedia } from '$/file_handler/FileHandler';
import { useShowErrorToast } from '@/composables/useShowErrorToast';
import GraphNode from '@/components/ui/GraphNode.vue';

const props = defineProps<{
  id: string;
  data: CustomNodeData;
}>();

const imgSrc = ref('');
const { node } = useNode(props.id);
const { updateNode } = useVueFlow();
const { showToast } = useShowErrorToast();
const input = ref<HTMLInputElement | null>(null);
const dimensions = computed(() =>
  node.style
    ? {
        width: (node.style as Styles).width + 'px',
        height: (node.style as Styles).height + 'px',
      }
    : {},
);

onMounted(async () => {
  if (!props.data.image) {
    return;
  }

  try {
    imgSrc.value = await OpenMedia(props.data.image);
  } catch (error) {
    showToast(error);
  }
});

onUnmounted(() => {
  URL.revokeObjectURL(imgSrc.value);
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
