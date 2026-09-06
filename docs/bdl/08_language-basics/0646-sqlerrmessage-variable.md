---
title: "SQLERRMESSAGE [variable]"
source: "fgl-topics/c_fgl_operators_SQLERRMESSAGE.html"
breadcrumb: "Language basics > Operators > List of expression elements > SQL related operators > SQLERRMESSAGE [variable]"
type: "concept"
---

# SQLERRMESSAGE [variable]

> The SQLERRMESSAGE predefined variable holds the error message corresponding to the last SQL error.

## Syntax

```
SQLERRMESSAGE
```

## Usage

The `SQLERRMESSAGE` predefined variable returns the error message if an
SQL error occurred.

The variable is `NULL` if the last SQL statement was successful.

## Example

```
MAIN
  DATABASE stores 
  WHENEVER ERROR CONTINUE
  SELECT foo FROM bar 
  DISPLAY SQLERRMESSAGE
END MAIN
```

## Related links

**Related concepts**  

[SQL execution diagnostics](../10_sql-support/0987-sql-execution-diagnostics.md "If an SQL statement execution fails, error description can be found in the sqlca.sqlcode, SQLSTATE, status and SQLERRMESSAGE predefined registers.")
