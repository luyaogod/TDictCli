---
title: "Database users"
source: "fgl-topics/c_fgl_odiagdmg_019.html"
breadcrumb: "SQL support > SQL database guides > Dameng® database server > Database concepts > Database users"
type: "concept"
description: "Informix® Until version 11.70.xC2, Informix database users must be created at the operating system level and must be members of the 'informix' group. Starting with 11.70.xC2, Informix supports ..."
---

# Database users

## Informix®

Until version 11.70.xC2, Informix database users must
be created at the operating system level and must be members of the 'informix' group.

Starting with 11.70.xC2, Informix supports
database-only users with the `CREATE USER` instruction, as in most other db
servers.

Any database user must have sufficient privileges to connect and use resources of the database;
user rights are defined with the `GRANT` command.

To get the database user associated to the current SQL connection with Informix, execute the following SQL
statement:

```
DEFINE p_username VARCHAR(50)
SELECT USER INTO p_username FROM systables WHERE tabid=1
```

## Dameng®

With Dameng, users must be created in the database server.

The SQL instruction to create a Dameng database user
is:

```
CREATE USER dmguser IDENTIFIED BY "fourjs"
    PASSWORD_POLICY 0 LIMIT PASSWORD_LIFE_TIME UNLIMITED;
```

User permissions can be granted with the GRANT command:

```
GRANT DBA TO dmguser;
GRANT RESOURCE TO dmguser;
```

## Solution

Create a Dameng database user for each Informix database user.

To get the database user associated to the current SQL connection with Dameng, execute the
following SQL statement:

```
DEFINE p_username VARCHAR(50)
SELECT USER INTO p_username FROM dual
```

See Dameng documentation for more details on database user creation and authentication.

## Related links

**Related concepts**  

[Database users and security](1002-database-users-and-security.md "Properly identifying database users allows to use database security and audit features.")
