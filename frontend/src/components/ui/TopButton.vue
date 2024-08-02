<template>
  <TooltipProvider>
    <Tooltip>
      <TooltipTrigger
        class="rounded-md p-1.5 hover:bg-zinc-700 hover:text-primary"
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

const emit = defineEmits<{
  (e: 'click'): void;
}>();
</script>
