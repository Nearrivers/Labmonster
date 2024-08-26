import { RenameFile } from "$/filetree/FileTree";
import { ref, computed, Ref } from "vue";
import { useShowErrorToast } from "../useShowErrorToast";
import { filetree } from "$/models";
import { useRoute } from "vue-router";

export function useFileNode(props: Ref<{ node: filetree.Node, path: string }>) {
  const route = useRoute()
  const { showToast } = useShowErrorToast();
  const fileName = ref(props.value.node.name);
  const input = ref<HTMLInputElement | null>(null);
  const ext = computed(() =>
    props.value.node.extension.replace('.', '').toLocaleUpperCase(),
  );

  const nodePath = ref(
    props.value.path ? props.value.path + '/' + props.value.node.name : props.value.node.name,
  );
  console.log(route.params.path, nodePath.value)

  const isActive = computed(() => route.params.path.includes(nodePath.value))

  const nodePathWithoutSpaces = computed(() =>
    nodePath.value.replaceAll(' ', '-'),
  );

  const updatedAt = computed(() => {
    const date = new Date(props.value.node.updatedAt);
    return `${date.toLocaleDateString()} à ${date.toLocaleTimeString()}`;
  })

  function selectInput(e: KeyboardEvent) {
    if (!input.value || e.key != "F2") {
      return
    }

    input.value.toggleAttribute('readonly')
    input.value.select()
    input.value.focus()
  }

  async function onBlur() {
    if (!input.value) {
      showToast('Input introuvable');
      return;
    }

    if (input.value.readOnly) {
      return
    }

    input.value.toggleAttribute('readonly');
    input.value.classList.add('cursor-pointer');
    input.value.classList.remove('cursor-text');

    try {
      const newName = fileName.value + props.value.node.extension
      await RenameFile(
        props.value.path,
        props.value.node.name + props.value.node.extension,
        newName
      );

    } catch (error) {
      showToast(error);
    }
  }

  return {
    nodePath,
    nodePathWithoutSpaces,
    input,
    onBlur,
    fileName,
    ext,
    updatedAt,
    selectInput,
    isActive
  }
}