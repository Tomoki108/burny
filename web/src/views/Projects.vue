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
.projects-container {
    padding: 20px;
}

h1 {
    font-size: 24px;
    margin-bottom: 20px;
}

.create-project-btn {
    background-color: #f0f0f0;
    border: none;
    padding: 10px 20px;
    cursor: pointer;
    margin-bottom: 20px;
}

.projects-list {
    display: flex;
    gap: 20px;
}

.project-card {
    background-color: #ffffff;
    padding: 20px;
    border: 1px solid #e0e0e0;
    border-radius: 8px;
    width: 300px;
}

.project-card h2 {
    font-size: 20px;
    margin-bottom: 10px;
}

.project-card p {
    margin-bottom: 10px;
}

.project-actions {
    display: flex;
    gap: 10px;
}

.project-actions button {
    background-color: #f0f0f0;
    border: none;
    padding: 10px;
    cursor: pointer;
}
</style>