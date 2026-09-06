---
title: "Database user authentication"
source: "fgl-topics/c_fgl_Connections_018.html"
breadcrumb: "SQL support > Database connections > Database user authentication"
type: "concept"
---

# Database user authentication

> Different database user authentication methods exist.

Connecting to a database server is not just specifying a database name: the current user must be
identified by the database server. Database users must be declared in the database
server and must be authenticated.

The typical user authentication is done by passing a login name and password at
connection time. Some database servers support external authentication methods, that do
not require login/password information (for example when db users are based on operating
system users), as well as delegated user authentication via credential tokens (for
example, when using an LDAP distinguished name). See database vendor specific
documentation for more details.

Additional user authentication solutions are provided to simplify migration from IBM®
Informix®databases, but are not recommended in
production for security reasons.

See also SQL adaptation guides for database vendor specific notes regarding user
authentication.

## Child topics

- [Specifying a user name and password with CONNECT](1091-specifying-a-user-name-and-password-with-connect.md)
- [Specifying a user name and password with DATABASE](1092-specifying-a-user-name-and-password-with-database.md)
- [User authentication callback function](1093-user-authentication-callback-function.md)
- [Order of precedence for database user specification](1094-order-of-precedence-for-database-user-specification.md)
