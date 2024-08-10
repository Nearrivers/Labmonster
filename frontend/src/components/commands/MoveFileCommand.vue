<template>
  <AppCommand
    :list="directories"
    :inputPlaceholder="'Saisir le nom d\'un dossier'"
    @select="(path: string) => onSelect(path, appCommand!.hideModal)"
    ref="appCommand"
  />
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useMoveFile } from '@/composables/Commands/useMoveFileCommand';
import AppCommand from '../ui/AppCommand.vue';

const props = defineProps<{
  selectedNode: HTMLLIElement | null;
}>();

const appCommand = ref<InstanceType<typeof AppCommand> | null>(null);
const { directories, onSelect } = useMoveFile(props.selectedNode);

defineExpose({
  showModal: () => appCommand.value?.showModal(),
  hideModal: () => appCommand.value?.hideModal(),
});
</script>
