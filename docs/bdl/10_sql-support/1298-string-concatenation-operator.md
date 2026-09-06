---
title: "String concatenation operator"
source: "fgl-topics/c_fgl_odiagmsv_043.html"
breadcrumb: "SQL support > SQL database guides > Microsoft™ SQL Server > Data manipulation > String concatenation operator"
type: "concept"
description: "Informix® The Informix concatenation operator is the double pipe ( || ): SELECT firstname || ' ' || lastname FROM employee Microsoft™ SQL Server The legacy Microsoft SQL Server concatenation operator ..."
---

# String concatenation operator

## Informix®

The Informix concatenation operator is the double pipe
( || ):

```
SELECT firstname || ' ' || lastname FROM employee
```

## Microsoft™ SQL Server

The legacy Microsoft SQL Server concatenation operator
is the plus sign:

```
SELECT firstname + ' ' + lastname FROM employee
```

Starting with SQL Server 2025, the ANSI standard double-pipe concatenation operator is supported
as with many other SQL database
servers:

```
SELECT firstname || ' ' || lastname FROM employee
```

## Solution

When connected to an SQL Server version prior to 2025, the database interface detects double-pipe
operators in SQL statements and converts them to a plus sign automatically. There is no FGLPROFILE
option to disable double-pipe to plus sign conversion.

When connected to an SQL Server version 2025 or higher, the database interface leaves the
double-pipe operators as is in SQL statements.

## Related links

**Related concepts**  

[String concatenation operators in SQL](1032-string-concatenation-operators-in-sql.md "The || operator is the standard to concatenate strings.")
