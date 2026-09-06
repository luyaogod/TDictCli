---
title: "STRING.trimWhiteSpace"
source: "fgl-topics/c_fgl_datatypes_STRING_trimWhiteSpace.html"
breadcrumb: "Library reference > Built-in packages > BDL data types package > STRING data type as class > STRING data type methods > STRING.trimWhiteSpace"
type: "concept"
---

# STRING.trimWhiteSpace

> Removes leading and trailing whitespace characters.

## Syntax

```
trimWhiteSpace( )
       RETURNS STRING
```

## Usage

The `trimWhiteSpace()` method removes the leading and trailing whitespace characters of the
current `STRING` variable and returns a new string.

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
  LET s = "\n\t Some text \n\t"
  DISPLAY "["||s.trimWhiteSpace()||"]"
END MAIN
```

Output:

```
[Some text]
```

## Related links

**Related concepts**  

[STRING.trim](2932-string-trim.md "Removes leading and trailing blank space (ASCII 32) characters.")
