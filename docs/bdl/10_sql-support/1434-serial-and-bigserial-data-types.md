---
title: "SERIAL and BIGSERIAL data types"
source: "fgl-topics/c_fgl_odiagpgs_008.html"
breadcrumb: "SQL support > SQL database guides > PostgreSQL > Data dictionary > SERIAL and BIGSERIAL data types"
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

## PostgreSQL

PostgreSQL provides native `SERIAL` and `BIGSERIAL` data types,
which are similar to Informix `SERIAL` and
`BIGSERIAL` types. However, native PostgreSQL serial types behave differently:

- It is not possible to define a start value directly such as `SERIAL(100)`. To
  modify the next generated serial, use the `setval('sequence',
  value, true/false)` SQL function. The name of the sequence can be found
  with the `pg_get_serial_sequence(tabname,
  colname)` SQL function.
- Unlike Informix, no serial value will be generated when the value zero is specified for the
  serial column during the `INSERT`: The PostgreSQL native serial types are based on
  default values. To get an auto-incremented value, you must omit the serial column in the
  `INSERT` statement.
- When a value is specified for the serial column in an `INSERT`, the underlying
  sequence will not be incremented. As result, by default, the next `INSERT` that does
  not specify the serial column may get a new sequence that was already inserted explicitly. In such
  case, the sequence must be reset with the `setval()` SQL function.

PostgreSQL identity columns:

- Auto-incremented "identity" columns can be defined
  with

  ```
  GENERATED { ALWAYS | BY DEFAULT } AS IDENTITY [ ( sequence-options ) ]
  ```
- A start value can be specified at table creation in the sequence-options with
  `(START WITH start-value)`
- Specifying zero for the identity column during `INSERT` will not generate a new
  serial value. To get a new generated incremented value, the serial column must be omitted in the
  `INSERT` statement.
- Specifying a non-zero value for the identity column during `INSERT` will not
  adapt the underlying sequence. This must be done manually with
  `setval('sequence', value, true/false)`

PostgreSQL sequences:

- The purpose of sequences is to provide unique integer numbers.
- To create a sequence, you must use the `CREATE SEQUENCE`
  statement:

  ```
  CREATE SEQUENCE tab1_seq START WITH 100;
  ```
- To get a new sequence value, you must use the `nextval()`
  function:

  ```
  INSERT INTO tab1 VALUES ( nextval('tab1_seq'), ... )
  ```
- To get the last generated number, PostgreSQL provides the `currval()`
  function:

  ```
  SELECT currval('tab1_seq')
  ```
- To reset a sequence, use the `setval('sequence',
  value, true/false)`
  function:

  ```
  SELECT setval('tab1_seq',200,false);
  ```

See PostgreSQL documentation for more details about serial types and sequences features in this
database engine.

## Solution

> **Note:**
>
> For best SQL portability when using different types of databases, consider using sequences as
> described in [Solution 3: Use native SEQUENCE database objects](1026-solution-3-use-native-sequence-database-objects.md).

The Informix `SERIAL` and
`BIGSERIAL` data types can be emulated with three different methods.

The method used to emulate serial types is defined by the
`ifxemul.datatype.serial.emulation` FGLPROFILE
parameter:

```
dbi.database.dbname.ifxemul.datatype.serial.emulation = {"native"|"regtable"|"trigseq"}
```

- `native`: uses the native PostgreSQL serial data types (`SERIAL`
  or `BIGSERIAL`).
- `regtable`: uses `INSERT` triggers with the
  `SERIALREG` table.
- `trigseq`: uses `INSERT` triggers with sequences.

The default emulation technique is "`native`".

> **Important:**
>
> Do not mix a serial emulation modes: Choose one of the provided serial
> emulations for all the tables of the application database. Do note create the tables with a given
> serial emulation, and later switch to another serial emulation for `INSERT`
> statements. This is not supported.

None of the provided methods use PostgreSQL `GENERATED ... AS IDENTITY` columns:
This feature does not bring sufficient added value compared to PostgreSQL native
`SERIAL` and `BIGSERIAL` types.

For serial emulations using PostgreSQL sequences ("`native`" and
"`trigseq`"), if the table with serial column is owned by another user, the current
user must have sufficient privileges to use, query and modify the underlying sequence. To grant
globally permissions on serial sequences in a given schema, use the following SQL
commands:

```
-- For existing sequences
GRANT USAGE, SELECT, UPDATE
   ON ALL SEQUENCES IN SCHEMA public
      TO user-name;
-- For future created sequences (if needed)
ALTER DEFAULT PRIVILEGES IN SCHEMA public
   GRANT USAGE, SELECT, UPDATE ON SEQUENCES
      TO user-name;
```

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

## Using the `native` serial emulation

