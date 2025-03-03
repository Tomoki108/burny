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
                <v-col class="project-card-new" cols="3" sm="6" md="4" @click="openNewProjectModal">
                    <h2>+ New Project</h2>
                </v-col>
            </v-row>
        </v-container>

        <!-- Use the generic ProjectModal for creating a new project -->
        <ProjectModal :show="newProjectModal" modalTitle="New Project" :project="defaultProject"
            @update:show="newProjectModal = $event" @submit="submitNewProject" />
    </ContentsContainer>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import ContentsContainer from '../components/ContentsContainer.vue'
import ProjectModal from '../components/ProjectModal.vue'
import { useProjectsStore } from '../stores/projects_store.ts'

const projectsStore = useProjectsStore()
const newProjectModal = ref(false)

// Default project object for a new project.
const defaultProject = ref({
    title: '',
    description: '',
    total_sp: 0,
    sprint_count: 1,
    sprint_duration: 1,
    start_date: '',
})

const openNewProjectModal = () => {
    newProjectModal.value = true
}

const submitNewProject = async (projectData: typeof defaultProject.value) => {
    try {
        // Here, you would call your API to create the project.
        // For example: await createProject(projectData)
        await projectsStore.fetchProjects()
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
    cursor: pointer;
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
    right: 20px;
    display: flex;
    gap: 10px;
}
</style>