<script setup lang="ts">
import { ref } from "vue";
import { authFetch } from "../api/http";

const title = ref("");

async function createCourse() {
  await authFetch("/api/admin/course", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ title: title.value }),
  });
}
</script>

<template>
  <div>
    <h1>Admin panel</h1>

    <input v-model="title" placeholder="Course title" />
    <button @click="createCourse">Create</button>
  </div>
</template>