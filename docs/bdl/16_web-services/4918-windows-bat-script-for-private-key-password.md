---
title: "Windows BAT script for private key password"
source: "fgl-topics/c_gws_ssl_configuration_008.html"
breadcrumb: "Web services > Reference > Web services FGLPROFILE configuration > Examples > Windows® BAT script for private key password"
type: "concept"
description: "Windows BAT script sample returning a password, depending on the .pem file passed as parameter."
---

# Windows BAT script for private key password

> Windows® BAT script sample returning a password, depending on the .pem file passed as parameter.

```
@echo off
REM -- Windows password script
IF "%1" == "Cert/MyPrivateKeyA.pem" GOTO KeyA
IF "%1" == "Cert/MyPrivateKeyB.pem" GOTO KeyB
GOTO end
:KeyA
ECHO PasswordA
GOTO end
:KeyB
ECHO PasswordB
GOTO end
:end
GOTO :EOF
```
