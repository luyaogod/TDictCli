---
title: "Controlling scrollgrid rendering"
source: "fgl-topics/c_fgl_ui_scrollgrids_rendering.html"
breadcrumb: "User interface > User interface programming > Scrollgrid views > Controlling scrollgrid rendering"
type: "concept"
---

# Controlling scrollgrid rendering

> Scrollgrid rendering can be controlled by the use of presentation styles and scrollgrid attributes.

## Scrollgrid resize control

By default scrollgrids are not resizable and the number of fixed rows is defined by the layout
element. The [`WANTFIXEDPAGESIZE`](1847-wantfixedpagesize-attribute.md "The WANTFIXEDPAGESIZE attribute controls the vertical resizing of a list element.") form file attribute controls the vertical resizing of the
list elements. Set this attribute to `NO`, in order to get a resizable
scrollgrid:

```
LAYOUT
SCROLLGRID ( WANTFIXEDPAGESIZE=NO )
{
...
```

![Resizable scrollgrid.](../_images/scrollgrid_standard.jpg)

*Resizable scrollgrid*

## Minimum number of scrollgrid lines

With a resizable scrollgrid, you can define the initial number of rows with the [`INITIALIPAGESIZE`](1792-initialpagesize-attribute.md "The INITIALPAGESIZE attribute defines the initial page size of a list element.") form file
attribute:

```
LAYOUT
SCROLLGRID ( WANTFIXEDPAGESIZE=NO, INITIALPAGESIZE=4 )
{
...
```

The `INITIALPAGESIZE` attribute is mainly useful in [modal windows](1567-configuring-windows-with-styles.md "Use the STYLE attribute to set a style for a window."), to define the initial window size.
However, regular windows will be sized from the [window container](1563-containers-for-program-windows.md "Program windows are displayed in window containers by the front-end."), and the form content will adapt to it.

## Paged scrollgrids (tile list)

A scrollgrid can be rendered as a tile list to fit records horizontally and vertically in a page
to respond to the container size when stretched or shrunk.

To enable a paged scrollgrid, in your .4st file, define the
`customWidget` style attribute to the value `pagedScrollGrid`:

```
<Style name="ScrollGrid.paged" >
    <StyleAttribute name="customWidget" value="pagedScrollGrid" />
</Style>
```

In the scrollgrid layout definition, use the style name as defined in your styles
file:

```
LAYOUT
SCROLLGRID ( WANTFIXEDPAGESIZE=NO, STYLE="paged" )
{
...
```

![Paged scrollgrid.](../_images/scrollgrid_paged_2.jpg)

*Paged scrollgrid*

For more details see the [customWidget](1646-scrollgrid-style-attributes.md) presention style attribute reference.

## Controlling element alignment inside a scrollgrid

To control the alignment of the elements inside the scrollgrid, use the
`itemsAlignment` presentation style:

```
<Style name="ScrollGrid.centered" >
    <StyleAttribute name="itemsAlignment" value="center" />
</Style>
```

In the scrollgrid layout definition, use the style name as defined in your styles
file:

```
LAYOUT
SCROLLGRID ( WANTFIXEDPAGESIZE=NO, STYLE="centered" )
{
...
```

This attribute applies to default scrollgrid rendering and paged scrollgrid rendering. However,
the possible values of `itemsAlignment` depend on the type of scrollgrid
rendering.

For more details see the [itemsAlignment](1646-scrollgrid-style-attributes.md) presention style attribute reference.

## Current row / current cell rendering

In a `SCROLLGRID` container, the highlighting of the current row (or current cell,
when focus granularity is at the cell level) can be controlled with style attributes:

```
  <Style name="ScrollGrid">
     <StyleAttribute name="highlightCurrentRow" value="no" />
     <StyleAttribute name="highlightColor" value="#AEFFAE" />
     <StyleAttribute name="highlightTextColor" value="black" />
  </Style>
```

When the above style attributes are used, the resizable scrollgrid looks like this:

![Scrollgrid row highlighting control.](../_images/scrollgrid_highlight_row_1.jpg)

*Scrollgrid row highlighting control*

For more details, see [Row and cell highlighting in SCROLLGRID](2340-row-and-cell-highlighting-in-scrollgrid.md "SCROLLGRID containers can be configured with presentation styles, to control row and cell highlighting, using specific foreground and background colors.").

## Current row visibility after dialog execution

When the dialog controlling the scrollgrid has finished, the current row may be deselected,
depending on the [`KEEP CURRENT ROW`](1965-display-array-instruction-configuration.md)
dialog attribute.

## Row aspect control

When using a `SCROLLGRID` with `WANTFIXEDPAGESIZE=NO`, use the
[`rowAspect`](1646-scrollgrid-style-attributes.md) style attribute to get a Material Design list rendering of the
rows:

```
<Style name="ScrollGrid.list" >
    <StyleAttribute name="rowAspect" value="list" />
</Style>
```

When this style attribute is used, the resizable scrollgrid looks like this:

![Resizable scrollgrid with card style.](../_images/scrollgrid_as_card.jpg)

*Resizable scrollgrid with card style*

## Related links

**Related concepts**  

[SCROLLGRID container](1723-scrollgrid-container.md "Defines a scrollable grid view widget.")
