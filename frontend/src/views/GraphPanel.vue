<template>
  <AppHeader />
  <RouterView v-slot="{ Component }">
    <component :is="Component" />
  </RouterView>
</template>

<script setup lang="ts">
import AppHeader from '@/components/AppHeader.vue';
import { RouterView } from 'vue-router';
</script>
