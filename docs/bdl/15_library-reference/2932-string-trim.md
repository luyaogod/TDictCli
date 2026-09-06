---
title: "STRING.trim"
source: "fgl-topics/c_fgl_datatypes_STRING_trim.html"
breadcrumb: "Library reference > Built-in packages > BDL data types package > STRING data type as class > STRING data type methods > STRING.trim"
type: "concept"
---

# STRING.trim

> Removes leading and trailing blank space (ASCII 32) characters.

## Syntax

```
trim( )
       RETURNS STRING
```

## Usage

The `trim()` method removes the leading and trailing blank space characters of the
current `STRING` variable and returns a new string.

The method removes only blank space (`ASCII(32)`) characters. Characters tab
(`\t`), newline (`\n`), carriage-return (`\r`) and
form-feed (`\f`) are not removed.

If the original `STRING` variable is `NULL`, the result
will be `NULL`.

## Example

```
MAIN
  DEFINE s STRING
  LET s = "   Some text   "
  DISPLAY "["||s.trim()||"]"
END MAIN
```

Output:

```
[Some text]
```

## Related links

**Related concepts**  

[STRING.trimWhiteSpace](2933-string-trimwhitespace.md "Removes leading and trailing whitespace characters.")
