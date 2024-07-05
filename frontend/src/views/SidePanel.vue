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
      <FileNode
        v-for="(file, index) in files"
        :node="file"
        :key="index"
        path=""
      />
    </ul>
  </ScrollArea>
  <FileContextMenu
    ref="fileContextMenu"
    :x="contextMenuX"
    :y="contextMenuY"
    :selected-node="selectedNode"
  />
  <FolderContextMenu />
</template>

<script setup lang="ts">
import { FilePlus2, FolderPlus } from 'lucide-vue-next';
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from '@/components/ui/tooltip';
import { CreateNewFile } from '$/filetree/FileTreeExplorer';
import { filetree } from '$/models';
import { ref, onMounted, nextTick } from 'vue';
import { GetFirstDepth } from '$/filetree/FileTreeExplorer';
import { ScrollArea } from '@/components/ui/scroll-area';
import FileNode from '@/components/sidepanel/FileNode.vue';
import FileContextMenu from '@/components/contextmenus/FileContextMenu.vue';
import FolderContextMenu from '@/components/contextmenus/FolderContextMenu.vue';

const files = ref<filetree.Node[]>([]);
const contextMenuX = ref(100);
const contextMenuY = ref(100);
const fileContextMenu = ref<InstanceType<typeof FileContextMenu> | null>(null);
const selectedNode = ref<HTMLLIElement | null>(null);

onMounted(async () => {
  try {
    files.value = await GetFirstDepth();
  } catch (error) {}
});

async function createNewFileAtRoot() {
  try {
    const newFileName = await CreateNewFile();
    files.value.push(
      new filetree.Node({
        name: newFileName,
        type: 'FILE',
        files: [],
      }),
    );
    files.value.sort((f1, f2) => {
      // Tri sur les types d'abord
      if (f1.type === 'DIR' && f2.type == 'FILE') {
        return 1;
      }

      if (f1.type === 'FILE' && f2.type == 'DIR') {
        return 1;
      }

      // Si les types sont les même, on trie sur le nom.
      // La fonction sort() sans fonction de comparaison custom trie les chaînes de caractères
      // dans l'ordre ASCII des caractères ce qui n'est pas le cas de ma fonction de tri en Go.
      // Cela causait une différence entre le tri réalisé par le backend et celui fait par le front
      // d'où l'implémentation de cette fonction de sort()
      if (f1.name < f2.name) {
        return -1;
      }

      if (f1.name == f2.name) {
        return 0;
      }

      if (f1.name > f2.name) {
        return 1;
      }

      return 0;
    });
  } catch (error) {
    console.log(error);
  }
}

async function onRightClick(event: MouseEvent) {
  contextMenuX.value = event.clientX;
  contextMenuY.value = event.clientY;
  console.log((event.target as HTMLElement).closest('li'));
  selectedNode.value = (event.target as HTMLElement).closest('li');
  await nextTick();
  fileContextMenu.value?.showPopover();
}

function onLeftClick(event: MouseEvent) {}
</script>
