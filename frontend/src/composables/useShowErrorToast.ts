import { ToastAction, useToast } from "@/components/ui/toast";
import { h } from "vue";

export function useShowErrorToast() {
  const { toast } = useToast()

  function showToast(error: unknown, title?: string) {
    let description = String(error)

    if (error instanceof Error) {
      description = error.message
    }

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
    showToast
  }
}