<template>
    <ContentsContainer title="Account" :alertCtx="alertCtx">
        <v-card title="Email" :text=authStore.email class="mb-5"></v-card>
        <v-card title="Password" text="********" class="mb-5"></v-card>
        <v-card class="mb-5">
            <v-card-title>API Key</v-card-title>
            <v-card-text v-if="apikeyExists">
                ************
            </v-card-text>
            <v-card-text v-else>
                <span class="color-muted">No API Key</span>
            </v-card-text>
            <v-card-actions>
                <button data-testid="create-apikey-button" :class="['button-small', { 'bg-color-muted': apikeyExists }]"
                    @click="submitCreateAPIKey" :disabled="apikeyExists">
                    Create
                </button>
                <button data-testid="delete-apikey-button"
                    :class="['button-small', { 'bg-color-muted': !apikeyExists }]" @click="submitDeleteAPIKey"
                    :disabled="!apikeyExists">
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
import { checkAPIKeyStatus, createAPIKey, deleteAPIKey } from '../api/apikey_api';
import { API_BASE_URL, ErrorResponse } from '../api/api_helper';
import { useAlertComposable } from '../composables/alert_composable.ts';


const authStore = useAuthStore();

const apikeyExists = ref(false);
const API_DOC_URL = API_BASE_URL.replace('api/v1', '') + 'swagger/index.html';

const { alertCtx, alert } = useAlertComposable()

onMounted(async () => {
    const status = await checkAPIKeyStatus();
    apikeyExists.value = status.exists;
});

const submitCreateAPIKey = async () => {
    try {
        const res = await createAPIKey();
        if (res instanceof ErrorResponse) {
            throw new Error(res.getMessage());
        }
        apikeyExists.value = true;
        alert("API Key created successfully", "success")
    } catch (error: any) {
        alert(`API Key creat failed: ${error.message}`, "error")
    }
};

const submitDeleteAPIKey = async () => {
    try {
        const res = await deleteAPIKey();
        if (res instanceof ErrorResponse) {
            throw new Error(res.getMessage());
        }
        apikeyExists.value = false;
        alert("API Key deleted successfully", "success")
    } catch (error: any) {
        alert(`API Key delete failed: ${error.message}`, "error")
    }
};
</script>
