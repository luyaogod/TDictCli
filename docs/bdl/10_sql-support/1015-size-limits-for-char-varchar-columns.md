---
title: "Size limits for CHAR/VARCHAR columns"
source: "fgl-topics/c_fgl_sql_programming_066.html"
breadcrumb: "SQL support > SQL programming > SQL portability > CHAR and VARCHAR types > Size limits for CHAR/VARCHAR columns"
type: "concept"
---

# Size limits for CHAR/VARCHAR columns

> Each database brand defines its own limits for CHAR/VARCHAR SQL types.

## CHAR and VARCHAR SQL types

To write portable SQL, you must use `CHAR` and `VARCHAR` type sizes
that does not exceed the limits of the target database brands (Informix, Oracle, SQL Server, etc).

The `CHAR` type is a blank-padded type: The string "abc" in a
`CHAR(200)` column will in fact get 197 additional trailing blancs. The
`VARCHAR` type is designed to store variable-length character strings. Thus, the
`CHAR(N)` type should be used for fixed-length string data (such
as `"ABC-123"` codes), while `VARCHAR(N)` should be
used for variable-length string data (such as names and descriptions). Furthermore, with some
database brands, in `VARCHAR` values, trailing spaces are [semantically significant](1017-trailing-blanks-in-char-varchar.md "How to cope with trailing blanks in CHAR(N) and VARCHAR(N) SQL columns and program variables?"), while trailing blanks of
`CHAR` SQL values are not significant.

With Informix, the `VARCHAR` SQL type was designed with a limit of 255 bytes. On
the other hand, the `CHAR` type has a much larger limit of 32767 bytes. In fact, it
should have been the other way around, since the blank-padded `CHAR` type is rather
for small fix-length data.

For large variable-length character strings, Informix has introduced the
`LVARCHAR` type, first with a limit of 2048 bytes and in younger IDS versions with a
limit of 32739 bytes.

Most database brands supported by Genero have a higher limit for `VARCHAR` than
the `CHAR` type. For example, Oracle DB `CHAR` has a limit of 2000,
while `VARCHAR2` has a limit of 4000 (bytes or chars, depending on the [length semantics](1013-byte-or-character-length-semantics.md "Length Semantics defines the unit used to express the length of a character string, the position of a given character, and the size of a character data type.")).

The Genero BDL `VARCHAR` type size limit is 65524 (like the `CHAR`
BDL type), and thus can store Informix `LVARCHAR` or any database brand
`VARCHAR` column data.

## What CHAR length should be used?

To store short and fix-length character data, use a `CHAR(N)`
column, where n will typically be 10 to 100 (bytes or characters): All database
brands supported by Genero can support such `CHAR` size.

## What VARCHAR length should be used?

To store variable-length character strings that do not exceed 255 bytes, use the
`VARCHAR(N)` type: All database brands supported by Genero can
support such `VARCHAR` size.

To store variable-length character strings that can exceed 255 bytes:

- If you target only non-Informix databases, use `VARCHAR(2048)` type in the
  `CREATE TABLE` / `ALTER TABLE` statements in BDL code: That will
  compile (no Informix 255 limit) and the SQL will also execute.
- If you want to support all databases including Informix, use
  `LVARCHAR(N)` type in `CREATE TABLE` /
  `ALTER TABLE` statements in BDL code: The `LVARCHAR` Informix-specific
  type will be converted to `VARCHAR(N)` with other database
  brands.

## Related links

**Related concepts**  

[Length semantics settings](../09_advanced-features/0881-length-semantics-settings.md "Length semantics settings")
