---
title: "Row highlighting in TREE"
source: "fgl-topics/c_fgl_tree_highlight_row_cell.html"
breadcrumb: "User interface > User interface programming > Tree views > Row highlighting in TREE"
type: "concept"
---

# Row highlighting in TREE

> TREE containers can be configured with presentation styles, to control row highlighting, using specific foreground and background colors.

## Purpose of row highlighting style attributes

`TREE` containers use a default rendering for the current row (i.e. tree node)
when controlled by `DISPLAY ARRAY` dialog.

The following style attributes can be combined to change the rendering of the current node in a
treeview container:

- `highlightCurrentRow (yes/no)`: Controls current row highlighting.
- `highlightColor`: Defines the background color.
- `highlightTextColor`: Defines the text/foreground color.

The `highlightColor` / `highlightTextColor` can get [color](1622-colors.md "When providing a value for style attributes that define color, you can specify a generic color name or its RGB value.") values. If these attributes are not specified
and highlighting is required, the front-end will use platform default highlighting colors.

A `TREE` style can for example be defined as
follows:

```
<Style name="Table.custom_style">
  <StyleAttribute name="highlightCurrentRow" value="yes"/>
  <StyleAttribute name="highlightColor" value="red"/>
  <StyleAttribute name="highlightTextColor" value="yellow"/>
</Style>
```

## Treeview controlled by DISPLAY ARRAY

When using a `DISPLAY ARRAY` to control a TREE container, the rendering of the
current row is as follows:

- By default, the current row is highlighted.
- If `highlightCurrentRow=yes`, the current row is highlighted.
- If `highlightCurrentRow=no`, the current row is not highlighted.
- The attribute `highlightCurrentCell` is ignored.

![Treeview rendering with with current row highlight color green and white text.](../_images/treeview_highlight_current_row_1.jpg)

*Treeview with current row highlight using DISPLAY ARRAY*

## Related links

**Related reference**  

[Tree style attributes](1649-tree-style-attributes.md "Tree presentation style attributes apply to the TREE container.")
