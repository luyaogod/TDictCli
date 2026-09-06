---
title: "Default exception handling"
source: "fgl-topics/c_fgl_Exceptions_012.html"
breadcrumb: "Advanced features > Exceptions > Default exception handling"
type: "concept"
---

# Default exception handling

> Default exception handling must be adapted to your programming pattern.

## Default program behavior when exception occurs

By default, when a [language or SQL error](0851-exception-classes.md "Exception classes define the kind of issues that can occur at runtime.") occurs,
the program stops and shows a message to the end user. However, by default, expression errors such
as type conversion errors, type overflows and division by zero will not stop the program. In other
words, the default is `WHENEVER ERROR STOP + WHENEVER ANY ERROR CONTINUE`.

In TUI mode, with default exception handlers, the error message is displayed in the terminal.

In GUI mode, with default exception handlers, the error message is displayed in a popup message
box, that the user can read before the program stops. If showing detailed error messages to the end
user is considered as a security risk (to prevent attacks), the original error message can be
replaced by a generic error message defined in the `gui.programStoppedMessage` [FGLPROFILE](../07_configuration/0486-fglprofile-entries-for-core-language.md "This is a summary of FGLPROFILE entries supported by the core BDL.")
entry:

```
gui.programStoppedMessage = "An unexpected error occurred, program will stop."
```

## Enforce expression error handling

By default, the `WHENEVER ANY ERROR` action is to `CONTINUE` the
program flow. This means that when a program makes for example a division by zero or when a string
cannot be converted to a number, the program continues. This can lead to unexpected program
behavior, but is the default to be backward compatible for legacy applications.

To make your code more robust, use `WHENEVER ANY ERROR STOP`.

Alternatively, in case of expression errors, you can force the runtime system to execute the
action defined with `WHENEVER ERROR` exception class, with the following [FGLPROFILE](../07_configuration/0483-the-fglprofile-file-s.md "FGLPROFILE environment variable defines Genero BDL configuration files")
entry:

```
fglrun.mapAnyErrorToError = true
```

When this entry is set to true, [expression errors](0851-exception-classes.md "Exception classes define the kind of issues that can occur at runtime.")
such as a division by zero will be trapped and execute the action defined by the last
`WHENEVER ERROR` instruction.

When using the default exception handler (`WHENEVER ERROR STOP`), the program with
stop at any expression error, and display the corresponding error message.

FGLPROFILE file:

```
fglrun.mapAnyErrorToError = true
```

Program
code:

```
MAIN
  DEFINE x INT
  WHENEVER ERROR CALL my_error_handler 
  LET x = 1 / 0   -- error handler will be called here 
  DISPLAY "It continues...."
END MAIN

FUNCTION my_error_handler()
  DISPLAY "Handler: ", status
END FUNCTION
```

## Related links

**Related reference**  

[Genero BDL errors](../15_library-reference/4483-genero-bdl-errors.md "System error messages sorted by error number.")
