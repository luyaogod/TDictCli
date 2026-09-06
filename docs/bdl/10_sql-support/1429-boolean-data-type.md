---
title: "BOOLEAN data type"
source: "fgl-topics/c_fgl_odiagpgs_015.html"
breadcrumb: "SQL support > SQL database guides > PostgreSQL > Data dictionary > BOOLEAN data type"
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

## PostgreSQL

PostgreSQL supports the `BOOLEAN` data type and stores `'t'` or
`'f'` values for `TRUE` and `FALSE` representation.

> **Note:**
>
> It is not possible to insert the integer values `1` or `0` into
> PostgreSQL `BOOLEAN` columns: Values must be `true`,
> `false`, `'t'`, `'f'`, `'1'` or
> `'0'`.

## Solution

The PostgreSQL database interface converts the `BOOLEAN` integer values to a
`CHAR(1)` values `'1'` or `'0'`.

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
