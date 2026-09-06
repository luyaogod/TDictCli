---
title: "SERIAL and BIGSERIAL data types"
source: "fgl-topics/c_fgl_odiagmsv_010.html"
breadcrumb: "SQL support > SQL database guides > Microsoft™ SQL Server > Data dictionary > SERIAL and BIGSERIAL data types"
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

## Microsoft™ SQL Server

Microsoft SQL Server provides
`IDENTITY` columns to auto-generate numeric values:

- When creating a table, the `IDENTITY` keyword must be specified after the column
  data type:

  ```
  CREATE TABLE tab1 ( k INTEGER IDENTITY, c NVARCHAR(10) )
  ```
- You can specify a start value and an increment with
  `IDENTITY(start,incr)`:

  ```
  CREATE TABLE tab1 ( k INTEGER IDENTITY(100,2), ...
  ```
- A new number is automatically created when inserting a new
  row:

  ```
  INSERT INTO tab1 ( c ) VALUES ( 'aaa' )
  ```
- To get the last generated number, Microsoft SQL
  SERVER provides the `SCOPE_IDENTITY()` function. The `@@IDENTITY`
  global T-SQL variable is not recommended, as it is scope-less.
- To put a specific value into an `IDENTITY` column, the `SET
  IDENTITY_INSERT` command must be used:

  ```
  SET IDENTITY_INSERT tab1 ON
  INSERT INTO tab1 ( k, c ) VALUES ( 100, 'aaa' )
  SET IDENTITY_INSERT tab1 OFF
  ```

Informix SERIALs and SQL Server
`IDENTITY` columns are quite similar; the main difference is that SQL Server does
not allow you to use the zero value for the identity column when inserting a new row.

Starting with version 2012, Microsoft SQL Server
supports sequences:

```
-- To create a sequence object:
CREATE SEQUENCE myseq START WITH 100 INCREMENT BY 1;

-- To get a new sequence value:
SELECT NEXT VALUE FOR myseq;

-- To find the current sequence value (last generated)
SELECT convert(bigint, current_value) FROM sys.sequences WHERE name = 'myseq';

-- To reset the sequence with a new start number:
ALTER SEQUENCE myseq START WITH 100;
```

## Solution

> **Note:**
>
> For best SQL portability when using different types of databases, consider using sequences as
> described in [Solution 3: Use native SEQUENCE database objects](1026-solution-3-use-native-sequence-database-objects.md).

To emulation Informix serials with SQL Server, you can
use three different solutions:

1. Native SQL Server `IDENTITY` columns (default).
2. `INSERT` triggers based on sequences (requires SQL Server 2012 and +).
3. `INSERT` triggers based on the `SERIALREG` table (for SQL Server
   prior to 2012).

The method used to emulate `SERIAL` types is defined by the
`ifxemul.datatype.serial.emulation` FGLPROFILE
parameter:

```
dbi.database.dbname.ifxemul.datatype.serial.emulation = {"native"|"trigseq"|"regtable"}
```

1. `native`: uses `IDENTITY` columns.
2. `trigseq`: uses `INSERT` triggers with sequences (`CREATE
   SEQUENCE`).
3. `regtable`: uses `INSERT` triggers with the
   `SERIALREG` table.

The default emulation technique is "`native`".

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

The native IDENTITY-based solution is faster, but does not allow explicit serial value
specification in insert statements; the other solutions are slower but allow explicit serial value
specification in `INSERT` statements.

> **Important:**
>
> The trigger-based solutions are provided to simplify the conversion from
> Informix, but are slower as the solution uses `IDENTITY` columns. To get best
> performances, we strongly recommend that you use native `IDENTITY` columns instead of
> triggers.

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

## 1. Using the `native` serial emulation

In order to use the "`native`" emulation, make sure that the following FGLPROFILE
entry is not defined (it is the default), or is set the entry to
`"native"`:

```
dbi.database.dbname.ifxemul.datatype.serial.emulation = "native"
```

In database creation scripts, all `SERIAL[(n)]` data types must be converted by
hand to `INTEGER IDENTITY[(n,1)]` data types, while `BIGSERIAL[(n)]`
data types must be converted by hand to `BIGINT IDENTITY[(n,1)]` data types.

Tables created from the BDL programs can use the `SERIAL` or
`BIGSERIAL` data type: When a BDL program executes a `CREATE [TEMP]
TABLE` with such column type, the database interface automatically converts to an identity
column.

The new generated `SERIAL` or `BIGSERIAL` value is available from
the `sqlca.sqlerrd[2]` variable. The database interface performs a `SELECT
SCOPE_IDENTITY()`. You can also get the last generated serial with the
`SCOPE_IDENTITY()` SQL function.

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

Since 2.10.06, SELECT \* FROM table INTO TEMP with original table having an IDENTITY column is
supported: The database driver converts the Informix SELECT INTO TEMP to the following sequence of statements:

```
SELECT selection-items INTO #table FROM ... WHERE 1=2
SET IDENTITY_ INSERT #table ON
INSERT INTO #table ( column-list ) SELECT original select clauses
SET IDENTITY_ INSERT #table OFF
```

See also [temporary tables](1291-temporary-tables.md).

## 2. Using the `trigseq` serial emulation (SQL Server 2012 and +)

