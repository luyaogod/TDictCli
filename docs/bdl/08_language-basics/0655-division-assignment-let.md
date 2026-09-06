---
title: "Division Assignment (/=) [LET]"
source: "fgl-topics/c_fgl_operators_compassign_divide.html"
breadcrumb: "Language basics > Operators > List of expression elements > Assignment operators > Division Assignment (/=) [LET]"
type: "concept"
---

# Division Assignment (/=) [LET]

> The /= operator used after the LET keyword, assigns a variable by dividing the current variable value by a numeric expression.

## Syntax

```
variable /= expr
```

> **Important:**
>
> This operator cannot be used as an element of an
> expression, it is only allowed as assignment operator with the [`LET`](0701-let.md "The LET statement assigns values to variables.") instruction.

## Usage

The `/=` assignment operator divides the value of the left-hand variable by the
right-hand numeric expression, and assigns the result to the variable.

## Example

```
MAIN
  DEFINE x INTEGER
  LET x = 500
  LET x /= 5
END MAIN
```
