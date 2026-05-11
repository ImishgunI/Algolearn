<script setup lang="ts">
import MainLayout from "../components/MainLayout.vue"; 
import { ref, onMounted } from "vue";
import { http } from "../api/http";

type Tab = "courses" | "lessons" | "users" | "comments" | "stats";
const activeTab = ref<Tab>("courses");

// Данные для вкладок
const courses = ref<any[]>([]);
const lessons = ref<any[]>([]);
const users = ref<any[]>([]);
const comments = ref<any[]>([]);
const stats = ref<any>({});

// Формы
const newCourse = ref({ title: "", description: "" });
const editingCourse = ref<any>(null);

const newLesson = ref({
  course_id: 1,
  title: "",
  theory: "",
  algorithm_type: "",
  order_index: 1,
});
const editingLesson = ref<any>(null);

const loading = ref(false);

// Загрузка данных по вкладке
async function loadCourses() {
  loading.value = true;
  try { courses.value = await http("/api/admin/courses"); }
  finally { loading.value = false; }
}

async function loadLessons(courseId?: number) {
  if (!courseId) { lessons.value = []; return; }
  loading.value = true;
  try { lessons.value = await http(`/api/admin/courses/${courseId}/lessons`); }
  finally { loading.value = false; }
}

async function loadUsers() {
  loading.value = true;
  try { users.value = await http("/api/admin/users"); }
  finally { loading.value = false; }
}

async function loadComments() {
  loading.value = true;
  try { comments.value = await http("/api/admin/comments"); }
  catch (e) { console.error(e); }
  finally { loading.value = false; }
}

async function loadStats() {
  loading.value = true;
  try { stats.value = await http("/api/admin/stats"); }
  finally { loading.value = false; }
}

async function switchTab(tab: Tab) {
  activeTab.value = tab;
  if (tab === "courses") await loadCourses();
  else if (tab === "users") await loadUsers();
  else if (tab === "stats") await loadStats();
  else if (tab === "comments") await loadComments();
  else if (tab === "lessons" && courses.value.length > 0) await loadLessons(courses.value[0].id);
}

// CRUD операции (без изменений)
async function createCourse() {
  await http("/api/admin/courses", { method: "POST", body: JSON.stringify(newCourse.value) });
  newCourse.value = { title: "", description: "" };
  await loadCourses();
}
async function updateCourse(course: any) {
  await http(`/api/admin/courses/${course.id}`, { method: "PUT", body: JSON.stringify(course) });
  editingCourse.value = null;
  await loadCourses();
}
async function deleteCourse(id: number) {
  await http(`/api/admin/courses/${id}`, { method: "DELETE" });
  await loadCourses();
}
async function createLesson(courseId: number) {
  await http(`/api/admin/courses/${courseId}/lessons`, { method: "POST", body: JSON.stringify(newLesson.value) });
  newLesson.value = { course_id: courseId, title: "", theory: "", algorithm_type: "", order_index: 1 };
  await loadLessons(courseId);
}
async function updateLesson(lesson: any) {
  await http(`/api/admin/lessons/${lesson.id}`, { method: "PUT", body: JSON.stringify(lesson) });
  editingLesson.value = null;
  await loadLessons(lesson.course_id);
}
async function deleteLesson(lessonId: number, courseId: number) {
  await http(`/api/admin/lessons/${lessonId}`, { method: "DELETE" });
  await loadLessons(courseId);
}
async function updateUserRole(userId: number, newRole: string) {
  await http(`/api/admin/users/${userId}/role`, { method: "PUT", body: JSON.stringify({ role: newRole }) });
  await loadUsers();
}
async function deleteComment(commentId: number) {
  await http(`/api/admin/comments/${commentId}`, { method: "DELETE" });
  await loadComments();
}

onMounted(() => loadCourses());
</script>

