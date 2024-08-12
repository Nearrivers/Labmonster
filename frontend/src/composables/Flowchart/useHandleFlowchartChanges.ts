import { SaveFile } from "$/filetree/FileTree";
import { graph } from "$/models";
import { NodeChange, useVueFlow } from "@vue-flow/core";
import { useShowErrorToast } from "../useShowErrorToast";
import { Ref, ref } from "vue";

export function useHandleFlowchartChanges(pathFromLabRoot: Ref<string>) {
  const isSaving = ref(false)
  const { showToast } = useShowErrorToast()
  const { onNodesChange, onEdgesChange, findNode, onViewportChangeEnd, toObject } = useVueFlow()

  async function Save() {
    try {
      isSaving.value = true
      await SaveFile(pathFromLabRoot.value, toObject() as unknown as graph.Graph)
    } catch (error) {
      showToast(error)
    } finally {
      isSaving.value = false
    }
  }

  onNodesChange(async (param: NodeChange[]) => {
    await Save()
  })
  onEdgesChange(async (_) => { await Save() })
  onViewportChangeEnd(async (_) => { await Save() })

  return {
    isSaving
  }
}