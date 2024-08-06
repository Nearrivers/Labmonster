<template>
  <BezierEdge
    :source-x="sourceX"
    :source-y="sourceY"
    :target-x="targetX"
    :target-y="targetY"
    :source-position="sourcePosition"
    :target-position="targetPosition"
    class="border-2 stroke-zinc-600 stroke-[3] ring-border hover:ring-2"
    :class="{ '!stroke-primary': isEdgeSelected }"
    :curvature="0.7"
    :interaction-width="20"
    :marker-end="markerEnd"
    :updatable="true"
  />
</template>

<script setup lang="ts">
import {
  BezierEdge,
  EdgeMouseEvent,
  EdgeProps,
  useVueFlow,
} from '@vue-flow/core';
import { computed } from 'vue';

const props = defineProps<EdgeProps>();
const { onEdgeMouseEnter, getSelectedEdges, onEdgeClick, updateEdge } =
  useVueFlow();

onEdgeClick((param: EdgeMouseEvent) => {
  console.log(props.id, isEdgeSelected.value);
});

const isEdgeSelected = computed(() =>
  getSelectedEdges.value.some((e) => e.id === props.id),
);
</script>
