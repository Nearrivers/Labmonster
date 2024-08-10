<template>
  <AppCommand
    :list="recentlyOpendedFiles"
    :input-placeholder="`Saisir le nom d'un fichier vers lequel basculer`"
    ref="appCommand"
    @select="(path: string) => onSelect(path, appCommand!.hideModal)"
  />
</template>

<script setup lang="ts">
import { useRecentlyOpendedCommand } from '@/composables/Commands/useRecentlyOpenedCommand';
import AppCommand from '../ui/AppCommand.vue';
import { ref } from 'vue';

const appCommand = ref<InstanceType<typeof AppCommand> | null>(null);
const { recentlyOpendedFiles, onSelect } =
  useRecentlyOpendedCommand(appCommand);
</script>
