import { Node, useVueFlow } from '@vue-flow/core';

export function useTopMenuActions() {
  const { zoomIn, zoomOut } = useVueFlow();

  function createNewNode(): Node {
    const id = Date.now().toString();
    const newNode: Node = {
      id,
      position: { x: 150, y: 50 },
      data: { hello: `Node ${id}` },
      type: 'custom',
      style: { width: 250, height: 60 }
    }

    return newNode
  }

  return {
    createNewNode,
    zoomIn,
    zoomOut,
  };
}
