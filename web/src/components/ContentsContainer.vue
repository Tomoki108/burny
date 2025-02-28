<template>
    <div class="layout-container">
        <Sidebar />
        <header class="app-header">
            <h1>{{ title }}</h1>
        </header>
        <main class="main-content">
            <slot></slot>
        </main>
    </div>
</template>

<script setup lang="ts">
import { defineProps } from 'vue'
import Sidebar from './Sidebar.vue'

defineProps<{
    title: string
}>()
</script>

<style scoped>
.layout-container {
    display: flex;
}

.app-header {
    position: fixed;
    top: 0;
    left: 200px;
    /* サイドバーの幅を考慮 */
    width: calc(100% - 200px);
    /* サイドバーの幅を考慮 */
    padding: 20px;
    background-color: var(--color-background);
    z-index: 1;
    /* サイドバーより上に表示 */
}

.app-header h1 {
    margin: 0;
    text-align: left;
    color: var(--color-primary);
}

.main-content {
    padding: 80px 20px 20px;
    /* ヘッダーが被らないように上部に余白を追加 */
}
</style>