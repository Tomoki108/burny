<template>
    <ContentsContainer :title="'Projects > ' + project.title" :alertCtx="alertCtx">
        <h2 class="mb-1">Term</h2>
        <p>{{ project.start_date }} to {{ projectEndDate }}, {{ project.sprint_count }} sprints</p>
        <h2 class="mt-3 mb-1">Description</h2>
        <p>{{ project.description }}</p>
        <h2 class="mt-3">Sprint Stats</h2>
        <v-table>
            <thead>
                <tr>
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
                <tr v-for="sprint in sprintsStore.getSprints()" :key="sprint.id">
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

    sprintsStore.fetchSprints(projectID);
});

// Update Sprint
const updateSprintModal = ref(false);
const updateSprint = ref({} as Sprint);

const openUpdateSprintModal = (sprint: Sprint) => {
    console.log(sprint);
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