import { FlowExportObject, useVueFlow } from "@vue-flow/core";
import { useEventListener } from "../useEventListener";
import { OpenFile, SaveMedia } from "$/filetree/FileTree";
import { CustomNodeData } from "@/types/CustomNodeData";
import { useRoute } from "vue-router";
import { computed, ref, watch } from "vue";
import { useShowErrorToast } from "../useShowErrorToast";

export function useFlowChart() {
  const path = ref('')
  const route = useRoute()
  const { updateNode, fromObject } = useVueFlow();
  const { showToast } = useShowErrorToast();
  useEventListener(window, 'paste', onPaste);

  watch(
    () => route.params.path,
    async () => {
      path.value = route.params.path as string;
      await loadGraph();
    },
    { immediate: true },
  );

  const fileName = computed(() =>
    route.params.path.slice(0, route.params.path.indexOf('.')),
  );

  async function loadGraph() {
    try {
      const path = route.params.path as string;
      const graph = await OpenFile(path);
      fromObject(graph as unknown as FlowExportObject);
    } catch (error) {
      showToast(error);
    }
  }

  async function onPaste(e: ClipboardEvent) {
    const id = (e.target as HTMLInputElement).id;

    if (!e.clipboardData) {
      return;
    }

    if (!e.clipboardData.files || e.clipboardData.files.length === 0) {
      return
    }

    const file = e.clipboardData.files[0];
    const mimeType = e.clipboardData.files[0].type;

    if (
      file.type.startsWith('image/')
    ) {
      handleImagePaste(id, mimeType, file)
      return
    }

    if (
      file.type.startsWith('video/')
    ) {
      const reader = new FileReader();
      reader.onload = async function (e) {
        try {
          const imagePath = await SaveMedia(
            path.value,
            mimeType,
            e.target?.result as string,
          );
          updateNode<CustomNodeData>(id, {
            type: 'video',
            data: {
              hasFrameDataSection: false,
              image: imagePath,
              text: '',
            },
          });
        } catch (error) {
          showToast(error);
        }
      };
      reader.onerror = function (e) {
        showToast(e.target?.error);
      };
      reader.readAsDataURL(file);
    }
  }

  function handleImagePaste(id: string, mimeType: string, file: File) {
    const reader = new FileReader();
    reader.onload = async function (e) {
      try {
        const imagePath = await SaveMedia(
          path.value,
          mimeType,
          e.target?.result as string,
        );
        updateNode<CustomNodeData>(id, {
          type: 'image',
          data: {
            hasFrameDataSection: false,
            image: imagePath,
            text: '',
          },
        });
      } catch (error) {
        showToast(error);
      }
    };
    reader.onerror = function (e) {
      showToast(e.target?.error);
    };
    reader.readAsDataURL(file);
  }

  return {
    path,
    fileName
  }
}