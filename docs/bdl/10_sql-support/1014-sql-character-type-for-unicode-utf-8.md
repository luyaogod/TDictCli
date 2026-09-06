---
title: "SQL character type for Unicode/UTF-8"
source: "fgl-topics/c_fgl_sql_programming_unicode.html"
breadcrumb: "SQL support > SQL programming > SQL portability > CHAR and VARCHAR types > SQL character type for Unicode/UTF-8"
type: "concept"
---

# SQL character type for Unicode/UTF-8

> This section explains database server specifics regarding Unicode / UTF-8 support with character string SQL types.

All database servers can store UNICODE data in character strings types, but there are some
specifics you must be aware of. Genero BDL programs typically use the `CHAR`
and `VARCHAR` types to store UTF-8 strings. But the corresponding SQL type may
have a different name, depending on the database server type. Use the correct SQL type when
creating your database tables. When the database uses a different UNICODE codeset as UTF-8 to
store the character string data, the database client or the Genero database driver take care
of the codeset conversion, as long as the runtime system and database client locale are
properly defined.

| Database Server Type | Char types to be used for Unicode/UTF-8 |
| --- | --- |
| IBM® Informix® | `CHAR` / `VARCHAR`, the database must be created with UTF-8 locale. |
| Microsoft™ SQL Server | `NCHAR` / `NVARCHAR`, to store UTF-16 data (drivers make the conversion for application codeset UTF-8)The `CHAR` / `VARCHAR` types typically store non-unicode data (using a single byte character set). Starting with SQL Server 2019, the database collation can have a \_UTF8 modifier to store UTF-8 in `CHAR` / `VARCHAR` columns.For more details, see [the topic about character data type usage with Microsoft SQL Server](1273-char-and-varchar-data-types.md). |
| Oracle® MySQL | `CHAR` / `VARCHAR` if the database locale is UTF-8.`NCHAR` / `NVARCHAR` if you need to use the national character set.For more details, see [the topic about character data type usage with Oracle MySQL](1327-char-and-varchar-data-types.md). |
| Oracle Database Server | `CHAR` / `VARCHAR2` if the database locale is UTF-8.`NCHAR` / `NVARCHAR2` if you need to use the national character set.For more details, see [the topic about character data type usage with Oracle DB](1373-char-and-varchar-data-types.md). |
| PostgreSQL | `CHAR` / `VARCHAR`, the database locale must be UTF-8.For more details, see [the topic about character data type usage with PostgreSQL](1430-char-and-varchar-data-types.md). |
| SQLite | `CHAR` / `VARCHAR` (data always stored in UTF-8).For more details, see [the topic about character data type usage with SQLite](1479-char-and-varchar-data-types.md). |
| Dameng® | `CHAR` / `VARCHAR`, the database locale must be UTF-8.For more details, see [the topic about character data type usage with Dameng](1228-char-and-varchar-data-types.md). |

## Related links

**Related concepts**  

[Language and character set settings](../09_advanced-features/0880-language-and-character-set-settings.md "Language and character set settings")

[Database client settings](../09_advanced-features/0885-database-client-settings.md "This section describes the settings defining the locale for the database client.")
