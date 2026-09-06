---
title: "errorlog()"
source: "fgl-topics/c_fgl_BuiltInFunctions_ERRORLOG.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > errorlog()"
type: "concept"
---

# errorlog()

> Copies the string passed as parameter into the error log file.

## Syntax

```
FUNCTION errorlog(
   message STRING )
```

1. message is the character string to be inserted in the error log file.

## Usage

The `errorlog()` function writes the passed string in the current error
log file. The error log file is defined by a call to the `startlog()` function.

> **Important:**
>
> Sensitive data may be written to the `startlog()` file. Make
> sure that the log output produced by `errorlog()` calls is written to files that can
> only be read by application administrators.

Use this function to identify errors in programs and to customize error handling. The error log
functions can also be used to trace the way a program is used in order to improve it, record work
habits, or help to detect attempts to breach security.

## Related links

**Related concepts**  

[startlog()](2790-startlog.md "Initializes error logging and opens the error log file passed as the parameter.")
