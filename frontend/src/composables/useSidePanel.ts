import { filetree } from "$/models";
import { ToastAction, useToast } from "@/components/ui/toast";
import { h } from "vue";

export function useSidePanel() {
  const { toast } = useToast()

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

  function showToast(description: string, title?: string) {
    toast({
      title,
      description,
      variant: 'destructive',
      action: h(
        ToastAction,
        {
          altText: 'Réessayer',
          onClick: () => location.reload(),
        },
        {
          default: () => 'Réessayer',
        },
      ),
    });
  }

  return {
    sortNodes,
    showToast
  }
}