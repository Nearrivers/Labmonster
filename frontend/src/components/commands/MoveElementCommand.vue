<template>
  <AppCommand
    :list="directories"
    :inputPlaceholder="'Saisir le nom d\'un dossier'"
    @select="
      (path: string) => {
        onSelect(path);
        appCommand?.hideModal();
      }
    "
    ref="appCommand"
  />
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useMoveElement } from '@/composables/Commands/useMoveElementCommand';
import AppCommand from '../ui/AppCommand.vue';
import { NodeElement } from '@/types/NodeElement';

const props = defineProps<{
  selectedNode: NodeElement | null;
}>();

const appCommand = ref<InstanceType<typeof AppCommand> | null>(null);
const { directories, loadDirectories, onSelect } = useMoveElement(
  props.selectedNode,
);

async function showModal() {
  await loadDirectories();
  appCommand.value?.showModal();
}

defineExpose({
  showModal,
  hideModal: () => appCommand.value?.hideModal(),
});
</script>
