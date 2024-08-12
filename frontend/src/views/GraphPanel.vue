<template>
  <AppHeader />
  <RouterView v-slot="{ Component }">
    <KeepAlive>
      <component :is="Component" />
    </KeepAlive>
  </RouterView>
</template>

<script setup lang="ts">
import AppHeader from '@/components/AppHeader.vue';
import { RouterView } from 'vue-router';
</script>
