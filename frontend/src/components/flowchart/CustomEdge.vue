<template>
  <BezierEdge
    :source-x="sourceX"
    :source-y="sourceY"
    :target-x="targetX"
    :target-y="targetY"
    :source-position="sourcePosition"
    :target-position="targetPosition"
    class="stroke-zinc-600 stroke-[3] ring-border hover:ring-2"
    :curvature="0.7"
    :interaction-width="20"
    :marker-end="markerEnd"
  />
  <div>test</div>
</template>

<script setup lang="ts">
import {
  BezierEdge,
  EdgeMouseEvent,
  EdgeProps,
  useEdge,
  useVueFlow,
} from '@vue-flow/core';

const { onEdgeMouseEnter, onEdgeClick, updateEdge } = useVueFlow();

defineProps<EdgeProps>();

onEdgeMouseEnter((param: EdgeMouseEvent) => {
  console.log(param);
});
</script>
