---
title: "Defining scrollgrid in the layout"
source: "fgl-topics/c_fgl_ui_scrollgrid_layout.html"
breadcrumb: "User interface > User interface programming > Scrollgrid views > Defining scrollgrid in the layout"
type: "concept"
---

# Defining scrollgrid in the layout

> Define scrollgrid views in the LAYOUT section of the form definition file.

## Designing scrollgrid views

The scrollgrid rows are defined as a template within an area delimited by curly brackets.

This is an example of a resizable scrollgrid definition using the `SCROLLGRID`
layout item. It is resizable because it has no fixed page size (`WANTFIXEDPAGESIZE=NO`).
Only one template is needed to define the records.

The position of the item tags is detected by the form compiler to build the scrollgrid. Item
types (widget to be used) and behavior are defined with form items in the
`ATTRIBUTES` section.

```
LAYOUT
SCROLLGRID (WANTFIXEDPAGESIZE=NO)
{
 Id:   [f1    ]
 Name: [f2           ][f3           ]
}
END
END
...
ATTRIBUTES
EDIT f1 = customer.cust_num;
EDIT f2 = customer.cust_fname, 
EDIT f3 = customer.cust_lname;
...
```

You can also define `SCROLLGRID` layout tags inside a `GRID` container,
beside other layout tags:

```
LAYOUT
GRID
{
<SCROLLGRID sg1                      >
 Id:   [f1    ]
 Name: [f2           ][f3           ]
...
}
END
END

ATTRIBUTES
SCROLLGRID sg1: scrollgrid1,
  WANTFIXEDPAGESIZE=NO,
  GRIDCHILDRENINPARENT;
...
```

> **Important:**
>
> Avoid Tab characters (ASCII 9) inside the curly-brace delimited area.
> If used, Tab characters are replaced with 8 blanks at compilation with
> fglform.

## Height of scrollgrid rows

The height of scrollgrid rows can be defined by adding empty tags underneath (this makes sense
only when using widgets that can get a height such as [`TEXTEDIT`](1747-textedit-item-definition.md "Defines attributes for a multi-line edit field.") or [`IMAGE`](1739-image-item-definition.md "Defines attributes for an area that can display an image resource.")).

In this example, the height of the image item type is defined by the number of rows in the layout
section (6 in our example).

```
LAYOUT
SCROLLGRID (WANTFIXEDPAGESIZE=NO)
{
 Icon [name        ] Index[index ]
 [img                            ]
 [                               ]
 [                               ]
 [                               ]
 [                               ]
 [                               ]
}
END
END
ATTRIBUTES
EDIT name = FORMONLY.name;
IMAGE img = FORMONLY.img, autoscale;
EDIT index = FORMONLY.index;
END
...
```

## Defining the rendering type of the scrollgrid

The scrollgrid can be configured to be fixed page size or resizable, depending on requirements or
the supported platform:

- To define a scrollgrid with a fixed page size, omit the `WANTFIXEDPAGESIZE=NO`
  attribute. The number of visible rows is defined by the number of row templates in the form
  layout.
- To define a resizable scrollgrid, the `WANTFIXEDPAGESIZE=NO` is set in the layout
  to allow the container to stretch vertically. Only one record row template needs to be defined. Then
  define the initial number of scrollgrid lines with the `INITIALPAGESIZE` form
  definition attribute.
- To render a scrollgrid as a paged responsive tile list, the `customWidget`
  presentation style is defined and the `WANTFIXEDPAGESIZE=NO` scrollgrid attribute is
  used.

For more details, see [Controlling scrollgrid rendering](2339-controlling-scrollgrid-rendering.md "Scrollgrid rendering can be controlled by the use of presentation styles and scrollgrid attributes.").

## Related links

**Related concepts**  

[SCROLLGRID container](1723-scrollgrid-container.md "Defines a scrollable grid view widget.")
