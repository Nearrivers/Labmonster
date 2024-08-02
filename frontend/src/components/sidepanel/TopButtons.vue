<template>
  <TopButton @click="emit('createFile')">
    <template #icon>
      <FilePlus2 :stroke-width="1.75" class="h-[18px] w-[18px]" />
    </template>
    <template #tooltip>Créer un nouveau setup</template>
  </TopButton>
  <TopButton @click="emit('createFolder')">
    <template #icon>
      <FolderPlus :stroke-width="1.75" class="h-[18px] w-[18px]" />
    </template>
    <template #tooltip>Créer un nouveau dossier</template>
  </TopButton>
</template>

<script setup lang="ts">
import TopButton from '@/components/ui/TopButton.vue';
import { FilePlus2, FolderPlus } from 'lucide-vue-next';

const emit = defineEmits<{
  (e: 'createFile'): void;
  (e: 'createFolder'): void;
}>();
</script>
