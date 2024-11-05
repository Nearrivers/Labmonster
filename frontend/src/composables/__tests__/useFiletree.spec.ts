import { node, watcher } from "$/models";
import { FsEvent } from "@/types/FsEvent";
import { afterAll, beforeAll, describe, expect, it, vi } from "vitest";
import { useFiletree } from "../useFiletree";
import { toRef } from "vue";

describe("useFiletree composable", () => {
  beforeAll(() => {
    vi.mock('../../../wailsjs/runtime')
  })

  describe("addElementInSidePanel tests", () => {
    it("should push a file at the end of the root array when the path is '.'", async () => {
      const e: FsEvent = {
        dataType: node.DataType.FILE,
        file: "graph.json",
        fileType: node.FileType.GRAPH,
        oldPath: '.',
        op: watcher.Op.CREATE,
        path: "."
      }

      const rootFiles: node.Node[] = []

      const { addElementInSidePanel } = useFiletree(toRef(rootFiles))
      await addElementInSidePanel(e)

      const file = rootFiles[0]

      expect(rootFiles.length).toBe(1)
      expect(file.name).toBe('graph')
      expect(file.extension).toBe('.json')
      expect(file.fileType).toBe(node.FileType.GRAPH)
      expect(file.type).toBe(node.DataType.FILE)
    })

    it("should push a file before the one already present in the root array as the elements are alphabetically sorted", async () => {
      const e: FsEvent = {
        dataType: node.DataType.FILE,
        file: "graph.json",
        fileType: node.FileType.GRAPH,
        oldPath: '.',
        op: watcher.Op.CREATE,
        path: "."
      }

      const rootFiles: node.Node[] = []
      const { addElementInSidePanel } = useFiletree(toRef(rootFiles))
      // adding the first element
      await addElementInSidePanel(e)

      // adding the second
      e.file = "abc.json"
      await addElementInSidePanel(e)

      const file = rootFiles[0]

      expect(file.name).toBe('abc')
    })

    it("should push a 4 files and keep the elements alphabetically sorted", async () => {
      const e: FsEvent = {
        dataType: node.DataType.FILE,
        file: "graph.json",
        fileType: node.FileType.GRAPH,
        oldPath: '.',
        op: watcher.Op.CREATE,
        path: "."
      }

      const rootFiles: node.Node[] = []
      const { addElementInSidePanel } = useFiletree(toRef(rootFiles))
      // adding the first element
      await addElementInSidePanel(e)

      // adding the second, should be second
      e.file = "abc.json"
      await addElementInSidePanel(e)

      // adding the third, should be first
      e.file = "aac.json"
      await addElementInSidePanel(e)

      // adding the fourth, should be at the end
      e.file = "zoink.json"
      await addElementInSidePanel(e)

      const firstFile = rootFiles[0]
      const lastFile = rootFiles.at(-1)

      expect(firstFile.name).toBe('aac')
      expect(lastFile?.name).toBe('zoink')
    })

    it('should push the new directory first as directories come before files regardless of their names', async () => {
      const e: FsEvent = {
        dataType: node.DataType.FILE,
        file: "graph.json",
        fileType: node.FileType.GRAPH,
        oldPath: '.',
        op: watcher.Op.CREATE,
        path: "."
      }

      const rootFiles: node.Node[] = []
      const { addElementInSidePanel } = useFiletree(toRef(rootFiles))
      await addElementInSidePanel(e)

      e.file = "newDir"
      e.dataType = node.DataType.DIR
      await addElementInSidePanel(e)

      const shouldBeDir = rootFiles[0]
      expect(shouldBeDir.name).toBe('newDir')
      expect(shouldBeDir.type).toBe(node.DataType.DIR)
    })

    it('should keep directories alphabetically sorted between each others', async () => {
      const e: FsEvent = {
        dataType: node.DataType.FILE,
        file: "graph.json",
        fileType: node.FileType.GRAPH,
        oldPath: '.',
        op: watcher.Op.CREATE,
        path: "."
      }

      const rootFiles: node.Node[] = []
      const { addElementInSidePanel } = useFiletree(toRef(rootFiles))
      await addElementInSidePanel(e)

      e.file = "newDir"
      e.dataType = node.DataType.DIR
      await addElementInSidePanel(e)

      e.file = "firstDir"
      e.dataType = node.DataType.DIR
      await addElementInSidePanel(e)

      const firstDir = rootFiles[0]
      const newDir = rootFiles[1]
      expect(newDir.name).toBe('newDir')
      expect(newDir.type).toBe(node.DataType.DIR)
      expect(firstDir.name).toBe('firstDir')
      expect(firstDir.type).toBe(node.DataType.DIR)
    })

    it('should add an element inside the right directory', async () => {
      const e: FsEvent = {
        dataType: node.DataType.FILE,
        file: "graph.json",
        fileType: node.FileType.GRAPH,
        oldPath: '.',
        op: watcher.Op.CREATE,
        path: "."
      }

      const newDirName = "newDir"

      const rootFiles: node.Node[] = []
      const { dirs, addDir, addElementInSidePanel } = useFiletree(toRef(rootFiles))
      await addElementInSidePanel(e)

      e.file = newDirName
      e.dataType = node.DataType.DIR
      await addElementInSidePanel(e)

      // Adding the directory inside the dirs Map
      addDir(newDirName, [])

      const subDirName = "subDir"
      e.file = subDirName
      // The path of the subDir should be /newDir with / being the root of the lab
      e.path = "newDir"
      e.dataType = node.DataType.DIR
      await addElementInSidePanel(e)

      const newDir = dirs.value.get(newDirName)
      expect(newDir).toBeDefined()
      expect(newDir?.length).toBe(1)
      expect(newDir![0].name).toBe(subDirName)
    })
  })

  /**
   * 
   * @param eltName - The element name with its extension
   * @param eltType - File or Directory
   */
  async function addElementHelper(eltName: string, eltType: node.DataType): Promise<{
    rootFiles: node.Node[],
    fileTreeComposable: ReturnType<typeof useFiletree>
  }> {
    const e: FsEvent = {
      dataType: eltType,
      file: eltName,
      fileType: node.FileType.GRAPH,
      oldPath: '.',
      op: watcher.Op.CREATE,
      path: "."
    }

    const rootFiles: node.Node[] = []
    const fileTreeComposable = useFiletree(toRef(rootFiles))
    await fileTreeComposable.addElementInSidePanel(e)

    return {
      rootFiles,
      fileTreeComposable
    }
  }

  describe("renameElementInSidePannel tests", () => {
    it('should rename the element', async () => {
      const fileToRename = "graph.json"
      const {
        rootFiles,
        fileTreeComposable
      } = await addElementHelper(fileToRename, node.DataType.FILE)

      const { renameElementInSidePannel } = fileTreeComposable

      const newFileName = "newName.json"
      const newName = newFileName.split(".")[0]
      const e: FsEvent = {
        dataType: node.DataType.FILE,
        file: newFileName,
        fileType: node.FileType.GRAPH,
        oldPath: fileToRename,
        op: watcher.Op.RENAME,
        path: "."
      }
      renameElementInSidePannel(e)

      const file = rootFiles[0]
      expect(file.name).toEqual(newName)
    })

    it('shoud not rename anything given we give an name that does not exists', async () => {
      const fileThatWontChange = "graph.json"
      const {
        rootFiles,
        fileTreeComposable
      } = await addElementHelper(fileThatWontChange, node.DataType.FILE)

      const { renameElementInSidePannel } = fileTreeComposable
      const fakeFileName = "fake.json"
      const fakePath = "fakepath.json"
      const e: FsEvent = {
        dataType: node.DataType.FILE,
        file: fakeFileName,
        fileType: node.FileType.GRAPH,
        oldPath: fakePath,
        op: watcher.Op.RENAME,
        path: "."
      }
      renameElementInSidePannel(e)

      const file = rootFiles[0]
      const fileName = fileThatWontChange.split('.')[0]
      expect(file.name).toBe(fileName)
      // Checking if another element was not added
      expect(rootFiles.length).toBe(1)
    })

    it('shoud rename a directory too', async () => {
      const dirToRename = "renDir"
      const {
        rootFiles,
        fileTreeComposable
      } = await addElementHelper(dirToRename, node.DataType.DIR)

      const { renameElementInSidePannel } = fileTreeComposable
      const newDirName = "newDirName"
      const e: FsEvent = {
        dataType: node.DataType.DIR,
        file: newDirName,
        fileType: node.FileType.UNSUPPORTED,
        oldPath: dirToRename,
        op: watcher.Op.RENAME,
        path: "."
      }
      renameElementInSidePannel(e)

      const dir = rootFiles[0]
      expect(dir.name).toBe(newDirName)
    })

    it('should rename outside of root', async () => {
      const parentDirName = "test"
      const {
        fileTreeComposable
      } = await addElementHelper(parentDirName, node.DataType.DIR)
      const { dirs, addDir, addElementInSidePanel, renameElementInSidePannel } = fileTreeComposable
      addDir(parentDirName, [])

      const subDirName = "subDir"

      const e: FsEvent = {
        dataType: node.DataType.DIR,
        file: subDirName,
        fileType: node.FileType.GRAPH,
        oldPath: "",
        op: watcher.Op.CREATE,
        path: parentDirName
      }
      addElementInSidePanel(e)
      addDir(subDirName, [])

      const newName = "newName"
      e.file = newName
      e.oldPath = parentDirName + "/" + subDirName
      e.op = watcher.Op.RENAME
      e.dataType = node.DataType.DIR
      e.path = parentDirName
      renameElementInSidePannel(e)

      const parentDir = dirs.value.get(parentDirName)!
      expect(parentDir).toBeDefined()
      expect(parentDir[0].name).toBe(newName)
    })
  })

  describe("deleteElementFromSidePannelWithPath tests", () => {
    it('should remove an element if it exists', async () => {
      const elementToDelete = "graph.json"
      const {
        rootFiles,
        fileTreeComposable
      } = await addElementHelper(elementToDelete, node.DataType.FILE)
      const { deleteElementFromSidePannelWithPath } = fileTreeComposable
      const e: FsEvent = {
        dataType: node.DataType.FILE,
        file: elementToDelete,
        fileType: node.FileType.GRAPH,
        oldPath: "",
        op: watcher.Op.REMOVE,
        path: "."
      }
      deleteElementFromSidePannelWithPath(e)

      expect(rootFiles.length).toBe(0)
    })

    it('should not remove an element that does not exists', async () => {
      const elementToDelete = "graph.json"
      const {
        rootFiles,
        fileTreeComposable
      } = await addElementHelper(elementToDelete, node.DataType.FILE)
      const { deleteElementFromSidePannelWithPath } = fileTreeComposable

      const fakeName = "fake.json"
      const e: FsEvent = {
        dataType: node.DataType.FILE,
        file: fakeName,
        fileType: node.FileType.GRAPH,
        oldPath: "",
        op: watcher.Op.REMOVE,
        path: "."
      }
      deleteElementFromSidePannelWithPath(e)

      const presentFileName = elementToDelete.split(".")[0]
      expect(rootFiles.length).toBe(1)
      expect(rootFiles[0].name).toBe(presentFileName)
    })
  })

  describe("deleteElementFromSidePannelWithOldPath tests", () => {
    it('should remove an element if it exists', async () => {
      const elementToDelete = "graph.json"
      const {
        rootFiles,
        fileTreeComposable
      } = await addElementHelper(elementToDelete, node.DataType.FILE)
      const {
        deleteFileFromSidePannelWithOldPath
      } = fileTreeComposable

      const e: FsEvent = {
        dataType: node.DataType.FILE,
        file: elementToDelete,
        fileType: node.FileType.GRAPH,
        oldPath: elementToDelete,
        op: watcher.Op.REMOVE,
        path: "."
      }
      deleteFileFromSidePannelWithOldPath(e)

      expect(rootFiles.length).toBe(0)
    })

    it('should delete even outside of root', async () => {
      const parentDirName = "test"
      const {
        fileTreeComposable
      } = await addElementHelper(parentDirName, node.DataType.DIR)
      const { dirs, addDir, addElementInSidePanel, deleteFileFromSidePannelWithOldPath } = fileTreeComposable
      addDir(parentDirName, [])

      const subDirName = "subDir"

      const e: FsEvent = {
        dataType: node.DataType.DIR,
        file: subDirName,
        fileType: node.FileType.GRAPH,
        oldPath: "",
        op: watcher.Op.CREATE,
        path: parentDirName
      }
      addElementInSidePanel(e)
      addDir(subDirName, [])

      const newName = "newName"
      e.file = newName
      e.oldPath = parentDirName + "/" + subDirName
      e.op = watcher.Op.MOVE
      e.dataType = node.DataType.DIR
      e.path = parentDirName

      deleteFileFromSidePannelWithOldPath(e)

      const parentDir = dirs.value.get(parentDirName)!
      expect(parentDir).toBeDefined()
      expect(parentDir.length).toBe(0)
    })

    it('should not remove an element that does not exists', async () => {
      const elementToDelete = "graph.json"
      const {
        rootFiles,
        fileTreeComposable
      } = await addElementHelper(elementToDelete, node.DataType.FILE)
      const { deleteFileFromSidePannelWithOldPath } = fileTreeComposable

      const fakeName = "fake.json"
      const e: FsEvent = {
        dataType: node.DataType.FILE,
        file: fakeName,
        fileType: node.FileType.GRAPH,
        oldPath: fakeName,
        op: watcher.Op.REMOVE,
        path: "."
      }
      deleteFileFromSidePannelWithOldPath(e)

      const presentFileName = elementToDelete.split(".")[0]
      expect(rootFiles.length).toBe(1)
      expect(rootFiles[0].name).toBe(presentFileName)
    })
  })

  describe('moveElementInSidePannel tests', () => {
    it('should move an element from a directory to another', async () => {
      const elementToMove = "graph.json"
      const {
        rootFiles,
        fileTreeComposable
      } = await addElementHelper(elementToMove, node.DataType.FILE)
      const {
        dirs,
        addDir,
        addElementInSidePanel,
        moveElementInSidePannel
      } = fileTreeComposable

      const subDirName = "subDir"

      const e: FsEvent = {
        dataType: node.DataType.DIR,
        file: subDirName,
        fileType: node.FileType.GRAPH,
        oldPath: "",
        op: watcher.Op.CREATE,
        path: "."
      }
      addElementInSidePanel(e)
      addDir(subDirName, [])

      e.file = elementToMove
      e.oldPath = elementToMove
      e.op = watcher.Op.MOVE
      e.dataType = node.DataType.FILE
      e.path = subDirName
      moveElementInSidePannel(e)

      const fileName = elementToMove.split(".")[0]
      const subDir = dirs.value.get(subDirName)!
      expect(rootFiles.length).toBe(1)
      expect(rootFiles[0].name).toBe(subDirName)
      expect(subDir).toBeDefined()
      expect(subDir.length).toBe(1)
      expect(subDir[0].name).toBe(fileName)
    })

    it('should move an element from a directory to root', async () => {
      const dirToMoveFrom = "newDir"
      const {
        rootFiles,
        fileTreeComposable
      } = await addElementHelper(dirToMoveFrom, node.DataType.DIR)
      const {
        dirs,
        addDir,
        addElementInSidePanel,
        moveElementInSidePannel
      } = fileTreeComposable

      addDir(dirToMoveFrom, [])

      const fileToMove = "graph.json"
      const fileMoveName = fileToMove.split('.')[0]
      const e: FsEvent = {
        dataType: node.DataType.FILE,
        file: fileToMove,
        fileType: node.FileType.GRAPH,
        oldPath: "",
        op: watcher.Op.CREATE,
        path: dirToMoveFrom
      }
      addElementInSidePanel(e)

      e.op = watcher.Op.MOVE
      e.file = fileToMove
      e.oldPath = dirToMoveFrom + "/" + fileToMove
      e.path = "."
      moveElementInSidePannel(e)

      const subDir = dirs.value.get(dirToMoveFrom)
      expect(subDir).toBeDefined()
      expect(subDir?.length).toBe(0)

      expect(rootFiles.length).toBe(2)
      expect(rootFiles[1].name).toBe(fileMoveName)
    })
  })

  afterAll(() => {
    vi.restoreAllMocks()
  })
})