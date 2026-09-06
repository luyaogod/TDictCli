---
title: "STRING.trimRight"
source: "fgl-topics/c_fgl_datatypes_STRING_trimRight.html"
breadcrumb: "Library reference > Built-in packages > BDL data types package > STRING data type as class > STRING data type methods > STRING.trimRight"
type: "concept"
---

# STRING.trimRight

> Removes trailing blank space (ASCII 32) characters.

## Syntax

```
trimRight( )
       RETURNS STRING
```

## Usage

The `trimRight()` method removes the trailing blank space characters of the
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
  LET s = "Some text     "
  DISPLAY "["||s.trimRight()||"]"
END MAIN
```

Output:

```
[Some text]
```

## Related links

**Related concepts**  

[STRING.trimRightWhiteSpace](2937-string-trimrightwhitespace.md "Removes trailing whitespace characters.")
