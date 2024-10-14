<template>
  <aside class="flex h-full flex-col border-r border-r-border bg-secondary">
    <header class="p-1">
      <TopButton
        @click="toggleSidePanel"
        :additionnalClasses="'text-muted-foreground hover:text-accent-foreground hover:bg-accent !p-1'"
      >
        <template #icon>
          <PanelLeft class="h-5 w-5" />
        </template>
        <template #tooltip> Basculer </template>
      </TopButton>
    </header>
    <section class="flex flex-1 flex-col gap-1 p-1">
      <ScreenrecordIcons />
      <TopButton
        :additionnalClasses="'text-muted-foreground mt-auto hover:text-accent-foreground hover:bg-accent !p-1'"
      >
        <template #icon>
          <Settings class="h-5 w-5" />
        </template>
        <template #tooltip> Paramètres </template>
      </TopButton>
    </section>
  </aside>
</template>

<script lang="ts" setup>
import ScreenrecordIcons from '@/components/AppRibbon/ScreenrecordIcons.vue';
import TopButton from '@/components/ui/TopButton.vue';
import { sidePanelToggled } from '@/events/ToggleSidePanel';
import { PanelLeft, Settings } from 'lucide-vue-next';

function toggleSidePanel() {
  sidePanelToggled.sidePanelToggled();
}
</script>
