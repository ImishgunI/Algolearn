<script setup lang="ts">
import { computed } from "vue";
import type { SortStep } from "../../types/execution";

const props = defineProps<{
  steps: SortStep[];
  currentStep: number;
}>();

const current = computed(() => props.steps[props.currentStep]);

function getBarClass(index: number) {
  if (!current.value) return "";

  if (current.value.swapping.includes(index)) return "swapping";
  if (current.value.active.includes(index)) return "active";
  if (current.value.pivot?.includes(index)) return "pivot";

  return "";
}
</script>

<template>
  <div class="bars-container">
    <div
      v-for="(bar, index) in current?.bars || []"
      :key="bar.id"
      class="bar"
      :class="getBarClass(index)"
      :style="{ height: bar.value * 20 + 'px' }"
    >
      {{ bar.value }}
    </div>
  </div>
</template>

<style scoped>
.bars-container {
  display: flex;
  align-items: flex-end;
  gap: 10px;
  height: 400px;
}

.bar {
  width: 30px;
  background: steelblue;
}

.bar.active {
  background: orange;
}

.bar.swapping {
  background: red;
}

.bar.pivot {
  background: purple;
}
</style>