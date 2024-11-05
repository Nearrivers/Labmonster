<template>
  <li
    class="relative mb-0.5 w-full"
    :data-path="nodePath"
    data-type="file"
    :data-extension="fileNode.extension"
    :data-file="fileNode.fileType"
    @keyup.stop="selectInput"
  >
    <TooltipProvider>
      <Tooltip>
        <TooltipTrigger
          class="h-7 w-full cursor-default justify-start rounded-md hover:bg-accent hover:text-accent-foreground"
          :class="{ 'bg-accent text-accent-foreground': isActive }"
        >
          <div
            class="flex w-full items-center gap-x-1 font-normal"
            :style="nodeStyle"
          >
            <NodeIcon :fileType="props.fileNode.fileType" />
            <input
              role="textbox"
              ref="input"
              class="w-full cursor-default overflow-hidden text-ellipsis whitespace-nowrap bg-transparent outline-none"
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
import { node } from '$/models';
import { useFileNode } from '@/composables/Nodes/useFileNode';
import { toRef } from 'vue';
import NodeIcon from './NodeIcon.vue';

const props = defineProps<{
  fileNode: node.Node;
  path: string;
  offset: number;
}>();

const {
  nodePath,
  nodeStyle,
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
