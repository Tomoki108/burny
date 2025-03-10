<template>
    <ContentsContainer :title="'Projects > ' + project.title">
        hoge
    </ContentsContainer>
</template>

<script setup lang="ts">
import ContentsContainer from '../components/ContentsContainer.vue';
import { useProjectsStore } from '../stores/projects_store';
import { ref, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import { type Project } from '../api/project_api';

const route = useRoute();
const projectsStore = useProjectsStore();
const project = ref({} as Project);

onMounted(async () => {
    await projectsStore.fetchProjects();
    const projectID = Number(route.params.id as string);
    project.value = projectsStore.getProject(projectID);
});
</script>