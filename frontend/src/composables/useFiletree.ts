import { node, watcher } from "$/models";
import { nextTick, provide, Ref, ref } from "vue";
import { EventsOn } from "../../wailsjs/runtime"
import { DirPath, FiletreeProvide, ShortNode } from "@/types/FiletreeProvide";
import { FsEvent } from "@/types/FsEvent";
import { useInputToggle } from "./ContextMenus/useInputToggle";

/**
 * Composable that reconciles the in-app filetree and the 
 * user's machine filesystem
 */
export function useFiletree(rootFiles: Ref<node.Node[]>) {
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
  EventsOn('fsop', async function (e: FsEvent) {
    switch (e.op) {
      case watcher.Op.CREATE:
        await addElementInSidePanel(e)
        break;
      case watcher.Op.REMOVE:
        deleteElementFromSidePannelWithPath(e)
        break
      case watcher.Op.RENAME:
        renameElementInSidePannel(e)
        break;
      case watcher.Op.MOVE:
        await moveElementInSidePannel(e)
        break
    }
  });

  async function addElementInSidePanel(e: FsEvent) {
    // If e.path === '.' means that the deletion happened at lab's root
    let dir: Array<ShortNode>
    const wasAdditionAtRoot = e.path === '.'
    if (wasAdditionAtRoot) {
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
    if (wasAdditionAtRoot) {
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
    if (index === -1) {
      console.error("Element index not found")
      return
    }

    // Removing the element
    dir.splice(index, 1)
  }

  function deleteFileFromSidePannelWithOldPath(e: FsEvent) {
    const { dir, index } = getDirAndIndexWithOldPath(e)

    if (!dir) {
      return
    }

    if (index === -1) {
      return
    }

    // Removing the element
    dir.splice(index, 1)
  }

  async function moveElementInSidePannel(e: FsEvent) {
    deleteFileFromSidePannelWithOldPath(e)
    await addElementInSidePanel(e)
  }

  function renameElementInSidePannel(e: FsEvent) {
    const { dir, index } = getDirAndIndexWithOldPath(e)

    if (!dir) {
      return
    }

    if (index === -1) {
      console.error("Element index not found")
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

  function getDirAndIndexWithOldPath(e: FsEvent) {
    let dir: Array<ShortNode>
    const lastIndexOfSlash = e.oldPath.lastIndexOf('/')

    const opAtRoot = !e.oldPath.includes('/')
    if (opAtRoot) {
      // handle lab's root delete
      dir = rootFiles.value

      if (!dir) {
        return {
          dir: undefined,
          index: -1
        }
      }

      const { oldFilepath, extension } = getOldPathAndExtAtRoot(e)
      const index = findElement(oldFilepath, e.dataType, extension, dir)
      return {
        dir,
        index
      }
    }

    dir = dirs.value.get(e.oldPath.slice(0, lastIndexOfSlash)) as ShortNode[]
    if (!dir) {
      return {
        dir: undefined,
        index: -1
      }
    }

    const { oldFilepath, extension } = getOldPathAndExt(e, lastIndexOfSlash)
    const index = findElement(oldFilepath, e.dataType, extension, dir)
    return {
      dir,
      index
    }
  }

  function getOldPathAndExtAtRoot(e: FsEvent) {
    let oldFilepath: string
    let extension: string
    oldFilepath = e.oldPath
    extension = ''

    if (e.dataType === node.DataType.DIR) {
      return {
        oldFilepath,
        extension
      }
    }

    const i = e.oldPath.lastIndexOf('.')
    oldFilepath = e.oldPath.slice(0, i)
    extension = e.oldPath.slice(i)

    return {
      oldFilepath,
      extension
    }
  }

  function getOldPathAndExt(e: FsEvent, lastIndexOfSlash: number) {
    let oldFilepath: string
    let extension: string

    oldFilepath = e.oldPath.slice(lastIndexOfSlash + 1)
    extension = ""

    if (e.dataType === node.DataType.DIR) {
      return {
        oldFilepath,
        extension
      }
    }

    const i = e.oldPath.lastIndexOf('.')
    oldFilepath = e.oldPath.slice(lastIndexOfSlash + 1, i)
    extension = e.oldPath.slice(i)

    return {
      oldFilepath,
      extension
    }
  }

  function add(event: ShortNode, dirs: ShortNode[]) {
    const i = findIndexToInsertAt(event, dirs)
    dirs.splice(i + 1, 0, event)
  }

  function findIndexToInsertAt(event: ShortNode, dirs: ShortNode[], low?: number, high?: number): number {
    if (dirs.length === 0) {
      return -1
    }

    low = low || 0
    high = high || dirs.length
    let median = Math.floor(low + (high - low) / 2)

    // if (high - low <= 1) {
    if (median === low) {
      if (less(dirs, median, event)) {
        return median
      }
      return -1
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