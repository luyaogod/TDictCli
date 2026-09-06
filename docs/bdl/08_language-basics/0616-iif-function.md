---
title: "IIF() [function]"
source: "fgl-topics/c_fgl_operators_IMMEDIATE_IF.html"
breadcrumb: "Language basics > Operators > List of expression elements > Comparison operators > IIF() [function]"
type: "concept"
---

# IIF() [function]

> The IIF() returns the second or third parameter depending on the boolean expression given as first argument.

## Syntax

```
IIF( bool-expr, true-expr, false-expr )
```

1. bool-expr is a boolean expression.
2. true-expr and false-expr are
   language expressions.

## Usage

The `IIF()` operator evaluates the first argument, the
returns the second argument if the first argument is true, otherwise it
returns the third argument.

This allows you to write the equivalent of the following [`IF` statement](0682-if.md "The IF instruction executes a group of statements conditionally."),
in a simple scalar expression:

```
IF bool-expr THEN
   RETURN true-expr
ELSE
   RETURN false-expr
END IF
```

## Example

```
MAIN
  DEFINE x VARCHAR(10)
  LET x = arg_val(1)
  DISPLAY IIF(x == "A", "Accepted", "Rejected")
END MAIN
```

## Related links

**Related concepts**  

[Boolean expressions](0594-boolean-expressions.md "This section covers boolean expression evaluation rules.")
