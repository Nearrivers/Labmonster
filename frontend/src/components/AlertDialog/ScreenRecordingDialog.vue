<template>
  <AlertDialog :open="isDialogOpen">
    <div data-test="dialog">
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>
            Enregistrement de l'écran en cours
          </AlertDialogTitle>
          <AlertDialogDescription asChild>
            <p>
              Cliquez sur "Sauvegarder" pour l'arrêter et enregistrer la vidéo
            </p>
          </AlertDialogDescription>
        </AlertDialogHeader>
        <AlertDialogFooter :class="'items-center'">
          <AlertDialogCancel
            @click="emit('cancelRecording')"
            data-test="cancel"
          >
            Annuler
          </AlertDialogCancel>
          <AlertDialogAction @click="emit('stopRecording')">
            <p>Sauvegarder</p>
          </AlertDialogAction>
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
import { ref, watchEffect } from 'vue';

const props = defineProps<{
  recorderState: RecordingState;
}>();

const emit = defineEmits<{
  (e: 'stopRecording'): void;
  (e: 'cancelRecording'): void;
}>();
const isDialogOpen = ref(true);

watchEffect(() => {
  if (props.recorderState !== 'inactive' && !isDialogOpen.value) {
    openDialog();
    return;
  }

  if (props.recorderState === 'inactive' && isDialogOpen.value) {
    closeDialog();
  }
});

function openDialog() {
  isDialogOpen.value = true;
}

function closeDialog() {
  isDialogOpen.value = false;
}

defineExpose({ openDialog, closeDialog });
</script>
