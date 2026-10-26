<script lang="ts">
import DiceRoller from "@/components/DiceRoller.vue"

export default {
  name: "App",
  components: { DiceRoller },
  data(): { diceRoutes: string[] } {
    return {
      diceRoutes: ["character", "sector"],
    }
  },
  computed: {
    diceActive(): boolean {
      const name = this.$route.name
      return typeof name === "string" && this.diceRoutes.includes(name)
    },
  },
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
  <DiceRoller :active="diceActive" />
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
