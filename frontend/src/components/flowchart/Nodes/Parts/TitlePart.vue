<template>
  <header class="flex w-full justify-between">
    <input
      ref="input"
      class="resize-none bg-transparent outline-none"
      :class="{ nodrag: isReadOnly, underline: !isReadOnly }"
      :style="{ width: `${title.length}ch` }"
      autocomplete="off"
      v-model="title"
      @keyup.enter=""
      @blur="isReadOnly = true"
      :readonly="isReadOnly"
    />
    <ToolbarButton
      @click="changeReadonlyState"
      :side="'right'"
      :side-offset="25"
    >
      <template #icon>
        <Pencil class="h-5 w-5" />
      </template>
      <template #tooltip> Modifier le titre </template>
    </ToolbarButton>
  </header>
</template>

<script setup lang="ts">
import { Pencil } from 'lucide-vue-next';
import { ref } from 'vue';
import ToolbarButton from '../../ToolbarButton.vue';

const isReadOnly = ref(true);
const input = ref<HTMLInputElement | null>(null);
const title = defineModel<string, string>({ default: 'Nouveau noeud' });

function changeReadonlyState() {
  isReadOnly.value = !isReadOnly.value;
  input.value?.select();
  input.value?.focus();
}
</script>
