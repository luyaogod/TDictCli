---
title: "Tree style attributes"
source: "fgl-topics/r_fgl_presentation_styles_tree_style_attributes.html"
breadcrumb: "User interface > Form definitions > Presentation styles > Style attributes reference > Tree style attributes"
type: "reference"
---

# Tree style attributes

> Tree presentation style attributes apply to the TREE container.

> **Note:**
>
> This topic lists presentation style attributes for a specific class of form
> element, [common
> presentation style attributes](1629-style-attributes-common-to-all-elements.md "Common presentation style attributes apply to any graphical element, such as windows, layout containers, or form items.") can also be used for this type of element.

> **Important:**
>
> In the .4st file, the style attribute for TREE
> containers must be specified with the `"Table"` class (in the AUI tree, TREE
> containers are represented with Table nodes)

## `allowWebSelection`

Users need to select items from a treeview. Once selected, the keyboard shortcut for copying data
can be used to copy the selection.

Values can be `"yes"` or `"no"` (default).

When `allowWebSelection="no"`, item selection requires the user to hold down the
CTRL key while dragging over the selection with the mouse.

When `allowWebSelection="yes"`, item selection requires a mouse drag only. There is
no need to press the CTRL key. However, the ability to drag-and-drop data is disabled.

## `alternateRows`

Controls background color rendering of odd/even rows.

Values can be `"yes"` (default) or `"no"`.

This style attribute applies only to classic table/tree rendering, as a
set of horizontal rows.

The style system also supports the [`:odd/:even` pseudo selectors](1612-pseudo-selectors.md "Pseudo selectors can be used to apply style only when some conditions are fulfilled."), to define specific rendering for table/tree
rows (not only background color). When `:odd` or `:even` is used, it
takes precence over the `alternateRows` style attribute. This solution is however
expensive in terms of front-end side processing; Consider using `alternateRows`
instead of `:odd/:even` pseudo selectors.

## `forceDefaultSettings`

By default, tables are reopened with column positions, visibility and sizes they had when the
window was closed. By setting this attribute to true, the saved settings are ignored and the table
gets the initial column layout. Note that the saved settings include also the sort columns, that
will impact on the order of the rows in the table.

Values can be `"yes"` or
`"no"` (default).

## `headerAlignment`

Defines the column header alignment in a table.

Values can be:

- `"default"` (default): use system default. Usually left aligned.
- `"left"` will force all column headers to be left aligned.
- `"center"` will force all column headers to be centered.
- `"right"` will force all column headers to be right aligned.
- `"auto"` will first try to align each column header following the
  `"justify"` attribute of the column. If no `"justify"`
  attribute is set, the column header will be aligned based on the type of data:
  right for numeric data, left for text data.

## `headerHidden`

Defines if the horizontal header must be visible in a treeview.

Values can be `"yes"` or `"no"` (default).

## `highlightColor`

Defines the highlight color of rows for the treeview, used for selected rows.

For possible values, see [Colors](1622-colors.md "When providing a value for style attributes that define color, you can specify a generic color name or its RGB value.").

For more details see [Row highlighting in TREE](2359-row-highlighting-in-tree.md "TREE containers can be configured with presentation styles, to control row highlighting, using specific foreground and background colors.").

## `highlightCurrentRow`

Indicates if the current row is highlighted in a `TREE`, with the system
default highlight colors, or the colors defined by the `highlightColor`/
`highlightTextColor` style attributes.

Values can be `"yes"` or `"no"` (default depends on front-end
and dialog type)

For more details see [Row highlighting in TREE](2359-row-highlighting-in-tree.md "TREE containers can be configured with presentation styles, to control row highlighting, using specific foreground and background colors.").

## `highlightTextColor`

Defines the highlighted text color of rows for the `TREE`, used for
selected rows.

For possible values, see [Colors](1622-colors.md "When providing a value for style attributes that define color, you can specify a generic color name or its RGB value.").

For more details see [Row highlighting in TREE](2359-row-highlighting-in-tree.md "TREE containers can be configured with presentation styles, to control row highlighting, using specific foreground and background colors.").

## `leftFrozenColumns`

Requires `"tableType"` set to `"frozenTable"`.

Defines how many columns are frozen, starting from the left of the treeview.

Values can be any numeric value matching with the number of columns.

Default is `"0"`.

## `resizeFillsEmptySpace`

Defines if the resize of the treeview adapts the size of the last column to avoid unused
space.

Values can be `"yes"` or `"no"` (default).

## `rightFrozenColumns`

Requires `"tableType"` set to `"frozenTable"`.

Defines how many columns are frozen, starting from the right of the treeview list part.

Values can be any numeric value matching with the number of columns.

Default is `"0"`.

## `rowActionTrigger`

Defines the physical event that will fire the row selection action ([`DOUBLECLICK`](2330-defining-the-action-for-a-row-choice.md "The row choice in a TABLE can be associated with a dedicated action.")) on a table, tree or
scrollgrid row.

Values can be:

- `"singleClick"`: Physical event is a single click on desktop platforms, and
  single tap on touchscreen devices.
- `"doubleClick"` (default): Physical event is a double click on desktop platforms,
  and double tap on touchscreen devices.

## `rowHover`

Controls specific background rendering of a row on mouse hover.

Values can be `"yes"` (default) or `"no"`.

Do not confuse with [row/cell highlighting](2328-row-and-cell-highlighting-in-table.md "TABLE containers can be configured with presentation styles, to control row and cell highlighting, using specific foreground and background colors.")
style attributes.

## `showGrid`

Indicates if the vertical and/or horizontal grid lines must be visible in a treeview.

This style attribute applies only to classic table/tree rendering, as a
set of horizontal rows.

Values can be:

- `"yes"`: Show vertical and horizontal lines.
- `"no"`: Do not show lines (this is the default with `DISPLAY
  ARRAY`)
- `"vertical"`: Show only vertical lines.
- `"horizontal"`: Show only horizontal lines.

## `tableType`

Defines the rendering type of the list part of the treeview.

Values can be:

- `"normal"` (default): Regular treeview rendering.
- `"frozenTable"`: Users can "freeze" some columns from scrolling, so that they
  always remain visible. Default frozen columns can be defined with
  `"leftFrozenColumns"` and `"rightFrozenColumns"` attributes.

## Related links

**Related concepts**  

[TREE item definition](1749-tree-item-definition.md "Defines attributes for a tree layout tag.")
