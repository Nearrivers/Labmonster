<template>
  <TooltipProvider>
    <Tooltip>
      <TooltipTrigger
        class="rounded-sm p-2 hover:bg-secondary hover:text-primary"
        :class="class"
        @click="emit('click')"
      >
        <slot name="icon" />
      </TooltipTrigger>
      <TooltipContent :side-offset="sideOffset" :side="side">
        <p><slot name="tooltip" /></p>
      </TooltipContent>
    </Tooltip>
  </TooltipProvider>
</template>

<script lang="ts" setup>
import {
  Tooltip,
  TooltipTrigger,
  TooltipContent,
  TooltipProvider,
} from '@/components/ui/tooltip';
import { Side } from 'radix-vue/dist/Popper';

withDefaults(
  defineProps<{
    class?: string;
    side?: Side;
    sideOffset?: number;
  }>(),
  {
    sideOffset: 15,
  },
);

const emit = defineEmits<{
  (e: 'click'): void;
}>();
</script>
