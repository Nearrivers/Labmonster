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
          <div class="flex items-center gap-x-1 pl-[14px] font-normal">
            <Image
              v-if="node.fileType === SupportedFiles.IMAGE"
              class="h-[14px] w-[14px] transition-transform"
            />
            <Clapperboard
              v-else-if="node.extension === '.mp4'"
              class="w-[14px] transition-transform"
            />
            <div
              role="textbox"
              ref="input"
              class="w-full cursor-pointer overflow-hidden text-ellipsis whitespace-nowrap bg-transparent [&_br]:hidden"
              :id="nodePathWithoutSpaces"
              @key.enter="input?.blur()"
              @blur.stop="onBlur"
              spellcheck="false"
              autocomplete="off"
            >
              {{ fileName }}
            </div>
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
import { Clapperboard, Image } from 'lucide-vue-next';
import { useFileNode } from '@/composables/Nodes/useFileNode';
import { toRef } from 'vue';
import { SupportedFiles } from '@/types/SupportedFiles';

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
