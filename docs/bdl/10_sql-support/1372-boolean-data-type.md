---
title: "BOOLEAN data type"
source: "fgl-topics/c_fgl_odiagora_015.html"
breadcrumb: "SQL support > SQL database guides > Oracle® Database > Data dictionary > BOOLEAN data type"
type: "concept"
description: "Informix® Informix supports the BOOLEAN data type, which can store 't' or 'f' values. Genero BDL implements the BOOLEAN data type in a different way: A BOOLEAN variable stores integer values 1 or 0 ..."
---

# BOOLEAN data type

## Informix®

Informix supports the `BOOLEAN` data type, which can store 't' or 'f'
values.

Genero BDL implements the `BOOLEAN` data type in a different way: A
`BOOLEAN` variable stores integer values `1` or `0`
(for `TRUE` or `FALSE`). This type is designed to hold the result of a
boolean expression.

## ORACLE

Oracle 23c supports a native `BOOLEAN` type in the SQL language. Older Oracle
versions like 21c only support a `BOOLEAN` type in the PL/SQL language.

Note that '1' and '0' char constants are casted by Oracle SQL as `BOOLEAN`
`TRUE/FALSE` equivalents.

## Solution

When using `dbmora_18` + Oracle client 18c, 19c, 21c or `dbmora_23`
+ Oracle client 23c with an Oracle server version prior to 23c (without SQL `BOOLEAN`
support), the ODI driver will always use `CHAR(1)` to store FGL
`BOOLEAN` values. The `BOOLEAN/CHAR(1)` type name conversion in DDL
statements can be disabled with the following FGLPROFILE
entry:

```
dbi.database.dsname.ifxemul.datatype.boolean = false
```

However,
SQL parameter and SQL fetch buffers can still be FGL `BOOLEAN` variables: These are
mapped to '1' and '0' values for `CHAR(1)` columns.

When using `dbmora_23` + Oracle client 23c with an Oracle 23c server (and future
versions), the ODI driver will by default use the native Oracle SQL `BOOLEAN` type
for FGL `BOOLEAN`. To get the former behavior `BOOLEAN/CHAR(1)`
conversion, use the following FGLPROFILE
entry:

```
dbi.database.dsname.ora.boolean.aschar = true
```

In this
case, the `dbi.database.dsname.ifxemul` and
`dbi.database.dsname.ifxemul.datatype.boolean` settings must be
left unset or set to `true`, to create `CHAR(1)` columns for
`BOOLEAN` in DDL statements. If the Informix emulation settings for boolean is
disabled, the use of `BOOLEAN` type in DDL statements will be accepted by Oracle 23c.
However, the ODI driver will use `CHAR(1)` for SQL Parameter and fetch buffer
bindings, and '1'/'0' values.

> **Important:**
>
> The `BOOLEAN/CHAR(1)` conversion rule applies to all SQL tables for a given SQL
> connection. If you have existing boolean data in several tables using `CHAR(1)`
> columns, you need to plan for an upgrade process in order to use the new native SQL
> `BOOLEAN` type of Oracle.

Note that the combination of `dbmora_18` + Oracle client 18c and Oracle server 23c
is possible, but is not supported by Genero: To connect to Oracler 23c server, use
`dbmora_23` + Oracle client 23c.

When using the `BOOLEAN` as `CHAR(1)` emulation, the
`CONSTRUCT` dialog will produce '1' and '0' constants for `BOOLEAN`
fields. When using the native Oracle 23c SQL `BOOLEAN` type,
`CONSTRUCT` will produce SQL conditions using `TRUE` and
`FALSE` constants.

To extract database schemas, the [fgldbsch](../13_programming-tools/2521-fgldbsch.md "The fgldbsch tool generates the database schema files from an existing database.") tool will do the following:

- When connecting to an Oracle server of a version prior to 23c, no `BOOLEAN` type
  is possible in SQL tables, and `CHAR(1)` will be seen as `CHAR(1)` in
  the resulting .sch file.
- When connecting to an Oracle 23c server (and future versions), `BOOLEAN` SQL typed
  columns will be extracted as FGL `BOOLEAN` type (45) by default, or as FGL
  `CHAR(1)` when using conversion option B. Here option B make only sense for legacy
  code where boolean values are handled with FGL `CHAR(1)` variables, assuming '1'/'0'
  values.

The `BOOLEAN` type translation can be controlled with the
following FGLPROFILE
entry:

```
dbi.database.dsname.ifxemul.datatype.boolean = { true | false }
```

For
more details see [IBM Informix emulation parameters in FGLPROFILE](1079-ibm-informix-emulation-parameters-in-fglprofile.md "Emulation of Informix specific SQL features can be controlled with FGLPROFILE entries.").

## Related links

**Related concepts**  

[Using portable data types](1008-using-portable-data-types.md "Only a limited set of data types are really portable across several database engines.")
