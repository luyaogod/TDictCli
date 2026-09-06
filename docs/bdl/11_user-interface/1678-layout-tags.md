---
title: "Layout tags"
source: "fgl-topics/c_fgl_FormSpecFiles_Layout_Tags.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file concepts > Form items > Form tags > Layout tags"
type: "concept"
---

# Layout tags

> Layout tags define layout areas for containers inside the frame of a grid-based container.

## Basics

A layout tag defines a layout region of a container, in the body frame of a [`GRID`](1722-grid-container.md "Defines a layout area based on a grid of cells.") container.

## Syntax

```
  <type  [identifier]      >
    content
[ <                        > ]
```

1. type defines the kind of layout tag to be inserted at this position.
2. identifier references a form item definition in the
   `ATTRIBUTES` section, it must be unique, but is optional.
3. content defines other form items inside the layout tag.
4. The (`< >`) ending the layout tag body is optional.

## Usage

While complex layout with nested frames can be defined with `HBOX` and
`VBOX` containers, it is sometimes more convenient to define a form with a
complex layout by using layout tags within a `GRID` container.

A layout tag has a type that defines the kind of container generated in the compiled
form.

A layout tag is delimited by angle brackets (`<>`), and contains the
tag type (`G`/`GROUP`, `T`/`TABLE`, etc)
and an optional identifier.

| Tag Type | Abbr. | Description |
| --- | --- | --- |
| `GROUP` | `G` | Defines a group box layout tag, resulting in the same presentation as the [`GROUP`](1719-group-container.md "Defines a layout area to group other layout elements together.") container. |
| `TABLE` | `T` | Defines a table view layout tag, resulting in the same presentation as the [`TABLE`](1724-table-container.md "Defines a re-sizable table designed to display a list of records.") container. |
| `TREE` | `N/A` | Defines a tree-view list view layout tag, resulting in the same presentation as the [`TREE`](1725-tree-container.md "The TREE container defines the presentation of a list of ordered records in a tree-view widget.") container. |
| `SCROLLGRID` | `S` | Defines a scrollable grid layout tag, resulting in the same presentation as the [`SCROLLGRID`](1723-scrollgrid-container.md "Defines a scrollable grid view widget.") container. |

The details of the layout tag definition are specified in the `ATTRIBUTE` section. Layout tags
must be identified by an item tag name. In the example, the layout tag named "g1" is defined in the
`ATTRIBUTE` section with the `GROUP` form item type to set the name
and text of the
element:

```
LAYOUT
GRID
{
<GROUP g1        >
[text1           ]
[                ]
[                ]
<                >
}
END
END
ATTRIBUTES
GROUP g1:group1, TEXT="Description";
TEXTEDIT text1=FORMONLY.text1;
END
```

In the `ATTRIBUTES` section, the group element definition is bound to the layout
item tag with the name "g1". The AUI element name will be "group1". This name can then be used in
programs to manipulate the group element, for example with [`ui.Form.setElementHidden("group1",1)`](../15_library-reference/3156-ui-form-setelementhidden.md "Show or hide form elements.").

The layout region is a rectangle, in which the width is defined by the length of the layout
tag, and the height by a closing tag (`< >`) . In the example, the layout
region is defined by the layout tag named
"group1".

```
<GROUP group1            >
                        
                         
<                        >
```

Form items must be placed inside the layout region. The `[ ]` square
brackets are not part of the form item's width and can be place at the same X position as the layout
tag delimiters:

```
<GROUP group1             >
 Item:     [f001          ]
 Quantity: [f002    ]    
 Date:     [f003          ]
<                         >
```

The `[ ]` delimiters are not counted to define the width of an item
tag. The width of the item is defined by the number of character between the square brackets. Thus,
this layout is valid and can be compiled:

```
<GROUP group1            >
[f001                    ]
[f002                    ]
 Static labels must fit!! 
<                        >
<TABLE table1            >
[colA  |colB             ]
[colA  |colB             ]
[colA  |colB             ]
[colA  |colB             ]
```

You can place several layout tags on the same layout line in order to split the frame
horizontally. This example defines six layout regions (four group boxes and two
tables):

```
<GROUP group1      ><GROUP group2               ><GROUP group4 >
 FName: [f001      ] Phone: [f004               ][f012         ]
 LName: [f002      ] EMail: [f005               ][             ]
<                  ><                           >[             ]
<GROUP group3                                   >[             ]
[f010                                           ][             ]
<                                               ><             >
<TABLE table1               ><TABLE table2                     >
[c11   |c12   |c13          ][c21   |c22                       ]
[c11   |c12   |c13          ][c21   |c22                       ]
[c11   |c12   |c13          ][c21   |c22                       ]
[c11   |c12   |c13          ][c21   |c22                       ]
<                           ><                                 >
```

The `< >` closing layout tag is optional. When not specified, the end
of the layout region is defined by the underlying layout tag or by the end of the current grid.
However, the ending tag must be specified if the form compiler cannot detect the end of the layout
region. This is usually the case with group layout tags.

