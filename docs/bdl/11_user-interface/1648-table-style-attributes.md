---
title: "Table style attributes"
source: "fgl-topics/r_fgl_presentation_styles_table_style_attributes.html"
breadcrumb: "User interface > Form definitions > Presentation styles > Style attributes reference > Table style attributes"
type: "reference"
---

# Table style attributes

> Table presentation style attributes apply to a TABLE container.

> **Note:**
>
> This topic lists presentation style attributes for a specific class of form
> element, [common
> presentation style attributes](1629-style-attributes-common-to-all-elements.md "Common presentation style attributes apply to any graphical element, such as windows, layout containers, or form items.") can also be used for this type of element.

## `allowWebSelection`

Users need to select items from a table. Once selected, the keyboard shortcut for copying data
can be used to copy the selection.

Values can be `"yes"` or `"no"` (default).

When `allowWebSelection="no"`, item selection requires the user to hold down the
CTRL key while dragging over the selection with the mouse.

When `allowWebSelection="yes"`, item selection requires a mouse drag only. There
is no need to press the CTRL key. However, the ability to drag-and-drop data is disabled.

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

Values can be `"yes"` or `"no"` (default).

By default, tables are reopened with column positions, visibility and sizes they had when the
window was closed. By setting this attribute to true, the saved settings are ignored and the table
gets the initial column layout. Note that the saved settings include also the sort columns, that
will impact on the order of the rows in the table.

## `headerAlignment`

Defines the label alignment in the column headers of a table.

Values can be:

- `"default"` (default): will use the system default. In most case it is left aligned.
- `"left"` will force all column headers to be left aligned.
- `"center"` will force all column headers to be centered.
- `"right"` will force all column headers to be right aligned.
- `"auto"` will first try to align each column header following the [`JUSTIFY`](1796-justify-attribute.md "The JUSTIFY attribute defines the alignment of a text field content, and table column headers.") attribute of the column. If
  no `JUSTIFY` attribute is set, the column header will be aligned based on the type of
  data: right for numeric data, left for text data.

## `headerHidden`

Defines if the column headers must be visible in a table. This applies to regular list rendering
and [flipped](1780-flipped-attribute.md "The FLIPPED attribute flips TABLE columns into rows.") table rendering.

Values can be `"yes"`
or `"no"` (default).

## `headerPosition`

Defines the position of column headers relative to the field position. This attribute applies
only to [flipped](1780-flipped-attribute.md "The FLIPPED attribute flips TABLE columns into rows.") table rendering.

Values can be `"top"` or `"left"` (default).

With `headerPosition` set to `"top"`, the `headerAlignment` attribute can be used to display column titles on the
top-left, top-center or top-right of the field values.

With `headerAlignment` set to `"auto"`, the column headers are
placed above the field values, depending on the row rendering type:

- With `rowAspect` set to `"default"`, the headers will be right
  aligned.
- With `rowAspect` set to `"list"`, the headers will be left
  aligned.

## `highlightColor`

Defines the highlight color of rows for the `TABLE`, used for selected rows.

For possible values, see [Colors](1622-colors.md "When providing a value for style attributes that define color, you can specify a generic color name or its RGB value.").

For more details see [Row and cell highlighting in TABLE](2328-row-and-cell-highlighting-in-table.md "TABLE containers can be configured with presentation styles, to control row and cell highlighting, using specific foreground and background colors.").

## `highlightCurrentCell`

Indicates if the current cell is highlighted in a `TABLE`, with system default highlight
colors, or the colors defined by the `highlightColor`/`highlightTextColor`
style attributes.

Values can be `"yes"`,`"no"` (default depends on front-end and dialog
type)

For more details see [Row and cell highlighting in TABLE](2328-row-and-cell-highlighting-in-table.md "TABLE containers can be configured with presentation styles, to control row and cell highlighting, using specific foreground and background colors.").

## `highlightCurrentRow`

Indicates if the current row is highlighted in a `TABLE`, with the system default
highlight colors, or the colors defined by the
`highlightColor`/`highlightTextColor` style attributes.

Values can be `"yes"` or `"no"` (default depends on front-end and dialog
type)

For more details see [Row and cell highlighting in TABLE](2328-row-and-cell-highlighting-in-table.md "TABLE containers can be configured with presentation styles, to control row and cell highlighting, using specific foreground and background colors.").

## `highlightTextColor`

Defines the highlighted text color of rows for the `TABLE`, used for selected rows.

For possible values, see [Colors](1622-colors.md "When providing a value for style attributes that define color, you can specify a generic color name or its RGB value.").

For more details see [Row and cell highlighting in TABLE](2328-row-and-cell-highlighting-in-table.md "TABLE containers can be configured with presentation styles, to control row and cell highlighting, using specific foreground and background colors.").

