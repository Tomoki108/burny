<template>
    <v-dialog v-model="ctx.show" class="dialog" closable max-width="500px">
        <v-sheet class="px-5 py-5">
            <h2>{{ ctx.title }}</h2>
            <p class="my-4">
                {{ ctx.text }}
            </p>

            <div class="mt-1 text-right">
                <button v-if="props.callback !== undefined" data-testid="dialog-proceed"
                    class="button-small-danger ml-20" @click="props.callback(); ctx.show = false">Proceed</button>
                <button data-testid="dialog-close"
                    :class="['button-small-close', { 'ml-2': props.callback !== undefined }]"
                    @click="ctx.show = false">Close</button>
            </div>
        </v-sheet>
    </v-dialog>
</template>

<script setup lang="ts">
import type { DialogContext } from '../composables/dialog_composable';

const props = defineProps<{
    ctx: DialogContext
    callback?: () => void
}>()

</script>