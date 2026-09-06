---
title: "fgldbutl.db_get_last_serial()"
source: "fgl-topics/c_fgl_utility_functions_DB_GET_LAST_SERIAL.html"
breadcrumb: "Library reference > Utility modules > fgldbutl: Database utility module > fgldbutl.db_get_last_serial()"
type: "concept"
---

# fgldbutl.db_get_last_serial()

> Retrieves the last generated serial for a given serial emulation and database table.

## Syntax

```
FUNCTION db_get_last_serial(
   emultype STRING,
   tabname STRING )
  RETURNS BIGINT
```

1. emultype is the serial emulation type (`"native"`,
   `"native2"`, `"regtable"`, `"trigseq"`).
2. tabname is the name of the database table (case senstitive, use
   lowercase).

## Usage

This function executes the SQL query to retrieve the last generated auto-incremented columns, for
the given serial emulation type and database table.

This function must be used after an `INSERT` statement producing a new serial
value, especially when the serial is based on the `BIGINT` type (because
`sqlca.sqlerrd[2]` can only hold `INTEGER` values).

When using this function, configure your connection to avoid the automatic retrieval of the last
generated serial to fill `sqlca.sqlerrd[2]`, by setting the following FGLPROFILE
entry:

```
dbi.database.dbname.ifxemul.datatype.serial.sqlerrd2 = false
```

## Example

```
IMPORT FGL fgldbutl
MAIN
  DEFINE ns BIGINT
  DATABASE mydb
  INSERT INTO mytable VALUES ( ns, 'a new sequence' )
  LET ns = db_get_last_serial("native","mytable")
END MAIN
```

## Related links

**Related concepts**  

[Auto-incremented columns (serials)](../10_sql-support/1023-auto-incremented-columns-serials.md "How to implement automatic record keys.")
