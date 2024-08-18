<template>
  <AppCtxMenu :y="y" :x="x" popover-id="flowpopover">
    <CtxSection>
      <CtxItem>
        <template #icon><DiamondPlus :stroke-width="1.75" /></template>
      </CtxItem>
    </CtxSection>
  </AppCtxMenu>
</template>

<script setup lang="ts">
import { ContextMenuProps } from '@/types/props/ContextMenuProps';
import AppCtxMenu from './AppCtxMenu.vue';
import CtxSection from './CtxSection.vue';
import CtxItem from './CtxItem.vue';
import { DiamondPlus } from 'lucide-vue-next';

const props = defineProps<ContextMenuProps>();
</script>
