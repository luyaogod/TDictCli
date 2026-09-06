---
title: "TEXT and BYTE (LOB) types"
source: "fgl-topics/c_fgl_odiagora_036.html"
breadcrumb: "SQL support > SQL database guides > Oracle® Database > Data dictionary > TEXT and BYTE (LOB) types"
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

## ORACLE

ORACLE provides `CLOB`, `BLOB`, and `BFILE` data
types to store large text and binary data.

The ORACLE large object types are similar to Informix
LOB types.

## Solution

The ORACLE database interface can convert BDL `TEXT` data to ORACLE
`CLOB` and `BYTE` data to ORACLE `BLOB`.

Genero `TEXT/BYTE` program variables have a limit of 2 gigabytes. Make sure that
the large object data does not exceed this limit.

The ORACLE `BFILE` type is not supported by Genero BDL.

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
