<script setup lang="ts">
import { user, logout } from "../store/user";
import { useRouter } from "vue-router";

const router = useRouter();

function go(path: string) {
  router.push(path);
}
</script>

<template>
  <div class="layout">
    <header class="navbar">
      <div class="logo" @click="go('/')">AlgoLearn</div>

      <nav class="nav-links">
        <a @click="go('/')">Главная</a>
        <a @click="go('/visualize')">Визуализатор</a>
      </nav>

      <div class="auth">
        <template v-if="!user">
          <button class="btn btn-secondary" @click="go('/login')">Войти</button>
          <button class="btn btn-primary" @click="go('/register')">Регистрация</button>
        </template>
        <template v-else>
          <button class="btn btn-secondary" @click="go('/profile')">
            {{ user.user_name }}
          </button>
          <button class="btn btn-danger" @click="logout">Выйти</button>
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
}

.logo {
  font-size: 28px;
  font-weight: 800;
  background: linear-gradient(90deg, #6366f1, #a5b4fc);
  -webkit-background-clip: text;
  background-clip: text;
  -webkit-text-fill-color: transparent;
  cursor: pointer;
}

.nav-links {
  display: flex;
  gap: 28px;
  font-weight: 500;
}

.nav-links a {
  color: var(--text-muted);
  text-decoration: none;
  transition: color 0.2s;
  cursor: pointer;
}

.nav-links a:hover {
  color: white;
}

.auth {
  margin-left: auto;
  display: flex;
  gap: 12px;
}

.main-content {
  padding: 40px;
  max-width: 1400px;
  margin: 0 auto;
}
</style>