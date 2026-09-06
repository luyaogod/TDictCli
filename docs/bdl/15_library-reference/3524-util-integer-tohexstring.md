---
title: "util.Integer.toHexString"
source: "fgl-topics/c_fgl_ext_util_Integer_toHexString.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Integer class > util.Integer methods > util.Integer.toHexString"
type: "concept"
---

# util.Integer.toHexString

> Returns the string representation of an INTEGER, as an unsigned integer in base 16.

## Syntax

```
util.Integer.toHexString(
  i INTEGER
 )
  RETURNS STRING
```

1. i is the source integer value to convert.

## Usage

The `util.Integer.toHexString()` method generates the hexadecimal representation
of the integer passed as parameter.

The resulting hexadecimal string represents the unsigned integer, in base 16.

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
    DISPLAY util.Integer.toHexString( 234 ) -- displays "ea"
END MAIN
```

## Related links

**Related concepts**  

[util.Integer.parseHexString](3518-util-integer-parsehexstring.md "Returns an INTEGER from its hexadecimal (base 16) string representation.")
