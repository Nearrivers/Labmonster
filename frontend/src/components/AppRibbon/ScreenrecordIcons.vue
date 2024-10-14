<template>
  <TopButton
    @click="startScreenRecording"
    :additionnalClasses="'text-muted-foreground hover:text-accent-foreground hover:bg-accent !p-1'"
    :tooltipSide="'right'"
    v-if="mediaRecorderState === 'inactive'"
  >
    <template #icon>
      <TvMinimalPlay class="h-5 w-5" />
    </template>
    <template #tooltip> Démarrer un enregistrement d'écran </template>
  </TopButton>
  <TopButton
    @click="stopScreenRecording"
    v-else
    :additionnalClasses="'bg-primary text-secondary hover:bg-primary hover:text-secondary !p-1'"
    :tooltipSide="'right'"
  >
    <template #icon> <MonitorX class="h-5 w-5" /> </template>
    <template #tooltip> Terminer l'enregistrement d'écran </template>
  </TopButton>
</template>

<script setup lang="ts">
import { MonitorX, TvMinimalPlay } from 'lucide-vue-next';
import TopButton from '../ui/TopButton.vue';
import { useScreenRecorder } from '@/composables/useScreenRecorder';

const { mediaRecorderState, stopScreenRecording, startScreenRecording } =
  useScreenRecorder();
</script>
