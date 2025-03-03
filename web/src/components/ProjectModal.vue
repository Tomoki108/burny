<template>
    <v-dialog v-model="isOpen" max-width="600px">
        <v-card class="project-modal">
            <v-card-title class="px-24">
                <h2>&nbsp{{ modalTitle }}</h2>
            </v-card-title>
            <v-card-text>
                <v-form ref="projectForm">
                    <v-text-field label="Title" v-model="localProject.title" required />
                    <v-textarea label="Description" v-model="localProject.description" />
                    <v-text-field label="Total SP" v-model.number="localProject.total_sp" type="number" required />
                    <v-text-field label="Sprint Count" v-model.number="localProject.sprint_count" type="number"
                        required />
                    <v-text-field label="Sprint Duration (weeks)" v-model.number="localProject.sprint_duration"
                        type="number" required />
                    <v-text-field label="Start Date" v-model="localProject.start_date" type="date" required />
                </v-form>
            </v-card-text>

            <v-card-actions>
                <button class="button-small" @click="onCancel">Cancel</button>
                <button class="button-small" color="primary" @click="onSubmit">Save</button>
            </v-card-actions>
        </v-card>
    </v-dialog>
</template>

<script setup lang="ts">
import { ref, watch, defineProps, defineEmits } from 'vue'

interface Project {
    title: string;
    description: string;
    total_sp: number;
    sprint_count: number;
    sprint_duration: number;
    start_date: string;
}

const props = defineProps<{
    show: boolean,
    modalTitle: string,
    project: Project
}>()

const emits = defineEmits<{
    (e: 'update:show', value: boolean): void;
    (e: 'submit', project: Project): void;
}>()

// Local copy to allow editing without immediately modifying parent data.
const localProject = ref<Project>({ ...props.project })

// Maintain dialog state synced with parent prop 'show'
const isOpen = ref(props.show)
watch(() => props.show, (newVal) => {
    isOpen.value = newVal
    if (newVal) {
        // reset localProject when opening modal
        localProject.value = { ...props.project }
    }
})
watch(isOpen, (val) => {
    emits('update:show', val)
})

const onCancel = () => {
    isOpen.value = false
}

const onSubmit = () => {
    emits('submit', localProject.value)
    isOpen.value = false
}
</script>

<style scoped>
.project-modal {
    padding: 15px;
}
</style>