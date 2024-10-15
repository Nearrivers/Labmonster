<template>
  <RouterView></RouterView>
  <CreateLab v-if="!isConfigFilePresent"></CreateLab>
  <Toaster />
  <RecentlyOpenedFileCommand />
</template>

<script lang="ts" setup>
import { useColorMode } from '@vueuse/core';
import CreateLab from './components/config/CreateLab.vue';
import { onMounted, ref } from 'vue';
import { CheckConfigPresenceAndLoadIt } from '$/config/AppConfig';
import { AddGame } from '$/games/GameRepository';
import { Toaster } from '@/components/ui/toast';
import RecentlyOpenedFileCommand from './components/commands/RecentlyOpenedFileCommand.vue';
import { repository } from '$/models';
import { useShowErrorToast } from './composables/useShowErrorToast';

const { showToast } = useShowErrorToast();
const isConfigFilePresent = ref(false);
const mode = useColorMode();

onMounted(async () => {
  try {
    isConfigFilePresent.value = await CheckConfigPresenceAndLoadIt();
    await AddGame({
      name: 'test',
      iconpath: 'test',
    } as repository.AddGameParams);
  } catch (error) {
    showToast(error);
  }
});
</script>
