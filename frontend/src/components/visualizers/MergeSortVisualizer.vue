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

  if (current.value.mergedRange && 
      index >= current.value.mergedRange[0] && 
      index <= current.value.mergedRange[1]) {
    return "merged";
  }

  if (current.value.leftRange && 
      index >= current.value.leftRange[0] && 
      index <= current.value.leftRange[1]) {
    return "left";
  }

  if (current.value.rightRange && 
      index >= current.value.rightRange[0] && 
      index <= current.value.rightRange[1]) {
    return "right";
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
      <span class="legend-item left">Левая половина</span>
      <span class="legend-item right">Правая половина</span>
      <span class="legend-item merged">Слияние</span>
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
  transition: height 0.4s ease, background 0.3s ease;
  display: flex;
  align-items: flex-end;
  justify-content: center;
  color: white;
  font-weight: 600;
}

.bar.left { background: #22c55e; }
.bar.right { background: #eab308; }
.bar.merged { 
  background: #a855f7; 
  transform: scale(1.05);
}

.bar-value {
  margin-bottom: 6px;
  font-size: 13px;
}

.legend {
  display: flex;
  justify-content: center;
  gap: 30px;
  margin-top: 20px;
  font-size: 0.95rem;
}

.legend-item {
  display: flex;
  align-items: center;
  gap: 8px;
}

.legend-item::before {
  content: '';
  display: inline-block;
  width: 14px;
  height: 14px;
  border-radius: 4px;
}

.legend-item.left::before { background: #22c55e; }
.legend-item.right::before { background: #eab308; }
.legend-item.merged::before { background: #a855f7; }
</style>