<template>
  <li
    class="w-full"
    :data-path="nodePath"
    data-type="file"
    :data-extension="node.extension"
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
          <div class="flex items-center gap-x-1 font-normal">
            <p class="w-[14px]"></p>
            <div
              role="textbox"
              ref="input"
              class="cursor-pointer overflow-hidden whitespace-nowrap bg-transparent [&_br]:hidden"
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
import { computed, ref } from 'vue';
import { cn } from '@/lib/utils';
import { buttonVariants } from '@/components/ui/button';
import { RenameFile } from '$/filetree/FileTree';
import { useShowErrorToast } from '@/composables/useShowErrorToast';

const props = defineProps<{
  node: filetree.Node;
  path: string;
}>();

const { showToast } = useShowErrorToast();
const fileName = ref(props.node.name);
const input = ref<HTMLDivElement | null>(null);
const nodePath = ref(
  props.path ? props.path + '/' + props.node.name : props.node.name,
);

const nodePathWithoutSpaces = computed(() =>
  nodePath.value.replaceAll(' ', '-'),
);

const updatedAt = computed(() => {
  const date = new Date(props.node.updatedAt);
  return `${date.toLocaleDateString()} à ${date.toLocaleTimeString()}`;
});

async function onBlur() {
  if (!input.value) {
    showToast('Input introuvable');
    return;
  }

  input.value.toggleAttribute('contenteditable');
  input.value.classList.add('cursor-pointer');
  input.value.classList.remove('cursor-text');
  fileName.value = input.value.innerText.trim();

  try {
    await RenameFile(
      props.path,
      props.node.name + props.node.extension,
      fileName.value + props.node.extension,
    );
  } catch (error) {
    showToast(error);
  }
}
</script>
