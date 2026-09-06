---
title: "Enabling/disabling SQL interruption"
source: "fgl-topics/c_fgl_programs_018.html"
breadcrumb: "Advanced features > Configuration options > OPTIONS (Runtime) > Enabling/disabling SQL interruption"
type: "concept"
---

# Enabling/disabling SQL interruption

> The OPTIONS SQL INTERRUPT instruction enables or disables SQL statement interruption.

## Syntax

```
OPTIONS SQL INTERRUPT { ON | OFF }
```

## Usage

The `OPTIONS SQL INTERRUPT` instruction controls interruption event
detection during the execution of long running SQL statements.

By default, SQL interruption is off.

Pay attention to the fact that not all database servers support SQL interruption.

Depending on the type of database server and client, SQL interruption handling
may require additional overhead. Consider enabling SQL interruption only for SQL
queries that can take a while to execute or can be blocked for a long time because
of concurrent access (locks).

If an SQL statement is interrupted, `sqlca.sqlcode` is set to -213
and the runtime system raises an error. Use a `TRY/CATCH` block or
`WHENEVER ERROR CONTINUE` to trap this SQL error.

## Example

```
MAIN
    DEFINE cnt INTEGER
    DEFER INTERRUPT -- Do not stop if interrupt signal is caught
    CONNECT TO "mydb"
    WHENEVER ERROR CONTINUE -- Continue in case of SQL interrupt error
    OPTIONS SQL INTERRUPT ON -- Enable SQL interruption
    SELECT COUNT(*) INTO cnt FROM stock -- Long running query 
    OPTIONS SQL INTERRUPT OFF -- Disable SQL interruption
    WHENEVER ERROR STOP -- Reset default exception handler
    IF sqlca.sqlcode == -213 THEN
       DISPLAY "SQL Statement interrupted by user"
    END IF
END MAIN
```

## Related links

**Related concepts**  

[Using SQL interruption](../10_sql-support/0993-using-sql-interruption.md "Interrupt long running SQL queries, or interrupt queries waiting for locked data.")
