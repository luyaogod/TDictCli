---
title: "Example 1: Defining and using constants"
source: "fgl-topics/c_fgl_Constants_006.html"
breadcrumb: "Language basics > Constants > Examples > Example 1: Defining and using constants"
type: "concept"
---

# Example 1: Defining and using constants

> This example shows how to define constants.

```
CONSTANT
    c1 ="Drink",       # Declares a STRING constant
    c2 = 4711,         # Declares an INTEGER constant
    n = 10,            # Declares an INTEGER constant
    x SMALLINT=1       # Declares a SMALLINT constant

DEFINE a ARRAY[n] OF INTEGER

MAIN
  CONSTANT c1 = "Hello"
  DEFINE i INTEGER
  FOR i=1 TO n 
      ...
  END FOR
  DISPLAY c1 || c2  # Displays "Hello4711"
END MAIN
```
