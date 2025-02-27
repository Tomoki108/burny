import { defineStore } from "pinia";
import { signIn } from "../api/auth";

export const useAuthStore = defineStore("auth", {
  state: () => ({
    token: "" as string, // JWT token を格納。ログインしていなければ空文字
  }),
  getters: {
    isAuthenticated: (state): boolean => !!state.token,
  },
  actions: {
    async signIn(email: string, password: string) {
      try {
        const data = await signIn(email, password);
        this.token = data.token;
      } catch (error) {
        // エラーハンドリングは適宜
        throw error;
      }
    },
    logout() {
      this.token = "";
    },
    // 以降のAPIリクエストでトークンを利用するためのヘッダ生成ヘルパー例
    getAuthHeader(): HeadersInit {
      return {
        "Content-Type": "application/json",
        Authorization: `Bearer ${this.token}`,
      };
    },
  },
});
