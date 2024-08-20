import { RenameFile } from "$/filetree/FileTree";
import { ref, computed, Ref, getCurrentInstance } from "vue";
import { useShowErrorToast } from "../useShowErrorToast";
import { filetree } from "$/models";

export function useFileNode(props: Ref<{ node: filetree.Node, path: string }>) {
  const { emit } = getCurrentInstance()!
  const { showToast } = useShowErrorToast();
  const fileName = ref(props.value.node.name);
  const input = ref<HTMLInputElement | null>(null);
  const nodePath = ref(
    props.value.path ? props.value.path + '/' + props.value.node.name : props.value.node.name,
  );
  const ext = computed(() =>
    props.value.node.extension.replace('.', '').toLocaleUpperCase(),
  );

  const nodePathWithoutSpaces = computed(() =>
    nodePath.value.replaceAll(' ', '-'),
  );

  const updatedAt = computed(() => {
    const date = new Date(props.value.node.updatedAt);
    return `${date.toLocaleDateString()} à ${date.toLocaleTimeString()}`;
  });

  async function onBlur() {
    if (!input.value) {
      showToast('Input introuvable');
      return;
    }

    if (input.value.readOnly) {
      return
    }

    input.value.toggleAttribute('readonly');
    input.value.classList.add('cursor-pointer');
    input.value.classList.remove('cursor-text');

    try {
      const newName = fileName.value + props.value.node.extension
      await RenameFile(
        props.value.path,
        props.value.node.name + props.value.node.extension,
        newName
      );

      emit('nodeRenamed', newName)
    } catch (error) {
      showToast(error);
    }
  }

  return {
    nodePath,
    nodePathWithoutSpaces,
    input,
    onBlur,
    fileName,
    ext,
    updatedAt
  }
}