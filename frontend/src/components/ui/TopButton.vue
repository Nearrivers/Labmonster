<template>
  <TooltipProvider>
    <Tooltip>
      <TooltipTrigger
        class="h-min rounded-md p-1.5 hover:bg-accent hover:text-accent-foreground"
        :class="additionnalClasses"
        @click="emit('click')"
      >
        <slot name="icon" />
      </TooltipTrigger>
      <TooltipContent :side="tooltipSide">
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
    additionnalClasses?: string;
    tooltipSide?: Side;
  }>(),
  {
    tooltipSide: 'top',
  },
);

const emit = defineEmits<{
  (e: 'click'): void;
}>();
</script>
