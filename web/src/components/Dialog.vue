<template>
    <v-dialog v-model="ctx.show" class="dialog" closable>
        <v-sheet class="px-5 py-5">
            <h2>{{ ctx.title }}</h2>
            <p class="my-4">
                {{ ctx.text }}
            </p>

            <div class="mt-1 text-right">
                <button data-testid="dialog-proceed" class="button-small-danger ml-20"
                    @click="callback(); ctx.show = false">Proceed</button>
                <button data-testid="dialog-cancel" class="button-small-cancel ml-2"
                    @click="ctx.show = false">Cancel</button>
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
    top: -67.5%;
    width: 20%;
}
</style>