<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";
import { login } from "../api/auth";
import { setToken } from "../utils/token";
import AuthLayout from "../components/AuthLayout.vue";

const email = ref("");
const password = ref("");
const router = useRouter();

async function submit() {
  const res = await login({
    email: email.value,
    password: password.value,
  });

  setToken(res.token);
  router.push("/visualize");
}
</script>

<template>
  <AuthLayout>
    <h2>Login</h2>

    <input v-model="email" placeholder="Email" />
    <input v-model="password" type="password" placeholder="Password" />

    <button @click="submit">Login</button>

    <p>
      No account?
      <router-link to="/register">Register</router-link>
    </p>
  </AuthLayout>
</template>

<style scoped>
input {
  width: 100%;
  margin: 8px 0;
  padding: 10px;
  border-radius: 6px;
  border: none;
}

button {
  width: 100%;
  padding: 10px;
  margin-top: 10px;
  background: steelblue;
  border: none;
  color: white;
  border-radius: 6px;
}
</style>