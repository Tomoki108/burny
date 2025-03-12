<template>
    <ContentsContainer :title="'Projects > ' + project.title" :alertCtx="alertCtx">
        <h2>Basic Info</h2>
        <p>{{ project.start_date }} to {{ projectEndDate }}, {{ project.sprint_count }} sprints, {{
            project.total_sp }} sp</p>
        <h2 class="mt-3">Description</h2>
        <p>{{ project.description }}</p>
        <h2 class="mt-3">Sprint Stats</h2>
        <v-table>
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
                    <td>{{ sprint.start_date }}</td>
                    <td>{{ sprint.end_date }}</td>
                    <td>{{ sprint.ideal_sp }}</td>
                    <td>{{ sprint.actual_sp }}</td>
                    <td>
                        <button class="button-small" v-if="isSprintStarted(sprint)"
                            @click.prevent="openUpdateSprintModal(sprint)">
                            Update
                        </button>
                    </td>
                </tr>
            </tbody>
        </v-table>

        <div class="mt-3">
            <Charts :sprints="sprintsStore.getSprints()" :total_sp="project.total_sp" />
        </div>

        <SprintModal :show="updateSprintModal" modalTitle="Update Sprint" :sprint="updateSprint"
            @update:show="updateSprintModal = $event" @submit="submitUpdateSprint" />
    </ContentsContainer>
</template>

<script setup lang="ts">
import ContentsContainer from '../components/ContentsContainer.vue';
import { useProjectsStore } from '../stores/projects_store';
import { ref, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import { type Project } from '../api/project_api';
import { getEndDate } from '../utils/project_helper';
import { type Sprint } from '../api/sprint_api';
import { isSprintStarted } from '../utils/sprint_helper';
import SprintModal from '../components/SprintModal.vue';
import { useSprintsStore } from '../stores/sprints_store';
import { useAlertComposable } from '../composables/alert_composable';
import { Chart as ChartJS, Title, Tooltip, Legend, LineElement, PointElement, LinearScale, CategoryScale } from 'chart.js';
import Charts from '../components/Charts.vue';

ChartJS.register(Title, Tooltip, Legend, LineElement, PointElement, LinearScale, CategoryScale);

const route = useRoute();

const projectsStore = useProjectsStore();
const project = ref({} as Project);
const projectEndDate = ref('');

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

const openUpdateSprintModal = (sprint: Sprint) => {
    updateSprintModal.value = true;
    updateSprint.value = sprint;
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
</script>
