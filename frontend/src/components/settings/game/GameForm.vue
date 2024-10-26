<template>
  <form
    class="mt-4 rounded-md border border-border p-4"
    @submit.prevent="onSubmit"
  >
    <h2 class="mb-2 border-b border-b-border pb-2 text-lg font-bold">
      Ajouter un nouveau jeu
    </h2>
    <Label class="block py-1.5" :class="{ 'text-red-500': isFormInvalid }">
      <span class="text-red-500">*</span> Nom
    </Label>
    <Input v-model="newGameParams.name" />
    <p class="py-1.5 text-sm opacity-65">
      Nom du jeu que vous souhaitez ajouter
    </p>
    <p class="text-sm text-red-500" v-if="isFormInvalid">Champ requis</p>
    <SelectFileInput
      v-model="newGameParams.iconpath"
      @file-choosen="onFileChoosen"
    >
      <template #default>Icone</template>
      <template #description>
        Icone du jeu. Vous pouvez choisir un lien ou un fichier depuis votre
        ordinateur
      </template>
    </SelectFileInput>
    <!-- Résultat -->
    <section class="flex items-center gap-4 py-2">
      <span class="text-sm text-muted-foreground"> Résultat: </span>
      <Avatar>
        <AvatarImage :src="src" alt="Icone du jeu" />
        <AvatarFallback>
          <Gamepad2 />
        </AvatarFallback>
      </Avatar>
      <p class="font-medium">
        {{ newGameParams.name }}
      </p>
    </section>
    <footer class="mt-2 flex justify-end gap-6">
      <Button type="button" variant="outline" @click="emits('cancel')">
        Annuler
      </Button>
      <Button type="submit">Valider</Button>
    </footer>
  </form>
</template>

<script lang="ts" setup>
import { AddGame } from '$/games/GameRepository';
import { repository } from '$/models';
import Button from '@/components/ui/button/Button.vue';
import Input from '@/components/ui/input/Input.vue';
import Label from '@/components/ui/label/Label.vue';
import { ref, computed, watchEffect } from 'vue';
import SelectFileInput from '../ui/SelectFileInput.vue';
import AvatarImage from '@/components/ui/avatar/AvatarImage.vue';
import Avatar from '@/components/ui/avatar/Avatar.vue';
import AvatarFallback from '@/components/ui/avatar/AvatarFallback.vue';
import { Gamepad2 } from 'lucide-vue-next';

const emits = defineEmits<{
  (e: 'submit'): void;
  (e: 'cancel'): void;
}>();

const newGameParams = ref<repository.AddGameParams>({ name: '', iconpath: '' });
const isFormSubmitted = ref(false);
const src = ref('');

const isFormInvalid = computed(
  () => isFormSubmitted.value && !newGameParams.value.name.trim(),
);

watchEffect(() => {
  try {
    // Checking if the url is a valid one. URL constructor throws otherwise
    new URL(newGameParams.value.iconpath);
    src.value = newGameParams.value.iconpath;
  } catch (error) {
    // If the function throws, its fine
  }
});

async function onSubmit() {
  isFormSubmitted.value = true;
  if (isFormInvalid.value) {
    return;
  }
  await AddGame({
    name: newGameParams.value.name,
    iconpath: src.value ? src.value : newGameParams.value.iconpath,
  });
  isFormSubmitted.value = false;
  emits('submit');
}

function onFileChoosen(base64File: string) {
  src.value = base64File;
}
</script>
