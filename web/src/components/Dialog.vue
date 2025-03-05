<template>
    <v-dialog v-model="ctx.show" class="dialog" closable>
        <v-sheet class="px-5 py-5 my-2 mx-5 w-2xs">
            <h2>Dialog</h2>
            <p class="my-4">
                {{ ctx.text }}
            </p>

            <div class="flex justify-end">
                <button class="button-small-danger ml-20" @click="callback(); ctx.show = false">Proceed</button>
                <button class="button-small-cancel ml-2" @click="ctx.show = false">Cancel</button>
            </div>
        </v-sheet>
    </v-dialog>
</template>

<script setup lang="ts">
import type { DialogContext } from '../composables/dialog_composable';

defineProps<{
    ctx: DialogContext
    callback: () => void
}>()

</script>

<style scoped>
.dialog {
    align-content: center;
    top: -80%;
    left: 45%;
}
</style>