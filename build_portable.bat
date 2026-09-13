@echo off
chcp 65001 >nul
rem Build the TDict Windows portable package: exe + EMPTY config + example + readme, zipped.
rem NOTE: 绝不打包本机 config.json(含真实 SSH/数据库凭据);便携版落地的是 config.empty.json,
rem       用户首次用 `tdict serve` 或手改自行配置。
rem NOTE: 也不打包 erp_data.db(含客户表字典/schema/企业码等数据);用户配好环境后自行
rem       `tdict db sync` 拉取本地字典库。
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

echo [2/4] Staging EMPTY config.json / config.example.json / README.md ...
copy /y config.empty.json "%STAGE%\config.json" >nul || (echo COPY empty config FAILED & exit /b 1)
copy /y config.example.json "%STAGE%\" >nul || (echo COPY config.example.json FAILED & exit /b 1)
copy /y README.md "%STAGE%\" >nul || (echo COPY README.md FAILED & exit /b 1)
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