The "`native`" mode is the default serial emulation mode, using the native
PostgreSQL `SERIAL` or `BIGSERIAL` data type. In this mode, the
original type name will be left untouched by the SQL Translator to get the PostgreSQL
`SERIAL`/`BIGSERIAL` column types, based on sequences and the
`DEFAULT` clause.

When creating a table with a serial column using start value, the PostgreSQL ODI driver will
convert the Informix type (`SERIAL(100)`) to a PostgreSQL native type, without the
start value (`SERIAL`), and after executing the `CREATE TABLE`
statement, the driver will perform a call to the
`setval('sequence-name',start-value,false)` SQL
function, to reset the underlying sequence. As result, the next inserted row will get the serial
value start-value.

`INSERT` statements that should generate an automatic serial must not use the
serial column and value as `NULL` or zero: When using a `NULL` value
explicity, PostgreSQL will report an non-null constraint error. When using zero, PostgreSQL will
insert that value as is without incrementing the serial as Informix does.

When the Informix serial emulation is enabled, the PostgreSQL ODI driver will automatically add a
`RETURNING` clause to the `INSERT` statements having a serial column.
The purpose of that `RETURNING` clause is to fetch the new generate serial value, and
reset the underlying sequence, if a serial value is explicitly provided and is greater than or equal
to the last generate serial. Therefore, `INSERT` statements can specify an explicit
value that is greater than zero, and the underlying sequence will be reset. When using the
`LOAD` instruction, the sequence for native serial/bigserial columns will also be
automatically reset.

If the serial emulation is disabled, or when inserting rows from outside a Genero program, the
underlying sequence will not be reset by PostgreSQL when a value is specified explicitly. In such
case, you need to manually reset the sequence
with:

```
SELECT setval(pg_get_serial_sequence('[schema.]tabname','colname'),
       (SELECT MAX(colname) FROM tabname), true)
```

After an `INSERT` execution, the last generated `SERIAL` value can
be found in the `sqlca.sqlerrd[2]` register, or with the following
query:

```
SELECT currval(pg_get_serial_sequence('[schema.]tabname','colname'));
```

> **Important:**
>
> If the PostgreSQL sequence was not yet used in the SQL session to generate a new serial number
> (typically because the `INSERT` specifies an explicit value for the serial column),
> the `currval()` function will produce an SQL error, that will cancel the current
> transaction. This problem does not occur using serial emulation and
> `sqlca.sqlerrd[2]`. Without serial emulation, make sure to call
> `currval()` only after an `INSERT` that does not provide a value for
> the serial column.

See also the PostgreSQL documentation for more details about the native serial types.

## Using the `regtable` serial emulation

With the "`regtable`" mode, the `SERIAL` data type is emulated with
a PostgreSQL `INTEGER` data type and `INSERT` triggers using the table
`SERIALREG`, which is dedicated to sequence production. Similarly,
`BIGSERIAL` is converted to `BIGINT` and `INSERT`
trigger.

The triggers can be created manually during the application database installation
procedure, or automatically from a BDL program: When a BDL program executes a `CREATE
[TEMP] TABLE` with a `SERIAL` column, the database interface automatically
converts the `SERIAL/BIGSERIAL` data types to `INTEGER/BIGINT` and
dynamically creates the triggers.

With the "`regtable`" emulation, the `SERIALREG` table must be
created as follows:

```
CREATE TABLE SERIALREG (
   TABLENAME VARCHAR(50) NOT NULL,
   LASTSERIAL DECIMAL(20,0) NOT NULL,
   PRIMARY KEY ( TABLENAME )
)
```

Make sure that other users can read from the `SERIALREG`
table:

```
GRANT SELECT ON SERIALREG TO PUBLIC
```

> **Important:**
>
> The `SERIALREG` table must be created before the triggers.
> The serial production is based on the `SERIALREG` table which registers the last
> generated number for each table. If you delete rows of this table, sequences will restart at 1 and
> you will get unexpected data.

In database creation scripts, all
`SERIAL[(start-value)]` data types must be converted
to `INTEGER` data types and you must create one trigger for each table. To know how
to write those triggers, you can create a small Genero program that creates a table with a
`SERIAL` column. Set the FGLSQLDEBUG environment variable and run the program. The
debug output will show you the native trigger creation command.

After an `INSERT` execution, the last generated `SERIAL` value can
be found in the `sqlca.sqlerrd[2]` register, or with the following query using the
`SERIALREG`
table:

```
SELECT lastserial FROM serialreg WHERE tablename = 'tabname'
```

With this emulation mode, `INSERT` statements using `NULL` for the
serial columns will produce a new serial
value:

```
INSERT INTO tab ( col1, col2 ) VALUES ( NULL, 'data' )
```

