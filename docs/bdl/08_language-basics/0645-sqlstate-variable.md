---
title: "SQLSTATE [variable]"
source: "fgl-topics/c_fgl_operators_SQLSTATE.html"
breadcrumb: "Language basics > Operators > List of expression elements > SQL related operators > SQLSTATE [variable]"
type: "concept"
---

# SQLSTATE [variable]

> The SQLSTATE predefined variable returns the code corresponding to the last SQL error.

## Syntax

```
SQLSTATE
```

## Usage

The `SQLSTATE` predefined variable returns the ANSI/ISO SQLSTATE code
when an SQL error occurred.

The `SQLSTATE` error code is a standard ANSI specification, but not
all database engines support this feature. Check the database server documentation for
more details.

The variable is `NULL` if the last SQL statement was successful.

## Example

```
MAIN
  DATABASE stores 
  WHENEVER ERROR CONTINUE
  SELECT foo FROM bar 
  DISPLAY SQLSTATE
END MAIN
```

## Related links

**Related concepts**  

[SQL execution diagnostics](../10_sql-support/0987-sql-execution-diagnostics.md "If an SQL statement execution fails, error description can be found in the sqlca.sqlcode, SQLSTATE, status and SQLERRMESSAGE predefined registers.")
