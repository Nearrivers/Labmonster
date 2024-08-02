<template>
  <ResizablePanelGroup direction="horizontal" class="h-[calc(100%-33px)]">
    <ResizablePanel :default-size="15" :min-size="10" collapsible>
      <SidePanel />
    </ResizablePanel>
    <ResizableHandle
      class="w-[2px] transition-all hover:bg-primary hover:ring-1 hover:ring-primary"
      :hit-area-margins="{ fine: 0, coarse: 1 }"
    />
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
