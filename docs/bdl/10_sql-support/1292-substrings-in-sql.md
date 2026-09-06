---
title: "Substrings in SQL"
source: "fgl-topics/c_fgl_odiagmsv_024.html"
breadcrumb: "SQL support > SQL database guides > Microsoft™ SQL Server > Data manipulation > Substrings in SQL"
type: "concept"
description: "Informix® Informix SQL statements can use subscripts on columns defined with the character data type: SELECT ... FROM tab1 WHERE col1[2,3] = 'RO' SELECT ... FROM tab1 WHERE col1[10] = 'R' -- Same as ..."
---

# Substrings in SQL

## Informix®

Informix SQL statements can use subscripts on columns
defined with the character data type:

```
SELECT ... FROM tab1 WHERE col1[2,3] = 'RO'
SELECT ... FROM tab1 WHERE col1[10] = 'R' -- Same as col1[10,10]
UPDATE tab1 SET col1[2,3] = 'RO' WHERE ...
SELECT ... FROM tab1 ORDER BY col1[1,3]
```

> **Important:**
>
> With other database servers than Informix, when the subscript notation is used to modify column values in
> `UPDATE` statement, or as `ORDER BY` element, you will get and SQL
> error:
>
> ```
> UPDATE tab1 SET col1[2,3] = 'RO' WHERE ...
> SELECT ... FROM tab1 ORDER BY col1[1,3]
> ```

Informix SQL provides a various set of substring functions:

- `SUBSTR( str-expr, start-pos [,
  length ] )`
- `SUBSTRB( str-expr, start-pos [,
  length ] )`
- `SUBSTRING( str-expr FROM start-pos [FOR
  length ] )`
- `SUBSTRING_INDEX( str-expr, delimiter,
  count )`

## Microsoft™ SQL Server

Microsoft SQL Server provides the
`SUBSTR()` function, to extract a substring from a string expression:

```
SELECT .... FROM tab1 WHERE SUBSTRING(col1,2,2) = 'RO'
SELECT SUBSTRING('Some text',6,3) FROM tab1 -- Gives 'tex'
```

## Solution

Replace all Informix `col[x,y]`
right-value expressions with `SUBSTRING(col,x,y-x+1)`.

Rewrite `UPDATE` and `ORDER BY` clauses using
`col[x,y]` expressions.

The Informix substring functions
`SUBSTR()`, `SUBSTRB()`, `SUBSTRING()` and
`SUBSTRING_INDEX()` are not converted by the ODI driver.

The translation of
`col[x,y]` expressions can be controlled with the following FGLPROFILE
entry:

```
dbi.database.dsname.ifxemul.colsubs = { true | false }
```

For more details see [IBM Informix emulation parameters in FGLPROFILE](1079-ibm-informix-emulation-parameters-in-fglprofile.md "Emulation of Informix specific SQL features can be controlled with FGLPROFILE entries.").

## Related links

**Related concepts**  

[Substring expressions](1041-substring-expressions.md "Handle substrings expressions with different database engines.")
