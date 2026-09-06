---
title: "LSTR() [function]"
source: "fgl-topics/c_fgl_operators_LSTR.html"
breadcrumb: "Language basics > Operators > List of expression elements > Character string operators > LSTR() [function]"
type: "concept"
---

# LSTR() [function]

> The LSTR() operator returns a localized string.

## Syntax

```
LSTR ( str-expr )
```

1. str-expr is a string expression.

## Usage

The `LSTR()` operator
returns a localized string corresponding to the identifier
passed as parameter.

Normally localized strings are automatically
replaced when using the `%"ident"` notation
in the source code. When the localized string identifier is
not known at compile time, use the `LSTR()` function.

## Example

```
MAIN
  DISPLAY LSTR ("str"||123)  -- loads string 'str123'
END MAIN
```

## Related links

**Related concepts**  

[Localized strings](../09_advanced-features/0901-localized-strings.md "Localized strings provide a means of writing applications in which the text of strings can be customized on site.")
