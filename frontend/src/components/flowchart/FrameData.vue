<template>
  <section class="flex justify-between border-t border-t-border p-2 text-sm">
    <p>Frame data:</p>
    <input
      type="number"
      class="w-10 bg-transparent text-right font-bold text-primary outline-none"
      :class="{
        'text-red-400': frameData < 0,
        '!text-blue-400': frameData > 0,
      }"
      name="frameData"
      v-model="frameData"
    />
  </section>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const frameData = ref(0);
</script>
