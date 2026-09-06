---
title: "SQL Server (MS ODBC) specific FGLPROFILE parameters"
source: "fgl-topics/c_fgl_dbvendor_param_snc.html"
breadcrumb: "SQL support > Database connections > Database type specific parameters in FGLPROFILE > SQL Server (MS ODBC) specific FGLPROFILE parameters"
type: "concept"
description: "dbi.database. dsname .snc.logintime Connection timeout (in seconds). dbi.database.stores.snc.logintime = 5 Set this parameter to raise an SQL error if the connection can not be established after the ..."
---

# SQL Server (MS ODBC) specific FGLPROFILE parameters

## `dbi.database.dsname.snc.logintime`

Connection timeout (in seconds).

```
dbi.database.stores.snc.logintime = 5
```

Set this parameter to raise an SQL error if the connection can not be
established after the given number of seconds.

The default is 5 seconds.

## `dbi.database.dsname.snc.prefetch.rows`

Maximum number of rows to be pre-fetched.

```
dbi.database.stores.snc.prefetch.rows = 50
```

Use this parameter to increase performance by defining the maximum
number of rows to be fetched into the db client buffer. However, the
bigger this parameter is, the more memory is used by each program.

The default is 10 rows.

## `dbi.database.dsname.snc.widechar`

Control single-char / wide-char mode usage for character string data.

Only set the `snc.widechar` parameter to `false`, if you use
char/varchar/text columns in the database, and the application locale is a multibyte character set
(such as BIG5).

```
dbi.database.stores.snc.widechar = false
```

By default, the SNC driver will select the expected char mode, depending on the current
application locale (LANG/LC\_ALL):

- In wide char mode, the SNC driver uses SQLWCHAR ODBC API functions, by converting the character
  data from the current locale to UCS/2, it will add the N prefix before string literals, and will
  bind SQL parameters with SQL\_C\_WCHAR and SQL\_WCHAR/SQL\_WVARCHAR types.
- In single char mode, the SNC driver will pass the character strings without conversion to
  "ASCII" ODBC API functions, it will leave the string literals (without adding the N prefix) and bind
  character string parameters with SQL\_C\_CHAR and SQL\_CHAR/SQL\_VARCHAR SQL types.

For more details, see also [SQL Server Adaptation Guide:
CHARACTER data types](1273-char-and-varchar-data-types.md).
