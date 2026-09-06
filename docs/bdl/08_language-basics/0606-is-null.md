---
title: "IS NULL"
source: "fgl-topics/c_fgl_operators_IS_NULL.html"
breadcrumb: "Language basics > Operators > List of expression elements > Comparison operators > IS NULL"
type: "concept"
---

# IS NULL

> The IS NULL operator checks for NULL values.

## Syntax

```
expr IS [NOT] NULL
```

1. expr can be any expression supported by the language.
2. The `NOT` keyword negates the comparison.

## Usage

The `IS NULL` operator can be used to test whether the left-hand expression is
null, while `IS NOT NULL` operator can be used to test for non-null values.

This operator applies to expressions that evaluate to primitive data types such
as [`INTEGER`](0562-integer.md "The INTEGER data type is used for storing large whole numbers."), [`VARCHAR`](0570-varchar-size.md "The VARCHAR data type is a variable-length character string data type, with a maximum size."), [`DATE`](0558-date.md "The DATE data type stores calendar dates with a Year/Month/Day representation."). It does not apply to the [`BYTE`](0555-byte.md "The BYTE data type stores any type of binary data, such as images or sounds.") and [`TEXT`](0569-text.md "The TEXT data type stores large text data.") types.

Structured variables defined with [`RECORD`](0715-records.md "Records allow structured program variables definitions."), [`DYNAMIC
ARRAY`](0729-arrays.md "Arrays (static or dynamic) allow you to handle an ordered collection of elements."), [`DICTIONARY`](0742-dictionaries.md "A dictionary holds an unordered collection of elements accessed by a key.") types
cannot be used with the `IS [NOT] NULL` operator.

Always use the `IS [NOT] NULL` operator to check for nulls: Comparing an
expression to `NULL` with the [`==`](0609-equal-to-or.md "The == operator checks for equality of two expressions or for two record variables. A single = can be used as alias for ==.") or [`!=`](0610-different-from-or.md "The != operator checks for non-equality of two expressions or for two record variables. The <> can be used as alias for !=.") operators evaluates to `NULL`, which is not
`TRUE`. The next code example displays twice the value 1/`TRUE`,
because both sub-expressions using the equal and not equal operators evaluate to
`NULL`:

```
MAIN
  DEFINE s STRING
  LET s = NULL
  DISPLAY ( (s != NULL) IS NULL )
  DISPLAY ( (s == NULL) IS NULL )
END MAIN
```

## Example

```
MAIN
  DEFINE n INTEGER
  LET n = NULL
  IF n IS NULL THEN
     DISPLAY "The variable is NULL."
  END IF
END MAIN
```

## Related links

**Related concepts**  

[NULL](0572-null.md "The NULL constant defines a non-value.")

[Expressions](0592-expressions.md "Shows the possible expressions supported in the language.")

[NVL() [function]](0615-nvl-function.md "The NVL() operator returns the second parameter if the first argument evaluates to NULL.")
