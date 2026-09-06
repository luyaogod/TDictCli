---
title: "fgl_db_driver_type()"
source: "fgl-topics/c_fgl_BuiltInFunctions_FGL_DB_DRIVER_TYPE.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > fgl_db_driver_type()"
type: "concept"
---

# fgl_db_driver_type()

> Returns the 3-letter identifier/code of the current database driver.

## Syntax

```
FUNCTION fgl_db_driver_type()
  RETURNS CHAR(3)
```

## Usage

This function can be called after connecting to a database server with the
`CONNECT` or `DATABASE` instructions, in order
to identify the type of the target database with the driver type.

Returned value is a 3-letter driver code, in lower case, such as "`ifx`",
"`ora`", "`snc`", etc.

See [the database drivers table](../10_sql-support/1073-database-driver-specification-driver.md) for more details
about the list of database driver types.

The function returns [`NULL`](../08_language-basics/0572-null.md "The NULL constant defines a non-value.") if there
is no current database driver (for example, if database connection is not yet established).

## Related links

**Related concepts**  

[Solution 1: Use database specific serial generators](../10_sql-support/1024-solution-1-use-database-specific-serial-generators.md "Solution 1: Use database specific serial generators")
