import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";
import ExecutionView from "../views/ExecutionView.vue";
import LoginView from "../views/LoginView.vue";
import { requireAuth } from "./guards";
import RegisterView from "../views/RegisterView.vue";
import AdminView from "../views/AdminView.vue";

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: "/", component: HomeView },
    { path: "/login", component: LoginView },
    { path: "/register", component: RegisterView },

    {
      path: "/visualize",
      component: ExecutionView,
      beforeEnter: requireAuth,
    },
    {
      path: "/admin",
      component: AdminView,
      beforeEnter: requireAuth,
    }
  ]
});

export default router;