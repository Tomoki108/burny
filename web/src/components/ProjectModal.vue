<template>
    <v-dialog v-model="isOpen" max-width="600px" persistent>
        <v-card class="project-modal">
            <v-card-title>
                <h2>&nbsp&nbsp{{ modalTitle }}</h2>
            </v-card-title>
            <v-card-text>
                <v-form ref="projectForm">
                    <v-text-field name="title" label="Title" v-model="localProject.title"
                        :rules="newRule('title').required().lte(50).rules" /> <br>
                    <v-textarea name="description" label="Description" v-model="localProject.description"
                        :rules="newRule('Description').lte(500).rules" />
                    <br>
                    <v-text-field name="total-sp" label="Total SP" v-model.number="localProject.total_sp" type="number"
                        :rules="newRule('Total SP').required().lte(1000).gt(0).rules" />
                    <br>
                    <v-text-field name="sprint-count" label="Sprint Count" v-model.number="localProject.sprint_count"
                        type="number" :rules="newRule('Sprint Count').required().gt(0).lte(100).rules" />
                    <v-alert type="info" text="Sprint Weeks and Start Date can't be updated after creation."></v-alert>
                    <br>
                    <v-select name="sprint-duration" :items="[1, 2, 3]" label="Sprint Weeks"
                        v-model.number="localProject.sprint_duration" :disabled="project.id !== 0"></v-select>
                    <br>
                    <v-text-field name="start-date" label="Start Date" v-model="localProject.start_date" type="date"
                        :rules="newRule('Start Date').required().rules" :disabled="project.id !== 0" />
                </v-form>
            </v-card-text>

            <v-card-actions>
                <v-spacer></v-spacer>
                <button class="button-small" data-testid="project-save" @click.prevent="onSubmit">Save</button>
                <button class="button-small-cancel" data-testid="project-cancel" @click="onCancel">Cancel</button>
            </v-card-actions>
        </v-card>
    </v-dialog>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import type { Project } from '../api/project_api';
import { newRule, type vForm } from '../utils/validation';

const props = defineProps<{
    show: boolean,
    modalTitle: string,
    project: Project
}>()

const emits = defineEmits<{
    (e: 'update:show', value: boolean): void;
    (e: 'submit', project: Project): void;
}>()

const projectForm = ref<vForm>({} as vForm)

// Local copy to allow editing without modifying parent data immediately.
const localProject = ref<Project>({ ...props.project })

// Maintain dialog state synced with parent prop 'show'
const isOpen = ref(props.show)
watch(() => props.show, (newVal) => {
    isOpen.value = newVal
    if (newVal) {
        localProject.value = { ...props.project }
    }
})
watch(isOpen, (val) => {
    emits('update:show', val)
})

const onCancel = () => {
    isOpen.value = false
}

const onSubmit = async () => {
    const result = await projectForm.value.validate()
    if (result.valid) {
        emits('submit', localProject.value)
        isOpen.value = false
    }
}
</script>

<style scoped>
.project-modal {
    padding: 15px;
}
</style>