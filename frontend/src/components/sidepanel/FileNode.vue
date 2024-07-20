<template>
  <li class="w-full" :data-path="nodePath" data-type="file">
    <TooltipProvider>
      <Tooltip>
        <TooltipTrigger
          :class="
            cn(
              buttonVariants({ variant: 'ghost', size: 'sm' }),
              'h-7 w-full justify-start rounded-md',
            )
          "
        >
          <div class="flex items-center gap-x-1 font-normal">
            <p class="w-[14px]"></p>
            <p>
              {{ node.name }}
            </p>
          </div>
        </TooltipTrigger>
        <TooltipContent as-child :side="'right'" :side-offset="30">
          <p>Fichier</p>
        </TooltipContent>
      </Tooltip>
    </TooltipProvider>
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
import { ref } from 'vue';
import { cn } from '@/lib/utils';
import { buttonVariants } from '@/components/ui/button';

const props = defineProps<{
  node: filetree.Node;
  path: string;
}>();

const nodePath = ref(
  props.path ? props.path + '/' + props.node.name : props.node.name,
);
</script>
