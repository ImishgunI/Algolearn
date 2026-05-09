<script setup lang="ts">
import { ref, computed, onUnmounted } from "vue";
import { executeAlgorithm, getExecution } from "../api/execution";
import CodeBlock from "../components/CodeBlock.vue";
import BubbleSortVisualizer from "../components/visualizers/BubbleSortVisualizer.vue";
import QuickSortVisualizer from "../components/visualizers/QuickSortVisualizer.vue";
import MergeSortVisualizer from "../components/visualizers/MergeSortVisualizer.vue";
import { algorithmRegistry } from "../algorithmCodes";
import type { SortStep, GraphStep, Step } from "../types/execution"

const input = ref("5, 3, 1, 8, 4, 2, 7");
const steps = ref<Step[]>([]);
const currentStep = ref(0);
const isPlaying = ref(false);
const speed = ref(600);
const isLoading = ref(false);
const selectedAlgorithm = ref("bubble_sort");

const algorithmMeta = computed(() => algorithmRegistry[selectedAlgorithm.value]);
const algorithmType = computed(() => algorithmMeta.value?.type || "other");
const currentCode = computed(() => algorithmMeta.value?.code || []);
const currentStepLines = computed(() => {
  const step = steps.value[currentStep.value];
  return step?.lines || [];
});

// ---------- Анимация ----------
let interval: number | null = null;

async function runAlgorithm() {
  isLoading.value = true;
  steps.value = [];
  currentStep.value = 0;

  try {
    const numbers = input.value.split(",").map(n => Number(n.trim()));
    console.log(selectedAlgorithm.value)
    const response = await executeAlgorithm(selectedAlgorithm.value, numbers);
    const ans = await getExecution(response.execution_id);

    const rawSteps = ans.steps || [];

    if (algorithmType.value === "sort") {
      steps.value = rawSteps.map((step: any): SortStep => ({
        type: "sort",
        bars: step.array.map((value: number, index: number) => ({
          value,
          id: index,
        })),
        active: step.active || [],
        swapping: step.swapping || [],
        lines: step.lines || [],
      }));
    } else if (algorithmType.value === "graph") {
      steps.value = rawSteps.map((step: any): GraphStep => ({
        type: "graph",
        graph: step.graph || { nodes: [], edges: [] },
        activeNodes: step.activeNodes || [],
        visitedNodes: step.visitedNodes || [],
        lines: step.lines || [],
      }));
    } else {
      steps.value = [];
    }

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

onUnmounted(() => pause());
</script>

<template>
  <div class="visualizer-page">
    <div class="header">
      <h1>⚡ Визуализатор алгоритмов</h1>
      <p class="subtitle">Введи числа через запятую и наблюдай за магией</p>
    </div>

    <!-- Панель управления -->
    <div class="controls-panel glass">
      <select v-model="selectedAlgorithm" class="algo-select">
        <option v-for="(_, key) in algorithmRegistry" :key="key" :value="key">
          {{ algorithmRegistry[key].displayName || key }}
        </option>
      </select>

      <div class="input-group" v-if="algorithmType === 'sort'">
        <div class="input-wrapper">
          <span class="input-icon">🔢</span>
          <input v-model="input" placeholder="5, 3, 1, 8, 4, 2, 7" @keyup.enter="runAlgorithm" />
        </div>
        <button class="btn btn-primary" @click="runAlgorithm" :disabled="isLoading">
          <span v-if="!isLoading">Запустить</span>
          <span v-else class="spinner"></span>
        </button>
      </div>
      <div v-else class="input-group">
        <button class="btn btn-primary" @click="runAlgorithm" :disabled="isLoading">
          Запустить алгоритм
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
    <BubbleSortVisualizer
      v-if="selectedAlgorithm === 'bubble_sort'"
      :steps="(steps as SortStep[])"
      :currentStep="currentStep"
    />

    <QuickSortVisualizer
      v-else-if="selectedAlgorithm === 'quick_sort'"
      :steps="(steps as SortStep[])"
      :currentStep="currentStep"
    />

    <MergeSortVisualizer
      v-else-if="selectedAlgorithm === 'merge_sort'"
      :steps="(steps as SortStep[])"
      :currentStep="currentStep"
    />

    <!-- <GraphVisualizer
      v-else-if="algorithmType === 'graph' && currentStepData?.type === 'graph'"
      :steps="(steps as GraphStep[])"
      :currentStep="currentStep"
    /> -->

    <!-- Блок кода -->
    <div v-if="currentCode.length" class="code-section">
      <h3>Исходный код</h3>
      <CodeBlock :codeLines="currentCode" :activeLines="currentStepLines" />
    </div>

      <!-- Управление плеером -->
    <div class="playback-controls">
      <button @click="reset" class="btn btn-secondary">↺ Сброс</button>
      <button @click="isPlaying ? pause() : play()" class="btn btn-primary play-btn">
        <span v-if="!isPlaying">▶ Анимировать</span>
        <span v-else>⏸ Пауза</span>
      </button>
      <div class="step-nav">
        <button class="btn btn-secondary" @click="currentStep = Math.max(0, currentStep - 1)">←</button>
        <button class="btn btn-secondary" @click="currentStep = Math.min(steps.length - 1, currentStep + 1)">→</button>
      </div>
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

.algo-select {
  background: rgba(255, 255, 255, 0.9);
  border: 1px solid var(--border);
  border-radius: var(--radius-sm);
  padding: 10px 16px;
  font-size: 15px;
  color: var(--text);
  cursor: pointer;
  outline: none;
}

.algo-select:focus {
  border-color: var(--primary);
  box-shadow: 0 0 0 3px var(--primary-glow);
}

.code-section {
  margin-top: 24px;
  background: var(--surface);
  padding: 16px;
  border-radius: var(--radius-sm);
}

.code-section h3 {
  margin: 0 0 12px;
  font-size: 1.1rem;
}

@keyframes spin { to { transform: rotate(360deg); } }
</style>