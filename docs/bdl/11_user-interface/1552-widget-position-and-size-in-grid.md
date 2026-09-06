---
title: "Widget position and size in grid"
source: "fgl-topics/c_fgl_form_rendering_grid_pos_size.html"
breadcrumb: "User interface > Form definitions > Form rendering > Grid-based layout > Widget position and size in grid"
type: "concept"
---

# Widget position and size in grid

> Form items render as widgets in the window, at a given position and with a given size.

To render form items, grid-based rendering follows the layout rules described below:

1. The position of the widgets in the virtual grid is defined by the
   `posX` and `posY` AUI tree attributes.
2. The number of virtual grid cells occupied by a widget is defined by the
   `gridWidth` and `gridHeight` AUI tree attributes.
3. The real size of a widget is defined by the `width` and `height`
   AUI tree attributes. Some form items allow to specify their size with the [`WIDTH`](1850-width-attribute.md "The WIDTH attribute forces an explicit width of a form element.") and [`HEIGHT`](1782-height-attribute.md "The HEIGHT attribute forces an explicit height for a form element.") attributes.
4. Empty lines and empty columns in the form layout definition take a size of 0 pixels.
5. The size of a cell in the virtual grid depends on the real size of the widgets
   inside the grid.
6. A widget's minimum size is computed via its real size and the [`SAMPLE`](1814-sample-attribute.md "The SAMPLE attribute defines the text to be used to compute the width of a form field widget.") attribute. The minimum width
   can be forced by setting the [`STRETCHMIN`](1824-stretchmin-attribute.md "The STRETCHMIN attribute defines the minimum stretching width.") attribute.
7. The preferred size of the widget is computed following the [`SIZEPOLICY`](1817-sizepolicy-attribute.md "The SIZEPOLICY attribute is a sizing directive based on the content of a form item.") attribute.
8. The final widget size is computed depending on the minimum and preferred size, to fill the cells
   in the grid.
9. A small spacing is applied in non-empty cells.

Take this form file:

```
LAYOUT
GRID
{
First Name  [fn     ]
Last Name   [ln     ]
}
END
END
ATTRIBUTES
EDIT fn = FORMONLY.fname;
EDIT ln = FORMONLY.lname;
END
```

It will produce the virtual grid as follows:

![Grid layout rules](../_images/layout03.jpg)

*Two labels and two fields placed in a grid: Grid view*

By default, empty grid rows and empty grid columns get no size when rendered on the front-end.
For example, in the above grid sample, the grid columns #10 and #11 are empty.

Visual result:

![Grid layout rules in Genero form](../_images/layout04_gbc.jpg)

*Two labels and two fields placed in a grid: Form view*
