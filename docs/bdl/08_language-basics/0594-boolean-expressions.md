---
title: "Boolean expressions"
source: "fgl-topics/c_fgl_Expressions_Boolean.html"
breadcrumb: "Language basics > Expressions > Boolean expressions"
type: "concept"
---

# Boolean expressions

> This section covers boolean expression evaluation rules.

Boolean expressions are a combination of [`AND`](0620-and.md "The AND operator is the logical intersection operator."), [`OR`](0621-or.md "The OR operator is the logical union operator."), [`NOT`](0619-not.md "The NOT operator performs a logical negation to invert a boolean expression.") boolean operators, as well as comparison
operators such as [`==`](0609-equal-to-or.md "The == operator checks for equality of two expressions or for two record variables. A single = can be used as alias for ==."), [`>=`](0614-greater-or-equal.md "The >= operator is provided to test whether a value or expression is greater than or equal to another.") or [`!=`](0610-different-from-or.md "The != operator checks for non-equality of two expressions or for two record variables. The <> can be used as alias for !=.").

The result of a boolean expression is a `TRUE` or `FALSE` boolean
value, but it can also be `NULL` if one of the operands is `NULL`.

A boolean value is typically used in an [`IF`](0682-if.md "The IF instruction executes a group of statements conditionally.") block, [`WHILE`](0685-while.md "The WHILE statement executes a block of statements until the specified condition becomes false.") block, or the `WHEN` clause in a [`CASE`](0677-case.md "The CASE instruction specifies statement blocks that must be executed conditionally.") block.

The language provides the [`TRUE`](0573-true.md "TRUE is a predefined constant to be used in boolean expressions.") and
[`FALSE`](0574-false.md "FALSE is a predefined constant to be used in boolean expressions.") predefined constants to
initialize boolean variables or return boolean values from functions.

There are three kind of boolean expressions:

- `expr AND expr`
- `expr OR expr`
- `NOT expr`

The (expr) operands of boolean expressions are boolean values.

> **Important:**
>
> The syntax and semantics of boolean expressions in Genero BDL programs is not
> the same as *Boolean conditions* in SQL, as SQL statements are executed by the database engine.

If both operands are `NULL`, the result is `NULL`. If one of the
operands is `NULL` and the other is non-null, the result depends on the operator type
and the non-null value. For more details, see [`AND`](0620-and.md "The AND operator is the logical intersection operator."), [`OR`](0621-or.md "The OR operator is the logical union operator."), [`NOT`](0619-not.md "The NOT operator performs a logical negation to invert a boolean expression.") operators.

Make sure that both operands using with a boolean or comparison operator are not null.

The following example shows a simple boolean expression using the `AND`
operator:

```
IF a AND b THEN
   DISPLAY "Both a and b are TRUE"
END IF
```

In the next example, a boolean expression uses two comparison
expressions:

```
IF (a == b) AND (a == c) THEN
   DISPLAY "a, b and c are equal"
END IF
```

The `NOT` operator will negate a boolean
expression:

```
IF NOT a THEN
   DISPLAY "a is FALSE"
END IF
```

Use a [`BOOLEAN`](0556-boolean.md "The BOOLEAN data type stores a logical value, TRUE or FALSE.") variable to store
the result of a boolean expression:

```
MAIN
  DEFINE b BOOLEAN
  LET b = ( "a" == "b" )  -- result is FALSE
END MAIN
```

> **Important:**
>
> It is bad practive to use non-boolean operands in boolean expressions, as in
> the following example:
>
> ```
> DEFINE s STRING, cnt INTEGER
> IF s AND cnt>0 THEN
>     ...
> END IF
> ```
>
> Good pratice (for character strings for example) is to do the following:
>
> ```
> IF LENGTH(s)>0 AND cnt>0 THEN
> ```
>
> or:
>
> ```
> IF s IS NOT NULL AND cnt>0 THEN
> ```

If the operand is not of type boolean, it has to be converted to a boolean. If a conversion is
required:

- Any [numeric value](0596-numeric-expressions.md "This section covers numeric expression evaluation rules.") evaluates to
  `FALSE`, if and only if the value is 0.
- Any [character string](0597-string-expressions.md "This section covers string expression evaluation rules.") value
  (`STRING`, `CHAR`, `VARCHAR`) follows the next
  rules:
  - If the string starts with a digit, then this conversion evaluates to `FALSE`, if
    and only if the string to integer conversion returns 0.
  - If the string does not start with a digit, then this conversion evaluates to
    `FALSE` if and only if the string has a length of 0.

  Consider using the expression `(LENGTH(string)>0)` or
  `string IS NOT NULL`, to check that a string contains characters,
  or convert the string to a numeric variable and then test the numeric value.
- [`DATE`](0598-date-expressions.md "This section covers date expression evaluation rules.") values can be converted to
  integers. `MDY(12,31,1899) = 0` and evaluates to `FALSE`. Any other
  date value is different from zero and evaluates to `TRUE`.
- Any other [data type](0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data.") produces a conversion error and
  raises the runtime error [-1260](../15_library-reference/4483-genero-bdl-errors.md).

Below a more complex example of boolean expressions:

```
MAIN
  DEFINE r BOOLEAN, c INTEGER
  LET c = 4
  LET r = ( c!=5 ) AND ( c==2 OR c==4 )
  IF ( r AND canReadFile("config.txt") ) THEN
     DISPLAY "OK"
  END IF
END MAIN
```

If an expression that returns `NULL` is the operand of the [`IS NULL`](0606-is-null.md "The IS NULL operator checks for NULL values.") operator, the value of the
boolean expression is `TRUE`:

```
MAIN
  DEFINE r INTEGER
  LET r = NULL
  IF r IS NULL THEN
     DISPLAY "TRUE"
  END IF
END MAIN
```

Boolean expressions in `CASE`, `IF`, or `WHILE`
statements evaluate to `FALSE`, if any element of the comparison is
`NULL`, except for operands of the `IS NULL` and the `IS NOT
NULL` operator.

If you include a boolean expression in a context where the runtime system expects a number, the
expression is evaluated, and is then converted to an integer by the rules `TRUE=1`
and `FALSE=0`.

```
MAIN
  DEFINE r INTEGER
  LET r = 4 + (1==0)    -- result is 4.
END MAIN
```

## Related links

**Related concepts**  

[Type conversions](0576-type-conversions.md "Explains primitive data type conversion rules of the language.")
