<script setup lang="ts">
import { ref } from "vue";
// import { executeAlgorithm, getExecution } from "../api/execution";

type Bar = {
  value: number;
  id: number;
};

type StepUI = {
  bars: Bar[];
  active: number[];
  swapping: number[];
};

const input = ref("5,3,1");
const steps = ref<StepUI[]>([]);
const currentStep = ref(0);
const playing = ref(false);
const speed = ref(500);
const phase = ref<"idle" | "swap">("idle");

let intervalId: any = null;

async function run() {
  const data = input.value.split(",").map(Number);

  const initialBars: Bar[] = data.map((value, index) => ({
    value,
    id: index,
  }));

  let currentBars = [...initialBars];
  const resultSteps: StepUI[] = [];

  const n = currentBars.length;

  for (let i = 0; i < n; i++) {
    for (let j = 0; j < n - i - 1; j++) {

      const active = [j, j + 1];

      if (currentBars[j].value > currentBars[j + 1].value) {
        const newBars = [...currentBars];

        [newBars[j], newBars[j + 1]] = [newBars[j + 1], newBars[j]];

        currentBars = newBars;

        resultSteps.push({
          bars: [...currentBars],
          active,
          swapping: [j, j + 1],
        });
      } else {
        resultSteps.push({
          bars: [...currentBars],
          active,
          swapping: [],
        });
      }
    }
  }

  steps.value = resultSteps;
  currentStep.value = 0;
}

function play() {
  if (playing.value) return;

  playing.value = true;

  intervalId = setInterval(() => {
    if (currentStep.value >= steps.value.length - 1) {
      pause();
      return;
    }

    const step = steps.value[currentStep.value];

    if (step.swapping.length) {
      phase.value = "swap";

      setTimeout(() => {
        phase.value = "idle";
      }, speed.value * 0.6);
    }

    currentStep.value++;
  }, speed.value);
}

function pause() {
  clearInterval(intervalId);
  playing.value = false;
}

function reset() {
  pause();
  currentStep.value = 0;
}

function getClass(index: number) {
  const step = steps.value[currentStep.value];
  if (!step) return "";

  if (step.swapping.includes(index)) return "swapping";
  if (step.active.includes(index)) return "active";

  return "";
}
</script>

<template>
  <div class="container">
    <h1>Algorithm Visualizer</h1>

    <!-- input -->
    <div class="controls">
      <input v-model="input" placeholder="5,3,1" />

      <button @click="run">Run</button>
      <button @click="play">Play</button>
      <button @click="pause">Pause</button>
      <button @click="reset">Reset</button>
    </div>

    <!-- speed -->
    <div class="controls">
      <label>Speed:</label>
      <input type="range" min="50" max="1000" v-model="speed" />
    </div>

    <!-- step info -->
    <p v-if="steps.length">
      Step: {{ currentStep }} / {{ steps.length - 1 }}
    </p>

    <!-- visualization -->
    <div class="bars">
      <div
        v-for="(bar, index) in steps[currentStep]?.bars || []"
        :key="bar.id"
        class="bar"
        :class="getClass(index)"
        :style="{ height: bar.value * 20 + 'px' }"
      />
    </div>
  </div>
</template>

<style scoped>
.container {
  padding: 20px;
  font-family: sans-serif;
}

.controls {
  margin-bottom: 10px;
  display: flex;
  gap: 10px;
}

.bars {
  position: relative;
  height: 250px;
  display: flex;
  align-items: flex-end;
  gap: 10px;
}

.bar {
  width: 30px;
  background: steelblue;
  transition: transform 0.4s ease, height 0.3s ease;
}

.bar.active {
  background: orange;
}

.bar.swapping {
  background: red;
  box-shadow: 0 0 10px rgba(255,0,0,0.6);
}

.bars-container {
  position: relative;
}

.label {
  position: absolute;
  top: -20px;
  font-size: 12px;
}
</style>