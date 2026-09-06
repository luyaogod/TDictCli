---
title: "Database users"
source: "fgl-topics/c_fgl_odiagora_022.html"
breadcrumb: "SQL support > SQL database guides > Oracle® Database > Database concepts > Database users"
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

## ORACLE

Oracle® users can be authenticated in
different ways: as database users, as operating system users or by delegating authentication to
another service, like Kerberos or LDAP.

Users must be created in the database with a `CREATE USER` command, to create a
user authenticated by the database
server:

```
CREATE USER username IDENTIFIED BY password
```

Oracle users can also be created with
the `IDENTIFIED EXTERNALLY`
clause:

```
CREATE USER username IDENTIFIED EXTERNALLY
```

In this case, Oracle trusts the
operating system to authenticate the user. See the Oracle documentation for OS user authentication configuration, especially
the `OS_AUTHENT_PREFIX` (empty string) and `REMOTE_OS_AUTHENT` (true)
server parameters. Note also that the Oracle user name needs to be specified in uppercase in the CREATE USER instruction, and gets
an additional prefix, depending on the operating system (domain name on Windows® platforms)

In Oracle, users can also be defined in
a central LDAP directory, with the `IDENTIFIED GLOBALLY`
clause:

```
CREATE USER username IDENTIFIED GLOBALLY AS 'distinguished_name'
```

Global users are registered and managed by an external LDAP service, and are identified by the
distinguished name (DN).

Oracle supports also proxy
authentication, by granting connection privileges through another "proxy
user":

```
ALTER USER username GRANT CONNECT THROUGH proxy_user
```

## Solution

Based on the application logic, you must create one or several Oracle users. Use database or external authentication.

## Identifying the current database user

To get the database user associated to the current SQL connection with Oracle, execute the following SQL
statement:

```
DEFINE p_username VARCHAR(50)
SELECT USER INTO p_username FROM DUAL
```

## Switching to the application schema

If several DB users are defined for the application, you might want to switch to a common schema
with the following FGLPROFILE
entry:

```
dbi.database.mydb.ora.schema = "app_owner"
```

## Connecting as SYSDBA or SYSOPER

An Oracle connection can also be
established as `SYSDBA` or `SYSOPER` users. This is possible by
specifying the `/SYSDBA` or `/SYSOPER` strings after the user name in
the `USER` clause of the `CONNECT TO` instruction.

For example:

```
CONNECT TO "orc1fox+driver='dbmora'"
    USER "orauser/SYSDBA" USING "fourjs"
```

## Using proxy authentication

If Oracle proxy authentication is
required, specify the `/PROXY_CLIENT:username` string after the
user name in the `USER` clause of the `CONNECT TO` instruction.

For example:

```
CONNECT TO "orc1fox+driver='dbmora'"
    USER "orauser/PROXY_CLIENT:appuser" USING "fourjs"
```

In the above example, the credentials of the orauser login will be used to
establish the connection, and then Oracle
will automatically switch to the user "appuser", assuming that the proxy
connection has been granted
with:

```
ALTER USER appuser GRANT CONNECT THROUGH orauser
```

## Related links

**Related concepts**  

[Database users and security](1002-database-users-and-security.md "Properly identifying database users allows to use database security and audit features.")

[Oracle DB specific FGLPROFILE parameters](1081-oracle-db-specific-fglprofile-parameters.md "Oracle DB specific FGLPROFILE parameters")
