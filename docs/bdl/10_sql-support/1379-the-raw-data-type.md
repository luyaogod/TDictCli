---
title: "The RAW data type"
source: "fgl-topics/c_fgl_odiagora_raw_type.html"
breadcrumb: "SQL support > SQL database guides > Oracle® Database > Data dictionary > The RAW data type"
type: "concept"
description: "ORACLE supports the RAW data type to hold binary data. This data type is for example used to return values from the SYS_GUID() SQL function. Solution The ORACLE RAW values can be converted to a ..."
---

# The RAW data type

ORACLE supports the `RAW` data type to hold binary data. This data type is
for example used to return values from the `SYS_GUID()` SQL function.

## Solution

The ORACLE `RAW` values can be converted to a character string in the hexadecimal
notation.

When fetching rows from the database, the database driver will automatically convert ORACLE
`RAW` values to hexadecimal. On the other hand, when using SQL parameters, the
database driver will convert hexadecimal `VARCHAR` strings to binary data.

Since each byte is represented with two characters in the hexadecimal notation, you must
define a `VARCHAR(N*2)` variable to hold the values of a native
`RAW(N)` column.

When extracting a database schema with the fgldbsch tool, the ORACLE
`RAW(N)` type is converted to `VARCHAR2(N*2)`.
