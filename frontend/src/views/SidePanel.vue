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
  <ScrollArea class="h-[95svh] pb-4">
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
import { FilePlus2, FolderPlus } from 'lucide-vue-next';
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from '@/components/ui/tooltip';
import { filetree } from '$/models';
import { ref, onMounted, nextTick } from 'vue';
import { ScrollArea } from '@/components/ui/scroll-area';
import FileNode from '@/components/sidepanel/FileNode.vue';
import DirNode from '@/components/sidepanel/DirNode.vue';
import FileContextMenu from '@/components/contextmenus/FileContextMenu.vue';
import DirContextMenu from '@/components/contextmenus/DirContextMenu.vue';
import { NEW_FILE_NAME } from '@/constants/NEW_FILE_NAME';
import {
  CreateNewFileAtRoot,
  GetSubDirAndFiles,
} from '$/filetree/FileTreeExplorer';
import { useSidePanel } from '@/composables/useSidePanel';
import { useShowErrorToast } from '@/composables/useShowErrorToast';
import { CheckConfigPresenceAndLoadIt } from '$/config/AppConfig';
import { useEventListener } from '@/composables/useEventListener';
import { configFileLoaded } from '@/events/ReloadFileExplorer';
import { CONFIG_FILE_LOADED } from '@/constants/event-names/CONFIG_FILE_LOADED';

const files = ref<filetree.Node[]>([]);
const { sortNodes } = useSidePanel();
const { showToast } = useShowErrorToast();
const contextMenuX = ref(100);
const contextMenuY = ref(100);
const fileContextMenu = ref<InstanceType<typeof FileContextMenu> | null>(null);
const selectedNode = ref<HTMLLIElement | null>(null);
useEventListener(configFileLoaded, CONFIG_FILE_LOADED, loadLabFiles);

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

async function loadLabFiles() {
  try {
    files.value = await GetSubDirAndFiles('');
  } catch (error) {
    showToast(String(error), 'Impossible de charger les fichiers');
  }
}

async function createNewFileAtRoot() {
  try {
    const newFileName = await CreateNewFileAtRoot(NEW_FILE_NAME);
    files.value.push(newFileName);
    files.value.sort(sortNodes);
  } catch (error) {
    showToast(String(error), 'Impossible de créer le fichier');
  }
}

async function onRightClick(event: MouseEvent) {
  contextMenuX.value = event.clientX;
  contextMenuY.value = event.clientY;
  selectedNode.value = (event.target as HTMLElement).closest('li');
  await nextTick();
  fileContextMenu.value?.showPopover();
}

function onLeftClick(event: MouseEvent) {}
</script>
