import {
  GetSubDirAndFiles,
  CreateNewFileAtRoot,
} from '$/filetree/FileTreeExplorer';
import { filetree } from '$/models';
import { NEW_FILE_NAME } from '@/constants/NEW_FILE_NAME';
import { nextTick, ref } from 'vue';
import { useShowErrorToast } from './useShowErrorToast';
import FileContextMenu from '@/components/contextmenus/FileContextMenu.vue';
import { CONFIG_FILE_LOADED } from '@/constants/event-names/CONFIG_FILE_LOADED';
import { configFileLoaded } from '@/events/ReloadFileExplorer';
import { useEventListener } from '@vueuse/core';
import DirContextMenu from '../components/contextmenus/DirContextMenu.vue';
import { sidePanelToggled } from '@/events/ToggleSidePanel';
import { SIDE_PANEL_TOGGLED } from '@/constants/event-names/SIDE_PANEL_TOGGLED';

export function useSidePanel() {
  const files = ref<filetree.Node[]>([]);
  const { showToast } = useShowErrorToast();
  const contextMenuX = ref(100);
  const contextMenuY = ref(100);
  const selectedNode = ref<HTMLLIElement | null>(null);
  const fileContextMenu = ref<InstanceType<typeof FileContextMenu> | null>(
    null,
  );
  const dirContextMenu = ref<InstanceType<typeof DirContextMenu> | null>(null);
  useEventListener(configFileLoaded, CONFIG_FILE_LOADED, loadLabFiles);

  function sortNodes(f1: filetree.Node, f2: filetree.Node) {
    // Tri sur les types d'abord
    if (f1.type === 'DIR' && f2.type == 'FILE') {
      return -1;
    }

    if (f1.type === 'FILE' && f2.type == 'DIR') {
      return 1;
    }

    if (f1.name < f2.name) {
      return -1;
    }

    if (f1.name == f2.name) {
      return 0;
    }

    if (f1.name > f2.name) {
      return 1;
    }

    return 0;
  }

  async function loadLabFiles() {
    try {
      files.value = await GetSubDirAndFiles('');
    } catch (error) {
      showToast(error, 'Impossible de charger les fichiers');
    }
  }

  async function createNewFileAtRoot() {
    try {
      const newFileName = await CreateNewFileAtRoot(NEW_FILE_NAME);
      files.value.push(newFileName);
      files.value.sort(sortNodes);
    } catch (error) {
      showToast(error, 'Impossible de créer le fichier');
    }
  }

  async function onRightClick(event: MouseEvent) {
    contextMenuX.value = event.clientX;
    contextMenuY.value = event.clientY;
    selectedNode.value = (event.target as HTMLElement).closest('li');
    await nextTick();

    if (selectedNode.value?.dataset.type === 'file') {
      fileContextMenu.value?.showPopover();
      return;
    }

    dirContextMenu.value?.showPopover();
  }

  function onLeftClick(event: MouseEvent) {}

  return {
    files,
    contextMenuX,
    contextMenuY,
    fileContextMenu,
    dirContextMenu,
    selectedNode,
    loadLabFiles,
    createNewFileAtRoot,
    onRightClick,
    onLeftClick,
    showToast,
  };
}
