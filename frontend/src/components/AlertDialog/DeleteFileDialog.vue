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
        <AlertDialogAction>Supprimer</AlertDialogAction>
        <AlertDialogCancel>Annuler</AlertDialogCancel>
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

defineProps<{
  fileTitle?: string;
  isDialogOpen: boolean;
}>();
</script>
