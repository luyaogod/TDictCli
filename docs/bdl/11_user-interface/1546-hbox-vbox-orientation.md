---
title: "HBOX/VBOX orientation"
source: "fgl-topics/c_fgl_form_responsive_orientation.html"
breadcrumb: "User interface > Form definitions > Form rendering > Responsive Layout > HBOX/VBOX orientation"
type: "concept"
---

# HBOX/VBOX orientation

> HBOX/VBOX containers can be defined to adapt to the screen size.

## HBOX/VBOX horizontal fit

`HBOX` and `VBOX` can be defined with the `ORIENTATION` attribute, to
force the sub-elements or sub-containers to align vertically or horizontally, depending on the display area width.

`HBOX` and `VBOX` elements can be oriented automatically for a given screen size,
depending on an abstract size specification, with the [`ORIENTATION@screen-size`](1543-layout-structure-for-responsive.md) attribute syntax.

An `HBOX` is by definition oriented horizontally, while a `VBOX` is by default oriented
vertically: It is not required to specify an `ORIENTATION@screen-size` attribute in
the natural direction of the box.

For the complete syntax, see [ORIENTATION attribute](1805-orientation-attribute.md "The ORIENTATION attribute defines whether an element displays vertically or horizontally.").

## Form example using the `ORIENTATION` attribute

The following example defines orientation attributes for a `VBOX` to get a responsive layout. With
this definition, the `GRID` will appear on top of the `TABLE` in small screens (using
vertical orientation). On medium and large screens, the `GRID` will appear on the left of the
`TABLE` (using horizontal orientation).

Notes:

1. The `VBOX` element is by essence a vertical box, and therefore would not require the
   `ORIENTATION@SMALL=VERTICAL` attribute. It is used here for better understanding.
2. The `VBOX` element also defines the `SPLITTER` attribute, to get a splitter between
   the grid and the table.
3. The `field11` has a `STRETCH=X` attribute, so that the grid elements strech to the
   available width, depending on the vertical or horizontal orientation of the box.

Form definition
file:

```
LAYOUT ( STYLE="main2" )
VBOX (
 ORIENTATION@SMALL  = VERTICAL,   -- (optional)
 ORIENTATION@MEDIUM = HORIZONTAL,
 ORIENTATION@LARGE  = HORIZONTAL,
 SPLITTER
)
GRID grid1
{
Field11: [f11         ]
Field12: [f12     ]
Field13: [f13         ]
}
END
TABLE table1
{
 Col21  Col22     Col23
[c21   |c22      |c23       ]
[c21   |c22      |c23       ]
[c21   |c22      |c23       ]
}
END
END
END

ATTRIBUTES

EDIT f11 = FORMONLY.field11, STRETCH=X;
COMBOBOX f12 = FORMONLY.field12, ITEMS=("aaa","bbb");
DATEEDIT f13 = FORMONLY.field13;

EDIT c21 = FORMONLY.col21;
EDIT c22 = FORMONLY.col22;
DATEEDIT c23 = FORMONLY.col23;

END

INSTRUCTIONS
SCREEN RECORD sr1(FORMONLY.col21 THRU FORMONLY.col23);
END
```

Visual result with a small screen size (orientation is vertical):

![Screenshot of form with vertical box in small window](../_images/rl_box_orient_gbc_1.jpg)

*Form with vertical box in small window*

Visual result with a medium window (horizontal orientation is used):

![Screenshot of form with horizontal box in medium window](../_images/rl_box_orient_gbc_2.jpg)

*Form with horizontal box in medium window*