This behavior is mandatory in order to support `INSERT` statements that do not use
the serial column:

```
INSERT INTO tab (col2) VALUES ('data')
```

When the Informix serial emulation is enabled, the PostgreSQL ODI driver will automatically add a
`RETURNING` clause to the `INSERT` statements having a serial column.
The purpose of that `RETURNING` clause is to return the new generate serial
value.

Check if your application uses tables with a `SERIAL/BIGSERIAL` column that can
contain a `NULL` value. Consider removing the serial column from the
`INSERT` statements.

## Using the `trigseq` serial emulation

With "`trigseq`", the `SERIAL` data type is emulated with a
PostgreSQL `INTEGER` data type and `INSERT` triggers using a sequence
`tablename_seq`. Similarly, `BIGSERIAL` is
converted to `BIGINT` and `INSERT` trigger.

The triggers can be created manually during the application database installation procedure, or
automatically from a BDL program: When a BDL program executes a `CREATE [TEMP] TABLE`
with a serial column, the database interface automatically converts the
`SERIAL/BIGSERIAL` data types to `INTEGER/BIGINT` and dynamically
creates the triggers.

In database creation scripts, all
`SERIAL[(start-value)]` data types must be converted
to `INTEGER` data types and you must create one trigger for each table. To know how
to write those triggers, you can create a small Genero program that creates a table with a
`SERIAL` column. Set the FGLSQLDEBUG environment variable and run the program. The
debug output will show you the native trigger creation command.

After an `INSERT` execution, the last generated `SERIAL` value can
be found in the `sqlca.sqlerrd[2]` register, or you can execute the following
query:

```
SELECT currval('[schema.]tabname'_seq);
```

> **Important:**
>
> If the PostgreSQL sequence was not yet used in the SQL session to generate a new serial number
> (typically because the `INSERT` specifies an explicit value for the serial column),
> the `currval()` function will produce an SQL error, that will cancel the current
> transaction. This problem does not occur using serial emulation and
> `sqlca.sqlerrd[2]`. Without serial emulation, make sure to call
> `currval()` only after an `INSERT` that does not provide a value for
> the serial column.

With the "`trigseq`" emulation mode, `INSERT` statements using
`NULL` for the serial column will produce a new serial
value:

```
INSERT INTO tab ( col1, col2 ) VALUES ( NULL, 'data' )
```

This behavior is mandatory in order to support `INSERT` statements which do not
use the serial column:

```
INSERT INTO tab (col2) VALUES ('data')
```

When the Informix serial emulation is enabled, the PostgreSQL ODI driver will automatically add a
`RETURNING` clause to the `INSERT` statements having a serial column.
The purpose of that `RETURNING` clause is to return the new generate serial
value.

Check if your application uses tables with a `SERIAL/BIGSERIAL` column that can
contain a `NULL` value. Consider removing the serial column from the
`INSERT` statements.

## Notes common to all serial emulation modes

Since `sqlca.sqlerrd[2]` is defined as `BIGINT`, it can hold values
from `SERIAL` (`INTEGER`) or `BIGSERIAL`
(`BIGINT`) auto incremented columns. The last generated serial can also be found with
the sequence pseudo-column `currval()`, or can it be fetched from the
`LASTSERIAL` column from the `SERIALREG` table when used.

> **Note:**
>
> The serial emulation assume that the same PostgreSQL schema is used when creating the tables
> and when executing the `INSERT` statements. Mixing or switching between PostgreSQL
> schemas is not supported.

For SQL portability, it is recommended to review `INSERT` statements to remove the
serial column from the list.

For example, the following statement:

```
INSERT INTO tab (col1,col2) VALUES ( 0 , p_value )
```

can be converted to:

```
INSERT INTO tab (col2) VALUES (p_value)
```

Static SQL `INSERT` using records defined from the schema file must also be
reviewed:

```
DEFINE rec LIKE tab.*
INSERT INTO tab VALUES (rec.*) -- will use the serial column
```

can be converted to:

```
INSERT INTO tab VALUES rec.* -- without parentheses, serial column is removed
```

> **Important:**
>
> When using the [Static SQL INSERT and UPDATE syntax using record.\* without parentheses](1115-static-sql-statements.md "Describes static SQL statements supported in the language."), make sure that you
> database schema files contain information about serials: This information can be lost when
> extracting the schema from a PostgreSQL database which does not use native serial emulation. See
> [Database Schema](../09_advanced-features/0791-database-schema.md "Defines database table structures with column type information to be reused in program variable definitions.") for more
> details about the serial flag in column type encoding (data type code must be 6)

## Related links

**Related concepts**  

[Auto-incremented columns (serials)](1023-auto-incremented-columns-serials.md "How to implement automatic record keys.")
