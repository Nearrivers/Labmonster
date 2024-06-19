<template>
  <li class="h-min w-full">
    <TooltipProvider>
      <Tooltip>
        <TooltipTrigger
          :class="
            cn(
              buttonVariants({ variant: 'ghost', size: 'sm' }),
              'h-6 w-full justify-start py-0',
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
            <p>
              {{ node.name }}
            </p>
          </div>
        </TooltipTrigger>
        <TooltipContent :side="'right'">
          <p v-if="node.type === 'DIR'">Dossier</p>
          <p v-else>Fichier</p>
        </TooltipContent>
      </Tooltip>
    </TooltipProvider>
    <ul v-show="isOpen" v-if="isFolder" class="w-full pl-[18.5px]">
      <FileNode
        v-for="(child, index) in node.files"
        :key="index"
        :node="child"
      ></FileNode>
    </ul>
  </li>
</template>

<script setup lang="ts">
import { filetree } from '$/models';
import { computed, ref } from 'vue';
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from '@/components/ui/tooltip';
import { ChevronRight } from 'lucide-vue-next';
import { cn } from '@/lib/utils';
import { buttonVariants } from '@/components/ui/button';

const props = defineProps<{
  node: filetree.Node;
}>();

const isOpen = ref(false);
const isFolder = computed(() => props.node.type === 'DIR');

function toggle() {
  isOpen.value = !isOpen.value;
}
</script>
