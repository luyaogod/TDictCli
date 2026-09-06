---
title: "MySQL DATETIME fractional seconds"
source: "fgl-topics/c_fgl_Migrate_to_300_mysql_dt_frac.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.00 upgrade guide > MySQL DATETIME fractional seconds"
type: "concept"
---

# MySQL DATETIME fractional seconds

> MySQL 5.6.4 TIME and DATETIME types support fractions of seconds that can be used to store DATETIME HOUR TO FRACTION(N) or DATETIME YEAR TO FRACTION(N).

Before Genero 3.00, the Oracle® MySQL
driver converted `DATETIME` types as follows:

- `DATETIME HOUR TO SECOND` was converted to MySQL `TIME`.
- Other `DATETIME` types were converted to MySQL `DATETIME`.

Starting with Genero 3.00, when creating a table in a BDL program with the `CREATE
TABLE` statement, if the MySQL server version is greater or equal to 5.6.4, the types are
converted differently, as follows:

- `DATETIME HOUR TO MINUTE` is converted to MySQL `TIME` (seconds
  set to 00).
- `DATETIME HOUR TO SECOND` is converted to MySQL `TIME`.
- `DATETIME HOUR TO FRACTION(n)` is converted to MySQL
  `TIME(n)`.
- `DATETIME YEAR TO MINUTE` is converted to MySQL `DATETIME`
  (seconds set to 00).
- `DATETIME YEAR TO SECOND` is converted to MySQL `DATETIME`.
- `DATETIME YEAR TO FRACTION(n)` is converted to MySQL
  `DATETIME(n)`.

This change has no impact your application when using `DATETIME HOUR TO SECOND` or
`DATETIME YEAR TO SECOND`.

However, it is now possible to store `DATETIME HOUR TO FRACTION(n)` and
`DATETIME YEAR TO FRACTION(n)` data. The `DATETIME YEAR TO
FRACTION(n)` is typically used to implement data modification timestamps to track user
changes.

> **Important:**
>
> More changes on `DATETIME` type mapping take place in [Genero BDL version 3.20](0165-datetime-sql-type-mappings.md "For some databases, the type mapping for DATETIME HOUR TO MINUTE has changed.").

## Related links

**Related concepts**  

[DATE and DATETIME data types](../10_sql-support/1329-date-and-datetime-data-types.md "DATE and DATETIME data types")

[Using portable data types](../10_sql-support/1008-using-portable-data-types.md "Only a limited set of data types are really portable across several database engines.")

[Primitive Data types](../08_language-basics/0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data.")

[Date/time literals in SQL statements](../10_sql-support/1033-date-time-literals-in-sql-statements.md "Good practices for date and time handling in SQL.")
