<template>
    <ContentsContainer :title="'Projects > ' + project.title">
        hoge
    </ContentsContainer>
</template>

<script setup lang="ts">
import ContentsContainer from '../components/ContentsContainer.vue';
import { useProjectsStore } from '../stores/projects_store';
import { computed } from 'vue';
import { useRoute } from 'vue-router';
const route = useRoute();

const projectsStore = useProjectsStore()
const project = computed(() => {
    const projectID = route.params.id as string
    return projectsStore.getProject(Number(projectID))
})

</script>