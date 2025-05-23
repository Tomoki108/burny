<template>
    <ContentsContainer :title="'Projects > ' + project.title" :alertCtx="alertCtx">
        <v-card title="Basic Info" class="mb-5" :text="projectBasicInfo">
            <button data-testid="edit-project-button" class="button-small ml-4 mb-4"
                @click.prevent="openUpdateProjectModal">
                Edit
            </button>
        </v-card>
        <v-card title="Description" class="mb-5" :text="project.description" />
        <v-card title="Sprint Stats" class="mb-5 p-5">
            <v-table class="m-5">
                <thead>
                    <tr>
                        <th class="text-left">
                            sprint_no
                        </th>
                        <th class="text-left">
                            start_date
                        </th>
                        <th class="text-left">
                            end_date
                        </th>
                        <th class="text-left">
                            ideal_sp
                        </th>
                        <th class="text-left">
                            actual_sp
                        </th>
                        <th class="text-left">
                            update
                        </th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="sprint, idx in sprintsStore.getSprints()" :key="sprint.id">
                        <td>Sprint {{ idx + 1 }}</td>
                        <td data-testid="start_date">{{ sprint.start_date }}</td>
                        <td data-testid="end_date">{{ sprint.end_date }}</td>
                        <td data-testid="ideal_sp">{{ sprint.ideal_sp }}</td>
                        <td data-testid="actual_sp">{{ sprint.actual_sp }}</td>
                        <td v-if="isSprintStarted(sprint)">
                            <button data-testid="update-sprint-button" class="button-small"
                                @click.prevent="openUpdateSprintModal(idx + 1, sprint)">
                                Update
                            </button>
                        </td>
                        <td v-else><span class="color-muted">not started</span></td>
                    </tr>
                </tbody>
            </v-table>
        </v-card>

        <div class="mt-3">
            <Charts :sprints="sprintsStore.getSprints()" :total_sp="project.total_sp" />
        </div>

        <SprintModal :show="updateSprintModal" :modalTitle="'Update Sprint ' + updateSprintNo" :sprint="updateSprint"
            @update:show="updateSprintModal = $event" @submit="submitUpdateSprint" />

        <ProjectModal :show="updateProjectModal" modalTitle="Edit Project" :project="updateProject"
            @update:show="updateProjectModal = $event" @submit="submitUpdateProject" />
    </ContentsContainer>
</template>

<script setup lang="ts">
import ContentsContainer from '../components/ContentsContainer.vue';
import { useProjectsStore } from '../stores/projects_store';
import { ref, onMounted, computed } from 'vue';
import { useRoute } from 'vue-router';
import { defaultProject, type Project } from '../api/project_api';
import { getEndDate } from '../utils/project_helper';
import { type Sprint } from '../api/sprint_api';
import { isSprintStarted } from '../utils/sprint_helper';
import SprintModal from '../components/SprintModal.vue';
import ProjectModal from '../components/ProjectModal.vue';
import { useSprintsStore } from '../stores/sprints_store';
import { useAlertComposable } from '../composables/alert_composable';
import Charts from '../components/Charts.vue';
import { Chart as ChartJS, Title, Tooltip, Legend, LineElement, PointElement, LinearScale, CategoryScale, Filler } from 'chart.js';

ChartJS.register(Title, Tooltip, Legend, Filler, LineElement, PointElement, LinearScale, CategoryScale,);

const route = useRoute();

const projectsStore = useProjectsStore();
const project = ref(defaultProject);
const projectEndDate = ref('');

const projectBasicInfo = computed(() => {
    return `${project.value.start_date} to ${projectEndDate.value}, ${project.value.sprint_count} sprints, ${project.value.total_sp} total sp`;
});

const sprintsStore = useSprintsStore();

const { alertCtx, alert } = useAlertComposable()

onMounted(async () => {
    await projectsStore.fetchProjects();
    const projectID = Number(route.params.id as string);
    project.value = projectsStore.getProject(projectID);
    projectEndDate.value = getEndDate(project.value);

    await sprintsStore.fetchSprints(projectID);
});

// Update Sprint
const updateSprintModal = ref(false);
const updateSprint = ref({} as Sprint);
const updateSprintNo = ref(0);

const openUpdateSprintModal = (sprintNo: number, sprint: Sprint) => {
    updateSprintModal.value = true;
    updateSprint.value = sprint;
    updateSprintNo.value = sprintNo;
};

const submitUpdateSprint = async (sprint: Sprint) => {
    try {
        await sprintsStore.updateSprint(sprint);
        updateSprintModal.value = false;
        alert("Sprint updated successfully", "success");
    } catch (error: any) {
        alert(`Sprint update failed: ${error.message}`, 'error');
    }
};

// Update Project
const updateProjectModal = ref(false);
const updateProject = ref<Project>({} as Project);

const openUpdateProjectModal = () => {
    updateProject.value = { ...project.value };
    updateProjectModal.value = true;
};

const submitUpdateProject = async (updatedProject: Project) => {
    try {
        await projectsStore.updateProject(updatedProject);
        project.value = updatedProject;
        projectEndDate.value = getEndDate(updatedProject);
        alert("Project updated successfully", "success");
    } catch (error: any) {
        alert(`Project update failed: ${error.message}`, 'error');
    }
};
</script>

<style scoped>
tr>td {
    min-width: 110px;
}
</style>