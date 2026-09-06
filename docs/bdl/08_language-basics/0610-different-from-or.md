---
title: "Different from (!= or <>)"
source: "fgl-topics/c_fgl_operators_DIFFERENT.html"
breadcrumb: "Language basics > Operators > List of expression elements > Comparison operators > Different from (!= or <>)"
type: "concept"
---

# Different from (!= or <>)

> The != operator checks for non-equality of two expressions or for two record variables. The <> can be used as alias for !=.

## Syntax 1: Expression comparison

```
expr != expr
```

1. expr can be any expression supported by the language.
2. `<>` can be used as alias for `!=`.

## Syntax 2: Record comparison

```
record1.* != record2.*
```

1. record1 and record2 are records with the same structure.
2. `<>` can be used as alias for `!=`.

## Usage

The `!=` operator evaluates whether two expressions or two records are different.

A less-than sign followed by a greater-than sign (`<>`) can be used as an
alias for the `!=` operator.

This operator applies to expressions that evaluate to primitive data types such
as [`INTEGER`](0562-integer.md "The INTEGER data type is used for storing large whole numbers."), [`VARCHAR`](0570-varchar-size.md "The VARCHAR data type is a variable-length character string data type, with a maximum size."), [`DATE`](0558-date.md "The DATE data type stores calendar dates with a Year/Month/Day representation."). It does not apply to the [`BYTE`](0555-byte.md "The BYTE data type stores any type of binary data, such as images or sounds.") and [`TEXT`](0569-text.md "The TEXT data type stores large text data.") types.

When comparing simple expressions (`expr !=
expr`), the result of the operator is `NULL` when one of
the operands is `NULL`.

When comparing two [records](0715-records.md "Records allow structured program variables definitions.") with the second syntax, the
runtime system compares all corresponding members of the records. If one pair of members are
different, the result of the operator is `TRUE`. When two corresponding members are
`NULL`, they are considered as equal. This second syntax allows you to compare all
members of records, but records must have the same structure.

## Example

```
MAIN
  DEFINE n INTEGER
  LET n=512
  IF n!=32 THEN
     DISPLAY "The variable is not equal to 32."
  END IF
END MAIN
```

## Related links

**Related concepts**  

[Expressions](0592-expressions.md "Shows the possible expressions supported in the language.")
