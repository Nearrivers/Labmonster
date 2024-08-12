<template>
  <div
    :style="{
      pointerEvents: 'all',
      position: 'absolute',
      transform: `translate(-50%, -100%) translate(${path[1]}px,${path[2]}px)`,
    }"
  >
    <textarea
      ref="input"
      role="textbox"
      aria-multiline="true"
      v-model="label"
      v-show="label || isEditing"
      class="box-content resize-none rounded-sm bg-background p-1 text-center text-primary outline-none ring-border focus:ring"
      @blur="isEditing = false"
      @input="handleLabelChange"
      @focus="handleLabelFocus"
    ></textarea>
    <EdgeToolbar @edit="onEdit" @remove="onRemove" />
  </div>
</template>

<script setup lang="ts">
import { nextTick, onMounted, ref, watch } from 'vue';
import EdgeToolbar from './EdgeToolbar.vue';
import { EdgeMouseEvent, useEdge, useVueFlow } from '@vue-flow/core';

const props = defineProps<{
  edgeId: string;
  path: [
    path: string,
    labelX: number,
    labelY: number,
    offsetX: number,
    offsetY: number,
  ];
}>();

const { edge } = useEdge(props.edgeId);
const label = defineModel<string>('label');
const input = ref<HTMLInputElement | null>(null);
const isEditing = ref(false);
const { onEdgeDoubleClick, addSelectedEdges } = useVueFlow();

onMounted(() => {
  label.value = edge.data.label;
});

watch(label, () => {
  if (!input.value) {
    return;
  }

  input.value.style.height = '24px';

  if (input.value.scrollHeight > 32) {
    input.value.style.height = input.value.scrollHeight + 'px';
  }
});

onEdgeDoubleClick((param: EdgeMouseEvent) => {
  if (param.edge.id === props.edgeId) {
    onEdit();
  }
});

async function onEdit() {
  if (input.value) {
    isEditing.value = true;
    await nextTick();
    input.value.select();
  }
}

async function onRemove() {
  label.value = '';
  await nextTick();
  isEditing.value = false;
  handleLabelChange();
}

function handleLabelChange() {
  edge.data = {
    ...edge.data,
    label: label.value,
  };
}

function handleLabelFocus() {
  if (!edge.selected) {
    addSelectedEdges([edge]);
  }
}
</script>
