---
title: "util.Integer.testBit"
source: "fgl-topics/c_fgl_ext_util_Integer_testBit.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Integer class > util.Integer methods > util.Integer.testBit"
type: "concept"
---

# util.Integer.testBit

> Returns TRUE, if in the INTEGER value, the bit at the designed position is set.

## Syntax

```
util.Integer.testBit(
  i INTEGER
  n INTEGER
 )
  RETURNS BOOLEAN
```

1. i is an integer value to check.
2. n is the bit position (LSB is at position zero).

## Usage

The `util.Integer.clearBit()` method returns `TRUE` if the passed
integer value has the bit set to 1 at the specified position.

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
    DISPLAY util.Integer.testBit( 3, 1 ) -- displays 1 (TRUE)
END MAIN
```

## Related links

**Related concepts**  

[util.Integer.clearBit](3514-util-integer-clearbit.md "Returns the INTEGER parameter with the bit at the designated position set to 0.")

[util.Integer.setBit](3519-util-integer-setbit.md "Returns the INTEGER parameter with the bit at the designated position set to 1.")
