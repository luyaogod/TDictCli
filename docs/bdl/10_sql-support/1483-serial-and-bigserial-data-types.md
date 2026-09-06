---
title: "SERIAL and BIGSERIAL data types"
source: "fgl-topics/c_fgl_odiagsqt_007.html"
breadcrumb: "SQL support > SQL database guides > SQLite > Data dictionary > SERIAL and BIGSERIAL data types"
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

## SQLite

SQLite supports the `AUTOINCREMENT` attribute for columns:

- Only one column must be declared as `INTEGER PRIMARY KEY AUTOINCREMENT`.
- SQLite (version 3.8.3.1) does not support `SEQUENCE` objects.
- SQLite (version 3.8.3.1) does not allow `AUTOINCREMENT` on
  `BIGINT` columns.
- To get the last generated number, SQLite provides the sqlite\_sequence table:

  ```
  SELECT seq FROM sqlite_sequence WHERE name='tabname'
  ```
- At `INSERT`, for the auto-incremented column:
  - When specifying a zero, SQLite will not generate a new sequence like Informix does.
  - When specifying a `NULL`, SQLite generates a new sequence; Informix denies nulls in serials.
  - When specifying a value different from zero and `NULL`, SQLite will use that
    value. The next `INSERT` statement not providing a value > 0 will produce a new
    auto-incremented value that is greater than the last inserted value.

## Solution

> **Note:**
>
> For best SQL portability when using different types of databases, consider using sequences as
> described in [Solution 3: Use native SEQUENCE database objects](1026-solution-3-use-native-sequence-database-objects.md).

When using SQLite, the `SERIAL` data type is converted to `INTEGER PRIMARY
KEY AUTOINCREMENT`.

The serial type emulation can be enabled or disabled with the
following FGLPROFILE
entries:

```
dbi.database.dbname.ifxemul.datatype.serial = {true|false}
dbi.database.dbname.ifxemul.datatype.serial8 = {true|false}
dbi.database.dbname.ifxemul.datatype.bigserial = {true|false}
```

For more details see [IBM Informix emulation parameters in FGLPROFILE](1079-ibm-informix-emulation-parameters-in-fglprofile.md "Emulation of Informix specific SQL features can be controlled with FGLPROFILE entries.").

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

The `sqlca.sqlerrd[2]` register is filled automatically after each
`INSERT` with the last generated number, by fetching the value from the
`sqlite_sequence` table.

> **Important:**
>
> SQLite (V 3.8) does not support auto-incremented `BIGINT`
> columns. Therefore, `BIGSERIAL` or `SERIAL8` cannot be supported.
> These Informix SQL types are converted to `BIGINT PRIMARY KEY AUTOINCREMENT`, but
> that will produce an SQL error `"AUTOINCREMENT is only allowed on an INTEGER PRIMARY
> KEY"`.

Because SQLite does not behave like Informix regarding
zero and `NULL` value specification for auto-incremented columns, all
`INSERT` statements must be reviewed to remove the `SERIAL` column
from the list.

For example, the following statement:

```
INSERT INTO tab (col1,col2) VALUES ( 0 , p_value)
```

Can be converted to:

```
INSERT INTO tab (col2) VALUES (p_value)
```

Static SQL INSERT using records defined from the schema file must also be reviewed:

```
DEFINE rec LIKE tab.*
INSERT INTO tab VALUES (rec.*) -- will use the serial column
```

Can be converted to:

```
INSERT INTO tab VALUES rec.* -- without parentheses, serial column is removed
```

## Related links

**Related concepts**  

[Auto-incremented columns (serials)](1023-auto-incremented-columns-serials.md "How to implement automatic record keys.")
