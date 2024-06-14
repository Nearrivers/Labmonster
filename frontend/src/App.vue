<template>
  <RouterView />
  <!-- <HelloWorld />
  <Button @click="QuitApp">Quitter l'application</Button> -->
</template>

<script lang="ts" setup>
import { useColorMode } from '@vueuse/core';
import HelloWorld from './components/HelloWorld.vue';
import { Button } from '@/components/ui/button';
import { QuitApp } from '../wailsjs/go/topmenu/TopMenu';

const mode = useColorMode();
</script>
