---
title: "Defining hbox tags in grids"
source: "fgl-topics/c_fgl_form_rendering_hbox_tags_def.html"
breadcrumb: "User interface > Form definitions > Form rendering > Grid-based layout > Using hbox tags to align form items > Defining hbox tags in grids"
type: "concept"
description: "An hbox tag is defined by using a : colon in an item tag delimited by square brackets: [f1 ][f2 ] -- No hbox tag [f1 |f2 ] -- No hbox tag [f1 :f2 ] -- f1 and f2 in hbox tag Next example creates an ..."
---

# Defining hbox tags in grids

An hbox tag is defined by using a `:` colon in an item tag delimited by square
brackets:

```
[f1  ][f2   ] -- No hbox tag
[f1  |f2   ]  -- No hbox tag
[f1  :f2   ]  -- f1 and f2 in hbox tag
```

Next example creates an hbox containing the fields "`a`", "`b`" and
"`c`". In the grid, the widgets for these fields won't be aligned with the widgets
created for the "`d`", "`e`" and "`f`" fields, because
it is the parent hbox that will align with the grid columns:

```
LAYOUT
GRID
{
[a:b:c   ]
[d|e|f   ]
}
END
END
ATTRIBUTES
EDIT a = FORMONLY.a;
EDIT b = FORMONLY.b;
EDIT c = FORMONLY.c;
EDIT d = FORMONLY.d;
EDIT e = FORMONLY.e;
EDIT f = FORMONLY.f;
END
```

![Grid view of using a HBox tag](../_images/layout14.jpg)

*Using an hbox tag*

![Widget size computations diagram](../_images/layout15.jpg)

*Widget size computations*

The three first widgets are then rendered independently, inside the hbox tag:

![HBox tag rendering 1 screenshot](../_images/layout61_gbc.jpg)

*HBox rendering 1*

Hbox tags are useful when the form contains large widgets in a small number of cells, that must
not be dependent to other widgets regarding grid column alignment. For example, when using a
`CHECKBOX` for field "d":

```
LAYOUT
GRID
{
[a:b:c   ]
[d|e|f   ]
}
END
END
ATTRIBUTES
EDIT a = FORMONLY.a;
EDIT b = FORMONLY.b;
EDIT c = FORMONLY.c;
CHECKBOX d = FORMONLY.d, TEXT="Checkbox";
EDIT e = FORMONLY.e;
EDIT f = FORMONLY.f;
END
```

The alignment in the grid would result to:

![Using an HBox tag diagram](../_images/layout16.jpg)

![HBox tag rendering 2 screenshot](../_images/layout62_gbc.jpg)

*Widget size computations*

The next example is now using a `CHECKBOX` for field "`d`", and
hbox tags to group fields "a" and "b" as well as "d" and "e", and uses regular item tags to align
the widgets of field "c" and "f" without an hbox tag:

```
LAYOUT
GRID
{
[a:b   ][c     ]
[d:e   ][f     ]
}
END
END
ATTRIBUTES
EDIT a = FORMONLY.a;
EDIT b = FORMONLY.b;
EDIT c = FORMONLY.c;
CHECKBOX d = FORMONLY.d, TEXT="Checkbox";
EDIT e = FORMONLY.e;
EDIT f = FORMONLY.f;
END
```

Rendering result fields "`c`" and "`f`" are now aligned:

![HBox tag rendering 3 screenshot](../_images/layout63_gbc.jpg)

*HBox rendering 3*

## Related links

**Related concepts**  

[Automatic HBox/VBox with splitter](1554-automatic-hbox-vbox-with-splitter.md "Horizontal and vertical boxes with splitter are created automatically when stretchable elements are set side by side.")
