---
title: "Addition Assignment (+=) [LET]"
source: "fgl-topics/c_fgl_operators_compassign_add.html"
breadcrumb: "Language basics > Operators > List of expression elements > Assignment operators > Addition Assignment (+=) [LET]"
type: "concept"
---

# Addition Assignment (+=) [LET]

> The += operator used after the LET keyword, assigns a variable by adding the current variable value to a numeric expression.

## Syntax

```
variable += expr
```

> **Important:**
>
> This operator cannot be used as an element of an
> expression, it is only allowed as assignment operator with the [`LET`](0701-let.md "The LET statement assigns values to variables.") instruction.

## Usage

The `+=` assignment operator adds the right-hand numeric expression to the value
of the left-hand variable and assigns the result to the variable.

## Example

In the next example, the `+=` operator is used to increment an integer variable by
100:

```
MAIN
  DEFINE x INTEGER
  LET x = 500
  LET x += 100
END MAIN
```
