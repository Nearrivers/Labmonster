import { Ref } from "vue";
import { useShowErrorToast } from "../useShowErrorToast";
import { useVueFlow } from "@vue-flow/core";
import AppCtxMenu from "@/components/ui/context-menu/AppCtxMenu.vue";

enum NodeType {
  Text = 'custom',
  Image = 'image',
  Video = 'video'
}

export function useFlowchartCtxMenu(ctxMenu: Ref<InstanceType<typeof AppCtxMenu> | null>) {
  const { addNodes, screenToFlowCoordinate } = useVueFlow()
  const { showToast } = useShowErrorToast();

  function showPopover() {
    ctxMenu.value?.showPopover();
  }

  function addTextNode(e: MouseEvent) {
    addNode(e.clientX, e.clientY, NodeType.Text)
  }

  function addImageNode(e: MouseEvent) {
    addNode(e.clientX, e.clientY, NodeType.Image)
  }

  function addVideoNode(e: MouseEvent) {
    addNode(e.clientX, e.clientY, NodeType.Video)
  }

  function addNode(x: number, y: number, nodeType: NodeType) {
    const flowCoordinates = screenToFlowCoordinate({
      x, y
    })

    addNodes({
      id: Date.now().toString(),
      position: flowCoordinates,
      type: nodeType
    })
    ctxMenu.value?.hidePopover()
  }

  return {
    showPopover,
    addTextNode,
    addImageNode,
    addVideoNode
  }
}