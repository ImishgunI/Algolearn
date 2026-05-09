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

  if (current.value.leftRange && index >= current.value.leftRange[0] && index <= current.value.leftRange[1]) {
    return "left";
  }

  if (current.value.rightRange && index >= current.value.rightRange[0] && index <= current.value.rightRange[1]) {
    return "right";
  }

  if (current.value.mergedRange && index >= current.value.mergedRange[0] && index <= current.value.mergedRange[1]) {
    return "merged";
  }

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
.bar.left {
  background: blue;
}
.bar.right {
  background: green;
}
.bar.merged {
  background: purple;
}
</style>