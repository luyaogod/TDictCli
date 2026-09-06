---
title: "Example 4: Defining methods"
source: "fgl-topics/c_fgl_Functions_example_4.html"
breadcrumb: "Language basics > Functions > Examples > Example 4: Defining methods"
type: "concept"
description: "This example shows how to define methods for a Rectangle type: PUBLIC TYPE Rectangle RECORD height, width FLOAT END RECORD PUBLIC FUNCTION (r Rectangle) area () RETURNS FLOAT RETURN r.height * r.width ..."
---

# Example 4: Defining methods

This example shows how to define methods for a `Rectangle`
type:

```
PUBLIC TYPE Rectangle RECORD
    height, width FLOAT
END RECORD

PUBLIC FUNCTION (r Rectangle) area () RETURNS FLOAT
    RETURN r.height * r.width
END FUNCTION

PUBLIC FUNCTION (r Rectangle) kind () RETURNS STRING
    RETURN "Rectangle"
END FUNCTION

PUBLIC FUNCTION (r Rectangle) setDimensions (w FLOAT, h FLOAT) RETURNS ()
    LET r.width = w
    LET r.height = h
END FUNCTION
```
