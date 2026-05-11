<script setup lang="ts">
import { user, fetchUser } from "../store/user";
import { useRouter, useRoute } from "vue-router";
import { watch } from "vue";
import { getToken } from "../utils/token";
import { logout } from "../store/user";

const router = useRouter();
const route = useRoute();


function go(path: string) {
  try {
    router.push(path);
  } catch (e) {
    console.error(e);
  }
}

watch(
  () => route.fullPath,
  async () => {
    if (getToken() && !user.value) {
      try {
        await fetchUser();
      } catch (e) {
        console.error("Failed to fetch user in MainLayout:", e);
      }
    }
  },
  { immediate: true }
);


</script>

<template>
  <div class="layout">
    <header class="navbar">
      <div class="logo" @click="go('/')">AlgoLearn</div>

      <nav class="nav-links">
        <a @click="go('/visualize')">Визуализатор</a>
      </nav>

      <div class="auth">
        <template v-if="!user">
          <button class="btn btn-secondary" @click="go('/login')">Войти</button>
          <button class="btn btn-primary" @click="go('/register')">Регистрация</button>
        </template>
        <template v-else>
          <!-- Администратор -->
          <template v-if="user.role === 'admin'">
            <button class="btn btn-secondary" @click="go('/admin')">Админ-панель</button>
            <button class="btn btn-danger" @click="logout">Выйти</button>
          </template>
          <!-- Обычный пользователь -->
          <template v-else>
            <button class="btn btn-secondary" @click="go('/profile')">
              {{ user.user_name }}
            </button>
          </template>
        </template>
      </div>
    </header>

    <main class="main-content">
      <slot />
    </main>
  </div>
</template>

<style scoped>
.layout {
  min-height: 100vh;
  background: var(--bg);
  color: var(--text);
}

.navbar {
  height: 70px;
  background: var(--surface);
  border-bottom: 1px solid var(--border);
  display: flex;
  align-items: center;
  padding: 0 40px;
  gap: 40px;
  backdrop-filter: blur(12px);
}

.logo {
  font-size: 28px;
  font-weight: 800;
  background: linear-gradient(90deg, #6366f1, #a5b4fc);
  -webkit-background-clip: text;
  background-clip: text;
  -webkit-text-fill-color: transparent;
  cursor: pointer;
  user-select: none;
}

.nav-links {
  display: flex;
  gap: 28px;
  font-weight: 500;
}

.nav-links a {
  color: var(--text);
  text-decoration: none;
  transition: color 0.2s;
  cursor: pointer;
}

.nav-links a:hover {
  color: var(--primary);
}

.auth {
  margin-left: auto;
  display: flex;
  gap: 12px;
  align-items: center;
}

.main-content {
  padding: 40px;
  max-width: 1400px;
  margin: 0 auto;
}

/* Кнопки (можно оставить глобальными, но продублируем для автономности) */
.btn {
  padding: 10px 20px;
  border-radius: 8px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s ease;
  border: none;
  font-size: 14px;
}

.btn-primary {
  background: linear-gradient(135deg, #6366f1, #4f46e5);
  color: white;
  box-shadow: 0 4px 12px rgba(99, 102, 241, 0.3);
}

.btn-primary:hover {
  transform: translateY(-1px);
  box-shadow: 0 8px 20px rgba(99, 102, 241, 0.4);
}

.btn-secondary {
  background: rgba(197, 201, 205, 0.9);
  color: var(--text);
  border: 1px solid var(--border);
}

.btn-secondary:hover {
  background: #e2e8f0;
}

.btn-danger {
  background: #ef4444;
  color: white;
}
.btn-danger:hover {
  background: #dc2626;
}
</style>