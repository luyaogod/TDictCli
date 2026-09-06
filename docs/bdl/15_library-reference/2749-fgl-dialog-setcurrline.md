---
title: "fgl_dialog_setcurrline()"
source: "fgl-topics/c_fgl_BuiltInFunctions_FGL_DIALOG_SETCURRLINE.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > fgl_dialog_setcurrline()"
type: "concept"
---

# fgl_dialog_setcurrline()

> This function moves to a specific row in a record list.

## Syntax

```
FUNCTION fgl_dialog_setcurrline(
  screenLine INTEGER,
  row INTEGER )
```

1. screenLine is the line number in the form screen array.
2. row is the row number in the program array variable.

## Usage

Moves to the row / screen line specified.
See [fgl\_set\_arr\_curr()](2782-fgl-set-arr-curr.md "Moves to a specific row in a record list.")
for more details.

To be called during a `DISPLAY ARRAY` or `INPUT
ARRAY` instruction, inside `BEFORE DISPLAY` /
`BEFORE INPUT` or `ON ACTION` /
`ON KEY` blocks only.

The screenLine parameter is ignored in GUI mode.

## Related links

**Related concepts**  

[Screen records / arrays](../11_user-interface/1676-screen-records-arrays.md "Form fields can be grouped in a screen record or screen array definition.")
