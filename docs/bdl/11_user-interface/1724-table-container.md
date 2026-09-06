---
title: "TABLE container"
source: "fgl-topics/c_fgl_FormSpecFiles_TABLE_container.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file structure > LAYOUT section > TABLE container"
type: "concept"
---

# TABLE container

> Defines a re-sizable table designed to display a list of records.

## Syntax

```
TABLE [identifier] [ ( attribute [,...] ) ]
{
 title [...]
[col-name   [|...]     ]
[...]
[aggr-name  [|...]     ]
}  
END
```

1. identifier defines the name of the element, it is optional but
   recommended.
2. attribute is an attribute for the element.
3. title is the text to be displayed as column title.
4. col-name is an identifier that references a form field.
5. aggr-name is an identifier that references an aggregate Field.

## Form attributes

`AGGREGATETEXT`,
`COMMENT`, `DOUBLECLICK`, [`FLIPPED`](1780-flipped-attribute.md "The FLIPPED attribute flips TABLE columns into rows."), `HIDDEN`, `FONTPITCH`, [`STRETCHCOLUMNS`](1822-stretchcolumns-attribute.md "The STRETCHCOLUMNS attribute makes all TABLE/TREE columns stretchable."), `STYLE`, `TAG`, `UNHIDABLECOLUMNS`, `UNMOVABLECOLUMNS`, `UNSIZABLECOLUMNS`, `UNSORTABLECOLUMNS`, `WANTFIXEDPAGESIZE`, `WIDTH`, `HEIGHT`.

## Style attributes

Common: [`backgroundColor`](1629-style-attributes-common-to-all-elements.md), [`border`](1629-style-attributes-common-to-all-elements.md), [`fontFamily`](1629-style-attributes-common-to-all-elements.md), [`fontSize`](1629-style-attributes-common-to-all-elements.md), [`fontStyle`](1629-style-attributes-common-to-all-elements.md), [`fontWeight`](1629-style-attributes-common-to-all-elements.md), [`textColor`](1629-style-attributes-common-to-all-elements.md), [`textDecoration`](1629-style-attributes-common-to-all-elements.md).

Class-specific: [`allowWebSelection`](1648-table-style-attributes.md), [`alternateRows`](1648-table-style-attributes.md), [`forceDefaultSettings`](1648-table-style-attributes.md), [`headerAlignment`](1648-table-style-attributes.md), [`headerHidden`](1648-table-style-attributes.md), [`headerPosition`](1648-table-style-attributes.md), [`highlightColor`](1648-table-style-attributes.md), [`highlightCurrentCell`](1648-table-style-attributes.md), [`highlightCurrentRow`](1648-table-style-attributes.md), [`highlightTextColor`](1648-table-style-attributes.md), [`leftFrozenColumns`](1648-table-style-attributes.md), [`reduceFilter`](1648-table-style-attributes.md), [`resizeFillsEmptySpace`](1648-table-style-attributes.md), [`rightFrozenColumns`](1648-table-style-attributes.md), [`rowActionTrigger`](1648-table-style-attributes.md), [`rowAspect`](1648-table-style-attributes.md), [`rowHover`](1648-table-style-attributes.md), [`showGrid`](1648-table-style-attributes.md), [`tableType`](1648-table-style-attributes.md).

## Usage

The `TABLE` container defines a list view element in a grid-based layout.

To create a table view in a grid layout, define the following elements in the form file:

1. The layout of the list, with a `TABLE` container in the
   `LAYOUT` section.
2. The column data types and field properties, in the `ATTRIBUTES`
   section.
3. The field list definition to group form fields together with a screen array, in the
   `INSTRUCTIONS` section.

For more details about this item type, see [TABLE item type](1703-table-item-type.md "Defines a list view widget.").

## Example

```
SCHEMA custdemo
LAYOUT ( TEXT="Customer list" )
TABLE ( TAG="normal" )
{
[c1     |c2                        |c3        |c4 ]
[c1     |c2                        |c3        |c4 ]
[c1     |c2                        |c3        |c4 ]
[c1     |c2                        |c3        |c4 ]
}
END
END
TABLES
  customer
END
ATTRIBUTES
  EDIT c1 = customer.cust_num, TITLE="Num";
  EDIT c2 = customer.cust_name, TITLE="Customer name";
  EDIT c3 = customer.phone, TITLE="Phone";
  EDIT c4 = customer.city, TITLE="City";
END
INSTRUCTIONS
SCREEN RECORD custlist(
    customer.cust_num,
    customer.cust_name,
    customer.phone,
    customer.city
);
END
```

## Related links

**Related concepts**  

[Grid-based layout](1549-grid-based-layout.md "A form file can define a grid-based layout within a tree of layout items.")

[Layout tags](1678-layout-tags.md "Layout tags define layout areas for containers inside the frame of a grid-based container.")

[Item tags](1679-item-tags.md "Item tags define the position and size in a grid-based container.")
