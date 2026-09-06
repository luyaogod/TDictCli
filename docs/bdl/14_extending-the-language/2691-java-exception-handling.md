---
title: "Java exception handling"
source: "fgl-topics/c_fgl_JavaBridge_037.html"
breadcrumb: "Extending the language > The Java interface > Advanced programming > Java exception handling"
type: "concept"
description: "In order to catch Java exceptions within programs, use a TRY/CATCH block . When a Java exception occurs, the runtime system sets the status variable to the error code -8306 . The Java exception ..."
---

# Java exception handling

In order to catch Java exceptions within programs, use a
[`TRY/CATCH` block](../09_advanced-features/0853-try-catch-block.md "Use TRY / CATCH blocks to trap runtime exceptions in a delimited code block.").

When a Java exception occurs, the runtime system sets the `status` variable
to the error code [-8306](../15_library-reference/4483-genero-bdl-errors.md).

The Java exception details (that is the name of the exception) can be found with the [ERR\_GET(status)](../15_library-reference/2732-err-get.md "Returns the text corresponding to an error number.") built-in function.

To get the Java exception type with `ERR_GET()`, do not execute other instructions
before querying for the error message, otherwise the `status` variable may be reset
to zero and the Java exception details would be lost.

To easily identify the type of the Java exceptions in your code, consider writing a library
function based on `ERR_GET()`, that recognizes most common Java exceptions,
and converts them to integer codes:

```
IMPORT JAVA java.lang.StringBuffer
MAIN
  DEFINE sb java.lang.StringBuffer
  LET sb = StringBuffer.create("abcdef")
  TRY
    CALL sb.deleteCharAt(50) -- out of bounds!
  CATCH
    DISPLAY err_get(status)
    EXIT PROGRAM 1
  END TRY
END MAIN
```

As a general pattern, do not use `TRY/CATCH` or `WHENEVER ERROR
CONTINUE` exception handlers if no exception is supposed to occur. By default the program
will then stop and display the Java exception details.
