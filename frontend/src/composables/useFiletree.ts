import { node, watcher } from "$/models";
import { nextTick, provide, Ref, ref } from "vue";
import { EventsOn } from "../../wailsjs/runtime"
import { DirPath, FiletreeProvide, ShortNode } from "@/types/FiletreeProvide";
import { FsEvent } from "@/types/FsEvent";
import { ShowToastFunc } from "./useShowErrorToast";
import { useInputToggle } from "./ContextMenus/useInputToggle";

/**
 * Composable that reconciles the in-app filetree and the 
 * user's machine filesystem
 */
export function useFiletree(rootFiles: Ref<node.Node[]>, showErrorToastFunc: ShowToastFunc) {
  const { toggleInput } = useInputToggle()
  /**
   * Map that holds the path to a directory as the key and 
   * the reference to that directory files array as value
   */
  const dirs = ref(new Map<DirPath, ShortNode[]>())

  /**
   * @param path - Path to the dir added
   * @param elements - Files and dirs of the parent dir whose path was just provide as the map's key
   */
  function addDir(path: DirPath, elements: ShortNode[]) {
    dirs.value.set(path, elements)
  }

  provide<FiletreeProvide>("dirs", {
    dirs,
    addDir,
  })

  // On every "operation" that happens in the filesystem, the go side will launch an
  // event that gets caught here.
  EventsOn('fsop', function (e: FsEvent) {
    switch (e.op) {
      case watcher.Op.CREATE:
        addElementInSidePanel(e)
        break;
      case watcher.Op.REMOVE:
        deleteElementFromSidePannelWithPath(e)
        break
      case watcher.Op.RENAME:
        renameElementInSidePannel(e)
        break;
      case watcher.Op.MOVE:
        moveElementInSidePannel(e)
        break
    }
  });

  async function addElementInSidePanel(e: FsEvent) {
    // If e.path === '.' means that the deletion happened at lab's root
    let dir: Array<ShortNode>
    const wasDeletionAtRoot = e.path === '.'
    if (wasDeletionAtRoot) {
      // handle lab's root delete
      dir = rootFiles.value as ShortNode[]
    } else {
      dir = dirs.value.get(e.path as DirPath) as ShortNode[]
    }

    // If we can't find it that means the user has not loaded
    // the dir files yet so we leave
    if (!dir) {
      return
    }

    if (e.dataType === node.DataType.FILE) {
      const i = e.file.lastIndexOf('.')
      const fileName = e.file.slice(0, i)
      const extension = e.file.slice(i)

      add({
        name: fileName,
        extension: extension,
        fileType: e.fileType,
        type: e.dataType,
        updatedAt: new Date()
      }, dir)

      return
    }

    add({
      name: e.file,
      extension: "",
      fileType: e.fileType,
      type: e.dataType,
      updatedAt: new Date()
    }, dir)

    await nextTick()
    if (wasDeletionAtRoot) {
      toggleInput(e.file, "dir")
      return
    }

    toggleInput(e.path + "/" + e.file, "dir")
  }

  function deleteElementFromSidePannelWithPath(e: FsEvent) {
    // If e.path === '.' means that the deletion happened at lab's root
    let dir: node.Node[] | ShortNode[]
    let oldFilepath: string
    if (e.path === '.') {
      // handle lab's root delete
      dir = rootFiles.value

      oldFilepath = e.file.slice(0, e.file.lastIndexOf('.'))
      if (e.dataType === node.DataType.DIR) {
        oldFilepath = e.file
      }
    } else {
      dir = dirs.value.get(e.path as DirPath) as ShortNode[]
      // Removing the extension
      oldFilepath = e.oldPath.slice(0, e.oldPath.lastIndexOf('.'))
      if (e.dataType === node.DataType.DIR) {
        oldFilepath = e.oldPath
      }
    }

    if (!dir) {
      return
    }

    // Finding the element using its path
    const el = document.querySelector(`[data-path="${oldFilepath}"`) as HTMLLIElement
    if (!el) {
      return
    }

    // Using the data-id attribute
    const index = parseInt(el.dataset.id!)
    if (isNaN(index)) {
      showErrorToastFunc("Element index not found")
      return
    }

    // Removing the element
    dir.splice(index, 1)
  }

  function deleteFileFromSidePannelWithOldPath(e: FsEvent) {
    // If e.oldpath doesn't include '/', that means that the deletion happened at lab's root
    let dir: node.Node[] | ShortNode[]
    let oldFilepath: string
    if (!e.oldPath.includes('/')) {
      // handle lab's root delete
      dir = rootFiles.value
      // Removing the extension
      oldFilepath = e.file.slice(0, e.file.lastIndexOf('.'))
    } else {
      // Removing the extension
      oldFilepath = e.oldPath.slice(0, e.oldPath.lastIndexOf('.'))
      dir = dirs.value.get(e.oldPath.slice(0, e.oldPath.lastIndexOf('/'))) as ShortNode[]
    }

    if (!dir) {
      return
    }

    // Finding the element using its path
    const el = document.querySelector(`[data-path="${oldFilepath}"`) as HTMLLIElement
    if (!el) {
      return
    }

    // Using the data-id attribute
    const index = parseInt(el.dataset.id!)
    if (isNaN(index)) {
      showErrorToastFunc("Element index not found")
      return
    }

    // Removing the element
    dir.splice(index, 1)
  }

  function moveElementInSidePannel(e: FsEvent) {
    deleteFileFromSidePannelWithOldPath(e)
    addElementInSidePanel(e)
  }

  function renameElementInSidePannel(e: FsEvent) {
    // If e.path === '.' means that the deletion happened at lab's root
    let dir: Array<ShortNode>
    if (!e.oldPath.includes('/')) {
      // handle lab's root delete
      dir = rootFiles.value
    } else {
      // Removing the extension
      dir = dirs.value.get(e.oldPath.slice(0, e.oldPath.lastIndexOf('/'))) as ShortNode[]
    }

    if (!dir) {
      return
    }

    // Removing the extension
    const oldFilepath = e.oldPath.slice(0, e.oldPath.lastIndexOf('.'))

    let selector = `[data-path="${oldFilepath}"]`
    if (e.dataType === node.DataType.FILE) {
      selector += '[data-type="file"]'
    } else {
      selector += '[data-type="directory"]'
    }

    // Finding the element using its path
    const el = document.querySelector(selector) as HTMLLIElement
    if (!el) {
      return
    }

    // Using the data-id attribute
    const index = parseInt(el.dataset.id!)
    if (isNaN(index)) {
      showErrorToastFunc("Element index not found")
      return
    }

    const i = e.file.lastIndexOf('.')
    const fileName = e.file.slice(0, i)
    const extension = e.file.slice(i)

    dir.splice(index, 1)

    add({
      name: fileName,
      extension,
      fileType: e.fileType,
      type: e.dataType,
      updatedAt: new Date()
    }, dir)
  }

  function add(event: ShortNode, dirs: ShortNode[]) {
    dirs.splice(findIndex(event, dirs) + 1, 0, event)
  }

  function findIndex(event: ShortNode, dirs: ShortNode[], low?: number, high?: number): number {
    low = low || 0
    high = high || dirs.length
    let pivot = Math.floor(low + (high - low) / 2)

    if (high - low <= 1 || (dirs[pivot].name === event.name && dirs[pivot].extension === event.extension)) {
      return pivot
    }

    if (less(dirs, pivot, event)) {
      return findIndex(event, dirs, pivot, high)
    } else {
      return findIndex(event, dirs, low, pivot)
    }
  }

  function less(dir: ShortNode[], i: number, event: ShortNode): boolean {
    const el = dir[i]

    if (el.type === event.type) {
      return new Intl.Collator().compare(el.name, event.name) < 0
    }

    if (el.type === node.DataType.DIR) {
      return true
    }

    return false
  }
}