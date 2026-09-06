---
title: "ON COLLAPSE block"
source: "fgl-topics/c_fgl_dialog_ON_COLLAPSE_3.html"
breadcrumb: "User interface > Dialog instructions > Declarative dialogs (DIALOG - at module level) > Using declarative dialogs > DIALOG data blocks > ON COLLAPSE block"
type: "concept"
description: "Syntax ON COLLAPSE ( row-index ) instruction [...] Usage The ON COLLAPSE block is executed when a tree view node is collapsed (closed). This data block is only used in DISPLAY ARRAY dialog blocks. ..."
---

# ON COLLAPSE block

## Syntax

```
ON COLLAPSE (row-index)
   instruction [...]
```

## Usage

The `ON COLLAPSE` block is executed when a tree view node is collapsed
(closed).

This data block is only used in `DISPLAY
ARRAY` dialog blocks.

This data block takes an `INTEGER` variable as argument, to get the index of the
collapsed row in the dynamic array containing tree nodes:

```
DEFINE row_index INTEGER
...
DISPLAY ARRAY tree_arr TO sr.*
    ...
    ON COLLAPSE (row_index)
       DISPLAY "Collapsed node: ", tree_arr[row_index].id
```

The `ON COLLAPSE` dialog data block can be used with a static tree view
`DISPLAY ARRAY` dialog, or with a dynamic tree view `DISPLAY ARRAY`
dialog.

When used in a static tree view dialog, `ON COLLAPSE` is a trigger indicating that
a node was collapsed.

In a dynamic tree view dialog, `ON COLLAPSE` is a trigger to update the dynamic
array model to remove the children nodes of the parent node that is collapsed. For more details, see
[Dynamic filling of very large trees](2362-dynamic-filling-of-very-large-trees.md "How to optimize the implementation of large tree-views?").

## Related links

**Related concepts**  

[ON EXPAND block](1969-on-expand-block.md "ON EXPAND block")
