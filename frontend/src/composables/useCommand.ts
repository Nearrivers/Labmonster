import { ComputedRef, Ref, ref } from "vue";

// Used with dialog elements.
// The dialog needs to have a ref="dialog" attribute
export function useCommand(activeLine: Ref<number>, computedList: ComputedRef<string[]>) {
  const dialog = ref<HTMLDialogElement | null>(null);

  function showModal() {
    dialog.value?.showModal();
  }

  function hideModal() {
    dialog.value?.classList.remove('open:animate-command-show');
    dialog.value?.classList.add('animate-command-hide');
    dialog.value?.addEventListener('animationend', onAnimationFinish);
  }

  function onAnimationFinish() {
    dialog.value?.classList.add('open:animate-command-show');
    dialog.value?.classList.remove('animate-command-hide');
    dialog.value?.close();
    dialog.value?.removeEventListener('animationend', onAnimationFinish);
  }

  function onKeyDown() {
    if (activeLine.value < computedList.value.length - 1) {
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

    activeLine.value = computedList.value.length - 1;
  }

  return {
    dialog,
    activeLine,
    showModal,
    hideModal,
    onKeyDown,
    onKeyUp
  }
}