<template>
  <ContextMenu>
    <ContextMenuTrigger as-child>
      <li class="h-7 w-full">
        <TooltipProvider>
          <Tooltip>
            <TooltipTrigger
              :class="
                cn(
                  buttonVariants({ variant: 'ghost', size: 'sm' }),
                  'h-full w-full justify-start rounded-md',
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
        </ul></li
    ></ContextMenuTrigger>
    <ContextMenuContent class="w-64">
      <ContextMenuItem inset>
        Back
        <ContextMenuShortcut>⌘[</ContextMenuShortcut>
      </ContextMenuItem>
      <ContextMenuItem inset disabled>
        Forward
        <ContextMenuShortcut>⌘]</ContextMenuShortcut>
      </ContextMenuItem>
      <ContextMenuItem inset>
        Reload
        <ContextMenuShortcut>⌘R</ContextMenuShortcut>
      </ContextMenuItem>
      <ContextMenuSub>
        <ContextMenuSubTrigger inset> More Tools </ContextMenuSubTrigger>
        <ContextMenuSubContent class="w-48">
          <ContextMenuItem>
            Save Page As...
            <ContextMenuShortcut>⇧⌘S</ContextMenuShortcut>
          </ContextMenuItem>
          <ContextMenuItem>Create Shortcut...</ContextMenuItem>
          <ContextMenuItem>Name Window...</ContextMenuItem>
          <ContextMenuSeparator />
          <ContextMenuItem>Developer Tools</ContextMenuItem>
        </ContextMenuSubContent>
      </ContextMenuSub>
      <ContextMenuSeparator />
      <ContextMenuCheckboxItem checked>
        Show Bookmarks Bar
        <ContextMenuShortcut>⌘⇧B</ContextMenuShortcut>
      </ContextMenuCheckboxItem>
      <ContextMenuCheckboxItem>Show Full URLs</ContextMenuCheckboxItem>
      <ContextMenuSeparator />
      <ContextMenuRadioGroup model-value="pedro">
        <ContextMenuLabel inset> People </ContextMenuLabel>
        <ContextMenuSeparator />
        <ContextMenuRadioItem value="pedro">
          Pedro Duarte
        </ContextMenuRadioItem>
        <ContextMenuRadioItem value="colm"> Colm Tuite </ContextMenuRadioItem>
      </ContextMenuRadioGroup>
    </ContextMenuContent>
  </ContextMenu>
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
import {
  ContextMenu,
  ContextMenuCheckboxItem,
  ContextMenuContent,
  ContextMenuItem,
  ContextMenuLabel,
  ContextMenuRadioGroup,
  ContextMenuRadioItem,
  ContextMenuSeparator,
  ContextMenuShortcut,
  ContextMenuSub,
  ContextMenuSubContent,
  ContextMenuSubTrigger,
  ContextMenuTrigger,
} from '@/components/ui/context-menu';
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
