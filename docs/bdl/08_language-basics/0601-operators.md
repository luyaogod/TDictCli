---
title: "Operators"
source: "fgl-topics/c_fgl_operators_001.html"
breadcrumb: "Language basics > Operators"
type: "concept"
---

# Operators

> Operators are basic syntax elements that appear in expressions.

In this documentation, there are different sort of syntax elements that are considered as
operators, such as classical comparison operators (`a==b`), arithmetic operators
(`a+b`), string manipulation operators (`s1||s2`) as well as
predefined variables like `SQLSTATE`, and utility operators like
`SFMT()` or `TODAY`.

Elements of an expression are evaluated by following precedence rules, from highest to lowest.
Use `()` parentheses to instruct the runtime system to evaluate the expression in a
different way than the default order of precedence.

We distinguish different kind of operators:

- Basic operators such as `+` , `||` ,
  `AND` , `ASCII` , which take a left operand, a right operand, or both
  left and right operands. When used together in the same expression, basic operators
  follow the [order of precedence list](0602-order-of-precedence.md "The order of precedence defines in which order the elements of an expression are evaluated.").
- Assignment operators such as `+=`, `||=` , are part
  of the syntax of the [`LET`](0701-let.md "The LET statement assigns values to variables.") statement,
  and can only be used withing this statement. Consequently, assignment operators are not concerned by
  precedence order rules.
- Function operators like `TODAY`, `IIF()`,
  `SFMT()`, where some take parameters enclosed in parentheses. Function
  operators act like functions by returning a value. Function operators have the
  lowest order of precedence when used with basic operators. Function
  operators can be identified with the [function] marker in the title of
  their reference topic.
- Variable operators like `SQLSTATE`, with the same predecence rules
  than function operators. Variable operators can be identified with
  the [variable] marker in the title of their reference topic.

## Related links

**Related concepts**  

[Expressions](0592-expressions.md "Shows the possible expressions supported in the language.")

[Parentheses: ()](0641-parentheses.md "Parentheses ( () ) force the evaluation of an expression before other operators.")

## Child topics

- [Order of precedence](0602-order-of-precedence.md): The order of precedence defines in which order the elements of an expression are evaluated.
- [Operator usage context](0603-operator-usage-context.md): Some operators are specific to a context.
- [List of expression elements](0604-list-of-expression-elements.md): This topic is the reference for language expressions.
