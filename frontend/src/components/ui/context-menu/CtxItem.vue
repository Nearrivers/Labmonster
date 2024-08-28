<template>
  <li
    class="relative flex cursor-default items-center justify-between rounded-sm px-2 py-1.5 hover:bg-accent"
  >
    <div class="flex items-center gap-2">
      <slot name="icon" :stroke-width="1.75" iconClass="h-[16px] w-4"></slot>
      <p><slot name="text"></slot></p>
    </div>
    <ChevronRight class="h-[16px] w-4" v-if="withSubmenu" />
    <CtxSubMenu v-if="withSubmenu">
      <slot name="submenus"></slot>
    </CtxSubMenu>
  </li>
</template>

<script setup lang="ts">
import { ChevronRight } from 'lucide-vue-next';
import CtxSubMenu from './CtxSubMenu.vue';

defineProps<{
  withSubmenu?: boolean;
}>();
</script>
