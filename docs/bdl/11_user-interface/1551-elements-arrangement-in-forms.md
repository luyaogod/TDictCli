---
title: "Elements arrangement in forms"
source: "fgl-topics/c_fgl_form_rendering_packed_grid.html"
breadcrumb: "User interface > Form definitions > Form rendering > Grid-based layout > Elements arrangement in forms"
type: "concept"
---

# Elements arrangement in forms

> When resizing a window, the content will either stretch with the window, be evenly distributed, or be packed in the top left position.

## Stretchable elements

If elements in the window can stretch, they will follow the window container and resize
accordingly. Some elements can stretch vertically, some can stretch horizontally, and some can
stretch in both directions.

The way resizable form items can grow is controlled by the [`STRETCH`](1821-stretch-attribute.md "The STRETCH attribute defines if the form element can grow or has a fixed size.") attribute.

The window content is packed horizontally, vertically or in both directions if none of the
elements can grow in that direction (horizontally, vertically or in both directions.)

The following form item types can stretch horizontally:

- `TABLE` / `TREE` containers (horizontally stretchable by
  design)
- `SCROLLGRID` containers with [tile-list
  view](2339-controlling-scrollgrid-rendering.md) rendering
- `IMAGE`, `TEXTEDIT`, and `WEBCOMPONENT` items with
  [`STRETCH=BOTH`](1821-stretch-attribute.md "The STRETCH attribute defines if the form element can grow or has a fixed size.") or [`STRETCH=X`](1821-stretch-attribute.md "The STRETCH attribute defines if the form element can grow or has a fixed size.")
- "Flat" items like `EDIT`, `BUTTON`, and `COMBOBOX`
  with [`STRETCH=X`](1821-stretch-attribute.md "The STRETCH attribute defines if the form element can grow or has a fixed size.") (see [Form item types](1683-form-item-types.md "The form item types defines the purpose of form elements."))

The following form item types can stretch vertically:

- `TABLE` / `TREE` containers when [`WANTFIXEDPAGESIZE`](1847-wantfixedpagesize-attribute.md "The WANTFIXEDPAGESIZE attribute controls the vertical resizing of a list element.") is not
  used
- `SCROLLGRID` containers when [`WANTFIXEDPAGESIZE=NO`](2339-controlling-scrollgrid-rendering.md "Scrollgrid rendering can be controlled by the use of presentation styles and scrollgrid attributes.")
- `IMAGE`, `TEXTEDIT`, and `WEBCOMPONENT` items with
  [`STRETCH=BOTH`](1821-stretch-attribute.md "The STRETCH attribute defines if the form element can grow or has a fixed size.") or [`STRETCH=Y`](1821-stretch-attribute.md "The STRETCH attribute defines if the form element can grow or has a fixed size.")

## Packing elements in HBOX/VBOX

By default, fixed-sized elements of a [`VBOX`](1718-vbox-container.md "Packs child layout elements vertically.") container are evenly distributed in the vertical direction and packed
to the left horizontally, and fixed-sized elements of an [`HBOX`](1717-hbox-container.md "Packs child layout elements horizontally.") container are
distributed evenly in the horizontal direction and centered vertically.

If you want to pack the fixed-sized elements to the top of a `VBOX` or to the left
of an `HBOX`, use the [`packed`](1640-hbox-style-attributes.md) style attribute for the box element.

Keep in mind `HBOX` and `VBOX` can change their orientation with
the [`ORIENTATION@media`](1805-orientation-attribute.md "The ORIENTATION attribute defines whether an element displays vertically or horizontally.") attribute. When the
`packed` style attribute is used for a box element that gets oriented in the opposite
direction, `packed` is applied as defined for the alignment direction selected for
the box.

For example, with the following form definition:

```
LAYOUT
VBOX vbox1 (STYLE="packme")
GRID grid1
{
Field11: [f11    ] Field12: [f12                ]
}
END
GRID grid2
{
Field21: [f21    ] Field22: [f22                ]
}
END
GRID grid3
{
Field31: [f31    ] Field32: [f32                ]
}
END
END
END
ATTRIBUTES
EDIT f11 = FORMONLY.field11;
EDIT f12 = FORMONLY.field12;
EDIT f21 = FORMONLY.field21;
EDIT f22 = FORMONLY.field22;
EDIT f31 = FORMONLY.field31;
EDIT f32 = FORMONLY.field32;
END
```

The form elements will be dispatched evenly in the vertical direction:

![Screenshot of VBOX without packed style attribute](../_images/layout81_gbc.jpg)

*VBOX without packed style attribute*

To pack the elements to the top of the vertical box, use the "packed" style attribute:

```
<Style name="VBox.packme">
    <StyleAttribute name="packed" value="yes" />
</Style>
```

![Screenshot of VBOX with packed style attribute](../_images/layout82_gbc.jpg)

*VBOX with packed style attribute*

## Placing group elements in parent GRID

In general, a [`GRID`](1722-grid-container.md "Defines a layout area based on a grid of cells.")
container can grow if any object inside the `GRID` can grow; however, there is an
exception to this rule. If there is a single [`GROUP`](1719-group-container.md "Defines a layout area to group other layout elements together.") container defined *without* the [`GRIDCHILDRENINPARENT`](1781-gridchildreninparent-attribute.md "The GRIDCHILDRENINPARENT attribute is used for a container to align its children to the parent container.")
attribute inside a `GRID` – and only the `GROUP` container in the
`GRID` – the grid can grow even if the objects inside the grid cannot grow.

This exception allows better rendering of a grouped grid.

![Screenshot of GROUP with GRIDCHILDRENINPARENT, note the white background of the group element](../_images/layout22_gbc.jpg)

*GROUP with GRIDCHILDRENINPARENT*

![Screenshot of GROUP without GRIDCHILDRENINPARENT, note the white background of the group element](../_images/layout21_gbc.jpg)

*GROUP without GRIDCHILDRENINPARENT*

Note the white background of the group element.

Here is the code supporting this explanation; comment out the
`GRIDCHILDRENINPARENT` attribute to view the difference.

```
LAYOUT
GRID
{
<GROUP g1   >
[f1        ]
[f2   ]
<           >
}
END
END

ATTRIBUTES
GROUP g1: group1, TEXT="Group 1"
        , GRIDCHILDRENINPARENT -- Comment to see the difference
        ;
EDIT f1 = FORMONLY.field1;
EDIT f2 = FORMONLY.field2;
END
```
