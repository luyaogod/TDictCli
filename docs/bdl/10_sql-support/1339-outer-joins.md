---
title: "Outer joins"
source: "fgl-topics/c_fgl_odiagmys_010.html"
breadcrumb: "SQL support > SQL database guides > Oracle® MySQL / MariaDB > Data manipulation > Outer joins"
type: "concept"
description: "Informix® OUTER() syntax In Informix SQL, outer joins can be defined in the FROM clause with the OUTER keyword: SELECT ... FROM a, OUTER (b) WHERE a.key = b.akey SELECT ... FROM a, OUTER(b,OUTER(c)) ..."
---

# Outer joins

## Informix® OUTER() syntax

In Informix SQL, outer joins can be defined in the
`FROM` clause with the `OUTER`
keyword:

```
SELECT ... FROM a, OUTER (b)
     WHERE a.key = b.akey

SELECT ... FROM a, OUTER(b,OUTER(c))
     WHERE a.key = b.akey 
       AND b.key1 = c.bkey1 AND b.key2 = c.bkey2
```

Informix also supports the ANSI outer join syntax,
which is the recommended way to specify outer joins with recent SQL database
engines:

```
SELECT ... FROM cust LEFT OUTER JOIN order
                      ON cust.key = order.custno
    WHERE ...
```

## Oracle® MySQL and MariaDB

MySQL and MariaDB support the ANSI outer join syntax:

```
SELECT ...
  FROM cust LEFT OUTER JOIN order
                  LEFT OUTER JOIN item
                  ON order.key = item.ordno
             ON cust.key = order.custno
 WHERE order.cdate > current date
```

## Solution

The Genero database drivers can convert Informix Informix `OUTER` specifications to ANSI outer joins.
> **Note:**
>
> For better SQL
> portability, use the ANSI outer join syntax instead of the old Informix `OUTER` syntax.

The outer join translation can be controlled with the following
FGLPROFILE
entry:

```
dbi.database.dsname.ifxemul.outers = { true | false }
```

For more details see [IBM Informix emulation parameters in FGLPROFILE](1079-ibm-informix-emulation-parameters-in-fglprofile.md "Emulation of Informix specific SQL features can be controlled with FGLPROFILE entries.").

- Prerequisites:
  1. In the `FROM` clause, the main table must be the first item and the outer tables
     must be listed from left to right in the order of outer levels.

     Example which does not
     work:

     ```
     ... FROM OUTER(tab2), tab1
     ```
  2. The outer join in the `WHERE` clause must use the table name as prefix:

     ```
     ... WHERE tab1.col1 = tab2.col2
     ```
- Restrictions:
  1. Statements composed by 2 or more `SELECT` instructions are not supported:

     ```
     SELECT ... UNION SELECT ...
     ```

     or:

     ```
     SELECT ... WHERE col IN (SELECT...)
     ```
  2. Additional conditions on outer table columns cannot be detected and therefore are not supported:

     ```
     ... FROM tab1, OUTER(tab2)
         WHERE tab1.col1 = tab2.col2
           AND tab2.colx > 10
     ```
  3. Using subscript in outer conditions:

     ```
     ... FROM tab1, OUTER(tab2)
         WHERE tab1.col1[1,3] = tab2.col2[1,3]
     ```
- Notes:
  1. Table aliases are detected in `OUTER` expressions.

     `OUTER`
     example with table alias:

     ```
     ... OUTER(tab1 alias1) ...
     ```
  2. In the outer join, `outertab.col` can be
     placed on both right or left sides of the equal sign:

     ```
     ... WHERE outertab.col1 = maintab.col2
     ```
  3. Table names detection is not case-sensitive:

     ```
     SELECT ... FROM tab1, TAB2
         WHERE tab1.col1 = tab2.col2
     ```
  4. [Temporary tables](1446-temporary-tables.md) are supported in
     `OUTER` specifications:

     ```
     CREATE TEMP TABLE tt1 ( ... )
     SELECT ... FROM tab1, OUTER(tt1) ...
     ```

## Related links

**Related concepts**  

[Outer joins](1040-outer-joins.md "Use standard ISO outer join syntax instead of the old IBM Informix OUTER() syntax.")
