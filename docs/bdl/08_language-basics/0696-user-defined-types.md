---
title: "User defined types"
source: "fgl-topics/c_fgl_variables_007.html"
breadcrumb: "Language basics > Variables > User defined types"
type: "concept"
---

# User defined types

> User defined types help to centralize the definition of complex structured data types.

A user defined type is created with the `TYPE` keyword. The next code example
defines a type as a dynamic array of a record, with the structure of a database table as defined in
the mydbschema.sch schema
file:

```
SCHEMA mydbschema
...
TYPE t_custlist DYNAMIC ARRAY OF RECORD LIKE customer.*
```

Variables can then be defined with the `"custlist"` user defined
type:

```
DEFINE cl t_custlist
```

The scope of a type can be global, local to a module or local to a function.

Variables can be defined with a type defined in the same scope, or in a higher level of
scope.

A typical usage is to declare user-defined types as `PUBLIC` in a module, so they
can be reused in other modules:

The myutils.4gl module defines the user type:

```
PUBLIC TYPE t_item_info RECORD
        id INTEGER,
        description VARCHAR(100),
        creation DATE,
        mandatory BOOLEAN,
        ...
    END RECORD
...
```

In the module importing the utility module, define variables using the type name, with utility
module as
prefix:

```
IMPORT FGL myutils
...
  DEFINE itemlist DYNAMIC ARRAY OF myutils.t_item_info
...
```

## Related links

**Related concepts**  

[Types](0751-types.md "Types can be defined by the programmer to centralize the definition of complex/structured variables.")
