<template>
    <div class="projects-container">
        <h1>Projects</h1>
        <button class="create-project-btn" @click="createProject">Create Project</button>
        <div class="projects-list">
            <div class="project-card" v-for="project in projectsStore.getProjects()" :key="project.id">
                <h2>{{ project.title }}</h2>
                <p>Sprint: {{ project.sprint_count }}</p>
                <p>{{ project.description }}</p>
                <div class="project-actions">
                    <button @click="updateProject(project.id)">Update</button>
                    <button @click="deleteProject(project.id)">Delete</button>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import { useProjectsStore } from '../stores/projects_store.ts'

const projectsStore = useProjectsStore()

const createProject = () => {
    // プロジェクト作成のロジックを追加
}

const updateProject = (id: number) => {
    // プロジェクト更新のロジックを追加
}

const deleteProject = (id: number) => {
    // プロジェクト削除のロジックを追加
}

onMounted(() => {
    projectsStore.fetchProjects()
})
</script>

<style scoped>
/* 画面レイアウト・コンテナ、カードの配置など、コンポーネント固有のスタイルのみを記述 */
/* 色やフォント、ボタン、入力フォームは全てglobal.cssを参照 */
.projects-container {
    padding: 20px;
}

.projects-list {
    display: flex;
    gap: 20px;
}

.project-card {
    background-color: var(--color-background);
    padding: 20px;
    border: 1px solid var(--color-muted);
    border-radius: 8px;
    width: 300px;
}

.project-card h2 {
    font-size: var(--font-size-large);
    margin-bottom: 10px;
}

.project-card p {
    margin-bottom: 10px;
}

.project-actions {
    display: flex;
    gap: 10px;
}
</style>