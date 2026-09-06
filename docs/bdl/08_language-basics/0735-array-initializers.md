---
title: "Array initializers"
source: "fgl-topics/c_fgl_Arrays_initializer.html"
breadcrumb: "Language basics > Arrays > Array initializers"
type: "concept"
---

# Array initializers

> Arrays can be initialized in their definition.

To initialize an array variable in its definition, use the equal sign followed by an array
initializer.

An array initializer is specified with brackets, with array elements separated by a
comma:

```
DEFINE arr DYNAMIC ARRAY OF RECORD
           cust_id INT,
           cust_name VARCHAR(50)
       END RECORD = [
          ( cust_id: 101, cust_name: "Mike TORN" ),
          ( cust_id: 234, cust_name: "Phil TOLLINS" )
       ]
```

For more details, see [Variable initializers](0691-variable-initializers.md "Variables can be initialized in their definition.").
