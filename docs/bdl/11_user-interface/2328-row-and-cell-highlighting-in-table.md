---
title: "Row and cell highlighting in TABLE"
source: "fgl-topics/c_fgl_table_highlight_row_cell.html"
breadcrumb: "User interface > User interface programming > Table views > Row and cell highlighting in TABLE"
type: "concept"
---

# Row and cell highlighting in TABLE

> TABLE containers can be configured with presentation styles, to control row and cell highlighting, using specific foreground and background colors.

## Purpose of row/cell highlighting style attributes

`TABLE` containers use a default rendering for the current row when controlled by
`DISPLAY ARRAY` dialog, and a default rendering for the current cell, when using
`INPUT ARRAY`, `INPUT` or `CONSTRUCT`.

The following style attributes can be combined to change the rendering of the current row or
current cell:

- `highlightCurrentRow (yes/no)`: Controls current row highlighting.
- `highlightCurrentCell (yes/no)`: Controls current cell highlighting.
- `highlightColor`: Defines the background color.
- `highlightTextColor`: Defines the text/foreground color.

The `highlightColor` / `highlightTextColor` can get [color](1622-colors.md "When providing a value for style attributes that define color, you can specify a generic color name or its RGB value.") values. If these attributes are not specified
and highlighting is required for the row or the cell, the front-end will use platform default
highlighting colors.

When using a dialog allowing user input such as `INPUT ARRAY`, cells can be edited
and text can be selected in editor-based field types. When a highlight colors are used at the cell
level, the front-ends will render the selected text accordingly.

A `TABLE` style can for example be defined as
follows:

```
<Style name="Table.custom_style">
  <StyleAttribute name="highlightCurrentRow" value="no"/>
  <StyleAttribute name="highlightCurrentCell" value="yes"/>
  <StyleAttribute name="highlightColor" value="red"/>
  <StyleAttribute name="highlightTextColor" value="yellow"/>
</Style>
```

## Regular DISPLAY ARRAY (no FOCUSONFIELD attribute)

When using a regular `DISPLAY ARRAY` with row-level focus granularity, the
rendering of the current row is as follows:

- By default, the current row is highlighted.
- If `highlightCurrentRow=yes`, the current row is highlighted.
- If `highlightCurrentRow=no`, the current row is not highlighted.
- The attribute `highlightCurrentCell` is ignored.

![Table rendering with with current row highlight color red and white text.](../_images/table_highlight_current_row_1.jpg)

*Table with current row highlight using DISPLAY ARRAY*

## DISPLAY ARRAY using FOCUSONFIELD attribute

When using a `DISPLAY ARRAY` with [`FOCUSONFIELD`](2305-field-level-focus-in-display-array.md "The DISPLAY ARRAY dialog supports cell-level focus with the FOCUSONFIELD.") attribute:

- By default, the current cell is highlighted (other cells of the current row are not).
- If `highlightCurrentRow=yes`, the current row is highlighted.
- If `highlightCurrentRow=no`, the current row is not highlighted
  (default).
- If `highlightCurrentCell=yes`, the current cell is highlighted (default).
- If `highlightCurrentCell=no`, the current cell is not highlighted.

![Table rendering with with current cell highlight color green and yellow text.](../_images/table_highlight_current_cell_1.jpg)

*Table with current cell highlight using DISPLAY ARRAY + FOCUSONFIELD*

## INPUT ARRAY

When using an `INPUT ARRAY` (and `INPUT`,
`CONSTRUCT`):

- By default, the current cell is not highlighted.
- If `highlightCurrentRow=yes`, the current row is highlighted.
- If `highlightCurrentRow=no`, the current row is not highlighted (default).
- If `highlightCurrentCell=yes`, the current cell is highlighted.
- If `highlightCurrentCell=no`, the current cell is not highlighted (default).

![Table rendering with with current cell highlight color yellow and black text.](../_images/table_highlight_current_cell_2.jpg)

*Table with current cell highlight using INPUT ARRAY*

## Related links

**Related concepts**  

[Row and cell highlighting in SCROLLGRID](2340-row-and-cell-highlighting-in-scrollgrid.md "SCROLLGRID containers can be configured with presentation styles, to control row and cell highlighting, using specific foreground and background colors.")

**Related reference**  

[Table style attributes](1648-table-style-attributes.md "Table presentation style attributes apply to a TABLE container.")
