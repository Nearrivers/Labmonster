<script setup lang="ts">
import { type HTMLAttributes, computed } from 'vue';
import {
  ScrollAreaCorner,
  ScrollAreaRoot,
  type ScrollAreaRootProps,
  ScrollAreaViewport,
} from 'radix-vue';
import ScrollBar from './ScrollBar.vue';
import { cn } from '@/lib/utils';

const props = defineProps<
  ScrollAreaRootProps & { class?: HTMLAttributes['class']; maxHeight?: string }
>();

const delegatedProps = computed(() => {
  const { class: _, ...delegated } = props;

  return delegated;
});
</script>

<template>
  <ScrollAreaRoot
    v-bind="delegatedProps"
    :class="cn('relative overflow-hidden', props.class)"
  >
    <ScrollAreaViewport
      v-if="!maxHeight"
      class="h-full w-full rounded-[inherit]"
    >
      <slot />
    </ScrollAreaViewport>
    <ScrollAreaViewport
      v-else
      class="w-full rounded-[inherit]"
      :class="maxHeight"
    >
      <slot />
    </ScrollAreaViewport>
    <ScrollBar />
    <ScrollAreaCorner />
  </ScrollAreaRoot>
</template>
