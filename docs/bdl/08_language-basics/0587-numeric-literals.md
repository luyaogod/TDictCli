---
title: "Numeric literals"
source: "fgl-topics/c_fgl_literals_DECIMAL.html"
breadcrumb: "Language basics > Literals > Numeric literals"
type: "concept"
---

# Numeric literals

> Numeric literals define values with a decimal part in an expression.

## Syntax

```
[+|-] digit[...] . digit[...] [ {e|E} [+|-] digit[...] ]
```

1. digit is a digit character from '0' to '9'.
2. Note that the decimal separator is always a dot, independently
   from DBMONEY.
3. The E notation can be used to specify the exponent.

## Usage

Numeric/decimal literals in base-10 notation, without blank spaces and commas, with a decimal
part after a dot.

Numeric literals can be used to specify values for [`DECIMAL(P,S)`](0560-decimal-p-s.md "The DECIMAL data type is provided to handle large numeric values with exact decimal storage."), [`MONEY(P,S)`](0564-money-p-s.md "The MONEY data type is provided to store currency amounts with exact decimal storage."), [`FLOAT`](0561-float.md "The FLOAT data type stores values as double-precision floating-point binary numbers with up to 16 significant digits.") and [`SMALLFLOAT`](0565-smallfloat.md "The SMALLFLOAT data type stores values as single-precision floating-point binary numbers with up to 8 significant digits.") data types.

## Example

```
MAIN
  DEFINE n DECIMAL(10,2)
  LET n = 12345.67
  LET n = -1.23456e-10
END MAIN
```

## Related links

**Related concepts**  

[Numeric expressions](0596-numeric-expressions.md "This section covers numeric expression evaluation rules.")
