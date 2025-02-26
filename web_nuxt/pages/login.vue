<template>
  <v-container class="fill-height" fluid>
    <v-row align="center" justify="center">
      <v-col cols="12" sm="8" md="4">
        <v-card>
          <v-card-title class="headline">Login</v-card-title>
          <v-card-text>
            <v-form @submit.prevent="onSubmit">
              <v-text-field v-model="email" label="Email" placeholder="youremail@burnyapp.io" required></v-text-field>
              <v-text-field v-model="password" label="Password" type="password" placeholder="Password"
                required></v-text-field>
              <v-btn type="submit" color="primary" block>Sign In</v-btn>
              <v-alert v-if="error" type="error" dense text>{{ error }}</v-alert>
            </v-form>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<!-- <script lang="ts">
export default {
  layout: 'empty',
  head() {
    return {
      title: 'Login',
    }
  },
}
</script> -->

<script type="setup" lang="ts">
import AppLayout from '@/Layouts/AppLayout';
import { ref } from 'vue'
import { useNuxtApp } from '#app'

defineOptions({ layout: "empty" })

const email = ref('')
const password = ref('')
const error = ref('')

const { $axios } = useNuxtApp()

const onSubmit = async () => {
  error.value = ''
  try {
    await $axios.post('/api/login', {
      email: email.value,
      password: password.value,
    })
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
  background-color: #fafafa;
  /* 全体の背景色 */
}

/* 背景の円を表現するための要素 */
.background-circle {
  position: absolute;
  top: -200px;
  left: -200px;
  width: 600px;
  height: 600px;
  background-color: #efefef;
  /* 円の色 */
  border-radius: 50%;
  z-index: -1;
  /* 背景として下に配置 */
}

/* 中央配置用のコンテナ */
.container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  width: 100%;
}

/* ログインフォームの枠 */
.login-box {
  background-color: #ffffff;
  width: 320px;
  padding: 40px 30px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  border-radius: 8px;
  text-align: center;
}

/* タイトル */
.login-box h1 {
  margin-bottom: 24px;
  font-size: 24px;
}

/* 入力フィールド */
.login-box input[type="text"],
.login-box input[type="password"] {
  width: 100%;
  padding: 12px;
  margin-bottom: 16px;
  font-size: 16px;
  border: 1px solid #ccc;
  border-radius: 4px;
}

/* ボタン共通スタイル */
.login-box button {
  width: 100%;
  padding: 12px;
  font-size: 16px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

/* Sign Inボタン */
.btn-signin {
  background-color: #d8b8c4;
  /* お好みのピンク系の色 */
  color: #fff;
  margin-bottom: 16px;
}

/* Create accountボタン */
.btn-create {
  background-color: #f2d5df;
  /* Sign Inより少し淡いピンク */
  color: #333;
}

/* "or" の部分 */
.or {
  margin: 16px 0;
  color: #666;
}
</style>
