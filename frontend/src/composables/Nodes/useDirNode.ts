import { GetSubDirAndFiles } from "$/file_handler/FileHandler";
import { node } from "$/models";
import { computed, inject, ref, Ref } from "vue";
import { useShowErrorToast } from "../useShowErrorToast";
import { FiletreeProvide } from "@/types/FiletreeProvide";

export function useDirNode(props: Ref<{
  dirNode: node.Node;
  path: string;
}>) {
  const files = ref<node.Node[]>([]);
  const isOpen = ref(false);
  const isFolder = computed(() => props.value.dirNode.type === 'DIR');
  const { showToast } = useShowErrorToast();
  const { addDir } = inject<FiletreeProvide>("dirs")!

  const nodePath = computed(() =>
    props.value.path ? props.value.path + '/' + props.value.dirNode.name : props.value.dirNode.name,
  );

  async function toggle() {
    try {
      let p = ''
      if (!props.value.path) {
        p = props.value.dirNode.name
      } else {
        p = props.value.path + '/' + props.value.dirNode.name
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