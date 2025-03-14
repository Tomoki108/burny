<template>
    <div class="modal-overlay" v-if="isVisible" @click.self="close">
        <div class="login-box">
            <h1 class="color-white">Burny 🐶</h1>
            <div class="tabs">
                <button @click="isSignUp = false" :class="{ active: !isSignUp }">Sign In</button>
                <button @click="isSignUp = true" :class="{ active: isSignUp }">Sign Up</button>
            </div>
            <form @submit.prevent="onSubmit">
                <label for="email">Email</label>
                <input type="text" id="email" v-model="email" placeholder="yourname@burny.page" required />
                <label for="password">Password</label>
                <input type="password" id="password" v-model="password" placeholder="************" required
                    minlength="8" maxlength="20" />
                <button type="submit" class="button">{{ isSignUp ? 'Sign Up' : 'Sign In' }}</button>
            </form>
            <p v-if="error" class="error">{{ error }}</p>
            <button class="close-button" @click="close">×</button>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { ref, defineProps, defineEmits, watch } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth_store'
import { PATH_PROJECTS } from '../router'
import { signUp } from '../api/auth_api'
import { ErrorResponse } from '../api/api_helper'

const props = defineProps({
    isVisible: {
        type: Boolean,
        required: true
    },
    initialSignUp: {
        type: Boolean,
        default: false
    }
})

const emit = defineEmits(['close', 'auth-success'])

const email = ref('')
const password = ref('')
const error = ref('')
const isSignUp = ref(props.initialSignUp)
const router = useRouter()
const authStore = useAuthStore()

// Watch for changes in the initialSignUp prop
watch(() => props.initialSignUp, (newVal) => {
    isSignUp.value = newVal
})

const onSubmit = async () => {
    error.value = ''

    try {
        if (isSignUp.value) {
            const response = await signUp(email.value, password.value)
            if (response instanceof ErrorResponse) {
                error.value = response.getMessage()
            } else {
                alert('Registration successful. Please sign in.')
                isSignUp.value = false
            }
        } else {
            await authStore.signIn(email.value, password.value)
            emit('auth-success')
            router.push(PATH_PROJECTS)
        }
    } catch (err) {
        error.value = isSignUp.value
            ? 'Registration failed. Please check your input.'
            : 'Login failed. Please check your credentials.'
    }
}

const close = () => {
    emit('close')
    // Clear form data
    email.value = ''
    password.value = ''
    error.value = ''
}
</script>

<style scoped>
h1 {
    text-align: center;
    margin-bottom: 20px;
    color: var(--color-text-light);
}

/* Modal styling */
.modal-overlay {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.5);
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 100;
}

.login-box {
    position: relative;
    background: linear-gradient(var(--color-tertiary), var(--color-secondary));
    padding: 40px;
    border-radius: 8px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.3);
    width: 400px;
    max-width: 90%;
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
    color: var(--color-error);
    font-size: var(--font-size-base);
    margin-top: 15px;
    text-align: center;
}

.close-button {
    position: absolute;
    top: 10px;
    right: 10px;
    background: none;
    border: none;
    color: var(--color-text-light);
    font-size: 24px;
    cursor: pointer;
    opacity: 0.7;
    transition: opacity 0.3s;
}

.close-button:hover {
    opacity: 1;
}

/* Responsive adjustments */
@media (max-width: 480px) {
    .login-box {
        padding: 20px;
    }
}
</style>