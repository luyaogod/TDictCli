---
title: "Assignment (:=)"
source: "fgl-topics/c_fgl_operators_ASSIGNMENT.html"
breadcrumb: "Language basics > Operators > List of expression elements > Assignment operators > Assignment (:=)"
type: "concept"
---

# Assignment (:=)

> The := operator assigns a variable with an expression and returns the result.

## Syntax

```
variable := expr
```

## Usage

The `:=` assignment operator puts a value in the left-hand variable and the
resulting value can again be used in an expression. In other programming languages, the
`:=` operator is known as the "walrus" operator.

The `:=` assignment operator is a different language element than the [`LET`](0701-let.md "The LET statement assigns values to variables.") instruction.

The `:=` assignment operator has the lowest precedence, it can be used at many
places and can simplify coding.

## Example

In the next example, the `:=` operator is used to increment the array index before
usage:

```
MAIN
  DEFINE arr DYNAMIC ARRAY OF STRING,
         idx INTEGER
  LET idx = 0
  LET arr[idx:=idx+1] = "One"
  LET arr[idx:=idx+1] = "Two"
  LET arr[idx:=idx+1] = "Three"
END MAIN
```

## Related links

**Related concepts**  

[Variables](0686-variables.md "Explains how to define program variables.")
