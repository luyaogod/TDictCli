---
title: "NOTFOUND"
source: "fgl-topics/c_fgl_programs_NOTFOUND.html"
breadcrumb: "Language basics > Predefined constants > NOTFOUND"
type: "concept"
---

# NOTFOUND

> NOTFOUND is a predefined constant used to check if an SQL statement returns rows.

## Syntax

```
NOTFOUND
```

## Usage

The `NOTFOUND` constant is used to test the
execution status of an SQL statement returning a result, to check whether rows have
been found.

The `NOTFOUND` constant is equal to 100.

You typically compare `sqlca.sqlcode` to `NOTFOUND`,
after a `SELECT` statement execution.

## Example

```
MAIN
  DATABASE stores
  SELECT tabid FROM systables WHERE tabid = 1
  IF sqlca.sqlcode = NOTFOUND THEN
    DISPLAY "No row was found"
  END IF
END MAIN
```

## Related links

**Related concepts**  

[SQL execution diagnostics](../10_sql-support/0987-sql-execution-diagnostics.md "If an SQL statement execution fails, error description can be found in the sqlca.sqlcode, SQLSTATE, status and SQLERRMESSAGE predefined registers.")

[SELECT](../10_sql-support/1123-select.md "Produces a result set from a query on database tables.")
