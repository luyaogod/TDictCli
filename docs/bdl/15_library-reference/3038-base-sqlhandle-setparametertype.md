---
title: "base.SqlHandle.setParameterType"
source: "fgl-topics/c_fgl_ClassSqlHandle_setParameterType.html"
breadcrumb: "Library reference > Built-in packages > The base package > The SqlHandle class > base.SqlHandle methods > base.SqlHandle.setParameterType"
type: "concept"
---

# base.SqlHandle.setParameterType

> Defines the type of an SQL parameter for this SQL handle.

## Syntax

```
setParameterType(
   index INTEGER,
   type STRING )
```

1. index is the ordinal position of the `?` SQL parameter (starts
   at 1).
2. type is the name of the primitive type to be used, such as
   `"DECIMAL(10,2)"`, `"DATETIME YEAR TO SECOND"`.

## Usage

The `setParameterType()` method defines the type of an SQL parameter specified
with a `?` place holder in the string passed to the [`prepare()`](3035-base-sqlhandle-prepare.md "Prepares an SQL statement for the SQL handle.") method.

The SQL statement must have been prepared with a `prepare()` call.

The method will raise error [-8131](4483-genero-bdl-errors.md),
if the index passed as parameter is lower than 1 or greater than the number of `?` parameter
placeholders in the prepared SQL statement.

The type of an SQL parameter is important and must match the type of the target SQL
column. After [`prepare()`](3035-base-sqlhandle-prepare.md "Prepares an SQL statement for the SQL handle."), if
`setParameterType()` is not used, the type of the SQL parameter is defined by the
expression passed to any [`setParameter()`](3037-base-sqlhandle-setparameter.md "Sets the value of an SQL parameter for this SQL handle.") call. However, it is also possible to pre-define the type of
an SQL parameter with the [`setParameterType()`](3038-base-sqlhandle-setparametertype.md "Defines the type of an SQL parameter for this SQL handle.") method. When calling `setParameter()`
with strings representing the actual values, it is mandatory to define the real type with
`setParameterType()`, to avoid any SQL type conversion issues. Once the type is
defined by `setParameterType()`, the same [type conversion rules](../08_language-basics/0578-data-type-conversion-reference.md "This topic lists type conversion rules for all data types.") apply as with a `LET variable =
string` instruction.

## Example

```
DEFINE sh base.SqlHandle
DEFINE pkey INTEGER
DEFINE value STRING
...
CALL sh.prepare("UPDATE customer SET cust_creadate=? WHERE cust_id=?")
CALL sh.setParameterType(1,"DATE")
CALL sh.setParameterType(2,"INTEGER")
...
LET value = TODAY -- DBDATE formatted string representing a date value
CALL h.setParameter(1,value) -- Type is defined by setParameterType()
CALL h.setParameter(2,pkey)  -- Type is defined by pkey INTEGER
CALL sh.execute()
```

For a complete example, see [Example 1: SqlHandle with simple SQL](3041-example-1-sqlhandle-with-simple-sql.md).
