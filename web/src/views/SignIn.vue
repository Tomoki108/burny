<template>
    <div class="container">
        <div class="login-box">
            <h1 class="color-white">Burny</h1>
            <div class="tabs">
                <button @click="isSignUp = false" :class="{ active: !isSignUp }">Sign In</button>
                <button @click="isSignUp = true" :class="{ active: isSignUp }">Sign Up</button>
            </div>
            <form @submit.prevent="onSubmit">
                <label for="email">Email</label>
                <input type="text" id="email" v-model="email" placeholder="yourname@burnuppro.io" required />
                <label for="password">Password</label>
                <input type="password" id="password" v-model="password" placeholder="************" required
                    minlength="8" maxlength="20" />
                <button type="submit" class="button">{{ isSignUp ? 'Sign Up' : 'Sign In' }}</button>
            </form>
            <p v-if="error" class="error">{{ error }}</p>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useAuthStore } from '../stores/auth_store'
import { PATH_PROJECTS } from '../router'
import { signUp } from '../api/auth_api'

const email = ref('')
const password = ref('')
const error = ref('')
const isSignUp = ref(false)
const router = useRouter()
const route = useRoute()
const authStore = useAuthStore()

// Check signup from URL query parameter
onMounted(() => {
    if (route.query.signup === 'true') {
        isSignUp.value = true
    }
})

const onSubmit = async () => {
    error.value = ''

    try {
        if (isSignUp.value) {
            await signUp(email.value, password.value)
            alert('Registration successful. Please sign in.')
            isSignUp.value = false
        } else {
            await authStore.signIn(email.value, password.value)

            // Redirect to the redirect path if it exists, otherwise to the projects page
            const redirectPath = route.query.redirect
                ? route.query.redirect.toString()
                : PATH_PROJECTS

            router.push(redirectPath)
        }
    } catch (err) {
        error.value = isSignUp.value
            ? 'Registration failed. Please check your input.'
            : 'Login failed. Please check your credentials.'
    }
}
</script>

<style scoped>
h1 {
    text-align: center;
    margin-bottom: 20px;
    color: var(--color-text-light);
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

.tabs {
    display: flex;
    margin-bottom: 20px;
    border-bottom: 1px solid rgba(255, 255, 255, 0.2);
}

.tabs button {
    flex: 1;
    background: none;
    border: none;
    padding: 10px;
    color: var(--color-text-light);
    opacity: 0.7;
    cursor: pointer;
    font-size: 16px;
    transition: all 0.3s;
}

.tabs button.active {
    opacity: 1;
    border-bottom: 2px solid var(--color-primary);
}

.error {
    color: #ff6b6b;
    margin-top: 15px;
    text-align: center;
}
</style>