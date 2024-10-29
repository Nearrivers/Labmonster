<template>
  <AppCtxMenu :x="x" :y="y" :popover-id="'filepopover'" ref="ctxMenu">
    <CtxSection>
      <CtxItem
        @click="onDuplicateClick(selectedNode?.dataset.path!, extension)"
      >
        <template #icon="{ strokeWidth, iconClass }">
          <Files :stroke-width="strokeWidth" :class="iconClass" />
        </template>
        <template #text>Dupliquer</template>
      </CtxItem>
      <CtxItem @click="onMoveClick">
        <template #icon="{ strokeWidth, iconClass }">
          <FolderTree :stroke-width="strokeWidth" :class="iconClass" />
        </template>
        <template #text>Déplacer le fichier vers...</template>
      </CtxItem>
    </CtxSection>
    <CtxSection>
      <CtxItem @click="toggleInput(selectedNode?.dataset.path!, 'file')">
        <template #icon="{ strokeWidth, iconClass }">
          <PencilLine :stroke-width="strokeWidth" :class="iconClass" />
        </template>
        <template #text>Renommer</template>
      </CtxItem>
      <CtxItem
        @click="onDeleteClick(selectedNode?.dataset.path!, extension)"
        class="text-red-500"
      >
        <template #icon="{ strokeWidth, iconClass }">
          <Trash2 :stroke-width="strokeWidth" :class="iconClass" />
        </template>
        <template #text>Supprimer</template>
      </CtxItem>
    </CtxSection>
    <template #commands>
      <DeleteFileDialog
        ref="deleteFileDialog"
        :path="selectedNode?.dataset.path"
        :extension="selectedNode?.dataset.extension"
      />
    </template>
  </AppCtxMenu>
</template>

<script setup lang="ts">
import { Files } from 'lucide-vue-next';
import { FolderTree, Trash2, PencilLine } from 'lucide-vue-next';
import { computed, ref } from 'vue';
import DeleteFileDialog from '../AlertDialog/DeleteFileDialog.vue';
import { useFileContextMenu } from '@/composables/ContextMenus/useFileContextMenu';
import AppCtxMenu from '../ui/context-menu/AppCtxMenu.vue';
import CtxSection from '../ui/context-menu/CtxSection.vue';
import CtxItem from '../ui/context-menu/CtxItem.vue';
import { NodeElement } from '@/types/NodeElement';

const props = defineProps<{
  x: number;
  y: number;
  selectedNode: NodeElement | null;
}>();

const emit = defineEmits<{
  (e: 'move'): void;
}>();

const ctxMenu = ref<InstanceType<typeof AppCtxMenu> | null>(null);
const extension = computed(() => props.selectedNode?.dataset.extension || '');

const deleteFileDialog = ref<InstanceType<typeof DeleteFileDialog> | null>(
  null,
);

const {
  showPopover,
  hidePopover,
  toggleInput,
  onDeleteClick,
  onDuplicateClick,
} = useFileContextMenu(ctxMenu, deleteFileDialog);

function onMoveClick() {
  emit('move');
  hidePopover();
}

defineExpose({
  showPopover,
  hidePopover,
});
</script>
