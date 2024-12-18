import { GetSubDirAndFiles, CreateFile } from '$/file_handler/FileHandler';
import { node } from '$/models';
import { NEW_FILE_NAME } from '@/constants/NEW_FILE_NAME';
import { nextTick, ref } from 'vue';
import { useShowErrorToast } from './useShowErrorToast';
import FileContextMenu from '@/components/contextmenus/FileContextMenu.vue';
import { CONFIG_FILE_LOADED } from '@/constants/event-names/CONFIG_FILE_LOADED';
import { configFileLoaded } from '@/events/ReloadFileExplorer';
import DirContextMenu from '../components/contextmenus/DirContextMenu.vue';
import { useRouter } from 'vue-router';
import { Routes } from '@/types/Routes';
import { SupportedFiles } from '@/types/SupportedFiles';
import { useEventListener } from './useEventListener';
import { CreateDirectory } from '$/dirhandler/DirHandler';
import { NEW_DIR_NAME } from '@/constants/NEW_DIR_NAME';
import { NodeElement } from '../types/NodeElement';

export function useSidePanel() {
  const files = ref<node.Node[]>([]);
  const router = useRouter();
  const { showToast } = useShowErrorToast();
  const contextMenuX = ref(100);
  const contextMenuY = ref(100);
  const selectedNode = ref<NodeElement | null>(null);
  const fileContextMenu = ref<InstanceType<typeof FileContextMenu> | null>(
    null,
  );
  const dirContextMenu = ref<InstanceType<typeof DirContextMenu> | null>(null);
  useEventListener(configFileLoaded, CONFIG_FILE_LOADED, loadLabFiles);

  async function loadLabFiles() {
    try {
      files.value = await GetSubDirAndFiles('');
    } catch (error) {
      showToast(error, 'Impossible de charger les fichiers');
    }
  }

  async function createNewFileAtRoot() {
    try {
      await CreateFile(NEW_FILE_NAME);
    } catch (error) {
      showToast(error, 'Impossible de créer le fichier');
    }
  }

  async function createNewDirAtRoot() {
    try {
      await CreateDirectory(NEW_DIR_NAME);
    } catch (error) {
      showToast(error, 'Impossible de créer le dossier');
    }
  }

  async function onRightClick(event: MouseEvent) {
    contextMenuX.value = event.clientX;
    contextMenuY.value = event.clientY;
    selectedNode.value = (event.target as HTMLElement).closest(
      'li',
    ) as unknown as NodeElement;
    await nextTick();

    if (selectedNode.value?.dataset.type === 'file') {
      fileContextMenu.value?.showPopover();
      return;
    }

    dirContextMenu.value?.showPopover();
  }

  function onLeftClick(event: MouseEvent) {
    const node = (event.target as HTMLElement).closest(
      'li',
    ) as unknown as NodeElement;

    selectedNode.value = node

    if (!node || node.dataset.type === 'directory') {
      return;
    }

    let name: Routes;
    const { path, extension } = node.dataset;
    switch (node.dataset.file) {
      case SupportedFiles.GRAPH:
        name = Routes.Flowchart;
        break;
      case SupportedFiles.IMAGE:
        name = Routes.Image;
        break;
      case SupportedFiles.VIDEO:
        name = Routes.Video;
        break;
      default:
        name = Routes.Unsupported;
        break;
    }

    router.push({
      name,
      params: { path: path?.includes('.json') ? path! : path! + extension },
    });
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
    createNewFolderAtRoot: createNewDirAtRoot,
    onRightClick,
    onLeftClick,
    showToast,
  };
}