## `imageBorderRadius`

The `imageBorderRadius` style attribute defines the
rounder border radius for images. This attribute accepts the same values as the CSS
`border-radius` property. For example, you can make circular images with a value of
`50%`, or get a little rounded corner with `12px`.

When used for a `TABLE` container, the `imageBorderRadius` style
attribute defines the border radius for [`IMAGECOLUMN`](1786-imagecolumn-attribute.md "The IMAGECOLUMN attribute defines the form field containing the image for the current field.") rendering.

## `leftFrozenColumns`

Requires `"tableType"` set to `"frozenTable"`.

Defines how many columns are frozen, starting from the left of the table.

Values can be any numeric value matching the number of columns.

Default is `"0"`.

## `reduceFilter`

Controls the usage of a reduce filter, to limit the visible rowset in the list container.

Values can be `"yes"` (default on mobile) or `"no"` (default on
desktop).

The reduce filter is only available when all data rows are in memory.

For more details see [List reduce filter](2327-list-reduce-filter.md "The reduce filter allows a user to limit the row set in the list by using a filter.").

## `resizeFillsEmptySpace`

Defines if the resize of the table adapts the size of the last column to avoid unused space.

Values can be `"yes"` or `"no"` (default).

## `rightFrozenColumns`

Requires `"tableType"` set to `"frozenTable"`.

Defines how many columns are frozen, starting from the right of the table.

Values can be any numeric value matching the number of columns.

Default is `"0"`.

## `rowActionTrigger`

Defines the physical event that will fire the row selection action ([`DOUBLECLICK`](2330-defining-the-action-for-a-row-choice.md "The row choice in a TABLE can be associated with a dedicated action.")) on a table, tree or
scrollgrid row.

Values can be:

- `"singleClick"`: Physical event is a single click on desktop platforms, and
  single tap on touchscreen devices.
- `"doubleClick"` (default): Physical event is a double click on desktop platforms,
  and double tap on touchscreen devices.

## `rowAspect`

Defines the aspect of `TABLE` rows.

Values can be:

- `"default"`: The rows get a standard/classic table row rendering, and the current
  row is highlighted according to the presentation styles dedicated for this purpose (such as `highlightColor`)
- `"list"`: When the `TABLE` is in [`FLIPPED`](1780-flipped-attribute.md "The FLIPPED attribute flips TABLE columns into rows.") mode, the rows are rendered
  as elements of a Material Design list, with a thin line at the bottom of each row, and instead of a
  background color, the current row gets a thick vertical bar on the left, using `highlightColor`. In addition, if the table row has the `IMAGECOLUMN` attribute specified on the
  first column, that image is displayed as a thumbnail in front of all column values rather than in
  the same table cell as the column on which it is defined.

See also [Row aspect control](2319-controlling-table-rendering.md).

## `rowHover`

Controls specific background rendering of a row on mouse hover.

Values can be `"yes"` (default) or `"no"`.

Do not confuse with [row/cell highlighting](2328-row-and-cell-highlighting-in-table.md "TABLE containers can be configured with presentation styles, to control row and cell highlighting, using specific foreground and background colors.")
style attributes.

## `showGrid`

Indicates if the vertical and/or horizontal grid lines must be visible in a table.

This style attribute applies only to classic table/tree rendering, as a
set of horizontal rows.

Values can be:

- `"yes"`: Show vertical and horizontal lines (this is the default with
  `INPUT ARRAY`)
- `"no"`: Do not show lines (this is the default with `DISPLAY
  ARRAY`)
- `"vertical"`: Show only vertical lines.
- `"horizontal"`: Show only horizontal lines.

## `tableType`

Defines the rendering type of the table.

> **Important:**
>
> The `tableType` style attribute cannot be changed dynamically,
> once the widget has been displayed.

Values can be:

- `"normal"` (default): Regular table rendering.
- `"frozenTable"`: Users can "freeze" some columns when scrolling, so that they
  always remain visible. Default frozen columns can be defined with
  `"leftFrozenColumns"` and `"rightFrozenColumns"`
  attributes.
- `"listView"` (deprecated): The table is rendered as a list view with two
  columns.
  > **Important:**
  >
  > `tableType listView` is deprecated: As replacement, consider using a [`TABLE`](2315-table-views.md "Describes how to implement table/list views.") with the [`FLIPPED`](1780-flipped-attribute.md "The FLIPPED attribute flips TABLE columns into rows.") attribute, in conjunction with `rowAspect` style attribute set to `"list"`.

## Related links

**Related concepts**  

[TABLE item definition](1746-table-item-definition.md "Defines attributes for a table layout tag.")

[STYLE attribute](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element.")

[Summary lines in tables](2329-summary-lines-in-tables.md "Table views can display a summary line, to show aggregate values for columns.")
