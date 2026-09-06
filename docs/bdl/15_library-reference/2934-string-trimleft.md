---
title: "STRING.trimLeft"
source: "fgl-topics/c_fgl_datatypes_STRING_trimLeft.html"
breadcrumb: "Library reference > Built-in packages > BDL data types package > STRING data type as class > STRING data type methods > STRING.trimLeft"
type: "concept"
---

# STRING.trimLeft

> Removes leading blank space (ASCII 32) characters.

## Syntax

```
trimLeft( )
       RETURNS STRING
```

## Usage

The `trimLeft()` method removes the leading blank space characters of the current
`STRING` variable and returns a new string.

The method removes only blank space (`ASCII(32)`) characters. Characters tab
(`\t`), newline (`\n`), carriage-return (`\r`) and
form-feed (`\f`) are not removed.

If the original `STRING` variable is `NULL`, the result
will be `NULL`.

## Example

```
MAIN
  DEFINE s STRING
  LET s = "     Some text"
  DISPLAY "["||s.trimLeft()||"]"
END MAIN
```

Output:

```
[Some text]
```

## Related links

**Related concepts**  

[STRING.trimLeftWhiteSpace](2935-string-trimleftwhitespace.md "Removes leading whitespace characters.")
