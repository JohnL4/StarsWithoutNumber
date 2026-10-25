The campaign store is a Pinia option store. It owns the WebSocket at `/ws`:
on a message, replace the campaign snapshot; actions send commands and apply
the same payload locally. What `ws` means is in
[web/DEVELOPMENT.md](../../DEVELOPMENT.md).

The server side of that broadcast is
[internal/httpapi](../../../internal/httpapi/README.md).

```ts
import { defineStore } from "pinia"
import type { Character } from "@/api/character"

export const useCampaignStore = defineStore("campaign", {
  state: () => ({
    characters: [] as Character[],
    selectedId: null as string | null,
  }),
  actions: {
    connect() {
      // Open the WebSocket and replace state from each snapshot.
    },
    select(id: string) {
      this.selectedId = id
    },
  },
})
```
