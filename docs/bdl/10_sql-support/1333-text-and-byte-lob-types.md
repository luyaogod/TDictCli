---
title: "TEXT and BYTE (LOB) types"
source: "fgl-topics/c_fgl_odiagmys_008.html"
breadcrumb: "SQL support > SQL database guides > Oracle® MySQL / MariaDB > Data dictionary > TEXT and BYTE (LOB) types"
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

## Oracle® MySQL and MariaDB

MySQL and MariaDB provide the following types to store large objects in the database:

| MySQL data type | Description |
| --- | --- |
| `TINYTEXT` | To store text data with a length < 2^8 bytes |
| `TEXT` | To store text data with a length < 2^16 bytes |
| `MEDIUMTEXT` | To store text data with a length < 2^24 bytes |
| `LONGTEXT` | To store text data with a length < 2^32 bytes |
| `TINYBLOB` | To store binary data with a length < 2^8 bytes |
| `BLOB` | To store binary data with a length < 2^16 bytes |
| `MEIDUMBLOB` | To store binary data with a length < 2^24 bytes |
| `LONGBLOB` | To store binary data with a length < 2^32 bytes |

## Solution

The MySQL and MariaDB database interface can convert BDL `TEXT` data to
`LONGTEXT` and `BYTE` data to `LONGBLOB`.

The `TEXT` and
`BYTE` types translation can be controlled with the following FGLPROFILE
entries:

```
dbi.database.dsname.ifxemul.datatype.text = { true | false }
dbi.database.dsname.ifxemul.datatype.byte = { true | false }
```

For more details see [IBM Informix emulation parameters in FGLPROFILE](1079-ibm-informix-emulation-parameters-in-fglprofile.md "Emulation of Informix specific SQL features can be controlled with FGLPROFILE entries.").

Genero `TEXT/BYTE` program variables have a limit of 2 gigabytes, make sure that
the large object data does not exceed this limit.

> **Note:**
>
> Because MySQL/MariaDB `CHAR` and `VARCHAR` cannot exceed 255
> bytes, we recommend that you use the MySQL `TEXT` type to store
> `CHAR/VARCHAR` values with a size larger than 255 bytes. When fetching
> `TEXT` columns from a MySQL database, these will be treated as
> `CHAR/VARCHAR` types by the MySQL database driver. See [CHAR/VARCHAR types](1327-char-and-varchar-data-types.md) for more details.

## Related links

**Related concepts**  

[Using portable data types](1008-using-portable-data-types.md "Only a limited set of data types are really portable across several database engines.")
