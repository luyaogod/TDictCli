---
title: "UNIX shell script for private key password"
source: "fgl-topics/c_gws_ssl_configuration_009.html"
breadcrumb: "Web services > Reference > Web services FGLPROFILE configuration > Examples > UNIX™ shell script for private key password"
type: "concept"
description: "UNIX shell script sample returning a password, depending on the .pem file passed as parameter."
---

# UNIX shell script for private key password

> UNIX® shell script sample returning a password, depending on the .pem file passed as parameter.

```
# UNIX password script
if [ "$1" == "Cert/MyPrivateKeyA.pem" ]
 then
  echo PasswordA
fi
if [ "$1" == "Cert/MyPrivateKeyB.pem" ]
 then
  echo PasswordB
fi
```
