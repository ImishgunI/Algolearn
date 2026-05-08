<script setup lang="ts">
import { ref, onUnmounted } from "vue";
import { executeAlgorithm, getExecution } from "../api/execution";

type Bar = {
  value: number;
  id: number;
};

type Step = {
  bars: Bar[];
  active: number[];
  swapping: number[];
  comparing?: number[];
};

const input = ref("5, 3, 1, 8, 4, 2, 7");
const steps = ref<Step[]>([]);
const currentStep = ref(0);
const isPlaying = ref(false);
const speed = ref(600);
const isLoading = ref(false);
const algorithmName = ref("Bubble Sort");

let interval: number | null = null;

async function runAlgorithm() {
  isLoading.value = true;
  steps.value = [];
  currentStep.value = 0;

  try {
    const numbers = input.value.split(",").map(n => Number(n.trim()));
    const response = await executeAlgorithm(numbers);
    const ans = await getExecution(response.execution_id)
    steps.value = ans.steps || [];
    currentStep.value = 0;
  } catch (err) {
    alert("Ошибка при запуске алгоритма");
    console.error(err);
  } finally {
    isLoading.value = false;
  }
}

function play() {
  if (isPlaying.value || steps.value.length === 0) return;
  isPlaying.value = true;
  interval = setInterval(() => {
    if (currentStep.value >= steps.value.length - 1) {
      pause();
      return;
    }
    currentStep.value++;
  }, speed.value);
}

function pause() {
  if (interval) clearInterval(interval);
  isPlaying.value = false;
}

function reset() {
  pause();
  currentStep.value = 0;
}

function getBarClass(index: number) {
  const step = steps.value[currentStep.value];
  if (!step) return "";
  if (step.swapping.includes(index)) return "swapping";
  if (step.active.includes(index)) return "active";
  if (step.comparing?.includes(index)) return "comparing";
  return "";
}

onUnmounted(() => pause());
</script>

<template>
  <div class="visualizer-page">
    <div class="header">
      <h1>⚡ Визуализатор алгоритмов</h1>
      <p class="subtitle">Введи числа через запятую и наблюдай за магией сортировки</p>
    </div>

    <!-- Панель управления -->
    <div class="controls-panel glass">
      <div class="input-group">
        <div class="input-wrapper">
          <span class="input-icon">🔢</span>
          <input 
            v-model="input" 
            placeholder="5, 3, 1, 8, 4, 2, 7"
            @keyup.enter="runAlgorithm"
          />
        </div>
        <button class="btn btn-primary" @click="runAlgorithm" :disabled="isLoading">
          <span v-if="!isLoading">Запустить</span>
          <span v-else class="spinner"></span>
        </button>
      </div>

      <div class="speed-control">
        <div class="speed-header">
          <span>🐢 Скорость анимации</span>
          <span class="speed-value">{{ speed }} мс</span>
        </div>
        <input type="range" v-model="speed" min="100" max="1200" step="50" class="range-slider" />
      </div>
    </div>

    <!-- Визуализация -->
    <div v-if="steps.length" class="visualization glass">
      <div class="info-bar">
        <div class="step-counter">
          <span class="badge">Шаг {{ currentStep }} / {{ steps.length - 1 }}</span>
          <span class="algo-badge">{{ algorithmName }}</span>
        </div>
      </div>

      <div class="bars-container">
        <div
          v-for="(bar, index) in steps[currentStep]?.bars || []"
          :key="bar.id"
          class="bar-wrapper"
        >
          <div
            class="bar"
            :class="getBarClass(index)"
            :style="{ height: bar.value * 20 + 'px' }"
          >
            <span class="bar-value">{{ bar.value }}</span>
          </div>
          <div class="bar-index">{{ index }}</div>
        </div>
      </div>

      <div class="playback-controls">
        <button @click="reset" class="btn btn-secondary">↺ Сброс</button>
        <button 
          @click="isPlaying ? pause() : play()" 
          class="btn btn-primary play-btn"
        >
          <span v-if="!isPlaying">▶ Анимировать</span>
          <span v-else>⏸ Пауза</span>
        </button>
        <div class="step-nav">
          <button class="btn btn-secondary" @click="currentStep = Math.max(0, currentStep-1)">←</button>
          <button class="btn btn-secondary" @click="currentStep = Math.min(steps.length-1, currentStep+1)">→</button>
        </div>
      </div>
    </div>

    <!-- Плейсхолдер -->
    <div v-else class="placeholder glass">
      <div class="placeholder-icon">📊</div>
      <p>Введи числа и нажми «Запустить», чтобы увидеть анимацию сортировки</p>
    </div>
  </div>
