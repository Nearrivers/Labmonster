import { ref, onMounted } from "vue";
import { GetDirectories, MoveFileToExistingDir } from "$/filetree/FileTree";
import { useShowErrorToast } from "../useShowErrorToast";

export function useMoveFile(selectedNode: HTMLLIElement | null) {
  const { showToast } = useShowErrorToast()
  const directories = ref<string[]>([]);

  onMounted(async () => {
    try {
      directories.value = await GetDirectories();
    } catch (error) {
      showToast(error);
    }
  });

  function moveDOMNode(newPath: string, oldPath: string) {
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

  async function onSelect(newPath: string, hideModalCb: () => void) {
    if (selectedNode) {
      const oldPath = selectedNode.dataset.path!;
      const extension = selectedNode.dataset.extension;

      try {
        await MoveFileToExistingDir(
          `${oldPath}${extension}`,
          newPath,
        );
        moveDOMNode(newPath, oldPath);
        hideModalCb();
      } catch (error) {
        showToast(error);
      }
    }
  }

  return {
    directories,
    onSelect
  }
}