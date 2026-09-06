---
title: "Automatic HBox/VBox with splitter"
source: "fgl-topics/c_fgl_form_rendering_auto_v_h_box.html"
breadcrumb: "User interface > Form definitions > Form rendering > Grid-based layout > Automatic HBox/VBox with splitter"
type: "concept"
---

# Automatic HBox/VBox with splitter

> Horizontal and vertical boxes with splitter are created automatically when stretchable elements are set side by side.

## When are HBox/VBox automatically created?

When using layout tags in a [`GRID`](1722-grid-container.md "Defines a layout area based on a grid of cells.") container, the fglform compiler will automatically
add hbox or vbox containers with splitters in the following conditions:

- An hbox is created when two or more stretchable elements are stacked side by side and touch each
  other (no space between).
- A vbox is created when two or more stretchable elements are stacked vertically and touch each
  other (no space between).

No hbox or vbox will be created if the elements are in a [`SCROLLGRID`](1723-scrollgrid-container.md "Defines a scrollable grid view widget.")
container.

## Example with tables and VBox+Splitter

This example defines two table item tags stacked vertically in a `GRID` container,
generating a vbox with splitter. The ending tags for the tables are omitted.

```
LAYOUT
GRID
{
<TABLE t1         >
[c11 |c12   |c13  ]
[c11 |c12   |c13  ]
[c11 |c12   |c13  ]
<                 >
<TABLE t2         >
[c21 |c22         ]
[c21 |c22         ]
[c21 |c22         ]
<                 >
}
END
END

ATTRIBUTES
TABLE t1: table1;
EDIT c11 = FORMONLY.col11, TITLE="C11";
EDIT c12 = FORMONLY.col12, TITLE="C12";
EDIT c13 = FORMONLY.col13, TITLE="C13";
TABLE t2: table2;
EDIT c21 = FORMONLY.col21, TITLE="C21";
EDIT c22 = FORMONLY.col22, TITLE="C22";
END

INSTRUCTIONS
SCREEN RECORD sr1(col11,col12,col13);
SCREEN RECORD sr2(col21,col22);
END
```

![Two tables with VBox/Splitter screenshot](../_images/layout31_gbc.jpg)

*Two table item tags producing an automatic vbox/splitter in between*

## Example with textedits and HBox+Splitter

This example defines a `GRID` with two stretchable [`TEXTEDIT`](1704-textedit-item-type.md "Defines a multi-line edit field.") fields placed side by
side, which would generate an automatic hbox with splitter.

- To make both widgets touch, you need to use a pipe delimiter in between the two item tags.
- The `TEXTEDIT` fields must be horizontally stretchable and therefore need a
  `STETCH=X` or `STRETCH=BOTH` attribute.

```
LAYOUT
GRID
{
[te1           |te2          ]
[              |             ]
[              |             ]
[              |             ]
[              |             ]
}
END
END

ATTRIBUTES
TEXTEDIT te1 = FORMONLY.field1, STRETCH=X;
TEXTEDIT te2 = FORMONLY.field2, STRETCH=X;
END
```

![Two textedits with HBox/Splitter screenshot](../_images/layout32_gbc.jpg)

*Two textedit item tags producing an automatic hbox/splitter in between*

## Related links

**Related concepts**  

[Layout tags](1678-layout-tags.md "Layout tags define layout areas for containers inside the frame of a grid-based container.")

[STRETCH attribute](1821-stretch-attribute.md "The STRETCH attribute defines if the form element can grow or has a fixed size.")
