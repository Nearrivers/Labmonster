<template>
  <header class="flex border-b border-b-border" style="--wails-draggable: drag">
    <section class="p-1">
      <TopButton @click="toggleSidePanel">
        <template #icon>
          <PanelLeft class="h-5 w-5 text-muted-foreground" />
        </template>
        <template #tooltip> Basculer </template>
      </TopButton>
    </section>
    <SysButtons />
  </header>
</template>

<script setup lang="ts">
import { sidePanelToggled } from '@/events/ToggleSidePanel';
import SysButtons from './AppHeader/SysButtons.vue';
import TopButton from './ui/TopButton.vue';
import { PanelLeft } from 'lucide-vue-next';

function toggleSidePanel() {
  sidePanelToggled.sidePanelToggled();
}
</script>
