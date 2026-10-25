Node is required on the machine that runs these commands. It is not installed
on the published host. Use Node `22.18.0` or newer on the 22 line, or
`24.12.0` or newer. Check with `node -v` from any directory.

Use the latest stable release of each named major. Do not install betas,
release candidates, or `dev` tags.

- Vue 3.5, Vue Router 5, and Pinia 4 come from the scaffold below. Do not pass
  `--vue-rc`, and do not install Router 4 or Pinia 2 or 3.
- The scaffold's `vite` dependency is the alias
  `npm:@voidzero-dev/vite-plus-core@latest`. Leave it, and leave
  `@vitejs/plugin-vue`. Do not run a separate `npm install vite`.
- `vue-tsc` 3 comes from the scaffold and checks `.vue` files. Leave the
  TypeScript pin at `~6.0.0`. TypeScript 7.0 is the latest compiler, and
  `vue-tsc` 3.3 cannot check `.vue` files against it. Move to TypeScript 7
  when a stable `vue-tsc` can.
- PixiJS stays on 8: `npm install pixi.js@8`.
- `vite-plugin-pwa` is installed from its `latest` tag.

Vite is the frontend development server and the production bundler.
`npm run dev` compiles each single-file component as the browser asks for it
and replaces a module in the open page after a save. `npm run build` writes
the static site to `web/dist`. The Go binary embeds that directory. Vite does
not run on the published host.

## Install

From the repository root. PowerShell needs the quoted `--`. Do not pass
`--force`. When asked what to do with the existing `web` directory, choose
Ignore. Remove deletes this file, the `src` READMEs, and `character.ts`.

```powershell
npm create vue@latest web '--' --ts --router --pinia
cd c:\Users\john-\Development\StarsWithoutNumber\web
npm install
npm install pixi.js@8
npm install -D vite-plugin-pwa
```

After the scaffold exists, add the PWA plugin and the dev proxy to
`vite.config.ts`. Leave the rest of the file as the scaffold wrote it. Icons
and a full manifest can wait until a home-screen icon matters.

```ts
import { VitePWA } from "vite-plugin-pwa"

export default defineConfig({
  plugins: [
    vue(),
    VitePWA({ registerType: "autoUpdate" }),
  ],
  server: {
    proxy: {
      "/api": "http://127.0.0.1:8080",
      "/ws": {
        target: "ws://127.0.0.1:8080",
        ws: true,
      },
    },
  },
})
```

`VitePWA({ registerType: 'autoUpdate' })` emits the service worker, so the app
can cache its shell. The install prompt is the browser asking to add the site
to the home screen. Chrome offers it once the site is a secure origin and the
manifest has a name plus a 192×192 icon and a 512×512 icon. `localhost` counts
as a secure origin, so a local PWA can install without TLS. Phones loading the
published app need HTTPS, which is covered in
[cmd/server](../cmd/server/README.md).

The canvas mount is in [src/views/README.md](src/views/README.md).

## Dev server

From `web`:

```powershell
npm run dev
```

Vite serves `http://localhost:5173` and proxies `/api` and `/ws` to the Go
process. `ws` stands for WebSocket, and `/ws` is the path for that
connection. `ws: true` on the proxy allows the browser to upgrade `/ws` from
HTTP to a WebSocket. `/api` stays ordinary HTTP. Nothing listens on `/ws`
until `cmd/server` exists. That server's route should use this same path.

## Package the frontend

From `web`:

```powershell
npm run build
```

That writes `web/dist`. The Linux binary build is in
[cmd/server](../cmd/server/README.md).
