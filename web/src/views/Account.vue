<template>
    <ContentsContainer title="Account">
        <v-card title="Email" :text=authStore.email class="mb-5"></v-card>
        <v-card title="Password" text="********" class="mb-5"></v-card>
        <v-card class="mb-5">
            <v-card-title>API Key</v-card-title>
            <v-card-text v-if="apikeyExists">
                ************
            </v-card-text>
            <v-card-text v-else>
                <span class="color-muted">No API key registered</span>
            </v-card-text>
            <v-card-actions>
                <button data-testid="create-apikey-button" :class="['button-small', { 'bg-color-muted': apikeyExists }]"
                    :disabled="apikeyExists">
                    Create
                </button>
                <button data-testid="delete-apikey-button"
                    :class="['button-small', { 'bg-color-muted': !apikeyExists }]" :disabled="!apikeyExists">
                    Delete
                </button>
            </v-card-actions>
        </v-card>
        <v-alert type="info" class="mt-3">API Document: {{ API_DOC_URL }}</v-alert>
    </ContentsContainer>
</template>

<script setup lang="ts">
import { useAuthStore } from '../stores/auth_store';
import ContentsContainer from '../components/ContentsContainer.vue';
import { onMounted, ref } from 'vue';
import { checkAPIKeyStatus } from '../api/apikey_api';
import { API_BASE_URL } from '../api/api_helper';

const authStore = useAuthStore();

const apikeyExists = ref(false);
const API_DOC_URL = API_BASE_URL.replace('api/v1', '') + 'swagger/index.html';

onMounted(async () => {
    const status = await checkAPIKeyStatus();
    apikeyExists.value = status.exists;
});
</script>
