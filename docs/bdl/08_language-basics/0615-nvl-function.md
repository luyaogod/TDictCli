---
title: "NVL() [function]"
source: "fgl-topics/c_fgl_operators_NULL_VALUE_SUBSTITUTION.html"
breadcrumb: "Language basics > Operators > List of expression elements > Comparison operators > NVL() [function]"
type: "concept"
---

# NVL() [function]

> The NVL() operator returns the second parameter if the first argument evaluates to NULL.

## Syntax

```
NVL( main-expr, subst-expr )
```

1. main-expr and subst-expr are
   any expression supported by the language.

## Usage

The `NVL()` operator evaluates the first argument, and returns the result if the
value is not null, otherwise it returns the second argument.

This allows you to write the equivalent of the following [`IF` statement](0682-if.md "The IF instruction executes a group of statements conditionally."), in a simple scalar expression:

```
IF main-expr IS NOT NULL THEN
  RETURN main-expr
ELSE
  RETURN subst-expr
END IF
```

## Example

```
MAIN
  DEFINE x VARCHAR(100)
  LET x = arg_val(1)
  DISPLAY "The argument value is: ", NVL(x, "NULL")
END MAIN
```

## Related links

**Related concepts**  

[Expressions](0592-expressions.md "Shows the possible expressions supported in the language.")
