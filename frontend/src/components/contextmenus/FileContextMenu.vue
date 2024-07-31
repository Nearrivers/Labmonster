<template>
  <ul
    id="filepopover"
    popover
    ref="menu"
    class="fixed z-10 rounded-md border border-border bg-background p-1.5 text-xs text-primary"
    :style="{ top: y + 'px', left: x + 'px' }"
  >
    <li
      class="flex cursor-pointer items-center gap-2 rounded-sm px-2 py-1.5 hover:bg-muted"
      @click="
        onDuplicateClick(
          selectedNode?.dataset.path!,
          fileName,
          selectedNode?.dataset.extension!,
        )
      "
    >
      <Files :stroke-width="1.75" class="h-[16px] w-4" />
      <p>Dupliquer</p>
    </li>
    <li
      class="flex cursor-pointer items-center gap-2 rounded-sm px-2 py-1.5 hover:bg-muted"
      @click="onMoveClick"
    >
      <FolderTree :stroke-width="1.75" class="h-[16px] w-4" />
      Déplacer le fichier vers...
    </li>
    <li
      class="flex cursor-pointer items-center gap-2 rounded-sm px-2 py-1.5 hover:bg-muted"
      @click="onRenameClick(selectedNode?.dataset.path!)"
    >
      <PencilLine :stroke-width="1.75" class="h-[16px] w-4" />
      Renommer
    </li>
    <li
      class="flex cursor-pointer items-center gap-2 rounded-sm px-2 py-1.5 text-red-500 hover:bg-muted"
      @click="onDeleteClick(selectedNode?.dataset.path!)"
    >
      <Trash2 :stroke-width="1.75" class="h-[16px] w-4" />
      Supprimer
    </li>
    <MoveFileCommand
      ref="moveFileCommand"
      :key="props.x"
      :selected-node="selectedNode"
    />
    <DeleteFileDialog
      ref="deleteFileDialog"
      :path="selectedNode?.dataset.path"
    />
  </ul>
</template>

<script setup lang="ts">
import { Files } from 'lucide-vue-next';
import { FolderTree, Trash2, PencilLine } from 'lucide-vue-next';
import { computed, ref } from 'vue';
import DeleteFileDialog from '../AlertDialog/DeleteFileDialog.vue';
import { ContextMenuProps } from '@/types/props/ContextMenuProps';
import { useNodeContextMenu } from '@/composables/ContextMenus/useNodeContextMenu';
import MoveFileCommand from '../commands/MoveFileCommand.vue';

const props = defineProps<
  ContextMenuProps & {
    selectedNode: HTMLLIElement | null;
  }
>();

const fileName = computed(
  () =>
    props.selectedNode?.dataset.path?.slice(
      props.selectedNode.dataset.path.lastIndexOf('/') + 1,
    ) || '',
);

const deleteFileDialog = ref<InstanceType<typeof DeleteFileDialog> | null>(
  null,
);
const moveFileCommand = ref<InstanceType<typeof MoveFileCommand> | null>(null);
const {
  menu,
  showPopover,
  hidePopover,
  onRenameClick,
  onDeleteClick,
  onDuplicateClick,
} = useNodeContextMenu(deleteFileDialog);

function onMoveClick() {
  moveFileCommand.value?.showModal();
  hidePopover();
}

defineExpose({
  showPopover,
  hidePopover,
});
</script>