</template>

<style scoped>
.visualizer-page {
  max-width: 1000px;
  margin: 0 auto;
  padding: 20px;
}

.header {
  text-align: center;
  margin-bottom: 30px;
}

.subtitle {
  color: var(--text-muted);
  font-size: 1.1rem;
}

.glass {
  background: var(--surface);
  backdrop-filter: blur(12px);
  border: 1px solid var(--surface-border);
  border-radius: var(--radius);
}

.controls-panel {
  padding: 24px;
  margin-bottom: 30px;
  display: flex;
  flex-wrap: wrap;
  gap: 24px;
  align-items: flex-end;
}

.input-group {
  display: flex;
  gap: 12px;
  flex: 2;
  min-width: 280px;
}

.input-wrapper {
  position: relative;
  flex: 1;
}

.input-icon {
  position: absolute;
  left: 14px;
  top: 50%;
  transform: translateY(-50%);
  opacity: 0.7;
}

.input-wrapper input {
  padding-left: 42px;
  width: 100%;
}

.speed-control {
  flex: 1;
  min-width: 200px;
}

.speed-header {
  display: flex;
  justify-content: space-between;
  margin-bottom: 8px;
  font-size: 0.9rem;
}

.speed-value {
  color: var(--accent);
  font-weight: 600;
}

.range-slider {
  width: 100%;
  accent-color: var(--primary);
}

.visualization {
  padding: 24px;
  margin-bottom: 30px;
}

.info-bar {
  margin-bottom: 20px;
}

.step-counter {
  display: flex;
  gap: 12px;
  align-items: center;
}

.badge {
  background: rgba(99, 102, 241, 0.15);
  padding: 6px 16px;
  border-radius: 30px;
  font-weight: 600;
  color: #c4b5fd;
}

.algo-badge {
  background: rgba(16, 185, 129, 0.15);
  padding: 6px 16px;
  border-radius: 30px;
  color: #6ee7b7;
  font-weight: 600;
}

.bars-container {
  display: flex;
  align-items: flex-end;
  justify-content: center;
  gap: 10px;
  height: 400px;
  padding: 20px 10px 10px;
  background: rgba(0, 0, 0, 0.2);
  border-radius: var(--radius-sm);
  position: relative;
}

.bar-wrapper {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 44px;
}

.bar {
  width: 100%;
  background: linear-gradient(180deg, #818cf8, #6366f1);
  border-radius: 8px 8px 0 0;
  box-shadow: 0 -4px 15px rgba(99, 102, 241, 0.4);
  transition: height 0.35s cubic-bezier(0.4, 0, 0.2, 1), background 0.2s, box-shadow 0.3s;
  display: flex;
  align-items: flex-start;
  justify-content: center;
  padding-top: 8px;
  font-weight: 700;
  font-size: 14px;
  color: white;
  text-shadow: 0 2px 4px rgba(0,0,0,0.4);
}

.bar.active {
  background: linear-gradient(180deg, #fbbf24, #d97706);
  box-shadow: 0 -4px 20px rgba(245, 158, 11, 0.6);
}

.bar.swapping {
  background: linear-gradient(180deg, #f87171, #dc2626);
  box-shadow: 0 -4px 20px rgba(239, 68, 68, 0.6);
  transform: scale(1.05);
}

.bar.comparing {
  background: linear-gradient(180deg, #34d399, #059669);
  box-shadow: 0 -4px 20px rgba(16, 185, 129, 0.6);
}

.bar-index {
  margin-top: 6px;
  font-size: 0.8rem;
  color: var(--text-muted);
}

.playback-controls {
  display: flex;
  gap: 16px;
  justify-content: center;
  margin-top: 28px;
  align-items: center;
}

.play-btn {
  min-width: 170px;
}

.step-nav {
  display: flex;
  gap: 8px;
}

.placeholder {
  padding: 60px 20px;
  text-align: center;
}

.placeholder-icon {
  font-size: 3rem;
  margin-bottom: 20px;
  opacity: 0.8;
}

.placeholder p {
  color: var(--text-muted);
  font-size: 1.1rem;
}

.spinner {
  width: 18px;
  height: 18px;
  border: 2px solid rgba(255,255,255,0.3);
  border-top-color: white;
  border-radius: 50%;
  animation: spin 0.7s linear infinite;
  display: inline-block;
}

@keyframes spin { to { transform: rotate(360deg); } }
</style>