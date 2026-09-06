---
title: "Hbox tags"
source: "fgl-topics/c_fgl_FormSpecFiles_HBox_Tags.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file concepts > Form items > Form tags > Hbox tags"
type: "concept"
---

# Hbox tags

> Hbox tags group several item tags within the same horizontal layout box, inside a grid-based container (GRID).

## Basics

An Hbox tag defines the position and size in a [`GRID`](1692-grid-item-type.md "Defines a layout area based on a grid of cells.") container for an horizontal box
containing several leaf form items.

The elements in the Hbox tag can use additional alignment rules to get the required visual
affect.

## Syntax

> **Note:**
>
> In the next(s) syntax diagram(s), the `[ ] { } |`
> symbols are part of the syntax.

```
[ element : ... ]
```

where element can be:

- An identifier referecing a form item definition in the
  `ATTRIBUTES`
  section:

  ```
  identifier
  ```

  The hyphen-minus
  can be placed after the identifier, to specify the real width of the
  element:

  ```
  identifier   -
  ```
- A string-literal:

  ```
  "character ..."
  ```
- A spacer:

  ```
  blank-char ...
  ```

1. The `:` colon is the delimiter for Hbox tag elements.
2. string-literal specifies a quoted text that defines a static label.
3. character is any character of the current encoding.
4. spacer is one or more blank characters specifying an invisible element that
   expands automatically.

## Usage

Hbox tags are provided to control the alignment of form items in a [grid](1549-grid-based-layout.md "A form file can define a grid-based layout within a tree of layout items."). Hbox tags allow you to stack form items
horizontally without the elements being influenced by elements above or below. In an Hbox, you can
mix form fields, static labels and spacers. A typical use of the hbox is to have zip-code / city
form fields side by side with predictable spacing in between.

An hbox tag is delimited by square brackets (`[]`) and must contain at least one
string-list or an identifier preceded or followed by a colon
(`:`). A string-list is a combination of
string-literals (quoted text) and spacers (blank characters).
The delimiter for hbox tag elements is the colon.

Hbox tags are not allowed for fields of
[Screen records / arrays](1676-screen-records-arrays.md "Form fields can be grouped in a screen record or screen array definition."); a form compiler error is generated.
The client needs a matrix element directly in a grid or a scrollgrid to perform the necessary
positioning calculations for the individual fields.

The following example shows simple hbox tags:

```
GRID
{
 ["Customer info:":    ]
 [f001       :         ]
 [  :f002              ]
 ["Name: "  :f003      ]
}
END
```

In this example:

1. The first hbox tag contains two elements: a static label and a spacer.
2. The second hbox tag contains two elements: a form field and a spacer.
3. The third hbox tag contains two elements: a spacer and a form field.
4. The fourth hbox tag contains two elements: a static label and a form field.

An hbox tag defines the position and width (in grid cells) of several form items grouped inside
an horizontal box. The position and width (in grid cells) of the horizontal box is defined by the
square brackets (`[]`) delimiting the hbox tag.

When using an identifier, you define the position of a form item which is
described in the [`ATTRIBUTES` section](1727-attributes-section.md "The ATTRIBUTES section describes properties of elements used in the form."). When using a string-list, you can
define static labels and/or spacers. The following example defines an hbox tag generating 7 items (a
static label, a spacer, a form item identified by `num`, a spacer, a static label, a
spacer and a form item identified by `name`):

```
GRID
{
 ["Num:"  :num  :  :"Name:"  :name        ]
}
END
```

A spacer is an invisible element that automatically expands. It can be used to align
elements left, right or center in the hbox. The following example defines 3 hboxes with the same
width. Each hbox contains one field. The first field is aligned to the left, the second is aligned
to the right and third is centered:

```
GRID
{
 [left :              ]
 [         :right     ]
 [     :centered:     ]
}
END

ATTRIBUTES
 LABEL left: label1, TEXT="LEFT";
 LABEL right: label2, TEXT="RIGHT";
 LABEL centered: label3, TEXT="CENTER";
END
```

