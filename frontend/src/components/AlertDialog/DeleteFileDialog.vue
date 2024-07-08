<template>
  <AlertDialog :open="isDialogOpen">
    <AlertDialogContent>
      <AlertDialogHeader>
        <AlertDialogTitle>Supprimer le fichier</AlertDialogTitle>
        <AlertDialogDescription>
          Confirmer la suppression de "{{ fileTitle }}"
        </AlertDialogDescription>
      </AlertDialogHeader>
      <AlertDialogFooter>
        <AlertDialogAction :class="'bg-red-600 text-white hover:bg-red-500'">
          <p>Supprimer</p>
        </AlertDialogAction>
        <AlertDialogCancel @click.prevent="closeDialog">
          Annuler
        </AlertDialogCancel>
      </AlertDialogFooter>
    </AlertDialogContent>
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
import { ref } from 'vue';

const isDialogOpen = ref(false);
const fileTitle = ref('');

function openDialog(filename: string) {
  fileTitle.value = filename;
  isDialogOpen.value = true;
}

function closeDialog() {
  isDialogOpen.value = false;
}

defineExpose({ openDialog, closeDialog });
</script>
