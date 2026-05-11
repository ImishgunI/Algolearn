<script setup lang="ts">
import MainLayout from "../components/MainLayout.vue";
import { ref,  onMounted } from "vue";
import { user, fetchUser, logout } from "../store/user";
import { updateProfile, updatePassword, getStats, getFavorites } from "../api/profile";
import type { Lesson } from "../api/lessons";

// ── Вкладки ──
type Tab = "profile" | "stats" | "favorites";
const activeTab = ref<Tab>("profile");

// ── Редактирование профиля ──
const isEditing = ref(false);
const editForm = ref({ user_name: "", user_surname: "", email: "" });
const editError = ref("");

// ── Смена пароля ──
const showPasswordForm = ref(false);
const passwordForm = ref({ old_password: "", new_password: "", confirm_password: "" });
const passwordError = ref("");
const passwordSuccess = ref("");

// ── Статистика ──
const stats = ref({ algorithms_executed: 0, favorites_count: 0 });
const loadingStats = ref(false);

// ── Избранное ──
const favorites = ref<Lesson[]>([]);
const loadingFavorites = ref(false);

// ── Загрузка данных при переключении вкладок ──
async function loadStats() {
  loadingStats.value = true;
  try {
    stats.value = await getStats();
  } catch (e) {
    console.error(e);
  } finally {
    loadingStats.value = false;
  }
}

async function loadFavorites() {
  loadingFavorites.value = true;
  try {
    favorites.value = await getFavorites();
  } catch (e) {
    console.error(e);
  } finally {
    loadingFavorites.value = false;
  }
}

// ── Действия ──
function startEdit() {
  if (!user.value) return;
  editForm.value = {
    user_name: user.value.user_name || "",
    user_surname: user.value.user_surname || "",
    email: user.value.email || "",
  };
  isEditing.value = true;
  editError.value = "";
}

async function saveProfile() {
  editError.value = "";
  try {
    await updateProfile(editForm.value);
    await fetchUser();                // обновить глобальный user
    isEditing.value = false;
  } catch (e: any) {
    editError.value = e?.message || "Ошибка сохранения";
  }
}

function cancelEdit() {
  isEditing.value = false;
  editError.value = "";
}

async function changePassword() {
  passwordError.value = "";
  passwordSuccess.value = "";
  if (passwordForm.value.new_password !== passwordForm.value.confirm_password) {
    passwordError.value = "Новые пароли не совпадают";
    return;
  }
  try {
    await updatePassword({
      old_password: passwordForm.value.old_password,
      new_password: passwordForm.value.new_password,
    });
    passwordSuccess.value = "Пароль успешно изменён";
    passwordForm.value = { old_password: "", new_password: "", confirm_password: "" };
    showPasswordForm.value = false;
  } catch (e: any) {
    passwordError.value = e?.message || "Ошибка смены пароля";
  }
}

// ── Инициализация ──
onMounted(() => {
  // Если пользователь не загружен, попытаться получить (но MainLayout уже делает watch)
  if (!user.value) {
    fetchUser().catch(() => {});
  }
});
</script>

