---
title: "INTERVAL data type"
source: "fgl-topics/c_fgl_odiagdmg_034.html"
breadcrumb: "SQL support > SQL database guides > Dameng® database server > Data dictionary > INTERVAL data type"
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

## Dameng®

Dameng provides the following interval data types:

| Dameng data type | Description |
| --- | --- |
| `INTERVAL year-month` | for storage of year/month interval values |
| `INTERVAL day-section` | for storage of day/second interval values |

## Solution

The `INTERVAL` data type and values are converted to the corresponding Dameng
interval types as follows:

| Informix data type | Dameng data type |
| --- | --- |
| `INTERVAL YEAR(n) TO YEAR` | `INTERVAL YEAR(n)` |
| `INTERVAL YEAR(n) TO MONTH` | `INTERVAL YEAR(n) TO MONTH` |
| `INTERVAL MONTH(n) TO MONTH` | `INTERVAL MONTH(n)` |
| `INTERVAL DAY(n) TO DAY` | `INTERVAL DAY(n)` |
| `INTERVAL DAY(n) TO HOUR` | `INTERVAL DAY(n) TO HOUR` |
| `INTERVAL DAY(n) TO MINUTE` | `INTERVAL DAY(n) TO MINUTE` |
| `INTERVAL DAY(n) TO SECOND` | `INTERVAL DAY(n) TO SECOND(0)` |
| `INTERVAL DAY(n) TO FRACTION(p)` | `INTERVAL DAY(n) TO SECOND(p)` |
| `INTERVAL HOUR(n) TO HOUR` | `INTERVAL HOUR(n)` |
| `INTERVAL HOUR(n) TO MINUTE` | `INTERVAL HOUR(n) TO MINUTE` |
| `INTERVAL HOUR(n) TO SECOND` | `INTERVAL HOUR(n) TO SECOND(0)` |
| `INTERVAL HOUR(n) TO FRACTION(p)` | `INTERVAL HOUR(n) TO SECOND(p)` |
| `INTERVAL MINUTE(n) TO MINUTE` | `INTERVAL MINUTE(n)` |
| `INTERVAL MINUTE(n) TO SECOND` | `INTERVAL MINUTE(n) TO SECOND(0)` |
| `INTERVAL MINUTE(n) TO FRACTION(p)` | `INTERVAL MINUTE(n) TO SECOND(p)` |
| `INTERVAL SECOND(n) TO SECOND` | `INTERVAL SECOND(n)` |
| `INTERVAL SECOND(n) TO FRACTION(p)` | `INTERVAL SECOND(n,p)` |
| `INTERVAL FRACTION TO FRACTION(p)` | `INTERVAL SECOND(1,p)` |

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
