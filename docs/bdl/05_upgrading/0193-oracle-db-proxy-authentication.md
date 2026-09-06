---
title: "Oracle DB Proxy Authentication"
source: "fgl-topics/c_fgl_Migrate_to_310_oracle_proxyauth.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.10 upgrade guide > Oracle DB Proxy Authentication"
type: "concept"
description: "Specifying a proxy user when connecting to Oracle DB."
---

# Oracle DB Proxy Authentication

> Specifying a proxy user when connecting to Oracle® DB.

Oracle DB supports proxy
authentication, a feature that allows a given DB user "A" to connect with the
credentials of another DB user "B" and to be seen as user "A" for the rest of the SQL
session.

In order to link a proxy user to another user, issue the following
command:

```
ALTER USER app_user GRANT CONNECT THROUGH proxy_user
```

An application can then connect with the credentials of `proxy_user`, and be
identified as `app_user` during the SQL session.

See Oracle DB documentation for more
details about proxy authentication.

Starting with Genero 3.10, it is possible to specify a "proxy client" at connection time, by
using the `/PROXY_CLIENT:user_name` suffix in the user name
parameter of the `CONNECT TO` instruction.

In the example, the application connects with the proxy user credentials, and the proxy client is
defined as
`app_user`:

```
CONNECT TO "myserver"
   USER "proxy_user/PROXY_CLIENT:app_user"
   USING "proxy_pswd"
```

A subsequent `SELECT USER FROM dual` would return `"APP_USER"`.

## Related links

**Related concepts**  

[Database users](../10_sql-support/1368-database-users.md "Database users")
