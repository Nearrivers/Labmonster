<template>
  <ul
    id="filepopover"
    popover
    ref="menu"
    class="fixed z-10 w-56 rounded-md border border-border bg-background p-1.5 text-sm text-primary"
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
    >
      <FolderTree :stroke-width="1.75" class="h-[16px] w-4" />
      Déplacer le fichier vers...
    </li>
    <li
      class="flex cursor-pointer items-center gap-2 rounded-sm px-2 py-1.5 hover:bg-muted hover:text-white"
    >
      <PencilLine :stroke-width="1.75" class="h-[16px] w-4" />
      Renommer
    </li>
    <li
      class="flex cursor-pointer items-center gap-2 rounded-sm px-2 py-1.5 text-red-500 hover:bg-muted"
      @click="onDeleteClick()"
    >
      <Trash2 :stroke-width="1.75" class="h-[16px] w-4" />
      Supprimer
    </li>
    <DeleteFileDialog
      ref="deleteFileDialog"
      :path="props.selectedNode?.dataset.path"
    ></DeleteFileDialog>
  </ul>
</template>

<script setup lang="ts">
import { Files } from 'lucide-vue-next';
import { FolderTree } from 'lucide-vue-next';
import { PencilLine } from 'lucide-vue-next';
import { Trash2 } from 'lucide-vue-next';
import { nextTick, ref } from 'vue';
import DeleteFileDialog from '../AlertDialog/DeleteFileDialog.vue';

const props = defineProps<{
  x: number;
  y: number;
  selectedNode: HTMLLIElement | null;
}>();

const menu = ref<any | null>(null);
const deleteFileDialog = ref<InstanceType<typeof DeleteFileDialog> | null>(
  null,
);

function showPopover() {
  menu.value?.showPopover();
}

function hidePopover() {
  menu.value?.hidePopover();
}

async function onDeleteClick() {
  hidePopover();
  await nextTick();
  deleteFileDialog.value?.openDialog(props.selectedNode?.dataset.path!);
}

function onFileDelete() {}

defineExpose({
  showPopover,
  hidePopover,
});
</script>
