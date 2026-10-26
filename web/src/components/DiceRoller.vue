<script lang="ts">
type DieOption = {
  sides: number
  label: string
}

export default {
  name: "DiceRoller",
  data(): {
    sides: number
    count: number
    faces: number[]
    rolledSides: number
    rolledCount: number
  } {
    return {
      sides: 6,
      count: 3,
      faces: [],
      rolledSides: 6,
      rolledCount: 3,
    }
  },
  computed: {
    dieOptions(): DieOption[] {
      return [
        { sides: 6, label: "d6" },
        { sides: 8, label: "d8" },
        { sides: 10, label: "d10" },
        { sides: 100, label: "d100" },
      ]
    },
    sortedFaces(): number[] {
      return [...this.faces].sort((a: number, b: number) => b - a)
    },
    sum(): number {
      return this.faces.reduce((total: number, face: number) => total + face, 0)
    },
    notation(): string {
      return `${this.rolledCount}d${this.rolledSides}`
    },
  },
  methods: {
    clampCount(value: number): number {
      const parsed = Math.floor(Number(value))
      if (!Number.isFinite(parsed)) {
        return 1
      }
      return Math.min(20, Math.max(1, parsed))
    },
    roll() {
      const count = this.clampCount(this.count)
      this.count = count
      const faces: number[] = []
      for (let i = 0; i < count; i++) {
        faces.push(Math.floor(Math.random() * this.sides) + 1)
      }
      this.faces = faces
      this.rolledSides = this.sides
      this.rolledCount = count
    },
  },
  mounted() {
    this.roll()
  },
}
</script>

<template>
  <form class="roller" @submit.prevent="roll">
    <label>
      Sides
      <select v-model.number="sides">
        <option v-for="option in dieOptions" :key="option.sides" :value="option.sides">
          {{ option.label }}
        </option>
      </select>
    </label>

    <label>
      Dice
      <input v-model.number="count" type="number" min="1" max="20" />
    </label>

    <button type="submit">Re-roll</button>
  </form>

  <p class="notation">{{ notation }}</p>

  <ol class="faces">
    <li v-for="(face, index) in sortedFaces" :key="index">{{ face }}</li>
  </ol>

  <p class="sum">Sum: {{ sum }}</p>
</template>

<style scoped>
.roller {
  display: flex;
  flex-wrap: wrap;
  align-items: end;
  gap: 1rem;
  margin-bottom: 1.5rem;
}

label {
  display: flex;
  flex-direction: column;
  gap: 0.25rem;
}

select,
input,
button {
  font: inherit;
  color: var(--color-text);
  background: var(--color-background-soft);
  border: 1px solid var(--color-border);
  border-radius: 4px;
  padding: 0.4rem 0.6rem;
}

input {
  width: 5rem;
}

button {
  cursor: pointer;
}

.notation,
.sum {
  margin-bottom: 0.75rem;
}

.faces {
  display: flex;
  flex-wrap: wrap;
  gap: 0.5rem;
  margin: 0 0 0.75rem;
  padding: 0;
  list-style: none;
}

.faces li {
  min-width: 2.5rem;
  padding: 0.4rem 0.6rem;
  text-align: center;
  border: 1px solid var(--color-border);
  border-radius: 4px;
  background: var(--color-background-mute);
}
</style>
