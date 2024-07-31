<template>
  <header class="flex justify-center gap-[2px] py-2 text-muted-foreground">
    <TooltipProvider>
      <Tooltip>
        <TooltipTrigger
          class="rounded-md p-1.5 hover:bg-zinc-700"
          @click="createNewFileAtRoot"
        >
          <FilePlus2 :stroke-width="1.75" class="h-[18px] w-[18px]" />
        </TooltipTrigger>
        <TooltipContent>
          <p>Créer un nouveau diagramme</p>
        </TooltipContent>
      </Tooltip>
    </TooltipProvider>
    <TooltipProvider>
      <Tooltip>
        <TooltipTrigger class="rounded-md p-1.5 hover:bg-zinc-700">
          <FolderPlus :stroke-width="1.75" class="h-[18px] w-[18px]" />
        </TooltipTrigger>
        <TooltipContent>
          <p>Créer un nouveau dossier</p>
        </TooltipContent>
      </Tooltip>
    </TooltipProvider>
  </header>
  <ScrollArea class="h-[95svh] pb-4" data-path="/">
    <ul
      class="w-full px-2 text-sm text-muted-foreground"
      v-if="files.length > 0"
      @contextmenu.prevent="onRightClick"
      @click="onLeftClick"
    >
      <template v-for="file in files" :key="file.name">
        <FileNode v-if="file.type === 'FILE'" :node="file" path="" />
        <DirNode v-if="file.type === 'DIR'" :node="file" path="" />
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
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from '@/components/ui/tooltip';
import { onMounted } from 'vue';
import { FilePlus2, FolderPlus } from 'lucide-vue-next';
import { ScrollArea } from '@/components/ui/scroll-area';
import FileNode from '@/components/sidepanel/FileNode.vue';
import DirNode from '@/components/sidepanel/DirNode.vue';
import { useSidePanel } from '@/composables/useSidePanel';
import { CheckConfigPresenceAndLoadIt } from '$/config/AppConfig';
import FileContextMenu from '@/components/contextmenus/FileContextMenu.vue';
import DirContextMenu from '@/components/contextmenus/DirContextMenu.vue';

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
