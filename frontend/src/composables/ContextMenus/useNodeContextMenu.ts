import { DuplicateFile } from '$/filetree/FileTree';
import { AppDialog } from '@/types/AppDialog';
import { nextTick, Ref, ref } from 'vue';
import { useShowErrorToast } from '../useShowErrorToast';
import AppCtxMenu from '@/components/contextmenus/AppCtxMenu.vue';

export function useNodeContextMenu(ctxMenu: Ref<InstanceType<typeof AppCtxMenu> | null>, deleteDialog: Ref<AppDialog | null>) {
  const { showToast } = useShowErrorToast();

  function showPopover() {
    ctxMenu.value?.showPopover();
  }

  function hidePopover() {
    ctxMenu.value?.hidePopover();
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
    ctxMenu,
    showPopover,
    hidePopover,
    onRenameClick,
    onDeleteClick,
    onDuplicateClick,
  };
}
