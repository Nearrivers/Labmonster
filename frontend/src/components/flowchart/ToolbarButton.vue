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
      <TooltipContent :side-offset="15">
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
