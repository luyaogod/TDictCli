---
title: "BOOLEAN data type"
source: "fgl-topics/c_fgl_odiagdmg_014.html"
breadcrumb: "SQL support > SQL database guides > Dameng® database server > Data dictionary > BOOLEAN data type"
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

## Dameng®

Dameng supports the `BIT` SQL type to store boolean values.

Possible values are 1/0 numbers respectively for true and false boolean values.

## Solution

The Dameng `BIT` type can be used to store `BOOLEAN` program
variable values.

However, the `TRUE/FALSE` keywords cannot be used in SQL statements. Review SQL
statements using `TRUE/FALSE` keywords for boolean columns and replace by 1/0, or use
an SQL parameter for type `BOOLEAN`.

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
