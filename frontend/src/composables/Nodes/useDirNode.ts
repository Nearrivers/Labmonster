import { GetSubDirAndFiles } from "$/filetree/FileTree";
import { filetree } from "$/models";
import { computed, inject, ref, Ref } from "vue";
import { useShowErrorToast } from "../useShowErrorToast";
import { FiletreeProvide } from "@/types/filetreeProvide";

export function useDirNode(props: Ref<{
  node: filetree.Node;
  path: string;
}>) {
  const files = ref<filetree.Node[]>([]);
  const isOpen = ref(false);
  const isFolder = computed(() => props.value.node.type === 'DIR');
  const { showToast } = useShowErrorToast();
  const { addDir } = inject<FiletreeProvide>("dirs")!

  const nodePath = computed(() =>
    props.value.path ? props.value.path + '/' + props.value.node.name : props.value.node.name,
  );

  async function toggle() {
    try {
      let p = ''
      if (!props.value.path) {
        p = props.value.node.name
      } else {
        p = props.value.path + '/' + props.value.node.name
      }

      files.value = await GetSubDirAndFiles(p)
      addDir(p, files.value)
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