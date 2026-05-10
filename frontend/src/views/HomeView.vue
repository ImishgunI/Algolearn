<script setup lang="ts">
import MainLayout from "../components/MainLayout.vue";

const courses = [
  { 
    id: 1, 
    title: "Сортировки", 
    desc: "Bubble, Quick, Merge — почувствуй ритм алгоритмов",
    icon: "🔄",
    color: "#7c3aed"
  },
  { 
    id: 2, 
    title: "Графы", 
    desc: "DFS, BFS, Dijkstra — исследуй связи",
    icon: "🌐",
    color: "#10b981"
  },
  { 
    id: 3, 
    title: "Деревья", 
    desc: "BST, AVL, Red-Black — баланс и структура",
    icon: "🌳",
    color: "#f59e0b"
  },
  {
    id: 4,
    title: "Поиск",
    desc: "Линейный, бинарный — найди элемент в массиве",
    icon: "🔍",
    color: "#6366f1"
  },
  {
    id: 5,
    title: "Два указателя",
    desc: "Эффективные техники работы с массивами",
    icon: "👆",
    color: "#14b8a6"
  },
  {
    id: 6,
    title: "Связный список",
    desc: "Односвязный, двусвязный, циклы и развороты",
    icon: "⛓️",
    color: "#8b5cf6"
  },
];
</script>

<template>
  <MainLayout>
    <div class="hero">
      <h1 class="gradient-text">AlgoLearn</h1>
      <p class="hero-subtitle">Интерактивная визуализация алгоритмов и структур данных</p>
      <div class="hero-decoration">
        <div class="floating-element" v-for="i in 3" :key="i"></div>
      </div>
    </div>

    <h2 class="section-title">📚 Доступные курсы</h2>
    
    <div class="courses-grid">
      <div 
        v-for="course in courses" 
        :key="course.id" 
        class="course-card"
        :style="{ '--card-color': course.color }"
      >
        <div class="card-icon" :style="{ background: course.color + '20', color: course.color }">
          {{ course.icon }}
        </div>
        <h3>{{ course.title }}</h3>
        <p>{{ course.desc }}</p>
        <button class="btn btn-primary mt-auto">
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
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
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