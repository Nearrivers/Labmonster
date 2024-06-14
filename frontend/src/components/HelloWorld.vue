<template>
  <main>
    <div id="result" class="result">{{ data.resultText }}</div>
    <div id="input" class="input-box">
      <input
        id="name"
        v-model="data.name"
        autocomplete="off"
        class="input"
        type="text"
      />
      <Button class="btn" @click="greet">Button</Button>
    </div>
  </main>
  <TestFlow></TestFlow>
</template>

<script lang="ts" setup>
import { reactive } from 'vue';
import { Greet } from '../../wailsjs/go/main/App';
import { Button } from '@/components/ui/button';
import TestFlow from '@/components/TestFlow.vue';

const data = reactive({
  name: '',
  resultText: 'Please enter your name below test 👇',
});

function greet() {
  Greet(data.name).then((result) => {
    data.resultText = result;
  });
}
</script>
