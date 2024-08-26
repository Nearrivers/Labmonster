<template>
  <TooltipProvider>
    <Tooltip>
      <TooltipTrigger
        class="rounded-md"
        :class="
          cn(
            buttonVariants({ variant: 'ghost', size: 'sm' }),
            'h-min rounded-md p-1.5',
            additionnalClasses,
          )
        "
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
