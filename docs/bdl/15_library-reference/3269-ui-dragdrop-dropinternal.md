---
title: "ui.DragDrop.dropInternal"
source: "fgl-topics/c_fgl_ClassDragDrop_dropInternal.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The DragDrop class > ui.DragDrop methods > ui.DragDrop.dropInternal"
type: "concept"
---

# ui.DragDrop.dropInternal

> Perform built-in row drop in trees.

## Syntax

```
dropInternal()
```

## Usage

In order to simplify drag &
drop programming in the same list, the `ui.DragDrop` class
provides the `dropInternal()` utility method, to be
called in the `ON DROP` block. This method will perform
all the row changes in the array and move row selection as well as
cell attributes.

When implementing drag & drop on a [tree-view](../11_user-interface/2352-tree-views.md "Describes how to implement tree views."),
dropping an element on the tree requires complex code in order to handle parent-child relationships.
Nodes can be inserted under a parent between two children, appended at the end of the children list,
and at different levels in the tree hierarchy. However, the `dropInternal()` method
can also be used on simple lists displayed in a regular `TABLE`.

A
call to `dropInternal()` will silently be ignored,
if the drag source is not the drop target, or if the method is called
in a different context as `ON DROP`.

For more
details about dropping elements in tree-views, see [Drag & drop](../11_user-interface/2370-drag-drop.md "Explains programming techniques for the drag & drop feature.").
