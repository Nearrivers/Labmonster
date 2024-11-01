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
    let extension: string
    if (e.path === '.') {
      // handle lab's root delete
      dir = rootFiles.value

      oldFilepath = e.file.slice(0, e.file.lastIndexOf('.'))
      extension = e.file.slice(e.file.lastIndexOf('.'))
      if (e.dataType === node.DataType.DIR) {
        oldFilepath = e.file
      }
    } else {
      dir = dirs.value.get(e.path as DirPath) as ShortNode[]
      // Removing the extension
      oldFilepath = e.oldPath.slice(0, e.oldPath.lastIndexOf('.'))
      extension = e.oldPath.slice(e.oldPath.lastIndexOf('.'))
      if (e.dataType === node.DataType.DIR) {
        oldFilepath = e.oldPath
        extension = ""
      }
    }

    if (!dir) {
      return
    }

    const index = findElement(oldFilepath, e.dataType, extension, dir)
    // // Finding the element using its path
    // const el = document.querySelector(`[data-path="${oldFilepath}"`) as HTMLLIElement
    // if (!el) {
    //   return
    // }

    // // Using the data-id attribute
    // const index = parseInt(el.dataset.id!)
    if (index === -1) {
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
    let extension: string
    if (!e.oldPath.includes('/')) {
      // handle lab's root delete
      dir = rootFiles.value
      // Removing the extension
      oldFilepath = e.file.slice(0, e.file.lastIndexOf('.'))
      extension = e.file.slice(e.file.lastIndexOf('.'))
    } else {
      // Removing the extension
      oldFilepath = e.oldPath.slice(0, e.oldPath.lastIndexOf('.'))
      extension = e.file.slice(e.file.lastIndexOf('.'))
      dir = dirs.value.get(e.oldPath.slice(0, e.oldPath.lastIndexOf('/'))) as ShortNode[]
    }

    if (!dir) {
      return
    }

    const index = findElement(oldFilepath, e.dataType, extension, dir)
    if (index === -1) {
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
    let oldFilepath: string
    let extension: string

    if (e.dataType === node.DataType.FILE) {
      const i = e.oldPath.lastIndexOf('.')
      oldFilepath = e.oldPath.slice(0, i)
      extension = e.oldPath.slice(i)
    } else {
      oldFilepath = e.oldPath
      extension = ''
    }

    const index = findElement(oldFilepath, e.dataType, extension, dir)
    // selecting the element in the DOM was faster
    // but hard to test so I used a binary search instead
    // let selector = `[data-path="${oldFilepath}"]`
    // if (e.dataType === node.DataType.FILE) {
    //   selector += '[data-type="file"]'
    // } else {
    //   selector += '[data-type="directory"]'
    // }

    // // Finding the element using its path
    // const el = document.querySelector(selector) as HTMLLIElement
    // if (!el) {
    //   return
    // }

    // // Using the data-id attribute
    // const index = parseInt(el.dataset.id!)
    if (index === -1) {
      showErrorToastFunc("Element index not found")
      return
    }

    dir.splice(index, 1)

    if (e.dataType === node.DataType.FILE) {
      const i = e.file.lastIndexOf('.')
      const fileName = e.file.slice(0, i)
      const extension = e.file.slice(i)

      add({
        name: fileName,
        extension,
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
  }

  function add(event: ShortNode, dirs: ShortNode[]) {
    const i = findIndexToInsertAt(event, dirs)
    dirs.splice(i + 1, 0, event)
  }

  function findIndexToInsertAt(event: ShortNode, dirs: ShortNode[], low?: number, high?: number): number {
    low = low || 0
    high = high || dirs.length
    let median = Math.floor(low + (high - low) / 2)

    if (high - low <= 1) {
      if (median === 0) {
        return -1
      }
      return median
    }

    if (less(dirs, median, event)) {
      return findIndexToInsertAt(event, dirs, median, high)
    } else {
      return findIndexToInsertAt(event, dirs, low, median)
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

  /**
   * @param name - Name of the element we're looking for
   * @param type - Directory or File
   * @param extension - File extension. Dir's ext = ""
   * @param dirs - Elements array to search into
   * @returns the index of the found element or -1 if not found
   */
  function findElement(name: string, type: node.DataType, extension: string, dirs: ShortNode[], low?: number, high?: number): number {
    low = low || 0
    high = high || dirs.length
    const median = Math.floor(low + (high - low) / 2)

    const dir = dirs[median]
    if (dir.name === name && dir.type === type && dir.extension === extension) {
      return median
    }

    if (high - low <= 1) {
      if (dirs[median].name === name && dirs[median].type === type && dir.extension === extension) {
        return median
      }

      return -1
    }

    if (compare(dirs, median, name, type, extension)) {
      return findElement(name, type, extension, dirs, median, high)
    } else {
      return findElement(name, type, extension, dirs, low, median)
    }
  }

  function compare(dirs: ShortNode[], i: number, name: string, type: node.DataType, extension: string): boolean {
    const el = dirs[i]

    if (el.type === type) {
      return new Intl.Collator().compare(el.name + el.extension, name + extension) < 0
    }

    if (el.type === node.DataType.DIR) {
      return true
    }

    return false
  }

  return {
    dirs,
    addDir,
    addElementInSidePanel,
    renameElementInSidePannel,
    moveElementInSidePannel,
    deleteElementFromSidePannelWithPath,
    deleteFileFromSidePannelWithOldPath,
  }
}