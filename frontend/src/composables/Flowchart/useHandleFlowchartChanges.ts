import { SaveFile } from '$/filetree/FileTree';
import { graph } from '$/models';
import { NodeChange, useVueFlow } from '@vue-flow/core';
import { useShowErrorToast } from '../useShowErrorToast';
import { Ref, ref } from 'vue';
import { onBeforeRouteUpdate } from 'vue-router';

export function useHandleFlowchartChanges(pathFromLabRoot: Ref<string>) {
  const isSaving = ref(false);
  const { showToast } = useShowErrorToast();
  const { onNodesChange, onEdgesChange, onViewportChangeEnd, toObject } =
    useVueFlow();

  async function Save() {
    try {
      isSaving.value = true;
      const graph = toObject() as unknown as graph.Graph;
      await SaveFile(pathFromLabRoot.value, graph);
    } catch (error) {
      showToast(error);
    } finally {
      isSaving.value = false;
    }
  }

  onNodesChange(async (param: NodeChange[]) => {
    await Save();
  });
  onEdgesChange(async (_) => {
    await Save();
  });
  onViewportChangeEnd(async (_) => {
    await Save();
  });

  onBeforeRouteUpdate(async () => {
    await Save();
  });

  return {
    isSaving,
  };
}
