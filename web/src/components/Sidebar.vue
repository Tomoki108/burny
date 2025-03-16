<template>
    <div class="sidebar">
        <nav>
            <ul>
                <li>
                    <h1 class="mt-n3 color-white">🐶 Burny</h1>
                </li>
                <router-link :to=PATH_PROJECTS custom v-slot="{ navigate }">
                    <li :class="{ 'active-li': isProjectsRouteActive }" @click="navigate">
                        <font-awesome-icon :icon="['fas', 'diagram-project']" /> Projects
                    </li>
                </router-link>
                <router-link :to=PATH_ACCOUNT custom v-slot="{ navigate, isActive }">
                    <li :class="{ 'active-li': isActive }" @click="navigate">
                        <font-awesome-icon :icon="['fas', 'user']" /> Account
                    </li>
                </router-link>
            </ul>
        </nav>
        <div class="signout" @click="signOut">
            <font-awesome-icon icon="sign-out-alt" /> Sign Out
        </div>
    </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth_store'
import { PATH_ACCOUNT, PATH_HOME, PATH_PROJECTS } from '../router'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'


const route = useRoute()
const router = useRouter()
const authStore = useAuthStore()

const isProjectsRouteActive = computed(() => route.path.startsWith("/projects"))

const signOut = () => {
    authStore.signOut()
    router.push(PATH_HOME)
}
</script>


<style scoped>
.sidebar {
    width: 10%;
    min-width: 210px;
    background: var(--color-tertiary-secondary);
    height: 100vh;
    position: sticky;
    top: 0;
    left: 0;
}

li {
    cursor: pointer;
    font-size: var(--font-size-large);
}

li:not(:first-child):hover,
div.signout:hover {
    color: var(--color-text-light);
    font-weight: bold;
    background-color: var(--color-tertiary);
}


.active-li {
    color: var(--color-text-light);
    font-weight: bold;
    background-color: var(--color-tertiary);
}

nav {
    margin-top: 20px;
}

nav ul {
    list-style: none;
}

nav ul li {
    padding-top: 20px;
    padding-right: 40px;
    padding-left: 40px;
    padding-bottom: 20px;
}

nav ul li a {
    text-decoration: none;
    color: var(--color-text);
    font-size: var(--font-size-large);
    display: flex;
    align-items: center;
}

nav ul li a svg {
    margin-right: 20px;
}

.signout {
    position: absolute;
    bottom: 20px;
    padding-top: 20px;
    padding-right: 40px;
    padding-left: 40px;
    padding-bottom: 20px;
    width: 100%;
    min-width: inherit;
}

.signout a {
    font-weight: 0 !important;
    display: flex;
    align-items: center;
    color: inherit;
    font-size: var(--font-size-large);
}

.signout a:hover {
    color: var(--color-text-light);
    font-weight: bold;
}

.signout a svg,
li svg {
    margin-right: 8px;
}
</style>