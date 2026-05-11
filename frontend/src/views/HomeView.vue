<script setup lang="ts">
import MainLayout from "../components/MainLayout.vue";
import { useRouter } from "vue-router";
import { getAllCourses, type Course } from "../api/courses";
import { ref, onMounted } from "vue";

const router = useRouter();
const courses = ref<Course[]>([]);

// Карта иконок и цветов по ID курса (можно расширять)
const iconMap: Record<number, string> = {
  1: "🔄", 2: "🌐", 3: "🌳", 4: "🔍", 5: "⬆️", 6: "⛓️"
};
const colorMap: Record<number, string> = {
  1: "#7c3aed", 2: "#10b981", 3: "#f59e0b", 4: "#6366f1", 5: "#14b8a6", 6: "#8b5cf6"
};

onMounted(async () => {
  try {
    courses.value = await getAllCourses();
  } catch (e) {
    console.error("Ошибка загрузки курсов:", e);
  }
});

function getCourseIcon(courseId: number): string {
  return iconMap[courseId] || "📘";
}
function getCourseColor(courseId: number): string {
  return colorMap[courseId] || "#6366f1";
}
function openCourse(course: Course) {
  if (course.id === 1) {
    router.push("/courses/sorting");
  } else {
    router.push(`/courses/${course.id}`);
  }
}
</script>

<template>
  <MainLayout>
    <div class="hero">
      <h1 class="gradient-text">AlgoLearn</h1>
      <p class="hero-subtitle">Интерактивная визуализация алгоритмов и структур данных</p>
    </div>

    <h2 class="section-title">📚 Доступные курсы</h2>

    <div class="courses-grid">
      <div
        v-for="course in courses"
        :key="course.id"
        class="course-card"
        :style="{ '--card-color': getCourseColor(course.id) }"
      >
        <div
          class="card-icon"
          :style="{
            background: getCourseColor(course.id) + '20',
            color: getCourseColor(course.id)
          }"
        >
          {{ getCourseIcon(course.id) }}
        </div>
        <h3>{{ course.title }}</h3>
        <p>{{ course.description || '' }}</p>
        <button class="btn btn-primary mt-auto" @click="openCourse(course)">
          Начать →
        </button>
      </div>
    </div>
  </MainLayout>
</template>

<style scoped>
.hero {
  position: relative;
  text-align: center;
  padding: 60px 20px 40px;
  overflow: hidden;
}

.gradient-text {
  font-size: 3.5rem;
  font-weight: 800;
  background: linear-gradient(135deg, #a78bfa, #38bdf8, #a78bfa);
  background-size: 200% 200%;
  -webkit-background-clip: text;
  background-clip: text;
  -webkit-text-fill-color: transparent;
  animation: gradientShift 4s ease infinite;
}

@keyframes gradientShift {
  0% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
  100% { background-position: 0% 50%; }
}

.hero-subtitle {
  font-size: 1.3rem;
  color: var(--text-muted);
  margin-top: 16px;
  max-width: 500px;
  margin-left: auto;
  margin-right: auto;
}

.section-title {
  margin: 20px 0 30px;
  font-size: 2rem;
}

.courses-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 28px;
  padding-bottom: 40px;
}

.course-card {
  background: var(--surface);
  backdrop-filter: blur(14px);
  border: 1px solid var(--surface-border);
  border-radius: var(--radius);
  padding: 32px 28px;
  display: flex;
  flex-direction: column;
  transition: all var(--transition);
  position: relative;
  overflow: hidden;
}

.course-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 4px;
  background: linear-gradient(to right, var(--card-color), transparent);
}

.course-card:hover {
  transform: translateY(-10px);
  box-shadow: 0 25px 40px rgba(0,0,0,0.5), 0 0 30px rgba(99, 102, 241, 0.25);
  border-color: rgba(255,255,255,0.15);
}

.card-icon {
  width: 64px;
  height: 64px;
  border-radius: 16px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 32px;
  margin-bottom: 24px;
  box-shadow: 0 8px 20px rgba(0,0,0,0.3);
}

h3 {
  font-size: 1.5rem;
  margin-bottom: 8px;
}

p {
  color: var(--text-muted);
  margin-bottom: 24px;
  flex: 1;
}

.mt-auto {
  margin-top: auto;
  align-self: flex-start;
}
</style>