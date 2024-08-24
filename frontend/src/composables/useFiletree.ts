import { filetree, watcher } from "$/models";
import { provide, ref } from "vue";
import { EventsOn } from "../../wailsjs/runtime"
import { DirPath, FiletreeProvide, ShortNode } from "@/types/filetreeProvide";

/**
 * Equivalent of the Event type in the watcher package of the golang side.
 * I have no use of the structs's method so the type is redeclared here
 */
type FsEvent = {
  oldPath: string
  path: string
  op: watcher.Op
  file: string
  fileType: filetree.FileType
  dataType: filetree.DataType
}

/**
 * Composable that reconcile the in-app filetree and the 
 * user's machine filesystem
 */
export function useFiletree() {
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
    addDir
  })

  // On every "operation" that happens in the filesystem, the go side will launch an
  // event that gets caught here.
  EventsOn('fsop', function (e: FsEvent) {
    switch (e.op) {
      case watcher.Op.CREATE:
        OnCreate(e)
      case watcher.Op.REMOVE:
      case watcher.Op.RENAME:
      case watcher.Op.MOVE:
    }
    console.log(e)
  });

  function OnCreate(e: FsEvent) {
    // We get the reference to the file array
    const dir = dirs.value.get(e.path as DirPath)

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
      type: e.dataType
    })
  }
}