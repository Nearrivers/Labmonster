<template>
  <TooltipProvider>
    <Tooltip>
      <TooltipTrigger as-child>
        <button
          class="block bg-secondary p-2 hover:bg-accent"
          @click="emit('click')"
        >
          <slot />
        </button>
      </TooltipTrigger>
      <TooltipContent :side="'left'">
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
