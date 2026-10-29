param(
    [string]$Key = "$env:USERPROFILE\.ssh\John-2026-10-27.pem",
    [Alias("Host")]
    [string]$DeployHost = "ubuntu@web.how-hard-can-it-be.com"
)

$ErrorActionPreference = "Stop"

$RepoRoot = $PSScriptRoot
$Remote = "/var/lib/swn"
$Binary = Join-Path $RepoRoot "swn"

if (-not (Test-Path -LiteralPath $Key)) {
    throw "SSH key not found: $Key"
}

Push-Location (Join-Path $RepoRoot "web")
try {
    npm run build
    if ($LASTEXITCODE -ne 0) {
        throw "npm run build failed with exit code $LASTEXITCODE"
    }
}
finally {
    Pop-Location
}

$env:GOOS = "linux"
$env:GOARCH = "arm64"
$env:CGO_ENABLED = "0"
try {
    go build -o $Binary ./cmd/server
    if ($LASTEXITCODE -ne 0) {
        throw "go build failed with exit code $LASTEXITCODE"
    }
}
finally {
    Remove-Item Env:GOOS, Env:GOARCH, Env:CGO_ENABLED -ErrorAction SilentlyContinue
}

$Dist = Join-Path $RepoRoot "web\dist"
if (-not (Test-Path -LiteralPath $Dist)) {
    throw "web/dist missing after build"
}

ssh -i $Key $DeployHost "rm -rf /tmp/swn.new /tmp/swn-dist.new"
if ($LASTEXITCODE -ne 0) {
    throw "remote tmp cleanup failed with exit code $LASTEXITCODE"
}

scp -i $Key $Binary "${DeployHost}:/tmp/swn.new"
if ($LASTEXITCODE -ne 0) {
    throw "scp binary failed with exit code $LASTEXITCODE"
}

scp -i $Key -r $Dist "${DeployHost}:/tmp/swn-dist.new"
if ($LASTEXITCODE -ne 0) {
    throw "scp dist failed with exit code $LASTEXITCODE"
}

$RemoteScript = @"
set -e
sudo install -m 755 /tmp/swn.new /usr/local/bin/swn
sudo rm -rf $Remote/web/dist
sudo mkdir -p $Remote/web
sudo mv /tmp/swn-dist.new $Remote/web/dist
sudo systemctl restart swn
sudo systemctl is-active swn
code=`$(curl -fsS -o /dev/null -w '%{http_code}' http://127.0.0.1:8080/)
echo "`$code"
test "`$code" = "200"
"@

$RemoteScript | ssh -i $Key $DeployHost bash -s
if ($LASTEXITCODE -ne 0) {
    throw "remote install/restart failed with exit code $LASTEXITCODE"
}

Write-Host "Deploy OK"
