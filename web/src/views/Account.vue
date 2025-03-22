<template>
    <ContentsContainer title="Account">
        <v-card title="Email" :text=authStore.email class="mb-5"></v-card>
        <v-card title="Password" text="********" class="mb-5"></v-card>
        <v-card title="API Key" text="*******" class="mb-5">
            <v-card-actions>
                <button data-testid="edit-project-button" class="button-small bg-color-muted">
                    Create
                </button>
                <button data-testid="delete-project-button" class="button-small-danger">
                    Delete
                </button>
            </v-card-actions>
        </v-card>
        <v-alert type="info" class="mt-3">API Document: http://localhost:1323/swagger/index.html</v-alert>


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