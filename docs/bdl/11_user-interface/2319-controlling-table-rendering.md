---
title: "Controlling table rendering"
source: "fgl-topics/c_fgl_ui_tables_rendering.html"
breadcrumb: "User interface > User interface programming > Table views > Controlling table rendering"
type: "concept"
---

# Controlling table rendering

> Table rendering can be controlled by the use of presentation styles and table attributes.

## Table width

By default, the width of a `TABLE` container is defined by the number of visible
columns in its layout.

Table column fields defined as [`HIDDEN`](1783-hidden-attribute.md "The HIDDEN attribute indicates if an element must be visible to the user.") will not contribute to computing the default table width.

In this example, the table is defined with three visible columns, which in turn define individual
field widths. The sum of the field widths defines the default table width.

```
GRID
{
<TABLE t1           >
[c1  |c2       |c3  ]
...
```

To specify explicitly the width of a table, use the [`WIDTH`](1850-width-attribute.md "The WIDTH attribute forces an explicit width of a form element.")
attribute:

```
TABLE t1 : table1, WIDTH = 5 COLUMNS, ... ;
```

The `COLUMNS` unit means a number of fields/columns in the `TABLE`,
not grid-columns.

The size of a `TABLE` (or `TREE`)
element has an impact on the size of a [modal
window](1567-configuring-windows-with-styles.md "Use the STYLE attribute to set a style for a window.") where the form of the table is displayed. However, regular windows will be sized from
the [window container](1563-containers-for-program-windows.md "Program windows are displayed in window containers by the front-end."), and the form
content will adapt to it.

## Table height

The height of a `TABLE` container is by default defined by the number of rows in
its layout. In the example, the default table height will be three
rows:

```
GRID
{
<TABLE t1           >
[c1  |c2       |c3  ]
[c1  |c2       |c3  ]
[c1  |c2       |c3  ]
<                   >
...
```

To specify explicitly the height of a table, use the [`HEIGHT`](1782-height-attribute.md "The HEIGHT attribute forces an explicit height for a form element.") attribute:

```
TABLE t1 : table1, HEIGHT = 10 LINES, ... ;
```

The `LINES` unit means a number of lines in the `TABLE`, not
grid-lines.

The size of a `TABLE` (or `TREE`)
element has an impact on the size of a [modal
window](1567-configuring-windows-with-styles.md "Use the STYLE attribute to set a style for a window.") where the form of the table is displayed. However, regular windows will be sized from
the [window container](1563-containers-for-program-windows.md "Program windows are displayed in window containers by the front-end."), and the form
content will adapt to it.

## Showing table grid lines

By default, the grid lines of `TABLE` container are shown, when the table is
controlled by an `INPUT ARRAY`. The table grid lines are hidden when using a
`DISPLAY ARRAY`.

Define the [`showGrid`](1648-table-style-attributes.md) style attribute, to indicate if the table widget must show
horizontal and/or vertical grid lines:

```
<Style name="Table.custom_style">
  <StyleAttribute name="showGrid" value="vertical"/>
</Style>
```

![Table rendered with grid lines by using the showGrid style attribute.](../_images/table_showgrid.jpg)

*Table with grid lines*

## Alternate background color of odd/even rows

A `TABLE` container can automatically display odd/even rows with a different
background color. This is the default rendering, and it can be disabled with the [`alternateRows`](1648-table-style-attributes.md)
style attribute:

```
<Style name="Table.custom_style">
  <StyleAttribute name="alternateRows" value="no"/>
</Style>
```

Consider using `alternateRows` instead of `:odd/:even` pseudo
selectors.

![Table rendering with the alternateRows style attribute.](../_images/table_alternaterows.jpg)

*Table with alternate row background*

## Distinguish row on mouse hover

By default, when the mouse hovers over a row, it gets a specific background color to distinguish
the row from others. This rendering can be controller with the [`rowHover`](1648-table-style-attributes.md) style
attribute:

```
<Style name="Table.custom_style">
  <StyleAttribute name="rowHover" value="no"/>
</Style>
```

## Current row / current cell rendering

In a `TABLE` container, the highlighting of the current row (or current cell, when
focus granularity is at the cell level) can be controlled with style attributes.

The current row / current cell style attributes apply also to stretchable
`SCROLLGRID` containers.

For more details, see [Row and cell highlighting in TABLE](2328-row-and-cell-highlighting-in-table.md "TABLE containers can be configured with presentation styles, to control row and cell highlighting, using specific foreground and background colors.").

## Frozen columns

When the Table style attribute `tableType` is set to `frozenTable`,
you can define a number of fixed columns on the left and right, respectively with the
`leftFrozenColumns` and `rightFrozenColumns` style
attributes:

```
<Style name="Table.custom_style">
  <StyleAttribute name="tableType" value="frozenTable"/>
  <StyleAttribute name="leftFrozenColumns" value="2"/>
  <StyleAttribute name="rightFrozenColumns" value="1"/>
</Style>
```

Frozen columns can also be used in [Treeviews](2352-tree-views.md "Describes how to implement tree views.").

See the reference topics [Table style attributes](1648-table-style-attributes.md "Table presentation style attributes apply to a TABLE container."),
[Tree style attributes](1649-tree-style-attributes.md "Tree presentation style attributes apply to the TREE container.").

## Stretchable columns

