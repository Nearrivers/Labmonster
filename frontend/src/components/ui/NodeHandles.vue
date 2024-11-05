<script setup lang="ts">
import { Handle, Position } from '@vue-flow/core';

defineProps<{
  nodeId: string;
}>();
</script>

<template>
  <Handle
    :id="nodeId + 'top'"
    type="source"
    :position="Position.Top"
    :class="'z-30 h-4 w-4 !cursor-pointer bg-accent-foreground opacity-0 hover:opacity-100'"
  />
  <Handle
    :id="nodeId + 'right'"
    type="source"
    :position="Position.Right"
    :class="'z-30 h-4 w-4 !cursor-pointer bg-accent-foreground opacity-0 hover:opacity-100'"
  />
  <Handle
    :id="nodeId + 'left'"
    type="source"
    :position="Position.Left"
    :class="'z-30 h-4 w-4 !cursor-pointer bg-accent-foreground opacity-0 hover:opacity-100'"
  />
  <Handle
    :id="nodeId + 'bot'"
    type="source"
    :position="Position.Bottom"
    :class="'z-30 h-4 w-4 !cursor-pointer bg-accent-foreground opacity-0 hover:opacity-100'"
  />
</template>
