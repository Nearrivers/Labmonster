<template>
  <section class="ml-auto" style="--wails-draggable: no-drag">
    <TopButton class="rounded-none" @click="MinimiseWindow">
      <template #icon>
        <Minus class="h-full w-7 p-1" />
      </template>
      <template #tooltip>Minimiser</template>
    </TopButton>
    <TopButton class="rounded-none" @click="onChangeSizeClick">
      <template #icon>
        <Maximize class="h-full w-7 p-1" />
      </template>
      <template #tooltip>Restaurer</template>
    </TopButton>
    <TopButton class="rounded-none hover:!bg-red-600" @click="QuitApp">
      <template #icon>
        <X class="h-full w-7 p-1" />
      </template>
      <template #tooltip>Fermer</template>
    </TopButton>
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
import TopButton from '../ui/TopButton.vue';

const { showToast } = useShowErrorToast();

async function onChangeSizeClick() {
  try {
    await MaximiseOrUnmaximiseWindow();
  } catch (error) {
    showToast(error);
  }
}
</script>