<template>
  <MainLayout>
    <div class="admin-page">
      <h1>Админ-панель</h1>

      <div class="tabs">
        <button v-for="tab in (['courses','lessons','users','comments','stats'] as Tab[])" :key="tab"
          :class="{ active: activeTab === tab }" @click="switchTab(tab)">
          {{ tab === 'courses' ? 'Курсы' : tab === 'lessons' ? 'Уроки' : tab === 'users' ? 'Пользователи' : tab === 'comments' ? 'Комментарии' : 'Статистика' }}
        </button>
      </div>

      <div v-if="loading" class="loading">Загрузка...</div>

      <!-- КУРСЫ -->
      <div v-if="activeTab === 'courses' && !loading" class="section-card">
        <h2>Курсы</h2>
        <div class="form-row">
          <input v-model="newCourse.title" placeholder="Название курса" class="modern-input" />
          <input v-model="newCourse.description" placeholder="Описание" class="modern-input" />
          <button class="btn-primary" @click="createCourse">Создать курс</button>
        </div>
        <div class="list">
          <div v-for="course in courses" :key="course.id" class="list-item">
            <div class="item-info">
              <strong>{{ course.title }}</strong> — {{ course.description }}
            </div>
            <div class="item-actions">
              <button class="btn-secondary" @click="editingCourse = course">Редактировать</button>
              <button class="btn-danger" @click="deleteCourse(course.id)">Удалить</button>
            </div>
          </div>
        </div>
        <div v-if="editingCourse" class="modal">
          <h3>Редактировать курс</h3>
          <input v-model="editingCourse.title" class="modern-input" />
          <input v-model="editingCourse.description" class="modern-input" />
          <div class="modal-actions">
            <button class="btn-primary" @click="updateCourse(editingCourse)">Сохранить</button>
            <button class="btn-secondary" @click="editingCourse = null">Отмена</button>
          </div>
        </div>
      </div>

      <!-- УРОКИ -->
      <div v-if="activeTab === 'lessons' && !loading" class="section-card">
        <h2>Уроки</h2>
        <div class="form-row">
          <select v-model="newLesson.course_id" class="modern-input">
            <option v-for="c in courses" :key="c.id" :value="c.id">{{ c.title }}</option>
          </select>
          <input v-model="newLesson.title" placeholder="Название урока" class="modern-input" />
          <textarea v-model="newLesson.theory" placeholder="Теория" class="modern-input" rows="2"></textarea>
          <input v-model="newLesson.algorithm_type" placeholder="Тип алгоритма" class="modern-input" />
          <input v-model.number="newLesson.order_index" type="number" placeholder="Порядок" class="modern-input" style="width:80px" />
          <button class="btn-primary" @click="createLesson(newLesson.course_id)">Добавить урок</button>
        </div>
        <div v-for="course in courses" :key="course.id" class="lesson-block">
          <h3>{{ course.title }} <button class="btn-secondary small" @click="loadLessons(course.id)">Загрузить уроки</button></h3>
          <div class="list">
            <div v-for="lesson in lessons.filter(l => l.course_id === course.id)" :key="lesson.id" class="list-item">
              <div class="item-info">
                <strong>{{ lesson.title }}</strong> ({{ lesson.algorithm_type }})
              </div>
              <div class="item-actions">
                <button class="btn-secondary" @click="editingLesson = lesson">Редактировать</button>
                <button class="btn-danger" @click="deleteLesson(lesson.id, lesson.course_id)">Удалить</button>
              </div>
            </div>
          </div>
        </div>
        <div v-if="editingLesson" class="modal">
          <h3>Редактировать урок</h3>
          <input v-model="editingLesson.title" class="modern-input" />
          <textarea v-model="editingLesson.theory" class="modern-input"></textarea>
          <input v-model="editingLesson.algorithm_type" class="modern-input" />
          <input v-model.number="editingLesson.order_index" type="number" class="modern-input" style="width:80px" />
          <div class="modal-actions">
            <button class="btn-primary" @click="updateLesson(editingLesson)">Сохранить</button>
            <button class="btn-secondary" @click="editingLesson = null">Отмена</button>
          </div>
        </div>
      </div>

      <!-- ПОЛЬЗОВАТЕЛИ -->
      <div v-if="activeTab === 'users' && !loading" class="section-card">
        <h2>Пользователи</h2>
        <div class="table-responsive">
          <table class="modern-table">
            <thead>
              <tr><th>ID</th><th>Имя</th><th>Email</th><th>Роль</th><th>Действия</th></tr>
            </thead>
            <tbody>
              <tr v-for="u in users" :key="u.id">
                <td>{{ u.id }}</td>
                <td>{{ u.user_name }} {{ u.user_surname }}</td>
                <td>{{ u.email }}</td>
                <td>{{ u.role }}</td>
                <td>
                  <select @change="updateUserRole(u.id, ($event.target as HTMLSelectElement).value)" :value="u.role" class="modern-input small-select">
                    <option value="user">User</option>
                    <option value="admin">Admin</option>
                  </select>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- КОММЕНТАРИИ -->
      <div v-if="activeTab === 'comments' && !loading" class="section-card">
        <h2>Комментарии</h2>
        <div class="list">
          <div v-for="c in comments" :key="c.id" class="comment-item">
            <div class="comment-header">
              <strong>{{ c.user_name }}</strong>
              <small>{{ new Date(c.created_at).toLocaleString() }}</small>
            </div>
            <p>{{ c.body }}</p>
            <button class="btn-danger" @click="deleteComment(c.id)">Удалить</button>
          </div>
          <div v-if="comments.length === 0">Комментариев пока нет.</div>
        </div>
      </div>

      <!-- СТАТИСТИКА -->
      <div v-if="activeTab === 'stats' && !loading" class="section-card">
        <h2>Общая статистика</h2>
        <div class="stats-grid">
          <div class="stat-card">
            <span class="stat-number">{{ stats.users || 0 }}</span>
            <span class="stat-label">Пользователей</span>
          </div>
          <div class="stat-card">
            <span class="stat-number">{{ stats.courses || 0 }}</span>
            <span class="stat-label">Курсов</span>
          </div>
          <div class="stat-card">
            <span class="stat-number">{{ stats.lessons || 0 }}</span>
            <span class="stat-label">Уроков</span>
          </div>
          <div class="stat-card">
            <span class="stat-number">{{ stats.comments || 0 }}</span>
            <span class="stat-label">Комментариев</span>
          </div>
        </div>
      </div>
    </div>
  </MainLayout>
