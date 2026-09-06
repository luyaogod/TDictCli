---
title: "util.Integer.setBit"
source: "fgl-topics/c_fgl_ext_util_Integer_setBit.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Integer class > util.Integer methods > util.Integer.setBit"
type: "concept"
---

# util.Integer.setBit

> Returns the INTEGER parameter with the bit at the designated position set to 1.

## Syntax

```
util.Integer.setBit(
  i INTEGER,
  n INTEGER
 )
  RETURNS INTEGER
```

1. i is the integer value to modify.
2. n is the bit position (LSB is at position zero).

## Usage

The `util.Integer.setBit()` method modifies the integer value by setting the bit
at the position passed as second parameter.

> **Note:**
>
> The position of the least significant bit (LSB) is the position zero:
>
> - In `00000001` (integer 1), the bit set to 1 is at position 0.
> - In `00100010` (integer 34), the bits set to 1 are at position 1 and 5.

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
    DISPLAY util.Integer.setBit(  0, 0 ) -- displays  1 (00000001)
    DISPLAY util.Integer.setBit(  1, 0 ) -- displays  1 (00000001)
    DISPLAY util.Integer.setBit( 16, 2 ) -- displays 20 (00010100)
END MAIN
```

## Related links

**Related concepts**  

[util.Integer.clearBit](3514-util-integer-clearbit.md "Returns the INTEGER parameter with the bit at the designated position set to 0.")

[util.Integer.testBit](3522-util-integer-testbit.md "Returns TRUE, if in the INTEGER value, the bit at the designed position is set.")
