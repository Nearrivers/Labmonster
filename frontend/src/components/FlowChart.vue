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

    <template #edge-custom="props">
      <CustomEdge v-bind="props" />
    </template>
    <FilePanel />
  </VueFlow>
</template>

<script setup lang="ts">
import { onActivated, ref, watch } from 'vue';
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
import { OpenFile } from '$/filetree/FileTreeExplorer';
import { useRoute } from 'vue-router';
import { useShowErrorToast } from '@/composables/useShowErrorToast';
import FilePanel from './flowchart/FilePanel.vue';
import { useHandleFlowchartChanges } from '@/composables/Flowchart/useHandleFlowchartChanges';

const { showToast } = useShowErrorToast();
const route = useRoute();
const nodes = ref<Node<CustomNodeData>[]>([]);
const edges = ref<Edge[]>([]);
const { addNodes, fromObject } = useVueFlow();
const { createNewNode, zoomIn, zoomOut } = useTopMenuActions(nodes);
useHandleFlowchartChanges(route.params.path as string);

watch(() => route.params.path, loadGraph, { immediate: true });

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
</script>
