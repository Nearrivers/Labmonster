import { Ref, ref, watch } from "vue";
import { useShowErrorToast } from "../useShowErrorToast";
import { useRouter } from "vue-router";
import { GetRecentlyOpenedFiles } from "$/filetree/FileTree";
import { useMagicKeys } from "@vueuse/core";
import AppCommand from "@/components/ui/AppCommand.vue";

export function useRecentlyOpendedCommand(appCmd: Ref<InstanceType<typeof AppCommand> | null>) {
  const router = useRouter()
  const keys = useMagicKeys()
  const { showToast } = useShowErrorToast()
  const recentlyOpendedFiles = ref<string[]>([])
  const CtrlO = keys['Ctrl+O']

  watch(CtrlO, async (v) => {
    if (!v || !appCmd.value) {
      return
    }

    try {
      recentlyOpendedFiles.value = await GetRecentlyOpenedFiles()
      appCmd.value.showModal()
    } catch (error) {
      showToast(error)
    }
  })


  function onSelect(path: string, hideModalCb: () => void) {
    router.push({
      name: 'flowchart',
      params: { path },
    });

    hideModalCb()
  }

  return {
    recentlyOpendedFiles,
    onSelect
  }
}