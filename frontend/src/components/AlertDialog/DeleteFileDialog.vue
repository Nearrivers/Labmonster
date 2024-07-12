<template>
  <AlertDialog asChild :open="isDialogOpen">
    <div data-test="dialog">
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>Supprimer le fichier</AlertDialogTitle>
          <AlertDialogDescription asChild>
            <p data-test="description">
              Confirmer la suppression de "{{ fileTitle }}"
            </p>
          </AlertDialogDescription>
        </AlertDialogHeader>
        <AlertDialogFooter>
          <AlertDialogAction
            :class="'bg-red-600 text-white hover:bg-red-500'"
            @click="onDeleteFile"
          >
            <p>Supprimer</p>
          </AlertDialogAction>
          <AlertDialogCancel @click.prevent="closeDialog">
            Annuler
          </AlertDialogCancel>
        </AlertDialogFooter>
      </AlertDialogContent>
    </div>
  </AlertDialog>
</template>

<script lang="ts" setup>
import { DeleteFile } from '$/filetree/FileTreeExplorer';
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from '@/components/ui/alert-dialog';
import { h, ref } from 'vue';
import { ToastAction, useToast } from '../ui/toast';

const props = defineProps<{
  path?: string;
}>();

const isDialogOpen = ref(false);
const fileTitle = ref('');
const { toast } = useToast();

function openDialog(filename: string) {
  fileTitle.value = filename;
  isDialogOpen.value = true;
}

function closeDialog() {
  isDialogOpen.value = false;
}

async function onDeleteFile() {
  try {
    await DeleteFile(props.path!);
    isDialogOpen.value = false;
    toast({
      description: `Fichier "${fileTitle.value}" supprimé avec succès`,
      duration: 5000,
    });
    removeNode();
  } catch (error) {
    toast({
      title: 'Suppression impossible',
      description: String(error),
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
}

// Suppression du fichier visuellement
function removeNode() {
  const nodeToRemove = document.querySelector(`[data-path="${props.path}"]`);
  if (nodeToRemove) {
    nodeToRemove.remove();
  }
}

defineExpose({ openDialog, closeDialog });
</script>
