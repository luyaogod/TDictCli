---
title: "Setting privileges"
source: "fgl-topics/c_fgl_odiagmsv_022.html"
breadcrumb: "SQL support > SQL database guides > Microsoft™ SQL Server > Database concepts > Setting privileges"
type: "concept"
description: "Informix® Informix users must have at least the CONNECT privilege to access the database: GRANT CONNECT TO username Application administration users need the RESOURCE privilege to create tables: GRANT ..."
---

# Setting privileges

## Informix®

Informix users must have at least the `CONNECT` privilege to access the
database:

```
GRANT CONNECT TO username
```

Application
administration users need the `RESOURCE` privilege to create
tables:

```
GRANT RESOURCE TO username
```

Since version 7.20,
Informix supports database
roles:

```
GRANT rolename TO username
```

## Microsoft™ SQL Server

Microsoft SQL Server supports user
groups, to grant or revoke permissions to more than one user at the same time.

See SQL Server documentation for more details.

## Solution

Informix and Microsoft SQL Server user privileges management are quite similar.
