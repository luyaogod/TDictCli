---
title: "UPDATE"
source: "fgl-topics/c_fgl_static_sql_UPDATE.html"
breadcrumb: "SQL support > Static SQL statements > UPDATE"
type: "concept"
---

# UPDATE

> Modifies rows of a database table.

## Syntax 1:

This is the most standard syntax,
working with all type of database engines.

```
UPDATE table-specification
   SET
       column = { variable | sql-expression }
       [,...] 
   [ sql-condition ]
```

## Syntax 2:

This syntax is not standard, but will be converted by the compiler to a portable UPDATE
syntax.

```
UPDATE table-specification 
   SET ( column [,...] )
     = ( { variable | sql-expression } [,...] )
   [ sql-condition ]
```

## Syntax 3:

This syntax is not portable, and
is not converted by the compiler.

```
UPDATE table-specification 
   SET [table.]*
     = ( { variable | sql-expression } [,...] )
   [ sql-condition ]
```

## Syntax 4:

This syntax requires a database schema specification with `SCHEMA` instruction,
and the corresponding database schema file. With this syntax, the columns of
`record.*` are expanded, without the `SERIAL` or
`BIGSERIAL` column:

```
UPDATE table-specification 
   SET { [table.]* | ( column [,...] ) }
     = record.*
   [ sql-condition ]
```

where table-specification
is:

```
[dbname[@dbserver]:][owner.]table
```

And sql-condition
is:

```
WHERE { condition | CURRENT OF cursor }
```

1. dbname identifies the database name.
2. dbserver identifies the database server (INFORMIXSERVER).
3. owner identifies the owner of the table, with optional double quotes.
4. table is the name of the database table.
5. column is a name of a table column.
6. variable is a program variable, a record member or an array member used as a
   parameter buffer to provide values.
7. sql-expression is an expression supported by the database server, this can be
   a literal or `NULL` for example.
8. record is the name of a record (followed by dot star in this syntax).
9. condition is an SQL expression to select the rows to be updated.
10. cursor is the identifier of a database cursor.

## Usage

The `UPDATE` SQL
statement can be used to modify one or more rows in the specified
database table.

It is recommended to avoid using the dbname,
dbserver and owner prefix of the table name to
ensure maximum SQL portability.

The third syntax is not standard and will not work with all database types. It is not
recommended.

The fourth syntax can be used if the [database schema
file](../09_advanced-features/0791-database-schema.md "Defines database table structures with column type information to be reused in program variable definitions.") has been generated with the correct data types. This is especially important when using
`SERIAL` columns or equivalent auto-incremented columns. The
fglcomp compiler will automatically extend the SQL text with the columns
identified by the record variable. The columns defined in the database schema file as SERIAL or
BIGSERIAL will be omitted in the generated column list.

*column* with
a subscript expression (`column[a,b]`) is not recommended
because most database servers do not support this notation.

For more details about the `WHERE CURRENT OF` clause,
see [Positioned updates/deletes](1156-positioned-updates-deletes.md "Describes row modification based on a FOR UPDATE cursor.").

## Example

```
MAIN
   DEFINE myrec RECORD
            key INTEGER,
            name CHAR(10),
            cdate DATE,
            comment VARCHAR(50)
         END RECORD
   DATABASE stock 
   LET myrec.key     = 123
   LET myrec.name    = "Katos"
   LET myrec.cdate   = TODAY
   LET myrec.comment = "xxxxxx"
   UPDATE items SET
        name    = myrec.name,
        cdate   = myrec.cdate,
        comment = myrec.comment 
    WHERE key = myrec.key 
END MAIN
```

## Related links

**Related concepts**  

[Positioned updates/deletes](1156-positioned-updates-deletes.md "Describes row modification based on a FOR UPDATE cursor.")
