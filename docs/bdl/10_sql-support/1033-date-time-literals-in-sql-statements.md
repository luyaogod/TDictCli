---
title: "Date/time literals in SQL statements"
source: "fgl-topics/c_fgl_sql_programming_081.html"
breadcrumb: "SQL support > SQL programming > SQL portability > Date/time literals in SQL statements"
type: "concept"
---

# Date/time literals in SQL statements

> Good practices for date and time handling in SQL.

## Date and time strings in SQL Statements

IBM® Informix®
allows you to specify date and time values as a quoted character string in a specific format,
depending upon DBDATE and GLS environment variables. For example, if `DBDATE=DMY4`,
the following statement specifies a valid `DATE` represented by a string
literal:

```
SELECT COUNT(*) FROM table WHERE date_col = '24/12/2005'
```

Other database servers do support date/time literals as quoted character strings, but the
date/time format specification is quite different. The parameter to specify the date/time format
can be a database parameter, an environment variable, or a session option.

In order to write portable SQL, use SQL parameters instead of string literals for date-time
values:

```
DEFINE cnt INTEGER
DEFINE adate DATE
LET adate = MDY(12,24,2005)
SELECT COUNT(*) INTO cnt FROM table
    WHERE date_col = adate
```

Or, when using dynamic SQL:

```
DEFINE cnt INTEGER
DEFINE adate DATE
LET adate = MDY(12,24,2005)
PREPARE s1 FROM "SELECT COUNT(*) FROM table WHERE date_col = ?"
EXECUTE s1 USING adate INTO cnt
```

Similarly, when fetching rows from the database server into program variables, IBM
Informix allows string literals (representing date values
in DBDATE format) to be fetched into `DATE`
variables:

```
DEFINE adate DATE
SELECT '24/12/2005' INTO adate FROM ...
```

With other database servers, consider casting the original date string to a real date value, to
avoid any conversion issue at `FETCH` time.

As a general rule, always store and handle date values in `DATE` columns and
variables, on both db server or program side.

## Date-time literals

IBM Informix
`DATETIME` and `INTERVAL` literals are not converted automatically
by the SQL translator of the database driver:

```
SELECT COUNT(*) FROM order WHERE ord_when > DATETIME (1999-10-12) YEAR TO DAY
```

Check your code, to detect where you are using such expressions in the SQL statements, and use
an SQL parameter instead.

## Informix-specific date/time keywords

SQL statements using expressions such as `TODAY`, `CURRENT`, and
`EXTEND` are specific to Informix SQL.

Database drivers try to translate date/time constant expressions to native SQL syntax, but this
is only provided to simplify migration.

Date/time expression translation can be controlled with the following FGLPROFILE
entries:

```
dbi.database.dsname.ifxemul.today = { true | false }
dbi.database.dsname.ifxemul.current = { true | false }
dbi.database.dsname.ifxemul.extend = { true | false }
```

> **Important:**
>
> To ease migration to a new database type, Informix-specific expressions such
> as `TODAY`, `CURRENT` and `EXTEND` are converted to
> native date/time expressions. However, the date/time returned by the native SQL function may use a
> different timezone / daylight saving time convention.

Check your code, to detect where you are using `TODAY/CURRENT/EXTEND` expressions
in the SQL statements, and consider using SQL parameters with program variables assigned with the
`TODAY/CURRENT/EXTEND` instruction of Genero BDL.

## Date-time expressions with parameters

Date-time arithmetic expressions using SQL parameters (USING variables) are not portable.

For example:

```
PREPARE s1 FROM "SELECT ... WHERE datecol < ? + 1"
```

Might generate an error with non-Informix databases.

## DATEs as a number of days

IBM Informix can
automatically convert integers to a DATE values, as a number of days since 12/31/1899 ( 1 = 01/01/1900 ).
This is however not supported by other database engines.

Check your code, to detect where you are using integers with DATE columns.

## Related links

**Related concepts**  

[DBDATE](../07_configuration/0508-dbdate.md "Defines the default display and input format for DATE values.")
