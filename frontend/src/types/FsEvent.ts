import { filetree, watcher } from "$/models"

/**
 * Equivalent of the Event type in the watcher package of the golang side.
 * I have no use of the structs's method so the type is redeclared here
 */
export type FsEvent = {
  oldPath: string
  path: string
  op: watcher.Op
  file: string
  fileType: filetree.FileType
  dataType: filetree.DataType
}