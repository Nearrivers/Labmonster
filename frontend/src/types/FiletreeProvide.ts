import { filetree, watcher } from "$/models"
import { Ref } from "vue"

/**
 * Type used for documentation purposes. Makes clear what kind of string
 * the map exepects as keys
 */
export type DirPath = string

/**
 * Utility type that removes useless keys from the type so I don't have
 * to create Partial types everywhere
 */
export type ShortNode = Omit<filetree.Node, 'convertValues' | 'updatedAt'>

export type FiletreeProvide = {
  dirs: Ref<Map<DirPath, ShortNode[]>>
  addDir: (path: DirPath, elements: ShortNode[]) => void
}