---
title: "SERIAL and BIGSERIAL data types"
source: "fgl-topics/c_fgl_odiagora_008.html"
breadcrumb: "SQL support > SQL database guides > Oracle® Database > Data dictionary > SERIAL and BIGSERIAL data types"
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

## ORACLE

Oracle® provides several solutions to
implement auto-incremented columns:

1. Sequence objects can be created to generate numbers (`CREATE SEQUENCE`,
   `seqname.currval`).
2. Sequences can be referenced in `DEFAULT ON NULL` column clauses.
3. Columns can be defined with the `GENERATE ... AS IDENTITY` clause. However, with
   pure Oracle SQL, there is no way to get the last generated number. This is only possible with PL/SQL
   by using the `RETURNING … INTO …` clause.

Details about Oracle sequences:

- Sequences are totally detached from tables.
- The purpose of sequences is to provide unique integer numbers.
- Sequences are identified by a sequence name.
- To create a sequence, you must use the `CREATE SEQUENCE` statement. Once a
  sequence is created, it is permanent (like a table).
- To get a new sequence value, you must use the `nextval` keyword, preceded by the
  name of the sequence. The `seqname.nextval` expression can be used
  in `INSERT`
  statements:

  ```
  INSERT INTO tab1 VALUES ( tab1_seq.nextval, ... )
  ```
- To get the last generated number, Oracle provides the `currval`
  keyword:

  ```
  SELECT seqname.currval FROM DUAL
  ```
- In order to improve performance, Oracle can handle a set of sequences in the cache (See `CREATE SEQUENCE` syntax
  in the Oracle documentation).

## Solution

> **Note:**
>
> For best SQL portability when using different types of databases, consider using sequences as
> described in [Solution 3: Use native SEQUENCE database objects](1026-solution-3-use-native-sequence-database-objects.md).

The `SERIAL` data type can be emulated with sequences used in
`INSERT` triggers or with the `DEFAULT ON NULL` clause.

The method used to emulate serial types is defined by the
`ifxemul.datatype.serial.emulation` FGLPROFILE
parameter:

```
dbi.database.dbname.ifxemul.datatype.serial.emulation = {"native"|"native2"|"regtable"}
```

- `native`: uses `INSERT` triggers with sequences.
- `native2`: uses `DEFAULT ON NULL` column clause with
  sequences.
- `regtable`: uses `INSERT` triggers with the
  `SERIALREG` table.

The default emulation technique is "`native`".

> **Important:**
>
> Do not mix a serial emulation modes: Choose one of the provided serial
> emulations for all the tables of the application database. Do note create the tables with a given
> serial emulation, and later switch to another serial emulation for `INSERT`
> statements. This is not supported.

> **Note:**
>
> Genero does use Oracle's `GENERATED ... AS IDENTITY` feature to emulate
> Informix SERIALs, because it is not possible to easily get the last generated/inserted number, in
> order to fill `sqlca.sqlerrd[2]`.

For serial emulations using Oracle sequences ("`native`" and
"`native2`"), if the table with serial column is owned by another user, the current
user must have select privileges to query the underlying sequence. To grant globally permissions on
serial sequences in a given schema, use the following SQL
command:

```
GRANT SELECT ON sequence-name TO PUBLIC;
```

The serial type emulation can be enabled or disabled with the
following FGLPROFILE
entries:

```
dbi.database.dbname.ifxemul.datatype.serial = {true|false}
dbi.database.dbname.ifxemul.datatype.serial8 = {true|false}
dbi.database.dbname.ifxemul.datatype.bigserial = {true|false}
```

> **Important:**
>
> The "`regtable`" emulation based on the
> `SERIALREG` table is provided to simplify the migration from Informix. We strongly
> recommend that you use the "`native`" or "`native2`" method instead.
> The "`native2`" method is the fastest solution when inserting a large number of rows
> in the database.

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

## Notes common to all serial emulation modes

When a BDL program executes a `CREATE [TEMP] TABLE` with a `SERIAL`
column, the Oracle interface automatically
creates the additional SQL objects (column clauses, sequences or triggers) to generate numbers when
an `INSERT` statement is performed.

Users executing programs which create tables with `SERIAL` columns must have the
`CONNECT` and `RESOURCE` roles assigned to create triggers and
sequences.

`SERIAL[(n)]` data types are converted to `NUMBER(10,0)`, while
`BIGSERIAL[(n)]` is replaced by `NUMBER(20,0)`.

For `SERIAL` and `BIGSERIAL` types, the
`sqlca.sqlerrd[2]` register is filled as expected with the last generated serial
value. Another option is to fetch the sequence pseudo-column `CURR_VAL` or fetch the
`LASTSERIAL` column from the `SERIALREG` table, if used.

Check whether your application uses tables with a `SERIAL` column that can contain
a `NULL` value: `INSERT` statements using `NULL` for
the `SERIAL` column will produce a new serial value:

```
INSERT INTO tab ( col1, col2 ) VALUES ( NULL, 'data' )
```

