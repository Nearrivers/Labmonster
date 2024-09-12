import { DuplicateFile } from '$/file_handler/FileHandler';
import { AppDialog } from '@/types/AppDialog';
import { nextTick, Ref } from 'vue';
import { useShowErrorToast } from '../useShowErrorToast';
import AppCtxMenu from '@/components/contextmenus/AppCtxMenu.vue';
import { useInputToggle } from './useInputToggle';

export function useFileContextMenu(
  ctxMenu: Ref<InstanceType<typeof AppCtxMenu> | null>,
  deleteDialog: Ref<AppDialog | null>,
) {
  const { toggleInput } = useInputToggle(hidePopover)
  const { showToast } = useShowErrorToast();

  function showPopover() {
    ctxMenu.value?.showPopover();
  }

  function hidePopover() {
    ctxMenu.value?.hidePopover();
  }

  async function onDuplicateClick(filepath: string, extension: string) {
    hidePopover();
    try {
      await DuplicateFile(filepath, extension);
    } catch (error) {
      showToast(error);
    }
  }

  async function onDeleteClick(filePath: string, extension: string) {
    console.log(filePath, extension)

    hidePopover();
    await nextTick();
    deleteDialog.value?.openDialog(filePath + extension);
  }

  return {
    showPopover,
    hidePopover,
    toggleInput,
    onDeleteClick,
    onDuplicateClick,
  };
}
