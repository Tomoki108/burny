<template>
    <v-dialog v-model="isOpen" max-width="400px" persistent>
        <v-card class="sprint-modal">
            <v-card-title>
                <h2 data-testid="sprint-modal-title">{{ modalTitle }}</h2>
            </v-card-title>
            <v-card-text>
                <v-form ref="sprintForm">
                    <v-text-field label="actual_sp" v-model.number="localSprint.actual_sp" type="number"
                        :rules="newRule('actual_sp').required().gte(0).lte(1000).rules" />
                </v-form>
            </v-card-text>
            <v-card-actions>
                <v-spacer></v-spacer>
                <button data-testid="sprint-update-button" class="button-small" @click.prevent="onSubmit">Save</button>
                <button class="button-small-close" @click="onCancel">Close</button>
            </v-card-actions>
        </v-card>
    </v-dialog>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import type { Sprint } from '../api/sprint_api'
import { newRule, type vForm } from '../utils/validation';

const props = defineProps<{
    show: boolean,
    modalTitle: string,
    sprint: Sprint
}>()

const emits = defineEmits<{
    (e: 'update:show', value: boolean): void;
    (e: 'submit', sprint: Sprint): void;
}>()

const sprintForm = ref<vForm>({} as vForm)
// Create a local copy to allow editing without immediately modifying the parent data
const localSprint = ref<Sprint>({ ...props.sprint })

// Sync the dialog open state with parent prop "show"
const isOpen = ref(props.show)
watch(() => props.show, newVal => {
    isOpen.value = newVal
    if (newVal) {
        localSprint.value = { ...props.sprint }
    }
})
watch(isOpen, val => {
    emits('update:show', val)
})

const onCancel = () => {
    isOpen.value = false
}

const onSubmit = async () => {
    const result = await sprintForm.value.validate()
    if (result.valid) {
        emits('submit', localSprint.value)
        isOpen.value = false
    }
}
</script>

<style scoped>
.sprint-modal {
    padding: 15px;
}
</style>