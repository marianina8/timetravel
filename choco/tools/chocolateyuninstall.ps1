$ErrorActionPreference = 'Stop'; # stop on all errors

# Remove any existing service with the same name
Write-Output "################################"
Write-Output "Removing any old $packageName service"
if (Get-Service $packageName -ErrorAction 'SilentlyContinue'){
        Stop-Service $packageName -Force
        (Get-WmiObject -Class Win32_Service -filter "Name='$packageName'").delete()
}
Write-Output "################################"