---
title: "ui.Dialog.appendNode"
source: "fgl-topics/c_fgl_ClassDialog_appendNode.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.appendNode"
type: "concept"
---

# ui.Dialog.appendNode

> Appends a new node in the specified tree-view.

## Syntax

```
appendNode(
   name STRING,
   parentIndex INTEGER
 ) RETURNS INTEGER
```

1. name is the name of the screen record, see [Identifying screen-arrays in ui.Dialog methods](3240-identifying-screen-arrays-in-ui-dialog-methods.md).
2. parentIndex is the index of the parent node in the program array (starts at
   1, zero appends a root node).
3. The method returns the array index of the new created node.

## Usage

The `appendNode()` method adds a new node under a given parent,
when the dialog controls a tree view.

This method must be used when modifying the array of a tree view during the
execution of the dialog, for example when implementing a dynamic tree with
`ON EXPAND` / `ON COLLAPSE` triggers. Before the
execution of the dialog, you can fill the program array directly. This includes
the context of `BEFORE DISPLAY` or `BEFORE DIALOG`
control blocks.

When adding rows for a tree view, the id of the parent node and new node matters because that
information is used to build the internal tree structure. When calling
`appendNode()`, you pass the index of the parent node under which the new node will
be appended. In the program array, the parent-id member of the new node will automatically be
initialized with the value of the id of the parent node identified by the index passed as parameter,
then the internal tree structure is rebuilt.

If the parent index is zero, a new root node will be appended.

The method returns the index of the new inserted node.

In the program
array, the parent-id member of the new node will automatically be
initialized with the value of the id member of the parent node
identified by the index.

```
DEFINE x INTEGER
DISPLAY ARRAY mytree TO sr.*
    ...
    ON EXPAND(id)
       LET x = DIALOG.appendNode("sr", id)
       LET mytree[x].name = SFMT("Item %1", x)
    ...
```

## Related links

**Related concepts**  

[Dynamic filling of very large trees](../11_user-interface/2362-dynamic-filling-of-very-large-trees.md "How to optimize the implementation of large tree-views?")
