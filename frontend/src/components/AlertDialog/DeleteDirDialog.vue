<template>
  <AlertDialog :open="isDialogOpen">
    <div data-test="dialog">
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>Supprimer le dossier</AlertDialogTitle>
          <AlertDialogDescription asChild>
            <p :id="'delete-' + fileTitle">
              Confirmer la suppression de "{{ fileTitle }}"
            </p>
            <p class="text-sm text-red-600">
              Attention: Ce dossier n'est pas vide.
            </p>
            <p class="text-sm text-red-600">
              Tous ses fichiers et dossiers seront supprimés avec lui
            </p>
          </AlertDialogDescription>
        </AlertDialogHeader>
        <AlertDialogFooter :class="'items-center !justify-start'">
          <div class="mr-auto flex items-center gap-2">
            <Checkbox id="never-ask" />
            <label for="never-ask" class="text-sm">Ne plus demander</label>
          </div>
          <AlertDialogAction
            :class="'bg-red-600 text-white hover:bg-red-500'"
            @click="onDeleteFile"
          >
            <p>Supprimer</p>
          </AlertDialogAction>
          <AlertDialogCancel @click.prevent="closeDialog" data-test="cancel">
            Annuler
          </AlertDialogCancel>
        </AlertDialogFooter>
      </AlertDialogContent>
    </div>
  </AlertDialog>
</template>

<script lang="ts" setup>
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
import Checkbox from '../ui/checkbox/Checkbox.vue';
import { DeleteDirectory } from '$/dirhandler/DirHandler';

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
    await DeleteDirectory(props.path!);
    toast({
      description: `Dossier "${fileTitle.value}" supprimé avec succès`,
      duration: 5000,
    });
  } catch (error) {
    toast({
      title: 'Suppression impossible',
      description: String(error),
      variant: 'destructive',
      action: h(
        ToastAction,
        {
          altText: 'Réessayer',
          onClick: async () => await onDeleteFile(),
        },
        {
          default: () => 'Réessayer',
        },
      ),
    });
  } finally {
    isDialogOpen.value = false;
  }
}

defineExpose({ openDialog, closeDialog });
</script>
