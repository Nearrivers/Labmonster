<template>
  <section class="flex justify-center gap-[2px] py-2 text-muted-foreground">
    <DataButtons
      @createFile="createNewFileAtRoot"
      @create-folder="createNewFolderAtRoot"
    />
  </section>
  <ScrollArea class="b-4 h-[90svh]" data-path="/">
    <ul
      class="w-full px-1 text-sm text-muted-foreground"
      v-if="files.length > 0"
      @click.right.prevent="onRightClick"
      @click.left="onLeftClick"
    >
      <template v-for="(file, index) in files" :key="file.name">
        <FileNode
          v-if="file.type === 'FILE'"
          :fileNode="file"
          path=""
          :data-id="index"
          :offset="14.5"
        />
        <DirNode
          v-else
          :dirNode="file"
          path=""
          :data-id="index"
          :offset="14.5"
        />
      </template>
    </ul>
  </ScrollArea>
  <FileContextMenu
    ref="fileContextMenu"
    :x="contextMenuX"
    :y="contextMenuY"
    :selected-node="selectedNode"
    @move="onMoveClick"
  />
  <DirContextMenu
    ref="dirContextMenu"
    :x="contextMenuX"
    :y="contextMenuY"
    :selected-node="selectedNode"
    @move="onMoveClick"
  />
  <MoveElementCommand
    :selectedNode="selectedNode"
    :key="contextMenuX"
    ref="moveElementCommand"
  />
</template>

<script setup lang="ts">
import { onMounted, ref, watch } from 'vue';
import { ScrollArea } from '@/components/ui/scroll-area';
import FileNode from '@/components/sidepanel/FileNode.vue';
import DirNode from '@/components/sidepanel/DirNode.vue';
import { useSidePanel } from '@/composables/useSidePanel';
import { CheckConfigPresenceAndLoadIt } from '$/config/AppConfig';
import FileContextMenu from '@/components/contextmenus/FileContextMenu.vue';
import DirContextMenu from '@/components/contextmenus/DirContextMenu.vue';
import { useFiletree } from '@/composables/useFiletree';
import { useMagicKeys } from '@vueuse/core';
import DataButtons from '@/components/sidepanel/DataButtons.vue';
import MoveElementCommand from '../components/commands/MoveElementCommand.vue';

const {
  files,
  contextMenuX,
  contextMenuY,
  fileContextMenu,
  dirContextMenu,
  selectedNode,
  loadLabFiles,
  onRightClick,
  createNewFileAtRoot,
  createNewFolderAtRoot,
  showToast,
  onLeftClick,
} = useSidePanel();

useFiletree(files, showToast);
const keys = useMagicKeys();
const F2 = keys['F2'];
const moveElementCommand = ref<InstanceType<typeof MoveElementCommand> | null>(
  null,
);

watch(F2, (v) => {
  if (!v) {
    return;
  }

  // console.log('test');
});

onMounted(async () => {
  try {
    const isConfigFilePresent = await CheckConfigPresenceAndLoadIt();

    if (isConfigFilePresent) {
      await loadLabFiles();
    }
  } catch (error) {
    showToast(error);
  }
});

async function onMoveClick() {
  await moveElementCommand.value?.showModal();
}
</script>
