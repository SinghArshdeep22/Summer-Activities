<script setup>
import { defineProps } from 'vue';
const props = defineProps({
    title: {
        type: String,
        required: true,
    },
});
</script>

<template>
    <h1>{{title}}</h1>
</template>