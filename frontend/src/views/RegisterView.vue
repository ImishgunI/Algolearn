<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";
import { register } from "../api/auth";
import AuthLayout from "../components/AuthLayout.vue";

const name = ref("");
const surname = ref("");
const email = ref("");
const password = ref("");
const loading = ref(false);
const router = useRouter();

async function submit() {
  if (!name.value || !surname.value || !email.value || !password.value) {
    alert("Пожалуйста, заполните все поля");
    return;
  }
  loading.value = true;
  try {
    await register({
      name: name.value,
      surname: surname.value,
      email: email.value,
      password: password.value,
    });
    alert("Регистрация прошла успешно! Теперь вы можете войти.");
    router.push("/login");
  } catch (err: any) {
    alert(err.message || "Ошибка регистрации. Проверьте данные.");
  } finally {
    loading.value = false;
  }
}
</script>

<template>
  <AuthLayout>
    <div class="auth-card glass">
      <div class="logo">🚀</div>
      <h2>Присоединяйся</h2>
      <p class="subtitle">Начни визуализировать алгоритмы</p>

      <form @submit.prevent="submit">
        <div class="row">
          <input v-model="name" type="text" placeholder="Имя" required />
          <input v-model="surname" type="text" placeholder="Фамилия" required />
        </div>
        <input v-model="email" type="email" placeholder="Email" required />
        <input v-model="password" type="password" placeholder="Пароль" required />

        <button class="btn btn-primary" type="submit" :disabled="loading">
          <span v-if="!loading">Зарегистрироваться</span>
          <span v-else class="spinner"></span>
        </button>
      </form>

      <p class="switch">
        Уже есть аккаунт? <router-link to="/login">Войти</router-link>
      </p>
    </div>
  </AuthLayout>
</template>

<style scoped>
.auth-card {
  width: 100%;
  max-width: 460px;
  padding: 40px 32px;
  border-radius: var(--radius);
  background: var(--surface);
  backdrop-filter: blur(18px);
  border: 1px solid var(--surface-border);
  box-shadow: 0 20px 50px rgba(0,0,0,0.4);
  text-align: center;
}

.logo {
  font-size: 2.2rem;
  margin-bottom: 16px;
}

.subtitle {
  color: var(--text-muted);
  margin-bottom: 32px;
}

.row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 14px;
  margin-bottom: 14px;
}

input {
  width: 100%;
  margin-bottom: 14px;
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

@keyframes spin { to { transform: rotate(360deg); } }

.switch {
  margin-top: 24px;
  color: var(--text-muted);
}

.switch a {
  color: var(--accent);
  text-decoration: none;
  font-weight: 500;
}
</style>