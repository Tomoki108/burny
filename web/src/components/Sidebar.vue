<template>
    <div>
        <!-- App bar -->
        <v-app-bar density="comfortable" color="var(--color-tertiary-secondary)" position="fixed">
            <v-app-bar-nav-icon @click="drawer = !drawer" color="white" data-testid="nav-drawer"></v-app-bar-nav-icon>
            <v-app-bar-title class="text-white app-title">🐶&nbsp;&nbsp;Burny</v-app-bar-title>
        </v-app-bar>

        <!-- Navigation drawer -->
        <v-navigation-drawer v-model="drawer" temporary>
            <v-list>
                <v-list-item :to="PATH_PROJECTS" :active="isProjectsRouteActive">
                    <template v-slot:prepend>
                        <font-awesome-icon :icon="['fas', 'diagram-project']" />
                    </template>
                    <v-list-item-title>Projects</v-list-item-title>
                </v-list-item>

                <v-list-item :to="PATH_ACCOUNT" data-testid="nav-account">
                    <template v-slot:prepend>
                        <font-awesome-icon :icon="['fas', 'user']" />
                    </template>
                    <v-list-item-title>Account</v-list-item-title>
                </v-list-item>

                <v-divider></v-divider>

                <v-list-item @click="signOut">
                    <template v-slot:prepend>
                        <font-awesome-icon icon="sign-out-alt" />
                    </template>
                    <v-list-item-title>Sign Out</v-list-item-title>
                </v-list-item>
            </v-list>
        </v-navigation-drawer>
    </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth_store'
import { PATH_ACCOUNT, PATH_HOME, PATH_PROJECTS } from '../router'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

const route = useRoute()
const router = useRouter()
const authStore = useAuthStore()

const drawer = ref(false)

const isProjectsRouteActive = computed(() => route.path.startsWith("/projects"))

const signOut = () => {
    authStore.signOut()
    router.push(PATH_HOME)
    drawer.value = false
}
</script>

<style scoped>
:deep(.v-list-item__prepend .svg-inline--fa) {
    margin-right: 12px;
    font-size: 1.2em;
}

.v-toolbar-title {
    margin-inline-start: 5px !important;
}

header {
    background: var(--color-tertiary-secondary) !important;
}
</style>