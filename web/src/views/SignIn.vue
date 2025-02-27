<!-- filepath: /Users/nagatatomoki/Dev/burny/web/src/views/SignIn.vue -->
<template>
    <div class="container">
        <div class="background-circle-left"></div>
        <div class="background-circle-right"></div>
        <div class="login-box">
            <h1>Burny</h1>
            <form @submit.prevent="onSubmit">
                <label for="email">Username</label>
                <input type="text" id="email" v-model="email" placeholder="yourname@burnuppro.io" required minlength="8"
                    maxlength="20" />
                <label for="password">Password</label>
                <input type="password" id="password" v-model="password" placeholder="************" required />
                <button type="submit">Sign In</button>
                <div class="or">or</div>
                <button type="button">Create account</button>
            </form>
            <p v-if="error" class="error">{{ error }}</p>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth_store'
import { PATH_PROJECTS } from '../router'

const email = ref('')
const password = ref('')
const error = ref('')

const router = useRouter()
const authStore = useAuthStore()

const onSubmit = async () => {
    error.value = ''
    try {
        await authStore.signIn(email.value, password.value)
        router.push(PATH_PROJECTS)
    } catch (err) {
        error.value = 'Login failed. Please check your credentials.'
    }
}
</script>

<style scoped>
/* Background circles */
.background-circle-left {
    position: absolute;
    top: -200px;
    left: -700px;
    width: 600px;
    height: 600px;
    background-color: #efefef;
    border-radius: 50%;
    z-index: 0;
}

.background-circle-right {
    position: absolute;
    bottom: 0;
    right: -600px;
    width: 400px;
    height: 400px;
    background-color: #efefef;
    border-radius: 50%;
    z-index: 0;
}

/* Container and login-box are specific to the SignIn view */
.container {
    position: relative;
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
    width: 100%;
}

.login-box {
    position: relative;
    z-index: 1;
    background-color: var(--color-background-grey);
    padding: 40px;
    border-radius: 8px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
    width: 400px;
}

.or {
    margin: 16px 0;
    color: var(--color-muted);
}
</style>