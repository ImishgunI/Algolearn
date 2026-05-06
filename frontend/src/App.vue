<script setup lang="ts">
import { onMounted } from "vue";
import { fetchUser } from "./store/user";

onMounted(() => {
  fetchUser();
});
</script>

<template>
  <router-view />
</template>