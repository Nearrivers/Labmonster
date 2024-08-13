<template>
  <VueFlow
    :nodes="nodes"
    class="h-[calc(100%-33px)]"
    auto-connect
    :edges="edges"
    :default-edge-options="{
      type: 'custom',
      markerEnd: {
        type: MarkerType.ArrowClosed,
        width: 10,
        height: 10,
        color: '#52525b',
      },
    }"
    :delete-key-code="'Delete'"
    :zoom-on-scroll="false"
    :zoom-on-double-click="false"
    :pan-on-drag="false"
    :pan-on-scroll="true"
    :pan-activation-key-code="'Space'"
    :zoom-activation-key-code="['Control', 'Space']"
    :select-nodes-on-drag="true"
  >
    <FlowchartButtons
      @add-node="onAddNode"
      @zoom-in="zoomIn({ duration: 200 })"
      @zoom-out="zoomOut({ duration: 200 })"
    />
    <Background :gap="30" />

    <template #node-custom="props">
      <CustomNode
        :id="props.id"
        :data="props.data"
        v-model:text="props.data.text"
      />
    </template>

    <template #node-image="props">
      <ImageNode
        :id="props.id"
        :data="props.data"
        :width="props.dimensions.width"
        :height="props.dimensions.height"
      />
    </template>

    <template #edge-custom="props">
      <CustomEdge v-bind="props" />
    </template>
    <FilePanel :isSaving="isSaving"> {{ fileName }} </FilePanel>
  </VueFlow>
</template>

<script setup lang="ts">
import { computed, onUnmounted, ref, watch } from 'vue';
import {
  Edge,
  FlowExportObject,
  MarkerType,
  Node,
  useVueFlow,
  VueFlow,
} from '@vue-flow/core';
import { Background } from '@vue-flow/background';
import { useTopMenuActions } from '@/composables/Flowchart/useTopMenuActions';
import FlowchartButtons from './flowchart/FlowchartControls.vue';
import CustomNode from './flowchart/CustomNode.vue';
import CustomEdge from './flowchart/CustomEdge.vue';
import { CustomNodeData } from '@/types/CustomNodeData';
import { OpenFile, SaveMedia } from '$/filetree/FileTree';
import { useRoute } from 'vue-router';
import { useShowErrorToast } from '@/composables/useShowErrorToast';
import FilePanel from './flowchart/FilePanel.vue';
import { useHandleFlowchartChanges } from '@/composables/Flowchart/useHandleFlowchartChanges';
import { useEventListener } from '@/composables/useEventListener';
import ImageNode from './flowchart/ImageNode.vue';

const path = ref('');
const route = useRoute();
const edges = ref<Edge[]>([]);
useEventListener(window, 'paste', foo);
const { showToast } = useShowErrorToast();
const nodes = ref<Node<CustomNodeData>[]>([]);
const { isSaving } = useHandleFlowchartChanges(path);
const { addNodes, updateNode, fromObject } = useVueFlow();
const { createNewNode, zoomIn, zoomOut } = useTopMenuActions();

async function foo(e: ClipboardEvent) {
  console.log();
  const id = (e.target as HTMLInputElement).id;

  if (!e.clipboardData) {
    return;
  }

  if (
    e.clipboardData.files &&
    e.clipboardData.files.length > 0 &&
    e.clipboardData.files[0].type.startsWith('image/')
  ) {
    const file = e.clipboardData.files[0];
    const mimeType = e.clipboardData.files[0].type;
    const reader = new FileReader();
    reader.onload = async function (e) {
      console.log(e.target?.result);
      try {
        const imagePath = await SaveMedia(
          path.value,
          mimeType,
          e.target?.result as string,
        );
        updateNode<CustomNodeData>(id, {
          type: 'image',
          data: {
            hasFrameDataSection: false,
            image: imagePath,
            text: '',
          },
        });
      } catch (error) {
        showToast(error);
      }
    };
    reader.onerror = function (e) {
      showToast(e.target?.error);
    };
    reader.readAsDataURL(file);
  }
}

const fileName = computed(() =>
  route.params.path.slice(0, route.params.path.indexOf('.')),
);

watch(
  () => route.params.path,
  async () => {
    path.value = route.params.path as string;
    await loadGraph();
  },
  { immediate: true },
);

function onAddNode() {
  addNodes(createNewNode());
}

async function loadGraph() {
  try {
    const path = route.params.path as string;
    const graph = await OpenFile(path);
    fromObject(graph as unknown as FlowExportObject);
  } catch (error) {
    showToast(error);
  }
}

onUnmounted(() => console.log('b'));
</script>
