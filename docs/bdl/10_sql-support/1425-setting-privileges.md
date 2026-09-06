---
title: "Setting privileges"
source: "fgl-topics/c_fgl_odiagpgs_022.html"
breadcrumb: "SQL support > SQL database guides > PostgreSQL > Database concepts > Setting privileges"
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

## PostgreSQL

PostgreSQL supports the concept of roles to grant or revoke permissions to a group
of users.

See PostgreSQL documentation for more details.

## Solution

Informix and PostgreSQL
user privileges management are quite similar.
