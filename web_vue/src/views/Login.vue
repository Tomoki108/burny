<template>
    <div class="container">
        <div class="background-circle"></div>
        <div class="login-box">
            <h1>Burny</h1>
            <form @submit.prevent="onSubmit">
                <label for="email">Username</label>
                <input type="text" id="email" v-model="email" placeholder="yourname@burnuppro.io" />

                <label for="password">Password</label>
                <input type="password" id="password" v-model="password" placeholder="************" />

                <button type="submit" class="btn-signin">Sign In</button>
                <div class="or">or</div>
                <button type="button" class="btn-create">Create account</button>
            </form>
            <p v-if="error" class="error">{{ error }}</p>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'

const email = ref('')
const password = ref('')
const error = ref('')

const onSubmit = async () => {
    error.value = ''
    try {
        const response = await fetch('http://localhost:1323/api/v1/sign_in', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({
                email: email.value,
                password: password.value,
            }),
        })
        if (!response.ok) {
            throw new Error('Login failed')
        }
        // 成功時の処理（例: リダイレクトなど）
    } catch (err) {
        error.value = 'Login failed. Please check your credentials.'
    }
}
</script>

<style scoped>
* {
    box-sizing: border-box;
    margin: 0;
    padding: 0;
}

body {
    font-family: sans-serif;
    background-color: #ffffff;
    /* 背景色を白に変更 */
}

.background-circle {
    position: absolute;
    top: -200px;
    left: -200px;
    width: 600px;
    height: 600px;
    background-color: #efefef;
    border-radius: 50%;
    z-index: -1;
}

.container {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
    width: 100%;
}

.login-box {
    background-color: #f0f0f0;
    /* フォームエリアの背景色をグレーに変更 */
    width: 320px;
    padding: 40px 30px;
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
    border-radius: 8px;
    text-align: center;
}

.login-box h1 {
    margin-bottom: 24px;
    font-size: 24px;
    color: #000000;
    /* "Burny" テキストを黒に変更 */
}

.login-box input[type="text"],
.login-box input[type="password"] {
    width: 100%;
    padding: 12px;
    margin-bottom: 16px;
    font-size: 16px;
    border: 1px solid #ccc;
    border-radius: 4px;
    color: #000;
    /* テキストカラーを黒に設定 */
}

.login-box button {
    width: 100%;
    padding: 12px;
    font-size: 16px;
    border: none;
    border-radius: 4px;
    cursor: pointer;
}

.btn-signin {
    background-color: #d8b8c4;
    color: #fff;
    margin-bottom: 16px;
}

.btn-create {
    background-color: #f2d5df;
    color: #fff;
}

.or {
    margin: 16px 0;
    color: #666;
}

.error {
    color: red;
    margin-top: 16px;
}
</style>