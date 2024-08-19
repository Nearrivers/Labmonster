<template>
  <AppCtxMenu :y="y" :x="x" popover-id="flowpopover" ref="ctxMenu">
    <CtxSection>
      <CtxItem @click="addTextNode">
        <template #icon="{ strokeWidth, iconClass }">
          <Plus :stroke-width="strokeWidth" :class="iconClass" />
        </template>
        <template #text>Ajouter un noeud texte</template>
      </CtxItem>
      <CtxItem @click="addImageNode">
        <template #icon="{ strokeWidth, iconClass }">
          <ImagePlus :stroke-width="strokeWidth" :class="iconClass" />
        </template>
        <template #text>Ajouter un noeud image</template>
      </CtxItem>
      <CtxItem @click="addVideoNode">
        <template #icon="{ strokeWidth, iconClass }">
          <Video :stroke-width="strokeWidth" :class="iconClass" />
        </template>
        <template #text>Ajouter un noeud vidéo</template>
      </CtxItem>
    </CtxSection>
  </AppCtxMenu>
</template>

<script setup lang="ts">
import { ContextMenuProps } from '@/types/props/ContextMenuProps';
import { ImagePlus, Plus, Video } from 'lucide-vue-next';
import { ref } from 'vue';
import { useFlowchartCtxMenu } from '@/composables/ContextMenus/useFlowchartCtxMenu';
import AppCtxMenu from '../ui/context-menu/AppCtxMenu.vue';
import CtxSection from '../ui/context-menu/CtxSection.vue';
import CtxItem from '../ui/context-menu/CtxItem.vue';

defineProps<ContextMenuProps>();

const ctxMenu = ref<InstanceType<typeof AppCtxMenu> | null>(null);
const { showPopover, addImageNode, addTextNode, addVideoNode } =
  useFlowchartCtxMenu(ctxMenu);

defineExpose({
  showPopover,
});
</script>
