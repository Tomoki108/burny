<!-- filepath: /Users/nagatatomoki/Dev/burny/web/src/views/Projects.vue -->
<template>
    <ContentsContainer title="Projects">
        <v-container>
            <v-row class="projects-list" dense>
                <v-col v-for="project in projectsStore.getProjects()" :key="project.id" cols="3" sm="6" md="4"
                    class="project-card">
                    <div>
                        <h2>{{ project.title }}</h2>
                        <p>Sprint: {{ project.sprint_count }}</p>
                        <p>{{ project.description }}</p>
                    </div>
                    <div class="project-actions">
                        <button class="button-small" @click="updateProject(project.id)">Update</button>
                        <button class="button-small" @click="deleteProject(project.id)">Delete</button>
                    </div>
                </v-col>
                <v-col class="project-card-new" cols="3" sm="6" md="4">
                    <h2>+ New Project</h2>
                </v-col>
            </v-row>
        </v-container>
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

.project-card,
.project-card-new {
    background: var(--color-tertiary-secondary);
    padding: 20px;
    border-radius: 4px;
    width: 350px;
    height: 200px;
    text-align: left;
    position: relative;
    color: var(--color-text-light);
}

.project-card-new {
    background: var(--coloer-primary-secondary);
    display: flex;
    align-items: center;
    justify-content: center;
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
    font-weight: bold;
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