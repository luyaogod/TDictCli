---
title: "Moving screen array rows with SCROLL"
source: "fgl-topics/c_fgl_MigI4GL_014.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > User interface topics > Moving screen array rows with SCROLL"
type: "concept"
description: "IBM Informix 4GL programs can use the SCROLL instruction to move screen array rows up and down, by preserving the TTY attributes."
---

# Moving screen array rows with SCROLL

> IBM® Informix® 4GL programs can use the SCROLL instruction to move screen array rows up and down, by preserving the TTY attributes.

Data records can be displayed in a screen array with a `DISPLAY
array[array-index].* TO screen-array[screen-line]` instruction,
optionally with the `ATTRIBUTES()` clause to use some TTY attributes like colors,
reverse and bold effects. When moving screen array rows with the `SCROLL`
instruction, I4GL actually uses the terminal scrolling capabilities to preserve the TTY attributes
in each row. This applies only to the rows visible on the screen.

In order to display application screens on different types of front-ends, Genero BDL handles user
interface elements in a more abstract way. A good replacement for `DISPLAY ... TO ...
ATTRIBUTES()` and `SCROLL` instruction is to use a regular `DISPLAY
ARRAY` or `INPUT ARRAY` dialog, with the
`DIALOG.setArrayAttributes()` method, to define decoration attributes for each
cell.

## Related links

**Related concepts**  

[ui.Dialog.setArrayAttributes](../15_library-reference/3219-ui-dialog-setarrayattributes.md "Define cell decoration attributes array for the specified list (singular or multiple dialogs).")
