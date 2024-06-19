<template>
  <ScrollArea class="h-[95svh]">
    <ul
      class="w-full px-2 text-sm text-muted-foreground"
      v-if="fileTree.length > 0"
    >
      <FileNode v-for="(file, index) in fileTree" :node="file" :key="index" />
    </ul>
    <ul class="w-full px-2 text-sm text-muted-foreground" v-else>
      <li class="w-full bg-red-50" v-for="i in [1, 2, 3]" :key="i"></li>
    </ul>
  </ScrollArea>
</template>

<script setup lang="ts">
import { GetFileTree } from '$/filetree/FileTreeExplorer';
import { filetree } from '$/models';
import { onMounted, ref } from 'vue';
import FileNode from './FileNode.vue';
import { ScrollArea } from '@/components/ui/scroll-area';

const fileTree = ref<filetree.Node[]>([]);

onMounted(async () => {
  try {
    fileTree.value = await GetFileTree();
  } catch (error) {}
});
</script>
