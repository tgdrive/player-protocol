@echo off
@echo.

rem Check for administrative privileges
net session >nul 2>&1
if %errorLevel% neq 0 (
    echo This script requires administrative privileges.
    echo Please run it as an administrator.
    pause
    exit /b
)

echo Deregistering vlc:// and potplayer:// protocols...
@echo.

rem Remove VLC protocol registration
reg delete HKCR\vlc /f
if %errorLevel% equ 0 (
    echo VLC protocol deregistered successfully.
) else (
    echo Failed to deregister VLC protocol. It may not have been registered.
)

@echo.

rem Remove PotPlayer protocol registration
reg delete HKCR\potplayer /f
if %errorLevel% equ 0 (
    echo PotPlayer protocol deregistered successfully.
) else (
    echo Failed to deregister PotPlayer protocol. It may not have been registered.
)

@echo.
echo Protocol deregistration process completed.
pause