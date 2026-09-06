---
title: "Example 4: TRY / CATCH in conjunction with WHENEVER"
source: "fgl-topics/c_fgl_Exceptions_018.html"
breadcrumb: "Advanced features > Exceptions > Examples > Example 4: TRY / CATCH in conjunction with WHENEVER"
type: "concept"
description: "This code illustrates the fact that a TRY/CATCH block can be used together with a WHENEVER instruction: The program first executes a WHENEVER ANY ERROR to define an error handler named foo and later ..."
---

# Example 4: TRY / CATCH in conjunction with WHENEVER

This code illustrates the fact that a `TRY/CATCH` block can be used
together with a `WHENEVER` instruction: The program first executes a
`WHENEVER ANY ERROR` to define an error handler named *foo* and
later it uses a `TRY/CATCH` block to trap expression errors.

In this example, we intentionally force a division by zero. After the `TRY/CATCH`
block, we force another division by zero error, which will call the *foo* error
handler:

```
MAIN
  DEFINE i INTEGER
  WHENEVER ANY ERROR CALL foo
  TRY
    DISPLAY "Next exception should be handled by the catch statement"
    LET i = i / 0
  CATCH
    DISPLAY "Exception caught, status: ", status
  END TRY
  -- Previous error handler is restored after the TRY - CATCH block 
  LET status = 0
  DISPLAY "Next exception should be handled by the foo function"
  LET i = i / 0
END MAIN

FUNCTION foo() RETURNS ()
  DISPLAY "Function foo called, status: ", status
END FUNCTION
```

Program output:

```
Next exception should be handled by the catch statement
Exception caught, status:      -1202
Next exception should be handled by the foo function 
Function foo called, status:      -1202
```
