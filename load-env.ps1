[CmdLetBinding(SupportsShouldProcess = $true, ConfirmImpact = "Low")]
param(
    [Parameter()]
    [string]
    $DotEnvFile = ".\.env",
    [Parameter()]
    [System.EnvironmentVariableTarget]
    $Target = [System.EnvironmentVariableTarget]::Process
)

Import-Module $PSScriptRoot\scripts\utils.psm1 -Force -Verbose:$false

Set-DotEnv -DotEnvFile $DotEnvFile -Target $Target -Verbose:$VerbosePreference -WhatIf:$WhatIfPreference
