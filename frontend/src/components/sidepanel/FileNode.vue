<template>
  <li class="w-full border-l border-muted-foreground">
    <TooltipProvider>
      <Tooltip>
        <TooltipTrigger
          class="mb-1 flex items-center rounded-md px-2 hover:bg-zinc-700"
        >
          <div
            :class="{ bold: isFolder }"
            class="flex items-center gap-1"
            @click="toggle"
          >
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
        <TooltipContent>
          <p>Créer un nouveau diagramme</p>
        </TooltipContent>
      </Tooltip>
    </TooltipProvider>
    <ul v-show="isOpen" v-if="isFolder" class="w-full px-3.5">
      <FileNode
        class="item"
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

const props = defineProps<{
  node: filetree.Node;
}>();

const isOpen = ref(false);
const isFolder = computed(() => props.node.files && props.node.files.length);

function toggle() {
  isOpen.value = !isOpen.value;
}
</script>
