<template>
  <aside class="border-r border-r-border bg-secondary">
    <header class="p-1">
      <TopButton
        @click="toggleSidePanel"
        :additionnalClasses="'text-muted-foreground hover:text-accent-foreground hover:bg-accent !p-1'"
      >
        <template #icon>
          <PanelLeft class="h-5 w-5" />
        </template>
        <template #tooltip> Basculer </template>
      </TopButton>
    </header>
    <section class="p-1">
      <TopButton
        @click="startScreenRecording"
        :additionnalClasses="'text-muted-foreground hover:text-accent-foreground hover:bg-accent !p-1'"
        :tooltipSide="'right'"
      >
        <template #icon>
          <TvMinimalPlay class="h-5 w-5" />
        </template>
        <template #tooltip> Démarrer une capture d'écran </template>
      </TopButton>
    </section>
  </aside>
</template>

<script lang="ts" setup>
import { SaveMedia } from '$/file_handler/FileHandler';
import TopButton from '@/components/ui/TopButton.vue';
import { useShowErrorToast } from '@/composables/useShowErrorToast';
import { sidePanelToggled } from '@/events/ToggleSidePanel';
import { PanelLeft, TvMinimalPlay } from 'lucide-vue-next';

const { showToast } = useShowErrorToast();

function toggleSidePanel() {
  sidePanelToggled.sidePanelToggled();
}

async function startScreenRecording() {
  let stream = await navigator.mediaDevices.getDisplayMedia({
    video: true,
  });

  const mime = 'video/webm';
  const mediaRecorder = new MediaRecorder(stream, {
    mimeType: mime,
  });

  const track = stream.getVideoTracks()[0];
  await track.applyConstraints({
    noiseSuppression: true,
    width: {
      min: 640,
      ideal: 1920,
    },
    height: {
      min: 480,
      ideal: 1080,
    },
    frameRate: {
      min: 30,
      ideal: 60,
    },
  });

  const chunks: Blob[] = [];
  mediaRecorder.addEventListener('dataavailable', function (e) {
    chunks.push(e.data);
  });

  mediaRecorder.addEventListener('stop', function () {
    const blob = new Blob(chunks, {
      type: chunks[0].type,
    });

    const reader = new FileReader();
    reader.onload = async function (e) {
      console.log(e.target?.result);

      try {
        const path = await SaveMedia(
          'root',
          'video/webm',
          e.target?.result as string,
        );

        console.log(path);
      } catch (error) {
        showToast(error);
      }
    };
    reader.onerror = function (e) {
      showToast(e.target?.error);
    };

    reader.readAsDataURL(blob);
  });

  mediaRecorder.start();
}
</script>
