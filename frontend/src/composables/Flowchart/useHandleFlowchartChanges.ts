import { SaveFile } from "$/filetree/FileTreeExplorer";
import { graph } from "$/models";
import { EdgeChange, NodeChange, useVueFlow, ViewportTransform } from "@vue-flow/core";
import { useShowErrorToast } from "../useShowErrorToast";

export function useHandleFlowchartChanges(pathFromLabRoot: string) {
  const { showToast } = useShowErrorToast()
  const { onNodesChange, onEdgesChange, onViewportChange, toObject } = useVueFlow()

  onNodesChange(async (param: NodeChange[]) => {
    if (param.length === 1) {
      const change = param[0]

      if (change.type === 'select') {
        return
      }

      try {
        await SaveFile(pathFromLabRoot, toObject() as unknown as graph.Graph)
      } catch (error) {
        showToast(error)
      }
      return
    }

    console.log('multiple changes', param)
  })

  onEdgesChange((param: EdgeChange[]) => {})

  onViewportChange((param: ViewportTransform) => {})
}