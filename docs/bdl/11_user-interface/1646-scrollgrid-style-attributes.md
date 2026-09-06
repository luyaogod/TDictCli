---
title: "ScrollGrid style attributes"
source: "fgl-topics/r_fgl_presentation_styles_scrollgrid_style_attributes.html"
breadcrumb: "User interface > Form definitions > Presentation styles > Style attributes reference > ScrollGrid style attributes"
type: "reference"
---

# ScrollGrid style attributes

> ScrollGrid presentation style attributes apply to SCROLLGRID container.

> **Note:**
>
> This topic lists presentation style attributes for a specific class of form
> element, [common
> presentation style attributes](1629-style-attributes-common-to-all-elements.md "Common presentation style attributes apply to any graphical element, such as windows, layout containers, or form items.") can also be used for this type of element.

## `customWidget`

Defines the display mode to be used for the `SCROLLGRID`.

Values can be:

- `"pagedScrollGrid"`: [Stretchable scrollgrids (defined with `WANTFIXEDPAGESIZE=NO`)](2339-controlling-scrollgrid-rendering.md "Scrollgrid rendering can be controlled by the use of presentation styles and scrollgrid attributes.") will be
  rendered as a responsive tile list. Each row is rendered as a tile. The page size of the scrollgrid
  defines the number of tiles in each page.

Default is to render the scrollgrid as a vertical list.

> **Important:**
>
> The `customWidget` ScrollGrid style attribute cannot be
> changed dynamically, once the widget has been displayed.

## `itemsAlignment`

Defines how items are aligned in [stretchable
scrollgrids (defined with `WANTFIXEDPAGESIZE=NO`)](2339-controlling-scrollgrid-rendering.md "Scrollgrid rendering can be controlled by the use of presentation styles and scrollgrid attributes.").

For the default rendering of stretchable scrollgrids, values of `itemsAlignment`
can be:

- `"stretch"` (default): Each element takes the whole width of the scrollgrid.
- `"left"`: Each element takes its minimal width and aligns to the left of the
  scrollgrid.
- `"center"`: Each element takes its minimal width and aligns to the center of the
  scrollgrid.
- `"right"`: Each element takes its minimal width and aligns to the right of the
  scrollgrid.

For [paged scrollgrids](1646-scrollgrid-style-attributes.md), values of `itemsAlignment` can be:

- `"left"` (default): Elements flow and align to the left of the scrollgrid.
- `"center"`: Elements flow and align to the center of the scrollgrid.
- `"right"`: Elements flow and align to the right of the scrollgrid.

## `highlightColor`

Defines the highlight color of rows for the `SCROLLGRID`, used for selected
rows.

For possible values, see [Colors](1622-colors.md "When providing a value for style attributes that define color, you can specify a generic color name or its RGB value.").

For more details see [Row and cell highlighting in SCROLLGRID](2340-row-and-cell-highlighting-in-scrollgrid.md "SCROLLGRID containers can be configured with presentation styles, to control row and cell highlighting, using specific foreground and background colors.").

## `highlightCurrentCell`

Indicates if the current cell is highlighted in a `SCROLLGRID`, with the
system default highlight colors, or the colors defined by the `highlightColor`/
`highlightTextColor` style attributes.

Values can be `"yes"`,`"no"` (default depends on front-end
and dialog type)

For more details see [Row and cell highlighting in SCROLLGRID](2340-row-and-cell-highlighting-in-scrollgrid.md "SCROLLGRID containers can be configured with presentation styles, to control row and cell highlighting, using specific foreground and background colors.").

## `highlightCurrentRow`

Indicates if the current row is highlighted in a `SCROLLGRID`, with system
default highlight colors, or the colors defined by the `highlightColor`/
`highlightTextColor` style attributes.

Values can be `"yes"` or `"no"` (default depends on front-end
and dialog type)

For more details see [Row and cell highlighting in SCROLLGRID](2340-row-and-cell-highlighting-in-scrollgrid.md "SCROLLGRID containers can be configured with presentation styles, to control row and cell highlighting, using specific foreground and background colors.").

## `highlightTextColor`

Defines the highlighted text color of rows for the `SCROLLGRID`, used for
selected rows.

For possible values, see [Colors](1622-colors.md "When providing a value for style attributes that define color, you can specify a generic color name or its RGB value.").

For more details see [Row and cell highlighting in SCROLLGRID](2340-row-and-cell-highlighting-in-scrollgrid.md "SCROLLGRID containers can be configured with presentation styles, to control row and cell highlighting, using specific foreground and background colors.").

## `reduceFilter`

Controls the usage of a reduce filter, to limit the visible rowset in the list container.

Values can be `"yes"` (default on mobile) or `"no"` (default on
desktop).

The reduce filter is only available when all data rows are in memory.

For more details see [List reduce filter](2327-list-reduce-filter.md "The reduce filter allows a user to limit the row set in the list by using a filter.").

> **Note:**
>
> To allow a reduce filter, the `SCROLLGRID` must be [scrollable](2336-understanding-scrollgrid-views.md "A scrollgrid view defines a graphical element to display a scrolling list of data records in a set of form fields positioned in a grid.").

## `rowActionTrigger`

Defines the physical event that will fire the row selection action ([`DOUBLECLICK`](2330-defining-the-action-for-a-row-choice.md "The row choice in a TABLE can be associated with a dedicated action.")) on a table, tree or
scrollgrid row.

Values can be:

- `"singleClick"`: Physical event is a single click on desktop platforms, and
  single tap on touchscreen devices.
- `"doubleClick"` (default): Physical event is a double click on desktop platforms,
  and double tap on touchscreen devices.

## `rowAspect`

Defines the aspect of `SCROLLGRID` rows.

Values can be:

- `"default"`: The rows get a standard/classic scrollgrid row rendering, and the
  current row is highlighted according to the presentation styles dedicated for this purpose (such as
  `highlightColor`)
- `"list"`: When the `SCROLLGRID` is resizable with [`WANTFIXEDPAGESIZE=NO`](1847-wantfixedpagesize-attribute.md "The WANTFIXEDPAGESIZE attribute controls the vertical resizing of a list element."), the
  rows are rendered as elements of a Material Design list, with a thin line at the bottom of each row,
  and instead of a background color, the current row gets a thick vertical bar on the left, using
  `highlightColor`. The scrollgrid can be decorated with `customWidget` set to `"pagedScrollGrid"`.

See also [Row aspect control](2339-controlling-scrollgrid-rendering.md).

## Related links

**Related concepts**  

[SCROLLGRID item definition](1743-scrollgrid-item-definition.md "Defines attributes for a scrollgrid layout tag.")
