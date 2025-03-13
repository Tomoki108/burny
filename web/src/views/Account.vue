<template>
    <ContentsContainer title="Account">
        <v-card title="Email" :text=authStore.email class="mb-5"></v-card>
        <v-card title="Password" text="********"></v-card>
    </ContentsContainer>
</template>

<script setup lang="ts">
import { useAuthStore } from '../stores/auth_store';
import ContentsContainer from '../components/ContentsContainer.vue';

const authStore = useAuthStore();
</script>

<style scoped>
.v-container {
    max-width: 800px;
}
</style>