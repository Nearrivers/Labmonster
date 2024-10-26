<template>
  <Dialog v-model:open="isDialogOpen">
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
import ScrollArea from '../ui/scroll-area/ScrollArea.vue';
import { useSettingsDialog } from '@/composables/Dialogs/useSettingsDialog';

const { isDialogOpen, currentTab, tabs } = useSettingsDialog(openDialog);

function openDialog() {
  isDialogOpen.value = true;
}

function closeDialog() {
  isDialogOpen.value = false;
}

defineExpose({ openDialog, closeDialog });
</script>
