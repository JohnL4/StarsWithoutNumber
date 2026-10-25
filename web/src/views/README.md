This view owns the canvas. One Pixi application is mounted from here and destroyed in `beforeUnmount`.

Views use the Options API. The sample below is the required style. Do not use
`<script setup>` or the Composition API.

```vue
<script lang="ts">
import { mapState } from "pinia"
import { createVtt } from "@/vtt/createVtt"
import { useCampaignStore } from "@/stores/campaign"
import CharacterSheet from "@/components/CharacterSheet.vue"

type Vtt = {
  select(id: string): void
  destroy(): void
}

export default {
  components: { CharacterSheet },
  computed: {
    ...mapState(useCampaignStore, ["selectedId"]),
  },
  data(): { vtt: Vtt | null } {
    return { vtt: null }
  },
  watch: {
    selectedId(id: string) {
      this.vtt?.select(id)
    },
  },
  mounted() {
    createVtt(this.$refs.host as HTMLElement).then((vtt: Vtt) => {
      this.vtt = vtt
    })
  },
  beforeUnmount() {
    this.vtt?.destroy()
  },
}
</script>

<template>
  <div class="stage">
    <div ref="host" class="canvas-host" />
    <CharacterSheet class="hud" />
  </div>
</template>

<style scoped>
.stage {
  position: relative;
  width: 100%;
  height: 100%;
}
.canvas-host {
  position: absolute;
  inset: 0;
}
.hud {
  position: absolute;
  inset: 0;
  pointer-events: none;
}
</style>
```
