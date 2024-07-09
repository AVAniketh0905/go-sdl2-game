@echo off
setlocal

REM Configuration
set "sdl_libs_path=C:\mingw64\bin"
set "outputdir=.\builds"
set "name_of_game=YourGameName"
set "info=false"

REM Check for the -i flag
if "%1" == "-i" (
    set "info=true"
)

REM Increment build number
if not exist "%outputdir%" mkdir "%outputdir%"
for /f "tokens=*" %%i in ('dir /B /AD "%outputdir%\build_v*" 2^>nul ^| findstr /R "^build_v[0-9][0-9]*$"') do (
    set "build_dir=%%i"
)
if not defined build_dir (
    set "i=0"
) else (
    set "i=%build_dir:~7%"
)
set /a i+=1

REM Define new build directory
set "newbuild=build_v%i%"
set "newbuild_path=%outputdir%\%newbuild%"
mkdir "%newbuild_path%"

REM Print info if flag is set
if "%info%" == "true" (
    echo Creating new build directory: %newbuild_path%
)

REM Build the Go program
go build -ldflags "-H=windowsgui" -o "%newbuild_path%\%name_of_game%.exe" .

REM Copy SDL DLL files
if "%info%" == "true" (
    echo Copying SDL DLL files from %sdl_libs_path% to %newbuild_path%
)
xcopy "%sdl_libs_path%\*.dll" "%newbuild_path%\" /Y >nul

REM Copy assets directory
if "%info%" == "true" (
    echo Copying assets directory to %newbuild_path%\assets
)
xcopy ".\assets" "%newbuild_path%\assets" /E /I /Y >nul

REM Zip the new build directory
if "%info%" == "true" (
    echo Zipping %newbuild_path% to %outputdir%\%newbuild%.zip
)
powershell -Command "Compress-Archive -Path '%newbuild_path%\*' -DestinationPath '%outputdir%\%newbuild%.zip'"

echo Build %newbuild% created and zipped successfully.

endlocal
pause