In the example, the table does not need an ending layout tag because it is defined by the
starting tag of the group, but the group needs an ending tag otherwise it would include the last
field (`field3`). Additionally, if `field3` had a different size, the
form compiler would raise an error because the group and the last field geometry would
conflict.

```
<TABLE table1     >
[colA  |colB      ]
[colA  |colB      ]
[colA  |colB      ]
[colA  |colB      ]
[colA  |colB      ]
[colA  |colB      ]
<GROUP group2     >
[field1           ]
[field2           ]
<                 >
[field3           ]
```

It is possible to mix container layout tags with singular form items. You typically put
form items using a large area of the form, such as `IMAGE` fields or
`TEXTEDIT` fields. The `[ ]` delimiters are not used to compute the
size of the singular form items:

```
<GROUP group1            >[image1         ]
 FName: [f001            ][               ]
 LName: [f002            ][               ]
<                        >[               ]
[textedit1                |               ]
[                         |               ]
[                         |               ]
```

Table layout tags can be embedded inside group layout
tags:

```
<GROUP group1           >
 <TABLE table1         >
 [colA  |colB          ]
 [colA  |colB          ] 
 [colA  |colB          ] 
 [colA  |colB          ]
<                       >
```

Hbox or vbox containers with splitter are [automatically created](1554-automatic-hbox-vbox-with-splitter.md "Horizontal and vertical boxes with splitter are created automatically when stretchable elements are set side by side.") by the form compiler in
these conditions:

- Hbox is created when two or more stretchable elements are stacked side by side and touch
  each other (no space between).
- Vbox is created when two or more stretchable elements are stacked vertically and touch
  each other (no space between).

Stretchable elements are containers such as `TABLE` containers, or form
items like `IMAGE` fields with the [`STRETCH`](1821-stretch-attribute.md "The STRETCH attribute defines if the form element can grow or has a fixed size.") attribute.

No hbox or vbox object will be created if the elements are in a `SCROLLGRID`
container.

This example defines two tables stacked vertically, generating a VBox with splitter (note
that ending tags are
omitted):

```
<TABLE table1     >
[colA  |colB      ]
[colA  |colB      ]
[colA  |colB      ]
[colA  |colB      ]
<TABLE table2     >
[colC  |colD      ]
[colC  |colD      ]
```

In this example, the layout defines two stretchable `TEXTEDIT` fields placed
side by side which would generate an automatic hbox with splitter. To make both textedits touch you
need to use a pipe delimiter in between:

```
[textedit1         |textedit2                ]
[                  |                         ]
[                  |                         ]
[                  |                         ]
```

The next layout example would make the
form compiler create an automatic vbox with splitter to hold `table2` and
`textedit1`, plus an hbox with splitter to hold `table1` and the first
VBox (We must use a pipe character to delimit the end of `colB` and
`textedit1` so that both tables can be placed side by
side):

```
<TABLE table1     ><TABLE table2           >
[colA  |colB      ][colC|colD              ]
[colA  |colB      ][colC|colD              ]
[colA  |colB      ][colC|colD              ]
[colA  |colB      |textedit1                ]
[colA  |colB      |                         ]
[colA  |colB      |                         ]
```

If you want to avoid automatic hbox or vbox with splitter creation, you must add blanks
between
elements:

```
<TABLE table1     >  <TABLE table2           >
[colA  |colB      ]  [colC|colD              ]
[colA  |colB      ]  [colC|colD              ]
[colA  |colB      ]  [colC|colD              ]
[colA  |colB      ]
[colA  |colB      ] [textedit1                ]
[colA  |colB      ] [                         ]
[colA  |colB      ] [                         ]
```

## Examples

The typical OK/Cancel window:

```
LAYOUT
GRID
{
<GROUP g1                   >
[com                        ]
<                           >
[            :bok   |bno    ]
}
END
END
ATTRIBUTES
LABEL com: comment;
BUTTON bok: accept;
BUTTON bno: cancel;
...
```

This example shows multiple uses of layout tags:

```
LAYOUT
GRID
{
<GROUP g1                          ><GROUP g2         >
 Ident: [f001   ] [f002            ] [text1           ]
 Addr:  [f003                      ] [                ]
<                                  ><                 >
<GROUP g3                                             >
[text2                                                ]
[                                                     ]
[                                                     ]
<                                                     >
<TABLE t1                                             >
  Num      Name               State  Value 
[col1    |col2              |col3  |col4              ]
[col1    |col2              |col3  |col4              ]
[col1    |col2              |col3  |col4              ]
[col1    |col2              |col3  |col4              ]
<                                                     >
}
END
END
ATTRIBUTES
GROUP g1:group1, TEXT="Customer";
GROUP g2:group2, TEXT="Comments";
TABLE t1:table1, UNSORTABLECOLUMNS;
...
```

## Related links

**Related concepts**  

[Form rendering](1538-form-rendering.md "The section explains the layout rules to render forms on graphical front-ends.")

[Grid-based layout](1549-grid-based-layout.md "A form file can define a grid-based layout within a tree of layout items.")
