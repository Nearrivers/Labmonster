<template>
  <AppHeader />
  <RouterView v-slot="{ Component }">
    <KeepAlive>
      <component :is="Component" />
    </KeepAlive>
  </RouterView>
  <!-- <FlowChart /> -->
</template>

<script setup lang="ts">
import AppHeader from '@/components/AppHeader.vue';
import FlowChart from '@/components/FlowChart.vue';
import { RouterView } from 'vue-router';
</script>
