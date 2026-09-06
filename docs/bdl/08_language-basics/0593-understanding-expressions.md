---
title: "Understanding expressions"
source: "fgl-topics/c_fgl_Expressions_002.html"
breadcrumb: "Language basics > Expressions > Understanding expressions"
type: "concept"
---

# Understanding expressions

> This is an introduction to language expressions.

## What is an expression ?

An expression is a sequence of operands, operators, and parentheses that
the runtime system can evaluate as a single value.
Operands are program variables, constants, functions
returning a single value and literal values. Operators are used for
arithmetic or string manipulation, and the parentheses
are used to overwrite precedence of operators.

## Language and SQL expressions

Expressions
in SQL statements are evaluated by the database server, not by the
runtime system. The set of operators that can appear
in SQL expressions resembles the set of language operators,
but they are not identical. A program can include SQL
operators, but these are restricted to SQL statements. Similarly,
most SQL operands are not valid in program expressions.
The SQL identifiers of databases, tables, or columns
can appear in a `LIKE` clause or field name in program
instructions, provided that these SQL identifiers
comply with the naming rules of language. Here are
some examples of SQL operands and operators that cannot appear in
other language expressions:

- SQL identifiers, such as column names
- The SQL keywords `USER` and `ROWID`
- Built-in or aggregate SQL functions that are not part of the language
- The `BETWEEN` and `IN` operators
- The `EXISTS`, `ALL`, `ANY`,
  or `SOME` keywords of SQL expressions

Conversely, you *cannot* include language-specific
operators in SQL expressions. For example:

- Arithmetic operators for exponentiation (`**`)
  and modulus (`MOD`)
- String operators `ASCII`, `COLUMN`,
  `SPACE`, `SPACES`,
  and `WORDWRAP`
- Field operators `FIELD_TOUCHED()`, `GET_FLDBUF()`,
  and `INFIELD()`
- The report operators `LINENO` and `PAGENO`

## Parentheses in expressions

Parentheses
are used as in algebra, to override the default order of precedence
of operators. In mathematics, this use of parentheses represents the
"associative" operator. It is, however, a convention in computer languages
to regard this use of parentheses as delimiters rather than as operators.
(Do not confuse this use of parentheses to specify *operator precedence*
with the use of parentheses to enclose arguments in function calls or to
delimit other lists.)

In this example, the variable `y` is assigned the value of 2.

```
LET y = 15 MOD 3 + 2
```

In this example, `y` is assigned the value of 0 because the parentheses
change the sequence of operations.

```
LET y = 15 MOD (3 + 2)
```

## Related links

**Related concepts**  

[Operators](0601-operators.md "Operators are basic syntax elements that appear in expressions.")

[Literals](0585-literals.md "Describes the syntax of literals (constant values) to be used in sources.")
