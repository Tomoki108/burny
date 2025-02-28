<template>
    <Sidebar />

    <main class="main-content">
        <header class="app-header">
            <h1>{{ title }}</h1>
        </header>
        <hr />
        <slot></slot>
    </main>
</template>

<script setup lang="ts">
import { defineProps } from 'vue'
import Sidebar from './Sidebar.vue'

defineProps<{
    title: string
}>()
</script>

<style scoped>
.app-header h1 {
    margin-bottom: 20px;
    text-align: left;
    color: var(--color-text);
    font-weight: bold;
}

.main-content {
    width: 90%;
    padding: 40px 40px 40px;
}

hr {
    border-top: 2px solid var(--color-muted);
    margin-bottom: 60px;
}
</style>