<template>
    <v-dialog v-model="isOpen" max-width="600px" persistent>
        <v-card class="project-modal">
            <v-card-title>
                <h2>&nbsp&nbsp{{ modalTitle }}</h2>
            </v-card-title>
            <v-card-text>
                <v-form ref="projectForm">
                    <v-text-field label="Title" v-model="localProject.title" :rules="[required('Title')]" />
                    <br>
                    <v-textarea label="Description" v-model="localProject.description" />
                    <br>
                    <v-text-field label="Total SP" v-model.number="localProject.total_sp" type="number"
                        :rules="[v => !!v || 'Total SP is required']" />
                    <br>
                    <v-text-field label="Sprint Count" v-model.number="localProject.sprint_count" type="number"
                        :rules="[v => !!v || 'Sprint Count is required']" />
                    <br>
                    <v-text-field label="Sprint Duration (weeks)" v-model.number="localProject.sprint_duration"
                        type="number" :rules="[v => !!v || 'Sprint Duration is required']" />
                    <br>
                    <v-text-field label="Start Date" v-model="localProject.start_date" type="date"
                        :rules="[v => !!v || 'Start Date is required']" />
                </v-form>
            </v-card-text>

            <v-card-actions>
                <v-spacer></v-spacer>
                <button class="button-small" @click.prevent="onSubmit">Save</button>
                <button class="button-small" @click="onCancel">Cancel</button>
            </v-card-actions>
        </v-card>
    </v-dialog>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import type { Project } from '../api/project_api';
import { required } from '../utils/validation_rules';

const props = defineProps<{
    show: boolean,
    modalTitle: string,
    project: Project
}>()

const emits = defineEmits<{
    (e: 'update:show', value: boolean): void;
    (e: 'submit', project: Project): void;
}>()

// フォームの参照を取得
const projectForm = ref()

// Local copy to allow editing without modifying parent data immediately.
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
    // Vue/Vuetify のバリデーションを実行
    if (projectForm.value.validate()) {
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