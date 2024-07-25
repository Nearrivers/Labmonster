<template>
  <ul
    id="filepopover"
    popover
    ref="menu"
    class="fixed z-10 rounded-md border border-border bg-background p-1.5 text-xs text-primary"
    :style="{ top: y + 'px', left: x + 'px' }"
  >
    <li
      class="flex cursor-pointer items-center gap-2 rounded-sm px-2 py-1.5 hover:bg-muted hover:text-white"
    >
      <Files :stroke-width="1.75" class="h-[16px] w-4" />
      <p>Dupliquer</p>
    </li>
    <li
      class="flex cursor-pointer items-center gap-2 rounded-sm px-2 py-1.5 hover:bg-muted hover:text-white"
      @click="onMoveClick"
    >
      <FolderTree :stroke-width="1.75" class="h-[16px] w-4" />
      Déplacer le fichier vers...
    </li>
    <li
      class="flex cursor-pointer items-center gap-2 rounded-sm px-2 py-1.5 hover:bg-muted hover:text-white"
      @click="onRenameClick(props.selectedNode?.dataset.path!)"
    >
      <PencilLine :stroke-width="1.75" class="h-[16px] w-4" />
      Renommer
    </li>
    <li
      class="flex cursor-pointer items-center gap-2 rounded-sm px-2 py-1.5 text-red-500 hover:bg-muted"
      @click="onDeleteClick(props.selectedNode?.dataset.path!)"
    >
      <Trash2 :stroke-width="1.75" class="h-[16px] w-4" />
      Supprimer
    </li>
    <MoveFileCommand
      ref="moveFileCommand"
      :oldPath="props.selectedNode?.dataset.path"
      :extension="props.selectedNode?.dataset.extension"
    />
    <DeleteFileDialog
      ref="deleteFileDialog"
      :path="props.selectedNode?.dataset.path"
    />
  </ul>
</template>

<script setup lang="ts">
import { Files } from 'lucide-vue-next';
import { FolderTree, Trash2, PencilLine } from 'lucide-vue-next';
import { ref } from 'vue';
import DeleteFileDialog from '../AlertDialog/DeleteFileDialog.vue';
import { ContextMenuProps } from '@/types/props/ContextMenuProps';
import { useNodeContextMenu } from '@/composables/ContextMenus/useNodeContextMenu';
import MoveFileCommand from '../commands/MoveFileCommand.vue';

const props = defineProps<
  ContextMenuProps & {
    selectedNode: HTMLLIElement | null;
  }
>();

const deleteFileDialog = ref<InstanceType<typeof DeleteFileDialog> | null>(
  null,
);
const moveFileCommand = ref<InstanceType<typeof MoveFileCommand> | null>(null);
const open = ref(false);
const { menu, showPopover, hidePopover, onRenameClick, onDeleteClick } =
  useNodeContextMenu(deleteFileDialog);

function onMoveClick() {
  moveFileCommand.value?.showModal();
  hidePopover();
}

defineExpose({
  showPopover,
  hidePopover,
});
</script>
