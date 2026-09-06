---
title: "User authentication callback function"
source: "fgl-topics/c_fgl_Connections_006.html"
breadcrumb: "SQL support > Database connections > Database user authentication > User authentication callback function"
type: "concept"
description: "When using the DATABASE (or CONNECT TO without the USER/USING clause) instruction, the database user name and password can be provided at runtime by a callback function. The callback function must be ..."
---

# User authentication callback function

When using the [`DATABASE`](1096-database.md "Opens a new database connection in unique-session mode.") (or [`CONNECT TO`](1099-connect-to.md "Opens a new database session in multi-session mode.") without the
`USER/USING` clause) instruction, the database user name and password can be provided
at runtime by a callback function. The callback function must be defined with the
`dbi.default.userauth.callback` FGLPROFILE
entry:

```
dbi.default.userauth.callback = "[module-name.]function-name"
```

This callback method is provided when a lot of programs use the `DATABASE`
instruction, and database user credentials are mandatory. If possible, use the `CONNECT
TO` instruction with the `USER`/`USING` clause instead of
`DATABASE`.
> **Note:**
>
> With the IBM® Informix® driver, when using the `DATABASE` instruction, the
> callback method is invoked, but the user name and password returned by the function are ignored:
> Only `CONNECT TO` will take the login parameters into account for IBM Informix, when no
> `USER/USING` clause is specified.

The callback function must have the following signature:

```
CALL function-name(dbspec STRING)
      RETURNS ( STRING, -- username
                STRING  -- password
              )
```

If you do not specify the module name, the callback function must be linked to the 42r program.
By using the "module-name.function-name" syntax in the FGLPROFILE entry, the
runtime system will automatically load the module. In both cases, the module must be located in a
directory where the runtime system can find it, defined by the [FGLLDPATH](../07_configuration/0528-fglldpath.md "Defines a list of paths to find program modules.") environment variable.

In the callback function body, the value of dbspec can be used to identify the
database source, read user name and encrypted password from FGLPROFILE entries with the [`fgl_getResource()`](../15_library-reference/2765-fgl-getresource.md "Returns the value of an FGLPROFILE entry.")
function, then decrypt password with the algorithm of your choice and return user name and decrypted
password.

## User authentication callback function for DATABASE:

```
FUNCTION getUserAuth(dbspec STRING) RETURNS (STRING,STRING)
  DEFINE un, ep STRING
  LET un = fgl_getResource("dbi.database."||dbspec||".username") 
  LET ep = fgl_getResource("dbi.database."||dbspec||".password.encrypted")
  RETURN un, decrypt_user_password(dbspec, un, ep)
END FUNCTION
```
