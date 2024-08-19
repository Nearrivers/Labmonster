<template>
  <Transition
    enter-active-class="transition-all duration-200"
    enter-from-class="scale-75"
    leave-active-class="transition-all duration-200"
    leave-to-class="scale-75"
  >
    <ul
      class="rouded-lg absolute right-0 z-20 origin-top-left border border-border"
      v-if="isDisplayed"
    >
      <slot></slot>
    </ul>
  </Transition>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const isDisplayed = ref(false);

function showSubmenu() {
  isDisplayed.value = true;
}

function hideSubmenu() {
  isDisplayed.value = false;
}

defineExpose({
  showSubmenu,
  hideSubmenu,
});
</script>
