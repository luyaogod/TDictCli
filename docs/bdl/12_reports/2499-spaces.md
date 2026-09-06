---
title: "SPACES"
source: "fgl-topics/c_fgl_reports_SPACES.html"
breadcrumb: "Reports > Report operators > SPACES"
type: "concept"
---

# SPACES

> Generates the given number of blank characters.

## Syntax

```
num-spaces SPACES
```

1. num-spaces is the number of spaces.

## Usage

This operator returns a string of blanks, equivalent to a quoted string containing
the specified number of blank spaces.

In a `PRINT` statement,
these blanks are inserted at the current character position.

Its operand must be an integer expression that returns a positive number, specifying
an offset (from the current character position) no greater than the difference (right
margin - current position). After `PRINT SPACES` has executed, the new
current character position has moved to the right by the specified number of characters.

Outside `PRINT` statements, `SPACES` and its operand
must appear within parentheses: `(n SPACES)`.

## Example

```
 ON EVERY ROW
    LET s = (6 SPACES), "=ZIP"
    PRINT r.fname, 2 SPACES, r.lname, s
```

## Related links

**Related concepts**  

[Integer expressions](../08_language-basics/0595-integer-expressions.md "This section covers integer expression evaluation rules.")
