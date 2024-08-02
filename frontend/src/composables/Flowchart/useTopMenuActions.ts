import { Node, useVueFlow } from '@vue-flow/core';
import { Ref } from 'vue';

export function useTopMenuActions(nodes: Ref<Node[]>) {
  const { zoomIn, zoomOut } = useVueFlow();

  function addNode() {
    const id = Date.now().toString();

    nodes.value.push({
      id,
      position: { x: 150, y: 50 },
      data: { label: `Node ${id}` },
      class: 'bg-primary border border-border',
    });
    console.log(nodes.value);
  }

  return {
    addNode,
    zoomIn,
    zoomOut,
  };
}
