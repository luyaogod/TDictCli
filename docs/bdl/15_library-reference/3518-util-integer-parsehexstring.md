---
title: "util.Integer.parseHexString"
source: "fgl-topics/c_fgl_ext_util_Integer_parseHexString.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Integer class > util.Integer methods > util.Integer.parseHexString"
type: "concept"
---

# util.Integer.parseHexString

> Returns an INTEGER from its hexadecimal (base 16) string representation.

## Syntax

```
util.Integer.parseHexString(
  s STRING
 )
  RETURNS INTEGER
```

1. s is the string in hexadecimal.

## Usage

The `util.Integer.parseHexString()` method scans the hexadecimal string and
converts it to an integer value.
> **Note:**
>
> The `parseHexString()` method is
> case-insensitive for A-F hexadecimal letters: The hexadecimal strings `"ae3f"`,
> `"aE3f"` and `"AE3F"` are equivalent.

If the hexadecimal string does not fit in an `INTEGER` or if it contains an
invalid hexadecimal representation, the method raises the numeric conversion error -1213, that can
be trapped with a `TRY/CATCH` or `WHENEVER ANY ERROR`.

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
    DISPLAY util.Integer.parseHexString("6F12") -- displays 28434
END MAIN
```

## Related links

**Related concepts**  

[util.Integer.toHexString](3524-util-integer-tohexstring.md "Returns the string representation of an INTEGER, as an unsigned integer in base 16.")
