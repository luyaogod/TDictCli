---
title: "Form item dependencies in grids"
source: "fgl-topics/c_fgl_form_rendering_grid_dependents.html"
breadcrumb: "User interface > Form definitions > Form rendering > Grid-based layout > Form item dependencies in grids"
type: "concept"
---

# Form item dependencies in grids

> Form items interact with each other in terms of width, depending on the front-end widget size.

This example illustrates how form items are dependent on each other inside the
grid.

```
GRID
{
  [a      ]
  [b      ]
}
END
```

This .per implies that form items `a` and `b`
start at the same position and have the same size, whatever `a` and
`b` are.

This rule leads to very different results, especially when a large widget is assigned into a
small number of cells.

Example:

```
LAYOUT
GRID
{
[a|b   ][f     ]
[c|d]   [e     ]
}
END
END
ATTRIBUTES
CHECKBOX a = FORMONLY.a, TEXT="A Checkbox";
EDIT b = FORMONLY.b;
EDIT c = FORMONLY.c;
CHECKBOX d = FORMONLY.d, TEXT="Another Checkbox";
EDIT e = FORMONLY.e;
EDIT f = FORMONLY.f;
END
```

The grid is computed with regard to the character cells in the form definition:

![Grid layout diagram](../_images/layout09.jpg)

*Grid layout*

![Grid layout with checkboxes diagram](../_images/layout10.jpg)

*Grid layout with checkboxes*

Then the minimum size of each widget and the layout is computed.

Cells (0,1) and (1,3) contain a checkbox; these checkboxes will enlarge columns 1 and 3.

![Enlarged columns diagram](../_images/layout11.jpg)

*Enlarged columns*

Because the `EDIT` field "`c`" is defined to have the same width as
checkbox "`a`", it will be much larger as expected.

![Resulting form screenshot 1](../_images/layout12_gbc.jpg)

*Resulting form 1*

To avoid this visual result, you must assign a realistic number of grid cells for each form item.
Consider however to use enough form cells to hold the values that can be displayed and entered in
the fields:

```
GRID
{
[a       |b   ][f     ]
[c|d          ][e     ]
}
END
```

Even if the grid area is wider in the source form file, the real graphical result will be
smaller.

![Resulting form screenshot 2](../_images/layout13_gbc.jpg)

*Resulting form 2*

Furthermore, you can add a [`STRETCH=X`](1821-stretch-attribute.md "The STRETCH attribute defines if the form element can grow or has a fixed size.") attribute to the field "`f`", and make the last
grid column stretch to the size of the parent container. Note that the underlying field
"`e`" will stretch as well, because it belongs to the same grid column as
"`f`":

```
EDIT f = formonly.f, STRETCH=X;
```

![Resulting form screenshot 2](../_images/layout41_gbc.jpg)

*Resulting form 3*

If the minimum width must be different from the number of grid cells used by the form item tag of
items using the `STRETCH=X` attribute, add the [`STRETCHMIN=nn`](1824-stretchmin-attribute.md "The STRETCHMIN attribute defines the minimum stretching width.")
attribute, to force the minimum width:

```
GRID
{
Customer ID: [cid     ]  <-- width is 8 cells
...
}
END
...
EDIT cid: customer.cust_id, STRETCH=X, STRETCHMIN=3;
...
```

## Related links

**Related concepts**  

[Item tags](1679-item-tags.md "Item tags define the position and size in a grid-based container.")

[Hbox tags](1680-hbox-tags.md "Hbox tags group several item tags within the same horizontal layout box, inside a grid-based container (GRID).")
