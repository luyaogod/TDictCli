---
title: "Order of precedence"
source: "fgl-topics/c_fgl_operators_014.html"
breadcrumb: "Language basics > Operators > Order of precedence"
type: "concept"
---

# Order of precedence

> The order of precedence defines in which order the elements of an expression are evaluated.

The following list describes the precedence order of expression elements in the meaning of
operators with operands such as
`operand1+operand2`.

For example, the `MOD` operator has a higher precedence as the `*`
operator. When computing an expression like `( 33 MOD 2 * 5 )`, the runtime system
first evaluates `(33 MOD 2) = 1` and then evaluates `(1 * 5) = 5`. The
order of evaluation can be changed this by using parentheses: `( 33 MOD ( 2 * 5 ) ) =
3`.

Some language elements such as compounds operators (`+=`) that can only be used in
the `LET` statement are not listed in the operators precedence table.

Language elements such as [`NVL()`](0615-nvl-function.md "The NVL() operator returns the second parameter if the first argument evaluates to NULL.") or predefined variales such as [`SQLSTATE`](0645-sqlstate-variable.md "The SQLSTATE predefined variable returns the code corresponding to the last SQL error.") are not part of this list:
These are similar to regular user functions or variables, and have the lowest order of precedence.
The function-like and variable-like language elements are respectively marked with the
[function] and [variable] indicators in there reference
topics.

| P | Syntax Element | A | Description | Example |
| --- | --- | --- | --- | --- |
| 19 | UNITS precision | L | Single-qualifier interval | 12 UNITS DAY |
| 18 | + (unary) | L | Unary plus | +324 |
| 18 | - (unary) | L | Unary minus | -324 |
| 17 | ** | L | Exponentiation | x ** 5 |
| 16 | MOD | L | Modulus | x MOD 2 |
| 16 | * | L | Multiplication | x * y |
| 16 | / | L | Division | x / y |
| 15 | + | L | Addition | x + y |
| 15 | - | L | Subtraction | x - y |
| 14 | \|\| | L | Concatenation | "Amount:" \|\| amount |
| 13 | [NOT] LIKE | L | String comparison | mystring LIKE "A%" |
| 13 | [NOT] MATCHES | L | String comparison | mystring MATCHES "A*" |
| 12 | [NOT] IN () | L | List comparison | var IN ('CA','NY') |
| 11 | < | L | Less than | var < 100 |
| 11 | <= | L | Less then or equal to | var <= 100 |
| 11 | > | L | Greater than | var > 100 |
| 11 | >= | L | Greater than or equal to | var >= 100 |
| 11 | == | L | Equals | var == 100 |
| 11 | <> or != | L | Not equal to | var <> 100 |
| 10 | ASCII | R | ASCII Character | ASCII 32 |
| 10 | COLUMN | R | Begin line mode display | PRINT COLUMN 32, "a" |
| 9 | SPACE[S] | L | Insert blank spaces | DISPLAY "a", 5 SPACES |
| 8 | CLIPPED | L | Delete trailing blanks | DISPLAY string CLIPPED |
| 8 | WORDWRAP [RIGHT MARGIN] | L | Report wordwrapping | PRINT string WORDWRAP |
| 8 | IS [NOT] NULL | L | Test for NULL | var IS NULL |
| 7 | NOT | L | Logical inverse | NOT expr |
| 6 | AND | L | Logical intersection | expr1 AND expr2 |
| 5 | OR | L | Logical union | expr1 OR expr2 |
| 4 | WHERE | L | Report aggregate condition | PRINT COUNT(*) WHERE var>0 |
| 3 | USING | L | Format character string | TODAY USING "yy/mm/dd" |
| 2 | := | L | Assignment | var := "abc" |
| 1 | INSTANCEOF | L | Type checking | var INSTANCEOF java.lang.Boolean |

In this table:

- The `P` column defines the precedence, from highest to lowest. Note that some
  operators have the same precedence (i.e. are equivalent in evaluation order).
- The `A` column defines the direction of associativity (L=left-associative,
  R=right-associative). Consider the expression `x % y % z`. If the operator
  `%` has left associativity, it would be interpreted as `(x % y) % z`.
  If the operator has right associativity, it would be interpreted as `x % (y % z)`.
