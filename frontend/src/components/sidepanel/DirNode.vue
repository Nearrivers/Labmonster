<template>
  <li class="w-full" :data-path="nodePath" data-type="directory">
    <TooltipProvider>
      <Tooltip>
        <TooltipTrigger
          class="h-7 w-full justify-start rounded-md hover:bg-accent hover:text-accent-foreground"
          @click="toggle"
        >
          <div class="flex items-center gap-x-1 pl-[14px] font-normal">
            <ChevronRight
              v-if="isFolder"
              class="w-[14px] transition-transform"
              :class="{ 'rotate-90': isOpen }"
            />
            <p>
              {{ node.name }}
            </p>
          </div>
        </TooltipTrigger>
        <TooltipContent :side="'right'" :side-offset="30">
          <p>Dossier</p>
        </TooltipContent>
      </Tooltip>
    </TooltipProvider>
    <ul v-show="isOpen" class="w-full pl-[18.5px]">
      <template v-for="(child, index) in files" :key="nodePath + child.name">
        <FileNode
          v-if="child.type == 'FILE'"
          :node="child"
          :path="nodePath"
          :data-id="index"
        />
        <DirNode v-else :node="child" :path="nodePath" :data-id="index" />
      </template>
    </ul>
  </li>
</template>

<script setup lang="ts">
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from '@/components/ui/tooltip';
import { filetree } from '$/models';
import { toRef } from 'vue';
import { ChevronRight } from 'lucide-vue-next';
import { cn } from '@/lib/utils';
import { buttonVariants } from '@/components/ui/button';
import FileNode from '@/components/sidepanel/FileNode.vue';
import DirNode from '@/components/sidepanel/DirNode.vue';
import { useDirNode } from '@/composables/Nodes/useDirNode';

const props = defineProps<{
  node: filetree.Node;
  path: string;
}>();

const { files, isOpen, isFolder, nodePath, toggle } = useDirNode(toRef(props));
</script>
