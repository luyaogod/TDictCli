---
title: "base.Application.getStackTrace"
source: "fgl-topics/c_fgl_ClassApplication_getStackTrace.html"
breadcrumb: "Library reference > Built-in packages > The base package > The Application class > base.Application methods > base.Application.getStackTrace"
type: "concept"
---

# base.Application.getStackTrace

> Returns the function call stack trace.

## Syntax

```
base.Application.getStackTrace()
  RETURNS STRING
```

## Usage

Use the `getStackTrace()` method, to print the stack trace to a log file.

> **Important:**
>
> Sensitive and personal data may be written to the output. Make sure that the log output is
> written to files that can only be read by application administrators.

This method returns a string containing a formatted list of the current function stack.

The output format of `getStackTrace()` method is for debug purpose only and can
change in future product releases.

You typically use this function in a [`WHENEVER
ERROR CALL`](../09_advanced-features/0850-whenever-directive.md "Use the WHENEVER directive to define how exceptions must be handled for the rest of the module.") handler.

```
MAIN
  WHENEVER ANY ERROR CALL my_handler
  CALL func1("abcdef",999)
END MAIN

FUNCTION func1(name,id)
  DEFINE name STRING, id INTEGER
  DEFINE r INTEGER
  LET r = div(5,0)
END FUNCTION

FUNCTION div(x,y)
  DEFINE x,y DECIMAL
  RETURN ( x / y )
END FUNCTION

FUNCTION my_handler()
  DISPLAY base.Application.getStackTrace()
END FUNCTION
```

Example of stack trace output:

```
#0 my_handler() at x.4gl:17
#1 div() at x.4gl:14
#2 func1() at x.4gl:9
#3 main() at x.4gl:3
```
