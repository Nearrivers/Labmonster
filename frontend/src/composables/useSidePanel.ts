import { filetree } from "$/models";

export function useSidePanel() {
  function sortNodes(f1: filetree.Node, f2: filetree.Node) {
    // Tri sur les types d'abord
    if (f1.type === 'DIR' && f2.type == 'FILE') {
      return -1;
    }

    if (f1.type === 'FILE' && f2.type == 'DIR') {
      return 1;
    }

    if (f1.name < f2.name) {
      return -1;
    }

    if (f1.name == f2.name) {
      return 0;
    }

    if (f1.name > f2.name) {
      return 1;
    }

    return 0;
  }


  return {
    sortNodes,
  }
}