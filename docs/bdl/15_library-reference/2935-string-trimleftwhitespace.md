---
title: "STRING.trimLeftWhiteSpace"
source: "fgl-topics/c_fgl_datatypes_STRING_trimLeftWhiteSpace.html"
breadcrumb: "Library reference > Built-in packages > BDL data types package > STRING data type as class > STRING data type methods > STRING.trimLeftWhiteSpace"
type: "concept"
---

# STRING.trimLeftWhiteSpace

> Removes leading whitespace characters.

## Syntax

```
trimLeftWhiteSpace( )
       RETURNS STRING
```

## Usage

The `trimLeftWhiteSpace()` method removes the leading whitespace characters of the current
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
  LET s = "\n\t     Some text"
  DISPLAY "["||s.trimLeftWhiteSpace()||"]"
END MAIN
```

Output:

```
[Some text]
```

## Related links

**Related concepts**  

[STRING.trimLeft](2934-string-trimleft.md "Removes leading blank space (ASCII 32) characters.")
