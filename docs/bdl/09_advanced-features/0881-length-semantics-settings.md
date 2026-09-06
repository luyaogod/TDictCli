---
title: "Length semantics settings"
source: "fgl-topics/c_fgl_localization_022.html"
breadcrumb: "Advanced features > Localization > Application locale > Defining the application locale > Length semantics settings"
type: "concept"
description: "Understanding length semantics The length semantics of character string data matters when using a multibyte character set. Length semantics involves data type length specification for database column ..."
---

# Length semantics settings

## Understanding length semantics

The length semantics of character string data matters when using a multibyte
character set. Length semantics involves data type length specification for database
column and program variable definitions, as well as string manipulations (for string
lengths, character positions, offsets and substring ranges).

In a single-byte characters set like ISO-8859-1, a character is encoded on one byte.
The length of a string can be counted in bytes or characters, the unit does not
matter. In other words, the length semantics is identical in bytes or characters,
with a single byte encoding. However, with a multibyte character set like UTF-8 or
BIG5, a character can be encoded on several bytes. In such case, the unit regarding
length semantics matters, because the number of bytes of a character string can be
different from the number of characters.

For multibyte characters sets, the language supports Byte Length
Semantics (BLS) and Character Length Semantics (CLS) specification. BLS or CLS
usage depends on the current character set of the application. BLS is typically used with a
character set such as BIG5, because for historical reasons programmers are used to counting 2 bytes
for each Chinese character. For UTF-8, which is a variable size encoding, the recommendation is to use
CLS instead. CLS simplifies data type definition and string handling when using UTF-8.

Programming areas concerned by length semantics are illustrated in the following code
example:

```
SCHEMA shop

# CREATE TABLE mytable (
#     k INT,
#     vc VARCHAR(10)
#         -- what is the unit for the column size and how many
#         -- characters can be stored in this column?
#  )

MAIN
  DEFINE buf, tmp VARCHAR(50)  -- what is the unit for the size?
  DEFINE rec RECORD LIKE mytable  -- what is the size of vc member?
  DEFINE str STRING, len INT

  DATABASE shop

  SELECT LENGTH(vc) INTO len -- What unit use string functions in SQL?
     FROM mytable WHERE k = 45

  LET buf = "abcdef..."  -- How many chars can this variable hold?

  DISPLAY length(buf)  -- In what unit is the length expressed?

  LET tmp = buf[1,5]  -- What is the unit for char positions?

  LET str = buf
  DISPLAY str.getLength() -- What is the unit for the length?
  DISPLAY str.getIndexOf("def") -- What is the unit for the offset?

END MAIN
```

## Using Byte Length Semantics

Byte Length Semantics must be used if the current locale defines a multibyte character set
different from UTF-8.

> **Important:**
>
> - Byte Length Semantics is the default on UNIX™ and
>   Windows® platforms.
> - Byte Length Semantics cannot be set on mobile platforms.

With BLS, the size of CHAR/VARCHAR program variables is expressed in byte units.
In a single-byte character set like ISO-8859-1, every character is encoded on a unique byte, so the
number of bytes equals the number of characters. When using BLS with a multibyte character set, you
must be aware of the storage size in byte units: Character encoding requires more than one byte, so
the number of bytes to store a multibyte string is greater than the number of characters.
For example, in a BIG5 encoding, one Chinese character needs 2 bytes, so if you want to hold a BIG5
string with a maximum of 10 Chinese characters, you must define a CHAR(20). When using UTF-8,
characters can take one or several bytes which can use two or three times more storage space than
character count. You need to choose the right expansion factor to define CHAR or VARCHAR variables
in byte units.

```
-- Using Byte Length Semantics
DEFINE code VARCHAR(10)  -- Can store 10 bytes / 10 single-byte chars.
```

In order to use BLS, you can define the [FGL\_LENGTH\_SEMANTICS](../07_configuration/0518-fgl-length-semantics.md "Defines the length semantics to be used in programs.") environment variable
to "BYTE", or just leave it unset, if BLS is the default on your platform. For example, on UNIX:

```
$ FGL_LENGTH_SEMANTICS="BYTE"
$ export FGL_LENGTH_SEMANTICS
```

## Using Char Length Semantics

Character Length Semantics is used with multibyte character sets such as UTF-8:
Migrating to UTF-8 by using CLS will allow you to leave the source code untouched,
even when doing complex string/substring manipulations.

The database typically also uses UTF-8 and CLS. If the database uses UTF-8 and only
supports BLS, programs can still use CLS with UTF-8.

> **Important:**
>
> Char Length Semantics is the default on iOS and Android™ mobile platforms, and cannot be changed (Byte Length Semantics
> cannot be used on mobile: only UTF-8 character set is allowed).

With CLS, the size of a CHAR/VARCHAR program variable is expressed in character
units, and the number of bytes needed to store these characters is allocated
automatically. A VARCHAR(10) variable will hold 10 characters, of any byte length.
Furthermore, language functions and class methods dealing with character string length
and positions will use character units.

```
-- Using Character Length Semantics
DEFINE name VARCHAR(10)  -- Can store 10 chars in UTF-8, or any encoding.
LET name = "Forêt"  -- 5 chars, that take 6 bytes in UTF-8
DISPLAY length(name)  -- Displays a length of 5 (characters)
DISPLAY "[",name[4,5],"]"    -- Displays [êt]
```

To enable Char Length Semantics, define the [FGL\_LENGTH\_SEMANTICS](../07_configuration/0518-fgl-length-semantics.md "Defines the length semantics to be used in programs.") environment variable
to "CHAR". For example, on UNIX:

```
$ FGL_LENGTH_SEMANTICS="CHAR"
$ export FGL_LENGTH_SEMANTICS
```

