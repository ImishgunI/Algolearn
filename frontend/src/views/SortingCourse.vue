<script setup lang="ts">
import MainLayout from "../components/MainLayout.vue";
import { sortingLessons } from "../data/sortingLessons";
import { useRouter, useRoute } from "vue-router";
import { computed } from "vue";

const router = useRouter();
const route = useRoute();

const currentId = computed(() => route.params.lessonId as string);

function goToLesson(id: string) {
  router.push(`/courses/sorting/${id}`);
}
</script>

<template>
  <MainLayout>
    <div class="course-layout">
      <aside class="sidebar">
        <h2>Сортировки</h2>
        <nav>
          <ul>
            <li v-for="lesson in sortingLessons" :key="lesson.id">
              <a
                :class="{ active: lesson.id === currentId }"
                @click="goToLesson(lesson.id)"
              >
                {{ lesson.title }}
              </a>
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
.course-layout {
  display: flex;
  gap: 32px;
  margin-top: 20px;
}
.sidebar {
  width: 240px;
  flex-shrink: 0;
  background: var(--surface);
  border-radius: var(--radius);
  padding: 20px;
  border: 1px solid var(--surface-border);
}
.sidebar h2 {
  font-size: 1.3rem;
  margin-bottom: 16px;
}
.sidebar ul {
  list-style: none;
  padding: 0;
}
.sidebar li {
  margin-bottom: 8px;
}
.sidebar a {
  display: block;
  padding: 8px 12px;
  border-radius: 6px;
  color: var(--text);
  text-decoration: none;
  cursor: pointer;
  transition: background 0.2s;
}
.sidebar a:hover {
  background: rgba(0,0,0,0.04);
}
.sidebar a.active {
  background: var(--primary);
  color: white;
  font-weight: 600;
}
.lesson-content {
  flex: 1;
  background: var(--surface);
  border-radius: var(--radius);
  padding: 32px;
  border: 1px solid var(--surface-border);
}
</style>