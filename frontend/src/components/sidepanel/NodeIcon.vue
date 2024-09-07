<template>
  <Image v-if="fileType === node.FileType.IMAGE" class="h-[14px] w-[14px]" />
  <Clapperboard
    v-else-if="fileType === node.FileType.VIDEO"
    class="h-[14px] w-[14px]"
  />
  <Workflow
    v-else-if="fileType === node.FileType.GRAPH"
    class="h-[14px] w-[14px]"
  />
  <FileX2 v-else class="h-[14px] w-[14px]" />
</template>

<script setup lang="ts">
import { node } from '$/models';
import { Clapperboard, Image, Workflow, FileX2 } from 'lucide-vue-next';

defineProps<{
  fileType: node.FileType;
}>();
</script>
