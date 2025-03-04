<template>
    <ContentsContainer title="Projects">
        <v-row>
            <v-col v-for="project in projectsStore.getProjects()" :key="project.id" lg="3" md="6" sm="12">
                <div class="project-card">
                    <h2>{{ project.title }}</h2>
                    <p>Sprint: {{ project.sprint_count }}</p>
                    <p>{{ project.description }}</p>
                    <div class="project-actions">
                        <button class="button-small" @click="updateProject(project.id)">Update</button>
                        <button class="button-small" @click="deleteProject(project.id)">Delete</button>
                    </div>
                </div>

            </v-col>
            <v-col lg="3" md="6" sm="12" @click="openNewProjectModal">
                <div class="project-card-new">
                    <h2>+ New Project</h2>
                </div>
            </v-col>
        </v-row>

        <ProjectModal :show="newProjectModal" modalTitle="New Project" :project="defaultProject"
            @update:show="newProjectModal = $event" @submit="submitNewProject" />
    </ContentsContainer>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import ContentsContainer from '../components/ContentsContainer.vue'
import ProjectModal from '../components/ProjectModal.vue'
import { useProjectsStore } from '../stores/projects_store.ts'
import { defaultProject, type Project } from '../api/project_api'

const projectsStore = useProjectsStore()
const newProjectModal = ref(false)


const openNewProjectModal = () => {
    newProjectModal.value = true
}

const submitNewProject = async (project: Project) => {
    try {
        console.log("Creating new project:", project)
        await projectsStore.createProject(project)
    } catch (error) {
        console.error("New project creation failed:", error)
    }
}

const updateProject = (id: number) => { /* ... */ }
const deleteProject = (id: number) => { /* ... */ }

onMounted(() => {
    projectsStore.fetchProjects()
})
</script>

<style scoped>
.project-card,
.project-card-new {
    background: var(--color-tertiary-secondary);
    padding: 20px;
    border-radius: 4px;
    min-width: 220px;
    width: auto;
    height: 200px;
    text-align: left;
    position: relative;
    color: var(--color-text-light);
    cursor: pointer;
}

.project-card-new {
    background: var(--coloer-primary-secondary);
    display: flex;
    align-items: center;
    justify-content: center;
}

.project-card-new:hover {
    background: var(--color-tertiary);
}

.project-actions {
    position: absolute;
    bottom: 20px;
    right: 20px;
    display: flex;
    gap: 10px;
}
</style>