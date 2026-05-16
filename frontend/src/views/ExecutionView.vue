<script setup lang="ts">
import { ref, computed, onUnmounted, onMounted } from "vue";
import { useRoute } from "vue-router";
import { executeAlgorithm, getExecution } from "../api/execution";
import { http } from "../api/http";
import CodeBlock from "../components/CodeBlock.vue";
import BubbleSortVisualizer from "../components/visualizers/BubbleSortVisualizer.vue";
import QuickSortVisualizer from "../components/visualizers/QuickSortVisualizer.vue";
import MergeSortVisualizer from "../components/visualizers/MergeSortVisualizer.vue";
import { algorithmRegistry } from "../algorithmCodes";
import type { SortStep, GraphStep, Step } from "../types/execution"
import CodeEditor from "../components/CodeEditor.vue";

const input = ref("5, 3, 1, 8, 4, 2, 7");
const steps = ref<Step[]>([]);
const currentStep = ref(0);
const isPlaying = ref(false);
const speed = ref(600);
const isLoading = ref(false);
const selectedAlgorithm = ref("bubble_sort");

const customCode = ref(`func Run(input []int) []Step {
  // Пример: пузырьковая сортировка
  var steps []Step
  n := len(input)
  arr := make([]int, n)
  copy(arr, input)
  for i := 0; i < n-1; i++ {
    for j := 0; j < n-i-1; j++ {
      steps = append(steps, Step{Array: append([]int(nil), arr...), Active: []int{j, j+1}})
      if arr[j] > arr[j+1] {
        arr[j], arr[j+1] = arr[j+1], arr[j]
        steps = append(steps, Step{Array: append([]int(nil), arr...), Swapping: []int{j, j+1}})
      }
    }
  }
  return steps
}`);

const algorithmMeta = computed(() => {
  if (selectedAlgorithm.value === "custom") {
    return { type: "sort", displayName: "Свой код (Go)", code: [] };
  }
  return algorithmRegistry[selectedAlgorithm.value] || { type: "sort", displayName: selectedAlgorithm.value, code: [] };
});

const algorithmType = computed(() => algorithmMeta.value?.type || "other");
const currentCode = computed(() => algorithmMeta.value?.code || []);
const currentStepLines = computed(() => {
  const step = steps.value[currentStep.value];
  return step?.lines || [];
});

// ---------- Анимация ----------
let interval: number | null = null;

async function runPredefinedAlgorithm() {
  const numbers = input.value.split(",").map(n => Number(n.trim()));
  const response = await executeAlgorithm(selectedAlgorithm.value, numbers);
  const ans = await getExecution(response.execution_id);
  const rawSteps = ans.steps || [];
  if (algorithmType.value === "sort") {
    steps.value = rawSteps.map((step: any): SortStep => ({
      type: "sort",
      bars: step.array.map((value: number, index: number) => ({ value, id: index })),
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
}

async function runCustomAlgorithm() {
  const numbers = input.value.split(",").map(n => Number(n.trim()));
  const response = await http("/api/execute-custom", {
    method: "POST",
    body: JSON.stringify({ code: customCode.value, input: numbers }),
  });
  const rawSteps = response.steps || [];
  steps.value = rawSteps.map((step: any): SortStep => ({
    type: "sort",
    bars: step.array.map((value: number, idx: number) => ({ value, id: idx })),
    active: step.active || [],
    swapping: step.swapping || [],
    lines: step.lines || [],
  }));
}

async function runAlgorithm() {
  isLoading.value = true;
  steps.value = [];
  currentStep.value = 0;
  try {
    if (selectedAlgorithm.value === "custom") {
      await runCustomAlgorithm();
    } else {
      await runPredefinedAlgorithm();
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

// Приём параметров из URL
const route = useRoute();
onMounted(() => {
  const algo = route.query.algorithm as string;
  const inp = route.query.input as string;
  if (algo) selectedAlgorithm.value = algo;
  if (inp) input.value = inp;
});
</script>

<template>
  <div class="visualizer-page">
    <div class="header">
      <h1>⚡ Визуализатор алгоритмов</h1>
      <p class="subtitle">Выберите алгоритм или напишите свой код на Go</p>
    </div>

    <!-- Панель управления -->
    <div class="controls-panel glass">
      <select v-model="selectedAlgorithm" class="algo-select">
        <option v-for="(meta, key) in algorithmRegistry" :key="key" :value="key">
          {{ meta.displayName || key }}
        </option>
        <option value="custom">Свой код (Go)</option>
      </select>

      <div class="input-group" v-if="algorithmType === 'sort' || selectedAlgorithm === 'custom'">
        <div class="input-wrapper">
          <span class="input-icon">🔢</span>
          <input v-model="input" placeholder="5, 3, 1, 8, 4, 2, 7" @keyup.enter="runAlgorithm" />
        </div>
        <button class="btn btn-primary" @click="runAlgorithm" :disabled="isLoading">
          <span v-if="!isLoading">Начать</span>
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

    <!-- Редактор для своего кода -->
    <div v-if="selectedAlgorithm === 'custom'" class="custom-code-section glass">
      <h3>Ваш алгоритм на Go</h3>
      <p class="hint">
        Определите функцию <code>Run(input []int) []Step</code>, где Step — структура с полями Array, Active, Swapping.
      </p>
      <CodeEditor v-model="customCode" />
    </div>

    <!-- Визуализация -->
    <BubbleSortVisualizer
      v-if="selectedAlgorithm === 'bubble_sort' || selectedAlgorithm === 'custom'"
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

    <!-- <GraphVisualizer ... /> -->

    <!-- Блок кода для встроенных алгоритмов -->
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

.custom-code-section {
  padding: 20px;
  margin-bottom: 24px;
}
.hint {
  font-size: 0.9rem;
  color: var(--text-muted);
  margin-bottom: 12px;
}
.hint code {
  background: var(--surface2);
  padding: 2px 6px;
  border-radius: 4px;
}

@keyframes spin { to { transform: rotate(360deg); } }
</style>