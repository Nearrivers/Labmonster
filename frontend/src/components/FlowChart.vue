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

    <template #node-video="props">
      <VideoNode
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
import { ref } from 'vue';
import { Edge, MarkerType, Node, useVueFlow, VueFlow } from '@vue-flow/core';
import { Background } from '@vue-flow/background';
import { useTopMenuActions } from '@/composables/Flowchart/useTopMenuActions';
import FlowchartButtons from './flowchart/FlowchartControls.vue';
import CustomNode from './flowchart/CustomNode.vue';
import CustomEdge from './flowchart/CustomEdge.vue';
import { CustomNodeData } from '@/types/CustomNodeData';
import FilePanel from './flowchart/FilePanel.vue';
import { useHandleFlowchartChanges } from '@/composables/Flowchart/useHandleFlowchartChanges';
import ImageNode from './flowchart/ImageNode.vue';
import { useFlowChart } from '@/composables/Flowchart/useFlowChart';
import VideoNode from './flowchart/VideoNode.vue';
import FlowchartContextMenu from './contextmenus/FlowchartContextMenu.vue';

const edges = ref<Edge[]>([]);
const nodes = ref<Node<CustomNodeData>[]>([]);
const { addNodes } = useVueFlow();
const { createNewNode, zoomIn, zoomOut } = useTopMenuActions();
const { path, fileName } = useFlowChart();
const { isSaving } = useHandleFlowchartChanges(path);

function onAddNode() {
  addNodes(createNewNode());
}
</script>