In order to use the serial emulation based on triggers and sequences, make sure that
all database users creating tables in program have permissions to create/drop
sequences and triggers.

Define the FGLPROFILE entry to enable `"trigseq"` serial
emulation:

```
dbi.database.dbname.ifxemul.datatype.serial.emulation = "trigseq"
```

In database creation scripts, all SERIAL[(n)] data types must be converted to INTEGER
data types, BIGSERIAL must be converted to BIGINT and you must create one trigger
for each table. To know how to write those triggers, you can create a small Genero
program that creates a table with a SERIAL column. Set the FGLSQLDEBUG environment
variable and run the program. The debug output will show you the native trigger
creation command using a sequence.

Tables created from the BDL programs can use the SERIAL data type. When a BDL program executes a
CREATE [TEMP] TABLE with a SERIAL column, the database interface automatically converts the
"SERIAL[(n)]" data type to "INTEGER" and creates the `INSERT` triggers. When using
BIGSERIAL[(n)], the column is converted to a BIGINT.

> **Important:**
>
> - SQL Server does not allow you to create triggers on temporary tables.
>   Therefore, you cannot create temp tables with a SERIAL column when using
>   this solution.
> - SELECT ... INTO TEMP statements using a table created with a SERIAL column
>   do not automatically create the SERIAL triggers in the temporary table. The
>   type of the column in the new table is INTEGER. Similarly, a BIGSERIAL
>   column becomes BIGINT.
> - When a table is dropped, all associated triggers are also dropped.
> - INSERT statements using NULL for the SERIAL column will produce a new serial
>   value, instead of using
>   NULL:
>
>   ```
>   INSERT INTO tab ( col1, col2 ) VALUES ( NULL, 'data' )
>   ```
>
>   This behavior is mandatory in order to support INSERT statements which do not use the serial
>   column:
>
>   ```
>   INSERT INTO tab (col2) VALUES ('data')
>   ```
>
>   Check if your application uses tables with a SERIAL column that can contain a NULL value.

## 3. Using the `regtable` serial emulation (SQL Server versions prior to 2012)

This solution is supported for SQL Server versions prior to 2012, if your server is SQL Server
2012 or +, consider using the "trigseq" emulation instead.

In order to use the serial emulation based on triggers and the SERIALREG table, make sure that
all database users creating tables in program have permissions to create/drop triggers.

Then, prepare the database and create the SERIALREG table as follows:

```
CREATE TABLE serialreg (
    tablename VARCHAR(50) NOT NULL,
    lastserial BIGINT NOT NULL,
    PRIMARY KEY ( tablename )
)
```

Make sure that other users can read from the `SERIALREG` table (using database
roles).

The SERIALREG table and columns have to be created with lower case names, since the SQL SERVER
database is created with case sensitive names, because triggers are using this table
in lower case.

Define the FGLPROFILE entry to enable `"regtable"` serial emulation:

```
dbi.database.dbname.ifxemul.datatype.serial.emulation = "regtable"
```

In database creation scripts, all SERIAL[(n)] data types must be converted to INTEGER data types,
BIGSERIAL must be converted to BIGINT and you must create one trigger for each table. To know how to
write those triggers, you can create a small Genero program that creates a table with a SERIAL
column. Set the FGLSQLDEBUG environment variable and run the program. The debug output will show you
the native trigger creation command using the SERIALREG table.

Tables created from the BDL
programs can use the SERIAL data type. When a BDL program executes a CREATE [TEMP] TABLE with a
SERIAL column, the database interface automatically converts the "SERIAL[(n)]" data type to
"INTEGER" and creates the `INSERT` triggers. When using BIGSERIAL[(n)], the column is
converted to a BIGINT.

> **Important:**
>
> - The serial production is based on the SERIALREG table which registers the last generated number
>   for each table. If you delete rows of this table, sequences will restart at 1 and you will get
>   unexpected data.
> - SQL Server does not allow you to create triggers on temporary tables. Therefore, you cannot
>   create temp tables with a SERIAL column when using this solution.
> - SELECT ... INTO TEMP statements using a table created with a SERIAL column do not automatically
>   create the SERIAL triggers in the temporary table. The type of the column in the new table is
>   INTEGER. Similarly, a BIGSERIAL column becomes BIGINT.
> - When a table is dropped, all associated triggers are also dropped.
> - INSERT statements using NULL for the SERIAL column will produce a new serial value, instead of
>   using NULL:
>
>   ```
>   INSERT INTO tab (col1,col2) VALUES ( NULL,'data')
>   ```
>
>   This behavior
>   is mandatory in order to support INSERT statements which do not use the serial
>   column:
>
>   ```
>   INSERT INTO tab (col2) VALUES ('data')
>   ```
>
>   Check if your
>   application uses tables with a SERIAL column that can contain a NULL value.

## Related links

**Related concepts**  

[Auto-incremented columns (serials)](1023-auto-incremented-columns-serials.md "How to implement automatic record keys.")

[The sqlca diagnostic record](0988-the-sqlca-diagnostic-record.md "The sqlca variable is a predefined record containing SQL statement execution information.")

[FGLPROFILE entries for core language](../07_configuration/0486-fglprofile-entries-for-core-language.md "This is a summary of FGLPROFILE entries supported by the core BDL.")
