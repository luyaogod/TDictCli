---
title: "Integer literals"
source: "fgl-topics/c_fgl_literals_INTEGER.html"
breadcrumb: "Language basics > Literals > Integer literals"
type: "concept"
---

# Integer literals

> Integer literals define a whole number in an expression.

## Syntax

```
[+|-] digit[...]
```

1. digit is a digit character from '0' to '9'.

## Usage

Integer literals are in base-10 notation, without blank spaces and commas and without a decimal
point.

Integer literals can be used to specify values for [`DECIMAL(P,0)`](0560-decimal-p-s.md "The DECIMAL data type is provided to handle large numeric values with exact decimal storage."), [`BIGINT`](0554-bigint.md "The BIGINT data type is used for storing very large whole numbers."), [`INTEGER`](0562-integer.md "The INTEGER data type is used for storing large whole numbers."), [`SMALLINT`](0566-smallint.md "The SMALLINT data type is used for storing small whole numbers.") and [`TINYINT`](0568-tinyint.md "The TINYINT data type is used for storing very small whole numbers.") data types.

## Example

```
MAIN
  DEFINE n INTEGER
  LET n = 1234567
END MAIN
```

## Related links

**Related concepts**  

[Integer expressions](0595-integer-expressions.md "This section covers integer expression evaluation rules.")
