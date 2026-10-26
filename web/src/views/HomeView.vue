<script lang="ts">
export default {
  name: "HomeView",
}
</script>

<template>
  <main class="vtt-home">
    <h1>Virtual tabletop</h1>
    <nav>
      <RouterLink to="/character">Character generation</RouterLink>
      <RouterLink to="/sector">Sector generation</RouterLink>
    </nav>
  </main>
</template>

<style scoped>
.vtt-home h1 {
  font-size: 2rem;
  color: var(--color-heading);
  margin-bottom: 1.5rem;
}

.vtt-home nav {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  gap: 0.75rem;
  font-size: 1.25rem;
}
</style>
