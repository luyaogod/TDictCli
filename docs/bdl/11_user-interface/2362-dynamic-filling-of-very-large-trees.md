---
title: "Dynamic filling of very large trees"
source: "fgl-topics/c_fgl_treeviews_010.html"
breadcrumb: "User interface > User interface programming > Tree views > Dynamic filling of very large trees"
type: "concept"
---

# Dynamic filling of very large trees

> How to optimize the implementation of large tree-views?

When a huge tree needs to be displayed, tree data filling can be
optimized by creating the nodes on demand. There is no need
to fill the complete program array with all possible nodes (down
to the last leaf), when only the first levels/branches of the
tree are displayed on the screen.

To implement a dynamically filled tree, first define an additional column in the
`TREE` container, to indicate whether a given node has children. That field will be
used to render a node with a [+] button, so as to allow the end user click on the node to expand it,
even if no child nodes are created yet.

In the `DISPLAY ARRAY` code, if a node is expanded (or collapsed), the dialog will
invoke the `ON EXPAND` or
`ON COLLAPSE` triggers, to let the
program add (or remove) rows in the array, and to adapt the tree data dynamically by following
navigation events.

```
DEFINE row_index INTEGER
...
DISPLAY ARRAY tree_arr TO sr.* ATTRIBUTES(UNBUFFERED)

  ON EXPAND (row_index)
    DISPLAY "EXPAND   - Expanded row is: ", row_index 
    -- Fill with children nodes for tree_arr[row_index]

  ON COLLAPSE (row_index)
    DISPLAY "COLLAPSE - Collapsed row is: ", row_index 
    -- Remove children nodes of tree_arr[row_index]

END DISPLAY
```

The program array can be filled directly before the dialog execution, but once the dialog has
started, use dialog methods such as `DIALOG.insertNode()` to modify the tree, otherwise internal tree definition will be
corrupted and information like multi-range selection flags and cell attributes will not be
synchronized. This is typically the case when implementing a dynamically-filled tree with `ON
EXPAND`/ `ON COLLAPSE` triggers.

## Related links

**Related concepts**  

[ui.Dialog.appendNode](../15_library-reference/3186-ui-dialog-appendnode.md "Appends a new node in the specified tree-view.")

[ui.Dialog.deleteNode](../15_library-reference/3191-ui-dialog-deletenode.md "Deletes a node from the specified tree-view.")

[Example 2: Dynamic tree view (filled on demand)](2369-example-2-dynamic-tree-view-filled-on-demand.md "Example 2: Dynamic tree view (filled on demand)")
