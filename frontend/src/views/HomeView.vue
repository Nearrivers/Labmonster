<template>
  <ResizablePanelGroup direction="horizontal">
    <ResizablePanel
      :default-size="15"
      :min-size="10"
      collapsible
      class="border-r border-zinc-700 bg-secondary"
    >
      <SidePanel></SidePanel>
    </ResizablePanel>
    <ResizableHandle />
    <ResizablePanel class="bg-primary-foreground">
      <GraphPanel></GraphPanel>
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
