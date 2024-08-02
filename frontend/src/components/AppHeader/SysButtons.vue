<template>
  <section class="ml-auto">
    <button class="p-2 hover:bg-secondary dark:hover:bg-zinc-700">
      <Minus class="h-4 w-6" />
    </button>
    <button class="p-2 hover:bg-secondary dark:hover:bg-zinc-700">
      <Maximize class="h-4 w-6" />
    </button>
    <button class="p-2 hover:bg-red-500">
      <X class="h-4 w-6" />
    </button>
  </section>
</template>

<script setup lang="ts">
import { Maximize, Minus, X } from 'lucide-vue-next';
</script>
