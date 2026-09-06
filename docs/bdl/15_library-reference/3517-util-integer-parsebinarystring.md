---
title: "util.Integer.parseBinaryString"
source: "fgl-topics/c_fgl_ext_util_Integer_parseBinaryString.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Integer class > util.Integer methods > util.Integer.parseBinaryString"
type: "concept"
---

# util.Integer.parseBinaryString

> Returns an INTEGER from its binary (base 2) string representation

## Syntax

```
util.Integer.parseBinaryString(
  s STRING
 )
  RETURNS INTEGER
```

1. s is the string in binary format (0 and 1).

## Usage

The `util.Integer.parseBinaryString()` method scans the string as a set of 0 and 1
digits, and converts it to an integer value.

If the binary representation does not fit in an `INTEGER` or if it contains an
invalid binary representation, the method raises the numeric conversion error -1213, that can be
trapped with a `TRY/CATCH` or `WHENEVER ANY ERROR`.

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
    DISPLAY util.Integer.parseBinaryString( "0010" ) -- displays 2
END MAIN
```

## Related links

**Related concepts**  

[util.Integer.toBinaryString](3523-util-integer-tobinarystring.md "Returns the string representation of an INTEGER, as an unsigned integer in base 2.")
