<script lang="ts">
import { useDraggable } from "@vueuse/core"
import { ref } from "vue"

type DieOption = {
  sides: number
  label: string
}

type KeepMode = "highest" | "lowest"

export default {
  name: "DiceRoller",
  props: {
    active: {
      type: Boolean,
      required: true,
    },
  },
  setup() {
    const panel = ref<HTMLElement | null>(null)
    const handle = ref<HTMLElement | null>(null)
    const { style } = useDraggable(panel, {
      handle,
      initialValue: {
        x: Math.max(16, window.innerWidth - 304),
        y: 72,
      },
      containerElement: document.documentElement,
      restrictInView: true,
      preventDefault: true,
    })
    return { panel, handle, dragStyle: style }
  },
  data(): {
    open: boolean
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
      open: false,
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
    show() {
      this.open = true
    },
    dismiss() {
      this.open = false
    },
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
  <Teleport to="body">
    <div class="dice-layer">
      <button v-show="active && !open" type="button" class="launcher" @click="show">
        Dice
      </button>

      <Transition name="dice-fade">
        <div v-show="active && open" ref="panel" class="panel" :style="dragStyle">
        <header class="header">
          <div ref="handle" class="handle">Dice</div>
          <button type="button" class="close" @click="dismiss">Close</button>
        </header>

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
        </div>
      </Transition>
    </div>
  </Teleport>
</template>

<style scoped>
.dice-layer {
  position: fixed;
  inset: 0;
  z-index: 30;
  pointer-events: none;
}

.launcher,
.panel {
  pointer-events: auto;
}

.launcher {
  position: fixed;
  top: 72px;
  right: 16px;
}

.panel {
  position: fixed;
  width: 18rem;
  padding: 0.75rem;
  background: var(--color-background-soft);
  border: 1px solid var(--color-border);
  border-radius: 6px;
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.18);
}

.dice-fade-enter-active,
.dice-fade-leave-active {
  transition: opacity 200ms ease;
}

.dice-fade-enter-from,
.dice-fade-leave-to {
  opacity: 0;
}

.header {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  margin-bottom: 0.75rem;
}

.handle {
  flex: 1;
  cursor: grab;
  user-select: none;
  touch-action: none;
}

.handle:active {
  cursor: grabbing;
}

.roller {
  display: flex;
  flex-wrap: wrap;
  align-items: end;
  gap: 0.75rem;
  margin-bottom: 0.75rem;
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
.sum,
.faces li {
  user-select: text;
}

.notation,
.sum {
  margin-bottom: 0.75rem;
}

.sum {
  margin-bottom: 0;
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
