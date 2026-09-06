---
title: "Database schema in globals"
source: "fgl-topics/c_fgl_Globals_005.html"
breadcrumb: "Advanced features > Globals > Database schema in globals"
type: "concept"
---

# Database schema in globals

> Globals files can define the database schema to be used by the compiler to resolve DEFINE ... LIKE statements.

The schema specification must appear before the `GLOBALS` keyword
starting the globals block.

The schema specification is propagated to the modules including the globals file defining
the database schema. These modules can use `DEFINE ... LIKE` without an explicit
`SCHEMA` instruction.

Furthermore, when using the `DATABASE` instruction instead of
`SCHEMA`, if the module including the globals contains the `MAIN`
block, the `DATABASE` specification of the globals file will be propagated and
result in an implicit database connection at runtime.

## Example

```
SCHEMA stores
GLOBALS
  DEFINE cust_rec LIKE customer.*
  ...
END GLOBALS
```

## Related links

**Related concepts**  

[Database schema](0791-database-schema.md "Defines database table structures with column type information to be reused in program variable definitions.")
