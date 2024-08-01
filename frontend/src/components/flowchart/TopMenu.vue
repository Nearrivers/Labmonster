<template>
  <Panel :position="'top-left'">
    <section
      class="flex overflow-hidden rounded-md border bg-background shadow-md dark:border-zinc-600 dark:bg-secondary"
    >
      <ControlButton @click="emit('addNode')">
        <Plus :stroke-width="1.75" class="h-4 w-4" />
        <template #tooltip> Ajouter un noeud </template>
      </ControlButton>
      <ControlButton @click="emit('zoomIn')">
        <ZoomIn :stroke-width="1.75" class="h-4 w-4" />
        <template #tooltip> Zoomer </template>
      </ControlButton>
      <ControlButton @click="emit('zoomOut')">
        <ZoomOut :stroke-width="1.75" class="h-4 w-4" />
        <template #tooltip> Dézoomer </template>
      </ControlButton>
    </section>
  </Panel>
</template>

<script setup lang="ts">
import { Panel } from '@vue-flow/core';
import { Plus, ZoomIn, ZoomOut } from 'lucide-vue-next';
import ControlButton from './TopMenu/ControlButton.vue';

const emit = defineEmits<{
  (e: 'addNode'): void;
  (e: 'zoomIn'): void;
  (e: 'zoomOut'): void;
}>();
</script>