</template>

<style scoped>
.admin-page {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
  color: var(--text);
}

h1 {
  font-size: 2.2rem;
  margin-bottom: 24px;
}

.tabs {
  display: flex;
  gap: 8px;
  margin-bottom: 24px;
}

.tabs button {
  padding: 10px 20px;
  border: 1px solid var(--border);
  background: var(--surface);
  color: var(--text);
  border-radius: 8px;
  cursor: pointer;
  font-weight: 500;
  transition: all 0.2s;
  font-size: 14px;
}

.tabs button.active {
  background: var(--primary);
  color: white;
  border-color: var(--primary);
}

.section-card {
  background: var(--surface);
  border: 1px solid var(--surface-border);
  border-radius: 14px;
  padding: 24px;
  margin-bottom: 24px;
  box-shadow: 0 4px 12px rgba(0,0,0,0.03);
  backdrop-filter: blur(4px);
}

h2 {
  font-size: 1.5rem;
  margin: 0 0 20px 0;
}

.form-row {
  display: flex;
  flex-wrap: wrap;
  gap: 12px;
  align-items: center;
  margin-bottom: 20px;
}

.modern-input, select.modern-input {
  background: rgba(255,255,255,0.9);
  border: 1px solid var(--border);
  border-radius: 8px;
  padding: 10px 14px;
  color: var(--text);
  font-size: 14px;
  outline: none;
  flex: 1 1 180px;
  transition: border-color 0.2s;
}

.modern-input:focus {
  border-color: var(--primary);
  box-shadow: 0 0 0 3px var(--primary-glow);
}

textarea.modern-input {
  flex: 2 1 300px;
}

button {
  border: none;
  border-radius: 8px;
  padding: 10px 18px;
  font-weight: 600;
  cursor: pointer;
  font-size: 13px;
  transition: all 0.2s;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  white-space: nowrap;
}

.btn-primary {
  background: linear-gradient(135deg, #6366f1, #4f46e5);
  color: white;
  box-shadow: 0 4px 12px rgba(99, 102, 241, 0.3);
}
.btn-primary:hover { transform: translateY(-1px); box-shadow: 0 8px 20px rgba(99, 102, 241, 0.4); }

.btn-secondary {
  background: rgba(241, 245, 249, 0.9);
  color: var(--text);
  border: 1px solid var(--border);
}
.btn-secondary:hover { background: #e2e8f0; }

.btn-danger {
  background: #ef4444;
  color: white;
}
.btn-danger:hover { background: #dc2626; }

.small {
  padding: 4px 12px;
  font-size: 12px;
}

.list {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.list-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 16px;
  background: rgba(0,0,0,0.01);
  border-radius: 10px;
  border: 1px solid var(--border);
}

.item-info {
  flex: 1;
  font-size: 14px;
}

.item-actions {
  display: flex;
  gap: 8px;
}

.comment-item {
  padding: 16px;
  border: 1px solid var(--border);
  border-radius: 10px;
  margin-bottom: 12px;
  background: rgba(0,0,0,0.01);
}

.comment-header {
  display: flex;
  justify-content: space-between;
  margin-bottom: 8px;
}

.modal {
  background: var(--surface);
  border: 1px solid var(--surface-border);
  border-radius: 14px;
  padding: 24px;
  margin-top: 20px;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.modal-actions {
  display: flex;
  gap: 12px;
  justify-content: flex-end;
}

.table-responsive {
  overflow-x: auto;
}

.modern-table {
  width: 100%;
  border-collapse: collapse;
}

.modern-table th, .modern-table td {
  padding: 12px 16px;
  text-align: left;
  border-bottom: 1px solid var(--border);
}

.modern-table th {
  background: rgba(0,0,0,0.02);
  font-weight: 600;
}

.small-select {
  width: 100px;
  padding: 8px 12px;
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(150px, 1fr));
  gap: 20px;
  margin-top: 16px;
}

.stat-card {
  background: rgba(0,0,0,0.02);
  border-radius: 12px;
  padding: 20px;
  text-align: center;
  border: 1px solid var(--border);
}

.stat-number {
  display: block;
  font-size: 2.2rem;
  font-weight: 700;
  color: var(--primary);
}

.stat-label {
  color: var(--text-muted);
  font-size: 0.9rem;
}

.loading {
  text-align: center;
  padding: 40px;
  color: var(--text-muted);
}
</style>