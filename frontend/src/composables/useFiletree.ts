import { filetree, watcher } from "$/models";
import { ref } from "vue";
import { EventsOn } from "../../wailsjs/runtime"

/**
 * Type used for documentation purposes. Makes clear what kind of string
 * the map exepects as keys
 */
type Path = string

/**
 * Equivalent of the Event type in the watcher package of the golang side.
 * I have no use of the structs's method so the type is redeclared here
 */
type FsEvent = {
  oldPath: string
  path: string
  op: watcher.Op
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
  const dirs = ref(new Map<Path, filetree.Node[]>())



  // On every "operation" that happens in the filesystem, the go side will launch an
  // event that gets caught here.
  EventsOn('fsop', function (e: FsEvent) {
    switch (e.op) {
      case watcher.Op.CREATE:
      case watcher.Op.REMOVE:
      case watcher.Op.RENAME:
      case watcher.Op.MOVE:
    }
    console.log(e)
  });
}