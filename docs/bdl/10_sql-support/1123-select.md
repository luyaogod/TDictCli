---
title: "SELECT"
source: "fgl-topics/c_fgl_static_sql_SELECT.html"
breadcrumb: "SQL support > Static SQL statements > SELECT"
type: "concept"
---

# SELECT

> Produces a result set from a query on database tables.

## Syntax

```
select-statement [ UNION [ALL] select-statement ] [...]
```

where select-statement
is:

```
SELECT [subset-clause] [duplicates-option]
         { * | select-list } 
  [ INTO variable [,...] ]
  FROM table-list [,...]
  [ WHERE condition ]
  [ GROUP BY column-list [ HAVING condition ] ]
  [ ORDER BY column [ASC|DESC] [,...] ]
```

where subset-clause is:

```
[ SKIP { integer | variable }]
[ {FIRST|MIDDLE|LIMIT} { integer | variable ]
```

where duplicates-option
is:

```
{ ALL
| DISTINCT
| UNIQUE
}
```

where select-list
is:

```
{ [@]table-specification.*
| [table-specification.]column
| literal
} [ [AS] column-alias ]
[,...]
```

where table-list
is:

```
{ table-name
| OUTER table-name
| OUTER ( table-name [,...] )
}
[,...]
```

where table-name is:

```
table-specification [ [AS] table-alias]
```

where table-specification is:

```
[dbname[@dbserver]:][owner.]table
```

where column-list is:

```
column-name [,...]
```

where column-name is:

```
[table.]column
```

1. dbname identifies the database name.
2. dbserver identifies the database server (INFORMIXSERVER).
3. owner identifies the owner of the table, with
   optional double quotes.
4. table is the name of the database table.
5. table-alias defines a new name to reference
   the table in the rest of the statement.
6. integer is an integer constant.
7. variable is a program variable.
8. column is a name of a table column.
9. column-alias defines a new name to reference
   the column in the rest of the statement.
10. condition is an SQL expression to select the
    rows to be deleted.

## Usage

It is recommended to avoid using the dbname,
dbserver and owner prefix of the table name to
ensure maximum SQL portability.

If the `SELECT` statement returns only
one row of data, you can write it directly as a procedural
instruction. However, you must use the `INTO` clause
to provide the list of variables where column values will
be fetched. The `INTO` clause provides the
list of fetch buffers. This clause is not part of the SQL
language sent to the database server; it is extracted from the statement
by the compiler.

```
MAIN
   DEFINE myrec RECORD
            key INTEGER,
            name CHAR(10),
            cdate DATE,
            comment VARCHAR(50)
         END RECORD
   DATABASE stock 
   LET myrec.key = 123
   SELECT name, cdate 
     INTO myrec.name, myrec.cdate
     FROM items 
     WHERE key=myrec.key
END MAIN
```

If the `SELECT` statement
returns more than one row of data, you must declare
a database cursor to process the result set.

```
MAIN
   DEFINE myrec RECORD
            key INTEGER,
            name CHAR(10),
            cdate DATE,
            comment VARCHAR(50)
         END RECORD
   DATABASE stock
   LET myrec.key = 123
   DECLARE c1 CURSOR FOR
     SELECT name, cdate
       FROM items
       WHERE key=myrec.key 
   OPEN c1
   FETCH c1 INTO myrec.name, myrec.cdate 
   CLOSE c1
END MAIN
```

The `SELECT` statement
can include the `INTO` clause, but it is strongly
recommended that you use that clause in the `FETCH` instruction
only.

The `SELECT INTO TEMP` statement creates temporary tables. Such a statement does
not return a result set.

## Related links

**Related concepts**  

[Result set processing](1148-result-set-processing.md "Shows how to fetch rows from a database query.")
