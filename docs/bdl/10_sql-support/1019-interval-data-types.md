---
title: "INTERVAL data types"
source: "fgl-topics/c_fgl_sql_programming_interval.html"
breadcrumb: "SQL support > SQL programming > SQL portability > INTERVAL data types"
type: "concept"
---

# INTERVAL data types

> Not all database brands support a native SQL type to store time duration.

The [`INTERVAL`](../08_language-basics/0563-interval-qual1-to-qual2.md "The INTERVAL data type stores spans of time as Year/Month or Day/Hour/Minute/Second/Fraction units.") data type in
Genero BDL is used to store an amount of time.

There are two classes of `INTERVAL` types:

- year-month interval, to represent a number of years and months like 9345 years, 10
  months.
- day-second interval, to represent a number of days and hours/minutes/seconds, such
  as 34 days, 5 hours, 20 minutes and 30.052 seconds.

While most database brands provide data types to store date/time information (as a point in
time), only a few have a real native SQL type to store time duration.

If the database engine cannot store real interval data, the Genero ODI driver will use a
`CHAR(50)` type, to provide a storage solution for `INTERVAL`
variables. However, with interval data stored in a character string column, we lose the features
offered by SQL to compute, filter, convert, and order time duration data in an efficient way. For
example, a table index is more efficient, when the underlying column type corresponds to the actual
stored data.

For maximum SQL portability, only use `INTERVAL` data, if the target database
engine provides a native SQL type to store such data.

Consider also to use the `INTERVAL` qualifiers to match the native interval
type:

- `INTERVAL YEAR(p) TO MONTH`
- `INTERVAL DAY(p) TO MINUTE`
- `INTERVAL DAY(p) TO SECOND`
- `INTERVAL DAY(p) TO FRACTION(5)`

| Database Server Type | Native interval type(s)? |
| --- | --- |
| IBM® Informix® | Yes, Genero BDL is based on Informix SQL data types... |
| Microsoft™ SQL Server | No, see [`INTERVAL` with Microsoft SQL Server](1276-interval-data-type.md) |
| Oracle® MySQL / MariadDB | No, see [`INTERVAL` with Oracle MySQL/MariaDB](1330-interval-data-type.md) |
| Oracle Database Server | Yes, see [`INTERVAL` with Oracle DB](1376-interval-data-type.md) |
| PostgreSQL | Yes, see [`INTERVAL` with PostgreSQL](1433-interval-data-type.md) |
| SQLite | No, see [`INTERVAL` with SQLite](1482-interval-data-type.md) |
| Dameng® | Yes, see [`INTERVAL` with Dameng](1231-interval-data-type.md) |

## Related links

**Related concepts**  

[Primitive Data types](../08_language-basics/0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data.")

[Using portable data types](1008-using-portable-data-types.md "Only a limited set of data types are really portable across several database engines.")
