---
title: "Fetching numbers into CHAR/VARCHAR"
source: "fgl-topics/c_fgl_MigI4GL_062.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > 4GL programming topics > Fetching numbers into CHAR/VARCHAR"
type: "concept"
---

# Fetching numbers into CHAR/VARCHAR

> I4GL and FGL format numbers in a different way when fetching direcly MONEY, DECIMAL column values into CHAR/VARCHAR variables.

With IBM® Informix® 4GL, when fetching SQL column values from a `MONEY` or
`DECIMAL` column into a `CHAR(n)` or `VARCHAR(n)`
program variable, string-to-number conversion rules apply, based on the DBMONEY or DBFORMAT
environment variable.

With Genero BDL, only DBMONEY is taken into account: DBFORMAT is ignored.

In such case, best practice is to fetch MONEY or DECIMAL data into a program variables defined
with the same data type. This can be easily achieved with a `DEFINE LIKE` statement,
using on the SQL column types of a schema file:

```
SCHEMA stock
...
DEFINE amount LIKE order.ord_amount
```

## Related links

**Related concepts**  

[Type conversions](../08_language-basics/0576-type-conversions.md "Explains primitive data type conversion rules of the language.")

[Database column types](../08_language-basics/0695-database-column-types.md "Simple variables and record structures can be defined from database columns types.")
