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
    <div class="h-full overflow-hidden p-2">
      <img
        :src="imgSrc"
        alt="Image du setup"
        class="h-full w-full object-contain"
      />
    </div>
  </GraphNode>
</template>

<script setup lang="ts">
import { useVueFlow } from '@vue-flow/core';
import { onMounted, ref } from 'vue';
import { CustomNodeData } from '@/types/CustomNodeData';
import { OpenMedia } from '$/filetree/FileTree';
import { useShowErrorToast } from '@/composables/useShowErrorToast';
import GraphNode from '@/components/ui/GraphNode.vue';

const props = defineProps<{
  id: string;
  data: CustomNodeData;
}>();

const imgSrc = ref('');
const { updateNode } = useVueFlow();
const { showToast } = useShowErrorToast();
const input = ref<HTMLInputElement | null>(null);

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
