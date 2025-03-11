<template>
    <ContentsContainer :title="'Projects > ' + project.title">
        <h2 class="mb-2">Term</h2>
        <p>{{ project.start_date }} to {{ projectEndDate }}, {{ project.sprint_count }} sprints</p>
        <h2 class="my-2">Description</h2>
        <p>{{ project.description }}</p>
        <h2 class="my-2">Sprint Stats</h2>
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
                <tr v-for="sprint in sprints" :key="sprint.id">
                    <td>{{ sprint.start_date }}</td>
                    <td>{{ sprint.end_date }}</td>
                    <td>{{ sprint.ideal_sp }}</td>
                    <td>{{ sprint.actual_sp }}</td>
                    <td>
                        <button class="button-small" @click="openUpdateSprintModal(sprint)">Update</button>
                    </td>
                </tr>
            </tbody>
        </v-table>
    </ContentsContainer>
</template>

<script setup lang="ts">
import ContentsContainer from '../components/ContentsContainer.vue';
import { useProjectsStore } from '../stores/projects_store';
import { ref, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import { type Project } from '../api/project_api';
import { getEndDate } from '../utils/project_helper';
import { type Sprint, fetchSprints } from '../api/sprint_api';

const route = useRoute();

const projectsStore = useProjectsStore();
const project = ref({} as Project);
const projectEndDate = ref('');

const sprints = ref([] as Sprint[]);

onMounted(async () => {
    await projectsStore.fetchProjects();
    const projectID = Number(route.params.id as string);
    project.value = projectsStore.getProject(projectID);
    projectEndDate.value = getEndDate(project.value);

    sprints.value = await fetchSprints(projectID);
});

const openUpdateSprintModal = (sprint: Sprint) => {
    console.log('openUpdateSprintModal', sprint);
};
</script>