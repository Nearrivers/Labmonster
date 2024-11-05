<template>
  <Teleport to="body">
    <dialog
      ref="dialog"
      class="w-fit max-w-[40rem] animate-command-hide rounded-lg border border-border bg-white text-black shadow-md duration-200 backdrop:bg-black/80 backdrop:transition-opacity backdrop:duration-200 open:animate-command-show dark:bg-secondary dark:text-white"
      @keydown.prevent.esc="hideModal"
    >
      <header class="flex justify-between border-b border-b-border px-1">
        <input
          type="text"
          class="h-11 w-full border-none bg-transparent px-1 py-3 text-sm outline-none"
          v-model="model"
          :placeholder="inputPlaceholder"
          autofocus
          autocomplete="off"
          @keydown.up="onKeyUp"
          @keydown.down="onKeyDown"
          @keydown.enter="emit('select', fuzzyFilteredList[activeLine])"
        />
        <X
          :stroke-width="1.75"
          class="cursor-pointer pt-2"
          @click="hideModal"
        />
      </header>
      <ScrollArea :max-height="'max-h-[80svh]'">
        <ul class="border-b border-b-border text-sm">
          <li
            role="option"
            v-for="(el, index) in fuzzyFilteredList"
            :key="el"
            class="m-1 cursor-default rounded-md px-2 py-1.5"
            :class="{ 'bg-accent': index === activeLine }"
            @mouseenter="activeLine = index"
            @click="emit('select', fuzzyFilteredList[index])"
          >
            {{ el }}
          </li>
          <li
            class="m-1 flex cursor-default items-center justify-between rounded-md bg-secondary px-2 py-1.5"
            v-if="fuzzyFilteredList.length === 0"
          >
            <p>
              {{ model }}
            </p>
            <p class="text-xs font-medium opacity-75">Entrée pour créer</p>
          </li>
        </ul>
      </ScrollArea>
      <footer
        class="flex flex-wrap justify-center gap-3 px-1 py-2 text-xs opacity-70"
      >
        <p class="flex items-center">
          <ArrowUp class="h-3 w-3" :stroke-width="4" />
          <ArrowDown class="mr-1 h-3 w-3" :stroke-width="4" /> pour naviguer
        </p>
        <p class="flex items-center">
          <CornerDownLeft class="mr-1 h-3 w-3" :stroke-width="4" /> pour valider
        </p>
        <p class="flex items-center">
          <span class="mr-1 font-bold"> shift </span>
          <CornerDownLeft class="mr-1 h-3 w-3" :stroke-width="4" />
          pour créer
        </p>
        <p class="flex items-center">
          <span class="mr-1 font-bold">esc</span>
          pour quitter
        </p>
      </footer>
    </dialog>
  </Teleport>
</template>

<script setup lang="ts">
import { ref, toRef } from 'vue';
import { ArrowDown, ArrowUp, CornerDownLeft, X } from 'lucide-vue-next';
import { useCommand } from '@/composables/Commands/useCommand';
import ScrollArea from '../ui/scroll-area/ScrollArea.vue';

const props = defineProps<{
  inputPlaceholder: string;
  list: string[];
}>();

const emit = defineEmits<{
  (e: 'select', selectedLine: string): void;
}>();

const model = ref('');
const {
  dialog,
  activeLine,
  fuzzyFilteredList,
  showModal,
  hideModal,
  onKeyDown,
  onKeyUp,
} = useCommand(
  model,
  toRef(() => props.list),
);

defineExpose({
  showModal,
  hideModal,
});
</script>
