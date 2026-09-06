---
title: "util.Integer.shiftRight"
source: "fgl-topics/c_fgl_ext_util_Integer_shiftRight.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Integer class > util.Integer methods > util.Integer.shiftRight"
type: "concept"
---

# util.Integer.shiftRight

> Returns the INTEGER value right-shifted by the given bit places.

## Syntax

```
util.Integer.shiftRight(
  i INTEGER,
  n SMALLINT
 )
  RETURNS INTEGER
```

1. i is the integer value to shift.
2. n is the right-shift distance in bits.

## Usage

The `util.Integer.shiftRight()` method shifts the bits to the right in the integer
value passed in the first parameter by the number of places specified by the second parameter.

```
   i   : 210 (00000000 11010010)
   n   : 3
--------------------------------
result :  26 (00000000 00011010)
```

Bits shifted out from the right end are lost.

Bits shifted in from the left end are set to 0.

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
    DISPLAY util.Integer.shiftRight(
                 util.Integer.parseBinaryString("1100"),
                 2 ) -- displays 3
END MAIN
```
