---
title: "The LENGTH() function"
source: "fgl-topics/c_fgl_odiagora_017.html"
breadcrumb: "SQL support > SQL database guides > Oracle® Database > Data manipulation > The LENGTH() function"
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

## ORACLE

Oracle supports the `LENGTH()` function, but there
are some differences with Informix
`LENGTH()`.

The Oracle `LENGTH()` function counts trailing
blanks. When using a `CHAR` column, values are blank padded, and the function returns
the size of the `CHAR` column. When using a `VARCHAR` column, trailing
blanks are significant, and the function returns the number of characters, including trailing
blanks.

Because ORACLE handles empty strings (`''`) as `NULL`
values, writing `LENGTH('')` is equivalent to `LENGTH(NULL)`. In this
case, the function returns `NULL`.

## Solution

Check if the trailing blanks are significant when using the `LENGTH()` SQL
function in your application.

To count the number of character by ignoring the trailing blanks, use the
`RTRIM()` function:

```
SELECT LENGTH(RTRIM(col1)) FROM table
```

SQL conditions which verify that the result of `LENGTH()` is greater that a given
number do not have to be changed, because the expression evaluates to false if the given string is
empty (`NULL>n`):

```
SELECT * FROM x WHERE LENGTH(col)>0
```

Only SQL conditions that compare the result of `LENGTH()` to zero will not work if
the column is `NULL`. You must check your BDL code for such conditions:

```
SELECT * FROM x WHERE LENGTH(col)=0
```

In this case, you must add a test to verify if the column is null:

```
SELECT * FROM x WHERE ( LENGTH(col)=0 OR col IS NULL )
```

In addition, when retrieving the result of a `LENGTH()` expression into a BDL
variable, you must check that the variable is not `NULL`.

In ORACLE, you can use the `NVL()` function in order to get a non-null value:

```
SELECT * FROM x WHERE NVL(LENGTH(c),0)=0
```

Since Informix supports the `NVL()`
function, you can write the same SQL for both Informix
and ORACLE, as shown in this example.

Consider to create a user-defined function that implements the behavior of the Informix
`LENGTH()`
function:

```
CREATE OR REPLACE FUNCTION vlength(
    value IN VARCHAR2
    )
    RETURN INTEGER
AUTHID CURRENT_USER
IS
BEGIN
    RETURN NVL(LENGTH(RTRIM(value)),0);
END;
/
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
