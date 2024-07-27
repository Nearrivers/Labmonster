import { ref, computed, watch, Ref, onMounted } from "vue";
import Fuse from 'fuse.js'
import { GetDirectories, MoveFileToExistingDir } from "$/filetree/FileTreeExplorer";
import { useShowErrorToast } from "./useShowErrorToast";

export function useMoveFile(selectedNode: HTMLLIElement | null, activeDir: Ref<number>) {
  const { showToast } = useShowErrorToast()
  const directories = ref<string[]>([]);
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
      showToast(error);
    }
  });

  function moveDOMNode(oldPath: string) {
    const newPath = fuzzyFoundDirs.value[activeDir.value];
    const nodeFolder = document.querySelector(`[data-path="${newPath}"] ul`);

    if (nodeFolder && selectedNode) {
      selectedNode.dataset.path = newPath + '/' + oldPath;
      nodeFolder.insertBefore(selectedNode, null);
      Array.from(nodeFolder.children)
        .sort(sortNodes)
        .forEach((node) => nodeFolder.appendChild(node));
    }
  }

  function getAttributeValue(e: Element, attributeName: string) {
    return e.attributes.getNamedItem(attributeName)!.value
  }

  function sortNodes(a: Element, b: Element) {
    const pathA = getAttributeValue(a, 'data-path')
    const typeA = getAttributeValue(a, 'data-type')

    const pathB = getAttributeValue(b, 'data-path')
    const typeB = getAttributeValue(b, 'data-type')

    if (typeA === 'directory' && typeB === 'file') {
      return -1
    }

    if (typeA === 'file' && typeB === 'directory') {
      return 1
    }

    if (pathA < pathB) {
      return -1
    }

    if (pathA === pathB) {
      return 0
    }

    if (pathA > pathB) {
      return 1
    }

    return 0
  }

  async function onSelect(hideModalCb: () => void) {
    if (selectedNode) {
      const oldPath = selectedNode.dataset.path!;
      const extension = selectedNode.dataset.extension;

      try {
        await MoveFileToExistingDir(
          `${oldPath}${extension}`,
          fuzzyFoundDirs.value[activeDir.value],
        );
        moveDOMNode(oldPath);
        hideModalCb();
      } catch (error) {
        showToast(error);
      }
    }
  }

  return {
    path,
    fuzzyFoundDirs,
    onSelect
  }
}