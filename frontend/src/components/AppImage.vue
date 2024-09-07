<template>
  <main class="h-full w-full">
    <header class="mb-4 py-2 text-center text-sm opacity-65">
      {{ path }}
    </header>
    <div class="flex w-full justify-center px-2">
      <img
        :src="src"
        alt="Fichier image"
        class="max-w-full object-cover text-center"
      />
    </div>
  </main>
</template>

<script setup lang="ts">
import { OpenMedia } from '$/file_handler/FileHandler';
import { useShowErrorToast } from '@/composables/useShowErrorToast';
import { computed, ref, watchEffect } from 'vue';
import { useRoute } from 'vue-router';

const src = ref('');
const route = useRoute();
const { showToast } = useShowErrorToast();

const path = computed(() =>
  (route.params.path as string).replaceAll('/', ' / '),
);

watchEffect(async () => {
  try {
    src.value = await OpenMedia(route.params.path as string);
  } catch (error) {
    showToast(error);
  }
});
</script>
