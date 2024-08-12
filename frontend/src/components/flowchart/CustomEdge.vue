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
    :selectable="true"
  />
  <EdgeLabelRenderer>
    <EdgeLabel :edge-id="id" :path="path" />
  </EdgeLabelRenderer>
</template>

<script setup lang="ts">
import {
  BezierEdge,
  EdgeLabelRenderer,
  EdgeProps,
  getBezierPath,
  useVueFlow,
} from '@vue-flow/core';
import { computed } from 'vue';
import EdgeLabel from './EdgeLabel.vue';

const props = defineProps<EdgeProps>();
const path = computed(() => getBezierPath(props));
const { getSelectedEdges } = useVueFlow();

const isEdgeSelected = computed(() =>
  getSelectedEdges.value.some((e) => e.id === props.id),
);
</script>
