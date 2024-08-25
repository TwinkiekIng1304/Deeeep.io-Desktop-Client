function Build-DCC {
    process {
        if (Test-Path -Path "dist") {
            Remove-Item -Path "dist" -Force -Recurse
        }

        go build -o dist/Deeeep.io-Desktop-Client.exe
        xcopy plugins dist\plugins\ /E /Q /Y
        $CompressParams = @{
            CompressionLevel = "Optimal"
            Force            = $true
            Path             = "dist\*"
            DestinationPath  = "Deeeep.io-Desktop-Client-Windows.zip"
        }
        Compress-Archive @CompressParams
    }
}

Build-DCC