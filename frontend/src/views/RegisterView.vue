<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";
import { register } from "../api/auth";
import AuthLayout from "../components/AuthLayout.vue";

const name = ref("");
const surname = ref("");
const email = ref("");
const password = ref("");

const router = useRouter();

async function submit() {
  await register({
    name: name.value,
    surname: surname.value,
    email: email.value,
    password: password.value,
  });

  router.push("/login");
}
</script>

<template>
  <AuthLayout>
    <h2>Register</h2>

    <input v-model="name" placeholder="Name" />
    <input v-model="surname" placeholder="Surname" />
    <input v-model="email" placeholder="Email" />
    <input v-model="password" type="password" placeholder="Password" />

    <button @click="submit">Create account</button>

    <p>
      Already have account?
      <router-link to="/login">Login</router-link>
    </p>
  </AuthLayout>
</template>

<style scoped>
input {
  width: 100%;
  margin: 6px 0;
  padding: 10px;
  border-radius: 6px;
  border: none;
}

button {
  width: 100%;
  padding: 10px;
  margin-top: 10px;
  background: green;
  border: none;
  color: white;
  border-radius: 6px;
}
</style>