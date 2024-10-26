<template>
  <div class="flex items-end justify-between border-b border-b-border pb-2">
    <h2 class="font-medium">
      <slot></slot>
    </h2>
    <slot name="actions"></slot>
  </div>
</template>

<script lang="ts" setup></script>
