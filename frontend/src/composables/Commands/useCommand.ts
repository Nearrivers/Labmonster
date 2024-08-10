import { computed, onMounted, Ref, ref, watch } from "vue";
import Fuse from 'fuse.js'

// Used with dialog elements.
// The dialog needs to have a ref="dialog" attribute
export function useCommand(model: Ref<string>, list: Readonly<Ref<string[]>>) {
  const activeLine = ref(0)
  const dialog = ref<HTMLDialogElement | null>(null);

  const fuzzyFilteredList = computed(() => {
    const fuse = new Fuse(list.value, { threshold: 0.35 })
    return model.value ? fuse.search(model.value).map((f) => f.item) : list.value
  })

  watch(
    () => fuzzyFilteredList.value.length,
    () => {
      activeLine.value = 0;
    },
  );

  function showModal() {
    dialog.value?.showModal();
  }

  function hideModal() {
    dialog.value?.classList.remove('open:animate-command-show');
    dialog.value?.classList.add('animate-command-hide');
    dialog.value?.addEventListener('animationend', onAnimationFinish);
  }

  function onAnimationFinish() {
    activeLine.value = 0
    dialog.value?.classList.add('open:animate-command-show');
    dialog.value?.classList.remove('animate-command-hide');
    dialog.value?.close();
    dialog.value?.removeEventListener('animationend', onAnimationFinish);
  }

  function onKeyDown() {
    if (activeLine.value < fuzzyFilteredList.value.length - 1) {
      activeLine.value++;
      return;
    }

    activeLine.value = 0;
  }

  function onKeyUp() {
    if (activeLine.value > 0) {
      activeLine.value--;
      return;
    }

    activeLine.value = fuzzyFilteredList.value.length - 1;
  }

  return {
    dialog,
    activeLine,
    fuzzyFilteredList,
    showModal,
    hideModal,
    onKeyDown,
    onKeyUp
  }
}