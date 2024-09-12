import { GetSubDirAndFiles } from '$/file_handler/FileHandler';
import { node } from '$/models';
import { computed, inject, ref, Ref } from 'vue';
import { useShowErrorToast } from '../useShowErrorToast';
import { FiletreeProvide } from '@/types/FiletreeProvide';
import { RenameDirectory } from '$/dirhandler/DirHandler';

export function useDirNode(
  props: Ref<{
    dirNode: node.Node;
    path: string;
  }>,
) {
  const files = ref<node.Node[]>([]);
  const isOpen = ref(false);
  const dirName = ref(props.value.dirNode.name);
  const isFolder = computed(() => props.value.dirNode.type === 'DIR');
  const input = ref<HTMLInputElement | null>(null);
  const { showToast } = useShowErrorToast();
  const { addDir } = inject<FiletreeProvide>('dirs')!;

  const nodePath = computed(() =>
    props.value.path
      ? props.value.path + '/' + props.value.dirNode.name
      : props.value.dirNode.name,
  );

  const nodePathWithoutSpaces = computed(() =>
    nodePath.value.replaceAll(' ', '-') + "-dir",
  );

  async function toggle() {
    try {
      let p = '';
      if (!props.value.path) {
        p = props.value.dirNode.name;
      } else {
        p = props.value.path + '/' + props.value.dirNode.name;
      }

      files.value = await GetSubDirAndFiles(p);
      addDir(p, files.value);
    } catch (error) {
      showToast(error);
    } finally {
      isOpen.value = !isOpen.value;
    }
  }

  async function onBlur() {
    if (!input.value) {
      showToast('Input introuvable');
      return;
    }

    if (input.value.readOnly) {
      return;
    }

    input.value.toggleAttribute('readonly');
    input.value.classList.add('cursor-pointer');
    input.value.classList.remove('cursor-text');

    try {
      const newName = dirName.value;
      await RenameDirectory(props.value.path, newName);
    } catch (error) {
      showToast(error);
    }
  }

  return {
    input,
    files,
    isOpen,
    isFolder,
    nodePath,
    toggle,
    dirName,
    nodePathWithoutSpaces,
    onBlur,
  };
}
