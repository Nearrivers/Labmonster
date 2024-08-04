import { Node, useVueFlow } from '@vue-flow/core';
import { Ref } from 'vue';

export function useTopMenuActions(nodes: Ref<Node[]>) {
  const { zoomIn, zoomOut } = useVueFlow();

  function createNewNode() {
    const id = Date.now().toString();
    const newNode = {
      id,
      position: { x: 150, y: 50 },
      data: { hello: `Node ${id}` },
      type: 'custom',
    }

    return newNode
  }

  return {
    createNewNode,
    zoomIn,
    zoomOut,
  };
}
