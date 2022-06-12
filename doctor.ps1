[CmdLetBinding()]
param()

Import-Module $PSScriptRoot\scripts\utils.psm1 -Force

Confirm-IsValidPSVersion -Edition "Core" -Version "7.0" -Verbose:$VerbosePreference
Confirm-CommandLineTools -JsonFile "doctor-tools.json" -Verbose:$VerbosePreference
