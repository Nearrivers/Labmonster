<template>
  <SettingsTitle> Jeux </SettingsTitle>
  <SettingsSection v-for="game in games" :key="game.id">
    <GameSection :game="game" @delete="loadGames" @edit="onEdit" />
  </SettingsSection>
  <GameForm @submit="loadGames" v-model="isEditing" :game="gameToEdit" />
  <div
    v-if="games.length === 0 && !isEditing"
    class="flex items-center justify-between py-4 text-sm"
  >
    <p>Aucun jeu n'est disponible</p>
    <Button variant="outline" @click="editNewGame"> Ajouter </Button>
  </div>
  <Button
    v-if="!isEditing && games.length > 0"
    variant="outline"
    @click="editNewGame"
    class="mt-4"
  >
    Ajouter un nouveau jeu
  </Button>
</template>

<script lang="ts">
export default {
  name: 'Jeux',
};
</script>

<script lang="ts" setup>
import { nextTick, onMounted, ref } from 'vue';
import SettingsTitle from './ui/SettingsTitle.vue';
import { Game } from '@/types/models/Game';
import { ListGames } from '$/games/GameRepository';
import SettingsSection from './ui/SettingsSection.vue';
import Button from '../ui/button/Button.vue';
import GameForm from './game/GameForm.vue';
import GameSection from './game/GameSection.vue';

const games = ref<Game[]>([]);
const isEditing = ref(false);
const gameToEdit = ref<Game>({ id: 0, name: '', iconpath: '' });

onMounted(async () => {
  await loadGames();
});

async function loadGames() {
  games.value = await ListGames();

  if (!games.value) {
    games.value = [];
  }
  isEditing.value = false;
}

async function editNewGame() {
  gameToEdit.value = { id: 0, name: '', iconpath: '' };
  await nextTick();
  isEditing.value = true;
}

function onEdit(game: Game) {
  gameToEdit.value = game;
  isEditing.value = true;
}
</script>
