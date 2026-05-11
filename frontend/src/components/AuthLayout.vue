<template>
  <div class="auth-page">
    <header class="auth-header">
      <router-link to="/" class="logo-link">← На главную</router-link>
    </header>
    <div class="auth-card">
      <slot />
    </div>
  </div>
</template>

<style scoped>
.auth-page {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  background: var(--bg);
  padding: 20px;
}

.auth-header {
  width: 100%;
  max-width: 420px;
  text-align: left;
  margin-bottom: 20px;
}

.logo-link {
  color: var(--primary);
  text-decoration: none;
  font-weight: 600;
  font-size: 14px;
  transition: opacity 0.2s;
}

.logo-link:hover {
  opacity: 0.8;
}

.auth-card {
  background: var(--surface);
  backdrop-filter: blur(18px);
  border: 1px solid var(--surface-border);
  padding: 40px 32px;
  border-radius: var(--radius);
  width: 100%;
  max-width: 420px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.05);
}
</style>