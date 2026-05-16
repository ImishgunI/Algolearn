<script setup lang="ts">
import { ref, onMounted, computed, watch } from "vue";
import { useRoute } from "vue-router";
import { getLesson, type Lesson } from "../api/lessons";
import { getComments, createComment, getFavoriteStatus, toggleFavorite } from "../api/lessons";
import { user } from "../store/user";
import CodeBlock from "../components/CodeBlock.vue";
import { algorithmRegistry } from "../algorithmCodes";

const route = useRoute();
const lessonId = computed(() => {
  const id = Number(route.params.lessonId);
  return isNaN(id) ? null : id;
});

const lesson = ref<Lesson | null>(null);
const comments = ref<any[]>([]);
const isFavorited = ref(false);
const newComment = ref("");


const currentCode = computed(() => {
  if (!lesson.value?.algorithm_type) return [];
  const meta = algorithmRegistry[lesson.value.algorithm_type];
  return meta?.code || [];
});

async function loadLessonData(id: number) {
  try {
    lesson.value = await getLesson(id);
    comments.value = await getComments(id);
    const fav = await getFavoriteStatus(id);
    isFavorited.value = fav.favorited;
  } catch (e) {
    console.error(e);
  }
}

watch(lessonId, (newId, oldId) => {
  if (newId && newId !== oldId) {
    loadLessonData(newId);
  }
});


onMounted(() => {
  if (lessonId.value) {
    loadLessonData(lessonId.value);
  }
});

async function handleAddComment() {
  if (!newComment.value.trim() || !lessonId.value) return;
  try {
    const cmt = await createComment(lessonId.value, newComment.value);
    comments.value.push(cmt);
    newComment.value = "";
  } catch (e) {
    console.error(e);
  }
}

async function handleToggleFavorite() {
  if (!lessonId.value) return;
  try {
    const res = await toggleFavorite(lessonId.value);
    isFavorited.value = res.favorited;
  } catch (e) {
    console.error(e);
  }
}
</script>

<template>
  <div v-if="lesson">
    <h1>{{ lesson.title }}</h1>
    <div class="text-content" v-html="lesson.theory.replace(/\n/g, '<br/>')" />

    <div v-if="currentCode.length" class="code-section">
      <h3>Пример кода</h3>
      <CodeBlock :codeLines="currentCode" :activeLines="[]" />
    </div>

    <div v-if="lesson.algorithm_type" class="action-section">
      <router-link
        :to="`/visualize?algorithm=${lesson.algorithm_type}&input=5,3,1,8,4,2,7`"
        class="btn btn-primary"
      >
        Открыть в визуализаторе
      </router-link>
      <button class="btn btn-secondary" @click="handleToggleFavorite">
        {{ isFavorited ? '❤️ В избранном' : '❣️ Добавить в избранное' }}
      </button>
    </div>

    <div class="comments-section">
      <h3>Комментарии</h3>
      <div class="comment-form" v-if="user">
        <textarea v-model="newComment" placeholder="Ваш комментарий..."></textarea>
        <button class="btn btn-primary" @click="handleAddComment">Отправить</button>
      </div>
      <div v-for="c in comments" :key="c.id" class="comment">
        <strong>{{ c.user_name }}&nbsp;</strong>
        <small>{{ new Date(c.created_at).toLocaleString() }}</small>
        <p>{{ c.body }}</p>
      </div>
    </div>
  </div>
  <div v-else-if="!lesson && !lessonId">Неверный ID урока</div>
  <div v-else>Загрузка...</div>
</template>

<style scoped>
.text-content { margin: 24px 0; line-height: 1.8; font-size: 1.05rem; }
.code-section { margin: 24px 0; }
.action-section { margin-top: 32px; display: flex; gap: 12px; align-items: center; }
.comments-section { margin-top: 40px; }
.comment-form { margin-bottom: 20px; }
textarea { width: 100%; padding: 10px; border-radius: 8px; border: 1px solid var(--border); background: var(--surface); color: var(--text); }
.comment { background: var(--surface); padding: 9px; border-radius: 8px; margin-bottom: 12px; }
.btn { display: inline-block; margin: 8px 8px 0 0; padding: 10px 20px; border-radius: 8px; font-weight: 600; cursor: pointer; }
.btn-primary { background: var(--primary); color: white; }
.btn-secondary { background: var(--surface2, rgba(0,0,0,0.08)); color: var(--text); }
</style>