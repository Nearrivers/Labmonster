<template></template>

<script lang="ts" setup>
import { onMounted, ref } from 'vue';
import { GetFileTree } from '$/filetree/FileTreeExplorer';

const fileTree = ref<{ [key: string]: Array<any> }>({ '': [] });

onMounted(async () => {
  try {
    fileTree.value = await GetFileTree();
    console.log(fileTree.value);
  } catch (error) {}
});
</script>
