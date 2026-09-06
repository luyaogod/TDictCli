---
title: "Example 3: Using function references"
source: "fgl-topics/c_fgl_Functions_example_3.html"
breadcrumb: "Language basics > Functions > Examples > Example 3: Using function references"
type: "concept"
description: "This example shows how to invoke functions dynamically by reference. The module \" compute.4gl \": PUBLIC TYPE compute_function FUNCTION(p1 INT, p2 INT) RETURNS INT PRIVATE DEFINE cf compute_function ..."
---

# Example 3: Using function references

This example shows how to invoke functions dynamically by reference.

The module
"compute.4gl":

```
PUBLIC TYPE compute_function FUNCTION(p1 INT, p2 INT) RETURNS INT

PRIVATE DEFINE cf compute_function
    
PUBLIC FUNCTION set_function( f compute_function )
    LET cf = f
END FUNCTION
    
PUBLIC FUNCTION compute(i1 INT, i2 INT) RETURNS INT
    DEFINE r INT
    IF cf IS NULL THEN
       DISPLAY "ERROR: Define the function with set_function(FUNCTION <name>)"
       EXIT PROGRAM 1
    END IF
    LET r = cf(i1, i2)
    DISPLAY SFMT("compute(%1, %2) = %3", i1, i2, r)
    RETURN r
END FUNCTION

PUBLIC FUNCTION add(p1 INT, p2 INT) RETURNS INT
    RETURN p1 + p2
END FUNCTION

PUBLIC FUNCTION sub(p1 INT, p2 INT) RETURNS INT
    RETURN p1 - p2
END FUNCTION
```

The main program:

```
IMPORT FGL compute

MAIN
    DEFINE r INT
    CALL compute.set_function(FUNCTION compute.add)
    LET r = compute(1, 2)
    CALL compute.set_function(FUNCTION compute.sub)
    LET r = compute(1, 2)
END MAIN
```
