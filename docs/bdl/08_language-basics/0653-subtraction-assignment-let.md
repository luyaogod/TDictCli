---
title: "Subtraction Assignment (-=) [LET]"
source: "fgl-topics/c_fgl_operators_compassign_subtract.html"
breadcrumb: "Language basics > Operators > List of expression elements > Assignment operators > Subtraction Assignment (-=) [LET]"
type: "concept"
---

# Subtraction Assignment (-=) [LET]

> The -= operator used after the LET keyword, assigns a variable by subtracting the current variable value to a numeric expression.

## Syntax

```
variable -= expr
```

> **Important:**
>
> This operator cannot be used as an element of an
> expression, it is only allowed as assignment operator with the [`LET`](0701-let.md "The LET statement assigns values to variables.") instruction.

## Usage

The `-=` assignment operator subtracts the right-hand numeric expression to the
value of the left-hand variable and assigns the result to the variable.

## Example

```
MAIN
  DEFINE x INTEGER
  LET x = 500
  LET x -= 100
END MAIN
```
