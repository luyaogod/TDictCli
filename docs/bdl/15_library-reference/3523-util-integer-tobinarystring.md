---
title: "util.Integer.toBinaryString"
source: "fgl-topics/c_fgl_ext_util_Integer_toBinaryString.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Integer class > util.Integer methods > util.Integer.toBinaryString"
type: "concept"
---

# util.Integer.toBinaryString

> Returns the string representation of an INTEGER, as an unsigned integer in base 2.

## Syntax

```
util.Integer.toBinaryString(
  i INTEGER
 )
  RETURNS STRING
```

1. i is the source integer value to convert.

## Usage

The `util.Integer.toBinaryString()` method generates the binary
representation of the integer passed as parameter.

The resulting binary string represents the unsigned integer, in base 2.

> **Note:**
>
> The result has no leading zeros.

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
    DISPLAY util.Integer.toBinaryString( 24 ) -- displays "11000"
END MAIN
```

## Related links

**Related concepts**  

[util.Integer.parseBinaryString](3517-util-integer-parsebinarystring.md "Returns an INTEGER from its binary (base 2) string representation")
