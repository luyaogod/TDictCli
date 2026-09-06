---
title: "fgl_sqldebug()"
source: "fgl-topics/c_fgl_BuiltInFunctions_FGL_SQLDEBUG.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > fgl_sqldebug()"
type: "concept"
---

# fgl_sqldebug()

> Sets the SQL debug level from program code.

## Syntax

```
FUNCTION fgl_sqldebug(
   level INTEGER )
```

1. level is the SQL debug level to set. Zero resets the level defined by
   FGLSQLDEBUG. -1 prints SQL debug information only for SQL statements producing an error.

## Usage

The [FGLSQLDEBUG](../07_configuration/0534-fglsqldebug.md "Defines the debug level for tracing SQL instructions.") environment variable can
be set, to get SQL debug information for all SQL statements executed by a program. The SQL debug
output is written to the stderr stream.

When program code can be modified and recompiled, the `fgl_sqldebug()` function
can be used to force SQL debug log to isolate a specific set of SQL instructions:

- When calling `fgl_sqldebug()` with a value different from zero, it has the same
  effect as setting FGLSQLDEBUG.
- When calling `fgl_sqldebug()` with zero, it resets the SQL debug level to the
  level defined by FGLSQLDEBUG.

## Related links

**Related concepts**  

[Debugging SQL statements](../10_sql-support/0995-debugging-sql-statements.md "The runtime system can display debug information for SQL statements executed by the program.")
