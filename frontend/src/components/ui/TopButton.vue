<template>
  <TooltipProvider>
    <Tooltip>
      <TooltipTrigger
        class="rounded-md p-1.5 hover:bg-secondary hover:bg-zinc-700 hover:text-primary"
        :class="class"
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

defineProps<{
  class?: string;
}>();

const emit = defineEmits<{
  (e: 'click'): void;
}>();
</script>
