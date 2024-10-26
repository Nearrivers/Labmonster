<template>
  <Dialog v-model:open="isDialogOpen">
    <DialogContent class="max-w-xl">
      <DialogHeader>
        <DialogTitle v-if="game">Modifier un jeu</DialogTitle>
        <DialogTitle v-else>Ajouter un jeu</DialogTitle>
        <DialogDescription>
          Créez ou modifiez votre jeu. Cliquez sur "Valider" une fois que vous
          avez fini
        </DialogDescription>
      </DialogHeader>
      <form @submit.prevent="onSubmit">
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
          @file-choosen="onFileChosen"
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
          <DialogClose as-child>
            <Button type="button" variant="outline"> Annuler </Button>
          </DialogClose>
          <Button type="submit">Valider</Button>
        </footer>
      </form>
    </DialogContent>
  </Dialog>
</template>

<script lang="ts" setup>
import { AddGame, UpdateGame } from '$/games/GameRepository';
import { repository } from '$/models';
import Button from '@/components/ui/button/Button.vue';
import Input from '@/components/ui/input/Input.vue';
import Label from '@/components/ui/label/Label.vue';
import { ref, computed, watchEffect, watch } from 'vue';
import SelectFileInput from '../ui/SelectFileInput.vue';
import AvatarImage from '@/components/ui/avatar/AvatarImage.vue';
import Avatar from '@/components/ui/avatar/Avatar.vue';
import AvatarFallback from '@/components/ui/avatar/AvatarFallback.vue';
import { Gamepad2 } from 'lucide-vue-next';
import Dialog from '@/components/ui/dialog/Dialog.vue';
import DialogContent from '@/components/ui/dialog/DialogContent.vue';
import DialogHeader from '@/components/ui/dialog/DialogHeader.vue';
import DialogDescription from '@/components/ui/dialog/DialogDescription.vue';
import { Game } from '@/types/models/Game';
import DialogTitle from '@/components/ui/dialog/DialogTitle.vue';
import DialogClose from '@/components/ui/dialog/DialogClose.vue';
import { useShowErrorToast } from '@/composables/useShowErrorToast';

const props = defineProps<{
  game: Game;
}>();
const emits = defineEmits<{
  (e: 'submit'): void;
}>();
const isDialogOpen = defineModel<boolean>();

const src = ref('');
const isFormSubmitted = ref(false);
const { showToast } = useShowErrorToast();
const newGameParams = ref<repository.AddGameParams>(
  props.game
    ? { name: props.game.name, iconpath: props.game.iconpath }
    : { name: '', iconpath: '' },
);

const isFormInvalid = computed(
  () => isFormSubmitted.value && !newGameParams.value.name.trim(),
);

watch(props, () => {
  newGameParams.value.name = props.game.name;

  // If the path is a valid url, we set it as the image src attribute
  if (checkIconpathIsValidUrl()) {
    newGameParams.value.iconpath = props.game.iconpath;
    return;
  }

  var base64regex =
    /^([0-9a-zA-Z+/]{4})*(([0-9a-zA-Z+/]{2}==)|([0-9a-zA-Z+/]{3}=))?$/;
  if (base64regex.test(props.game.iconpath) && props.game.iconpath) {
    newGameParams.value.iconpath = 'Icone de ' + newGameParams.value.name;
    src.value = props.game.iconpath;
    return;
  }

  newGameParams.value.iconpath = props.game.iconpath;
  src.value = '';
});

watchEffect(() => {
  checkIconpathIsValidUrl();
});

function checkIconpathIsValidUrl(): boolean {
  try {
    // Checking if the url is a valid one. URL constructor throws otherwise
    new URL(newGameParams.value.iconpath);
    src.value = newGameParams.value.iconpath;
    return true;
  } catch (error) {
    // If the function throws, its fine
    return false;
  }
}

async function onSubmit() {
  isFormSubmitted.value = true;
  if (isFormInvalid.value) {
    return;
  }

  const game = {
    name: newGameParams.value.name,
    iconpath: src.value ? src.value : newGameParams.value.iconpath,
  };

  if (props.game.name) {
    try {
      await UpdateGame({
        id: props.game.id,
        ...game,
      });
      isFormSubmitted.value = false;
      newGameParams.value = { name: '', iconpath: '' };
      src.value = '';
      emits('submit');
    } catch (error) {
      showToast(error);
    } finally {
      return;
    }
  }

  try {
    await AddGame(game);
    isFormSubmitted.value = false;
    newGameParams.value.name = '';
    newGameParams.value.iconpath = '';
    src.value = '';
    emits('submit');
  } catch (error) {
    showToast(error);
  }
}

function onFileChosen(base64File: string) {
  src.value = base64File;
}
</script>
