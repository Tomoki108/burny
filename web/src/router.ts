import { createRouter, createWebHistory } from "vue-router";
import type { RouteRecordRaw } from "vue-router";
import { useAuthStore } from "./stores/auth_store";

export const PATH_HOME = "/";
export const PATH_PROJECTS = "/projects";
export const PATH_PROJECT_DETAIL = "/projects/:id";

const routes: Array<RouteRecordRaw> = [
  // non auth required routes
  {
    path: PATH_HOME,
    name: "Home",
    component: () => import("./views/Home.vue"),
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
    // Redirect to home page where the modal can be shown instead
    next({ path: PATH_HOME, query: { auth: "required" } });
  } else {
    next();
  }
});

export default router;
