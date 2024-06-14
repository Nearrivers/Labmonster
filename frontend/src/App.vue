<template>
  <RouterView></RouterView>
</template>

<script lang="ts" setup>
import { useColorMode } from '@vueuse/core';

const mode = useColorMode();
</script>
