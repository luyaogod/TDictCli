---
title: "ON EXPAND block"
source: "fgl-topics/c_fgl_dialog_ON_EXPAND_3.html"
breadcrumb: "User interface > Dialog instructions > Declarative dialogs (DIALOG - at module level) > Using declarative dialogs > DIALOG data blocks > ON EXPAND block"
type: "concept"
description: "Syntax ON EXPAND ( row-index ) instruction [...] Usage The ON EXPAND block is executed when a tree view node is expanded (opened). This data block is only used in DISPLAY ARRAY dialog blocks. This ..."
---

# ON EXPAND block

## Syntax

```
ON EXPAND (row-index)
   instruction [...]
```

## Usage

The `ON EXPAND` block is executed when a tree view node is expanded (opened).

This data block is only used in `DISPLAY
ARRAY` dialog blocks.

This data block takes an `INTEGER` variable as argument, to get the index of the
expanded row in the dynamic array containing tree nodes:

```
DEFINE row_index INTEGER
...
DISPLAY ARRAY tree_arr TO sr.*
    ...
    ON EXPAND (row_index)
       DISPLAY "Expanded node: ", tree_arr[row_index].id
```

The `ON EXPAND` dialog data block can be used with a static tree view
`DISPLAY ARRAY` dialog, or with a dynamic tree view `DISPLAY ARRAY`
dialog.

When used in a static tree view dialog, `ON EXPAND` is a trigger indicating that a
node was expanded.

In a dynamic tree view dialog, `ON EXPAND` is a trigger to fill the dynamic array
model with children nodes for the parent node that is expanded. For more details, see [Dynamic filling of very large trees](2362-dynamic-filling-of-very-large-trees.md "How to optimize the implementation of large tree-views?").

## Related links

**Related concepts**  

[ON COLLAPSE block](1970-on-collapse-block.md "ON COLLAPSE block")
