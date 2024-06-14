<template>
  <RouterView></RouterView>
  <AlertDialog :open="isDialogOpen">
    <AlertDialogContent>
      <AlertDialogHeader class="!text-center">
        <AlertDialogTitle>LabMonster</AlertDialogTitle>
        <AlertDialogDescription>Version 0.0.1</AlertDialogDescription>
      </AlertDialogHeader>
      <div class="grid gap-4 py-4" v-if="!dir">
        <div class="grid grid-cols-4 items-center gap-4">
          <div class="col-span-3">
            <p class="leading-7 [&:not(:first-child)]:mt-6">
              Créer un nouveau Lab
            </p>
            <p class="text-sm text-muted-foreground">
              Crée un nouveau "Lab" dans un dossier
            </p>
          </div>
          <Button @click="getLabDirectory">Créer</Button>
        </div>
        <div class="grid grid-cols-4 items-center gap-4">
          <div class="col-span-3">
            <p class="leading-7 [&:not(:first-child)]:mt-6">
              Ouvrir un dossier comme lab
            </p>
            <p class="text-sm text-muted-foreground">
              Défini un dossier comme "Lab"
            </p>
          </div>
          <Button variant="outline">Ouvrir</Button>
        </div>
      </div>
      <div class="flex items-center gap-4 py-4" v-else>
        <div class="col-span-3">
          <p class="leading-7 [&:not(:first-child)]:mt-6">Emplacement</p>
          <p class="text-sm">
            <span class="text-muted-foreground">
              Votre nouveau Lab sera placé dans le dossier
            </span>
            <span class="font-semibold">{{ dir }}</span>
          </p>
        </div>
        <div class="col-span-1 flex justify-center">
          <Button>Valider</Button>
        </div>
      </div>
    </AlertDialogContent>
  </AlertDialog>
</template>

<script lang="ts" setup>
import { useColorMode } from '@vueuse/core';
import { Button } from '@/components/ui/button';
import {
  AlertDialog,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogHeader,
  AlertDialogTitle,
} from '@/components/ui/alert-dialog';
import { ref } from 'vue';
import { OpenCreateLabDialog } from '$/config/AppConfig';

const isDialogOpen = ref(true);
const dir = ref('');

async function getLabDirectory() {
  try {
    dir.value = await OpenCreateLabDialog();
  } catch (error) {}
}

const mode = useColorMode();
</script>