## Length Semantics in SQL

On the database server side, the length semantics used for character data types
varies from one vendor to another. Some databases use BLS, other use CLS, and other support both
semantics.
> **Important:**
>
> The length semantics used by the Genero runtime (defined by
> FGL\_LENGTH\_SEMANTICS) should match the length semantics used by the database server: If a database
> column is defined as a CHAR of 10 bytes, the corresponding BDL program variable should also be
> defined as a CHAR of 10 bytes (using FGL\_LENGTH\_SEMANTICS=BYTE, the default). If the database column
> is defined as a CHAR of 10 characters, the variable should be defined as a CHAR of 10 characters
> (using FGL\_LENGTH\_SEMANTICS=CHAR).

This table shows the character data type length semantics of supported database
servers:

| Database Engine | Length semantics in character data types | Summary |
| --- | --- | --- |
| Oracle® | Supports both Byte or Character Length Semantics in character type definition, can be defined globally for the database or at column level (with the `CHAR(10 BYTE\|CHAR)` syntax).Character string data is stored in database character set for CHAR /VARCHAR columns and in national character set for NCHAR /NVARCHAR columns.See [Oracle DB Guide](../10_sql-support/1373-char-and-varchar-data-types.md) for more details. | BLS/CLS |
| Informix® | Uses Byte Length Semantics for the size of character columns. Can apply a ratio when creating columns, based on the SQL\_LOGICAL\_CHAR server configuration parameter.Character string data is stored in the database character set defined by DB\_LOCALE. | BLS |
| Microsoft® SQL Server | CHAR / VARCHAR sizes are specified in bytes; data is stored in the character set defined by the database collation.NCHAR / NVARCHAR sizes are specified in characters; data is stored in UCS-2.See [Microsoft SQL Server Guide](../10_sql-support/1273-char-and-varchar-data-types.md) for more details. | BLS/CLS |
| PostgreSQL | Uses Character Length Semantics for the size of character columns.Character string data is stored in the database character set defined by WITH ENCODING of CREATE DATABASE. | CLS |
| Oracle® MySQL / Maria DB | Uses Character Length Semantics for the size of character columns.Character string data is stored in the server character set defined by a configuration parameter.See [Oracle MySQL / MariadDB Guide](../10_sql-support/1327-char-and-varchar-data-types.md) for more details. | CLS |
| SQLite | Uses Character Length Semantics for the size of character columns.Character string data is stored in UTF-8.See [SQLite Guide](../10_sql-support/1479-char-and-varchar-data-types.md) for more details. | CLS |

Other SQL elements like functions and operators are affected by the length semantic.
For example, Informix LENGTH() function
always returns a number of bytes, while Oracle's LENGTH() function returns a number
of characters (use LENGTHB() to get the number of bytes with Oracle).

It is important to understand properly how the database servers handle multibyte
character sets. Check your database server reference manual. In most documentations you
will find a "Localization" chapter which describes those concepts in detail.

## Extracting database schemas

[Database schema files](0791-database-schema.md "Defines database table structures with column type information to be reused in program variable definitions.")
(.sch) are used to resolve column data types when compiling
.4gl modules and .per form files. This file contains size
information for CHAR and VARCHAR types. It is important to identify the unit used by the database
columns, to properly define CHAR/VARCHAR variables in programs and fields in forms.

Most database engines (like Oracle DB, SQL Server, PostgreSQL, SQLite) provide
catalog tables with column size information in character units. In this case, the
fgldbsch tool extracts the column sizes in character units, without further
conversion. If the column sizes is provided in bytes by catalog tables, fgldbsch
will try to detect character length semantic usage in the database and apply a reduction factor to
convert the number of bytes to chars.

For example, with Informix, when using the
SQL\_LOGICAL\_CHAR onconfig parameter, fgldbsch will convert the size stored in
bytes in `syscolumns.collength` column, to a number of characters, by dividing that
number of bytes by the value of SQL\_LOGICAL\_CHAR.

As a result - independently from the length semantics used in your programs - the
CHAR/VARCHAR type sizes in the schema file are always expressed in character units. When using Byte
Length Semantics, this makes no difference in a single-byte locale, because one character occupies a
single byte. In a multibyte encoding (UTF-8) with BLS, this method guarantees that the program
variable will not hold more ASCII characters than the database column can hold. When using Character
Length Semantics with a multibyte character set, the size in characters will define character type
variables in the same unit.

For example, with BLS, a VARCHAR(10 *(bytes or chars)*) column will define a VARCHAR(10
*(bytes)*) in programs. With CLS, a VARCHAR(10 *(chars)*) column will define a VARCHAR(10
*(chars)*) in programs.

## Moving from single-byte to UTF-8

Migration to Unicode (UTF-8) is facilitated with Char Length Semantics:

1. Verify that your database uses Char Length Semantics.
2. Convert your sources and string files from your single-byte locale to UTF-8 (iconv).
3. Enable Char Length Semantics with FGL\_LENGTH\_SEMANTICS=CHAR.
4. Compile and run your programs untouched.

## Related links

**Related concepts**  

[SQL character type for Unicode/UTF-8](../10_sql-support/1014-sql-character-type-for-unicode-utf-8.md "This section explains database server specifics regarding Unicode / UTF-8 support with character string SQL types.")

[Multibyte character sets (MBCS)](0873-multibyte-character-sets-mbcs.md "Multibyte character sets (MBCS)")

[Single-byte character sets (SBCS)](0871-single-byte-character-sets-sbcs.md "Single-byte character sets (SBCS)")

[Character size unit and length semantics](0874-character-size-unit-and-length-semantics.md "Character size unit and length semantics")

[Manipulating character strings](0976-manipulating-character-strings.md "Manipulating character strings")
