
export function useInputToggle(hidePopoverFunc?: () => void) {
  function toggleInput(path: string, type: 'file' | 'dir') {
    if (hidePopoverFunc) {
      hidePopoverFunc()
    }

    const inputPath = path.replaceAll(' ', '-') + "-" + type;
    const fileInput = document.getElementById(inputPath) as HTMLInputElement;

    if (fileInput) {
      fileInput.toggleAttribute('readonly');
      fileInput.classList.remove('cursor-pointer');
      fileInput.classList.add('cursor-text');
      fileInput.select();
      fileInput.focus();
    }
  }

  return {
    toggleInput
  }
}