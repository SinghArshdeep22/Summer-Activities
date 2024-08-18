<script setup>
import {computed, ref, watch} from 'vue'
import Title from './components/Title.vue'; 

const testo = ref("")
function changeText() {
  testo.value = 'Changed Text';
};
const comp = computed(() => "--> " + testo.value + " <--");
const wat = watch(
  () => testo.value, 
  (newValue, oldValue) => {
    console.log("newValue: ", newValue);
    console.log("oldValue: ", oldValue);
  });
</script>

<template>
<div>{{ testo }}</div>
<div>{{ comp }}</div>
<input type="text" v-model="testo">
<button @click="changeText">Click me!</button>
<TitleVue :title="'Hello World'"></TitleVue>
</template>