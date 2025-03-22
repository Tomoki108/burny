<template>
    <v-dialog v-model="ctx.show" class="dialog" max-width="500px">
        <v-sheet class="px-5 py-5">
            <h2>API Key</h2>
            <p class="my-4"> Copy the API Key below and keep it safe. <b>You can't see it again after closing this
                    dialog.</b></p>
            <v-card :text="ctx.rawKey" class="my-4" />
            <div class="mt-1 text-right">
                <button data-testid="copy-apikey-button" class="button-small" @click="copy()">
                    <span v-if="!copied">Copy</span>
                    <span v-else>Copied!!</span>
                </button>
                <button data-testid="dialog-close" class="button-small-close ml-2"
                    @click="ctx.show = false; ctx.rawKey = ''">Close</button>
            </div>
        </v-sheet>
    </v-dialog>
</template>

<script setup lang="ts">
import type { ApikeyContext } from '../composables/apikey_composable';
import { useClipboard } from '@vueuse/core'
import { computed, ref } from 'vue'

const props = defineProps<{
    ctx: ApikeyContext
}>()

const source = computed(() => {
    return props.ctx.rawKey
})

const { copy, copied } = useClipboard({ source })
</script>