---
title: "The LENGTH() function"
source: "fgl-topics/c_fgl_odiagdmg_042.html"
breadcrumb: "SQL support > SQL database guides > Dameng® database server > Data manipulation > The LENGTH() function"
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

## Dameng®

Dameng supports the `LENGTH()` function,
but there are some differences with Informix
`LENGTH()`.

The Dameng `LENGTH()` function counts
trailing blanks. When using a `CHAR` column, values are blank padded, and the function
returns the size of the `CHAR` column. When using a `VARCHAR` column,
trailing blanks are significant, and the function returns the number of characters, including
trailing blanks.

## Solution

Check if the trailing blanks are significant when using the `LENGTH()` SQL function
in your application.

To count the number of characters by ignoring the trailing blanks, you must use the
`RTRIM()` function:

```
SELECT LENGTH(RTRIM(col1)) FROM table
```

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
