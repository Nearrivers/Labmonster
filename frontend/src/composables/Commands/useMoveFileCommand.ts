import { ref, onMounted } from "vue";
import { MoveFileToExistingDir } from "$/file_handler/FileHandler";
import { GetDirectories } from "$/dirhandler/DirHandler";
import { useShowErrorToast } from "../useShowErrorToast";

export function useMoveFile(selectedNode: HTMLLIElement | null) {
  const { showToast } = useShowErrorToast()
  const directories = ref<string[]>([]);

  onMounted(async () => {
    try {
      directories.value = await GetDirectories();
    } catch (error) {
      showToast(error);
    }
  });


  async function onSelect(newPath: string, hideModalCb: () => void) {
    if (selectedNode) {
      const oldPath = selectedNode.dataset.path!;
      const extension = selectedNode.dataset.extension;

      try {
        await MoveFileToExistingDir(
          `${oldPath}${extension}`,
          newPath,
        );
        hideModalCb();
      } catch (error) {
        showToast(error);
      }
    }
  }

  return {
    directories,
    onSelect
  }
}