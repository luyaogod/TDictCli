---
title: "CHAR and VARCHAR types"
source: "fgl-topics/c_fgl_sql_programming_063.html"
breadcrumb: "SQL support > SQL programming > SQL portability > CHAR and VARCHAR types"
type: "concept"
---

# CHAR and VARCHAR types

> Using the CHAR and VARCHAR data types with different databases.

The `CHAR` and `VARCHAR` types are designed to store character
strings, but all database servers do not have the same semantics for these types.

The maximum size, supported characters sets and length semantics of `CHAR` and
`VARCHAR` types can be very different from one database system to another. Consider
using character types and sizes that are common across all the database systems you target.

> **Important:**
>
> Most database engine brands allow to create table columns using a different character encoding
> than the database. This is not supported with Genero: All table columns must use the same character
> encoding defined at the database level.

The behavior of database servers may differ in the following areas related to
`CHAR`/`VARCHAR` types.

- [Byte or Character Length semantics?](1013-byte-or-character-length-semantics.md "Length Semantics defines the unit used to express the length of a character string, the position of a given character, and the size of a character data type.")
- [SQL character type for Unicode/UTF-8](1014-sql-character-type-for-unicode-utf-8.md "This section explains database server specifics regarding Unicode / UTF-8 support with character string SQL types.")
- [Size limits for CHAR/VARCHAR columns](1015-size-limits-for-char-varchar-columns.md "Each database brand defines its own limits for CHAR/VARCHAR SQL types.")
- [Empty strings and NULLs](1016-empty-strings-and-nulls.md "Depending on the context, an empty string ( '' ) can be considered as NULL or NOT NULL.")
- [Trailing blanks in CHAR/VARCHAR](1017-trailing-blanks-in-char-varchar.md "How to cope with trailing blanks in CHAR(N) and VARCHAR(N) SQL columns and program variables?")
- [What should you do?](1018-what-should-you-do.md "This section contains facts and tips to consider regarding character data types and locale settings.")

| Database Server Type | Character types topic |
| --- | --- |
| IBM® Informix® | [CHAR/VARCHAR in IBM Informix](1186-what-are-the-supported-ibm-informix-sql-features.md) |
| Microsoft™ SQL Server | [CHAR/VARCHAR in SQL Server](1273-char-and-varchar-data-types.md) |
| Oracle® MySQL / MariadDB | [CHAR/VARCHAR in Oracle MySQL](1327-char-and-varchar-data-types.md) |
| Oracle Database Server | [CHAR/VARCHAR in Oracle DB](1373-char-and-varchar-data-types.md) |
| PostgreSQL | [CHAR/VARCHAR in PostgreSQL](1430-char-and-varchar-data-types.md) |
| SQLite | [CHAR/VARCHAR in SQLite](1479-char-and-varchar-data-types.md) |
| Dameng® | [CHAR/VARCHAR in Dameng](1228-char-and-varchar-data-types.md) |

## Child topics

- [Byte or Character Length semantics?](1013-byte-or-character-length-semantics.md): Length Semantics defines the unit used to express the length of a character string, the position of a given character, and the size of a character data type.
- [SQL character type for Unicode/UTF-8](1014-sql-character-type-for-unicode-utf-8.md): This section explains database server specifics regarding Unicode / UTF-8 support with character string SQL types.
- [Size limits for CHAR/VARCHAR columns](1015-size-limits-for-char-varchar-columns.md): Each database brand defines its own limits for CHAR/VARCHAR SQL types.
- [Empty strings and NULLs](1016-empty-strings-and-nulls.md): Depending on the context, an empty string ( '' ) can be considered as NULL or NOT NULL.
- [Trailing blanks in CHAR/VARCHAR](1017-trailing-blanks-in-char-varchar.md): How to cope with trailing blanks in CHAR(N) and VARCHAR(N) SQL columns and program variables?
- [What should you do?](1018-what-should-you-do.md): This section contains facts and tips to consider regarding character data types and locale settings.
