---
title: "ui.Dialog.insertNode"
source: "fgl-topics/c_fgl_ClassDialog_insertNode.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.insertNode"
type: "concept"
---

# ui.Dialog.insertNode

> Inserts a new node in the specified tree.

## Syntax

```
insertNode(
   name STRING,
   index INTEGER )
```

1. name is the name of the screen record, see [Identifying screen-arrays in ui.Dialog methods](3240-identifying-screen-arrays-in-ui-dialog-methods.md).
2. index is the index of the next sibling node
   in the program array (starts at 1).

## Usage

The `insertNode()` method
is similar to [`insertRow()`](3205-ui-dialog-insertrow.md "Inserts a new row in the specified list."),
except that it has to be used when the list dialog controls a [tree view](../11_user-interface/2352-tree-views.md "Describes how to implement tree views.").

This method
must be used when modifying the array of a tree view during the execution
of the dialog, for example when implementing a [dynamic tree](../11_user-interface/2362-dynamic-filling-of-very-large-trees.md "How to optimize the implementation of large tree-views?") with `ON
EXPAND` / `ON COLLAPSE` triggers. Before the
execution of the dialog, you can fill the program array directly.
This includes the context of `BEFORE DISPLAY` or `BEFORE
DIALOG` control blocks.

When adding rows for a tree view, the id of the parent node and new node matters because that
information is used to build the internal tree structure. When calling
`insertNode()`, you pass the index of the next sibling node. In the program array,
the parent-id member of the new node will be automatically initialized with the value of the
parent-id of the next sibling node, then the internal tree structure is rebuilt.
