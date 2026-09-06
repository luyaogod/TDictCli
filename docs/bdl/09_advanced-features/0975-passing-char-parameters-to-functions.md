---
title: "Passing CHAR parameters to functions"
source: "fgl-topics/c_fgl_optimization_011.html"
breadcrumb: "Advanced features > Optimization > Optimize your programs > Passing CHAR parameters to functions"
type: "concept"
description: "Simple data types such as CHAR are passed by value to functions (this means the value of the caller variable is copied on the stack, and then copied back into a local variable of the called function.) ..."
---

# Passing CHAR parameters to functions

Simple data types such as `CHAR` are passed by value to functions (this means the
value of the caller variable is copied on the stack, and then copied back into a local variable of
the called function.) When large data types are used, this can introduce a performance issue.

For example, the following code defines a logging function that takes a [`CHAR(2000)`](../08_language-basics/0557-char-size.md "The CHAR data type is a fixed-length character string data type.") as parameter:

```
FUNCTION log_msg( msg )
  DEFINE msg CHAR(2000)
  CALL myLogChannel.writeLine(msg)
END FUNCTION
```

The function is then typically used by passing a log message as parameter:

```
CALL log_msg( "Start processing..." )
```

When performing this call, the runtime system copies 19 characters on the stack, calls the
function, and then copies the value into the local variable. Since the values in
`CHAR` variables must always have a length matching the variable definition size, the
runtime system fills the remaining 1981 positions with blanks. As result, each time you call this
function, a 2000 character-long variable is created on the stack.

By using a [`VARCHAR(2000)`](../08_language-basics/0570-varchar-size.md "The VARCHAR data type is a variable-length character string data type, with a maximum size.") (or a
[`STRING`](../08_language-basics/0567-string.md "The STRING data type is a variable-length, dynamically allocated character string data type, without limitation.")) data type in this function,
you optimize the execution, because no trailing blanks need to be added for these types.

Consider also using `VARCHAR` instead of `CHAR` for large database
columns. Due to the Informix® SQL history
and the Informix
`VARCHAR(255)` limit, there may be `CHAR()` columns when the storage
size required is more than 255 bytes. For more details, see [Trailing blanks in CHAR/VARCHAR](../10_sql-support/1017-trailing-blanks-in-char-varchar.md "How to cope with trailing blanks in CHAR(N) and VARCHAR(N) SQL columns and program variables?").

## Related links

**Related concepts**  

[Manipulating character strings](0976-manipulating-character-strings.md "Manipulating character strings")
