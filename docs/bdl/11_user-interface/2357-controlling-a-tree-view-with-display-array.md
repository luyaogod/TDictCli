---
title: "Controlling a tree-view with DISPLAY ARRAY"
source: "fgl-topics/c_fgl_treeviews_007.html"
breadcrumb: "User interface > User interface programming > Tree views > Controlling a tree-view with DISPLAY ARRAY"
type: "concept"
---

# Controlling a tree-view with DISPLAY ARRAY

> A DISPLAY ARRAY dialog needs to be used to control a tree-view.

After the program array has been filled, you must execute a `DISPLAY ARRAY`
dialog.

The following code example implements a `DISPLAY ARRAY` binding the program array
called *tree\_arr* to the *sr* screen-array, attaching the dialog to the tree table defined
in the
form:

```
CALL fill_tree(tree_arr)
DISPLAY ARRAY tree_arr TO sr.* ATTRIBUTES(UNBUFFERED)
  BEFORE ROW
    DISPLAY "Current row is: ", DIALOG.getCurrentRow("sr")
END DISPLAY
```

It is not possible to use the `DISPLAY ARRAY` paged mode (`ON FILL
BUFFER`) when the decoration is a tree view list. The dialog needs the complete set of open
nodes with parent/child relation to handle the tree view display, with the paged mode only a given
window of the dataset is known by the dialog. If you use the paged mode in `DISPLAY
ARRAY` with a tree view as decoration, the program will raise an error at runtime.

However, tree-views can be filled dynamically with `ON EXPAND` / `ON
COLLAPSE` triggers.

## Related links

**Related concepts**  

[Dynamic filling of very large trees](2362-dynamic-filling-of-very-large-trees.md "How to optimize the implementation of large tree-views?")
