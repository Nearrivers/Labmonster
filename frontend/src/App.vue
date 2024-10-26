<template>
  <RouterView></RouterView>
  <CreateLab v-if="!isConfigFilePresent"></CreateLab>
  <Toaster />
  <RecentlyOpenedFileCommand />
  <SettingsDialog />
</template>

<script lang="ts" setup>
import { useColorMode } from '@vueuse/core';
import CreateLab from './components/config/CreateLab.vue';
import SettingsDialog from '@/components/dialogs/SettingsDialog.vue';
import { onMounted, ref } from 'vue';
import { CheckConfigPresenceAndLoadIt } from '$/config/AppConfig';
import { Toaster } from '@/components/ui/toast';
import RecentlyOpenedFileCommand from './components/commands/RecentlyOpenedFileCommand.vue';
import { useShowErrorToast } from './composables/useShowErrorToast';

const { showToast } = useShowErrorToast();
const isConfigFilePresent = ref(false);
// const mode = useColorMode();

onMounted(async () => {
  try {
    isConfigFilePresent.value = await CheckConfigPresenceAndLoadIt();
  } catch (error) {
    showToast(error);
  }
});
</script>
