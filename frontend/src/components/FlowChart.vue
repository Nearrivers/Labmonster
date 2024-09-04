<template>
  <VueFlow
    :nodes="nodes"
    class="h-full"
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
    :multi-selection-key-code="['Shift']"
    @click.right.prevent="onFlowRightClick"
  >
    <FlowchartButtons
      @add-node="onAddNode"
      @zoom-in="zoomIn({ duration: 200 })"
      @zoom-out="zoomOut({ duration: 200 })"
    />
    <Background :gap="30" />

    <template #node-custom="props">
      <TextNode
        :id="props.id"
        :data="props.data"
        v-model:text="props.data.text"
      />
    </template>

    <template #node-image="props">
      <ImageNode :id="props.id" :data="props.data" />
    </template>

    <template #node-video="props">
      <VideoNode :id="props.id" :data="props.data" />
    </template>

    <template #edge-custom="props">
      <CustomEdge v-bind="props" />
    </template>
    <FilePanel :isSaving="isSaving"> {{ fileName }} </FilePanel>
  </VueFlow>
  <FlowchartContextMenu
    :x="contextMenuX"
    :y="contextMenuY"
    popover-id="flowchartPopover"
    ref="ctxMenu"
  />
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { Edge, MarkerType, Node, useVueFlow, VueFlow } from '@vue-flow/core';
import { Background } from '@vue-flow/background';
import { useTopMenuActions } from '@/composables/Flowchart/useTopMenuActions';
import FlowchartButtons from './flowchart/FlowchartControls.vue';
import CustomEdge from './flowchart/CustomEdge.vue';
import { CustomNodeData } from '@/types/CustomNodeData';
import FilePanel from './flowchart/FilePanel.vue';
import { useHandleFlowchartChanges } from '@/composables/Flowchart/useHandleFlowchartChanges';
import { useFlowChart } from '@/composables/Flowchart/useFlowChart';
import FlowchartContextMenu from './contextmenus/FlowchartContextMenu.vue';
import { EventsOn } from '../../wailsjs/runtime/runtime';
import TextNode from './flowchart/Nodes/TextNode.vue';
import VideoNode from './flowchart/Nodes/VideoNode.vue';
import ImageNode from './flowchart/Nodes/ImageNode.vue';

const edges = ref<Edge[]>([]);
const contextMenuX = ref(100);
const contextMenuY = ref(100);
const nodes = ref<Node<CustomNodeData>[]>([]);
const { addNodes } = useVueFlow();
const { createNewNode, zoomIn, zoomOut } = useTopMenuActions();
const { path, fileName, onFsEvent } = useFlowChart();
const { isSaving } = useHandleFlowchartChanges(path);
const ctxMenu = ref<InstanceType<typeof FlowchartContextMenu> | null>(null);

function onFlowRightClick(e: MouseEvent) {
  contextMenuX.value = e.clientX;
  contextMenuY.value = e.clientY;
  ctxMenu.value?.showPopover();
}

EventsOn('fsop', onFsEvent);

function onAddNode() {
  addNodes(createNewNode());
}
</script>
