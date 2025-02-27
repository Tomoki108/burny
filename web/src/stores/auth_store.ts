import { defineStore } from "pinia";
import { signIn } from "../api/auth_api";
import { jwtDecode } from "jwt-decode";

interface DecodedToken {
  exp: number;
}

export const useAuthStore = defineStore("auth", {
  state: () => ({
    token: localStorage.getItem("token") || "",
  }),
  getters: {
    isAuthenticated: (state): boolean =>
      !!state.token && !isTokenExpired(state.token),
  },
  actions: {
    async signIn(email: string, password: string) {
      try {
        const data = await signIn(email, password);
        this.token = data.token;
        localStorage.setItem("token", data.token);
      } catch (error) {
        throw error;
      }
    },
    signOut() {
      this.token = "";
      localStorage.removeItem("token");
    },
  },
});

function isTokenExpired(token: string): boolean {
  const decoded = jwtDecode<DecodedToken>(token);
  return decoded.exp * 1000 < Date.now();
}
