<template>
    <ContentsContainer title="Projects" :alertCtx="alertCtx">
        <v-row>
            <v-col v-for="project in projectsStore.getProjects()" :key="project.id" lg="3" md="6" sm="6" xs="12">
                <router-link :to="'/projects/' + project.id" :props="project">
                    <div class="project-card" :data-testid="'project-card-' + project.id">
                        <h2 class="mb-2">{{ project.title }}</h2>
                        <p>{{ project.start_date }} to {{ getEndDate(project) }}, {{ project.sprint_count }} sprints, {{
                            project.total_sp }} total sp</p>
                        <p class="mb-2"></p>
                        <p class="project-description">{{ truncateStr(project.description, 70) }}</p>
                        <div class="project-actions">
                            <button :data-testid="'edit-project-button-' + project.id" class="button-small"
                                @click.prevent="openUpdateProjectModal(project)">Edit
                            </button>
                            <button :data-testid="'delete-project-button-' + project.id" class="button-small-danger"
                                @click.prevent="dialog(`Delete Project`, `Are you shure
                                to delete project ${project.title}?`)">Delete
                                <Dialog :ctx="dialogCtx" :callback="() => submitDeleteProject(project.id)">
                                </Dialog>
                            </button>
                        </div>
                    </div>
                </router-link>

            </v-col>
            <v-col lg="3" md="6" sm="6" @click="openNewProjectModal">
                <div data-testid="new-project-button" class="project-card-new">
                    <h2>+ New Project</h2>
                </div>
            </v-col>
        </v-row>

        <ProjectModal :show="newProjectModal" modalTitle="New Project" :project="defaultProject"
            @update:show="newProjectModal = $event" @submit="submitNewProject" />
        <ProjectModal :show="updateProjectModal" modalTitle="Edit Project" :project="updateProject"
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
import Dialog from '../components/Dialog.vue'
import { useDialogComposable } from '../composables/dialog_composable'
import { getEndDate } from '../utils/project_helper'
import { truncateStr } from '../utils/string_helper'

const projectsStore = useProjectsStore()
const { alertCtx, alert } = useAlertComposable()
const { dialogCtx, dialog } = useDialogComposable()

onMounted(() => {
    projectsStore.fetchProjects()
})

// Create Project
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

// Update Project
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

// Delete Project
const submitDeleteProject = (id: number) => {
    try {
        projectsStore.deleteProject(id)
        alert("Project deleted successfully", "success")
    } catch (error) {
        alert(`Project deletion failed. Please retry`, 'error')
    }
}
</script>

<style scoped>
.project-card,
.project-card-new {
    background: var(--color-tertiary-secondary);
    border-radius: 4px;
    min-width: 220px;
    width: auto;
    height: 245px;
    text-align: left;
    position: relative;
    color: var(--color-text-light);
    cursor: pointer;
}

.project-card-new {
    padding: 20px;
}

.project-card {
    padding: 20px 20px 70px 20px;
}

.project-card-new {
    background: var(--color-primary-secondary);
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