<script lang="ts">
export default {
  name: "App",
}
</script>

<template>
  <nav class="app-nav">
    <RouterLink to="/">VTT</RouterLink>
    <RouterLink to="/character">Character</RouterLink>
    <RouterLink to="/sector">Sector</RouterLink>
    <RouterLink to="/hello">Original Hello</RouterLink>
  </nav>
  <RouterView />
</template>

<style scoped>
.app-nav {
  display: flex;
  flex-wrap: wrap;
  gap: 0.5rem 0;
  margin-bottom: 2rem;
  font-size: 1rem;
}

.app-nav a {
  display: inline-block;
  padding: 0 1rem;
  border-left: 1px solid var(--color-border);
}

.app-nav a:first-of-type {
  border: 0;
  padding-left: 0;
}

.app-nav a.router-link-exact-active {
  color: var(--color-text);
}

.app-nav a.router-link-exact-active:hover {
  background-color: transparent;
}
</style>
