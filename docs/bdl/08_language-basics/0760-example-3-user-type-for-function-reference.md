---
title: "Example 3: User type for function reference"
source: "fgl-topics/c_fgl_user_types_example_3.html"
breadcrumb: "Language basics > Types > Examples > Example 3: User type for function reference"
type: "concept"
description: "This example shows how to define a user type from a function type. TYPE t_func_ref FUNCTION (p1 INT, p2 INT) RETURNS INT MAIN DEFINE fr t_func_ref LET fr = FUNCTION add DISPLAY fr(100,200) END MAIN ..."
---

# Example 3: User type for function reference

This example shows how to define a user type from a function type.

```
TYPE t_func_ref FUNCTION (p1 INT, p2 INT) RETURNS INT

MAIN
    DEFINE fr t_func_ref
    LET fr = FUNCTION add
    DISPLAY fr(100,200)
END MAIN

FUNCTION add(p1 INT, p2 INT) RETURNS INT
    RETURN (p1 + p2)
END FUNCTION
```
