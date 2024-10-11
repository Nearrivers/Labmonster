import { useShowErrorToast } from "./useShowErrorToast";
import { SaveMedia } from "$/file_handler/FileHandler";

export function useScreenRecorder() {
  const chunks: Blob[] = []
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

  async function startScreenRecording() {
    const stream = await navigator.mediaDevices.getDisplayMedia({
      video: true,
      audio: true,
    });

    const mediaRecorder = new MediaRecorder(stream, {
      mimeType: mime,
    });

    const track = stream.getVideoTracks()[0];
    await track.applyConstraints(trackConstraints);

    const audioTracks = stream.getAudioTracks()
    console.log(audioTracks[0])

    mediaRecorder.addEventListener('dataavailable', pipeStreamData);
    mediaRecorder.addEventListener('stop', outputStreamIntoFile);
    mediaRecorder.start();
  }

  function pipeStreamData(e: BlobEvent) {
    chunks.push(e.data)
  }

  function outputStreamIntoFile() {
    const blob = new Blob(chunks, {
      type: 'video/webm',
    });

    const reader = new FileReader();
    reader.onload = async function (e) {
      try {
        await SaveMedia('root', 'video/webm', e.target?.result as string);
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
    startScreenRecording
  }
}