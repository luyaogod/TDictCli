---
title: "INSERT"
source: "fgl-topics/c_fgl_static_sql_INSERT.html"
breadcrumb: "SQL support > Static SQL statements > INSERT"
type: "concept"
---

# INSERT

> Creates a new row in a database table.

## Syntax 1:

This is the most standard syntax, works with all types of database
engines.

```
INSERT INTO table-specification [ ( column [,...] ) ]
{
 VALUES ( { variable | sql-expression } [,...] )
|
 select-statement
}
```

## Syntax 2:

The fglcomp compiler will automatically generate a standard INSERT statement
with the complete list of members of the record. The generated SQL will depend on the definition of
the
record.

```
INSERT INTO table-specification VALUES ( record.* )
```

## Syntax 3:

This syntax requires a database schema specification with the `SCHEMA`
instruction, and the corresponding database schema file. With this syntax, the columns of
`record.*` are expanded, without the `SERIAL` or
`BIGSERIAL`
column:

```
INSERT INTO table-specification VALUES record.*
```

where table-specification is:

```
[dbname[@dbserver]:][owner.]table
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
8. select-statement is a static `SELECT` statement with or
   without parameters as variables.
9. record is the name of a record (followed by dot star in this syntax).

## Usage

The `INSERT` SQL statement can be used to create a row in a specified database
table.

The dbname, dbserver and owner prefix
of the table name should be avoided for maximum SQL portability.

With `INSERT INTO table VALUES ...`, the statement inserts a
row in the table with the values specified in [variables](../08_language-basics/0686-variables.md "Explains how to define program variables."), as [literals](../08_language-basics/0585-literals.md "Describes the syntax of literals (constant values) to be used in sources."), or with
`NULL`. When using a [record](../08_language-basics/0715-records.md "Records allow structured program variables definitions."), you can
specify all record members with the dot-star notation (`record.*`).

The third syntax can be used to avoid serial column usage in the value list. The record member
corresponding to a column defined as `SERIAL` or `BIGSERIAL` in the
[schema file](../09_advanced-features/0791-database-schema.md "Defines database table structures with column type information to be reused in program variable definitions.") will be removed by the compiler. This
is useful when using databases like Microsoft™ SQL Server,
where `IDENTITY` columns must be omitted in `INSERT` statements.

When using a select-statement, the statement inserts all rows returned in
the result set of the `SELECT` statement. The columns returned by the result set must
match the column number and data types of the target table. For SQL portability, it is not
recommended that you use this syntax.

## Example

Example using `INSERT` syntax #2:

```
MAIN
   DEFINE myrec RECORD
            key INTEGER,
            name VARCHAR(50),
            cdate DATE,
            comment VARCHAR(250)
         END RECORD
   DATABASE stock 
   LET myrec.key     = 123
   LET myrec.name    = "Katos"
   LET myrec.cdate   = TODAY
   LET myrec.comment = "Perfect jumpsuit, green color"
   INSERT INTO goods VALUES ( myrec.* )
END MAIN
```

Example using database schema and `INSERT` syntax #3 to avoid serial column usage
in the SQL parameters (assuming `customer.customer_num` is a `SERIAL`
column):

```
SCHEMA stores

MAIN
   DEFINE myrec RECORD LIKE customer.*
   DATABASE stores
   LET myrec.customer_num     = NULL
   LET myrec.fname            = "Robert"
   LET myrec.lname            = "Parrisan"
   INSERT INTO customer VALUES myrec.*  -- without serial column 
   LET myrec.customer_num = sqlca.sqlerrd[2] -- new serial value
END MAIN
```

## Related links

**Related concepts**  

[Dynamic SQL management](1140-dynamic-sql-management.md "Explains how to execute and manage SQL statements at runtime.")
