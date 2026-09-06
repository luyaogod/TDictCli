---
title: "util.Strings.collate"
source: "fgl-topics/c_fgl_ext_util_Strings_collate.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Strings class > util.Strings methods > util.Strings.collate"
type: "concept"
---

# util.Strings.collate

> Compares two strings using locale collation rules.

## Syntax

```
util.Strings.collate(
    s1 STRING,
    s2 STRING
 ) RETURNS INTEGER
```

1. s1 is the string to compare to s2.
2. s2 is the string to compare to s1.

## Usage

The `util.Strings.collate()` method compares two strings by following the
collation rules defined by the current [application
locale](../09_advanced-features/0864-application-locale.md "The application locale defines the language and codeset for your application.").

For details about the collation rules that can be used, read the OS
documentation about `setlocale/LC_COLLATE` category.

Note for example that in a French locale, the "é" character has a lower weight than the "è"
character. However, these characters will be considered as equivalent, when other characters follow
these characters in the strings to be compared. For example, "éa" is considered as lower than "èb",
because in this case, "é" and "è" are equivalent and "a" has a lower weight than "b".

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
    CALL test("é","è")
    CALL test("aé","aè")
    CALL test("éa","èa")
    CALL test("éb","èa")
END MAIN

FUNCTION test(s1 CHAR(10), s2 CHAR(10))
    DISPLAY s1, " / ", s2, " => ", util.Strings.collate(s1,s2)
END FUNCTION
```

Output:

```
$ echo $LC_COLLATE 
fr_FR.utf8

$ fglrun -i
Charmap      : UTF-8
Multibyte    : yes
Stateless    : yes
Length Semantics : CHAR

$ fglcomp -M collate.4gl

$ fglrun collate.42m
é         / è         =>          -1
aé        / aè        =>          -1
éa        / èa        =>          -1
éb        / èa        =>          13
```

## Related links

**Related concepts**  

[util.Strings.collateNumeric](3558-util-strings-collatenumeric.md "Compares two strings using locale collation rules and sequences of numerical digits.")
