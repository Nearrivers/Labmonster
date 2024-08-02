<template>
  <TooltipProvider>
    <Tooltip>
      <TooltipTrigger as-child>
        <button
          class="p-2 hover:bg-secondary dark:hover:bg-zinc-700"
          @click="emit('click')"
        >
          <slot />
        </button>
      </TooltipTrigger>
      <TooltipContent>
        <p><slot name="tooltip" /></p>
      </TooltipContent>
    </Tooltip>
  </TooltipProvider>
</template>

<script setup lang="ts">
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from '@/components/ui/tooltip';

const emit = defineEmits<{
  (e: 'click'): void;
}>();
</script>
