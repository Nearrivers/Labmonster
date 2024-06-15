<template>
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
      <div class="flex items-center justify-between gap-4 py-4" v-else>
        <div class="col-span-3">
          <p class="leading-7 text-muted-foreground [&:not(:first-child)]:mt-6">
            Emplacement
          </p>
          <p class="text-sm font-semibold">
            {{ dir }}
          </p>
        </div>
        <Button @click="createConfigFile(dir)">Valider</Button>
      </div>
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

const dir = ref('');

const isDialogOpen = ref(true);
async function getLabDirectory() {
  try {
    dir.value = await OpenCreateLabDialog();
  } catch (error) {}
}

async function createConfigFile(path: string) {
  try {
    await CreateAppConfig(path);
    isDialogOpen.value = false;
  } catch (error) {}
}
</script>
