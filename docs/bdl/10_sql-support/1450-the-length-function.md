---
title: "The LENGTH() function"
source: "fgl-topics/c_fgl_odiagpgs_017.html"
breadcrumb: "SQL support > SQL database guides > PostgreSQL > Data manipulation > The LENGTH() function"
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

## PostgreSQL

PostgreSQL supports the `LENGTH()` function,
which is similar to Informix
`LENGTH()`.

The PostgreSQL `LENGTH()` function ignores trailing
blanks.

When passing `NULL` as parameter, the PostgreSQL `LENGTH()`
function returns `NULL`.

## Solution

The SQL `LENGTH()` function name can be used with PostgreSQL.

## Related links

**Related concepts**  

[The LENGTH() function in SQL](1045-the-length-function-in-sql.md "The semantics of the LENGTH() SQL function differs according to the database engine.")
