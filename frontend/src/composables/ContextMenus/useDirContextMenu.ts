import AppCtxMenu from "@/components/ui/context-menu/AppCtxMenu.vue";
import { AppDialog } from "@/types/AppDialog";
import { Ref } from "vue";
import { useShowErrorToast } from "../useShowErrorToast";
import { CreateDirectory } from "$/dirhandler/DirHandler";
import { NEW_DIR_NAME } from "@/constants/NEW_DIR_NAME";
import { NEW_FILE_NAME } from "@/constants/NEW_FILE_NAME";
import { CreateFile } from "$/file_handler/FileHandler";

export function useDirContextMenu(ctxMenu: Ref<InstanceType<typeof AppCtxMenu> | null>, deleteDialog: AppDialog | null) {
  const { showToast } = useShowErrorToast()

  function showPopover() {
    ctxMenu.value?.showPopover();
  }

  function hidePopover() {
    ctxMenu.value?.hidePopover();
  }

  async function createNewSetup(path: string) {
    try {
      const newFile = path + "/" + NEW_FILE_NAME
      const file = await CreateFile(newFile)
    } catch (error) {
      showToast(error)
    }
  }

  async function createNewDirectory(path: string) {
    try {
      const newDir = path + "/" + NEW_DIR_NAME
      await CreateDirectory(newDir)
    } catch (error) {
      showToast(error)
    }
  }

  return {
    showPopover,
    hidePopover,
    createNewDirectory,
    createNewSetup
  }
}