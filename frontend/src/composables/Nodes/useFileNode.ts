import { RenameFile } from '$/file_handler/FileHandler';
import { ref, computed, Ref, reactive } from 'vue';
import { useShowErrorToast } from '../useShowErrorToast';
import { useRoute } from 'vue-router';
import { node } from '$/models';

export function useFileNode(props: Ref<{ fileNode: node.Node; path: string, offset?: number }>) {
  const route = useRoute();
  const { showToast } = useShowErrorToast();
  const fileName = ref(props.value.fileNode.name);
  const input = ref<HTMLInputElement | null>(null);
  const ext = computed(() =>
    props.value.fileNode.extension.replace('.', '').toLocaleUpperCase(),
  );

  const nodeStyle = reactive({
    paddingLeft: `${props.value.offset}px`
  })

  const nodePath = ref(
    props.value.path
      ? props.value.path + '/' + props.value.fileNode.name
      : props.value.fileNode.name,
  );


  const isActive = computed(
    () =>
      route.params.path &&
      decodeURI(route.params.path as string) ===
      nodePath.value + props.value.fileNode.extension,
  );

  const nodePathWithoutSpaces = computed(() =>
    nodePath.value.replaceAll(' ', '-') + "-file",
  );

  const updatedAt = computed(() => {
    const date = new Date(props.value.fileNode.updatedAt);
    return `${date.toLocaleDateString()} à ${date.toLocaleTimeString()}`;
  });

  function selectInput(e: KeyboardEvent) {
    if (!input.value || e.key != 'F2') {
      return;
    }

    input.value.toggleAttribute('readonly');
    input.value.select();
    input.value.focus();
  }

  async function onBlur() {
    if (!input.value) {
      showToast('Input introuvable');
      return;
    }

    if (input.value.readOnly) {
      return;
    }

    input.value.toggleAttribute('readonly');
    input.value.classList.add('cursor-default');
    input.value.classList.remove('cursor-text');

    // We don't rename 
    const newName = fileName.value.trim()
      ? fileName.value + props.value.fileNode.extension
      : props.value.fileNode.name + props.value.fileNode.extension

    try {
      await RenameFile(
        props.value.path,
        props.value.fileNode.name + props.value.fileNode.extension,
        newName,
      );
    } catch (error) {
      showToast(error);
    } finally {
      if (fileName.value.trim()) {
        return
      }

      input.value.value = props.value.fileNode.name
    }
  }

  return {
    nodePath,
    nodeStyle,
    nodePathWithoutSpaces,
    input,
    onBlur,
    fileName,
    ext,
    updatedAt,
    selectInput,
    isActive,
  };
}
