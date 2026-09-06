---
title: "Database users"
source: "fgl-topics/c_fgl_odiagpgs_021.html"
breadcrumb: "SQL support > SQL database guides > PostgreSQL > Database concepts > Database users"
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

## PostgreSQL

PostgreSQL users must be registered in the database with the `CREATE USER`
command:

```
dbname=# CREATE USER user-name PASSWORD 'password';
CREATE USER
dbname=# GRANT ALL PRIVILEGES ON DATABASE dbname TO user-name;
GRANT
```

## Solution

Based on the application logic (if it is a multiuser application), you have to create one or
several PostgreSQL users.

To get the database user associated to the current SQL connection with PostgreSQL, execute the
following SQL statement:

```
DEFINE p_username VARCHAR(50)
PREPARE stmt FROM "SELECT CURRENT_USER"
EXECUTE stmt INTO p_username
```

## Related links

**Related concepts**  

[Database users and security](1002-database-users-and-security.md "Properly identifying database users allows to use database security and audit features.")
