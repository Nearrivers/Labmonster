import { DuplicateFile } from '$/filetree/FileTreeExplorer';
import { AppDialog } from '@/types/AppDialog';
import { nextTick, Ref, ref } from 'vue';
import { useShowErrorToast } from '../useShowErrorToast';

export function useNodeContextMenu(deleteDialog: Ref<AppDialog | null>) {
  const menu = ref<any | null>(null);
  const { showToast } = useShowErrorToast();

  function showPopover() {
    menu.value?.showPopover();
  }

  function hidePopover() {
    menu.value?.hidePopover();
  }

  async function onRenameClick(filePath: string) {
    hidePopover();
    const inputPath = filePath.replaceAll(' ', '-');
    const fileInput = document.getElementById(inputPath) as HTMLDivElement;

    if (fileInput) {
      fileInput.toggleAttribute('contenteditable');
      fileInput.classList.remove('cursor-pointer');
      fileInput.classList.add('cursor-text');
      fileInput.focus();
    }
  }

  async function onDuplicateClick(filepath: string, extension: string) {
    hidePopover();
    try {
      const file = await DuplicateFile(filepath, extension);
    } catch (error) {
      showToast(error);
    }
  }

  async function onDeleteClick(filePath: string, extension: string) {
    hidePopover();
    await nextTick();
    deleteDialog.value?.openDialog(filePath + extension);
  }

  return {
    menu,
    showPopover,
    hidePopover,
    onRenameClick,
    onDeleteClick,
    onDuplicateClick,
  };
}