<template>
  <MainLayout>
    <div class="profile-page">
      <h1>Личный кабинет</h1>

      <!-- Вкладки -->
      <div class="tabs">
        <button
          v-for="tab in (['profile','stats','favorites'] as Tab[])"
          :key="tab"
          :class="{ active: activeTab === tab }"
          @click="activeTab = tab; if (tab === 'stats') loadStats(); if (tab === 'favorites') loadFavorites()"
        >
          {{ tab === 'profile' ? 'Профиль' : tab === 'stats' ? 'Статистика' : 'Избранное' }}
        </button>
      </div>

      <!-- Вкладка Профиль -->
      <div v-if="activeTab === 'profile'" class="profile-card glass">
        <div class="avatar">
          {{ user?.user_name?.[0] || user?.email?.[0] || "?" }}
        </div>

        <!-- Просмотр -->
        <div v-if="!isEditing" class="info">
          <h2>{{ user?.user_name }} {{ user?.user_surname }}</h2>
          <p class="email">{{ user?.email }}</p>
          <span class="role-badge">{{ user?.role || "пользователь" }}</span>
          <div class="actions">
            <button class="btn btn-primary" @click="startEdit">Редактировать</button>
            <button class="btn btn-secondary" @click="showPasswordForm = !showPasswordForm">
              {{ showPasswordForm ? 'Отмена' : 'Сменить пароль' }}
            </button>
            <button class="btn btn-danger" @click="logout">Выйти</button>
          </div>
        </div>

        <!-- Редактирование -->
        <form v-else class="edit-form" @submit.prevent="saveProfile">
          <div class="form-group">
            <label>Имя</label>
            <input v-model="editForm.user_name" type="text" required />
          </div>
          <div class="form-group">
            <label>Фамилия</label>
            <input v-model="editForm.user_surname" type="text" required />
          </div>
          <div class="form-group">
            <label>Email</label>
            <input v-model="editForm.email" type="email" required />
          </div>
          <div v-if="editError" class="error-message">{{ editError }}</div>
          <div class="actions">
            <button class="btn btn-primary" type="submit">Сохранить</button>
            <button class="btn btn-secondary" type="button" @click="cancelEdit">Отмена</button>
          </div>
        </form>

        <!-- Смена пароля -->
        <form v-if="showPasswordForm" class="password-form" @submit.prevent="changePassword">
          <h3>Смена пароля</h3>
          <div class="form-group">
            <label>Старый пароль</label>
            <input v-model="passwordForm.old_password" type="password" required />
          </div>
          <div class="form-group">
            <label>Новый пароль</label>
            <input v-model="passwordForm.new_password" type="password" required minlength="6" />
          </div>
          <div class="form-group">
            <label>Подтвердите новый пароль</label>
            <input v-model="passwordForm.confirm_password" type="password" required minlength="6" />
          </div>
          <div v-if="passwordError" class="error-message">{{ passwordError }}</div>
          <div v-if="passwordSuccess" class="success-message">{{ passwordSuccess }}</div>
          <button class="btn btn-primary" type="submit">Обновить пароль</button>
        </form>
      </div>

      <!-- Вкладка Статистика -->
      <div v-if="activeTab === 'stats'" class="stats-card glass">
        <h2>Статистика</h2>
        <div v-if="loadingStats">Загрузка...</div>
        <div v-else class="stats-grid">
          <div class="stat-item">
            <span class="stat-number">{{ stats.algorithms_executed }}</span>
            <span class="stat-label">запусков алгоритмов</span>
          </div>
          <div class="stat-item">
            <span class="stat-number">{{ stats.favorites_count }}</span>
            <span class="stat-label">избранных уроков</span>
          </div>
        </div>
      </div>

      <!-- Вкладка Избранное -->
      <div v-if="activeTab === 'favorites'" class="favorites-card glass">
        <h2>Избранные уроки</h2>
        <div v-if="loadingFavorites">Загрузка...</div>
        <ul v-else-if="favorites.length">
          <li v-for="fav in favorites" :key="fav.id">
            <router-link :to="`/courses/sorting/${fav.id}`">{{ fav.title }}</router-link>
          </li>
        </ul>
        <p v-else>Нет избранных уроков</p>
      </div>
    </div>
  </MainLayout>
</template>

<style scoped>
.profile-page {
  max-width: 700px;
  margin: 40px auto;
}

h1 {
  margin-bottom: 24px;
}

.tabs {
  display: flex;
  gap: 8px;
  margin-bottom: 24px;
}

.tabs button {
  background: var(--surface);
  border: 1px solid var(--surface-border);
  color: var(--text);
  padding: 10px 20px;
  border-radius: 8px;
  cursor: pointer;
  font-weight: 500;
  transition: 0.2s;
}

.tabs button.active {
  background: var(--primary);
  color: white;
}

.glass {
  background: var(--surface);
  backdrop-filter: blur(12px);
  border: 1px solid var(--surface-border);
  border-radius: var(--radius);
  padding: 32px;
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
  text-align: center;
  margin-bottom: 24px;
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

.actions {
  margin-top: 24px;
  display: flex;
  gap: 12px;
  justify-content: center;
  flex-wrap: wrap;
}

.btn {
  padding: 10px 20px;
  border-radius: 8px;
  font-weight: 600;
  cursor: pointer;
  border: none;
  transition: 0.2s;
}

.btn-primary { background: var(--primary); color: white; }
.btn-primary:hover { background: var(--primary-hover); }
.btn-secondary { background: var(--surface2); color: var(--text); }
.btn-secondary:hover { background: #e2e8f0; }
.btn-danger { background: #ef4444; color: white; }
.btn-danger:hover { background: #dc2626; }

.edit-form, .password-form {
  text-align: left;
  margin-top: 20px;
}

.form-group {
  margin-bottom: 16px;
}

.form-group label {
  display: block;
  margin-bottom: 6px;
  font-weight: 500;
  color: var(--text);
}

.form-group input {
  width: 100%;
  padding: 12px;
  border: 1px solid var(--border);
  border-radius: 8px;
  background: var(--surface);
  color: var(--text);
  font-size: 15px;
}

.error-message {
  color: #ef4444;
  margin-bottom: 12px;
  font-size: 14px;
}

.success-message {
  color: #10b981;
  margin-bottom: 12px;
  font-size: 14px;
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(160px, 1fr));
  gap: 20px;
  margin-top: 20px;
}

.stat-item {
  background: rgba(0,0,0,0.02);
  border-radius: 12px;
  padding: 20px;
  text-align: center;
}

.stat-number {
  display: block;
  font-size: 2.5rem;
  font-weight: 700;
  color: var(--primary);
}

.stat-label {
  color: var(--text-muted);
  font-size: 0.95rem;
}

.favorites-card ul {
  list-style: none;
  padding: 0;
}

.favorites-card li {
  margin-bottom: 10px;
}

.favorites-card a {
  color: var(--primary);
  text-decoration: none;
}
.favorites-card a:hover {
  text-decoration: underline;
}
</style>