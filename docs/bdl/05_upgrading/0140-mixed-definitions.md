---
title: "Mixed definitions"
source: "fgl-topics/c_fgl_Migrate_to_400_mixed_definitions.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 4.00 upgrade guide > Mixed definitions"
type: "concept"
---

# Mixed definitions

> Definition of functions, variables, constants and types can be mixed.

Before Genero BDL version 4.00, the language syntax forced you to define variables, constants and
types before functions, methods and reports:

1. Types ([`TYPE`](../08_language-basics/0751-types.md "Types can be defined by the programmer to centralize the definition of complex/structured variables.")), Constants ([`CONSTANT`](../08_language-basics/0710-constants.md "The definition of constants allows to centralize common static values.")), Variables ([`DEFINE`](../08_language-basics/0686-variables.md "Explains how to define program variables."))
2. Functions/Methods ([`FUNCTION`](../08_language-basics/0761-functions.md "Describes user defined functions.")),
   Reports ([`REPORT`](../12_reports/2461-reports.md)), Declarative dialogs
   ([`DIALOG`](../11_user-interface/2150-declarative-dialogs-dialog-at-module-level.md "DIALOG/END DIALOG defined at module level implement declarative dialogs that can be used in procedural dialogs."))

Starting with version 4.00, the definition of functions, methods, reports, variables, constants
and types can now be mixed in any order, to group definitions:

```
PUBLIC CONSTANT DEFAULT_SCOPE = "basic"
PRIVATE DEFINE current_scope STRING
FUNCTION set_current_scope(s STRING) RETURNS ()
    LET current_scope = s
END FUNCTION

PRIVATE DEFINE current_rate DECIMAL(10,2)
FUNCTION set_current_rate(r DECIMAL(10,2)) RETURNS ()
    LET current_rate = r
END FUNCTION
```

Any symbol can be forward referenced (a symbol can be used before its lexical
declaration):

```
FUNCTION check_customer_order()
    -- In next line, cust_rec usage is a forward reference
    IF order_rec.cust_id != cust_rec.cust_num THEN
       ...
    END IF
END FUNCTION
...
-- Customer 
DEFINE cust_rec RECORD LIKE customer.* -- Definition
```

Below another example, defining two types with methods in the same module:

```
TYPE t_cust RECORD
        pkey INT,
        name VARCHAR(50)
     END RECORD

FUNCTION (c t_cust) set_defaults() RETURNS ()
    LET c.pkey = generate_cust_id()
    LET c.name = "<undefined>"
END FUNCTION

TYPE t_ord RECORD
        pkey INT,
        cust_id INT,
        crea_date DATE
     END RECORD

FUNCTION (o t_ord) set_defaults() RETURNS ()
    LET o.pkey = generate_ord_id()
    LET o.cust_id = -1
    LET o.crea_date = TODAY
END FUNCTION
```

For more details, see [Content of a .4gl module](../09_advanced-features/0788-content-of-a-4gl-module.md "A module defines a set of program elements.").
