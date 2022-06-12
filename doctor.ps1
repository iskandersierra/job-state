Import-Module $PSScriptRoot\scripts\utils.psm1 -Force

Confirm-IsValidPSVersion -Edition "Core" -Version "7.0"

Confirm-CommandLineTool -Title "Git" -Command "git" -Version "2.33" -DownloadUrl "https://git-scm.com/downloads"
# Confirm-CommandLineTool -Title "Microsoft Tye" -Command "tye" -Version "0.11.0" -VersionArgs "--version" -VersionPattern "^[0-9\.]+" -DownloadUrl "https://www.nuget.org/packages/Microsoft.Tye"
Confirm-CommandLineTool -Title "Docker" -Command "docker" -Version "20.10" -VersionArgs "--version" -VersionPattern "version ([0-9\.]+)" -VersionPatternGroup 1 -DownloadUrl "https://www.docker.com/products/docker-desktop/"

Confirm-CommandLineTool -Title "Make" -Command "make" -Version "4.3" -VersionArgs "--version" -VersionPattern "GNU Make ([0-9\.]+)" -VersionPatternGroup 1 -Optional -DownloadUrl "https://community.chocolatey.org/packages/make"

Confirm-CommandLineTool -Title "Go" -Command "go" -Version "1.18" -VersionArgs "version" -VersionPattern "go version go([0-9\.]+)" -VersionPatternGroup 1 -Optional -DownloadUrl "https://golang.org/dl/"
Confirm-CommandLineTool -Title "Air" -Command "air" -Optional -DownloadUrl "https://github.com/cosmtrek/air#installation"
