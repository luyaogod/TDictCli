---
title: "Database users and security"
source: "fgl-topics/c_fgl_sql_programming_058.html"
breadcrumb: "SQL support > SQL programming > SQL portability > Database users and security"
type: "concept"
---

# Database users and security

> Properly identifying database users allows to use database security and audit features.

To get the benefit of the database server security features, you should identify each
physical user as a database user.

Some applications use a single database user for different end users, to avoid user management
and connection issues in the database. This is not good practice, because all user-related features
of the database are unusable. Furthermore, the single db user often has all database privileges and
thus can lead in security issues.

According to the type of server, you must do this steps to create
a database user:

1. Define the user as an operating system user.
2. Declare the user in the database server.
3. Grant database access privileges.

Each database server has its specific users management and data access privilege mechanisms.
Check the vendor documentation for security features and make sure you can define the users,
groups, and privileges in all database servers you want to use.

Programs can identify the current user by executing an SQL query using a keyword or SQL function
that returns the database user name. In most database brands, the keyword is `USER`.
For example, with Oracle DB:

```
SELECT USER INTO p_username FROM DUAL
```

| Database Server Type | SQL Keyword / Function | DB Users topic |
| --- | --- | --- |
| IBM® Informix® | `USER` / `CURRENT_USER` | [DB Users in IBM Informix](1186-what-are-the-supported-ibm-informix-sql-features.md) |
| Microsoft™ SQL Server | `CURRENT_USER` | [DB Users in SQL Server](1268-database-users.md) |
| Oracle® MySQL / MariadDB | `CURRENT_USER` / `CURRENT_USER()` | [DB Users in Oracle MySQL](1323-database-users.md) |
| Oracle Database Server | `USER` | [DB Users in Oracle DB](1368-database-users.md) |
| PostgreSQL | `USER` / `CURRENT_USER` | [DB Users in PostgreSQL](1424-database-users.md) |
| SQLite | N/A | [DB Users in SQLite](1475-database-users.md) |
| Dameng® | `USER` | [DB Users in Dameng](1223-database-users.md) |

## Related links

**Related concepts**  

[Database user authentication](1090-database-user-authentication.md "Different database user authentication methods exist.")
