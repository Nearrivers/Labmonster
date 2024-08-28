<template>
  <header class="flex bg-secondary" style="--wails-draggable: drag">
    <section class="p-1">
      <TopButton
        @click="toggleSidePanel"
        :additionnalClasses="'text-muted-foreground hover:text-accent-foreground hover:bg-accent'"
      >
        <template #icon>
          <PanelLeft class="h-5 w-5" />
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
