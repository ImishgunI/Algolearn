<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";
import { login } from "../api/auth";
import AuthLayout from "../components/AuthLayout.vue";
import { setToken } from "../utils/token"

const email = ref("");
const password = ref("");
const loading = ref(false);
const router = useRouter();

async function submit() {
  loading.value = true;
  try {
    const res = await login({ email: email.value, password: password.value });
    setToken(res.access_token);
    router.push("/");
  } catch (err) {
    alert("Ошибка входа");
  } finally {
    loading.value = false;
  }
}
</script>

<template>
  <AuthLayout>
    <div class="auth-card glass">
      <div class="logo">AlgoLearn</div>
      <h2>С возвращением</h2>
      <p class="subtitle">Продолжи изучение алгоритмов</p>

      <div class="input-wrapper">
        <span class="icon">📧</span>
        <input v-model="email" type="email" placeholder="Email" />
      </div>

      <div class="input-wrapper">
        <span class="icon">🔒</span>
        <input v-model="password" type="password" placeholder="Пароль" />
      </div>

      <button class="btn btn-primary" @click="submit" :disabled="loading">
        <span v-if="!loading">Войти</span>
        <span v-else class="spinner"></span>
      </button>

      <p class="switch">
        Нет аккаунта? <router-link to="/register">Создать</router-link>
      </p>
    </div>
  </AuthLayout>
</template>

<style scoped>
.auth-card {
  width: 100%;
  max-width: 420px;
  padding: 40px 32px;
  border-radius: var(--radius);
  background: var(--surface);
  backdrop-filter: blur(18px);
  border: 1px solid var(--surface-border);
  box-shadow: 0 20px 50px rgba(0, 0, 0, 0.4);
  text-align: center;
}

.logo {
  font-size: 2rem;
  margin-bottom: 20px;
}

h2 {
  margin: 0;
}

.subtitle {
  color: var(--text-muted);
  margin-bottom: 32px;
}

.input-wrapper {
  position: relative;
  margin-bottom: 18px;
}

.input-wrapper .icon {
  position: absolute;
  left: 16px;
  top: 50%;
  transform: translateY(-50%);
  font-size: 18px;
  opacity: 0.7;
}

.input-wrapper input {
  padding-left: 48px;
  width: 100%;
}

.btn {
  width: 100%;
  margin-top: 10px;
  padding: 14px;
}

.spinner {
  width: 20px;
  height: 20px;
  border: 2px solid rgba(255,255,255,0.3);
  border-top-color: white;
  border-radius: 50%;
  animation: spin 0.7s linear infinite;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}

.switch {
  margin-top: 24px;
  color: var(--text-muted);
}

.switch a {
  color: var(--accent);
  text-decoration: none;
  font-weight: 500;
}

.switch a:hover {
  text-decoration: underline;
}
</style>