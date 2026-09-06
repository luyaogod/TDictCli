---
title: "Layout structure for responsive"
source: "fgl-topics/c_fgl_form_responsive_struct.html"
breadcrumb: "User interface > Form definitions > Form rendering > Responsive Layout > Layout structure for responsive"
type: "concept"
---

# Layout structure for responsive

> Use the LAYOUT structure to define forms that adapt to screen sizes.

## Understanding the layout structure for responsive

To achieve best responsive layout, a form should be defined with a combination of small,
individual, stretchable containers, that are organized in a tree of `HBOX/VBOX`
elements:

```
VBOX
|-- HBOX
|   |-- TABLE 
|   +-- GRID
+-- GRID
```

In `GRID` containers, define fields as stretchable, to allow fields to resize
depending on the screen size. Containers such as `TABLE` are by essence
stretchable.

`HBOX/VBOX` and leaf containers can be defined with
`HIDDEN@screen-size` and/or
`ORIENTATION@screen-size` attributes, to describe how these should
behave depending on the size of the screen.

## `GRID` as stretchable leaf container

In a responsive layout form, define `GRID` containers with basic form items such
as labels and fields, without using [container
layout tags](1678-layout-tags.md "Layout tags define layout areas for containers inside the frame of a grid-based container.") such as `<GROUP g1 >`.

Inside the `GRID` containers, elements can be organized by position. Right-most
elements can be made [horizontally
stretchable](1544-horizontal-stretching.md "Define stretchable form elements to achieve responsive layout."), to occupy all the screen width.

![Diagram of a responsive form using grids](../_images/rl_concept_grid_1.jpg)

*Responsive form with labels and fields in grid*

## Screen-size dependant attributes

Responsive layout introduces a new syntax to indicate for what screen or window width an
attribute must take effect.

For example, the `HIDDEN` attribute can get the
`@screen-size`
specifier:

```
HIDDEN@screen-size
```

The `@screen-size` specifier can be one of
`@SMALL`, `@MEDIUM`, `@LARGE`.

The interpretation of these abtract size qualifiers is done dynamically, depending on the
device type and display area size/width.

The exact screen size/width corresponding to `SMALL/MEDIUM/LARGE` can be
configured in the front-end, but it is not recommended to change the defaults. The predefined sizes
correspond typically to a given device type: `SMALL` and `MEDIUM` for
smartphones and tablets, `LARGE` for desktop computer screens. This abstract size can
also reflect the screen orientation on any device type.

For the complete syntax, see [HIDDEN attribute](1783-hidden-attribute.md "The HIDDEN attribute indicates if an element must be visible to the user.").

## Responsive form layout definition example

In the next example:

```
LAYOUT
VBOX vb1
  HBOX hb1 ( ORIENTATION@SMALL = VERTICAL )
    TABLE tb1
       ...
    END
    GRID gd1 ( HIDDEN@SMALL )
       ...
    END
  END
  GRID gd2 ( HIDDEN@SMALL, HIDDEN@MEDIUM )
     ...
  END
END
END
```

1. `vb1` organizes `hb1` and `gd2` vertically, no
   matter the screen size or orientation.
2. `hb1` organizes `tb1` and `gd1` horizontally by
   default, but vertically on small screens.
3. `gd1` will be hidden on small screens.
4. `gd2` will be hidden on small and medium screens.

## Related links

**Related concepts**  

[LAYOUT section](1715-layout-section.md "The LAYOUT section defines the graphical alignment of the form by using a tree of layout containers.")
