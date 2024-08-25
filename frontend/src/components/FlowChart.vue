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
import CustomNode from './flowchart/CustomNode.vue';
import CustomEdge from './flowchart/CustomEdge.vue';
import { CustomNodeData } from '@/types/CustomNodeData';
import FilePanel from './flowchart/FilePanel.vue';
import { useHandleFlowchartChanges } from '@/composables/Flowchart/useHandleFlowchartChanges';
import ImageNode from './flowchart/ImageNode.vue';
import { useFlowChart } from '@/composables/Flowchart/useFlowChart';
import VideoNode from './flowchart/VideoNode.vue';
import FlowchartContextMenu from './contextmenus/FlowchartContextMenu.vue';
import { useRoute, useRouter } from 'vue-router';
import { EventsOn } from '../../wailsjs/runtime/runtime';
import { FsEvent } from '@/types/FsEvent';
import { watcher } from '$/models';
import { Routes } from '@/types/Routes';

const router = useRouter();
const edges = ref<Edge[]>([]);
const contextMenuX = ref(100);
const contextMenuY = ref(100);
const nodes = ref<Node<CustomNodeData>[]>([]);
const { addNodes } = useVueFlow();
const { createNewNode, zoomIn, zoomOut } = useTopMenuActions();
const { path, fileName } = useFlowChart();
const { isSaving } = useHandleFlowchartChanges(path);
const ctxMenu = ref<InstanceType<typeof FlowchartContextMenu> | null>(null);

function onFlowRightClick(e: MouseEvent) {
  contextMenuX.value = e.clientX;
  contextMenuY.value = e.clientY;
  ctxMenu.value?.showPopover();
}

const route = useRoute();

EventsOn('fsop', (e: FsEvent) => {
  const filePath = e.path + '/' + e.file;
  if (
    // We skip if the operation is a file creation
    e.op === watcher.Op.CREATE || // OR
    // If the operation is a delete but the deleted file is not the one currently opened, we skip
    (filePath != route.params.path && e.op === watcher.Op.REMOVE) || // OR
    // If the operation is a move or a rename and the old path is different from the one of the file currently opened, we skip
    (e.oldPath != route.params.path &&
      (e.op === watcher.Op.MOVE || e.op === watcher.Op.RENAME))
  ) {
    return;
  }

  if (e.op === watcher.Op.REMOVE) {
    router.push({ name: Routes.NotOpened });
    return;
  }

  router.push({ name: Routes.Flowchart, params: { path: filePath } });
});

function onAddNode() {
  addNodes(createNewNode());
}
</script>
