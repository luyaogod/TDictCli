---
title: "INTERVAL data type"
source: "fgl-topics/c_fgl_odiagpgs_038.html"
breadcrumb: "SQL support > SQL database guides > PostgreSQL > Data dictionary > INTERVAL data type"
type: "concept"
description: "Informix® Informix provides the INTERVAL data type to store a value that represents a span of time. INTERVAL types are divided into two classes: Intervals of year-month class such as INTERVAL YEAR(5) ..."
---

# INTERVAL data type

## Informix®

Informix provides the `INTERVAL` data type to store a value that represents a
span of time.

`INTERVAL` types are divided into two classes:

- Intervals of year-month class such as `INTERVAL YEAR(5) TO
  MONTH`
- Intervals of day-time class such as `INTERVAL DAY(9) TO SECOND`

`INTERVAL` columns can be defined with various time units, by specifying a
start and end qualifier. For example, you can define an interval to store a number of hours and
minutes with `INTERVAL HOUR(n) TO MINUTE`, where
n defines the maximum number of digits for the hours unit.

The values of Informix `INTERVAL` can be represented with a
character string literal, or as `INTERVAL()`
literals:

```
'-9834 15:45:12.345'  -- an INTERVAL DAY(6) TO FRACTION(3)
'7623-11'   -- an INTERVAL YEAR(9) TO MONTH
INTERVAL(18734:45) HOUR(5) TO MINUTE
INTERVAL(-7634-11) YEAR(5) TO MONTH
```

## PostgreSQL

PostgreSQL provides an `INTERVAL` data type which is equivalent to the Informix `INTERVAL`:

- It is possible to specify the interval class / precision with `YEAR`,
  `MONTH`, `DAY`, `HOUR`, `MINUTE` and
  `SECOND[(p)]` fields.
- Fractional part of seconds can be defined with up to 6 digits.
- The interval value range is from -178000000 to +178000000 years.
- Input and output format can be controlled with the `SET interval style`
  command.

Unlike Informix `INTERVAL`, the PostgreSQL `INTERVAL` type can be
used without any qualifiers when defining table columns. The qualifiers are mostly specified as an
indicator, but are not strick interval type specification as in
Informix:

```
test1=> create table mytable ( i interval hour to minute );
CREATE TABLE
test1=> insert into mytable values ( interval '9999-10 555 11:22:33' );
INSERT 0 1
test1=> select * from mytable;
                  i                   
--------------------------------------
 9999 years 10 mons 555 days 11:22:00
(1 row)
```

When selecting directly columns defined with interval qualifiers, PostgreSQL provides detail type
information. However, when using aggregate SQL functions such as `SUM()` on interval
columns, the resulting type will be a simple INTERVAL type without
details:

```
create table tab1 ( pkey int, dur interval hour to minute);
\d tab1

insert into tab1 values ( 101, interval '999:12');
insert into tab1 values ( 102, interval '100:00');

create view v1 as select sum(dur) from tab1;
select * from v1;
\d v1

create view v2 as select cast(sum(dur) as interval hour to minute) from tab1;
select * from v2;
\d v2

drop view v1;
drop view v2;
drop table tab1;
```

With the above SQL code, the `\d v1` command which is describing the column types
of view `v1`, will show the `"interval"` type name, without further
type info.

## Solution

Use the following conversion rules to map Informix numeric types to PostgreSQL numeric types:

| Informix data type | PostgreSQL data type |
| --- | --- |
| `INTERVAL YEAR[(p)] TO MONTH` | `INTERVAL YEAR TO MONTH` |
| `INTERVAL YEAR[(p)] TO YEAR` | `INTERVAL YEAR` |
| `INTERVAL MONTH[(p)] TO MONTH` | `INTERVAL MONTH` |
| `INTERVAL DAY[(p)] TO FRACTION(n)` | `INTERVAL DAY TO SECOND(n)` |
| `INTERVAL DAY[(p)] TO SECOND` | `INTERVAL DAY TO SECOND(0)` |
| `INTERVAL DAY[(p)] TO MINUTE` | `INTERVAL DAY TO MINUTE` |
| `INTERVAL DAY[(p)] TO HOUR` | `INTERVAL DAY TO HOUR` |
| `INTERVAL DAY[(p)] TO DAY` | `INTERVAL DAY` |
| `INTERVAL HOUR[(p)] TO FRACTION(n)` | `INTERVAL HOUR TO SECOND(n)` |
| `INTERVAL HOUR[(p)] TO SECOND` | `INTERVAL HOUR TO SECOND(0)` |
| `INTERVAL HOUR[(p)] TO MINUTE` | `INTERVAL HOUR TO MINUTE` |
| `INTERVAL HOUR[(p)] TO HOUR` | `INTERVAL HOUR` |
| `INTERVAL MINUTE[(p)] TO FRACTION(n)` | `INTERVAL MINUTE TO SECOND(n)` |
| `INTERVAL MINUTE[(p)] TO SECOND` | `INTERVAL MINUTE TO SECOND(0)` |
| `INTERVAL MINUTE[(p)] TO MINUTE` | `INTERVAL MINUTE` |
| `INTERVAL SECOND[(p)] TO FRACTION(n)` | `INTERVAL SECOND(n)` |
| `INTERVAL SECOND[(p)] TO SECOND` | `INTERVAL SECOND(0)` |
| `INTERVAL FRACTION TO FRACTION(n)` | `INTERVAL SECOND(n)` |

The PostgreSQL database interface converts the Informix-style `INTERVAL` type to
the native PostgreSQL `INTERVAL` type.

> **Important:**
>
> The PostgreSQL database driver forces the interval style session parameter
> to 'iso\_8601', this is required to insert and fetch interval database with the libpq CAPI functions.
> You must not change this setting during program execution.

While PostgreSQL intervals support up to 9 digits for the higher unit like Informix, year values range from -178000000 to +178000000 only. This
limitation exists in PostgreSQL 8.4 and may be solved in future versions.

When using aggregate SQL functions such as `SUM()`, since PostgreSQL does not
provide detailed type information, the ODI driver must select a corresponding FGL interval type,
which will be `INTERVAL YEAR(9) TO MONTH`. Fetching such aggregate function result
into a variable using a different `INTERVAL` class will fail. To workaround this
problem, use a `CAST()` operator. This `CAST()` expression can be used
in static SQL, as long as the `INTERVAL` type specification matches an Informix
type:

```
SELECT CAST( SUM(duration) AS INTERVAL HOUR TO MINUTE )
   INTO total_duration
   FROM process
```

The
`INTERVAL` types translation can be controlled with the following FGLPROFILE
entry:

```
dbi.database.dsname.ifxemul.datatype.interval = { true | false }
```

For more details see [IBM Informix emulation parameters in FGLPROFILE](1079-ibm-informix-emulation-parameters-in-fglprofile.md "Emulation of Informix specific SQL features can be controlled with FGLPROFILE entries.").

## Related links

**Related concepts**  

[Using portable data types](1008-using-portable-data-types.md "Only a limited set of data types are really portable across several database engines.")

[Type conversion rules](1452-type-conversion-rules.md "Type conversion rules")
