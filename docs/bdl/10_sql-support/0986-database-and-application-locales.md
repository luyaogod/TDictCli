---
title: "Database and application locales"
source: "fgl-topics/c_fgl_sql_programming_locale.html"
breadcrumb: "SQL support > SQL programming > SQL basics > Database and application locales"
type: "concept"
---

# Database and application locales

> Choose the right locale for your database and application programs.

When designing a database and the application programs that connect to this database, it is
important to consider what character set will be used in the database (server-side), and the
character set used by database client application programs.

> **Important:**
>
> Most database engine brands allow to create table columns using a different character encoding
> than the database. This is not supported with Genero: All table columns must use the same character
> encoding defined at the database level.

Most database servers support character set conversion between the client application and the
database server. Thus, it is possible to use a different character set in the application and in the
database, for example:

1. Database uses ISO88591, application programs use ISO88591
2. Database uses UTF-8, application programs use ISO88591
3. Database uses UTF-8, application programs use UTF-8

However, best practice is to use the same character set encoding in the application programs and
in the database.

> **Important:**
>
> The locale definition of the database client software must match the application locale defined
> for Genero. For more details, see [Database client settings](../09_advanced-features/0885-database-client-settings.md "This section describes the settings defining the locale for the database client.").

Another important concept to be considered is the length semantics that will be used
in the application programs and in the database. The length semantic defines if a CHAR(10) can hold
10 bytes or 10 characters (using X bytes).

For a detailed discussion about character sets and length semantics, see [CHAR and VARCHAR types](1012-char-and-varchar-types.md "Using the CHAR and VARCHAR data types with different databases.").

## Related links

**Related concepts**  

[Application locale](../09_advanced-features/0864-application-locale.md "The application locale defines the language and codeset for your application.")
