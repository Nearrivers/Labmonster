import { GetSubDirAndFiles } from "$/filetree/FileTree";
import { filetree } from "$/models";
import { computed, ref, Ref } from "vue";
import { useShowErrorToast } from "../useShowErrorToast";

export function useDirNode(props: Ref<{
  node: filetree.Node;
  path: string;
}>) {
  const files = ref<filetree.Node[]>([]);
  const isOpen = ref(false);
  const isFolder = computed(() => props.value.node.type === 'DIR');
  const { showToast } = useShowErrorToast();

  const nodePath = computed(() =>
    props.value.path ? props.value.path + '/' + props.value.node.name : props.value.node.name,
  );

  async function toggle() {
    try {
      files.value = props.value.path
        ? await GetSubDirAndFiles(props.value.path + '/' + props.value.node.name)
        : await GetSubDirAndFiles(props.value.node.name);
    } catch (error) {
      showToast(error);
    } finally {
      isOpen.value = !isOpen.value;
    }
  }

  return {
    files, isOpen, isFolder, nodePath, toggle
  }
}