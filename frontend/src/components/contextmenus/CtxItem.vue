<template>
  <li
    class="flex cursor-default items-center gap-2 rounded-sm px-2 py-1.5 hover:bg-muted"
  >
    <slot name="icon" :stroke-width="1.75" iconClass="h-[16px] w-4"></slot>
    <p><slot name="text"></slot></p>
  </li>
</template>

<script setup lang="ts"></script>