This behavior is mandatory in order to support `INSERT` statements that do not
use the serial column:

```
INSERT INTO tab (col2) VALUES ('data')
```

For SQL portability, it is recommended to review `INSERT` statements to remove
the `SERIAL` column from the list. For example, the following
statement:

```
INSERT INTO tab (col1,col2) VALUES (0, p_value)
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

can be converted
to:

```
INSERT INTO tab VALUES rec.* -- without parentheses, serial column is removed
```

When using the [Static SQL
INSERT or UPDATE syntax using record.\* without parentheses](1115-static-sql-statements.md "Describes static SQL statements supported in the language."), make sure that your database schema
files contain information about serials. This information can be lost when extracting the schema
from an Oracle database. See [Database Schema](../09_advanced-features/0791-database-schema.md "Defines database table structures with column type information to be reused in program variable definitions.") for more
details about the serial flag in column type encoding (data type code must be 6).

If the "`native`" or "`regtable`" emulation is used,
inserting rows with Oracle tools like
SQL\*Plus or SQL\*Loader will execute the `INSERT` triggers. When loading big tables,
you can disable triggers with `ALTER TRIGGER [ENABLE | DISABLE]` (see Oracle documentation for more details). After
reactivation of the serial triggers, the SERIAL sequences must be re-initialized (use
`serialpkg.create_sequence('tab','col')`) or re-execute the PL/SQL script containing
the sequence and trigger creation.

## Using the `native` serial emulation

Each table having a `SERIAL` column needs an `INSERT
TRIGGER` and a `SEQUENCE` dedicated to `SERIAL`
generation.

To know how to write those sequences and triggers, you can create a small Genero program that
creates a table with a serial column. Set the FGLSQLDEBUG environment variable and run the program.
The debug output will show you the native SQL commands to create the sequence and the trigger.

For [temporary tables](1390-temporary-tables.md),
the trigger and the sequence are dropped automatically after a "`DROP TABLE temptab`"
or when the program disconnects from the database.

## Using the `native2` serial emulation

With this emulation, a `SERIAL` type is converted to a `DEFAULT
ON NULL` clause using a sequence created automatically by the database driver, for
example:

```
CREATE TABLE t1 ( mykey SERIAL(100), .... )
```

is converted to:

```
CREATE SEQUENCE t1_srl INCREMENT BY 1 START WITH 100
CREATE TABLE t1 (mykey NUMBER(10,0) DEFAULT ON NULL t1_srl.nextval , ...
```

For [temporary tables](1390-temporary-tables.md),
the sequence is dropped automatically after a "`DROP TABLE temptab`" or when the
program disconnects from the database.

> **Note:**
>
> The `native2` serial emulation uses the `DEFAULT ON NULL`
> clause, supported by Oracle, starting from version 12.1. When using this serial emulation with
> an Oracle database version prior to 12.1, a `CREATE TABLE` statement using a
> serial will produce the SQL error [-6370](../15_library-reference/4483-genero-bdl-errors.md).

## Using the `regtable` serial emulation

Each table having a `SERIAL` column needs an `INSERT TRIGGER`
which uses the `SERIALREG` table dedicated to `SERIAL` registration.

Prepare the database and create the `SERIALREG` table as
follows:

```
CREATE TABLE serialreg (
    tablename VARCHAR2(50) NOT NULL,
    lastserial NUMBER(20,0) NOT NULL,
    PRIMARY KEY ( tablename )
)
```

Make sure that other users can read from the `SERIALREG`
table:

```
GRANT SELECT ON serialreg TO PUBLIC
```

> **Important:**
>
> The
> `SERIALREG` table must exist in the database before creating the serial triggers, and
> have `SELECT` permission for other users to fill the
> `sqlca.sqlerrd[2]` register.

In database creation scripts, all `SERIAL[(n)]` data types must be converted to
`INTEGER` data types and you must create one trigger for each table.
`SERIAL8/BIGSERIAL` columns must be converted to `NUMBER(20,0)`. To
know how to write those triggers, you can create a small Genero program that creates a table with a
`SERIAL` column. Set the FGLSQLDEBUG environment variable and run the program. The
debug output will show you the native trigger creation command.

The serial production is based on the `SERIALREG` table which
registers the last generated number for each table. If you delete rows of this table, sequences will
restart at start values and you might get duplicated values.

For [temporary tables](1390-temporary-tables.md),
the trigger is dropped automatically after a "`DROP TABLE temptab`" or when the
program disconnects from the database.

## Related links

**Related concepts**  

[Auto-incremented columns (serials)](1023-auto-incremented-columns-serials.md "How to implement automatic record keys.")

[FGLPROFILE entries for core language](../07_configuration/0486-fglprofile-entries-for-core-language.md "This is a summary of FGLPROFILE entries supported by the core BDL.")

[The sqlca diagnostic record](0988-the-sqlca-diagnostic-record.md "The sqlca variable is a predefined record containing SQL statement execution information.")

[Debugging SQL statements](0995-debugging-sql-statements.md "The runtime system can display debug information for SQL statements executed by the program.")
