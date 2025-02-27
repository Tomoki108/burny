import { createRouter, createWebHistory } from "vue-router";
import type { RouteRecordRaw } from "vue-router";
import { useAuthStore } from "../stores/auth";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/projects",
    name: "Projects",
    component: () => import("../views/Projects.vue"),
    meta: { requiresAuth: true },
  },
  {
    path: "/sign_in",
    name: "SignIn",
    component: () => import("../views/SignIn.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

// グローバルなナビゲーションガードで認証をチェック
router.beforeEach((to, from, next) => {
  const authStore = useAuthStore();
  if (to.meta.requiresAuth && !authStore.isAuthenticated) {
    next({ path: "/sign_in", query: { redirect: to.fullPath } });
  } else {
    next();
  }
});

export default router;
