---
title: "GRIDCHILDRENINPARENT attribute"
source: "fgl-topics/c_fgl_FSFAttributes_GRIDCHILDRENINPARENT.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > GRIDCHILDRENINPARENT attribute"
type: "concept"
---

# GRIDCHILDRENINPARENT attribute

> The GRIDCHILDRENINPARENT attribute is used for a container to align its children to the parent container.

## Syntax

```
GRIDCHILDRENINPARENT
```

## Usage

By default, child elements of a grid container are aligned locally inside the container layout
cells. With the `GRIDCHILDRENINPARENT` attribute, you can force children to be
aligned in a vertical or horizontal direction, based on the layout cells in the parent container of
the container to which you assign this attribute.

The `GRIDCHILDRENINPARENT` attribute applies only to [`GROUP`](1738-group-item-definition.md "Defines attributes for a group-box layout tag.") and [`SCROLLGRID`](1743-scrollgrid-item-definition.md "Defines attributes for a scrollgrid layout tag.") containers
used inside a parent [`GRID`](1722-grid-container.md "Defines a layout area based on a grid of cells.") container.

When the group or scrollgrid containers are placed vertically over each other, the alignment
applies on parent grid columns, and when the containers are placed side by side horizontally, the
alignment applies on parent grid rows.

## Example

With this form definition, the elements in the four group boxes will align vertically
and horizontally to the parent grid cells:

```
LAYOUT
GRID
{
<G ga       ><G gb       >
 Some text
[a          ]  b[b     ]
<           ><           >
<G gc       ><G gd       >
[c          ]  d[d     ]
<           ><           >
}
END
END
ATTRIBUTES
GROUP ga: GRIDCHILDRENINPARENT;
GROUP gb: GRIDCHILDRENINPARENT;
GROUP gc: GRIDCHILDRENINPARENT;
GROUP gd: GRIDCHILDRENINPARENT;
EDIT a = FORMONLY.f_a;
EDIT b = FORMONLY.f_b;
EDIT c = FORMONLY.f_c;
EDIT d = FORMONLY.f_d;
END
```

## Related links

**Related concepts**  

[Form rendering](1538-form-rendering.md "The section explains the layout rules to render forms on graphical front-ends.")

[Layout tags](1678-layout-tags.md "Layout tags define layout areas for containers inside the frame of a grid-based container.")
