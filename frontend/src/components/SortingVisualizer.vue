<script setup lang="ts">
import { computed } from "vue";

type Bar = { value: number; id: number };

type Step = {
  bars: Bar[];
  active: number[];
  swapping: number[];
};

const props = defineProps<{
  steps: Step[];
  currentStep: number;
  maxHeight?: number;
}>();

const current = computed(() => props.steps[props.currentStep] || { bars: [], active: [], swapping: [] });

function getBarClass(index: number) {
  if (current.value.swapping.includes(index)) return "swapping";
  if (current.value.active.includes(index)) return "active";
  return "";
}
</script>

<template>
  <div class="bars-container">
    <div
      v-for="(bar, index) in current.bars"
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
  min-height: 50px;
}

.bar-index {
  margin-top: 6px;
  font-size: 0.8rem;
  color: var(--text-muted);
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
</style>