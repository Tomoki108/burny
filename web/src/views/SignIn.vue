<template>
    <div class="container">
        <div class="login-box">
            <h1>Burny</h1>
            <form @submit.prevent="onSubmit">
                <label for="email">Username</label>
                <input type="text" id="email" v-model="email" placeholder="yourname@burnuppro.io" required minlength="8"
                    maxlength="20" />
                <label for="password">Password</label>
                <input type="password" id="password" v-model="password" placeholder="************" required />
                <button type="submit" class="button">Sign In</button>
                <div class="or">or</div>
                <button type="button" class="button">Create account</button>
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
h1 {
    text-align: center;
    margin-bottom: 20px;
    color: var(--color-text-light);
}

.or {
    margin: 16px 0;
    text-align: center;
}


/* Container and login-box are specific to the SignIn view */
.container {
    background-color: var(--color-secondary);
    margin: 0 auto;
    position: relative;
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
    width: 100%;
    z-index: 1;
}

.login-box {
    position: relative;
    z-index: 1;
    background: linear-gradient(var(--color-tertiary), var(--color-secondary));
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