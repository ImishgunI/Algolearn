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

// Для форм создания/редактирования
const newCourse = ref({ title: "", description: "" });
const editingCourse = ref<any>(null);

const newLesson = ref({
  course_id: 1, // по умолчанию
  title: "",
  theory: "",
  algorithm_type: "",
  order_index: 1,
});
const editingLesson = ref<any>(null);

// Загрузчики
const loading = ref(false);

// Загрузка данных по вкладке
async function loadCourses() {
  loading.value = true;
  try {
    courses.value = await http("/api/admin/courses");
  } finally {
    loading.value = false;
  }
}

async function loadLessons(courseId?: number) {
  if (!courseId) {
    lessons.value = [];
    return;
  }
  loading.value = true;
  try {
    lessons.value = await http(`/api/admin/courses/${courseId}/lessons`);
  } finally {
    loading.value = false;
  }
}

async function loadUsers() {
  loading.value = true;
  try {
    users.value = await http("/api/admin/users");
  } finally {
    loading.value = false;
  }
}

async function loadComments() {
  loading.value = true;
  try {
    comments.value = await http("/api/admin/comments");
  } catch (e) {
    console.error(e);
  } finally {
    loading.value = false;
  }
}

async function loadStats() {
  loading.value = true;
  try {
    stats.value = await http("/api/admin/stats");
  } finally {
    loading.value = false;
  }
}

// Переключение вкладок с загрузкой
async function switchTab(tab: Tab) {
  activeTab.value = tab;
  if (tab === "courses") await loadCourses();
  else if (tab === "users") await loadUsers();
  else if (tab === "stats") await loadStats();
  else if (tab === "comments") await loadComments();
  else if (tab === "lessons") {
    // загружаем уроки первого курса (или по выбору)
    if (courses.value.length > 0) {
      await loadLessons(courses.value[0].id);
    }
  }
}

// CRUD: курсы
async function createCourse() {
  await http("/api/admin/courses", {
    method: "POST",
    body: JSON.stringify(newCourse.value),
  });
  newCourse.value = { title: "", description: "" };
  await loadCourses();
}

async function updateCourse(course: any) {
  await http(`/api/admin/courses/${course.id}`, {
    method: "PUT",
    body: JSON.stringify(course),
  });
  editingCourse.value = null;
  await loadCourses();
}

async function deleteCourse(id: number) {
  await http(`/api/admin/courses/${id}`, { method: "DELETE" });
  await loadCourses();
}

// CRUD: уроки
async function createLesson(courseId: number) {
  await http(`/api/admin/courses/${courseId}/lessons`, {
    method: "POST",
    body: JSON.stringify(newLesson.value),
  });
  newLesson.value = { course_id: courseId, title: "", theory: "", algorithm_type: "", order_index: 1 };
  await loadLessons(courseId);
}

async function updateLesson(lesson: any) {
  await http(`/api/admin/lessons/${lesson.id}`, {
    method: "PUT",
    body: JSON.stringify(lesson),
  });
  editingLesson.value = null;
  await loadLessons(lesson.course_id);
}

async function deleteLesson(lessonId: number, courseId: number) {
  await http(`/api/admin/lessons/${lessonId}`, { method: "DELETE" });
  await loadLessons(courseId);
}

// Пользователи: смена роли
async function updateUserRole(userId: number, newRole: string) {
  await http(`/api/admin/users/${userId}/role`, {
    method: "PUT",
    body: JSON.stringify({ role: newRole }),
  });
  await loadUsers();
}

// Комментарии: удаление
async function deleteComment(commentId: number) {
  await http(`/api/admin/comments/${commentId}`, { method: "DELETE" });
  await loadComments();
}

onMounted(async () => {
  await loadCourses(); // начальная загрузка для вкладки "courses"
});
</script>

