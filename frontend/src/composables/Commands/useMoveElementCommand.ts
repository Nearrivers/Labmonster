import { ref, onMounted } from 'vue';
import { MoveFileToExistingDir } from '$/file_handler/FileHandler';
import { GetDirectories, MoveDir } from '$/dirhandler/DirHandler';
import { useShowErrorToast } from '../useShowErrorToast';
import { NodeElement } from '../../types/NodeElement';

export function useMoveElement(selectedNode: NodeElement | null) {
  const { showToast } = useShowErrorToast();
  const directories = ref<string[]>([]);

  onMounted(async () => {
    await loadDirectories();
  });

  async function loadDirectories() {
    try {
      directories.value = await GetDirectories();
    } catch (error) {
      showToast(error);
    }
  }

  async function onSelect(newPath: string) {
    if (!selectedNode) {
      return;
    }

    const oldPath = selectedNode.dataset.path;
    const type = selectedNode.dataset.type;

    try {
      if (type === 'file') {
        const extension = selectedNode.dataset.extension;
        await MoveFileToExistingDir(`${oldPath}${extension}`, newPath);
      }

      if (type === 'directory') {
        console.log(oldPath, newPath, selectedNode);
        await MoveDir(oldPath, newPath);
        await loadDirectories();
      }
    } catch (error) {
      showToast(error);
    }
  }

  return {
    directories,
    loadDirectories,
    onSelect,
  };
}
