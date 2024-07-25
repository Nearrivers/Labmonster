<template>
  <li class="w-full" :data-path="nodePath" data-type="directory">
    <TooltipProvider>
      <Tooltip>
        <TooltipTrigger
          :class="
            cn(
              buttonVariants({ variant: 'ghost', size: 'sm' }),
              'h-7 w-full justify-start rounded-md',
            )
          "
          @click="toggle"
        >
          <div class="flex items-center gap-x-1 font-normal">
            <ChevronRight
              v-if="isFolder"
              class="w-[14px] transition-transform"
              :class="{ 'rotate-90': isOpen }"
            />
            <p v-else class="w-[14px]"></p>
            <p>
              {{ node.name }}
            </p>
          </div>
        </TooltipTrigger>
        <TooltipContent :side="'right'" :side-offset="30">
          <p>Dossier</p>
          <TooltipArrow />
        </TooltipContent>
      </Tooltip>
    </TooltipProvider>
    <ul v-if="isOpen" class="w-full pl-[18.5px]">
      <template v-for="child in files" :key="child.name">
        <FileNode v-if="child.type == 'FILE'" :node="child" :path="nodePath" />
        <DirNode v-else :node="child" :path="nodePath" />
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
import { computed, ref } from 'vue';
import { ChevronRight } from 'lucide-vue-next';
import { cn } from '@/lib/utils';
import { buttonVariants } from '@/components/ui/button';
import { GetSubDirAndFiles } from '$/filetree/FileTreeExplorer';
import FileNode from '@/components/sidepanel/FileNode.vue';
import DirNode from '@/components/sidepanel/DirNode.vue';
import { TooltipArrow } from 'radix-vue';
import { useShowErrorToast } from '@/composables/useShowErrorToast';

const props = defineProps<{
  node: filetree.Node;
  path: string;
}>();

const files = ref<filetree.Node[]>([]);
const isOpen = ref(false);
const isFolder = computed(() => props.node.type === 'DIR');
const nodePath = ref(
  props.path ? props.path + '/' + props.node.name : props.node.name,
);
const { showToast } = useShowErrorToast();

async function toggle() {
  try {
    files.value = props.path
      ? await GetSubDirAndFiles(props.path + '/' + props.node.name)
      : await GetSubDirAndFiles(props.node.name);
  } catch (error) {
    showToast(String(error));
  } finally {
    isOpen.value = !isOpen.value;
  }
}
</script>
