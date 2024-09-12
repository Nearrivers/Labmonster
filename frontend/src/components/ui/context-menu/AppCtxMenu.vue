<template>
  <ul
    :id="$props.popoverId"
    popover
    ref="menu"
    class="fixed z-10 min-w-52 origin-top-left animate-command-show rounded-lg border border-accent bg-secondary text-xs dark:text-primary"
    :style="{ top: y + 'px', left: x + 'px' }"
  >
    <slot></slot>
    <slot name="commands"></slot>
  </ul>
</template>

<script lang="ts" setup>
import { ContextMenuProps } from '@/types/props/ContextMenuProps';
import { ref } from 'vue';

defineProps<ContextMenuProps>();

const menu = ref<any | null>(null);

function showPopover() {
  menu.value?.showPopover();
}

function hidePopover() {
  menu.value?.hidePopover();
}

defineExpose({
  showPopover,
  hidePopover,
});
</script>
