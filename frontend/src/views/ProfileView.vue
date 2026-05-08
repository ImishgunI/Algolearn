<script setup lang="ts">
import MainLayout from "../components/MainLayout.vue";
import { user, logout } from "../store/user";
</script>

<template>
  <MainLayout>
    <div class="profile-page">
      <h1>Профиль</h1>

      <div class="profile-card glass">
        <div class="avatar">
          {{ user?.user_name?.[0] || user?.email?.[0] || "?" }}
        </div>

        <div class="info">
          <h2>{{ user?.user_name }} {{ user?.user_surname }}</h2>
          <p class="email">{{ user?.email }}</p>
          <span class="role-badge">{{ user?.role || "пользователь" }}</span>
        </div>

        <button class="btn btn-danger" @click="logout">
          Выйти из аккаунта
        </button>
      </div>
    </div>
  </MainLayout>
</template>

<style scoped>
.profile-page {
  max-width: 500px;
  margin: 40px auto;
}

h1 {
  margin-bottom: 24px;
}

.profile-card {
  background: var(--surface);
  backdrop-filter: blur(12px);
  border: 1px solid var(--surface-border);
  border-radius: var(--radius);
  padding: 40px 32px;
  text-align: center;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.05);
}

.avatar {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  background: linear-gradient(135deg, #6366f1, #a5b4fc);
  color: white;
  font-size: 32px;
  font-weight: 700;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 auto 20px;
}

.info {
  margin-bottom: 32px;
}

.info h2 {
  margin: 0 0 8px;
  font-size: 1.8rem;
}

.email {
  color: var(--text-muted);
  margin-bottom: 12px;
}

.role-badge {
  display: inline-block;
  background: rgba(99, 102, 241, 0.1);
  color: var(--primary);
  padding: 4px 16px;
  border-radius: 30px;
  font-weight: 600;
  font-size: 14px;
}

.btn {
  padding: 12px 28px;
  border-radius: 8px;
  font-weight: 600;
  cursor: pointer;
  border: none;
  transition: all 0.2s;
}

.btn-danger {
  background: #ef4444;
  color: white;
}

.btn-danger:hover {
  background: #dc2626;
}
</style>