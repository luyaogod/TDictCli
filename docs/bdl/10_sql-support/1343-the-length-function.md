---
title: "The LENGTH() function"
source: "fgl-topics/c_fgl_odiagmys_022.html"
breadcrumb: "SQL support > SQL database guides > Oracle® MySQL / MariaDB > Data manipulation > The LENGTH() function"
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

## Oracle® MySQL and MariaDB

MySQL and MariaDB support the `LENGTH()` function, which is similar to Informix `LENGTH()`.

The MySQL/MariaDB `LENGTH()` function ignores trailing blanks by default, except
when `sql_mode` parameter defines `PAD_CHAR_TO_FULL_LENGTH`.

When passing `NULL` as parameter, the MySQL `LENGTH()`
function returns `NULL`.

## Solution

The SQL `LENGTH()` function name can be used with MySQL/MariaDB.

Note that the trailing blanks of `CHAR` values are counted, if the
`sql_mode` parameter defines `PAD_CHAR_TO_FULL_LENGTH`. This affects
`CHAR` columns, not string literals, considered as `VARCHAR` where
trailing blanks are counted.

## Related links

**Related concepts**  

[The LENGTH() function in SQL](1045-the-length-function-in-sql.md "The semantics of the LENGTH() SQL function differs according to the database engine.")
