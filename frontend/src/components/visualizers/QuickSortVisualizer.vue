<script setup lang="ts">
import { computed } from "vue";
import type { SortStep } from "../../types/execution";

const props = defineProps<{
  steps: SortStep[];
  currentStep: number;
}>();

const current = computed(() => props.steps[props.currentStep]);

function getBarClass(index: number): string {
  if (!current.value) return "";

  if (current.value.swapping?.includes(index)) return "swapping";
  if (current.value.active?.includes(index)) return "active";
  if (current.value.pivot?.includes(index)) return "pivot";
  if (current.value.partition && 
      index >= current.value.partition[0] && 
      index <= current.value.partition[1]) {
    return "partition";
  }

  return "";
}
</script>

<template>
  <div class="visualization">
    <div class="bars-container">
      <div
        v-for="(bar, index) in current?.bars || []"
        :key="bar.id"
        class="bar-wrapper"
      >
        <div
          class="bar"
          :class="getBarClass(index)"
          :style="{ height: Math.min(bar.value * 18, 380) + 'px' }"
        >
          <span class="bar-value">{{ bar.value }}</span>
        </div>
        <div class="bar-index">{{ index }}</div>
      </div>
    </div>

    <div class="legend">
      <span class="legend-item pivot">● Pivot</span>
      <span class="legend-item partition">● Partition</span>
    </div>
  </div>
</template>

<style scoped>
.bars-container {
  display: flex;
  align-items: flex-end;
  justify-content: center;
  gap: 8px;
  height: 420px;
  background: rgba(15, 23, 42, 0.6);
  border-radius: 16px;
  padding: 30px 20px;
}

.bar-wrapper {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 40px;
}

.bar {
  width: 100%;
  background: #6366f1;
  border-radius: 8px 8px 0 0;
  transition: all 0.3s ease;
  display: flex;
  align-items: flex-end;
  justify-content: center;
  color: white;
  font-weight: 600;
}

.bar.active { background: #f59e0b; }
.bar.swapping { 
  background: #ef4444; 
  transform: scale(1.1);
}
.bar.pivot { 
  background: #8b5cf6; 
  transform: scale(1.08);
}
.bar.partition {
  background: #22c55e;
  opacity: 0.85;
}

.legend {
  display: flex;
  justify-content: center;
  gap: 30px;
  margin-top: 20px;
}
</style>