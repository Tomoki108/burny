<template>
    <ContentsContainer :title="'Projects > ' + project.title">
        <h2 class="my-2">Term</h2>
        <p>{{ project.start_date }} to {{ projectEndDate }}, {{ project.sprint_count }} sprints</p>
        <h2 class="my-2">Description</h2>
        <p>{{ project.description }}</p>
    </ContentsContainer>
</template>

<script setup lang="ts">
import ContentsContainer from '../components/ContentsContainer.vue';
import { useProjectsStore } from '../stores/projects_store';
import { ref, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import { type Project } from '../api/project_api';
import { getEndDate } from '../utils/project_helper';

const route = useRoute();
const projectsStore = useProjectsStore();
const project = ref({} as Project);
const projectEndDate = ref('');


onMounted(async () => {
    await projectsStore.fetchProjects();
    const projectID = Number(route.params.id as string);
    project.value = projectsStore.getProject(projectID);
    projectEndDate.value = getEndDate(project.value);
});
</script>