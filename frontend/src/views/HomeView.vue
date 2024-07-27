<template>
  <ResizablePanelGroup direction="horizontal">
    <ResizablePanel :default-size="15" :min-size="10" collapsible>
      <SidePanel />
    </ResizablePanel>
    <ResizableHandle />
    <ResizablePanel>
      <GraphPanel />
    </ResizablePanel>
  </ResizablePanelGroup>
</template>

<script setup lang="ts">
import {
  ResizableHandle,
  ResizablePanel,
  ResizablePanelGroup,
} from '@/components/ui/resizable';
import SidePanel from './SidePanel.vue';
import GraphPanel from './GraphPanel.vue';
</script>
