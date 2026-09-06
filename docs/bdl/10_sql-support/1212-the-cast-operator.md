---
title: "The :: cast operator"
source: "fgl-topics/c_fgl_odiagifx_024.html"
breadcrumb: "SQL support > SQL database guides > IBM® Informix® > Unsupported IBM® Informix® SQL features > The :: cast operator"
type: "concept"
description: "IBM® Informix® SQL implements the :: cast operator and the CAST() expressions to do an explicit cast of a value: CREATE TABLE tab ( v INTEGER ) INSERT INTO tab VALUES ( 123456::INTEGER ) SELECT ..."
---

# The :: cast operator

IBM® Informix®
SQL implements the :: cast operator and the CAST() expressions to do an explicit cast of
a value:

```
CREATE TABLE tab ( v INTEGER )
INSERT INTO tab VALUES ( 123456::INTEGER )
SELECT 'abcdef'::CHAR(20)||'.' FROM tab 
SELECT CAST('abcdef' AS CHAR(20))||'.' FROM tab
```

Genero BDL does not support the :: cast operator in the static SQL grammar. However, the
CAST() expressions are allowed. If you need to use the :: cast operator, you must use [Dynamic SQL](1140-dynamic-sql-management.md "Explains how to execute and manage SQL statements at runtime.") to perform such
queries.

## Related links

**Related concepts**  

[Static SQL statements](1115-static-sql-statements.md "Describes static SQL statements supported in the language.")
