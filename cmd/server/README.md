This process will embed `web/dist`, serve the API, and broadcast campaign
snapshots on a WebSocket at `/ws`. There is no `main` package here yet, so
`go build ./cmd/server` fails until that file exists. What `ws` means is in
[web/DEVELOPMENT.md](../../web/DEVELOPMENT.md).

The process listens on `127.0.0.1`. Caddy, or another reverse proxy,
terminates HTTPS in front of it so phones can install the PWA. A minimal site
block:

```
example.com {
	reverse_proxy 127.0.0.1:8080
}
```

Build the frontend first. Those commands are in
[web/DEVELOPMENT.md](../../web/DEVELOPMENT.md). From the repo root, with a
`main` package in this directory that embeds `web/dist`:

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
service user can write, such as `/var/lib/swn`. Run the binary there.

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
