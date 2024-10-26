<template>
  <Dialog :open="true">
    <DialogContent class="flex h-[70%] max-w-5xl gap-0 overflow-hidden p-0">
      <DialogTitle class="sr-only">Paramètres</DialogTitle>
      <DialogDescription class="sr-only">
        Gestion des paramètres
      </DialogDescription>
      <aside class="w-1/4 bg-sidebar p-2 text-sidebar-foreground">
        <h2 class="p-2.5 text-xs font-semibold opacity-65">Paramètres</h2>
        <ul class="flex flex-col gap-1 text-sm">
          <li
            v-for="[name, tab] in tabs.entries()"
            class="flex cursor-default items-center gap-2 rounded-md px-2 py-1 hover:bg-accent dark:hover:bg-sidebar-accent"
            :class="{ 'bg-accent dark:bg-sidebar-accent': currentTab === name }"
            :key="name"
            @click="currentTab = name"
          >
            <component :is="tab.icon" class="w-4"></component>
            {{ name }}
          </li>
        </ul>
      </aside>
      <ScrollArea class="h-full flex-1">
        <div class="px-12 py-6">
          <component :is="tabs.get(currentTab)!.tab"></component>
        </div>
      </ScrollArea>
    </DialogContent>
  </Dialog>
</template>

<script lang="ts" setup>
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogTitle,
} from '@/components/ui/dialog';
import { FunctionalComponent, ref } from 'vue';
import GameSettings from '../settings/GameSettings.vue';
import MediaSettings from '../settings/MediaSettings.vue';
import { Gamepad2, GitFork, MonitorPlay, Wrench } from 'lucide-vue-next';
import GeneralSettings from '../settings/GeneralSettings.vue';
import { DefinedComponent } from '@vue/test-utils/dist/types';
import GraphSettings from '../settings/GraphSettings.vue';
import ScrollArea from '../ui/scroll-area/ScrollArea.vue';

type SettingsTab = {
  icon: FunctionalComponent;
  tab: DefinedComponent;
};

const currentTab = ref('Jeux');
const tabs = new Map<string, SettingsTab>();
tabs.set(GeneralSettings.name!, { icon: Wrench, tab: GeneralSettings });
tabs.set(GameSettings.name!, { icon: Gamepad2, tab: GameSettings });
tabs.set(MediaSettings.name!, { icon: MonitorPlay, tab: MediaSettings });
tabs.set(GraphSettings.name!, { icon: GitFork, tab: GraphSettings });

const isDialogOpen = ref(true);

function openDialog() {
  isDialogOpen.value = true;
}

function closeDialog() {
  isDialogOpen.value = false;
}

defineExpose({ openDialog, closeDialog });
</script>
