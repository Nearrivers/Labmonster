import { filetree, watcher } from "$/models";
import { provide, Ref, ref } from "vue";
import { EventsOn } from "../../wailsjs/runtime"
import { DirPath, FiletreeProvide, ShortNode } from "@/types/filetreeProvide";
import { FsEvent } from "@/types/FsEvent";
import { ShowToastFunc } from "./useShowErrorToast";

/**
 * Composable that reconciles the in-app filetree and the 
 * user's machine filesystem
 */
export function useFiletree(rootFiles: Ref<filetree.Node[]>, showErrorToastFunc: ShowToastFunc) {
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
        OnCreate(e)
        break;
      case watcher.Op.REMOVE:
        onDelete(e)
        break
      case watcher.Op.RENAME:
      case watcher.Op.MOVE:
        onMoveOrRename(e)
        break
    }
    console.log(e)
  });

  function OnCreate(e: FsEvent) {
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

    const i = e.file.lastIndexOf('.')
    const fileName = e.file.slice(0, i)
    const extension = e.file.slice(i)

    dir.push({
      name: fileName,
      extension: extension,
      fileType: e.fileType,
      type: e.dataType,
      updatedAt: new Date().toLocaleDateString()
    })
  }

  function onDelete(e: FsEvent) {
    // If e.path === '.' means that the deletion happened at lab's root
    let dir: filetree.Node[] | ShortNode[]
    if (e.path === '.') {
      // handle lab's root delete
      dir = rootFiles.value
    } else {
      dir = dirs.value.get(e.path as DirPath) as ShortNode[]
    }

    if (!dir) {
      return
    }

    // Removing the extension
    const oldFilepath = e.oldPath.slice(0, e.oldPath.lastIndexOf('.'))
    // Finding the element using its path
    const el = document.querySelector(`[data-path="${oldFilepath}"`) as HTMLLIElement
    if (!el) {
      showErrorToastFunc("Element not found")
    }

    // Using the data-id attribute
    const index = parseInt(el.dataset.id!)
    if (isNaN(index)) {
      showErrorToastFunc("Element index not found")
    }

    // Removing the element
    dir.splice(index, 1)
  }

  function onMoveOrRename(e: FsEvent) {
    onDelete(e)
    OnCreate(e)
  }
}