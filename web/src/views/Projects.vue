<!-- filepath: /Users/nagatatomoki/Dev/burny/web/src/views/Projects.vue -->
<template>
    <ContentsContainer title="Projects">
        <button class="button-small" @click="createProject()">Delete</button>

        <div class="projects-list">
            <div class="project-card" v-for="project in projectsStore.getProjects()" :key="project.id">
                <div>
                    <h2>{{ project.title }}</h2>
                    <p>Sprint: {{ project.sprint_count }}</p>
                    <p>{{ project.description }}</p>
                </div>
                <div class="project-actions">
                    <button class="button-small" @click="updateProject(project.id)">Update</button>
                    <button class="button-small" @click="deleteProject(project.id)">Delete</button>
                </div>
            </div>
        </div>
    </ContentsContainer>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import ContentsContainer from '../components/ContentsContainer.vue'
import { useProjectsStore } from '../stores/projects_store.ts'

const projectsStore = useProjectsStore()

const createProject = () => { /* ... */ }
const updateProject = (id: number) => { /* ... */ }
const deleteProject = (id: number) => { /* ... */ }

onMounted(() => {
    projectsStore.fetchProjects()
})
</script>

<style scoped>
/* 画面レイアウト・コンテナ、カードの配置など、コンポーネント固有のスタイルのみを記述 */
/* 色やフォント、ボタン、入力フォームは全てglobal.cssを参照 */

.projects-container {
    /* 左パディングを 0 にしてヘッダータイトルと左をそろえる */
    padding: 20px 20px 20px 0;
}

.projects-list {
    display: flex;
    gap: 20px;
}

.project-card {
    background-color: var(--color-background);
    padding: 20px;
    border: 1.5px solid var(--color-muted);
    border-radius: 18px;
    width: 300px;
    text-align: left;
    position: relative;
    padding-bottom: 60px;
    box-shadow: 0.2rem 0.2rem var(--color-shadow);
}

.project-actions {
    position: absolute;
    bottom: 20px;
    /* 下部に配置 */
    right: 20px;
    /* 右端に配置 */
    display: flex;
    gap: 10px;
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