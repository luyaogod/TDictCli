---
title: "Anonymous types"
source: "fgl-topics/c_fgl_user_types_005.html"
breadcrumb: "Language basics > Types > Anonymous types"
type: "concept"
---

# Anonymous types

> Anonymous types are created from DEFINE instructions.

Unlike [user-defined types](0752-understanding-type-definition.md "This is an introduction to types.") which are created with
the `TYPE` instruction and are identified by a name, anonymous types are
created automatically when defining variables with an explicit type specification.

For example:

```
DEFINE name VARCHAR(50)
DEFINE cust_rec RECORD
           cust_id INTEGER,
           cust_name VARCHAR(50)
       END RECORD
DEFINE cust_list DYNAMIC ARRAY OF RECORD LIKE customer.*
```

The above code will define 3 anonymous types for the module:

1. A type for `VARCHAR(50)`.
2. A structured type for the `RECORD` definition, including `cust_id`
   and `cust_name` fields.
3. An array type using a structured type for elements.

When defining several variables with the same type specification, the compiler will only create
one anonymous type, that will be reused for each variable definition.

Anonymous types are for example used when doing variable introspection with the reflection API,
when returning a type from the [`reflect.Value.getType()`](../15_library-reference/4371-reflect-value-gettype.md "Returns the reflect.Type object of a reflect.Value object.") method of a value object created from a
variable.
