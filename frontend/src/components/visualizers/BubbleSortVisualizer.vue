<script setup lang="ts">
import { computed } from "vue";
import type { SortStep } from "../../types/execution";

const props = defineProps<{
  steps: SortStep[];
  currentStep: number;
}>();

const currentStepData = computed(() => {
  return props.steps[props.currentStep] || {
    bars: [],
    active: [],
    swapping: [],
    lines: [],
  };
});

function getBarClass(index: number) {
  if (currentStepData.value.swapping.includes(index)) return "swapping";
  if (currentStepData.value.active.includes(index)) return "active";
  return "";
}
</script>

<template>
  <div class="bars-container">
    <div
      v-for="(bar, index) in currentStepData.bars"
      :key="bar.id"
      class="bar-wrapper"
    >
      <div
        class="bar"
        :class="getBarClass(index)"
        :style="{ height: Math.min(bar.value * 20, 360) + 'px' }"
      >
        <span class="bar-value">{{ bar.value }}</span>
      </div>
      <div class="bar-index">{{ index }}</div>
    </div>
  </div>
</template>

<style scoped>
.bars-container {
  display: flex;
  align-items: flex-end;
  justify-content: center;
  gap: 10px;
  height: 400px;
  padding: 20px 10px 10px;
  background: rgba(0, 0, 0, 0.2);
  border-radius: var(--radius-sm);
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
  transition: 0.3s;
  display: flex;
  justify-content: center;
  color: white;
}

.bar.active {
  background: orange;
}

.bar.swapping {
  background: red;
}
</style>