---
title: "PostgreSQL DATETIME type mapping change"
source: "fgl-topics/c_fgl_Migrate_to_300_postgres_datetime.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.00 upgrade guide > PostgreSQL DATETIME type mapping change"
type: "concept"
---

# PostgreSQL DATETIME type mapping change

> Conversion of DATETIME type with fractional seconds to PostgreSQL TIME(N)/TIMESTAMP(N) was invalid and has been reviewed.

Before Genero 3.00, the PostgreSQL driver converted `DATETIME` types as
follows:

- `DATETIME HOUR TO MINUTE` was converted to `TIMESTAMP(3) WITHOUT TIME
  ZONE`
- `DATETIME HOUR TO SECOND` was converted to `TIME(0) WITHOUT TIME
  ZONE`
- `DATETIME HOUR TO FRACTION(n)` was converted to `TIME(n+1) WITHOUT TIME
  ZONE`
- `DATETIME YEAR TO MINUTE` was converted to `TIMESTAMP(3) WITHOUT TIME
  ZONE`
- `DATETIME YEAR TO SECOND` was converted to `TIMESTAMP(3) WITHOUT TIME
  ZONE`
- `DATETIME YEAR TO FRACTION(n)` was converted to `TIMESTAMP(n+1) WITHOUT
  TIME ZONE`

Starting with Genero 3.00, when creating a table in a BDL program with CREATE TABLE, the types
are converted in a different way:

- `DATETIME HOUR TO MINUTE` is converted to `TIME(0) WITHOUT TIME
  ZONE` (seconds set to 00).
- `DATETIME HOUR TO SECOND` is converted to `TIME(0) WITHOUT TIME
  ZONE`.
- `DATETIME HOUR TO FRACTION(n)` is converted to `TIME(n) WITHOUT TIME
  ZONE`.
- `DATETIME YEAR TO MINUTE` is converted to `TIMESTAMP(0) WITHOUT TIME
  ZONE` (seconds set to 00).
- `DATETIME YEAR TO SECOND` is converted to `TIMESTAMP(0) WITHOUT TIME
  ZONE`.
- `DATETIME YEAR TO FRACTION(n)` is converted to `TIMESTAMP(n) WITHOUT TIME
  ZONE`.

This bug fix introduces an incompatibility and can have an impact on applications using
`DATETIME HOUR TO MINUTE`, `DATETIME HOUR TO FRACTION(n)` or
`DATETIME YEAR TO FRACTION(n)`. If you are using one of these types, consider
reviewing your database schema, to modify the column types according to the new SQL type conversion
rules.

> **Important:**
>
> More changes on `DATETIME` type mapping take place in [Genero BDL version 3.20](0165-datetime-sql-type-mappings.md "For some databases, the type mapping for DATETIME HOUR TO MINUTE has changed.").

## Related links

**Related concepts**  

[DATE and DATETIME data types](../10_sql-support/1432-date-and-datetime-data-types.md "DATE and DATETIME data types")

[Using portable data types](../10_sql-support/1008-using-portable-data-types.md "Only a limited set of data types are really portable across several database engines.")

[Primitive Data types](../08_language-basics/0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data.")

[Date/time literals in SQL statements](../10_sql-support/1033-date-time-literals-in-sql-statements.md "Good practices for date and time handling in SQL.")
