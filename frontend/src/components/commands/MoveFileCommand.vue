<template>
  <Teleport to="body">
    <dialog
      ref="dialog"
      class="w-full max-w-[30rem] animate-command-hide rounded-lg border border-border bg-white text-black shadow-md duration-200 backdrop:bg-black/80 backdrop:transition-opacity backdrop:duration-200 open:animate-command-show dark:bg-black dark:text-white"
      @keydown.prevent.esc="hideModal"
    >
      <header class="flex justify-between border-b border-b-border px-1">
        <input
          type="text"
          class="h-11 w-full border-none bg-transparent px-1 py-3 text-sm outline-none"
          v-model="path"
          placeholder="Saisir le nom d'un dossier"
          autofocus
          autocomplete="off"
          @keydown.up="onKeyUp"
          @keydown.down="onKeyDown"
          @keydown.enter="onSelect"
        />
        <X
          :stroke-width="1.75"
          class="cursor-pointer pt-2"
          @click="hideModal"
        />
      </header>
      <ul class="border-b border-b-border text-sm">
        <li
          role="option"
          v-for="(dir, index) in fuzzyFoundDirs"
          class="m-1 cursor-default rounded-md px-2 py-1.5"
          :class="{ 'bg-secondary': index === activeDir }"
          @mouseenter="activeDir = index"
          @click="onSelect"
        >
          {{ dir }}
        </li>
        <li
          class="m-1 flex cursor-default items-center justify-between rounded-md bg-secondary px-2 py-1.5"
          v-if="fuzzyFoundDirs.length === 0"
        >
          <p>
            {{ path }}
          </p>
          <p class="text-xs font-medium opacity-75">Entrée pour créer</p>
        </li>
      </ul>
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
          <span class="mr-1 font-bold"> shift </span
          ><CornerDownLeft class="mr-1 h-3 w-3" :stroke-width="4" />
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
import {
  GetDirectories,
  MoveFile,
  RenameFile,
} from '$/filetree/FileTreeExplorer';
import { useShowErrorToast } from '@/composables/useShowErrorToast';
import { computed, onMounted, ref, watch } from 'vue';
import { ArrowDown, ArrowUp, CornerDownLeft, X } from 'lucide-vue-next';
import { useCommand } from '@/composables/useCommand';
import Fuse from 'fuse.js';

const props = defineProps<{
  oldPath?: string;
  extension?: string;
}>();

const { dialog, showModal, hideModal } = useCommand();
const { showToast } = useShowErrorToast();
const directories = ref<string[]>([]);
const activeDir = ref(0);
let fuse: Fuse<string>;
const path = ref('');

const fuzzyFoundDirs = computed(() =>
  path.value ? fuse.search(path.value).map((f) => f.item) : directories.value,
);

watch(
  () => fuzzyFoundDirs.value.length,
  () => {
    activeDir.value = 0;
  },
);

onMounted(async () => {
  try {
    directories.value = await GetDirectories();
    fuse = new Fuse(directories.value, { threshold: 0.35 });
  } catch (error) {
    showToast(String(error));
  }
});

function onKeyDown() {
  if (activeDir.value < directories.value.length - 1) {
    activeDir.value++;
    return;
  }

  activeDir.value = 0;
}

function onKeyUp() {
  if (activeDir.value > 0) {
    activeDir.value--;
    return;
  }

  activeDir.value = directories.value.length - 1;
}

async function onSelect() {
  try {
    await MoveFile(
      `${props.oldPath}${props.extension}`,
      fuzzyFoundDirs.value[activeDir.value],
    );
  } catch (error) {
    showToast(String(error));
  }
}

defineExpose({
  showModal,
  hideModal,
});
</script>
