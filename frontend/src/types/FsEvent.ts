import { node, watcher } from "$/models"

/**
 * Equivalent of the Event type in the watcher package of the golang side.
 * I have no use of the structs's method so the type is redeclared here
 */
export type FsEvent = {
  /**
   * Old path to the file. Used in REMOVE and MOVE related functions
   */
  oldPath: string
  /**
   * Path to the parent folder of the file for REMOVE operations
   * For MOVE and RENAME operations, this is the new full folder of the file
   */
  path: string
  /**
   * Operation. See enum
   */
  op: watcher.Op
  /**
   * Name of the file, with its extension
   */
  file: string
  /**
   * Kind of file
   */
  fileType: node.FileType
  /**
   * DIR or FILE
   */
  dataType: node.DataType
}