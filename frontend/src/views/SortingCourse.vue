<script setup lang="ts">
import MainLayout from "../components/MainLayout.vue";
import { ref, onMounted, computed } from "vue";
import { useRoute, useRouter } from "vue-router";
import { getLessonsByCourse, type Lesson } from "../api/lessons";

const route = useRoute();
const router = useRouter();
const currentId = computed(() => {
  const id = Number(route.params.lessonId);
  return isNaN(id) ? null : id;
});

const lessons = ref<Lesson[]>([]);
const courseID = 1; // ID курса "Сортировки" в таблице courses

onMounted(async () => {
  try {
    lessons.value = await getLessonsByCourse(courseID);
    // Если URL без ID и уроки есть — перейти на первый урок
    if (!currentId.value && lessons.value.length > 0) {
      router.replace(`/courses/sorting/${lessons.value[0].id}`);
    }
  } catch (e) {
    console.error(e);
  }
});
</script>

<template>
  <MainLayout>
    <div class="course-layout">
      <aside class="sidebar">
        <h2>Сортировки</h2>
        <nav>
          <ul>
            <li v-for="lesson in lessons" :key="lesson.id">
              <router-link
                :to="`/courses/sorting/${lesson.id}`"
                :class="{ active: lesson.id === currentId }"
              >
                {{ lesson.title }}
              </router-link>
            </li>
          </ul>
        </nav>
      </aside>
      <main class="lesson-content">
        <router-view />
      </main>
    </div>
  </MainLayout>
</template>

<style scoped>
.course-layout { display: flex; gap: 32px; margin-top: 20px; }
.sidebar { width: 240px; flex-shrink: 0; background: var(--surface); border-radius: var(--radius); padding: 20px; border: 1px solid var(--surface-border); }
.sidebar h2 { font-size: 1.3rem; margin-bottom: 16px; }
.sidebar ul { list-style: none; padding: 0; }
.sidebar li { margin-bottom: 8px; }
.sidebar a { display: block; padding: 8px 12px; border-radius: 6px; color: var(--text); text-decoration: none; transition: background 0.2s; }
.sidebar a:hover { background: rgba(0,0,0,0.04); }
.sidebar a.active { background: var(--primary); color: white; font-weight: 600; }
.lesson-content { flex: 1; background: var(--surface); border-radius: var(--radius); padding: 32px; border: 1px solid var(--surface-border); }
</style>