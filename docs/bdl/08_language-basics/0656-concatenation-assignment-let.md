---
title: "Concatenation Assignment (||=) [LET]"
source: "fgl-topics/c_fgl_operators_compassign_concat.html"
breadcrumb: "Language basics > Operators > List of expression elements > Assignment operators > Concatenation Assignment (||=) [LET]"
type: "concept"
---

# Concatenation Assignment (||=) [LET]

> The ||= operator used after the LET keyword, assigns a variable by concatenating the current variable value to an expression.

## Syntax

```
variable ||= expr
```

> **Important:**
>
> This operator cannot be used as an element of an
> expression, it is only allowed as assignment operator with the [`LET`](0701-let.md "The LET statement assigns values to variables.") instruction.

## Usage

The `||=` assignment operator concatenates the value of the left-hand variable to
the right-hand expression, and assigns the result to the variable.

Unlike the [`,=`
operator](0657-append-assignment-let.md "The ,= operator used after the LET keyword, assigns a variable by concatenating the current variable value with the formatted version of an expression."), with `||=`, the right-hand expression is concatenated wihout any
additional left space padding.

If one of operands of `||=` is `NULL`, the resulting string will be
`NULL`.

The target variable should be a `VARCHAR`, `STRING` or
`TEXT` variable.

Note that `CHAR(size)` variables are blank-padded to
size characters. As result, it's useless to apply the `||=`
assignment operator, as the value of the left-hand operand already fills the `CHAR`
size completely.

## Example

```
MAIN
  DEFINE x STRING
  LET x = "abc"
  LET x ||= "def"
  DISPLAY x
END MAIN
```

## Related links

**Related concepts**  

[Concatenate (||)](0632-concatenate.md "The || operator makes a string concatenation.")
