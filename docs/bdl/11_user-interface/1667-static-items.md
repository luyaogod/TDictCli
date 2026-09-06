---
title: "Static items"
source: "fgl-topics/c_fgl_FormSpecFiles_Static_Items.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file concepts > Form items > Static items"
type: "concept"
---

# Static items

> A static item defines a simple form item as a final grid element that does not change.

A static item is a form element, such as text (typically, a field label), that is defined
directly in a [`GRID`](1722-grid-container.md "Defines a layout area based on a grid of cells.") (or
[`SCROLLGRID`](1723-scrollgrid-container.md "Defines a scrollable grid view widget."))
container in the form `LAYOUT` section.

Static items are identified by the fglform compiler and converted to an
AUI tree node element in the resulting .42f file.

## Simple texts

It is possible to define simple texts and field labels in the form layout:

```
LAYOUT
GRID
{
A simple text
}
END
END
```

To simplify internationalization, consider using named [static labels](1696-label-item-type.md "Defines a simple text area to display a read-only value.") instead of hard-coded text in the form
layout.

## Horizontal lines

You define a horizontal line with a sequence of hyphen-minus (`-`) characters
in a grid:

```
LAYOUT
GRID
{
This is a horizontal line: ------------
}
END
END
```

Horizontal lines are mainly provided for TUI mode applications. While horizontal lines will be
represented by some GUI front-ends, it is not a typical practice in common graphical
applications.

## Related links

**Related concepts**  

[LAYOUT section](1715-layout-section.md "The LAYOUT section defines the graphical alignment of the form by using a tree of layout containers.")
