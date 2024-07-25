import { ToastAction, useToast } from "@/components/ui/toast";
import { h } from "vue";

export function useShowErrorToast() {
  const { toast } = useToast()

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
    showToast
  }
}