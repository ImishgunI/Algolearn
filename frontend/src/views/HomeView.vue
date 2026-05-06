<script setup lang="ts">
import { useRouter } from "vue-router";
import { user, isAdmin } from "../store/user";
import { logout } from "../api/auth";
const router = useRouter();

function goToVisualizer() {
  router.push("/visualize");
}
</script>

<template>
  <div class="home">
    <div v-if="user">
      <p>Welcome, {{ user.user_id }}</p>
      <button v-if="isAdmin()">Create Course (Admin)</button>
      <button @click="logout">Logout</button>
    </div>

    <div class="hero">
      <h1>Algorithm Visualizer</h1>

      <p>
        Учебная система для визуализации алгоритмов сортировки и структур данных.
      </p>

      <div class="buttons">
        <button @click="goToVisualizer">Start Visualizing</button>
        <button disabled>Login (soon)</button>
      </div>
    </div>

    <div class="features">
      <div class="card">
        <h3>Sorting Algorithms</h3>
        <p>Bubble, Quick, Merge и другие</p>
      </div>

      <div class="card">
        <h3>Step-by-step</h3>
        <p>Пошаговое выполнение алгоритмов</p>
      </div>

      <div class="card">
        <h3>Interactive UI</h3>
        <p>Play, pause, speed control</p>
      </div>
    </div>
  </div>
</template>

<style scoped>
.home {
  min-height: 100vh;
  padding: 60px;
  font-family: sans-serif;
  background: #0f172a;
  color: white;
}

.hero {
  text-align: center;
  margin-bottom: 60px;
}

.hero h1 {
  font-size: 48px;
  margin-bottom: 10px;
}

.hero p {
  opacity: 0.7;
  margin-bottom: 20px;
}

.buttons {
  display: flex;
  gap: 10px;
  justify-content: center;
}

button {
  padding: 10px 20px;
  border: none;
  cursor: pointer;
  border-radius: 8px;
  background: steelblue;
  color: white;
}

button:hover {
  background: #1d4ed8;
}

button:disabled {
  opacity: 0.4;
  cursor: not-allowed;
}

.features {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 20px;
}

.card {
  background: #1e293b;
  padding: 20px;
  border-radius: 12px;
}
</style>