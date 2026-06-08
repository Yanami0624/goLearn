Set-StrictMode -Version Latest
$ErrorActionPreference = "Stop"

$repoRoot = Split-Path -Parent $PSScriptRoot
$binDir = Join-Path $repoRoot "bin"
$exePath = Join-Path $binDir "gocodes-gin.exe"
$ruleName = "Gocodes Gin Dev 8080"

$principal = New-Object Security.Principal.WindowsPrincipal([Security.Principal.WindowsIdentity]::GetCurrent())
$isAdmin = $principal.IsInRole([Security.Principal.WindowsBuiltInRole]::Administrator)
if (-not $isAdmin) {
	Write-Error "Please run this script in an elevated PowerShell window."
}

New-Item -ItemType Directory -Force -Path $binDir | Out-Null
go build -o $exePath $repoRoot

$existingRule = Get-NetFirewallRule -DisplayName $ruleName -ErrorAction SilentlyContinue
if ($existingRule) {
	Remove-NetFirewallRule -DisplayName $ruleName
}

New-NetFirewallRule `
	-DisplayName $ruleName `
	-Direction Inbound `
	-Action Allow `
	-Program $exePath `
	-Protocol TCP `
	-LocalPort 8080 `
	-Profile Private

Write-Host "Firewall rule installed for $exePath on TCP 8080, Private profile only."
