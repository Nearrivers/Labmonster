<template>
  <HelloWorld />
</template>

<script lang="ts" setup>
import { useColorMode } from '@vueuse/core';
import HelloWorld from './components/HelloWorld.vue';

const mode = useColorMode();
</script>
