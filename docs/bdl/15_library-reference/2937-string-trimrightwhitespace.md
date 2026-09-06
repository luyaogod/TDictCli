---
title: "STRING.trimRightWhiteSpace"
source: "fgl-topics/c_fgl_datatypes_STRING_trimRightWhiteSpace.html"
breadcrumb: "Library reference > Built-in packages > BDL data types package > STRING data type as class > STRING data type methods > STRING.trimRightWhiteSpace"
type: "concept"
---

# STRING.trimRightWhiteSpace

> Removes trailing whitespace characters.

## Syntax

```
trimRightWhiteSpace( )
       RETURNS STRING
```

## Usage

The `trimRightWhiteSpace()` method removes the trailing whitespace characters of the current
`STRING` variable and returns a new string.

The method considers as whitespace characters all characters less than or equal to blank space
(`ASCII(32)`). This includes tab (`\t`), newline
(`\n`), carriage-return (`\r`) and form-feed
(`\f`).

If the original `STRING` variable is `NULL`, the result
will be `NULL`.

## Example

```
MAIN
  DEFINE s STRING
  LET s = "Some text  \n\t"
  DISPLAY "["||s.trimRightWhiteSpace()||"]"
END MAIN
```

Output:

```
[Some text]
```

## Related links

**Related concepts**  

[STRING.trimRight](2936-string-trimright.md "Removes trailing blank space (ASCII 32) characters.")
