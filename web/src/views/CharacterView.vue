<script lang="ts">
import DiceRoller from "@/components/DiceRoller.vue"

export default {
  name: "CharacterView",
  components: { DiceRoller },
}
</script>

<template>
  <main>
    <h1>Character generation</h1>
    <DiceRoller />
  </main>
</template>

<style scoped>
h1 {
  font-size: 2rem;
  color: var(--color-heading);
  margin-bottom: 1.5rem;
}
</style>
