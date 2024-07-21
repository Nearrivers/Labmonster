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
            <input
              ref="input"
              type="text"
              v-model="fileName"
              class="cursor-pointer bg-transparent"
              disabled
              :id="nodePathWithoutSpaces"
              @blur="onBlur"
              autocomplete="off"
            />
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
import { computed, ref } from 'vue';
import { cn } from '@/lib/utils';
import { buttonVariants } from '@/components/ui/button';
import { RenameFile } from '$/filetree/FileTreeExplorer';

const props = defineProps<{
  node: filetree.Node;
  path: string;
}>();

const fileName = ref(props.node.name);
const input = ref<HTMLInputElement | null>(null);
const nodePath = ref(
  props.path ? props.path + '/' + props.node.name : props.node.name,
);

const nodePathWithoutSpaces = computed(() =>
  nodePath.value.replaceAll(' ', '-'),
);

async function onBlur() {
  if (input.value) {
    input.value.toggleAttribute('disabled');
    input.value.classList.add('cursor-pointer');
    input.value.classList.remove('cursor-text');
  }

  console.log(props.path, fileName.value, props.node.name);
  try {
    await RenameFile(props.path, props.node.name, fileName.value);
  } catch (error) {
    console.log(error);
  }
}
</script>
