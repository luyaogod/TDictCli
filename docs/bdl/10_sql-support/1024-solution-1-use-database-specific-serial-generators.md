---
title: "Solution 1: Use database specific serial generators"
source: "fgl-topics/c_fgl_sql_programming_074.html"
breadcrumb: "SQL support > SQL programming > SQL portability > Auto-incremented columns (serials) > Solution 1: Use database specific serial generators"
type: "concept"
description: "Principle In accordance with the target database, you must use the appropriate native serial generation method. Get the database type with the fgl_db_driver_type() built-in function and use the ..."
---

# Solution 1: Use database specific serial generators

## Principle

In accordance with the target database, you must use the appropriate native serial generation
method. Get the database type with the `fgl_db_driver_type()` built-in function
and use the appropriate SQL statements to insert rows with serial generation.

This solution uses the native auto-increment feature of the target database and is fast at
execution time, but is not very convenient as it requires different code for each database type to
be written. Solutions for auto-incremented columns are vendor-specific. It is of course not
realistic to use this solution in a large application with hundreds of tables.

When possible, the `sqlca.sqlerrd[2]` register will be filled with the last
generated serial, after the execution of the SQL `INSERT` instruction.

| Database Server Type | Serial type support |
| --- | --- |
| IBM® Informix® | Yes, [this is a native Informix feature](1186-what-are-the-supported-ibm-informix-sql-features.md) |
| Microsoft™ SQL Server | Emulated, [see details](1277-serial-and-bigserial-data-types.md) |
| Oracle® MySQL / MariadDB | Emulated, [see details](1331-serial-and-bigserial-data-type.md) |
| Oracle Database Server | Emulated, [see details](1377-serial-and-bigserial-data-types.md) |
| PostgreSQL | Emulated, [see details](1434-serial-and-bigserial-data-types.md) |
| SQLite | Emulated, [see details](1483-serial-and-bigserial-data-types.md) |
| Dameng® | Emulated, [see details](1232-serial-and-bigserial-data-types.md) |

## Implementation

1. Create the database objects required for serial generation in the target database (for example,
   create tables with `SERIAL` columns in IBM Informix, tables with `IDENTITY` columns in
   SQL Server, and `SEQUENCE` database objects in Oracle).
2. Adapt your programs to use the native sequence generators in accordance
   with the database type.

## Example

```
DEFINE t1rec RECORD
          id    INTEGER,
          name  CHAR(50),
          cdate DATE
    END RECORD

CASE fgl_db_driver_type()
 WHEN "ifx"
   INSERT INTO t1 ( id, name, cdate )
          VALUES ( 0, t1rec.name, t1rec.cdate )
   LET t1rec.id = sqlca.sqlerrd[2]
 WHEN "ora"
   INSERT INTO t1 ( id, name, cdate )
          VALUES ( t1seq.nextval, t1rec.name, t1rec.cdate )
   SELECT t1seq.currval INTO t1rec.id FROM dual
 WHEN "msv"
   INSERT INTO t1 ( name, cdate )
           VALUES ( t1rec.name, t1rec.cdate )
   PREPARE s FROM "SELECT @@IDENTITY"
   EXECUTE s INTO t1rec.id
END CASE
```

As
you can see in this example, this solution requires database engine specific coding. Querying the
last generated serial can be centralized in a function, but the insert statements would still need
to be specific to the type of database.

## Related links

**Related concepts**  

[fgl\_db\_driver\_type()](../15_library-reference/2737-fgl-db-driver-type.md "Returns the 3-letter identifier/code of the current database driver.")

[The sqlca diagnostic record](0988-the-sqlca-diagnostic-record.md "The sqlca variable is a predefined record containing SQL statement execution information.")
