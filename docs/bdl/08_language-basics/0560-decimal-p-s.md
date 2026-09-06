---
title: "DECIMAL(p,s)"
source: "fgl-topics/c_fgl_datatypes_DECIMAL.html"
breadcrumb: "Language basics > Primitive Data types > DECIMAL(p,s)"
type: "concept"
---

# DECIMAL(p,s)

> The DECIMAL data type is provided to handle large numeric values with exact decimal storage.

## Syntax

```
DECIMAL [ ( precision[,scale]) ]
```

1. precision defines the number of significant digits (limit is
   32, default is 16).
2. scale defines the number of digits to the right of the decimal
   point.
3. When no scale is specified, the data type defines a floating
   point number.
4. When no (precision, scale) is specified,
   it defaults to `DECIMAL(16)`.

## Usage

Use the `DECIMAL` data type when you need to store values that have
a fixed number of digits on the right and left of the decimal point (`DECIMAL(p,s)`),
or to store a floating point decimal with an exact number of significant digits
(`DECIMAL(p)`).

`DEC`, `DECIMAL` and `NUMERIC` are
synonyms.

`DECIMAL` variables are initialized to `NULL` in
functions, modules and globals.

When using `DECIMAL(p,s)` with a precision and scale, you define a decimal for
fixed point arithmetic, with p significant digits and s digits on the right of the decimal point.
For example, `DECIMAL(8,2)` can hold the value 123456.78 (8 (p) = 6 digits on the
left + 2 (s) digits on the right of the decimal point).

When using `DECIMAL(p)` with a precision but no scale, you define a floating-point
number with p significant digits. For example, `DECIMAL(8)` can store 12345678, as
well as 0.12345678.

In most database implementations, the decimal data type always has a fixed number of decimal
digits. Use `DECIMAL` types with precision and scale to implement portable code, and
avoid mistakes if default sizes apply when precisions and/or scale are omitted in SQL statements.
For example, with Oracle®, a
`NUMBER(p)` is equivalent to a `DECIMAL(p,0)` in BDL, not
`DECIMAL(p)`.

When using `DECIMAL` without a precision and scale, it defaults to
`DECIMAL(16)`, a floating-point number with a precision of 16 digits.

```
MAIN
  DEFINE d1 DECIMAL(10,4)
  DEFINE d2 DECIMAL(10,3)
  LET d1 = 1234.4567
  LET d2 = d1 / 3 -- Rounds decimals to 3 digits 
  DISPLAY d1, d2
END MAIN
```

`DECIMAL` values can be converted to strings based on the [DBFORMAT](../07_configuration/0511-dbformat.md "Defines the characters to be used for the currency symbol, decimal and thousands separators for numeric values.") (or DBMONEY) environment variable
(defines the decimal separator) setting.

## Value ranges

The largest absolute value that a `DECIMAL(p,s)` can store without errors is
10p-s - 10s. The stored value can have up to 30 significant decimal
digits in its fractional part, or up to 32 digits to the left of the decimal point.

When using `DECIMAL(p,s)` the range of values is defined by the p, the
number of significant digits. For example, a variable defined as
`DECIMAL(5,3)` can store values in the range -99.999 to 99.999. The
smallest positive non zero value is 0.001.

When using `DECIMAL(p)` the magnitude can range from -N\*10-124 to
N\*10124, where N can have up to p significant digits and be 0<N<10. For
example, a variable defined as `DECIMAL(5)` can store values in the range
-9.9999e-124 to 9.9999e+124. The smallest positive non zero value is 9.9999e-130.

## Exceptions

When the default exception handler is used, if you try to assign a value larger than the
decimal definition (for example, 12345.45 into `DECIMAL(4,2)` ), no
out of range error occurs, and the variable is assigned with `NULL`.
If `WHENEVER ANY ERROR` is used, it raises error [-1226](../15_library-reference/4483-genero-bdl-errors.md). If you do not
use `WHENEVER ANY ERROR`, the status variable is not set to -1226.

Data type conversion can be controlled by catching the runtime exceptions. For more
details, see [Handling type conversion errors](0579-handling-type-conversion-errors.md "Runtime errors can be handled on type conversion failures.").

## Computation and rounding rules

When computing or converting decimal values, the "round half away from zero" rule will
apply: If the fraction of the value v is exactly 0.5, then r = v + 0.5 if v is positive, and r = v -
0.5 if v is negative. For example, when the result must be rounded to a whole number, 23.5 gets
rounded to 24, and -23.5 gets rounded to -24.

In the next example, the division result of 11 / 3 gives the infinite decimal value
3.666666... (with an infinite decimal part). However, this value cannot be stored in a fixed
point decimal type. When stored in a `DECIMAL(10,2)`, the value will be
rounded to 3.67, and when multiplying 3.67 by 3, the result will be 11.01, instead of 11:

```
MAIN
    DEFINE v DECIMAL(10,2)
    LET v = 11 / 3
    DISPLAY "1. v = ", v USING "---&.&&&&&&&&"
    LET v = v * 3
    DISPLAY "2. v = ", v USING "---&.&&&&&&&&"
END MAIN
```

Output:

```
1. v =    3.67000000
2. v =   11.01000000
```

## High-precision math functions

A couple of precision math functions are available, to be used with DECIMAL values.
These functions have a higher precision as the standard C library functions based on C
double data type, which is equivalent to `FLOAT`:

- [FGL\_DECIMAL\_TRUNCATE()](../15_library-reference/2738-fgl-decimal-truncate.md "Returns a decimal truncated to the precision passed as parameter.")
- [FGL\_DECIMAL\_SQRT()](../15_library-reference/2739-fgl-decimal-sqrt.md "Computes the square root of the decimal passed as parameter.")
- [FGL\_DECIMAL\_EXP()](../15_library-reference/2740-fgl-decimal-exp.md "Returns the value of Euler's constant (e) raised to the power of the decimal passed as parameter.")
- [FGL\_DECIMAL\_LOGN()](../15_library-reference/2741-fgl-decimal-logn.md "Returns the natural logarithm of the decimal passed as parameter.")
- [FGL\_DECIMAL\_POWER()](../15_library-reference/2742-fgl-decimal-power.md "Raises decimal to the power of the real exponent.")

## Related links

**Related concepts**  

[MONEY(p,s)](0564-money-p-s.md "The MONEY data type is provided to store currency amounts with exact decimal storage.")
