<template>
  <div class="flex h-full w-full flex-col">
    <AppHeader />
    <div class="h-full w-full flex-grow bg-secondary p-2 pl-1 pt-0">
      <div class="h-full w-full rounded-md bg-background">
        <RouterView v-slot="{ Component }">
          <component :is="Component" />
        </RouterView>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import AppHeader from '@/components/AppHeader.vue';
import { RouterView, useRoute, useRouter } from 'vue-router';
import { EventsOn } from '../../wailsjs/runtime/runtime';
import { node, watcher } from '$/models';
import { FsEvent } from '@/types/FsEvent';
import { Routes } from '@/types/Routes';

const route = useRoute();
const router = useRouter();
EventsOn('fsop', onFsEvent);

function onFsEvent(e: FsEvent) {
  if (e.dataType === node.DataType.DIR) {
    return;
  }

  const routePath = route.params.path as string;

  let filePath = e.path + '/' + e.file;
  if (e.path === '.') {
    filePath = e.file;
  }

  if (
    // If the operation is a delete but the deleted file is not the one currently opened, we skip
    (filePath != routePath && e.op === watcher.Op.REMOVE) || // OR
    // If the operation is a move or a rename and the old path is different from the one of the file currently opened, we skip
    (e.oldPath != routePath.replace('./', '') &&
      (e.op === watcher.Op.MOVE || e.op === watcher.Op.RENAME))
  ) {
    return;
  }

  if (e.op === watcher.Op.REMOVE) {
    router.push({ name: Routes.NotOpened });
    return;
  }

  router.push({ name: Routes.Flowchart, params: { path: filePath } });
}
</script>
