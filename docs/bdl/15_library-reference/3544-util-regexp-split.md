---
title: "util.Regexp.split"
source: "fgl-topics/c_fgl_ext_util_Regexp_split.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Regexp class > util.Regexp methods > util.Regexp.split"
type: "concept"
---

# util.Regexp.split

> Splits a string around matches of the current regular expression.

## Syntax

```
util.Regexp.split(
  subject STRING
 )
  RETURNS DYNAMIX ARRAY OF STRING
```

1. subject is the string to be scanned.

## Usage

This method scans the string passed as first parameter and generates a dynamic array with all
tokens terminated by separators matching the current compiled regular expression.

The tokens can be terminated by the substring that matches the regular expression, or is
terminated by the end of the string.

If the expression does not match any part of the input, then the resulting array has just one
element, which is a copy of this string.

If the regular expression is `NULL`, it matches an empty separator at every
position, splitting the source string into its individual characters. The resulting array then
contains one element per character, preceded and followed by an additional empty string element. For
example, the source string `"abc"` will be splitted into array elements `[ "",
"a", "b", "c", "" ]` .

Splitting with an `NULL` regular expression follows the runtime character set.
With a multibyte character set such as UTF-8, the string is split into characters and multibyte
characters are kept together; with a single-byte character set, the string is split into individual
bytes. This behavior depends on the [runtime locale](../09_advanced-features/0864-application-locale.md "The application locale defines the language and codeset for your application.")
(LC\_ALL/LANG), not on the FGL\_LENGTH\_SEMANTICS setting.

## Example

```
IMPORT util
MAIN
    DEFINE re util.Regexp
    DEFINE tokens DYNAMIC ARRAY OF STRING
    DEFINE x INTEGER
    LET re = util.Regexp.compile(`[ ;:]`)
    LET tokens = re.split("aa bb;cc:dd")
    FOR x=1 TO tokens.getLength()
        DISPLAY tokens[x]
    END FOR
END MAIN
```

Output:

```
aa
bb
cc
dd
```

## Related links

**Related concepts**  

[util.Regexp.compile](3532-util-regexp-compile.md "Compiles a regular expression and returns a util.Regexp object.")

[util.Regexp.replaceFirst](3543-util-regexp-replacefirst.md "Substitutes the match of the regular expression with the replacement string.")

[STRING.split](2928-string-split.md "Splits the current string around matches of the given regular expression.")
