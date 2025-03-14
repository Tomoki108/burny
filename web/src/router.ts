import { createRouter, createWebHashHistory } from "vue-router";
import type { RouteRecordRaw } from "vue-router";
import { useAuthStore } from "./stores/auth_store";

export const PATH_HOME = "/";
export const PATH_ACCOUNT = "/account";
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
    path: PATH_ACCOUNT,
    name: "Account",
    component: () => import("./views/Account.vue"),
  },
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
  // GCS静的ホスティングでの動作を改善するためにハッシュモードに変更
  history: createWebHashHistory(),
  routes,
});

// グローバルなナビゲーションガードで認証をチェック
router.beforeEach((to, _, next) => {
  const authStore = useAuthStore();
  if (to.meta.requiresAuth && !authStore.isAuthenticated) {
    // Redirect to home page where the modal can be shown instead
    next({ path: PATH_HOME, query: { auth: "required" } });
  } else {
    next();
  }
});

export default router;
