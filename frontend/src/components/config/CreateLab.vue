<template>
  <AlertDialog :open="isDialogOpen">
    <AlertDialogContent>
      <AlertDialogHeader class="!text-center">
        <AlertDialogTitle>Labmonster</AlertDialogTitle>
        <AlertDialogDescription>Version 0.1.0</AlertDialogDescription>
      </AlertDialogHeader>
      <div class="grid gap-4 pt-4" v-if="!dir">
        <div class="grid grid-cols-4 items-center gap-4">
          <div class="col-span-3">
            <p class="leading-7 [&:not(:first-child)]:mt-6">
              Créer un nouveau Lab
            </p>
            <p class="text-xs text-muted-foreground">
              Sélectionnez un emplacement où mettre vos fichiers
            </p>
          </div>
          <Button @click="getLabDirectory"> Parcourir </Button>
        </div>
      </div>
      <div class="flex items-center justify-between gap-4 pt-4" v-else>
        <div>
          <p class="leading-7 text-muted-foreground [&:not(:first-child)]:mt-6">
            Emplacement
          </p>
          <p class="text-sm font-semibold">
            {{ dir }}
          </p>
        </div>
        <div class="flex gap-2">
          <Button @click="getLabDirectory" variant="outline">
            Parcourir
          </Button>
          <Button @click="createConfigFile(dir)">Valider</Button>
        </div>
      </div>
      <footer class="mt-4 flex w-full justify-center text-sm">
        <button @click="Quit()" class="hover:underline">
          Quitter l'application
        </button>
      </footer>
    </AlertDialogContent>
  </AlertDialog>
</template>

<script setup lang="ts">
import { Button } from '@/components/ui/button';
import {
  AlertDialog,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogHeader,
  AlertDialogTitle,
} from '@/components/ui/alert-dialog';
import { CreateAppConfig } from '$/config/AppConfig';
import { ref } from 'vue';
import { OpenCreateLabDialog } from '$/config/AppConfig';
import { useShowErrorToast } from '@/composables/useShowErrorToast';
import { configFileLoaded } from '@/events/ReloadFileExplorer';
import { Quit } from '../../../wailsjs/runtime/runtime';

const dir = ref('');
const { showToast } = useShowErrorToast();
const isDialogOpen = ref(true);

async function getLabDirectory() {
  try {
    const path = await OpenCreateLabDialog();

    if (!dir.value) {
      dir.value = path;
    }
  } catch (error) {
    showToast(error);
  }
}

async function createConfigFile(path: string) {
  try {
    await CreateAppConfig(path);
    isDialogOpen.value = false;
    configFileLoaded.configFileLoaded();
  } catch (error) {
    showToast(error);
  }
}
</script>
