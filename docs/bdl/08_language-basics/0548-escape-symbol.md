---
title: "Escape symbol"
source: "fgl-topics/c_fgl_language_features_escape_symbol.html"
breadcrumb: "Language basics > Syntax features > Escape symbol"
type: "concept"
---

# Escape symbol

> Backslash ( \ ) is the escape character of Genero BDL.

The Genero BDL compiler treats a backslash ( `\` ) as the default escape symbol,
and treats the immediately following symbol as a literal, except for special characters such as
`\r` or `\t`.

To specify anything that includes a literal backslash, enter double
( `\\` ) backslashes wherever a single backslash is
required. Similarly, use `\\\\` to represent a literal
double backslash.

```
MAIN
  DISPLAY "\a"   -- displays a
  DISPLAY "\r"   -- displays CR
  DISPLAY "\n"   -- displays NL
  DISPLAY "\ta"  -- displays <tab>a
  DISPLAY "\\"   -- displays \
  DISPLAY "\\\\" -- displays \\
END MAIN
```

## Related links

**Related concepts**  

[Text literals](0588-text-literals.md "Text literals define a character string in an expression.")
