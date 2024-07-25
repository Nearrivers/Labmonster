import { ref } from "vue";

// Used with dialog elements.
// The dialog needs to have a ref="dialog" attribute
export function useCommand() {
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

  return {
    dialog,
    showModal,
    hideModal
  }
}