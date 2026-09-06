@echo off
chcp 65001 >nul
rem Build the TDict Windows portable package: exe + config + db + readme, zipped.
setlocal
cd /d "%~dp0"

set STAGE=dist\tdict-portable
set GOPROXY=https://goproxy.cn,direct

if exist dist rmdir /s /q dist
mkdir "%STAGE%"

echo [1/4] Building tdict.exe ...
call go build -trimpath -o tdict.exe .
if errorlevel 1 (
    echo BUILD FAILED
    exit /b 1
)

echo [2/4] Copying config.json / erp_data.db ...
copy /y config.json "%STAGE%\" >nul || (echo COPY config.json FAILED & exit /b 1)
copy /y erp_data.db "%STAGE%\" >nul || (echo COPY erp_data.db FAILED & exit /b 1)
if exist "%STAGE%\tdict.exe" del /q "%STAGE%\tdict.exe"
copy /y tdict.exe "%STAGE%\" >nul || (echo COPY tdict.exe FAILED & exit /b 1)

echo [3/4] Packing zip ...
python -c "import zipfile, os; root='dist/tdict-portable'; z=zipfile.ZipFile('dist/tdict-portable.zip','w',zipfile.ZIP_DEFLATED); [z.write(os.path.join(root,n),'tdict-portable/'+n) for n in os.listdir(root)]; z.close()" >nul 2>nul
if errorlevel 1 (
    echo   python unavailable, falling back to PowerShell ...
    powershell -NoProfile -Command "Compress-Archive -Path '%STAGE%' -DestinationPath 'dist\tdict-portable.zip' -Force" >nul 2>nul
    if errorlevel 1 (
        echo ZIP FAILED
        exit /b 1
    )
)

echo [4/4] Done: dist\tdict-portable.zip
dir /b dist
endlocal
