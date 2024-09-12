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
    console.log(e)
    switch (e.op) {
      case watcher.Op.CREATE:
        createFileInSidePanel(e)
        break;
      case watcher.Op.REMOVE:
        deleteElementFromSidePannelWithPath(e)
        break
      case watcher.Op.RENAME:
        renameFileInSidePannel(e)
        break;
      case watcher.Op.MOVE:
        moveFileInSidePannel(e)
        break
    }
  });

  async function createFileInSidePanel(e: FsEvent) {
    // If e.path === '.' means that the deletion happened at lab's root
    let dir: Array<object>
    if (e.path === '.') {
      // handle lab's root delete
      dir = rootFiles.value
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

      dir.push({
        name: fileName,
        extension: extension,
        fileType: e.fileType,
        type: e.dataType,
        updatedAt: new Date()
      })

      return
    }

    dir.push({
      name: e.file,
      extension: "",
      fileType: e.fileType,
      type: e.dataType,
      updatedAt: new Date()
    })

    if (e.dataType === node.DataType.DIR) {
      await nextTick()
      toggleInput(e.file, "dir")
    }
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

    console.log(oldFilepath)
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

  function moveFileInSidePannel(e: FsEvent) {
    deleteFileFromSidePannelWithOldPath(e)
    createFileInSidePanel(e)
  }

  function renameFileInSidePannel(e: FsEvent) {
    // If e.path === '.' means that the deletion happened at lab's root
    let dir: Array<object>
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

    const i = e.file.lastIndexOf('.')
    const fileName = e.file.slice(0, i)
    const extension = e.file.slice(i)

    dir.splice(index, 1, {
      name: fileName,
      extension,
      fileType: e.fileType,
      type: e.dataType,
      updatedAt: new Date()
    })
  }
}