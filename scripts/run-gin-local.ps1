Set-StrictMode -Version Latest
$ErrorActionPreference = "Stop"

$repoRoot = Split-Path -Parent $PSScriptRoot
$binDir = Join-Path $repoRoot "bin"
$exePath = Join-Path $binDir "gocodes-gin.exe"

New-Item -ItemType Directory -Force -Path $binDir | Out-Null

go build -o $exePath $repoRoot
& $exePath
