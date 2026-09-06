---
title: "The cell grid concept"
source: "fgl-topics/c_fgl_form_rendering_grid_concept.html"
breadcrumb: "User interface > Form definitions > Form rendering > Grid-based layout > The cell grid concept"
type: "concept"
---

# The cell grid concept

> Form elements align to each other by following the cells of a virtual grid.

In a .per form specification file, the `LAYOUT`
section defines a tree of layout containers, which hold form items such as labels and form fields.

The [`GRID`](1722-grid-container.md "Defines a layout area based on a grid of cells.") container
can be used to define a grid of cells that hold form items: in the layout tree, the
`GRID` container acts as a leaf node, which holds the visible widgets (fields,
buttons, and so on).

> **Note:**
>
> [`SCROLLGRID`](1700-scrollgrid-item-type.md "Defines a scrollable grid view widget.")
> and [`GROUP`](1693-group-item-type.md "Defines a layout area to group other layout elements together.") containers defined by
> layout tags inside a grid without the [`GRIDCHILDRENINPARENT`](1781-gridchildreninparent-attribute.md "The GRIDCHILDRENINPARENT attribute is used for a container to align its children to the parent container.") attribute, are similar to `GRID`
> containers in regards to the layout rules describe in this section.

The .per form specification file defines a form layout based on a character
grid, each character defines a cell of the
grid:

```
GRID
{
First Name [fname  ]
Last Name  [lname  ]
}
END
```

The .per file layout specification can be shown in a character grid.

![Character grid diagram](../_images/layout01.jpg)

*Character grid of a form layout*

With a fixed-font based front-end (such as a dumb terminal), the forms appear within a screen
where each cell is identified by x and y coordinates, as in the `SCREEN` section of
the form specification file. There is no particular layout issue, as all characters can be displayed
at the same (relative) position as in the source form file.

With the graphical front-end, text-based forms must be displayed in a graphical window using
fonts with a proportional size. In a proportional font, the field label "Key" has a different
graphical length than the label "Num", despite having the same number of characters.

In the compiled version of the form specification file, all form items get coordinates in a
virtual grid (defined by `posX` and `posY` attributes), and the number
of cells the item occupies in the grid (in the `gridWidth` and
`gridHeight` attributes):

![Grid positioning diagram](../_images/layout02.jpg)

*Grid positioning*

The "First Name" and "Last Name" texts are identified as whole labels, even if the words "First"
and "Name" (or "Last" and "Name") are not joined in the form definition, because the form compiler
considers a single blank as a word separator within labels.

## Related links

**Related concepts**  

[GRID container](1722-grid-container.md "Defines a layout area based on a grid of cells.")
