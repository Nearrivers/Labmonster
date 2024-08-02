<template>
  <section class="ml-auto" style="--wails-draggable: no-drag">
    <button
      class="p-2 hover:bg-secondary dark:hover:bg-zinc-700"
      @click="MinimiseWindow"
    >
      <Minus class="h-4 w-6" />
    </button>
    <button
      class="p-2 hover:bg-secondary dark:hover:bg-zinc-700"
      @click="onChangeSizeClick"
    >
      <Maximize class="h-4 w-6" />
    </button>
    <button class="p-2 hover:bg-red-500" @click="QuitApp">
      <X class="h-4 w-6" />
    </button>
  </section>
</template>

<script setup lang="ts">
import { Maximize, Minus, X } from 'lucide-vue-next';
import {
  MaximiseOrUnmaximiseWindow,
  MinimiseWindow,
  QuitApp,
} from '$/topmenu/TopMenu';
import { useShowErrorToast } from '@/composables/useShowErrorToast';

const { showToast } = useShowErrorToast();

async function onChangeSizeClick() {
  try {
    await MaximiseOrUnmaximiseWindow();
  } catch (error) {
    showToast(error);
  }
}
</script>
