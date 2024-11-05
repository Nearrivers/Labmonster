<template>
  <div class="flex items-center gap-4">
    <Avatar>
      <AvatarImage :src="game.iconpath" :alt="'Icone de ' + game.name" />
      <AvatarFallback>
        <Gamepad2 />
      </AvatarFallback>
    </Avatar>
    <p>
      {{ game.name }}
    </p>

    <div class="ml-auto flex gap-2 text-muted-foreground">
      <TopButton :additionnal-classes="'!p-0'" :tooltip-side="'bottom'">
        <template #icon>
          <Settings class="h-6 w-6 p-1" />
        </template>
        <template #tooltip>Options</template>
      </TopButton>
      <TopButton
        :additionnal-classes="'!p-0'"
        :tooltip-side="'bottom'"
        @click="emit('edit', game)"
      >
        <template #icon>
          <Pencil class="h-6 w-6 p-1" />
        </template>
        <template #tooltip>Modifier</template>
      </TopButton>
      <TopButton
        :additionnal-classes="'!p-0'"
        :tooltip-side="'bottom'"
        @click="removeGame"
      >
        <template #icon>
          <Trash class="h-6 w-6 p-1" />
        </template>
        <template #tooltip>Supprimer</template>
      </TopButton>
    </div>
  </div>
</template>

<script setup lang="ts">
import { DeleteGame } from '$/games/GameRepository';
import Avatar from '@/components/ui/avatar/Avatar.vue';
import AvatarFallback from '@/components/ui/avatar/AvatarFallback.vue';
import AvatarImage from '@/components/ui/avatar/AvatarImage.vue';
import TopButton from '@/components/ui/TopButton.vue';
import { Game } from '@/types/models/Game';
import { Gamepad2, Pencil, Settings, Trash } from 'lucide-vue-next';

const props = defineProps<{
  game: Game;
}>();

const emit = defineEmits<{
  (e: 'edit', game: Game): void;
  (e: 'delete'): void;
}>();

async function removeGame() {
  try {
    await DeleteGame(props.game.id);
    emit('delete');
  } catch (error) {}
}
</script>
