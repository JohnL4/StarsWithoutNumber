`go run ./cmd/server`, from the repo root, listens on `127.0.0.1:8080` and
serves `web/dist` from disk. `/api` and `/ws` respond 404. Embed remains how
the Linux binary will carry the frontend. What `ws` means is in
[web/DEVELOPMENT.md](../../web/DEVELOPMENT.md).

`go run` compiles and serves. After a Go edit, stop that process and start it
again. `npm run dev` does not recompile Go, and `go build` is not part of
that loop. A frontend build is not required for an API change while the
browser is on port 5173. The npm commands are in
[web/DEVELOPMENT.md](../../web/DEVELOPMENT.md).

The process listens on `127.0.0.1`. Caddy, or another reverse proxy,
terminates HTTPS in front of it so phones can install the PWA. A minimal site
block:

```
example.com {
	reverse_proxy 127.0.0.1:8080
}
```

`go build` produces a binary. Build the frontend first. Those commands are in
[web/DEVELOPMENT.md](../../web/DEVELOPMENT.md). From the repo root:

```powershell
$env:GOOS = "linux"
$env:GOARCH = "amd64"
$env:CGO_ENABLED = "0"
go build -o swn ./cmd/server
```

Set `GOARCH` to `arm64` when the host is ARM. Reset both variables before
building a Windows binary again:

```powershell
Remove-Item Env:GOOS
Remove-Item Env:GOARCH
```

Copy `swn` to the Linux machine. The database file is described in
[data/README.md](../../data/README.md). Create an empty data directory the
service user can write, such as `/var/lib/swn`. This stub resolves `web/dist`
from its working directory, so `go run` belongs at the repo root.

```ini
[Unit]
Description=Stars Without Number
After=network.target

[Service]
ExecStart=/usr/local/bin/swn
WorkingDirectory=/var/lib/swn

[Install]
WantedBy=multi-user.target
```
