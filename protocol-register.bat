@echo off
@echo.
if not exist "%~dp0player-protocol.exe" (
  echo Warning: Can't find player-protocol.exe.
  echo Did you compile it successfully?
  @echo.
  pause
  exit /b
)

echo If you see "ERROR: Access is denied." then you need to right click and use "Run as Administrator".
@echo.
echo Associating vlc:// and potplayer:// with player-protocol.exe...

rem Register VLC protocol
reg add HKCR\vlc /ve /t REG_SZ /d "URL:VLC Protocol" /f
reg add HKCR\vlc /v "URL Protocol" /t REG_SZ /d "" /f
reg add HKCR\vlc\shell\open\command /ve /t REG_SZ /d "\"%~dp0player-protocol.exe\" %%1" /f

rem Register PotPlayer protocol
reg add HKCR\potplayer /ve /t REG_SZ /d "URL:PotPlayer Protocol" /f
reg add HKCR\potplayer /v "URL Protocol" /t REG_SZ /d "" /f
reg add HKCR\potplayer\shell\open\command /ve /t REG_SZ /d "\"%~dp0player-protocol.exe\" %%1" /f

@echo.
echo Protocols registered successfully.
pause