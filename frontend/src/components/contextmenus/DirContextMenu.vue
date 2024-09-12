<template>
  <AppCtxMenu :y="y" :x="x" popover-id="dirpopover" ref="ctxMenu">
    <CtxSection>
      <CtxItem @click="createNewSetup(path)">
        <template #icon="{ strokeWidth, iconClass }">
          <SquarePen :stroke-width="strokeWidth" :class="iconClass" />
        </template>
        <template #text>Nouveau setup</template>
      </CtxItem>
      <CtxItem @click="createNewDirectory(path)">
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
      <CtxItem @click="toggleInput(path, 'dir')">
        <template #icon="{ strokeWidth, iconClass }">
          <PencilLine :stroke-width="strokeWidth" :class="iconClass" />
        </template>
        <template #text>Renommer</template>
      </CtxItem>
      <CtxItem class="text-red-500" @click="onDeleteClick(path)">
        <template #icon="{ strokeWidth, iconClass }">
          <Trash2 :stroke-width="strokeWidth" :class="iconClass" />
        </template>
        <template #text>Supprimer</template>
      </CtxItem>
    </CtxSection>
    <template #commands>
      <DeleteDirDialog ref="deleteDialog" :path="selectedNode?.dataset.path" />
    </template>
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
import { computed, ref } from 'vue';
import AppCtxMenu from '../ui/context-menu/AppCtxMenu.vue';
import CtxSection from '../ui/context-menu/CtxSection.vue';
import CtxItem from '../ui/context-menu/CtxItem.vue';
import { useDirContextMenu } from '@/composables/ContextMenus/useDirContextMenu';
import DeleteDirDialog from '../AlertDialog/DeleteDirDialog.vue';

const props = defineProps<{
  x: number;
  y: number;
  selectedNode: HTMLLIElement | null;
}>();

const deleteDialog = ref<InstanceType<typeof DeleteDirDialog> | null>(null);
const ctxMenu = ref<InstanceType<typeof AppCtxMenu> | null>(null);
const {
  showPopover,
  hidePopover,
  createNewSetup,
  createNewDirectory,
  onDeleteClick,
  toggleInput,
} = useDirContextMenu(ctxMenu, deleteDialog);

const path = computed(() => props.selectedNode?.dataset.path || '');

defineExpose({
  showPopover,
  hidePopover,
});
</script>
