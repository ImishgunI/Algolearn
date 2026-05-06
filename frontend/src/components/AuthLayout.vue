<template>
  <div class="auth">
    <div class="card">
      <slot />
    </div>
  </div>
</template>

<style scoped>
.auth {
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background: #0f172a;
  font-family: sans-serif;
}

.card {
  width: 350px;
  padding: 30px;
  border-radius: 12px;
  background: #1e293b;
  color: white;
  box-shadow: 0 10px 30px rgba(0,0,0,0.3);
}
</style>