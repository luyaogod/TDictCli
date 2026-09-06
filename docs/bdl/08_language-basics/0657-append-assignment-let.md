---
title: "Append Assignment (,=) [LET]"
source: "fgl-topics/c_fgl_operators_compassign_comma.html"
breadcrumb: "Language basics > Operators > List of expression elements > Assignment operators > Append Assignment (,=) [LET]"
type: "concept"
---

# Append Assignment (,=) [LET]

> The ,= operator used after the LET keyword, assigns a variable by concatenating the current variable value with the formatted version of an expression.

## Syntax

```
variable ,= expr
```

> **Important:**
>
> This operator cannot be used as an element of an
> expression, it is only allowed as assignment operator with the [`LET`](0701-let.md "The LET statement assigns values to variables.") instruction.

## Usage

The `,=` assignment operator concatenates the value of the left-hand variable to a
formatted form of the right-hand expression, and assigns the result to the variable.

The right-hand expression gets formatted with left space padding, according to the data type of
the expression. For example, an `INTEGER` expression gets spaces on the left of the
value, to always occupy 11 character positions. The [`USING`](0634-using.md "The USING operator converts date and numeric values to a string based on a formatting mask.") operator can be specified to control value formatting.

The target variable should be a `VARCHAR` or `STRING` variable.

Note that `CHAR(size)` variables are blank-padded to
size characters. As result, it's useless to apply the `,=`
assignment operator, as the value of the left-hand operand already fills the `CHAR`
size completely.

`TEXT` variables cannot be used with the `,=` formatting assignment
operator.

## Example

```
MAIN
  DEFINE x STRING
  LET x = "abc"
  LET x ,= 123
  DISPLAY x
END MAIN
```

## Related links

**Related concepts**  

[Formatting data](0580-formatting-data.md "Explains data to string conversion options of the language.")
