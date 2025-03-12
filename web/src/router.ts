import { createRouter, createWebHistory } from "vue-router";
import type { RouteRecordRaw } from "vue-router";
import { useAuthStore } from "./stores/auth_store";

export const PATH_HOME = "/";
export const PATH_SIGN_IN = "/sign_in";
export const PATH_PROJECTS = "/projects";
export const PATH_PROJECT_DETAIL = "/projects/:id";

const routes: Array<RouteRecordRaw> = [
  // non auth required routes
  {
    path: PATH_HOME,
    name: "Home",
    component: () => import("./views/Home.vue"),
  },
  {
    path: PATH_SIGN_IN,
    name: "SignIn",
    component: () => import("./views/SignIn.vue"),
  },
  // auth required routes
  {
    path: PATH_PROJECTS,
    name: "Projects",
    component: () => import("./views/Projects.vue"),
    meta: { requiresAuth: true },
  },
  {
    path: PATH_PROJECT_DETAIL,
    name: "ProjectDetail",
    component: () => import("./views/ProjectDetail.vue"),
    meta: { requiresAuth: true },
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
