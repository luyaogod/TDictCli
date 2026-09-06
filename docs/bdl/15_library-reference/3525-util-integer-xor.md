---
title: "util.Integer.xor"
source: "fgl-topics/c_fgl_ext_util_Integer_xor.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Integer class > util.Integer methods > util.Integer.xor"
type: "concept"
---

# util.Integer.xor

> Returns the result of a bitwise XOR on two INTEGER values.

## Syntax

```
util.Integer.xor(
  x INTEGER,
  y INTEGER   
 )
  RETURNS INTEGER
```

1. x is an integer value.
2. y is an integer value.

## Usage

The `util.Integer.xor()` method makes a bitwise XOR operation with the integer
values passed as
parameter:

```
 x     : 210 (00000000 00000000 00000000 11010010)
 y     : 135 (00000000 00000000 00000000 10000111)
--------------------------------------------------
result :  85 (00000000 00000000 00000000 01010101)
```

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
    DISPLAY util.Integer.xor( 6, 2 ) -- displays 4
END MAIN
```

## Related links

**Related concepts**  

[util.Integer.and](3512-util-integer-and.md "Returns the result of a bitwise AND on two INTEGER values.")

[util.Integer.andNot](3513-util-integer-andnot.md "Returns the result of a bitwise AND of the 1st INTEGER and the inverted 2nd INTEGER.")

[util.Integer.not](3515-util-integer-not.md "Returns the INTEGER value with all bits inverted.")

[util.Integer.or](3516-util-integer-or.md "Returns the result of a bitwise OR on two INTEGER values.")
