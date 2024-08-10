import { GetSubDirAndFiles, CreateFile } from '$/filetree/FileTree';
import { filetree } from '$/models';
import { NEW_FILE_NAME } from '@/constants/NEW_FILE_NAME';
import { nextTick, ref } from 'vue';
import { useShowErrorToast } from './useShowErrorToast';
import FileContextMenu from '@/components/contextmenus/FileContextMenu.vue';
import { CONFIG_FILE_LOADED } from '@/constants/event-names/CONFIG_FILE_LOADED';
import { configFileLoaded } from '@/events/ReloadFileExplorer';
import { useEventListener } from '@vueuse/core';
import DirContextMenu from '../components/contextmenus/DirContextMenu.vue';
import { useRouter } from 'vue-router';

export function useSidePanel() {
  const files = ref<filetree.Node[]>([]);
  const router = useRouter();
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
      const file = await CreateFile(NEW_FILE_NAME);
      // removing .json from the file name
      file.name = file.name.slice(0, file.name.indexOf('.'));
      files.value.push(file);
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

  function onLeftClick(event: MouseEvent) {
    const node = (event.target as HTMLElement).closest('li');

    if (!node) {
      return;
    }

    if (node.dataset.type === 'file') {
      const { path, extension } = node.dataset;
      router.push({
        name: 'flowchart',
        params: { path: path?.includes('.json') ? path! : path! + extension },
      });
    }
  }

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
