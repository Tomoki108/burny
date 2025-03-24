import { defineStore } from "pinia";
import { signIn } from "../api/auth_api";
import { jwtDecode } from "jwt-decode";
import { ErrorResponse } from "../api/api_helper";

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
        const res = await signIn(email, password);
        if (res instanceof ErrorResponse) {
          throw new Error(res.getMessage());
        }

        this.token = res.token;
        this.email = email;
        localStorage.setItem("token", res.token);
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
