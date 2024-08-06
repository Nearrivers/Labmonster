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
      <CustomNode :id="props.id" :data="props.data" />
    </template>

    <template #edge-custom="props">
      <CustomEdge v-bind="props" />
    </template>
  </VueFlow>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import {
  Edge,
  EdgeChange,
  MarkerType,
  NodeChange,
  useVueFlow,
  VueFlow,
  VueFlowStore,
} from '@vue-flow/core';
import { Background } from '@vue-flow/background';
import { useTopMenuActions } from '@/composables/Flowchart/useTopMenuActions';
import FlowchartButtons from './flowchart/FlowchartControls.vue';
import CustomNode from './flowchart/CustomNode.vue';
import CustomEdge from './flowchart/CustomEdge.vue';

const nodes = ref([
  {
    id: '1',
    position: { x: 25, y: 90 },
    type: 'custom',
    data: { hello: 'test' },
  },
  {
    id: '2',
    position: { x: 45, y: 200 },
    type: 'custom',
    data: { hello: 'autre test' },
  },
]);
const edges = ref<Edge[]>([
  {
    id: '1->2',
    source: '1',
    target: '2',
    type: 'custom',
  },
]);
const { addNodes, onNodesChange, onEdgesChange, onInit, toObject } =
  useVueFlow();
const { createNewNode, zoomIn, zoomOut } = useTopMenuActions(nodes);

function onAddNode() {
  addNodes(createNewNode());
}

onInit((param: VueFlowStore) => {
  console.log(toObject());
});

onNodesChange((param: NodeChange[]) => {});
onEdgesChange((param: EdgeChange[]) => {});
</script>

<style scoped>
.vue-flow__node-custom.selected {
  background-color: red !important;
}
</style>
