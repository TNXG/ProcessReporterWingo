cargo build --release

if %COMPUTERNAME%==TNXG-PC (
    upx --best --lzma .\target\release\GetSmtcInfo.exe
)

copy .\target\release\GetSmtcInfo.exe .\