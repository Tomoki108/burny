<template>
    <Navigation />
    <main class="main-content">
        <v-alert v-if="alertCtx?.show" :type="alertCtx?.type" :text="alertCtx?.text" closable class="alert-dialog" />
        <header class="app-header">
            <h1>{{ title }}</h1>
        </header>
        <hr />
        <slot></slot>
    </main>
</template>

<script setup lang="ts">
import Navigation from './Navigation.vue'
import type { AlertContext } from '../composables/alert_composable'

defineProps<{
    title: string
    alertCtx?: AlertContext,
}>()
</script>

<style scoped>
.main-content {
    width: 100%;
    padding: 25px 40px 40px 40px;
    margin-top: 64px;
}

.app-header h1 {
    margin-bottom: 20px;
    text-align: left;
    color: var(--color-text);
    font-weight: bold;
}

hr {
    border-top: 2px solid var(--color-muted);
    margin-bottom: 30px;
}

.alert-dialog {
    margin-top: -15px;
    margin-bottom: 15px;
}
</style>