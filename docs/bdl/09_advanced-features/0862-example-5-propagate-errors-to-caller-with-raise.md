---
title: "Example 5: Propagate errors to caller with RAISE"
source: "fgl-topics/c_fgl_Exceptions_019.html"
breadcrumb: "Advanced features > Exceptions > Examples > Example 5: Propagate errors to caller with RAISE"
type: "concept"
description: "This example shows the usage of WHENEVER ... RAISE to propagate a potential exception to the caller. The division by zero error will occur in the divide() function, but appear as if the division ..."
---

# Example 5: Propagate errors to caller with RAISE

This example shows the usage of `WHENEVER ... RAISE` to propagate a potential
exception to the caller. The division by zero error will occur in the `divide()`
function, but appear as if the division operator was used directly at the line calling this
function:

```
MAIN
    TRY
        DISPLAY "Next function call will generate an exception"
        DISPLAY divide(100, 0)
    CATCH
        DISPLAY "Error: ", status
    END TRY
END MAIN

FUNCTION divide(a FLOAT, b FLOAT) RETURNS FLOAT
    WHENEVER ANY ERROR RAISE
    RETURN a / b
END FUNCTION
```

Program output:

```
Next function call will generate an exception 
Error:    -1202
```
