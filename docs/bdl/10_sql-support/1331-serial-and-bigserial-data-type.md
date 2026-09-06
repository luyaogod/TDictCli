---
title: "SERIAL and BIGSERIAL data type"
source: "fgl-topics/c_fgl_odiagmys_009.html"
breadcrumb: "SQL support > SQL database guides > Oracle® MySQL / MariaDB > Data dictionary > SERIAL and BIGSERIAL data type"
type: "concept"
description: "Informix® Informix supports the SERIAL , BIGSERIAL data types to produce automatic integer sequences: SERIAL can produce 32 bit integers ( INTEGER ) BIGSERIAL can produce 64 bit integers ( BIGINT ) ..."
---

# SERIAL and BIGSERIAL data type

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

## Oracle® MySQL and MariaDB

MySQL supports the `AUTO_INCREMENT` column definition option as well as the
`SERIAL` keyword:

- In `CREATE TABLE`, you specify a auto-incremented column with the
  `AUTO_INCREMENT` attribute
- Auto-incremented columns have the same behavior as Informix `SERIAL` columns
- A start value can be defined with `ALTER TABLE tabname AUTO_INCREMENT =
  value`
- The column must be part of the primary key, or the first column of an index.
- `SERIAL` is a synonym for `BIGINT UNSIGNED NOT NULL AUTO_INCREMENT
  UNIQUE`.

## Solution

> **Note:**
>
> For best SQL portability when using different types of databases, consider using sequences as
> described in [Solution 3: Use native SEQUENCE database objects](1026-solution-3-use-native-sequence-database-objects.md).

Note that MySQL (8.0) does not support `CREATE SEQUENCE`, while MariaDB (10.3)
supports this feature. However, with MySQL it is possible to create a table dedicated to sequence
generation, by using an `AUTO_INCREMENT` column.

The Informix `SERIAL` data type is
emulated with MySQL `AUTO_INCREMENT` option.

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

## Using the `native` serial emulation (only option)

After an insert, `sqlca.sqlerrd[2]` holds the last generated serial value. You can
also use the `LAST_INSERT_ID()` SQL function to get the last generated sequence.

`AUTO_INCREMENT` columns must be primary keys. This is handled automatically when
you create a table in a BDL program.

Like Informix, MySQL allows you to specify a zero for
auto-incremented columns. However, for SQL portability, it is recommended to review
`INSERT` statements to remove the `SERIAL` column from the list.

For example, the following
statement:

```
INSERT INTO tab (col1,col2) VALUES ( 0, p_value)
```

can be converted to:

```
INSERT INTO tab (col2) VALUES (p_value)
```

Static SQL `INSERT` using records defined from the schema file must also be
reviewed:

```
DEFINE rec LIKE tab.*
INSERT INTO tab VALUES ( rec.* ) -- will use the serial column
```

can be converted to:

```
INSERT INTO tab VALUES rec.* -- without parentheses, serial column is removed
```

## Related links

**Related concepts**  

[Auto-incremented columns (serials)](1023-auto-incremented-columns-serials.md "How to implement automatic record keys.")
