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
</template>

<script lang="ts" setup>
    import {reactive} from "vue"
    import {Greet} from "../../wailsjs/go/main/App"
    import {Button} from "@/components/ui/button"

    const data = reactive({
        name: "",
        resultText: "Please enter your name below test 👇",
    })

    function greet() {
        Greet(data.name).then((result) => {
            data.resultText = result
        })
    }
</script>
