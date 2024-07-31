<template>
  <ul
    id="dirpopover"
    popover
    ref="menu"
    class="fixed z-10 min-w-52 rounded-lg border border-border bg-background text-xs text-primary"
    :style="{ top: y + 'px', left: x + 'px' }"
  >
    <li>
      <ul class="border-b border-b-border p-1.5">
        <li
          class="flex cursor-default items-center gap-2 rounded-sm px-2 py-1.5 hover:bg-muted"
        >
          <SquarePen :stroke-width="1.75" class="h-[16px] w-4" />
          <p>Nouveau setup</p>
        </li>
        <li
          class="flex cursor-default items-center gap-2 rounded-sm px-2 py-1.5 hover:bg-muted"
        >
          <FolderOpen :stroke-width="1.75" class="h-[16px] w-4" />
          <p>Nouveau dossier</p>
        </li>
      </ul>
    </li>
    <li>
      <ul class="border-b border-b-border p-1.5">
        <li
          class="flex cursor-default items-center gap-2 rounded-sm px-2 py-1.5 hover:bg-muted"
        >
          <Files :stroke-width="1.75" class="h-[16px] w-4" />
          <p>Dupliquer</p>
        </li>
        <li
          class="flex cursor-default items-center gap-2 rounded-sm px-2 py-1.5 hover:bg-muted"
        >
          <FolderTree :stroke-width="1.75" class="h-[16px] w-4" />
          Déplacer le dossier vers...
        </li>
      </ul>
    </li>
    <li>
      <ul class="p-1.5">
        <li
          class="flex cursor-default items-center gap-2 rounded-sm px-2 py-1.5 hover:bg-muted"
        >
          <PencilLine :stroke-width="1.75" class="h-[16px] w-4" />
          Renommer
        </li>
        <li
          class="flex cursor-default items-center gap-2 rounded-sm px-2 py-1.5 text-red-500 hover:bg-muted"
        >
          <Trash2 :stroke-width="1.75" class="h-[16px] w-4" />
          Supprimer
        </li>
      </ul>
    </li>

    <!-- <MoveFileCommand
      ref="moveFileCommand"
      :key="props.x"
      :selected-node="selectedNode"
    />
    <DeleteFileDialog
      ref="deleteFileDialog"
      :path="selectedNode?.dataset.path"
      :extension="selectedNode?.dataset.extension"
    /> -->
  </ul>
</template>

<script setup lang="ts">
import { ContextMenuProps } from '@/types/props/ContextMenuProps';
import {
  Files,
  FolderOpen,
  FolderTree,
  PencilLine,
  SquarePen,
  Trash2,
} from 'lucide-vue-next';
import { ref } from 'vue';

const props = defineProps<ContextMenuProps>();

const menu = ref<any | null>(null);

function showPopover() {
  menu.value?.showPopover();
}

function hidePopover() {
  menu.value?.hidePopover();
}

defineExpose({
  showPopover,
  hidePopover,
});
</script>
