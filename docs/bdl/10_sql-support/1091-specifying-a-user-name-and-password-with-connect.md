---
title: "Specifying a user name and password with CONNECT"
source: "fgl-topics/c_fgl_Connections_019.html"
breadcrumb: "SQL support > Database connections > Database user authentication > Specifying a user name and password with CONNECT"
type: "concept"
description: "In order to specify a user name and password, use the CONNECT instruction with the USER / USING clause: MAIN DEFINE uname, upswd STRING CALL login_dialog() RETURNING uname, upswd CONNECT TO \"stock\" ..."
---

# Specifying a user name and password with CONNECT

In order to specify a user name and password, use the [`CONNECT`](1099-connect-to.md "Opens a new database session in multi-session mode.") instruction with the `USER`/`USING`
clause:

```
MAIN
  DEFINE uname, upswd STRING
  CALL login_dialog() RETURNING uname, upswd 
  CONNECT TO "stock" USER uname USING upswd
  ...
END MAIN
```

This is the recommended way to connect to a database server.

With some database types, it is possible to use an external user authentication service,
such as Kerberos, SSL/TLS, LDAP-based directory services. To connect as an external user,
configure database client settings to authenticate the external user and perform the
CONNECT TO instruction without specifying a login/password:

```
CONNECT TO "stock"
```

For more details, see for example [database user
handling in the Oracle SQL Adaptation Guide](1368-database-users.md).
