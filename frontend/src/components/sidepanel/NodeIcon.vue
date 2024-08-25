<template>
  <Image
    v-if="fileType === filetree.FileType.IMAGE"
    class="h-[14px] w-[14px]"
  />
  <Clapperboard
    v-else-if="fileType === filetree.FileType.VIDEO"
    class="h-[14px] w-[14px]"
  />
  <Workflow
    v-else-if="fileType === filetree.FileType.GRAPH"
    class="h-[14px] w-[14px]"
  />
  <FileX2 v-else class="h-[14px] w-[14px]" />
</template>

<script setup lang="ts">
import { filetree } from '$/models';
import { Clapperboard, Image, Workflow, FileX2 } from 'lucide-vue-next';

defineProps<{
  fileType: filetree.FileType;
}>();
</script>
