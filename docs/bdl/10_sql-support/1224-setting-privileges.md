---
title: "Setting privileges"
source: "fgl-topics/c_fgl_odiagdmg_020.html"
breadcrumb: "SQL support > SQL database guides > Dameng® database server > Database concepts > Setting privileges"
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

## Dameng®

Dameng users must have at least the `RESOURCE` permission to access the database
and create tables (for temporary table
emulation):

```
GRANT RESOURCE TO user
```

See Dameng documentation for more details.

## Solution

Informix and Dameng
user privileges management are quite similar.

See also [Temporary Tables](1241-temporary-tables.md)
