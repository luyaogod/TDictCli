---
title: "Setting privileges"
source: "fgl-topics/c_fgl_odiagora_023.html"
breadcrumb: "SQL support > SQL database guides > Oracle® Database > Database concepts > Setting privileges"
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

## ORACLE

ORACLE supports the concept of roles to group privileges which then can be assigned
to users.

ORACLE users do not have to explicitly set a role, they are assigned to a default privilege
domain (set of roles). More than one role can be enabled at a time with ORACLE.

Informix database privileges do NOT correspond exactly
to ORACLE `CONNECT`, `RESOURCE` and DBA roles. However, roles can be
created with equivalent privileges.

ORACLE users must have at least the `CREATE SESSION` privilege to access the
database. This privilege is part of the `CONNECT`
role:

```
GRANT CONNECT TO (PUBLIC|username)
```

## Solution

Create a role which groups Informix
`CONNECT` privileges, and assign this role to the application
users:

```
CREATE ROLE ifx_connect IDENTIFIED BY oracle;
GRANT CREATE SESSION, ALTER SESSION, CREATE ANY VIEW, ... TO ifx_connect; 
GRANT ifx_connect TO user1;
```
