---
title: "The LENGTH() function"
source: "fgl-topics/c_fgl_odiagmsv_036.html"
breadcrumb: "SQL support > SQL database guides > Microsoft™ SQL Server > Data manipulation > The LENGTH() function"
type: "concept"
description: "Informix® Informix provides the LENGTH() function to count the number of bytes of a character string expression: SELECT LENGTH(\"aaa\"), LENGTH(col1) FROM table Informix LENGTH() does not count the ..."
---

# The LENGTH() function

## Informix®

Informix provides the `LENGTH()`
function to count the number of bytes of a character string expression:

```
SELECT LENGTH("aaa"), LENGTH(col1) FROM table
```

Informix `LENGTH()` does not count the
trailing blanks for `CHAR` or `VARCHAR` expressions, while Oracle
counts the trailing blanks.

Informix `LENGTH()` returns 0 when the
given string is empty. That means, `LENGTH('')=0`.

## Microsoft™ SQL Server

Microsoft SQL Server supports the
`LEN()` function, but there are some differences with Informix `LENGTH()`.

> **Note:**
>
> Do not confuse `LEN()` with `DATALEN()`, which returns the data
> size used for storage (number of bytes).

Like Informix, Microsoft SQL Server ignores trailing blanks when
computing the length of a string.

When passing `NULL` as parameter, the SQL Server `LEN()` function
returns `NULL`.

## Solution

The database driver is able to replace `LENGTH()` by the `LEN()`
function name.

The translation of
`LENGTH()` expressions can be controlled with the following FGLPROFILE
entry:

```
dbi.database.dsname.ifxemul.length = { true | false }
```

For more details see [IBM Informix emulation parameters in FGLPROFILE](1079-ibm-informix-emulation-parameters-in-fglprofile.md "Emulation of Informix specific SQL features can be controlled with FGLPROFILE entries.").

## Related links

**Related concepts**  

[The LENGTH() function in SQL](1045-the-length-function-in-sql.md "The semantics of the LENGTH() SQL function differs according to the database engine.")
