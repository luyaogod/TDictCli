---
title: "WIDTH attribute"
source: "fgl-topics/c_fgl_FSFAttributes_WIDTH.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > WIDTH attribute"
type: "concept"
---

# WIDTH attribute

> The WIDTH attribute forces an explicit width of a form element.

## Syntax

```
WIDTH = value [CHARACTERS|COLUMNS|POINTS|PIXELS]
```

1. value is an integer that defines the width of the element, according to the
   specified unit.

## Usage

By default, the width of an element is defined by the size of the [form item tag](1679-item-tags.md "Item tags define the position and size in a grid-based container."). Use the `WIDTH`
attribute, to define a specific width for a form item.

The `WIDTH` attribute is only available for some types of resizable form items,
such as [`TABLE`](1703-table-item-type.md "Defines a list view widget."), [`TREE`](1706-tree-item-type.md "Defines a tree view widget."), [`IMAGE`](1695-image-item-type.md "Defines an area that can display an image resource."), [`WEBCOMPONENT`](1708-webcomponent-item-type.md "Defines a specialized form item that holds an external component.").

> **Note:**
>
> As a general rule, consider not specifying a unit, to default to relative
> characters/lines/columns, instead of specifying exact pixels or points. This is especially important
> for mobile devices, where the screen resolution can significantly vary depending on the smartphone
> or tablet model.

If you don't specify a size unit for the `WIDTH` attribute, it defaults to
`CHARACTERS`, which defines a width based on the characters size in the current
font.

The `CHARACTERS` unit corresponds to the CSS `"em"` unit and
specifies a width based on the font-size used by the form.

The `PIXELS` unit defines a number of CSS `"px"` pixels. In this
context, the size of a pixel is defined by web standards and is independent of the screen
resolution.

The `POINTS` unit specifies an absolute CSS `"pt"` size defined as
follows: 72 points = 1 inch.

The `COLUMNS` unit defines a number of grid columns for `IMAGE`
fields. The size of a grid column depends on the font-size used by the elements in the form. For
`TABLE/TREE` containers, the `COLUMNS` units is interpreted as a
number of first field columns.

For sizable items like `IMAGE` and `WEBCOMPONENT`, the default
width is defined by the number of horizontal characters used in the form item tag. Overwrite this
default by specifying the `WIDTH` attribute:

```
IMAGE img1: image1, WIDTH = 20, HEIGHT = 12;
```

For `TABLE`/`TREE` containers, the default width is defined by the
fields/columns used in the table definition. Overwrite the default by specifying the `WIDTH =
x COLUMNS` attribute. This will give a small initial width for tables
with a large number of columns:

```
TABLE t1: table1, WIDTH = 5 COLUMNS;
```

The size of an element like `TABLE`,
`WEBCOMPONENT` or `IMAGE` has an impact on the size of a [modal window](1567-configuring-windows-with-styles.md "Use the STYLE attribute to set a style for a window.") where the form of the table is
displayed. However, regular windows will be sized from the [window container](1563-containers-for-program-windows.md "Program windows are displayed in window containers by the front-end."), and the form content will
adapt to it.

## Related links

**Related concepts**  

[Controlling the image layout](1585-controlling-the-image-layout.md "Explains how image form items can be sized in different front-end layout systems.")

[Controlling table rendering](2319-controlling-table-rendering.md "Table rendering can be controlled by the use of presentation styles and table attributes.")

[Web component grid layout](2382-web-component-grid-layout.md "Web component grid layout")

[Item tags](1679-item-tags.md "Item tags define the position and size in a grid-based container.")

[Form rendering basics](1539-form-rendering-basics.md "Get the essentials about form rendering.")

[HEIGHT attribute](1782-height-attribute.md "The HEIGHT attribute forces an explicit height for a form element.")
