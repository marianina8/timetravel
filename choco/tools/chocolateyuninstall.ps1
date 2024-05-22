$ErrorActionPreference = 'Stop'

$packageName = 'timetravel'
$toolsDir = "$(Split-Path -Parent $MyInvocation.MyCommand.Definition)"
$exePath = Join-Path $toolsDir 'timetravel.exe'

# Remove the binary from the tools directory
Write-Host "Removing $exePath"
if (Test-Path $exePath) {
    Remove-Item -Path $exePath -Force
} else {
    Write-Host "$exePath does not exist."
}

# Remove from PATH
Write-Host "Uninstalling $packageName from PATH"
Uninstall-BinFile -Name 'timetravel'

Write-Host "$packageName has been uninstalled."
