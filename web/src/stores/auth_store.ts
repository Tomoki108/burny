import { defineStore } from "pinia";
import { signIn } from "../api/auth_api";
import { jwtDecode } from "jwt-decode";

interface DecodedToken {
  exp: number;
  email?: string;
  sub?: string;
}

export const useAuthStore = defineStore("auth", {
  state: () => ({
    token: localStorage.getItem("token") || "",
    email: localStorage.getItem("email") || "",
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
        this.email = email;
        localStorage.setItem("token", data.token);
        localStorage.setItem("email", email);
      } catch (error) {
        throw error;
      }
    },
    signOut() {
      this.token = "";
      this.email = "";
      localStorage.removeItem("token");
      localStorage.removeItem("email");
    },
  },
});

function isTokenExpired(token: string): boolean {
  const decoded = jwtDecode<DecodedToken>(token);
  return decoded.exp * 1000 < Date.now();
}
