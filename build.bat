go build -ldflags="-s -w"

if %COMPUTERNAME%==TNXG-PC (
    upx --best --lzma .\ProcessReporterWingo.exe
)