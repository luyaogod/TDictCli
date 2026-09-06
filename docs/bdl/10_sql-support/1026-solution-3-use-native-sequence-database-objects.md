---
title: "Solution 3: Use native SEQUENCE database objects"
source: "fgl-topics/c_fgl_sql_programming_076.html"
breadcrumb: "SQL support > SQL programming > SQL portability > Auto-incremented columns (serials) > Solution 3: Use native SEQUENCE database objects"
type: "concept"
description: "Principle Most recent database engines support SEQUENCE database objects. If all database server types you want to use support sequences, it is recommended that you use this solution. Table 1. ..."
---

# Solution 3: Use native SEQUENCE database objects

## Principle

Most recent database engines support `SEQUENCE` database objects. If all database
server types you want to use support sequences, it is recommended that you use this solution.

| Database Server Type | Supports sequences |
| --- | --- |
| IBM® Informix® | Yes |
| Microsoft™ SQL Server | Yes |
| Oracle® MySQL | No, [see details](1331-serial-and-bigserial-data-type.md). |
| MariadDB | Yes |
| Oracle Database Server | Yes |
| PostgreSQL | Yes |
| SQLite | No, [see details](1483-serial-and-bigserial-data-types.md). |
| Dameng® | Yes |

## Implementation

1. Create a `SEQUENCE` object for each table that previously used a
   `SERIAL` column in the IBM Informix database.
2. In database creation scripts (`CREATE TABLE`), replace all
   `SERIAL` types by `INTEGER` (or `BIGINT` if you need
   large integers).
3. Adapt your programs to retrieve a new sequence before inserting a new row. Consider writing a
   function to retrieve a new sequence number, using dynamic SQL to pass the name of the sequence as
   parameter, and adapt to the target database specifics to retrieve a single row (see example
   below).

For databases not supporting `SEQUENCE` objects, the code uses `SELECT
MAX(col)+1` to get a new number, assuming that it's a database like
SQLite used by a single process, where duplicates cannot occur. In a multi-user/process case, for
databases that do not support sequences, consider using the [user-made sequence table](1025-solution-2-generate-serial-numbers-from-your-own-sequence-ta.md) solution.

## Example

```
MAIN
  DEFINE item_rec RECORD
           item_num BIGINT,
           item_name VARCHAR(40)
         END RECORD
  DEFINE x INT
  DATABASE test1
  CREATE TABLE item (
           item_num BIGINT NOT NULL PRIMARY KEY,
           item_name VARCHAR(50)
          )
  CALL sequence_create("item","item_num")
  FOR x=1 TO 5
      LET item_rec.item_num = sequence_next("item","item_num")
      DISPLAY "New sequence: ", item_rec.item_num
      LET item_rec.item_name = "Item#" || item_rec.item_num
      INSERT INTO item VALUES ( item_rec.* )
  END FOR
  DROP TABLE item
  CALL sequence_drop("item")
END MAIN

PRIVATE FUNCTION _supports_sequences() RETURNS BOOLEAN
  DEFINE dt STRING
  LET dt = fgl_db_driver_type()
  RETURN ( dt!="sqt" AND dt!="mys" )
END FUNCTION

PRIVATE FUNCTION _get_next_pkey(tabname STRING, colname STRING) RETURNS BIGINT
  DEFINE sql STRING
  DEFINE val BIGINT
  LET sql = SFMT("SELECT MAX(%1.%2)+1 FROM %1",tabname,colname)
  PREPARE stmt1 FROM sql
  EXECUTE stmt1 INTO val
  RETURN NVL(val,1)
END FUNCTION

PUBLIC FUNCTION sequence_create(tabname STRING, colname STRING)
  DEFINE val BIGINT
  IF NOT _supports_sequences() THEN RETURN END IF
  LET val = _get_next_pkey(tabname,colname)
  EXECUTE IMMEDIATE SFMT("CREATE SEQUENCE %1_seq START WITH %2",tabname,val)
END FUNCTION

PUBLIC FUNCTION sequence_drop(tabname STRING)
  IF NOT _supports_sequences() THEN RETURN END IF
  EXECUTE IMMEDIATE SFMT("DROP SEQUENCE %1_seq",tabname)
END FUNCTION

PRIVATE FUNCTION _unique_row_condition() RETURNS STRING
    CASE fgl_db_driver_type()
        WHEN "ifx" RETURN " FROM systables WHERE tabid=1"
        WHEN "pgs" RETURN " FROM pg_class WHERE relname='pg_class'"
        WHEN "ora" RETURN " FROM dual"
        OTHERWISE RETURN " "
    END CASE
END FUNCTION

PUBLIC FUNCTION sequence_next(tabname STRING, colname STRING) RETURNS BIGINT
  DEFINE sql STRING
  DEFINE seqname STRING
  DEFINE val BIGINT
  LET seqname = tabname||"_seq"
  CASE
     WHEN NOT _supports_sequences()
          -- WARNING: Duplicates can be created in multi-user databases, must
          -- be controlled with PRIMARY KEY or UNIQUE constraint!
          RETURN _get_next_pkey(tabname,colname)
     WHEN fgl_db_driver_type()=="pgs"
          LET sql = SFMT("SELECT nextval('%1') %2",seqname,_unique_row_condition())
     WHEN fgl_db_driver_type()=="esm"
       OR fgl_db_driver_type()=="ftm"
       OR fgl_db_driver_type()=="snc"
       OR fgl_db_driver_type()=="ntz"
          LET sql = SFMT("SELECT NEXT VALUE FOR %1",seqname)
     OTHERWISE
          LET sql = SFMT("SELECT %1.nextval %2",seqname,_unique_row_condition())
  END CASE
  TRY
      PREPARE seq FROM sql
      EXECUTE seq INTO val
  CATCH
      DISPLAY SFMT("ERROR: Could not produce new sequence for %1", tabname)
      EXIT PROGRAM 1
      RETURN NULL
  END TRY
  RETURN val
END FUNCTION
```

## Related links

**Related concepts**  

[INTEGER](../08_language-basics/0562-integer.md "The INTEGER data type is used for storing large whole numbers.")

[BIGINT](../08_language-basics/0554-bigint.md "The BIGINT data type is used for storing very large whole numbers.")
