import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";
import ExecutionView from "../views/ExecutionView.vue";
import LoginView from "../views/LoginView.vue";
import { requireAuth } from "./guards";
import RegisterView from "../views/RegisterView.vue";
import AdminView from "../views/AdminView.vue";
import ProfileView from "../views/ProfileView.vue";
import SortingCourse from "../views/SortingCourse.vue";
import LessonView from "../views/LessonView.vue";

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: "/", component: HomeView },
    { path: "/login", component: LoginView },
    { path: "/register", component: RegisterView },
    { path: "/profile", component: ProfileView },
    {
      path: "/visualize",
      component: ExecutionView,
      beforeEnter: requireAuth,
    },
    {
      path: "/admin",
      component: AdminView,
      beforeEnter: requireAuth,
    }, 
    {
      path: "/courses/sorting",
      component: SortingCourse,
      children: [
        {
          path: ":lessonId",
          component: LessonView,
          props: true,
        },
        {
          path: "",
          redirect: "/courses/sorting/intro",
        },
      ],
    },
  ]
});

export default router;