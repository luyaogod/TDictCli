---
title: "INTERVAL data type"
source: "fgl-topics/c_fgl_odiagmsv_037.html"
breadcrumb: "SQL support > SQL database guides > Microsoft™ SQL Server > Data dictionary > INTERVAL data type"
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

## Microsoft™ SQL Server

Microsoft SQL Server does not provide a
data type corresponding the Informix
`INTERVAL` data type.

## Solution

The `INTERVAL` data type and values are converted `CHAR(50)` column
with Microsoft SQL Server.

`INTERVAL` values can be stored and retrieved from the database. However, since
Microsoft SQL Server does not support a native interval
type, arithmetics cannot be performed on the database side in SQL statements.

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
