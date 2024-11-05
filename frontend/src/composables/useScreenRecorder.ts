import { useShowErrorToast } from "./useShowErrorToast";
import { SaveMedia } from "$/file_handler/FileHandler";
import { onUnmounted, ref } from "vue";

export function useScreenRecorder() {
  let wasRecordCanceled = false
  let stream: MediaStream;
  const mediaRecorder = ref<MediaRecorder>()

  let chunks: Blob[] = []
  const mime = 'video/webm; codecs=h264'
  const { showToast } = useShowErrorToast();
  const trackConstraints: MediaTrackConstraints = {
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
  }

  onUnmounted(() => {
    if (!mediaRecorder.value) {
      return
    }

    mediaRecorder.value.removeEventListener('dataavailable', pipeStreamData);
    mediaRecorder.value.removeEventListener('stop', outputStreamIntoFile);
    mediaRecorder.value.addEventListener('pause', pauseRecorderState)
  })

  const mediaRecorderState = ref<RecordingState>('inactive')

  async function startScreenRecording() {
    stream = await navigator.mediaDevices.getDisplayMedia({
      video: true,
      audio: true,
    });

    mediaRecorder.value = new MediaRecorder(stream, {
      mimeType: mime,
    });

    const track = stream.getVideoTracks()[0];
    await track.applyConstraints(trackConstraints);

    chunks = []
    mediaRecorder.value.addEventListener('dataavailable', pipeStreamData);
    mediaRecorder.value.addEventListener('pause', pauseRecorderState)
    mediaRecorder.value.addEventListener('stop', outputStreamIntoFile);

    try {
      mediaRecorder.value.start();
      mediaRecorderState.value = "recording"
    } catch (error) {
      showToast("Impossible de lancer l'enregistrement")
    }
  }

  function stopScreenRecording() {
    if (mediaRecorderState.value != 'inactive') {
      stream.getTracks().forEach((track) => track.stop())
      return
    }

    showToast("Aucun enregistrement en cours")
  }

  function cancelScreenRecording() {
    wasRecordCanceled = true
    stopScreenRecording()
  }

  function pauseRecorderState() {
    mediaRecorderState.value = 'paused'
  }

  function pipeStreamData(e: BlobEvent) {
    chunks.push(e.data)
  }

  function outputStreamIntoFile() {
    if (!mediaRecorder.value || wasRecordCanceled) {
      wasRecordCanceled = false
      mediaRecorderState.value = 'inactive'
      return
    }

    const fileName = prompt("Entrez un nom pour votre vidéo", "")

    mediaRecorder.value.removeEventListener('dataavailable', pipeStreamData);
    mediaRecorder.value.removeEventListener('stop', outputStreamIntoFile);
    mediaRecorder.value.addEventListener('pause', pauseRecorderState)
    const blob = new Blob(chunks, {
      type: 'video/webm',
    });

    const reader = new FileReader();
    reader.onload = async function (e) {
      try {
        await SaveMedia(fileName!, 'root', 'video/webm', e.target?.result as string);
        mediaRecorderState.value = 'inactive'
        chunks = []
      } catch (error) {
        showToast(error);
      }
    };

    reader.onerror = function (e) {
      showToast(e.target?.error);
    };

    reader.readAsDataURL(blob);
  }

  return {
    mediaRecorderState,
    startScreenRecording,
    stopScreenRecording,
    cancelScreenRecording
  }
}