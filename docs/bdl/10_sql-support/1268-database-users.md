---
title: "Database users"
source: "fgl-topics/c_fgl_odiagmsv_021.html"
breadcrumb: "SQL support > SQL database guides > Microsoft™ SQL Server > Database concepts > Database users"
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

## Microsoft™ SQL Server

Before a user can access an SQL Server database, the system administrator (SA) must add the
user's login to the SQL Server Login list and add a user name for that database. The user name is a
name that is assigned to a login ID for the purpose of allowing that user to access a specified
database. Database users are members of a user group; the default group is 'public'.

Microsoft SQL Server offers two authentication
modes:

1. The SQL Server authentication mode, which requires a login name and a password
2. The Windows authentication mode, which
   uses the security mechanisms within Windows when validating
   login connections. With this mode, user do not have to enter a login ID and password - their login
   information is taken directly from the network connection.

## Solution

Both SQL Server and Windows
authentication methods can be used to allow BDL program users to connect to
Microsoft SQL Server and access a
specific database.

If you don't specify the `USER`/`USING`
clause in the `CONNECT TO` instruction, operating system
authentication takes place.

To get the database user associated to the current SQL connection with Microsoft SQL Server, execute the following SQL
statement:

```
DEFINE p_username VARCHAR(50)
PREPARE stmt FROM "SELECT CURRENT_USER"
EXECUTE stmt INTO p_username
```

See SQL Server documentation for more details on database logins and users.

## Related links

**Related concepts**  

[Database users and security](1002-database-users-and-security.md "Properly identifying database users allows to use database security and audit features.")
