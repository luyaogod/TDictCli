---
title: "Horizontal stretching"
source: "fgl-topics/c_fgl_form_responsive_stretch.html"
breadcrumb: "User interface > Form definitions > Form rendering > Responsive Layout > Horizontal stretching"
type: "concept"
---

# Horizontal stretching

> Define stretchable form elements to achieve responsive layout.

## Purpose of horizontal stretching

For best user experience, form elements must stretch horizontally, to fill the whole screen of a
mobile device (especially when switching between portrait and landscape mode). Horizontal stretching
is also useful, when enlarging the window of a desktop application or the browser window using a web
application.

To get horizontal stretching for a set of fields aligned vertically, define the
`STRETCH=X` attribute for one of the fields of the set. When the display area is
enlarged, the group of fields will then stretch proportionally to each other, following the [grid-layout](1549-grid-based-layout.md "A form file can define a grid-based layout within a tree of layout items.") alignment rules.

In the next example, the `cust_id`, `cust_name` and
`cust_company` fields belong to the same vertical alignment group, since they all
start at the same position.

By defining the [`STRETCH=X`](1821-stretch-attribute.md "The STRETCH attribute defines if the form element can grow or has a fixed size.")
attribute for the `cust_id` field, you indicate that this field set can stretch
horizontally. Since cust\_name and cust\_company will stretch as well because of grid layout rules, it
is recommended to use the [`SCROLL`](1815-scroll-attribute.md "The SCROLL attribute can be used to enable horizontal scrolling in a character field.") attribute for character string fields:

```
LAYOUT
GRID
{
ID:      [f01     ]       Transaction id: [f10    ]
Name:    [f02          ]  Date:           [f11       ]
Company: [f03          ]
}
END
END
ATTRIBUTES
EDIT f01 = customer.cust_id, STRETCH=X;
EDIT f02 = customer.cust_name, SCROLL;
EDIT f03 = customer.cust_company, SCROLL;
...
END
```

## STRETCH attribute on LAYOUT element

A global `STRETCH` attribute can be defined on the `LAYOUT`
element.

All form items that accept the `STRETCH` attribute and do not have this
attribute defined locally, will inherit from the stretch value defined at the `LAYOUT`
level.

```
LAYOUT ( STRETCH=X )
...
```

On desktop, for the `LAYOUT` element, the `STRETCH` attribute is
`NONE` by default. On mobile, the default is `STRETCH=X`. To override
the default on mobile, define `STRETCH=NONE` on the `LAYOUT` element,
and use the `STRETCH=X` at form item
level:

```
-- Mobile form bypassing default STRETCH=X for LAYOUT
LAYOUT ( STRETCH=NONE )
VBOX
GRID
{
[f1    ]
...
ATTRIBUTES
EDIT f1 = customer.cust_num, STRETCH=X, ...
```

`GRID` and `SCROLLGRID` elements using the
`STRETCH=X` attribute can also define a [`STRETCHMIN`](1824-stretchmin-attribute.md "The STRETCHMIN attribute defines the minimum stretching width.") value, to overwrite
the default minimal widget width defined by the number of cells used in the form item
tag:

```
EDIT f1 = customer.cust_num, STRETCH=X, STRETCHMIN=5;
```

## TABLE and TREE container stretching

`TABLE` and `TREE` containers can by default stretch vertically and
horizontally.

To force a table or tree view container to keep its original width or height (or both), use the
`STRETCH` attribute in the container definition.

For example, with small screens, a table will only stretch vertically and keep its original
width, by using the following definition:

```
LAYOUT
TABLE ( STRETCH@SMALL=Y )
...
```

See also [c\_fgl\_ui\_tables\_rendering.html#c\_fgl\_ui\_tables\_rendering\_\_section\_table\_stretch](2319-controlling-table-rendering.md).

## TABLE and TREE stretchable columns

`TABLE` and `TREE` containers can use columns that can stretch or
shrink horizontally, when the width of the window changes.

To get stretchable columns, form fields defining table columns must use the
`STRETCH=X` attribute, or the `TABLE/TREE` container can define the
`STRETCHCOLUMNS` attribute to make all columns stretchable. Each column can define a
minimum and maximum stretch width by using the `STRETCHMIN` and
`STRETCHMAX` attributes:

```
LAYOUT
TABLE table1
{
[c1    |c2          |c3         ...
...
ATTRIBUTES
EDIT c1 = customer.cust_num;
EDIT c2 = customer.cust_name, STRETCH=X, STRETCHMAX=40, ...
```

`STRETCH`, `STRETCHCOLUMNS`, `STRETCHMIN` and
`STRETCHMAX` can be combined with a `@screen-size`
modifier, to enable these stretching attributes for specific screen sizes.

```
TABLE table1 ( STRETCHCOLUMNS@LARGE )
```

For more details, see [Stretchable columns](2319-controlling-table-rendering.md).

## Example with multiple grids and table

The next form file example defines:

- An `HBOX` containing a `VBOX` and a `GRID`
  (`grid2`).
- The inner `VBOX` contains another `GRID` (`grid1`)
  and a `TABLE`.
- The first field (`field11`) in `grid1` gets a
  `STRETCH=X` attribute, to allow horizontal stretching for the whole
  `GRID`.
- The `TEXTEDIT` `field21` gets a `STRETCH=Y`
  attribute, to allow vertical stretching of that field (we do not want that this field stretches
  horizontally).
- The `TABLE` can stretch vertically and horizontally by design, but the columns do
  not stretch (horizontally) by default.
- The second table column (`col22`) uses a `STRETCH=X` attribute to
  make it stretch, when the parent table stretches, and `STRETCHMAX=50` indicates that
  this column cannot stretch more than 50 grid units.

With this layout definition, the form elements will occupy all the space of the window:

```
LAYOUT
HBOX
VBOX
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
GRID grid2
{
 Field21
[f21         ]
[            ]
[            ]
[            ]
 Field22
[f22     ]
 Field23
[f23         ]
}
END
END

ATTRIBUTES

EDIT f11 = FORMONLY.field11, STRETCH=X;
COMBOBOX f12 = FORMONLY.field12, ITEMS=("aaa","bbb");
DATEEDIT f13 = FORMONLY.field13;

TEXTEDIT f21 = FORMONLY.field21, STRETCH=Y;
COMBOBOX f22 = FORMONLY.field22, ITEMS=("aaa","bbb");
DATEEDIT f23 = FORMONLY.field23;

EDIT c21 = FORMONLY.col21;
EDIT c22 = FORMONLY.col22, STRETCH=X, STRETCHMAX=50;
DATEEDIT c23 = FORMONLY.col23;

END

INSTRUCTIONS
SCREEN RECORD sr1(FORMONLY.col21 THRU FORMONLY.col23);
END
```

Visual result with a small window:

![Screenshot of form with stretcheable fields in small window](../_images/rl_stretch_gbc_1.jpg)

*Form with stretcheable fields in small window*

Visual result with a larger window:

![Screenshot of form with stretcheable fields in larger window](../_images/rl_stretch_gbc_2.jpg)

*Form with stretcheable fields in larger window*
