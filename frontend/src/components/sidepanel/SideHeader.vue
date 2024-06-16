<template>
  <header class="flex justify-center gap-[2px] py-2 text-muted-foreground">
    <TooltipProvider>
      <Tooltip>
        <TooltipTrigger class="rounded-md p-1.5 hover:bg-zinc-700">
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
</template>

<script setup lang="ts">
import { FilePlus2, FolderPlus } from 'lucide-vue-next';
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from '@/components/ui/tooltip';
</script>
