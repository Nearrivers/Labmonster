<template>
  <Image v-if="fileType === SupportedFiles.IMAGE" class="h-[14px] w-[14px]" />
  <Clapperboard
    v-else-if="fileType === SupportedFiles.VIDEO"
    class="h-[14px] w-[14px]"
  />
  <Workflow
    v-else-if="fileType === SupportedFiles.GRAPH"
    class="h-[14px] w-[14px]"
  />
  <FileX2 v-else class="h-[14px] w-[14px]" />
</template>

<script setup lang="ts">
import { SupportedFiles } from '@/types/SupportedFiles';
import { Clapperboard, Image, Workflow, FileX2 } from 'lucide-vue-next';

defineProps<{
  fileType: SupportedFiles;
}>();
</script>
