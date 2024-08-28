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
      <TooltipContent>
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
import { cn } from '@/lib/utils';
import { buttonVariants } from '@/components/ui/button';

defineProps<{
  additionnalClasses?: string;
}>();

const emit = defineEmits<{
  (e: 'click'): void;
}>();
</script>
