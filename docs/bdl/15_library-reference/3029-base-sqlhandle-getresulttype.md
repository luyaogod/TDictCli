---
title: "base.SqlHandle.getResultType"
source: "fgl-topics/c_fgl_ClassSqlHandle_getResultType.html"
breadcrumb: "Library reference > Built-in packages > The base package > The SqlHandle class > base.SqlHandle methods > base.SqlHandle.getResultType"
type: "concept"
---

# base.SqlHandle.getResultType

> Returns the Genero type name of a column in the result set produced by the SQL statement.

## Syntax

```
getResultType( index INTEGER )
 RETURNS STRING
```

1. index is the ordinal position of the result set column (starts at 1).

## Usage

Call the `getResultType()` method to query the type of a column in the result set,
after executing the SQL statement with the [`open()`](3031-base-sqlhandle-open.md "Opens the SQL handle (SELECT or INSERT cursor).") method and fetching a row with [`fetch()`](3020-base-sqlhandle-fetch.md "Fetches a new row from the SQL result set.").

The method takes the position of the column as the parameter.

The type name is a string that represents a Genero type. For example, "`INTEGER`",
"`DECIMAL(10,2)`", "`DATE`", "`DATETIME YEAR TO
SECOND`".

For `DATETIME` types, when the second qualifier is a `FRACTION`
with a precision of 3, the type name returned is "`DATETIME … TO FRACTION`" without
precision.

The column type returned by the `getResultType()` method can differ, depending on
the type of database server, the character set locale, and the length semantics used.

The native column type is provided by the database client software. This native type is then
converted to a Genero type. For example, if you define a column with Oracle's native DATE type, the
resulting type returned by `getResultType()` will be a Genero `DATETIME YEAR
TO SECOND`, which corresponds to [Oracle's DATE
(YYYY-MM-DD hh:mm:ss)](../10_sql-support/1375-date-and-datetime-data-types.md).

Furthermore, depending on the character set and the length semantics used, the size of
`CHAR/VARCHAR` types can differ, in order to get a Genero type that is large enough
to hold the maximum character string the query may return. The resulting type can for example be
different when using UTF-8 with `CHAR` or with `BYTE` length
semantics. When using a single byte character set and `BYTE` length semantics, or
when using UTF-8 and `CHAR` length semantics, the char size returned by
`getResultType()` will always match the size of the database column in character
units. For more details, see [Length semantics settings](../09_advanced-features/0881-length-semantics-settings.md).

## Example

```
DEFINE sh base.SqlHandle, i INT
...
FOR i=1 TO sh.getResultCount()
    DISPLAY sh.getResultType(i)
END FOR
```

For a complete example, see [Example 2: SqlHandle with result set SQL](3042-example-2-sqlhandle-with-result-set-sql.md).

## Related links

**Related concepts**  

[Defining the application locale](../09_advanced-features/0879-defining-the-application-locale.md "This section describes the settings defining the application locale, changing the behavior of the compilers and runtime system.")
