---
title: "BEFORE/AFTER FIELD field prefix"
source: "fgl-topics/c_fgl_MigI4GL_050.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > User interface topics > BEFORE/AFTER FIELD field prefix"
type: "concept"
---

# BEFORE/AFTER FIELD field prefix

> Recent I4GL versions deny the specification of a screen-record prefix in BEFORE FIELD and AFTER FIELD clauses, Genero BDL allows this for backward compatibility.

With recent IBM® Informix® 4GL versions, an interactive instruction defining a `BEFORE FIELD` or
`AFTER FIELD` clause with a screen record or table name field prefix will produce a
compiler error.

For example, with c4gl 7.51.FC2:

```
MAIN

    DEFINE rec RECORD
               cust_id INTEGER,
               cust_name VARCHAR(50)
           END RECORD

    INPUT BY NAME rec.*
       BEFORE FIELD customer.cust_name
|
|      When doing INPUT BY NAME or INPUT ARRAY, the BEFORE/AFTER field names
| can be specified only by the field name suffix.  Screen array and screen
| record elements are not allowed.
| See error number -4391.
         MESSAGE "Customer name"
    END INPUT

END MAIN
```

With Genero BDL, the field prefix is allowed and ignored.

The field prefix in an `INPUT BY NAME` statement is useless.

Fixing this difference in Genero BDL would force to change the code, with no added value.

## Related links

**Related concepts**  

[BEFORE FIELD block](../11_user-interface/1942-before-field-block.md "BEFORE FIELD block")

[AFTER FIELD block](../11_user-interface/1944-after-field-block.md "AFTER FIELD block")
