<template>
  <Panel :position="'top-right'" class="flex flex-col gap-2">
    <section
      class="overflow-hidden rounded-md border bg-background shadow-md dark:border-zinc-600 dark:bg-secondary"
    >
      <ControlButton @click="emit('zoomIn')">
        <ZoomIn :stroke-width="1.75" class="h-4 w-4" />
        <template #tooltip> Zoomer </template>
      </ControlButton>
      <ControlButton @click="emit('zoomOut')">
        <ZoomOut :stroke-width="1.75" class="h-4 w-4" />
        <template #tooltip> Dézoomer </template>
      </ControlButton>
    </section>
    <section
      class="flex overflow-hidden rounded-md border bg-background shadow-md dark:border-zinc-600 dark:bg-secondary"
    >
      <ControlButton>
        <CircleHelp :stroke-width="1.75" class="h-4 w-4" />
        <template #tooltip> Aide </template>
      </ControlButton>
    </section>
    <section
      class="flex overflow-hidden rounded-md border bg-background shadow-md dark:border-zinc-600 dark:bg-secondary"
    >
      <ControlButton>
        <Settings :stroke-width="1.75" class="h-4 w-4" />
        <template #tooltip> Paramètres du setup </template>
      </ControlButton>
    </section>
  </Panel>
</template>

<script setup lang="ts">
import { Panel } from '@vue-flow/core';
import { CircleHelp, Settings, ZoomIn, ZoomOut } from 'lucide-vue-next';
import ControlButton from './FlowchartControls/ControlButton.vue';

const emit = defineEmits<{
  (e: 'zoomIn'): void;
  (e: 'zoomOut'): void;
}>();
</script>
