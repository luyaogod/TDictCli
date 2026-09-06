---
title: "Large OBjects (LOBs) data types"
source: "fgl-topics/c_fgl_sql_programming_096.html"
breadcrumb: "SQL support > SQL programming > SQL portability > Large OBjects (LOBs) data types"
type: "concept"
---

# Large OBjects (LOBs) data types

> Use TEXT and BYTE FGL types to store database character and binary large objects.

Genero BDL provides the [`TEXT`](../08_language-basics/0569-text.md "The TEXT data type stores large text data.") data
type to store character LOBs and the [`BYTE`](../08_language-basics/0555-byte.md "The BYTE data type stores any type of binary data, such as images or sounds.") data type to store binary LOBs.

The `TEXT` and `BYTE` types can be used to handle large objects
data from various database brands. However, some database types are not supported.

| Database Server Type | LOBs support with TEXT/BYTE |
| --- | --- |
| IBM® Informix® | Yes, [native SQL feature](1186-what-are-the-supported-ibm-informix-sql-features.md) |
| Microsoft™ SQL Server | Yes, [see details](1282-text-and-byte-lob-types.md) |
| Oracle® MySQL / MariadDB | Yes, [see details](1333-text-and-byte-lob-types.md) |
| Oracle Database Server | Yes, [see details](1380-text-and-byte-lob-types.md) |
| PostgreSQL | Yes, [see details](1436-text-and-byte-lob-types.md) |
| SQLite | Yes, [see details](1486-text-and-byte-lob-types.md) |
| Dameng® | Yes, [see details](1234-text-and-byte-lob-types.md) |
