function Confirm-IsValidPSVersion {
  [CmdletBinding()]
  param (
    [Parameter(Mandatory = $true)]
    [string]
    $Edition,
    [Parameter(Mandatory = $true)]
    [string]
    $Version
  )

  Write-Verbose ("❏ Checking if PowerShell is $Edition v$($Version)+")

  if ($PSVersionTable.PSEdition -ne $Edition) {
    Write-Host ("✕ This script requires PowerShell $Edition edition, " +
      "but current edition is $PSVersionTable.PSEdition") `
      -ForegroundColor Red

    exit 1
  }

  $SemVersion = $null
  if (-not [System.Management.Automation.SemanticVersion]::TryParse($Version, [ref] $SemVersion)) {
    Write-Host ("✕ The expected version '$Version' cannot be parsed.") `
      -ForegroundColor Red

    exit 1
  }

  if ($SemVersion -gt $PSVersionTable.PSVersion) {
    Write-Host ("✕ This script requires PowerShell $Version or later, " +
      "but current version is $($PSVersionTable.PSVersion)") `
      -ForegroundColor Red

    exit 1
  }

  Write-Verbose ("✓ PowerShell is $($PSVersionTable.PSEdition) v$($PSVersionTable.PSVersion)")
}


function Exit-OnSwitch {
  param (
    [Parameter()]
    [int]
    $ExitCode = 1,
    [Parameter()]
    [switch]
    $ExitOnError
  )

  if ($ExitOnError -eq $true) {
    exit $ExitCode
  }
}


function Confirm-CommandLineTool {
  [CmdletBinding()]
  param (
    [Parameter(Mandatory = $true)]
    [string]
    $Title,
    [Parameter(Mandatory = $true)]
    [string]
    $Command,
    [Parameter()]
    [System.Version]
    $Version,
    [Parameter()]
    [string]
    $VersionArgs,
    [Parameter()]
    [string]
    $VersionPattern,
    [Parameter()]
    [int]
    $VersionPatternGroup = 0,
    [Parameter()]
    [string]
    $DownloadUrl,
    [Parameter()]
    [switch]
    $Optional,
    [Parameter()]
    [switch]
    $ExitOnError
  )

  Write-Verbose ("❏ Checking if $Title is installed")

  $AppInfo = Get-Command $Command -ErrorAction SilentlyContinue
  if (-not $?) {
    if ($Optional -eq $false) {
      Write-Host ("❌ $Title is not installed: $DownloadUrl") -ForegroundColor Red
      Exit-OnSwitch -ExitOnError:$ExitOnError
      return
    } else {
      Write-Host ("❓ $Title is not installed (optional): $DownloadUrl") -ForegroundColor Yellow
      return
    }
  }

  Write-Verbose ("✓ Found $Title at $($AppInfo.Source) with version $($AppInfo.Version)")

  if ($PSBoundParameters.ContainsKey("Version")) {
    $ActualVersion = $AppInfo.Version
    if ($PSBoundParameters.ContainsKey("VersionArgs")) {
      $VersionResult = ""
      Invoke-Expression -Command "$Command $VersionArgs" -OutVariable VersionResult -ErrorAction SilentlyContinue | Out-Null
      if (-not $?) {
        Write-Host ("❌ $Title version cannot be determined: $DownloadUrl") -ForegroundColor Red
        Exit-OnSwitch -ExitOnError:$ExitOnError
        return
      }

      if ($PSBoundParameters.ContainsKey("VersionPattern")) {
        $Match = [regex]::Match($VersionResult, $VersionPattern)
        if ($Match.Success -and ($VersionPatternGroup -lt $Match.Groups.Count)) {
          $VersionResult = $Match.Groups[$VersionPatternGroup].Value
        } else {
          Write-Host ("❌ $Title version cannot be determined: $DownloadUrl") -ForegroundColor Red
          Exit-OnSwitch -ExitOnError:$ExitOnError
          return
        }
      } ### if ($PSBoundParameters.ContainsKey("VersionPattern"))

      if (-not [System.Version]::TryParse($VersionResult, [ref] $ActualVersion)) {
        Write-Host ("❌ $Title version cannot be parsed: $DownloadUrl") -ForegroundColor Red
        Exit-OnSwitch -ExitOnError:$ExitOnError
        return
      }
    } ### if ($PSBoundParameters.ContainsKey("VersionArgs"))

    if ($Version -ne $null) {
      if ($ActualVersion -lt $Version) {
        Write-Host ("❌ $Title is installed, but version is $($ActualVersion) < $Version`: $DownloadUrl") `
          -ForegroundColor Red
        Exit-OnSwitch -ExitOnError:$ExitOnError
        return
      }
    } ### if ($Version -ne $null)
  } ### if ($PSBoundParameters.ContainsKey("Version"))

  if ($null -ne $ActualVersion) {
    Write-Host ("✅ $Title is installed with version $($ActualVersion)")
  } else {
    Write-Host ("✅ $Title is installed")
  }
}


function Set-DotEnv {
  [CmdLetBinding(SupportsShouldProcess = $true, ConfirmImpact = "Low")]
  param(
    [Parameter()]
    [string]
    $DotEnvFile = ".\.env",
    [Parameter()]
    [System.EnvironmentVariableTarget]
    $Target = [System.EnvironmentVariableTarget]::Process
  )

  if (!(Test-Path $DotEnvFile)) {
    Write-Verbose ("$DotEnvFile does not exist")
    return
  }

  $DotEnvContent = Get-Content $DotEnvFile -ErrorAction Stop
  if ($null -eq $DotEnvContent -or $DotEnvContent.Count -eq 0) {
    Write-Verbose ("$DotEnvFile is empty")
    return
  }

  $LineCount = 0
  foreach ($Line in $DotEnvContent) {
    $LineCount += 1

    # Skip empty lines and comments
    if ([string]::IsNullOrWhiteSpace($Line) -or $Line -match "^\s*#") {
      continue
    }

    # Split the line into key and value
    if (!($Line -match "^([^=`"']+)=(.*)$")) {
      Write-Verbose ("$DotEnvFile line $LineCount` is invalid: $Line")
      continue
    }

    $Key = $matches[1]
    $Value = $matches[2]

    Write-Verbose ("Setting $Key to $Value")

    if ($PSCmdlet.ShouldProcess("Env Var $Key", "Set Value $Value")) {
      [System.Environment]::SetEnvironmentVariable($Key, $Value, $Target) | Out-Null
    }
  }
}
