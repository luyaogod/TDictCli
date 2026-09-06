---
title: "STRETCH attribute"
source: "fgl-topics/c_fgl_FSFAttributes_STRETCH.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > STRETCH attribute"
type: "concept"
---

# STRETCH attribute

> The STRETCH attribute defines if the form element can grow or has a fixed size.

## Syntax 1: Vertical and horizontal stretch

```
STRETCH[@screen-size] = { NONE | X | Y | BOTH }
```

where screen-size can
be:

```
{ SMALL | MEDIUM | LARGE }
```

1. screen-size is a screen size selector that indicates when
   the stretch attribute must apply, depending on the size of the screen. Several
   `STRETCH@screen-size` attributes can be used for the same
   element.
2. Values `NONE`, `X`, `Y` and `BOTH`
   apply to form items that can stretch horizontally and vertically such as `TEXTEDIT`,
   `IMAGE`, `WEBCOMPONENT`.

## Syntax 2: Horizontally-only stretch

```
STRETCH[@screen-size] = { NONE | X }
```

where screen-size can
be:

```
{ SMALL | MEDIUM | LARGE }
```

1. screen-size is a screen size selector that indicates when
   the stretch attribute must apply, depending on the size of the screen. Several
   `STRETCH@screen-size` attributes can be used for the same
   element.
2. Values `NONE`, `X` apply to form items that can only stretch
   horizontally such as `BUTTON`, `EDIT`, `COMBOBOX`, as
   well as to the `LAYOUT` section.

## Usage

The `STRETCH` attribute specifies if the size of the form element is fixed, or
depends on the width or height of the window. The remaining space in the window is distributed among
the form elements that can stretch.

`STRETCH` options description:

- `STRETCH=NONE`: Element cannot resize in any direction.
- `STRETCH=Y`: Element can resize vertically only.
- `STRETCH=X`: Element can resize horizontally only.
- `STRETCH=BOTH`: Element can resize vertically and horizontally.

In a `GRID` container where [grid-based aligment rules](1553-form-item-dependencies-in-grids.md "Form items interact with each other in terms of width, depending on the front-end widget size.") take precedence over size and proportions, the
`STRETCH` attribute set on a given item applies to the entire grid-column (for
`STRETCH=X/BOTH`) or to the entire grid-row (for `STRETCH=Y/BOTH`),
meaning that all elements that belong to the same grid column or row will stretch accordingly.

`STRETCH=X` is typically used on `GRID` elements to achieve [responsive layout](1544-horizontal-stretching.md "Define stretchable form elements to achieve responsive layout.").

Fields that can stretch horizontally should also use the [`SCROLL`](1815-scroll-attribute.md "The SCROLL attribute can be used to enable horizontal scrolling in a character field.") attribute.

For elements inside `GRID`/`SCROLLGRID` containers,
`STRETCH=X` attribute can be combined with the [`STRETCHMIN`](1824-stretchmin-attribute.md "The STRETCHMIN attribute defines the minimum stretching width.") attribute, to define a
minimum width for the stretchable element.

For `TABLE/TREE` columns, the `STRETCH=X` attribute can be combined
with [`STRETCHMIN`](1824-stretchmin-attribute.md "The STRETCHMIN attribute defines the minimum stretching width.") and [`STRETCHMAX`](1823-stretchmax-attribute.md "The STRETCHMAX attribute defines the maximum stretching width for a TABLE/TREE column.") attributes, to define
a minimum and maximum width for the stretchable element. The [`STRETCHCOLUMNS`](1822-stretchcolumns-attribute.md "The STRETCHCOLUMNS attribute makes all TABLE/TREE columns stretchable.") attribute can
be used at the `TABLE/TREE` item definition to make all columns stretchable.

`STRETCH` can be set on the [`LAYOUT`](1715-layout-section.md "The LAYOUT section defines the graphical alignment of the form by using a tree of layout containers.") element, to specify a default stretch behaviour for all children
elements of the form. This default behaviour can be overridden by a local definition of
`STRETCH` on a form item. For the `LAYOUT` node, the default is
`STRETCH=X` for mobile devices and `STRETCH=NONE` for desktop
front-ends.

By default, resizable form items like [`IMAGE`](1695-image-item-type.md "Defines an area that can display an image resource."), [`TEXTEDIT`](1704-textedit-item-type.md "Defines a multi-line edit field.") have a fixed width and height. Use the `STRETCH`
attribute to allow the widget to resize vertically, horizontally, or in both directions.

By default, [`WEBCOMPONENT`](1708-webcomponent-item-type.md "Defines a specialized form item that holds an external component.") form items stretch in both directions
(`STRETCH=BOTH` is the default).

To create a responsive layout form, the `STRETCH` attribute can be combined with
the `@screen-size` selector, to get different stretch behaviors
depending on the screen size.

## Example

```
IMAGE i01 = FORMONLY.picture, STRETCH=BOTH;
EDIT f01 = customer.cust_name, STRETCH@LARGE=X, STRETCH@MEDIUM=NONE;
```

## Related links

**Related concepts**  

[Controlling the image layout](1585-controlling-the-image-layout.md "Explains how image form items can be sized in different front-end layout systems.")

[Controlling table rendering](2319-controlling-table-rendering.md "Table rendering can be controlled by the use of presentation styles and table attributes.")

[Responsive Layout](1541-responsive-layout.md "Forms can be designed to adapt to the front-end screen possibilities.")
