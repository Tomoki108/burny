<template>
    <ContentsContainer title="Projects" :alertCtx="alertCtx">
        <v-row>
            <v-col v-for="project in projectsStore.getProjects()" :key="project.id" lg="3" md="6" sm="12">
                <div class="project-card">
                    <h2>{{ project.title }}</h2>
                    <p>Sprint: {{ project.sprint_count }}</p>
                    <p>{{ project.description }}</p>
                    <div class="project-actions">
                        <button class="button-small" @click="openUpdateProjectModal(project)">Update</button>
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
        <ProjectModal :show="updateProjectModal" modalTitle="Update Project" :project="updateProject"
            @update:show="updateProjectModal = $event" @submit="submitUpdateProject" />
    </ContentsContainer>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import ContentsContainer from '../components/ContentsContainer.vue'
import ProjectModal from '../components/ProjectModal.vue'
import { useProjectsStore } from '../stores/projects_store.ts'
import { defaultProject, type Project } from '../api/project_api'
import { useAlertComposable } from '../composables/alert_composable.ts'

const projectsStore = useProjectsStore()

const { alertCtx, alert } = useAlertComposable()

// Create Modal
const newProjectModal = ref(false)

const openNewProjectModal = () => {
    newProjectModal.value = true
}

const submitNewProject = async (project: Project) => {
    try {
        await projectsStore.createProject(project)
        alert("Project created successfully", "success")
    } catch (error: any) {
        alert(`Project creation failed: ${error.message}`, "error")
    }
}

// Update Modal
const updateProject = ref<Project>({} as Project)
const updateProjectModal = ref(false)

const openUpdateProjectModal = (project: Project) => {
    updateProject.value = project
    updateProjectModal.value = true
}

const submitUpdateProject = async (project: Project) => {
    try {
        await projectsStore.updateProject(project)
        alert("Project updated successfully", "success")
    } catch (error: any) {
        alert(`Project update failed: ${error.message}`, 'error')
    }
}

const deleteProject = (id: number) => {
    try {
        projectsStore.deleteProject(id)
        alert("Project deleted successfully", "success")
    } catch (error) {
        alert(`Project deletion failed. Please retry`, 'error')
    }
}

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