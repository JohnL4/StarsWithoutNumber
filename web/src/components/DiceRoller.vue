<script lang="ts">
type DieOption = {
  sides: number
  label: string
}

type KeepMode = "highest" | "lowest"

export default {
  name: "DiceRoller",
  data(): {
    sides: number
    count: number
    keep: number
    keepMode: KeepMode
    lastCount: number
    faces: number[]
    rolledSides: number
    rolledCount: number
  } {
    return {
      sides: 6,
      count: 3,
      keep: 3,
      keepMode: "highest",
      lastCount: 3,
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
    keptCount(): number {
      return Math.min(this.clampKeep(this.keep), this.faces.length)
    },
    sum(): number {
      return this.sortedFaces.reduce((total: number, face: number, index: number) => {
        return this.isKept(index) ? total + face : total
      }, 0)
    },
    notation(): string {
      return `${this.rolledCount}d${this.rolledSides}`
    },
  },
  watch: {
    count(value: number | string) {
      if (!this.isCompleteCount(value)) {
        return
      }
      const n = this.clampCount(value)
      const previous = this.lastCount
      this.lastCount = n
      if (this.keep === previous || this.keep > n) {
        this.keep = n
      }
    },
  },
  methods: {
    isCompleteCount(value: number | string): boolean {
      if (typeof value === "string" && value.trim() === "") {
        return false
      }
      return Number.isFinite(Number(value))
    },
    clampCount(value: number | string): number {
      const parsed = Math.floor(Number(value))
      if (!Number.isFinite(parsed)) {
        return 1
      }
      return Math.min(20, Math.max(1, parsed))
    },
    clampKeep(value: number | string): number {
      const max = this.clampCount(this.count)
      const parsed = Math.floor(Number(value))
      if (!Number.isFinite(parsed)) {
        return max
      }
      return Math.min(max, Math.max(1, parsed))
    },
    isKept(index: number): boolean {
      const total = this.sortedFaces.length
      const keep = this.keptCount
      if (keep >= total) {
        return true
      }
      if (this.keepMode === "lowest") {
        return index >= total - keep
      }
      return index < keep
    },
    roll() {
      const count = this.clampCount(this.count)
      this.count = count
      this.lastCount = count
      this.keep = this.clampKeep(this.keep)
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

    <div class="keep">
      <select v-model="keepMode" aria-label="Keep highest or lowest">
        <option value="highest">Highest</option>
        <option value="lowest">Lowest</option>
      </select>
      <input
        v-model.number="keep"
        type="number"
        min="1"
        :max="clampCount(count)"
        aria-label="Dice to keep"
      />
    </div>

    <button type="submit">Roll</button>
  </form>

  <p class="notation">{{ notation }}</p>

  <ol class="faces">
    <li
      v-for="(face, index) in sortedFaces"
      :key="index"
      :class="{ dropped: !isKept(index) }"
    >
      {{ face }}
    </li>
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

label,
.keep {
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

.faces li.dropped {
  opacity: 0.4;
}
</style>
