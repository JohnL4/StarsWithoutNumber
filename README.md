# Stars Without Number

* https://sine-nomine-publishing.myshopify.com/collections/stars-without-number
* https://www.drivethrurpg.com/en/product/226996/stars-without-number-revised-edition

## Todo/Roadmap

- [ ] Deploy to Amazon
- [ ] Character generation
  - [ ] Pauses for user interaction (swap two? add 14?)
  - [ ] Look up rule for 14
- [ ] Sector generation -- needed for character generation, but more complicated
- [ ] Learn:
  - [ ] Learn Vue (duh)
  - [ ] Learn TypeScript (and modern Javascript)
  - [ ] Learn modern html, css
  - [ ] Learn Go
  - [ ] Vite? Playwright?


## Technical details

Go and SQLite are the authority for characters, campaigns, and anything that
must survive a refresh. Pinia holds the latest snapshot the server sent. Pixi
holds sprites, the camera, and tweens.

The domain is the rules code, such as `internal/character`. To call the domain
is to run that code. Handlers do this, then save through the store. Domain
packages import neither HTTP nor SQL.

`go.mod` names this repository `github.com/JohnL4/StarsWithoutNumber` and
records the Go version. That name is why `cmd/cli` can import
`github.com/JohnL4/StarsWithoutNumber/internal/character`: the import path
starts with the module path, and the rest is the directory under the repo
root. `go.sum` appears when a third-party module is downloaded and stores
checksums. This module has no third-party requirements yet.

## Layout

- [cmd/cli](cmd/cli/README.md) prints one rolled character.
- [cmd/server](cmd/server/README.md) will embed the built frontend, serve the
  API, and broadcast on a WebSocket. Publishing steps live there.
- [internal/character](internal/character/README.md) holds dice and rules.
- [internal/store](internal/store/README.md) will hold SQLite repositories.
- [internal/httpapi](internal/httpapi/README.md) will hold JSON handlers and
  the broadcast.
- [web/DEVELOPMENT.md](web/DEVELOPMENT.md) is the frontend install, dev server,
  and production build, including the Options API rule for new Vue components
  and Pinia stores.
- [web/src/api](web/src/api/README.md) is the fetch client and the TypeScript
  character shape.
- [web/src/vtt](web/src/vtt/README.md) is the virtual tabletop.
- [web/src/components](web/src/components/README.md) is the HTML HUD.
- [web/src/views](web/src/views/README.md) mounts the canvas.
- [web/src/stores](web/src/stores/README.md) is the Pinia campaign store.
- [data](data/README.md) is the runtime SQLite directory.