The form items defining columns in a [`TABLE`](1703-table-item-type.md "Defines a list view widget.") or [`TREE`](1706-tree-item-type.md "Defines a tree view widget.") container can use a [`STRETCH=X`](1821-stretch-attribute.md "The STRETCH attribute defines if the form element can grow or has a fixed size.") attribute, to have the column stretch when the container width
changes.

To make all columns stretchable, define the `TABLE/TREE` form item with the [`STRETCHCOLUMNS`](1822-stretchcolumns-attribute.md "The STRETCHCOLUMNS attribute makes all TABLE/TREE columns stretchable.") attribute.
Some columns can then be made non-stretchable individually with `STRETCH=NONE`.

For stretchable columns, a minimum and maximum stretch width can be defined with the [`STRETCHMIN`](1824-stretchmin-attribute.md "The STRETCHMIN attribute defines the minimum stretching width.") and [`STRETCHMAX`](1823-stretchmax-attribute.md "The STRETCHMAX attribute defines the maximum stretching width for a TABLE/TREE column.") attributes. The value
for these attributes represents a number of item-tag cells. By default, the width of a column is
defined by the number of cells used in the column [item-tag](1679-item-tags.md "Item tags define the position and size in a grid-based container.").

`STRETCH`, `STRETCHCOLUMNS`, `STRETCHMIN` and
`STRETCHMAX` can be combined with a `@screen-size`
modifier, to enable these stretching attributes for specific screen sizes.

`TABLE` definition example using stretchable columns:

```
LAYOUT
...
TABLE ...
{
[c1  |c2    |c3      ]
}
END
...
ATTRIBUTES
EDIT c1 = customer.cust_num;
EDIT c2 = customer.cust_name, STRETCH=X,
     STRETCHMIN@SMALL=5, STRETCHMAX@LARGE=30 ;
EDIT c3 = customer.cust_zipcode;
...
```

In the above example, only the `c2 (cust_name)` column will stretch, when the
table width increases. For small screens, the minimum stretch width will be 5 grid cells, and for
large screens the maximum stretch width will be 30 cells.

See also [Horizontal stretching](1544-horizontal-stretching.md "Define stretchable form elements to achieve responsive layout.").

## Flip columns into rows

`TABLE` containers can define the [`FLIPPED`](1780-flipped-attribute.md "The FLIPPED attribute flips TABLE columns into rows.") attribute, to render the table in a way that all cells of a given
row get packed vertically. In this rendering mode, the table is thinner and can be displayed on
small screens.

To pack row cells vertically, some form item widgets may not keep the same size (especially the
height) as in classical table display. For example, the height of [`RADIOGROUP`](1699-radiogroup-item-type.md "Defines a mutual exclusive set of options field."), [`IMAGE`](1695-image-item-type.md "Defines an area that can display an image resource.") and [`TEXTEDIT`](1704-textedit-item-type.md "Defines a multi-line edit field.") form items will be
smaller. Consider using these item types with care, in a table where the `FLIPPED`
attribute applies.

The `FLIPPED` attribute can be combined with a
`@screen-size` modifier, to get this rendering for specific screen
sizes (typically with small
screens):

```
TABLE table1 ( FLIPPED@SMALL )
```

![Table with flipped rows arrangement.](../_images/table_flipped_layout.jpg)

*Table with flipped rows layout*

## Row aspect control

When using a `TABLE` with `FLIPPED` mode in action, use the [`rowAspect`](1648-table-style-attributes.md) style attribute to get a Material Design list rendering of the
rows:

```
<Style name="Table.list" >
    <StyleAttribute name="rowAspect" value="list" />
</Style>
```

When this style attribute is used, the flipped table looks like this:

![Table with flipped rows and list row aspect style.](../_images/table_flipped_as_card.jpg)

*Table with flipped rows and list row aspect style*

With the `rowAspect` set to `list`, if the table row has the
`IMAGECOLUMN` attribute
specified on the first column, the image is displayed as a thumbnail in front of all column values
rather than in the same table cell as the column on which it is defined.

![Table with flipped rows and list row aspect style + IMAGECOLUMN.](../_images/table_flipped_as_card_imagecolumn.jpg)

*Table with flipped rows and list row aspect style + IMAGECOLUMN*

## Column title rendering

When a `TABLE` uses the `FLIPPED` attribute, you can control the
visibility, position and alignment of the column titles respectively with the [`headerHidden`](1648-table-style-attributes.md), [`headerPosition`](1648-table-style-attributes.md) and [`headerAlignment`](1648-table-style-attributes.md) style
attributes:

```
<Style name="Table.list" >
     <StyleAttribute name="rowAspect" value="list" />
     <!-- default: StyleAttribute name="headerHidden" value="no" /-->
     <StyleAttribute name="headerPosition" value="top" />
     <StyleAttribute name="headerAlignment" value="center" />
</Style>
```

![Table with flipped rows, list row aspect style and column title alignment.](../_images/table_flipped_as_card_headers.jpg)

*Table with flipped rows, list row aspect style and column title alignment*

## Current row visibility after dialog execution

When the dialog controlling the table has finished, the current row may be deselected, depending
on the [`KEEP CURRENT ROW`](1965-display-array-instruction-configuration.md) dialog
attribute.

## Related links

**Related concepts**  

[TABLE container](1724-table-container.md "Defines a re-sizable table designed to display a list of records.")
