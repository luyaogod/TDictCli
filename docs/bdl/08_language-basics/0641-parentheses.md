---
title: "Parentheses: ()"
source: "fgl-topics/c_fgl_operators_PARENTHESES.html"
breadcrumb: "Language basics > Operators > List of expression elements > Associative syntax operators > Parentheses: ()"
type: "concept"
---

# Parentheses: ()

> Parentheses ( () ) force the evaluation of an expression before other operators.

## Syntax

```
( expr [...] )
```

1. expr is a language expression.

## Usage

Parentheses can be used to change
the order in which expression elements are evaluated, to bypass the
precedence of operators.

Parentheses can also be used
to ease the readability of the code in a complex expression.

## Example

```
MAIN
  DEFINE n INTEGER
  LET n = ( ( 3 + 2 ) * 2 )
  IF n=10 AND ( n<=0 OR n>=20 ) THEN
    DISPLAY "OK"
  END IF
END MAIN
```

## Related links

**Related concepts**  

[Order of precedence](0602-order-of-precedence.md "The order of precedence defines in which order the elements of an expression are evaluated.")

[Expressions](0592-expressions.md "Shows the possible expressions supported in the language.")
