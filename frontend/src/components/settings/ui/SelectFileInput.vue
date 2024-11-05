<template>
  <Label class="block py-1.5"><slot></slot></Label>
  <div
    class="flex rounded-sm border border-border ring-ring ring-offset-background has-[:focus-visible]:ring-2 has-[:focus]:ring-2 has-[:focus]:ring-offset-2"
  >
    <Input
      v-model="filePathModel"
      class="rounded-e-none border-y-0 border-l-0 border-r outline-none focus-visible:!ring-0 focus-visible:!ring-offset-0"
    >
    </Input>
    <TooltipProvider>
      <Tooltip>
        <TooltipTrigger tabindex="-1">
          <label for="iconpath">
            <Button
              @click.prevent="fileInput?.click()"
              type="button"
              variant="outline"
              class="rounded-none border-none bg-transparent"
            >
              <FolderOpen class="w-4" />
            </Button>
          </label>
          <input
            ref="fileInput"
            class="sr-only"
            type="file"
            name="iconpath"
            id="iconpath"
            @change="selectFile"
            tabindex="-1"
          />
        </TooltipTrigger>
        <TooltipContent :side="'top'"> Chercher un fichier </TooltipContent>
      </Tooltip>
    </TooltipProvider>
  </div>
  <p class="py-1.5 text-sm opacity-65">
    <slot name="description"></slot>
  </p>
</template>

<script lang="ts" setup>
import Button from '@/components/ui/button/Button.vue';
import Input from '@/components/ui/input/Input.vue';
import Label from '@/components/ui/label/Label.vue';
import Tooltip from '@/components/ui/tooltip/Tooltip.vue';
import TooltipContent from '@/components/ui/tooltip/TooltipContent.vue';
import TooltipProvider from '@/components/ui/tooltip/TooltipProvider.vue';
import TooltipTrigger from '@/components/ui/tooltip/TooltipTrigger.vue';
import { useShowErrorToast } from '@/composables/useShowErrorToast';
import { FolderOpen } from 'lucide-vue-next';
import { ref } from 'vue';

const emit = defineEmits<{
  (e: 'fileChoosen', base64File: string): void;
}>();
const [filePathModel, _] = defineModel<string, string>({ required: true });
const fileInput = ref<HTMLInputElement | null>(null);
const { showToast } = useShowErrorToast();

function selectFile(e: Event) {
  const file = (e.target as HTMLInputElement).files![0];

  const reader = new FileReader();
  reader.onload = async function (e) {
    filePathModel.value = file.name;
    if (e.target) {
      emit('fileChoosen', e.target.result as string);
    }
  };
  reader.onerror = function (e) {
    showToast(e.target?.error);
  };
  reader.readAsDataURL(file);
}
</script>
