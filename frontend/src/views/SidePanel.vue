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
    >
      <FileNode v-for="(file, index) in files" :node="file" :key="index" />
    </ul>
    <ul class="w-full px-2 text-sm text-muted-foreground" v-else>
      <li class="w-full bg-red-50" v-for="i in [1, 2, 3]" :key="i"></li>
    </ul>
  </ScrollArea>
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
import { ref, onMounted } from 'vue';
import { GetFirstDepth } from '$/filetree/FileTreeExplorer';
import { ScrollArea } from '@/components/ui/scroll-area';
import FileNode from '@/components/sidepanel/FileNode.vue';

const files = ref<filetree.Node[]>([]);

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
      if (f1.type === 'DIR' && f2.type == 'FILE') {
        return 1;
      }

      if (f1.type === 'FILE' && f2.type == 'DIR') {
        return 1;
      }

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
</script>
