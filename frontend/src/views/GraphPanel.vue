<template>
  <div class="flex h-full w-full flex-col">
    <AppHeader />
    <div class="h-full w-full flex-grow bg-secondary p-2 pl-1 pt-0">
      <div class="h-full w-full rounded-md bg-background">
        <RouterView v-slot="{ Component }">
          <component :is="Component" />
        </RouterView>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import AppHeader from '@/components/AppHeader.vue';
import { RouterView } from 'vue-router';
</script>
