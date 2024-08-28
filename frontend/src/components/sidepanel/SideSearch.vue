<template>
  <header class="flex h-10 w-full items-center text-muted-foreground">
    <Search />
    <input
      type="text"
      name="sideSearch"
      id="sideSearch"
      class="border-b border-border bg-transparent"
    />
  </header>
</template>

<script setup lang="ts">
import { Search } from 'lucide-vue-next';
</script>
