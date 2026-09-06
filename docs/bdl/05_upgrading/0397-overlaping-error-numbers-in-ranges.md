---
title: "Overlaping error numbers in ranges"
source: "fgl-topics/c_fgl_MigI4GL_061.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > 4GL programming topics > Overlaping error numbers in ranges"
type: "concept"
---

# Overlaping error numbers in ranges

> Genero BDL uses error numbers in a range of Informix SQL errors.

Genero FGL uses the following error number ranges:

- `-6999:-6000 / 6000:6999` (FGL, WRT)
- `-8999:-8000 / 8000:8999` (FGL)
- `-15999:-15000 / 15000:15999` (GWS)

Recent Informix servers defines SQL errors in the range -8999 to -8000, such
as:

```
 -8128   Minimum length of a VARCHAR variable must be smaller than the maximum size.
```

In this range, the errors must be interpreted accoridng to the context: After executing an SQL
statement when connected to an Informix server, it's an Informix SQL error. In any other context,
it's a Genero BDL error.

Note also that Genero BDL runtime errors will only set the `status` variable,
while SQL errors will set the `sqlca.sqlcode` and `status`
variable.

## Related links

**Related concepts**  

[SQL execution diagnostics](../10_sql-support/0987-sql-execution-diagnostics.md "If an SQL statement execution fails, error description can be found in the sqlca.sqlcode, SQLSTATE, status and SQLERRMESSAGE predefined registers.")
