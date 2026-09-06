---
title: "SERIAL and BIGSERIAL data types"
source: "fgl-topics/c_fgl_odiagdmg_008.html"
breadcrumb: "SQL support > SQL database guides > Dameng® database server > Data dictionary > SERIAL and BIGSERIAL data types"
type: "concept"
description: "Informix® Informix supports the SERIAL , BIGSERIAL data types to produce automatic integer sequences: SERIAL can produce 32 bit integers ( INTEGER ) BIGSERIAL can produce 64 bit integers ( BIGINT ) ..."
---

# SERIAL and BIGSERIAL data types

## Informix®

Informix supports the `SERIAL`,
`BIGSERIAL` data types to produce automatic integer sequences:

- `SERIAL` can produce 32 bit integers (`INTEGER`)
- `BIGSERIAL` can produce 64 bit integers (`BIGINT`)
- `SERIAL8` is a synonym for `BIGSERIAL`

Steps to use serials with Informix:

1. Create the table with a column using `SERIAL`, or
   `BIGSERIAL`.
2. To generate a new serial, no value or a zero value is specified in the `INSERT`
   statement:

   ```
   INSERT INTO tab1 ( c ) VALUES ( 'aa' )
   INSERT INTO tab1 ( k, c ) VALUES ( 0, 'aa' )
   ```
3. After `INSERT`, the new value of a `SERIAL` or
   `BIGSERIAL` column is provided in `sqlca.sqlerrd[2]`, or it can be
   fetched with a `SELECT dbinfo('sqlca.sqlerrd1')` query. Since
   `sqlca.sqlerrd[1-6]` is defined as `BIGINT`, when using Informix CSDK
   15 with the `dbmifx_15` driver, the last generated `BIGSERIAL` will be
   available in `sqlca.sqlerrd[2]`.

Informix allows you to insert rows with a value
different from zero for a serial column. Using an explicit value will automatically increment the
internal serial counter, to avoid conflicts with future `INSERT` statements that are
using a zero value:

```
SQL statement                                               Internal serial counter
------------------------------------------------------------------------------------------
CREATE TABLE tab ( pkey SERIAL, name VARCHAR(50) );                    0
INSERT INTO tab VALUES (  0, 'aaa' );                                  1
INSERT INTO tab VALUES ( 10, 'bbb' );                                 10
INSERT INTO tab VALUES (  0, 'ccc' );                                 11
DELETE FROM tab;                                                      11
INSERT INTO tab VALUES (  0, 'ddd' );                                 12
```

## Dameng®

Dameng supports `IDENTITY(start,incr)`
columns:

```
CREATE TABLE tab ( k INTEGER IDENTITY(100,1), name VARCHAR(50) )
```

To get the last generated idetity value after an `INSERT`, Dameng provides the
`IDENTITY_CURRENT('tabname')` function.

To put a specific value into an `IDENTITY` column, the `SET
IDENTITY_INSERT` command must be used:

```
SET IDENTITY_INSERT tab1 ON
INSERT INTO tab1 ( k, c ) VALUES ( 100, 'aaa' )
SET IDENTITY_INSERT tab1 OFF
```

On the next `INSERT` not specifying the identity column, the new auto-incremented
value will be adapted to be greater than the largest value inserted by previous
`INSERT` statements.

Dameng supports
`SEQUENCES`:

```
CREATE SEQUENCE sq1 START WITH 100
```

To create a new sequence number, use the `seqname.nextval()`
expression:

```
INSERT INTO table VALUES ( sq1.nextval(), ... )
```

## Solution

> **Note:**
>
> For best SQL portability when using different types of databases, consider using sequences as
> described in [Solution 3: Use native SEQUENCE database objects](1026-solution-3-use-native-sequence-database-objects.md).

To emulate Informix serials with Dameng, the ODI driver uses `IDENTITY` columns.

In database creation scripts, all `SERIAL[(n)]` data types must be converted by
hand to `INTEGER IDENTITY[(n,1)]` data types, while `BIGSERIAL[(n)]`
data types must be converted by hand to `BIGINT IDENTITY[(n,1)]` data types.

Tables created from the BDL programs can use the `SERIAL` or BIGSERIAL data type:
When a BDL program executes a CREATE [TEMP] TABLE with such column type, the database interface
automatically converts to an identity column.

The new generated `SERIAL` or `BIGSERIAL` value is available from
the `sqlca.sqlerrd[2]` variable. The database interface performs a `SELECT
IDENT_CURRENT('tabname')`. You can also retrieve the last generated
serial with the `IDENT_CURRENT()` SQL function.

By default, you cannot specify the identity column in `INSERT` statements; All
`INSERT` statements must be reviewed to remove the identity column from the list.

For example, the following statement:

```
INSERT INTO tab (col1,col2) VALUES (0, p_value)
```

must be converted to:

```
INSERT INTO tab (col2) VALUES (p_value)
```

Static SQL `INSERT` using records defined from the schema file (`DEFINE rec
LIKE tab.*`) must also be reviewed:

```
INSERT INTO tab VALUES (rec.*) -- will use the serial column
```

must be converted to:

```
INSERT INTO tab VALUES rec.* -- without parentheses, serial column is removed
```

> **Important:**
>
> Do not mix a serial emulation modes: Choose one of the provided serial
> emulations for all the tables of the application database. Do note create the tables with a given
> serial emulation, and later switch to another serial emulation for `INSERT`
> statements. This is not supported.

The serial type emulation can be enabled or disabled with the
following FGLPROFILE
entries:

```
dbi.database.dbname.ifxemul.datatype.serial = {true|false}
dbi.database.dbname.ifxemul.datatype.serial8 = {true|false}
dbi.database.dbname.ifxemul.datatype.bigserial = {true|false}
```

## Disabling automatic serial retrieval for sqlca.sqlerrd[2]

For Informix compatibility, when the
`SERIAL` type emulation is active, the ODI drivers automatically execute another SQL
query (or do a DB client API call when possible) after each `INSERT` statement, to
get the last generated serial, and fill the `sqlca.sqlerrd[2]` register. This results
in some overhead that can be avoided, if the `sqlca.sqlerrd[2]` register is not used
by the program.

When serial emulation is required (to create temp tables with a serial column during program
execution), and the `sqlca.sqlerrd[2]` register does not need to be filled,
(typically because you use your own method to retrieve the last generated serial), you can set the
`ifxemul.datatype.serial.sqlerrd2` FGLPROFILE entry to false. This will avoid the
automatic retrieval of last serial value to fill `sqlca.sqlerrd[2]`:

```
dbi.database.dbname.ifxemul.datatype.serial.sqlerrd2 = false
```

The above FGLPROFILE entry is useless, if Informix `SERIAL` type emulation is
disabled with:

```
dbi.database.dbname.ifxemul.datatype.serial = false
```

See also [fgldbutl.db\_get\_last\_serial()](../15_library-reference/2803-fgldbutl-db-get-last-serial.md "Retrieves the last generated serial for a given serial emulation and database table.").

## Related links

**Related concepts**  

[Auto-incremented columns (serials)](1023-auto-incremented-columns-serials.md "How to implement automatic record keys.")
