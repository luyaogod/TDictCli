---
title: "TEXT and BYTE (LOB) types"
source: "fgl-topics/c_fgl_odiagdmg_031.html"
breadcrumb: "SQL support > SQL database guides > Dameng® database server > Data dictionary > TEXT and BYTE (LOB) types"
type: "concept"
description: "Informix® Informix provides the TEXT , BYTE , CLOB and BLOB data types to store very large texts or binary data. Legacy Informix 4GL applications typically use the TEXT and BYTE types. Genero BDL does ..."
---

# TEXT and BYTE (LOB) types

## Informix®

Informix provides the `TEXT`,
`BYTE`, `CLOB` and `BLOB` data types to store very
large texts or binary data.

Legacy Informix 4GL applications typically use the
`TEXT` and `BYTE` types.

Genero BDL does not support the Informix
`CLOB` and `BLOB` types.

## Dameng®

Dameng supports the `BLOB` and
`CLOB` data types for large objects (LOB) storage.

The maximum storage size for a Dameng `BLOB/CLOB` is 1Mega bytes.

## Solution

The Dameng database interface can convert BDL
`TEXT` data to Dameng `CLOB` and `BYTE` data to
Dameng `BLOB`.

In programs, when creating a table with `TEXT` or `BYTE` columns,
the ODI driver will convert the type to the native Dameng SQL types `CLOB` or
`BLOB`.

The `TEXT` and
`BYTE` types translation can be controlled with the following FGLPROFILE
entries:

```
dbi.database.dsname.ifxemul.datatype.text = { true | false }
dbi.database.dsname.ifxemul.datatype.byte = { true | false }
```

For more details see [IBM Informix emulation parameters in FGLPROFILE](1079-ibm-informix-emulation-parameters-in-fglprofile.md "Emulation of Informix specific SQL features can be controlled with FGLPROFILE entries.").

## Related links

**Related concepts**  

[Using portable data types](1008-using-portable-data-types.md "Only a limited set of data types are really portable across several database engines.")
