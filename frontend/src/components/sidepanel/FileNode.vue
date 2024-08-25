<template>
  <li
    class="w-full"
    :data-path="nodePath"
    data-type="file"
    :data-extension="node.extension"
    :data-file="node.fileType"
  >
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
          <div class="flex w-full items-center gap-x-1 pl-[14px] font-normal">
            <NodeIcon :fileType="props.node.fileType" />
            <input
              role="textbox"
              ref="input"
              class="w-full cursor-pointer overflow-hidden text-ellipsis whitespace-nowrap bg-transparent outline-none"
              :id="nodePathWithoutSpaces"
              @key.enter.prevent="input?.blur()"
              @blur.stop="onBlur"
              spellcheck="false"
              autocomplete="off"
              v-model="fileName"
              readonly
            />
          </div>
        </TooltipTrigger>
        <TooltipContent as-child :side="'right'" :side-offset="20">
          <div>
            <p class="text-xs">Dernière modification le: {{ updatedAt }}</p>
            <p class="text-center text-xs">{{ ext }}</p>
          </div>
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
import { cn } from '@/lib/utils';
import { buttonVariants } from '@/components/ui/button';
import { useFileNode } from '@/composables/Nodes/useFileNode';
import { toRef } from 'vue';
import NodeIcon from './NodeIcon.vue';

const props = defineProps<{
  node: filetree.Node;
  path: string;
}>();

const {
  nodePath,
  ext,
  fileName,
  input,
  nodePathWithoutSpaces,
  onBlur,
  updatedAt,
} = useFileNode(toRef(props));
</script>
