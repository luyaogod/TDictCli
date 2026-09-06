---
title: "LOAD and UNLOAD instructions"
source: "fgl-topics/c_fgl_sql_programming_086.html"
breadcrumb: "SQL support > SQL programming > SQL portability > LOAD and UNLOAD instructions"
type: "concept"
---

# LOAD and UNLOAD instructions

> The LOAD and UNLOAD instructions can produce different data formats depending on the database server type.

The Genero BDL runtime system implements the `LOAD` and `UNLOAD`
instructions, by using and `INSERT` (for `LOAD`) or
`SELECT` (for `UNLOAD`) SQL commands.

To handle the values to be used in the `INSERT` statements for
`LOAD`, or to hold the data fetched from the `SELECT` for
`UNLOAD`, the runtime system needs to allocate data buffers dynamically. This
operation requires the description of the column types in order to work.

As each database defines its native set of SQL data types, you must pay attention to the BDL type
that results from the native database column type.

Depending on the native SQL data type, data formatting may be different from Informix® SQL / Genero BDL.

For example, when using Oracle DB, if the table contains a `DATE` column, the
`LOAD` and `UNLOAD` instruction will use the date/time format
`YYYY-MM-DD hh:mm:ss`. Since the native Oracle `DATE` type can be used
to store both Informix / BDL
`DATE` or `DATETIME YEAR TO SECOND` values, and as
`LOAD` / `UNLOAD` need to make date/time to string conversions when
reading from or writing to unload files, it impossible to distinguish the Informix/BDL native date
type formats.

When using `LOAD/UNLOAD`, if the target database server provides the exact
equivalent date types as the native Informix/BDL `DATE`, the date values will use the
DBDATE format setting.

| Database Server Type | LOAD/UNLOAD support |
| --- | --- |
| IBM® Informix | Yes, [native SQL feature](1186-what-are-the-supported-ibm-informix-sql-features.md) |
| Microsoft™ SQL Server | Yes, [with some limitations](1305-load-and-unload.md) |
| Oracle® MySQL / MariadDB | Yes, [see details](1352-load-and-unload.md) |
| Oracle Database Server | Yes, [with some limitations](1408-load-and-unload.md) |
| PostgreSQL | Yes, [see details](1459-load-and-unload.md) |
| SQLite | Yes, [see details](1501-load-and-unload.md) |
| Dameng® | Yes, [see details](1253-load-and-unload.md) |

To export table rows from a database brand that is different from the target database, format the
data, to prepare for the native column types of the destination database.

For example, to unload from Informix and load into Oracle DB, format date/time
columns:

```
FUNCTION ifx_type_to_expr_for_oracle(tn STRING) RETURNS STRING
    CASE
    WHEN tn == "DATE"                              RETURN "TO_CHAR(%1,'%%Y-%%m-%%d 00:00:00')"
    WHEN tn == "DATETIME YEAR TO YEAR"             RETURN "%1||'-01-01 00:00:00'"
    WHEN tn == "DATETIME YEAR TO MONTH"            RETURN "%1||'-01 00:00:00'"
    WHEN tn == "DATETIME YEAR TO DAY"              RETURN "%1||' 00:00:00'"
    WHEN tn == "DATETIME YEAR TO MINUTE"           RETURN "%1||':00'"
    WHEN tn == "DATETIME HOUR TO SECOND"           RETURN "'1900-01-01 '||%1"
    WHEN tn == "DATETIME HOUR TO MINUTE"           RETURN "'1900-01-01 '||%1||':00'"
    WHEN tn MATCHES "DATETIME HOUR TO FRACTION*"   RETURN "'1900-01-01 '||%1"
    WHEN tn MATCHES "INTERVAL YEAR* TO YEAR"       RETURN "%1-00"
    WHEN tn MATCHES "INTERVAL YEAR* TO MONTH"      RETURN "%1"
    WHEN tn MATCHES "INTERVAL HOUR* TO MINUTE"     RETURN "'0 '||%1||':00'"
    WHEN tn MATCHES "INTERVAL HOUR* TO SECOND"     RETURN "'0 '||%1"
    WHEN tn MATCHES "INTERVAL HOUR* TO FRACTION*"  RETURN "'0 '||%1"
    OTHERWISE                                      RETURN "%1"
    END CASE
END FUNCTION

FUNCTION build_select(tn STRING) RETURNS STRING
    DEFINE x INTEGER
    DEFINE sqlcmd base.StringBuffer
    DEFINE curs base.SqlHandle
    LET curs = base.SqlHandle.create()
    CALL curs.prepare("SELECT * FROM "||tn)
    CALL curs.open()
    LET sqlcmd = base.StringBuffer.create()
    CALL sqlcmd.append("SELECT ")
    FOR x=1 TO curs.getResultCount()
        IF x>1 THEN
            CALL sqlcmd.append( ", " )
        END IF
        CALL sqlcmd.append(
             SFMT( ifx_type_to_expr(curs.getResultType(x)),
                   curs.getResultName(x) )
             )
    END FOR
    CALL sqlcmd.append( " FROM " || tn )
    RETURN sqlcmd.toString()
END FUNCTION
```

## Related links

**Related concepts**  

[SQL LOAD and UNLOAD](1175-sql-load-and-unload.md "Describes the instructions to export/import information from/to a database.")
