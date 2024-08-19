<template>
  <AppCtxMenu :y="y" :x="x" popover-id="dirpopover" ref="ctxMenu">
    <CtxSection>
      <CtxItem>
        <template #icon="{ strokeWidth, iconClass }">
          <SquarePen :stroke-width="strokeWidth" :class="iconClass" />
        </template>
        <template #text>Nouveau setup</template>
      </CtxItem>
      <CtxItem>
        <template #icon="{ strokeWidth, iconClass }">
          <FolderOpen :stroke-width="strokeWidth" :class="iconClass" />
        </template>
        <template #text>Nouveau dossier</template>
      </CtxItem>
    </CtxSection>
    <CtxSection>
      <CtxItem>
        <template #icon="{ strokeWidth, iconClass }">
          <Files :stroke-width="strokeWidth" :class="iconClass" />
        </template>
        <template #text>Dupliquer</template>
      </CtxItem>
      <CtxItem>
        <template #icon="{ strokeWidth, iconClass }">
          <FolderTree :stroke-width="strokeWidth" :class="iconClass" />
        </template>
        <template #text>Déplacer le dossier vers...</template>
      </CtxItem>
    </CtxSection>
    <CtxSection>
      <CtxItem>
        <template #icon="{ strokeWidth, iconClass }">
          <PencilLine :stroke-width="strokeWidth" :class="iconClass" />
        </template>
        <template #text>Renommer</template>
      </CtxItem>
      <CtxItem class="text-red-500">
        <template #icon="{ strokeWidth, iconClass }">
          <Trash2 :stroke-width="strokeWidth" :class="iconClass" />
        </template>
        <template #text>Supprimer</template>
      </CtxItem>
    </CtxSection>
  </AppCtxMenu>
</template>

<script setup lang="ts">
import {
  Files,
  FolderOpen,
  FolderTree,
  PencilLine,
  SquarePen,
  Trash2,
} from 'lucide-vue-next';
import { ref } from 'vue';
import AppCtxMenu from '../ui/context-menu/AppCtxMenu.vue';
import CtxSection from '../ui/context-menu/CtxSection.vue';
import CtxItem from '../ui/context-menu/CtxItem.vue';

defineProps<{
  x: number;
  y: number;
}>();

const ctxMenu = ref<InstanceType<typeof AppCtxMenu> | null>(null);

defineExpose({
  showPopover: () => ctxMenu.value!.showPopover(),
  hidePopover: () => ctxMenu.value!.hidePopover(),
});
</script>
