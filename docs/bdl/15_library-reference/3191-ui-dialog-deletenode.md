---
title: "ui.Dialog.deleteNode"
source: "fgl-topics/c_fgl_ClassDialog_deleteNode.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.deleteNode"
type: "concept"
---

# ui.Dialog.deleteNode

> Deletes a node from the specified tree-view.

## Syntax

```
deleteNode(
   name STRING,
   index INTEGER )
```

1. name is the name of the screen record, see [Identifying screen-arrays in ui.Dialog methods](3240-identifying-screen-arrays-in-ui-dialog-methods.md).
2. index is the index of the node in the program
   array that has to be deleted (starts at 1).

## Usage

The `deleteNode()` method
is similar to [`deleteRow()`](3192-ui-dialog-deleterow.md "Deletes a row from the specified list."),
except that it has to be used when the dialog controls a [tree view](../11_user-interface/2352-tree-views.md "Describes how to implement tree views.").

This method
must be used when modifying the array of a tree view during the execution
of the dialog, for example when implementing a [dynamic tree](../11_user-interface/2362-dynamic-filling-of-very-large-trees.md "How to optimize the implementation of large tree-views?") with `ON
EXPAND` / `ON COLLAPSE` triggers. Before the
execution of the dialog, you can fill the program array directly.
This includes the context of `BEFORE DISPLAY` or `BEFORE
DIALOG` control blocks.

The main difference with [`deleteRow()`](3192-ui-dialog-deleterow.md "Deletes a row from the specified list.") is
that `deleteNode()` will remove recursively all
child nodes before removing the node identified by index.

If
the index is zero, all root nodes will be deleted from the tree.
