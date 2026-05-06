<template>
  <div class="auth-page">
    <div class="auth-card">
      <slot />
    </div>
  </div>
</template>

<style scoped>
.auth-page {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: var(--bg);
}

.auth-card {
  background: var(--surface);
  padding: 40px 32px;
  border-radius: 16px;
  width: 100%;
  max-width: 420px;
  box-shadow: 0 10px 15px -3px rgb(0 0 0 / 0.1);
}
</style>