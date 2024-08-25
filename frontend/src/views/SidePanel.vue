<template>
  <header class="flex justify-center gap-[2px] py-2 text-muted-foreground">
    <TopButtons @createFile="createNewFileAtRoot" />
  </header>
  <ScrollArea class="h-[90svh] pb-4" data-path="/">
    <ul
      class="w-full px-2 text-sm text-muted-foreground"
      v-if="files.length > 0"
      @click.right.prevent="onRightClick"
      @click.left="onLeftClick"
    >
      <template v-for="(file, index) in files" :key="file.name">
        <FileNode
          v-if="file.type === 'FILE'"
          :node="file"
          path=""
          :data-id="index"
        />
        <DirNode v-else :node="file" path="" :data-id="index" />
      </template>
    </ul>
  </ScrollArea>
  <FileContextMenu
    ref="fileContextMenu"
    :x="contextMenuX"
    :y="contextMenuY"
    :selected-node="selectedNode"
  />
  <DirContextMenu ref="dirContextMenu" :x="contextMenuX" :y="contextMenuY" />
</template>

<script setup lang="ts">
import { onMounted } from 'vue';
import { ScrollArea } from '@/components/ui/scroll-area';
import FileNode from '@/components/sidepanel/FileNode.vue';
import DirNode from '@/components/sidepanel/DirNode.vue';
import { useSidePanel } from '@/composables/useSidePanel';
import { CheckConfigPresenceAndLoadIt } from '$/config/AppConfig';
import FileContextMenu from '@/components/contextmenus/FileContextMenu.vue';
import DirContextMenu from '@/components/contextmenus/DirContextMenu.vue';
import TopButtons from '@/components/sidepanel/TopButtons.vue';
import { useFiletree } from '@/composables/useFiletree';

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
  showToast,
  onLeftClick,
} = useSidePanel();

useFiletree(files, showToast);

onMounted(async () => {
  try {
    const isConfigFilePresent = await CheckConfigPresenceAndLoadIt();

    if (isConfigFilePresent) {
      await loadLabFiles();
    }
  } catch (error) {
    showToast(String(error));
  }
});
</script>
