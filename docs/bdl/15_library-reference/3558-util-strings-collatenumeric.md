---
title: "util.Strings.collateNumeric"
source: "fgl-topics/c_fgl_ext_util_Strings_collateNumeric.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Strings class > util.Strings methods > util.Strings.collateNumeric"
type: "concept"
---

# util.Strings.collateNumeric

> Compares two strings using locale collation rules and sequences of numerical digits.

## Syntax

```
util.Strings.collateNumeric(
    s1 STRING,
    s2 STRING
 ) RETURNS INTEGER
```

1. s1 is the string to compare to s2.
2. s2 is the string to compare to s1.

## Usage

The `util.Strings.collateNumeric()` method compares two strings by following the
collation rules defined by the current [application
locale](../09_advanced-features/0864-application-locale.md "The application locale defines the language and codeset for your application."), and by interpreting any sequence of numerical digits as a positive integer value.

For details about the collation rules that can be used, read the OS
documentation about `setlocale/LC_COLLATE` category.

Sequence of numerical digits are grouped to form whole numbers that will be compared. For
example, `"106M"` is greater than `"51M"`, because
`106` is greater than `51`.

Decimal numbers are not detected, because decimal separators are treated as regular characters:
When comparing `"10.5M"` to `"10.56M"`, the parser identifies the
whole numbers before the dots (`10=10`), then compares the dots
(`.=.`), and then compares the whole number after the dots
(`5<56`). As result, `"10.5M"` compared `"10.56M"`
will produce a negative result. When comparing `"10.500M"` to
`"10.56M"`, the whole numbers after the dots will produce a positive result because
`500` is greater than `56`. However, the result will match decimal
number comparison, if the values are formatted. For example, comparing `"10.500M"` to
`"10.506M"` gives a negative number, because `500<506`.

Negative and positive signs are not included in the whole number detection: When comparing
"`-10M`" to "`+10M`", the minus and plus signs are considered as other
non-digit characters, and the current locale settings and collation rules apply
(`LC_COLLATE`).

When `s1<s2`, the method returns a negative integer,
when `s1>s2`, the result is a positive integer, and when `s1==s2`, the
method returns zero. If one of the strings is `NULL`, it is considered as the lowest
possible value. If both strings are `NULL`, they are considered as equal. When
strings differ, the result may be greater than `1` or lower than `-1`,
and may vary, depending on the locale settings and collation rules of the system. Consequently, do
not compare the result of this method to numbers other than zero.

## Example

Source "collate.4gl":

```
IMPORT util

MAIN
    CALL test(     "10M",       "1M" )
    CALL test(    "150M",     "100M" )
    CALL test(   "10.5M",   "10.56M" )
    CALL test( "10.500M",   "10.56M" )
    CALL test( "10.500M",  "10.560M" )
    CALL test(    "-10M",     "+10M" )
END MAIN

FUNCTION test(s1 CHAR(10), s2 CHAR(10))
    DISPLAY s1, " / ", s2, " => ", util.Strings.collateNumeric(s1,s2)
END FUNCTION
```

Output on Microsoft™ Windows®, with UTF-8 locale:

```
$ fglrun -i
Charmap      : UTF-8
Multibyte    : yes
Stateless    : yes
Length Semantics : CHAR

$ fglrun collate.42m
10M        / 1M         =>           1    -- expected: 10 > 1
150M       / 100M       =>           1    -- expected: 150 > 100
10.5M      / 10.56M     =>          -1    -- expected: 10.5 < 10.56
10.500M    / 10.56M     =>           1    -- unexpected: 10.500 < 10.56
10.500M    / 10.560M    =>          -1    -- correct, because of formatting
-10M       / +10M       =>          -1    -- minus/plus signs processed as chars
```

## Related links

**Related concepts**  

[util.Strings.collate](3557-util-strings-collate.md "Compares two strings using locale collation rules.")
