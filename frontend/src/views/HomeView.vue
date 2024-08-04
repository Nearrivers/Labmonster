<template>
  <ResizablePanelGroup direction="horizontal" class="h-[calc(100%-33px)]">
    <ResizablePanel
      :default-size="15"
      :min-size="10"
      collapsible
      ref="resizablePanel"
      id="sidePanel"
    >
      <SidePanel />
    </ResizablePanel>
    <ResizableHandle
      class="w-[2px] transition-all hover:bg-primary hover:ring-1 hover:ring-primary"
      :hit-area-margins="{ fine: 0, coarse: 1 }"
    />
    <ResizablePanel class="h-full">
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
import { SIDE_PANEL_TOGGLED } from '@/constants/event-names/SIDE_PANEL_TOGGLED';
import { sidePanelToggled } from '@/events/ToggleSidePanel';
import { useEventListener } from '@/composables/useEventListener';
import { ref } from 'vue';

const resizablePanel = ref<InstanceType<typeof ResizablePanel> | null>(null);
useEventListener(sidePanelToggled, SIDE_PANEL_TOGGLED, toggleSidePanel);

function toggleSidePanel() {
  resizablePanel.value?.isCollapsed
    ? resizablePanel.value?.expand()
    : resizablePanel.value?.collapse();
}
</script>
