<template>
  <ul class="w-full px-2 text-sm text-muted-foreground [&>li]:!border-none">
    <FileNode v-for="(file, index) in fileTree" :node="file" :key="index" />
  </ul>
</template>

<script setup lang="ts">
import { GetFileTree } from '$/filetree/FileTreeExplorer';
import { filetree } from '$/models';
import { onMounted, ref } from 'vue';
import FileNode from './FileNode.vue';

const fileTree = ref<filetree.Node[]>([]);

onMounted(async () => {
  try {
    fileTree.value = await GetFileTree();
  } catch (error) {}
});
</script>