When you use string literals, the quotes define where the label starts and stops. If there is
free space after the quote that ends the label, then it is filled by a spacer. Consider this
example:

```
GRID
{
 [    :"Label1"     ]
 [         :"Label2"]
}
END
```

In this example:

1. The first line contains a spacer, followed by the static label, followed by another spacer.
   The quotation marks end the string literal; a colon is not required to delimit the label from
   the final spacer.
2. The second line contains a spacer, followed by the static label. Because there is no empty
   space between the end of the static label and the closing bracket of the hbox tag
   ( `]` ).

A typical use of hbox tags is to vertically align some form items - that must appear on the
same line - with one or more form items that appear on the other lines:

```
GRID
{
 Id:      [num  :"Name: ":name        ]
 Address: [street                   : ]
          [zip-code:city              ]
 Phones:  [phone       :fax           ]
}
END
```

In this example, the form compiler will generate a grid containing 7 elements (3 labels and
4 hboxes):

1. The label "Id:",
2. A first hbox defines 3 cells, where:
   - The field 'num' will occupy the cell (1,1),
   - The label "Name:" will occupy the cell (2,1),
   - The field 'name' will occupy the cell (3,1).
3. The label "Address:" will occupy cell (1,2),
4. A second hbox defines 1 cell, where:
   - The field 'street' will occupy the cell (1,1).
5. A third hbox defines 2 cells, where:
   - The field 'zip-code' will occupy the cell (1,1),
   - The field 'city' will occupy the cell (2,1).
6. The label "Phones:" will occupy cell (1,4),
7. A fourth hbox defines 2 cells, where:
   - The field 'phone' will occupy the cell (1,1),
   - The field 'fax' will occupy the cell (2,1).

Inside an hbox tag, the positions and widths of elements are independent of other hboxes. It is
not possible to align elements over hboxes. The position of items inside an hbox depends on the
spacer and the real size of the elements. The following example does not align the items as you
would expect, following the character positions in the layout definition:

```
GRID
{
  ["Num:     " :fnum :        ]
  ["Name:    " :fname         ]
}
END
```

A big advantage in using elements in an hbox is that the fields get their real sizes from the
.per definition. The following example illustrates the case:

```
GRID
{
 MMMMM
[f1   ]
[f2 : ]
}
END
```

Here all items will occupy the same number of grid columns (5). The MMMMM static label will have
the largest width and define the width of the 5 grid cells. The first field is defined with a normal
item tag, and expands to the width of the 5 grid cells. The line 5 defines an hbox that will expand
to the size of the 5 grid cells, according to the static label, but its child element - the field f2
- gets a size corresponding to the number of characters used before the ':' colon (that is 3
characters).

If the default width generated by the form compiler does not fit, the `-` hyphen
symbol can be used to define the real width of the item. In this example, the hbox tag occupies 20
grid cells, the first form item gets a width of 5, and the second form item gets a width of 3:

```
GRID
{
 12345678901234567890
[f1   - :f2 -    :   ]
}
END
```

The `-` hyphen size indicator is especially useful in `BUTTONEDIT`,
`DATEEDIT`, and `COMBOBOX` form fields, for which the default width
computed by the form compiler may not fit. See [Widget width inside hbox tags](1559-widget-width-inside-hbox-tags.md) for more details.

In this example, a static label is positioned above a `TEXTEDIT` field. The label
will be centered over the `TEXTEDIT` field, and will remain centered as the field
expands or contracts with the resizing of the window.

```
GRID
{
 [ :"label": ]
 [textedit   ]
}
END

ATTRIBUTES
 TEXTEDIT textedit = formonly.textedit, STRETCH=BOTH;
END
```

## Related links

**Related concepts**  

[Grid-based layout](1549-grid-based-layout.md "A form file can define a grid-based layout within a tree of layout items.")

[Layout tags](1678-layout-tags.md "Layout tags define layout areas for containers inside the frame of a grid-based container.")
