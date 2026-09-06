---
title: "util.Integer.shiftLeft"
source: "fgl-topics/c_fgl_ext_util_Integer_shiftLeft.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Integer class > util.Integer methods > util.Integer.shiftLeft"
type: "concept"
---

# util.Integer.shiftLeft

> Returns the INTEGER value left-shifted by the given bit places.

## Syntax

```
util.Integer.shiftLeft(
  i INTEGER,
  n SMALLINT
 )
  RETURNS INTEGER
```

1. i is the integer value to shift.
2. n is the left-shift distance in bits.

## Usage

The `util.Integer.shiftLeft()` method shifts the bits to the left in the integer
value passed in the first parameter by the number places specified by the second
parameter.

```
   i   : 210 (00000000 11010010)
   n   : 3
--------------------------------
result :1680 (00000110 10010000)
```

Bits shifted out from the left end are lost.

Bits shifted in from the right end are set to 0.

> **Note:**
>
> Bitwise methods provided by the `util.Integer` class are based on the
> `INTEGER` type. Consider the following facts when using these methods:
>
> 1. The `INTEGER` type is a four-byte signed integer: If the bit at position 31 is
>    set to 1, the corresponding `INTEGER` value will be negative. Thus,
>    `util.Integer.not(0)` produces the `INTEGER` value `-1`
>    (`11111111 11111111 11111111 11111111`).
> 2. The `NULL` value for the `INTEGER` type is represented internally
>    with the value `0x80000000` (`10000000 00000000 00000000 00000000`).
>    When `NULL` is used with the `util.Integer` bitwise methods, it will
>    be interpreted as `0x80000000` instead of a null value. However, if the result of the
>    bitwise operation produces the value `0x80000000`, it will be interpreted as
>    `NULL` when used in an expression.

## Example

```
IMPORT util
MAIN
    DISPLAY util.Integer.shiftLeft(
                 util.Integer.parseBinaryString("11"),
                 3 ) -- displays 24
END MAIN
```
