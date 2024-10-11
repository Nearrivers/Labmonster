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
    <section class="flex flex-1 flex-col p-1">
      <TopButton
        @click="startScreenRecording"
        :additionnalClasses="'text-muted-foreground hover:text-accent-foreground hover:bg-accent !p-1'"
        :tooltipSide="'right'"
      >
        <template #icon>
          <TvMinimalPlay class="h-5 w-5" />
        </template>
        <template #tooltip> Démarrer une capture d'écran </template>
      </TopButton>
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
import TopButton from '@/components/ui/TopButton.vue';
import { useScreenRecorder } from '@/composables/useScreenRecorder';
import { sidePanelToggled } from '@/events/ToggleSidePanel';
import { PanelLeft, Settings, TvMinimalPlay } from 'lucide-vue-next';

const { startScreenRecording } = useScreenRecorder();

function toggleSidePanel() {
  sidePanelToggled.sidePanelToggled();
}
</script>
