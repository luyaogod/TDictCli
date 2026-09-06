---
title: "base.SqlHandle.getResultValue"
source: "fgl-topics/c_fgl_ClassSqlHandle_getResultValue.html"
breadcrumb: "Library reference > Built-in packages > The base package > The SqlHandle class > base.SqlHandle methods > base.SqlHandle.getResultValue"
type: "concept"
---

# base.SqlHandle.getResultValue

> Returns the value of a column in the result set produced by the SQL statement.

## Syntax

```
getResultValue( index INTEGER )
 RETURNS fgl-type
```

1. index is the ordinal position of the result set column (starts at 1).
2. fgl-type is one of the [primitive data
   types](../08_language-basics/0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data.").

## Usage

Call the `getResultValue()` method to get the value of a column in the result set,
after executing the SQL statement with the [`open()`](3031-base-sqlhandle-open.md "Opens the SQL handle (SELECT or INSERT cursor).") method and fetching a row with [`fetch()`](3020-base-sqlhandle-fetch.md "Fetches a new row from the SQL result set.").

The method takes the position of the column as the parameter.

The value returned can be assigned to a program variable of the type corresponding to the type
name returned by [`getResultType()`](3029-base-sqlhandle-getresulttype.md "Returns the Genero type name of a column in the result set produced by the SQL statement.").

> **Important:**
>
> `TEXT` and `BYTE` values are returned by reference. In order to get
> the value of a `TEXT` or `BYTE` column, define a variable of this type
> and assign the `getResultValue()` return. The returned `TEXT` or
> `BYTE` variable is already located in memory, there is no need to
> `LOCATE` the variable before calling
> `getResultValue()`.
>
> ```
> DEFINE p_text TEXT
> ...
> LET p_text = h.getResultValue(3)
> ...
> ```

## Example

```
DEFINE sh base.SqlHandle, i INT
...
FOR i=1 TO sh.getResultCount()
    DISPLAY sh.getResultValue(i)
END FOR
```

For a complete example, see [Example 2: SqlHandle with result set SQL](3042-example-2-sqlhandle-with-result-set-sql.md).