<template>
  <MainLayout>
    <div class="admin-page">
      <h1>Админ-панель</h1>

      <div class="tabs">
        <button
          v-for="tab in (['courses','lessons','users','comments','stats'] as Tab[])"
          :key="tab"
          :class="{ active: activeTab === tab }"
          @click="switchTab(tab)"
        >
          {{ tab === 'courses' ? 'Курсы' : tab === 'lessons' ? 'Уроки' : tab === 'users' ? 'Пользователи' : tab === 'comments' ? 'Комментарии' : 'Статистика' }}
        </button>
      </div>

      <div v-if="loading" class="loading">Загрузка...</div>

      <!-- Курсы -->
      <div v-if="activeTab === 'courses' && !loading">
        <h2>Курсы</h2>
        <div class="form">
          <input v-model="newCourse.title" placeholder="Название курса" />
          <input v-model="newCourse.description" placeholder="Описание" />
          <button class="btn" @click="createCourse">Создать курс</button>
        </div>
        <ul>
          <li v-for="course in courses" :key="course.id">
            <strong>{{ course.title }}</strong> — {{ course.description }}
            <button class="btn small" @click="editingCourse = course">Редактировать</button>
            <button class="btn small btn-danger" @click="deleteCourse(course.id)">Удалить</button>
          </li>
        </ul>
        <!-- модалка редактирования -->
        <div v-if="editingCourse" class="modal">
          <h3>Редактировать курс</h3>
          <input v-model="editingCourse.title" />
          <input v-model="editingCourse.description" />
          <button class="btn" @click="updateCourse(editingCourse)">Сохранить</button>
          <button class="btn" @click="editingCourse = null">Отмена</button>
        </div>
      </div>

      <!-- Уроки -->
      <div v-if="activeTab === 'lessons' && !loading">
        <h2>Уроки</h2>
        <div class="form">
          <select v-model="newLesson.course_id">
            <option v-for="c in courses" :key="c.id" :value="c.id">{{ c.title }}</option>
          </select>
          <input v-model="newLesson.title" placeholder="Название урока" />
          <textarea v-model="newLesson.theory" placeholder="Теория"></textarea>
          <input v-model="newLesson.algorithm_type" placeholder="Тип алгоритма (bubble_sort)" />
          <input v-model.number="newLesson.order_index" type="number" placeholder="Порядок" />
          <button class="btn" @click="createLesson(newLesson.course_id)">Добавить урок</button>
        </div>

        <div v-for="course in courses" :key="course.id">
          <h3>{{ course.title }}</h3>
          <button @click="loadLessons(course.id)">Загрузить уроки</button>
          <ul v-if="lessons.length">
            <li v-for="lesson in lessons" :key="lesson.id">
              <strong>{{ lesson.title }}</strong> ({{ lesson.algorithm_type }})
              <button @click="editingLesson = lesson">Редактировать</button>
              <button class="btn-danger" @click="deleteLesson(lesson.id, lesson.course_id)">Удалить</button>
            </li>
          </ul>
        </div>
        <!-- модалка редактирования урока -->
        <div v-if="editingLesson" class="modal">
          <h3>Редактировать урок</h3>
          <input v-model="editingLesson.title" />
          <textarea v-model="editingLesson.theory"></textarea>
          <input v-model="editingLesson.algorithm_type" />
          <input v-model.number="editingLesson.order_index" type="number" />
          <button @click="updateLesson(editingLesson)">Сохранить</button>
          <button @click="editingLesson = null">Отмена</button>
        </div>
      </div>

      <!-- Пользователи -->
      <div v-if="activeTab === 'users' && !loading">
        <h2>Пользователи</h2>
        <table>
          <thead>
            <tr>
              <th>ID</th><th>Имя</th><th>Email</th><th>Роль</th><th>Действия</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="u in users" :key="u.id">
              <td>{{ u.id }}</td>
              <td>{{ u.user_name }} {{ u.user_surname }}</td>
              <td>{{ u.email }}</td>
              <td>{{ u.role }}</td>
              <td>
                <select @change="updateUserRole(u.id, ($event.target as HTMLSelectElement).value)" :value="u.role">
                  <option value="user">User</option>
                  <option value="admin">Admin</option>
                </select>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Комментарии -->
      <div v-if="activeTab === 'comments' && !loading">
        <h2>Комментарии</h2>
        <ul>
          <li v-for="c in comments" :key="c.id">
            <strong>{{ c.user_name }}</strong> ({{ new Date(c.created_at).toLocaleString() }})
            <p>{{ c.body }}</p>
            <button @click="deleteComment(c.id)">Удалить</button>
          </li>
        </ul>
      </div>

      <!-- Статистика -->
      <div v-if="activeTab === 'stats' && !loading">
        <h2>Общая статистика</h2>
        <ul>
          <li>Пользователи: {{ stats.users }}</li>
          <li>Курсы: {{ stats.courses }}</li>
          <li>Уроки: {{ stats.lessons }}</li>
          <li>Комментарии: {{ stats.comments }}</li>
        </ul>
      </div>
    </div>
  </MainLayout>
</template>

<style scoped>
.admin-page {
  max-width: 1100px;
  margin: 0 auto;
  padding: 20px;
}
.tabs { display: flex; gap: 8px; margin-bottom: 20px; }
.tabs button {
  padding: 8px 16px;
  border: 1px solid var(--border);
  background: var(--surface);
  color: var(--text);
  border-radius: 6px;
  cursor: pointer;
}
.tabs button.active { background: var(--primary); color: white; }
.form { margin-bottom: 20px; display: flex; gap: 8px; flex-wrap: wrap; align-items: center; }
input, select, textarea { padding: 8px; border: 1px solid var(--border); border-radius: 4px; color: var(--text); background: var(--surface); }
.btn { background: var(--primary); color: white; border: none; padding: 8px 16px; border-radius: 4px; cursor: pointer; }
.btn-danger { background: #ef4444; color: white; }
.btn-danger:hover { background: #dc2626; }
.small { font-size: 0.85rem; padding: 4px 8px; }
.modal { background: var(--surface); border: 1px solid var(--border); padding: 20px; margin: 10px 0; }
.loading { text-align: center; padding: 40px; }
</style>