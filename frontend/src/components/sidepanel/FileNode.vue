<template>
  <li
    class="mb-0.5 w-full"
    :data-path="nodePath"
    data-type="file"
    :data-extension="node.extension"
    :data-file="node.fileType"
    @keyup.stop="selectInput"
  >
    <TooltipProvider>
      <Tooltip>
        <TooltipTrigger
          class="h-7 w-full justify-start rounded-md hover:bg-accent hover:text-accent-foreground"
          :class="{ 'bg-accent text-accent-foreground': isActive }"
        >
          <div class="flex w-full items-center gap-x-1 pl-[14px] font-normal">
            <NodeIcon :fileType="props.node.fileType" />
            <input
              role="textbox"
              ref="input"
              class="w-full cursor-pointer overflow-hidden text-ellipsis whitespace-nowrap bg-transparent outline-none"
              :id="nodePathWithoutSpaces"
              @blur.stop="onBlur"
              @keyup.enter="input?.blur()"
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
  selectInput,
  isActive,
} = useFileNode(toRef(props));
</script>
