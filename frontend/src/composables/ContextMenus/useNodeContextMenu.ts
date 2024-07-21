import { AppDialog } from "@/types/AppDialog";
import { nextTick, Ref, ref } from "vue";

export function useNodeContextMenu(deleteDialog: Ref<AppDialog | null>) {
  const menu = ref<any | null>(null);

  function showPopover() {
    menu.value?.showPopover();
  }

  function hidePopover() {
    menu.value?.hidePopover();
  }

  async function onRenameClick(filePath: string) {
    hidePopover()
    const inputPath = filePath.replaceAll(' ', '-')
    const fileInput = document.getElementById(inputPath) as HTMLInputElement

    if (fileInput) {
      fileInput.toggleAttribute('disabled')
      fileInput.classList.remove('cursor-pointer')
      fileInput.classList.add('cursor-text')
      fileInput.focus()
      fileInput.select()
    }
  }

  async function onDeleteClick(filePath: string) {
    hidePopover();
    await nextTick();
    deleteDialog.value?.openDialog(filePath);
  }

  return {
    menu,
    showPopover,
    hidePopover,
    onRenameClick,
    onDeleteClick
  }
}