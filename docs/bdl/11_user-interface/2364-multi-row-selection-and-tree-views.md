---
title: "Multi-row selection and tree-views"
source: "fgl-topics/c_fgl_treeviews_012.html"
breadcrumb: "User interface > User interface programming > Tree views > Multi-row selection and tree-views"
type: "concept"
---

# Multi-row selection and tree-views

> Multi-row selection can be used with a DISPLAY ARRAY controlling a TREE container.

Due to the tree-view ergonomic differences with [regular table containers](2331-multi-row-selection-in-tables.md "Multi-row selection can be used with a DISPLAY ARRAY controlling a TABLE container."), the selection of tree nodes follows some specific rules:

- When selecting a range of nodes, only visible nodes will get the selection flag. For example, if
  you select all nodes with Ctrl-A, and if the root node is collapsed, only the root node will be
  selected. This applies also when selecting nodes by program with the `DIALOG.setSelectionRange()`.
- When the `DISPLAY ARRAY` controlling the tree-view implements [dynamic filling](2362-dynamic-filling-of-very-large-trees.md "How to optimize the implementation of large tree-views?"), collapsing a parent node will de-select all
  child nodes, as these are removed from the data model.

![Screenshot of treeview using multi-row selection](../_images/treeview_multirowsel_1.jpg)

*Treeview container with multiple selected rows*

See [Multiple row selection](2314-multiple-row-selection.md "Multiple row selection allows the end user to select several rows within a list of records.") for details about enabling multi-row
selection in a `DISPLAY ARRAY`.
