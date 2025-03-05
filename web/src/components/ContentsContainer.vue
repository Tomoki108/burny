<template>
    <Sidebar />
    <main class="main-content">
        <v-alert v-if="alertCtx.show" :type="alertCtx.type" :text="alertCtx.text" closable class="alert-dialog" />
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
import type { AlertContext } from '../composables/alert_composable'

defineProps<{
    title: string
    alertCtx: AlertContext,
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

.alert-dialog {
    margin-top: -15px;
    margin-bottom: 15px;
}
</style>