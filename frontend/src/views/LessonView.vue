<script setup lang="ts">
import { computed } from "vue";
import { useRoute } from "vue-router";
import { sortingLessons } from "../data/sortingLessons";
import CodeBlock from "../components/CodeBlock.vue";

const route = useRoute();
const lessonId = computed(() => route.params.lessonId as string);
const lesson = computed(() => sortingLessons.find((l) => l.id === lessonId.value));

const currentCode = computed(() => {
  if (!lesson.value?.codeExample) return [];
  return lesson.value.codeExample.split("\n");
});
</script>

<template>
  <div v-if="lesson">
    <h1>{{ lesson.title }}</h1>
    <div class="text-content" v-html="lesson.content.replace(/\n/g, '<br/>')"></div>

    <div v-if="lesson.codeExample" class="code-section">
      <h3>Пример кода (Go)</h3>
      <CodeBlock :codeLines="currentCode" :activeLines="[]" />
    </div>

    <div v-if="lesson.algorithm" class="action-section">
      <router-link
        :to="`/visualize?algorithm=${lesson.algorithm}&input=5,3,1,8,4,2,7`"
        class="btn btn-primary"
      >
        Открыть в визуализаторе
      </router-link>
    </div>
  </div>
  <div v-else>
    <p>Урок не найден</p>
  </div>
</template>

<style scoped>
.text-content {
  margin: 24px 0;
  line-height: 1.8;
  font-size: 1.05rem;
}
.code-section {
  margin: 24px 0;
}
.action-section {
  margin-top: 32px;
}
.btn-primary {
  display: inline-block;
  padding: 12px 24px;
  background: linear-gradient(135deg, #6366f1, #4f46e5);
  color: white;
  border-radius: 8px;
  text-decoration: none;
  font-weight: 600;
}
.btn-primary:hover {
  background: linear-gradient(135deg, #4f46e5, #4338ca);
}
</style>